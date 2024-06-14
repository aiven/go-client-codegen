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

func (d *Doc) getSchema(path string) (*Schema, error) {
	name := strings.Split(path, "/")[3]

	schema := d.Components.Schemas[name]
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
	s := lowerFirst(p.Summary)
	c := jen.Comment(fmt.Sprintf("%s %s", p.FuncName, s))
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
	Required    bool        `json:"required"`
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
	Type          SchemaType         `json:"type"`
	Properties    map[string]*Schema `json:"properties"`
	Items         *Schema            `json:"items"`
	RequiredProps []string           `json:"required"`
	Enum          []any              `json:"enum"`
	Default       any                `json:"default"`
	MinItems      int                `json:"minItems"`
	Ref           string             `json:"$ref"`
	CamelName     string             `json:"for-hash-only!"`
	required      bool
	name          string
	propertyNames []string
	parent        *Schema
	in, out       bool // Request or Response DTO
}

// nolint:funlen,gocognit,gocyclo // It is easy to maintain and read, we don't need to split it
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

	if s.isObject() {
		switch {
		case s.isIn():
			s.CamelName += "In"
		case s.isOut():
			s.CamelName += "Out"
		}
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
		suffix := parts[len(parts)-1]

		if len(parts) > 1 && (suffix == "at" || suffix == "time") {
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

	if s.isObject() || s.isEnum() {
		for s.parent != nil {
			v, ok := scope[s.CamelName]
			if !ok {
				break
			}

			if v.hash() == s.hash() {
				// This is a duplicate
				return
			}

			s.CamelName += "Alt"
		}

		scope[s.CamelName] = s
	}
}

func (s *Schema) isPrivate() bool {
	return s.parent == nil && s.out && len(s.Properties) == 1
}

func (s *Schema) hash() string {
	if s.isEnum() {
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

func (s *Schema) isNestedArray() bool {
	return s.isArray() && s.Items.isArray()
}

func (s *Schema) isScalar() bool {
	switch s.Type {
	case SchemaTypeString, SchemaTypeInteger, SchemaTypeNumber, SchemaTypeBoolean, SchemaTypeTime:
		return true
	}

	return false
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
	switch {
	case s.Type == SchemaTypeAny:
		return jen.Any()
	case s.isEnum():
		return jen.Id(s.CamelName)
	case s.isScalar():
		scalar := getScalarType(s)
		if !s.required && s.Type != SchemaTypeString {
			return jen.Op("*").Add(scalar)
		}

		return scalar
	}

	switch {
	case s.isArray():
		a := jen.Index()
		if !(s.required || s.isOut()) {
			a = jen.Op("*").Index()
		}

		// No pointers for complex objects
		switch {
		case s.isNestedArray():
			// but not nested array
		case s.Items.isObject() || s.Items.isArray():
			return a.Id(s.Items.CamelName)
		}

		return a.Add(getType(s.Items))
	case s.isObject():
		if !s.required {
			return jen.Id("*" + s.CamelName)
		}

		return jen.Id(s.CamelName)
	case s.isMap():
		a := jen.Map(jen.String())
		if !(s.required || s.isOut()) {
			a = jen.Op("*").Map(jen.String())
		}

		if isMapString(s) {
			return a.String()
		} else {
			return a.Any()
		}
	default:
		panic(fmt.Errorf("unknown type %q for %q and parent %q", s.Type, s.name, s.parent.name))
	}
}

func mustMarshal(s any) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	return string(b)
}

// isMapString for hacking schemaless maps
func isMapString(s *Schema) bool {
	return s.name == "tags"
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
