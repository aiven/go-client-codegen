package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
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
	s := strings.ToLower(p.Summary[:1]) + p.Summary[1:]
	c := jen.Comment(fmt.Sprintf("%s %s", p.FuncName, s))
	c.Line().Comment(fmt.Sprintf("%s %s %s", p.OperationID, p.Method, p.Path))
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
	required      bool
	hash          string
	name          string
	camelName     string
	properties    []*Schema
	parent        *Schema
	in, out       bool // Request or Response DTO
}

func (s *Schema) init(scope map[string]*Schema, prefix, name string) {
	s.hash = mustMarshal(s)
	name = strings.ReplaceAll(name, `\`, "")
	s.name = name
	s.camelName = strcase.ToCamel(name)

	pluralPrefix := prefix + "s"
	switch s.camelName {
	case prefix, pluralPrefix:
	default:
		s.camelName = trimPrefix(prefix, s.camelName)
	}

	if s.isEnum() && !strings.HasSuffix(s.camelName, "Type") {
		s.camelName += "Type"
	}

	if s.out {
		delete(s.Properties, "errors")
		delete(s.Properties, "message")
	}

	if s.isObject() || s.isEnum() {
		for _, k := range s.RequiredProps {
			p, ok := s.Properties[k]
			if ok {
				p.required = true
			}
		}

		for {
			other, ok := scope[s.camelName]
			if !ok || s.hash == other.hash {
				break
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
		s.Items.init(scope, s.camelName, toSingle(name))
	}

	if s.Type == SchemaTypeString {
		parts := strings.Split(s.name, "_")
		switch parts[len(parts)-1] {
		case "at", "time":
			s.Type = SchemaTypeTime
		}
	}

	keys := lo.Keys(s.Properties)
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})

	s.properties = make([]*Schema, 0, len(s.Properties))
	for _, k := range keys {
		p := s.Properties[k]
		p.parent = s
		p.init(scope, s.camelName, k)
		s.properties = append(s.properties, p)
	}

	sort.SliceStable(s.properties, func(i, j int) bool {
		return s.properties[i].camelName < s.properties[j].camelName
	})
}

func (s *Schema) isObject() bool {
	return s.Type == SchemaTypeObject
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
	return s.isObject() && len(s.Properties) == 0
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
		if len(s.Items.properties) != 0 {
			return jen.Index().Id(s.Items.camelName)
		}
		return jen.Index().Add(getType(s.Items))
	case s.isObject():
		if s.isMap() {
			if isMapString(s) {
				return jen.Map(jen.String()).String()
			} else {
				return jen.Map(jen.String()).Any()
			}
		}
		return jen.Id("*" + s.camelName)
	default:
		panic(fmt.Errorf("unknown type %q", s.Type))
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
