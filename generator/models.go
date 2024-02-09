//go:build generator

package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
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

// nolint:funlen,gocognit // It is easy to maintain and read, we don't need to split it
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

		if !strings.HasSuffix(s.CamelName, enumTypeSuffix) {
			s.CamelName += enumTypeSuffix
		}

		// When it is just "Type" it is useless
		if s.CamelName == enumTypeSuffix {
			s.CamelName = s.parent.CamelName + s.CamelName
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

	if s.isPrivate() {
		s.CamelName = lowerFirst(s.CamelName)
	}

	if s.parent != nil && s.parent.isPrivate() {
		s.CamelName = strcase.ToCamel(s.parent.CamelName)
	}

	if s.Type == SchemaTypeString {
		parts := strings.Split(s.name, "_")
		suffix := parts[len(parts)-1]

		if len(parts) > 1 && (suffix == "at" || suffix == "time") {
			s.Type = SchemaTypeTime
		}
	}

	if s.isArray() {
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

	if s.isObject() {
		keys := sortedKeys(scope)
	outer:
		for len(keys) > 0 {
			for _, k := range keys {
				other := scope[k]
				if other.hash() == s.hash() {
					continue
				}

				// A duplicate
				if other.CamelName == s.CamelName {
					s.CamelName += "Alt"

					continue outer
				}
			}

			break outer
		}

		scope[s.hash()] = s
	}

	if s.isEnum() {
		// Enums compared by enum list
		// In case if they are equal, they must have the same name
		other, ok := scope[s.hash()]
		if ok {
			s.CamelName = other.CamelName
		} else {
			scope[s.hash()] = s
		}
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
	return len(s.Enum) != 0 && s.isIn()
}

func (s *Schema) root() *Schema {
	if s.parent == nil {
		return s
	}

	return s.parent.root()
}

func (s *Schema) isIn() bool {
	return s.root().in
}

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
	if s.isEnum() {
		return jen.Id(s.CamelName)
	}

	if s.isScalar() {
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
		if s.Items.isObject() || s.Items.isArray() {
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
