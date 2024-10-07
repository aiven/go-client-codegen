//go:build generator

package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

const docSite = "https://api.aiven.io/doc"

// Doc represents a parsed OpenAPI document.
type Doc struct {
	Paths      map[string]map[string]*Path `json:"paths"`
	Components struct {
		Schemas    map[string]*Schema    `json:"schemas"`
		Parameters map[string]*Parameter `json:"parameters"`
	} `json:"components"`
}

const componentIndex = 3

func (d *Doc) getSchema(path string) (*Schema, error) {
	chunks := strings.Split(path, "/")
	if len(chunks) < componentIndex+1 {
		return nil, fmt.Errorf("invalid schema path %q", path)
	}

	name := chunks[componentIndex]
	schema := d.Components.Schemas[name]
	for i := componentIndex + 1; i < len(chunks); i++ {
		switch k := chunks[i]; k {
		case "items":
			schema = schema.Items
		case "properties":
			schema = schema.Properties[chunks[i+1]]
			i++
		default:
			return nil, fmt.Errorf("unknown schema path %v: %s", chunks, k)
		}
	}

	if schema == nil {
		return nil, fmt.Errorf("schema %q not found", path)
	}

	schema.name = name

	return schema, nil
}

// Content represents a request or response body.
type Content map[string]*struct {
	Schema struct {
		Ref string `json:"$ref"`
	} `json:"schema"`
}

// Path represents a parsed OpenAPI path.
type Path struct {
	ID          string
	Path        string
	Method      string
	FuncName    string
	Tags        []string     `json:"tags"`
	OperationID string       `json:"operationId"`
	Parameters  []*Parameter `json:"parameters"`
	Summary     string       `json:"summary"`
	Deprecated  bool         `json:"deprecated"`
	In          struct {
		Content Content `json:"content"`
	} `json:"requestBody"`
	Out struct {
		OK struct {
			Content Content `json:"content"`
		} `json:"200"`
		NoContent struct {
			Content Content `json:"content"`
		} `json:"204"`
	} `json:"responses"`
}

// Comment returns a comment for the path.
func (p *Path) Comment() *jen.Statement {
	// IDE highlights any coincidence method names
	// For instance, there is always "List" method and that's a common verb
	// But IDE will highlight it as a reference
	// Lowers first letter
	s := fmt.Sprintf("%s %s", p.FuncName, lowerFirst(p.Summary))
	if p.Deprecated {
		s = "Deprecated: " + s
	}

	c := jen.Comment(s)
	c.Line().Comment(fmt.Sprintf("%s %s", p.Method, p.Path))

	if p.Tags[0] == "" {
		c.Line().Comment(fmt.Sprintf("%s/#operation/%s", docSite, p.OperationID))
	} else {
		c.Line().Comment(fmt.Sprintf("%s/#tag/%s/operation/%s", docSite, p.Tags[0], p.OperationID))
	}

	return c
}

// ParameterIn represents a parameter location.
type ParameterIn string

const (
	// ParameterInPath represents a path parameter location.
	ParameterInPath ParameterIn = "path"
)

