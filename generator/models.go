package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
)

const docSite = "https://api.aiven.io/doc"

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

type Content map[string]*struct {
	Schema struct {
		Ref string `json:"$ref"`
	} `json:"schema"`
}

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

type ParameterIn string

const (
	ParameterInPath ParameterIn = "path"
)

type Parameter struct {
	Ref         string      `json:"$ref"`
	In          ParameterIn `json:"in"`
	Required    bool        `json:"required"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Schema      *Schema     `json:"schema"`
}

type SchemaType string

const (
	SchemaTypeObject  = "object"
	SchemaTypeArray   = "array"
	SchemaTypeString  = "string"
	SchemaTypeInteger = "integer"
	SchemaTypeNumber  = "number"
	SchemaTypeBoolean = "boolean"
	SchemaTypeTime    = "time"
)

type Schema struct {
	Type          SchemaType         `json:"type"`
	Properties    map[string]*Schema `json:"properties"`
	Items         *Schema            `json:"items"`
	RequiredProps []string           `json:"required"`
	Enum          []any              `json:"enum"`
	Default       any                `json:"default"`
	MinItems      int                `json:"minItems"`
	Ref           string             `json:"$ref"`
	required      bool
	hash          string
	name          string
	camelName     string
	propertyNames []string
	parent        *Schema
	in, out       bool // Request or Response DTO
}

func (s *Schema) init(doc *Doc, scope map[string]*Schema, prefix, name string) {
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

	if s.out {
		delete(s.Properties, "errors")
		delete(s.Properties, "message")
	}

	s.name = name
	s.hash = mustMarshal(s)

	s.camelName = strcase.ToCamel(s.name)
	if s.isEnum() && !strings.HasSuffix(s.camelName, "Type") {
		s.camelName += "Type"
	}

	if s.isObject() || s.isEnum() {
		for _, k := range sortedKeys(scope) {
			other := scope[k]
			if s.parent == nil {
				break
			}

			if s.hash == other.hash {
				// fixme: need some "deep copy" here
				*s = *other
				// Must preserve the name, because this is how the parent field is called
				s.name = name
				return
			}

			if s.camelName != other.camelName {
				continue
			}

			parent := s.parent
			if parent.isArray() {
				parent = parent.parent
				if !strings.HasSuffix(s.camelName, "Item") {
					s.camelName += "Item"
					continue
				}
			}

			s.camelName = parent.camelName + s.camelName
		}
		scope[s.camelName] = s
	}

	if s.isArray() {
		s.Items.parent = s
		s.Items.required = true // a workaround to not have slices with pointers
		s.Items.init(doc, scope, s.camelName, toSingle(name))
	}

	if s.Type == SchemaTypeString {
		parts := strings.Split(s.name, "_")
		switch parts[len(parts)-1] {
		case "at", "time":
			s.Type = SchemaTypeTime
		}
	}

	s.propertyNames = sortedKeys(s.Properties)
	for _, k := range s.propertyNames {
		p := s.Properties[k]
		p.parent = s
		p.init(doc, scope, s.camelName, k)
	}

	// KafkaTopicConfig hacks.
	// Because of deduplication in the scope, each field gets the first rendered type.
	// For example, each int field becomes DeleteRetentionMs
	if s.isObject() && slices.Equal(s.propertyNames, []string{"source", "synonyms", "value"}) {
		p := s.Properties["value"]
		if p.Enum == nil {
			delete(scope, s.camelName)
			s.camelName = "TopicConfig" + upperFirst(string(p.Type))
			scope[s.camelName] = s
		}
	}
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
	return len(s.Enum) != 0
}

func (s *Schema) isOut() bool {
	p := s.parent
	for p != nil {
		if p.out {
			return true
		}
		p = p.parent
	}
	return false
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

func getType(s *Schema) *jen.Statement {
	if s.isEnum() {
		return jen.Id(s.camelName)
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
		if len(s.Items.Properties) != 0 {
			return a.Id(s.Items.camelName)
		}
		return a.Add(getType(s.Items))
	case s.isObject():
		return jen.Id("*" + s.camelName)
	case s.isMap():
		if isMapString(s) {
			return jen.Map(jen.String()).String()
		} else {
			return jen.Map(jen.String()).Any()
		}
	default:
		panic(fmt.Errorf("unknown type %q for %q and parent %q", s.Type, s.name, s.parent.name))
	}
}

func mustMarshal(s *Schema) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(fmt.Errorf("err marshal %q: %w", s.name, err))
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

func upperFirst(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func sortedKeys[T any](m map[string]T) []string {
	keys := maps.Keys(m)
	sort.Strings(keys)
	return keys
}