// Parameter represents a parsed OpenAPI parameter.
type Parameter struct {
	Ref         string      `json:"$ref"`
	In          ParameterIn `json:"in"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Schema      *Schema     `json:"schema"`
}

// SchemaType represents a schema type.
type SchemaType string

const (
	// SchemaTypeObject represents an object schema type.
	SchemaTypeObject = "object"
	// SchemaTypeArray represents an array schema type.
	SchemaTypeArray = "array"
	// SchemaTypeString represents a string schema type.
	SchemaTypeString = "string"
	// SchemaTypeInteger represents an integer schema type.
	SchemaTypeInteger = "integer"
	// SchemaTypeNumber represents a number schema type.
	SchemaTypeNumber = "number"
	// SchemaTypeBoolean represents a boolean schema type.
	SchemaTypeBoolean = "boolean"
	// SchemaTypeTime represents a time schema type.
	SchemaTypeTime = "time"
	// SchemaTypeAny internal type for anyOf
	SchemaTypeAny = "any"
)

// Schema represents a parsed OpenAPI schema.
type Schema struct {
	Type                 SchemaType         `json:"type"`
	Properties           map[string]*Schema `json:"properties"`
	AdditionalProperties *Schema            `json:"additionalProperties"`
	Items                *Schema            `json:"items"`
	RequiredProps        []string           `json:"required"`
	Enum                 []any              `json:"enum"`
	Default              any                `json:"default"`
	MinItems             int                `json:"minItems"`
	Ref                  string             `json:"$ref"`
	Description          string             `json:"description"`
	CamelName            string             `json:"for-hash-only!"`
	required             bool
	name                 string
	propertyNames        []string
	parent               *Schema
	in, out              bool      // Request or Response DTO
	hasCollision         bool      // Means this struct has a collision with another one with different type of fields
	duplicates           []*Schema // Refs to structs with exactly the same fields
}

//nolint:funlen,gocognit,gocyclo // It is easy to maintain and read, we don't need to split it
func (s *Schema) init(doc *Doc, scope map[string]*Schema, name string) {
	if s.Ref != "" {
		other, err := doc.getSchema(s.Ref)
		if err != nil {
			panic(err)
		}

		*s = *other
	}

	for _, k := range s.RequiredProps {
		p, ok := s.Properties[k]
		if ok {
			p.required = true
		}
	}

	// Removes fields that are managed by the client,
	// and converted to a golang err object.
	// We don't need to expose them in DTO
	if s.out {
		delete(s.Properties, "errors")
		delete(s.Properties, "message")
	}

	s.name = name
	s.CamelName = strcase.ToCamel(s.name)

	if s.isEnum() {
		const enumTypeSuffix = "Type"

		betterName := getEnumName(s)
		if betterName != s.name {
			s.CamelName = cleanEnumName.ReplaceAllString(strcase.ToCamel(betterName), "") + s.CamelName
		}

		if !strings.Contains(s.CamelName, enumTypeSuffix) {
			s.CamelName += enumTypeSuffix
		}
	}

	// Adds suffix to reduce name collision
	suffix := ""
	switch {
	case s.isIn():
		suffix = "In"
	case s.isOut():
		suffix = "Out"
	}

	if s.isObject() {
		s.CamelName += suffix
	}

	// Cleans duplicates like StatusStatus
	s.CamelName = dedupCamelName(s.CamelName)

	// Makes structure private
	if s.isPrivate() {
		s.CamelName = lowerFirst(s.CamelName)
	}

	if s.parent != nil && s.parent.isPrivate() {
		s.CamelName = strcase.ToCamel(s.parent.CamelName)
	}

	// Some cases just impossible to cover
	switch s.CamelName {
	case "MessageFormatVersionValueType":
		s.CamelName = "MessageFormatVersionType"
	case "ServiceKafkaConnectConnectorStatusStateType":
		s.CamelName = "ServiceKafkaConnectConnectorStateType"
	case "PeeringConnectionStateType", "VpcPeeringConnectionWithResourceGroupStateType",
		"VpcPeeringConnectionWithRegionStateType":
		s.CamelName = "VpcPeeringConnectionStateType"
	case "ServiceSchemaRegistryGlobalConfigGetOut", "ServiceSchemaRegistryGlobalConfigPutOut",
		"ServiceSchemaRegistrySubjectConfigGetOut", "ServiceSchemaRegistrySubjectConfigPutOut":
		if s.isEnum() {
			s.CamelName = "CompatibilityType"
		}
	}

	if s.Type == SchemaTypeString {
		parts := strings.Split(s.name, "_")
		suffx := parts[len(parts)-1]

		if len(parts) > 1 && (suffx == "at" || suffx == "time") {
			s.Type = SchemaTypeTime
		}
	}

	if s.isArray() {
		// a workaround for invalid schema
		// fixme: on the backend
		if s.Items == nil {
			s.Items = &Schema{Type: SchemaTypeAny}
		}

		s.Items.parent = s
		s.Items.required = true // a workaround to not have slices with pointers
		s.Items.init(doc, scope, toSingle(name))
	}

	if s.isObject() {
		s.propertyNames = sortedKeys(s.Properties)
		for _, k := range s.propertyNames {
			p := s.Properties[k]
			p.parent = s
			p.init(doc, scope, k)
		}
	}

	if s.isObject() || s.isEnum() { //nolint:nestif
		for s.parent != nil {
			v, ok := scope[s.CamelName]
			if !ok {
				break
			}

			if v.hash() == s.hash() {
				v.duplicates = append(v.duplicates, s)
				return
			}

			// Resolves name collision
			// Takes parent's name as prefix or uses parent's name
			parent := s.parent
			if s.parent.isArray() {
				parent = parent.parent
			}

			if parent.isPrivate() {
				s.CamelName = strcase.ToCamel(parent.CamelName)
			} else {
				s.CamelName = strings.TrimSuffix(strcase.ToCamel(parent.CamelName), suffix) + s.CamelName
			}

			// Marks all have collision
			// We don't know in the beginning that there will be a collision
			// That's why we need this "duplicates" field
			v.hasCollision = true
			for _, d := range v.duplicates {
				d.hasCollision = true
			}
			s.hasCollision = true
		}

		scope[s.CamelName] = s
	}
}

// isPrivate returns true when a struct is just a wrapper for one field,
// so we can just return the field value making things less nested
func (s *Schema) isPrivate() bool {
	return s.parent == nil && s.out && len(s.Properties) == 1
}

// hash is for comparison
func (s *Schema) hash() string {
	if s.isEnum() {
		// Compares enums by values
		return mustMarshal(s.Enum)
	}

	return mustMarshal(s)
}

func (s *Schema) isObject() bool {
	return s.Type == SchemaTypeObject && len(s.Properties) != 0
}

func (s *Schema) isArray() bool {
	return s.Type == SchemaTypeArray
}

// isMap schemaless map
func (s *Schema) isMap() bool {
	return s.Type == SchemaTypeObject && len(s.Properties) == 0
}

func (s *Schema) isEnum() bool {
	return len(s.Enum) != 0
}

func (s *Schema) root() *Schema {
	if s.parent == nil {
		return s
	}

	return s.parent.root()
}

// isIn is request object
func (s *Schema) isIn() bool {
	return s.root().in
}

// isOut is response object
func (s *Schema) isOut() bool {
	return s.root().out
}

func (s *Schema) lowerCamel() string {
	return strcase.ToLowerCamel(s.CamelName)
}

func (s *Schema) level() int {
	level := 0
	p := s.parent
	for p != nil {
		level++
		p = p.parent
	}
	return level
}

// isAnonymous returns true when a struct should be rendered anonymous to reduce scope noise
func (s *Schema) isAnonymous() bool {
	return s.hasCollision && s.isObject() && s.level() > 3
}

func getScalarType(s *Schema) *jen.Statement {
	switch s.Type {
	case SchemaTypeString:
		return jen.String()
	case SchemaTypeInteger:
		return jen.Int()
	case SchemaTypeNumber:
		return jen.Float64()
	case SchemaTypeBoolean:
		return jen.Bool()
	case SchemaTypeTime:
		return jen.Qual("time", "Time")
	default:
		panic(fmt.Errorf("unknown type %q", s.Type))
	}
}

// getType returns go type with/wo a pointer
func getType(s *Schema) *jen.Statement {
	if s.isEnum() {
		return jen.Id(s.CamelName)
	}

	switch s.Type {
	case SchemaTypeAny:
		return jen.Any()
	case SchemaTypeString, SchemaTypeInteger, SchemaTypeNumber, SchemaTypeBoolean, SchemaTypeTime:
		return withPointer(getScalarType(s), s.required)
	}

	switch {
	case s.isArray():
		return withPointer(jen.Index(), s.required || s.isOut()).Add(getType(s.Items))
	case s.isObject():
		o := jen.Id(s.CamelName)
		if s.isAnonymous() {
			o = fmtStruct(s)
		}

		return withPointer(o, s.required)
	case s.isMap():
		a := withPointer(jen.Map(jen.String()), s.required || s.isOut())
		if s.AdditionalProperties != nil {
			s.AdditionalProperties.required = true
			return a.Add(getType(s.AdditionalProperties))
		} else if s.name == "tags" {
			// tags are everywhere in the schema, better not to use the patch
			return a.String()
		} else {
			return a.Any()
		}
	default:
		panic(fmt.Errorf("unknown type %q for %q and parent %q", s.Type, s.name, s.parent.name))
	}
}

func withPointer(j *jen.Statement, required bool) *jen.Statement {
	if required {
		return j
	}
	return jen.Op("*").Add(j)
}

func mustMarshal(s any) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func lowerFirst(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func sortedKeys[T any](m map[string]T) []string {
	keys := maps.Keys(m)
	sort.Strings(keys)

	return keys
}

var cleanEnumName = regexp.MustCompile("(Create|Get|Update|Delete|Stop|Cancel|Verify|Put)")

// getEnumName enum can't have just "state" name, drills to the root until finds something
func getEnumName(s *Schema) string {
	switch s.name {
	case "type", "value", "state", "status":
		if s.parent != nil {
			return getEnumName(s.parent)
		}
	}

	return s.name
}

var camelFinder = regexp.MustCompile("[A-Z]+[a-z]+")

func dedupCamelName(src string) string {
	result := make([]string, 0)
	for _, s := range camelFinder.FindAllString(src, -1) {
		if !slices.Contains(result, s) {
			result = append(result, s)
		}
	}

	return strings.Join(result, "")
}
