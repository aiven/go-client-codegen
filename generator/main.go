//go:build generator

// Package main is the generator of the client code.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

const (
	generatedHeader   = "Code generated by Aiven. DO NOT EDIT."
	configPrefix      = "GEN"
	defaultAPIVersion = "v1"
)

type envConfig struct {
	Module           string `envconfig:"MODULE" default:"github.com/aiven/go-client-codegen"`
	Package          string `envconfig:"PACKAGE" default:"aiven"`
	HandlerDir       string `envconfig:"HANDLER_DIR" default:"handler"`
	ConfigFile       string `envconfig:"CONFIG_FILE" default:"config.yaml"`
	ClientFile       string `envconfig:"CLIENT_FILE" default:"client_generated.go"`
	OpenAPIFile      string `envconfig:"OPENAPI_FILE" default:"openapi.json"`
	OpenAPIPatchFile string `envconfig:"OPENAPI_PATCH_FILE" default:"openapi_patch.yaml"`
}

var (
	pathClean      = regexp.MustCompile(`\{[^{]+}`)
	pathVersioning = regexp.MustCompile(`^/v[0-9]/`)
)

var strFormatters = map[SchemaType]string{
	SchemaTypeInteger: "%d",
	SchemaTypeNumber:  "%f",
	SchemaTypeString:  "%s",
	SchemaTypeBoolean: "%t",
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

	err := exec()
	if err != nil {
		log.Err(err).Send()
	}
}

const (
	doerName             = "doer"
	handlerTypeName      = "Handler"
	queryParamName       = "query"
	queryParamTypeSuffix = "Query"
	queryParamArraySize  = 2
)

//nolint:funlen,gocognit,gocyclo // It's a generator, it's supposed to be long, and we won't expand it.
func exec() error {
	cfg := new(envConfig)

	err := envconfig.Process(configPrefix, cfg)
	if err != nil {
		return err
	}

	config, err := readConfig(cfg.ConfigFile)
	if err != nil {
		return err
	}

	// Check for duplicate endpoints
	err = checkDuplicateEndpoints(config)
	if err != nil {
		return err
	}

	// Reads OpenAPI file and applies a patch
	docBytes, err := readOpenAPIPatched(cfg.OpenAPIFile, cfg.OpenAPIPatchFile)
	if err != nil {
		return err
	}

	doc := new(Doc)
	err = json.Unmarshal(docBytes, doc)
	if err != nil {
		return err
	}

	// To validate all operation ids in the config exist in the OpenAPI spec
	// OperationID => Package name
	configOperationIDs := make(map[string]string)
	for pkg, idList := range config {
		for _, id := range idList {
			configOperationIDs[id] = pkg
		}
	}

	pkgs := make(map[string][]*Path)
	for path := range doc.Paths {
		v := doc.Paths[path]
		for meth, p := range v {
			if !pathVersioning.MatchString(path) {
				path = fmt.Sprintf("/%s%s", defaultAPIVersion, path)
			}

			p.Path = path
			p.Method = strings.ToUpper(meth)
			p.ID = p.OperationID

			pkg, ok := configOperationIDs[p.ID]
			if !ok {
				log.Warn().Msgf("%q id not found in config!", p.ID)
				continue
			}

			// Removes the operation id from the map to see which are not used
			delete(configOperationIDs, p.ID)

			pkgs[pkg] = append(pkgs[pkg], p)
			params := make([]*Parameter, 0)

			for _, ref := range p.Parameters {
				parts := strings.Split(ref.Ref, "/")
				name := parts[len(parts)-1]

				param, ok := doc.Components.Parameters[name]
				if !ok {
					return fmt.Errorf("param %q not found", ref.Ref)
				}

				if param.Name == "version_id" {
					param.Schema.Type = SchemaTypeInteger
				}

				param.Ref = ref.Ref
				params = append(params, param)
			}

			p.Parameters = params
		}
	}

	if len(configOperationIDs) > 0 {
		return fmt.Errorf("config has unused operation ids: %s", strings.Join(sortedKeys(configOperationIDs), ", "))
	}

	ctx := jen.Id("ctx").Qual("context", "Context")
	doer := jen.Comment(doerName + " http client").Line().Type().Id(doerName).Interface(
		jen.Id("Do").Params(
			ctx,
			jen.List(jen.Id("operationID"), jen.Id("method"), jen.Id("path")).String(),
			jen.Id("in").Any(),
			jen.Id(queryParamName).Op("...").Add(fmtQueryParamType()),
		).Parens(jen.List(jen.Index().Byte(), jen.Error())),
	).Line()

	clientFields := make([]jen.Code, 0, len(pkgs))
	clientValues := jen.Dict{}
	clientTypeValues := make([]jen.Code, 0, len(pkgs))

	for _, pkg := range sortedKeys(pkgs) {
		paths := pkgs[pkg]
		fileName := strings.ToLower(pkg)
		handlerName := pkg + handlerTypeName
		newHandler := "New" + handlerTypeName
		scope := make(map[string]*Schema)

		for _, p := range paths {
			p.FuncName = p.OperationID
		}

		sort.SliceStable(paths, func(i, j int) bool {
			return paths[i].FuncName < paths[j].FuncName
		})

		file := jen.NewFile(fileName)
		file.HeaderComment(generatedHeader)

		// Creates the handler's type (interface)
		// Reserves the line in the file
		handlerType := file.Type().Id(handlerTypeName)

		// Adds private types (interfaces)
		file.Add(doer)

		// Creates the "new" constructor
		file.Func().Id(newHandler).Params(jen.Id(doerName).Id(doerName)).Id(handlerName).Block(
			jen.Return(jen.Id(handlerName).Values(jen.Id(doerName))),
		)

		// Creates the handler's implementation
		file.Type().Id(handlerName).Struct(jen.Id(doerName).Id(doerName))

		var typeMethods []jen.Code
		for _, path := range paths {
			// todo: support 204
			out := path.Out.OK.Content["application/json"]
			if out == nil && path.Out.NoContent.Content == nil {
				log.Printf("%q has no json response. Skipping", path.OperationID)
				continue
			}

			// Method's schemas and query params
			schemas := make([]*Schema, 0)
			queryParams := make([]*Schema, 0)

			// Interface and implementation args
			funcArgs := []jen.Code{ctx}

			// Collects params: in path and in query
			// Adds to schemas to render enums
			for _, p := range path.Parameters {
				p.Schema.required = true
				p.Schema.init(doc, scope, p.Name)

				if p.In == ParameterInPath {
					schemas = append(schemas, p.Schema)
					param := jen.Id(p.Schema.lowerCamel()).Add(getType(p.Schema))
					funcArgs = append(funcArgs, param)
					continue
				}

				queryParams = append(queryParams, p.Schema)

				// Adds param function (request modifier)
				var code *jen.Statement
				code, err = fmtQueryParam(path.FuncName, p)
				if err != nil {
					return err
				}
				file.Add(code)
			}

			in := path.In.Content["application/json"]
			if in != nil {
				var schemaIn *Schema

				schemaIn, err = doc.getSchema(in.Schema.Ref)
				if err != nil {
					return err
				}

				schemaIn.in = true
				schemaIn.init(doc, scope, path.FuncName)
				schemas = append(schemas, schemaIn)
				funcArgs = append(funcArgs, jen.Id("in").Id("*"+schemaIn.CamelName))
			}

			// Adds queryParams options
			if len(queryParams) > 0 {
				funcArgs = append(funcArgs, jen.Id(queryParamName).Op("...").Add(fmtQueryParamType()))
			}

			typeMeth := jen.Id(path.FuncName).Params(funcArgs...)
			structMeth := jen.Func().Params(jen.Id("h").Id("*" + handlerName)).Id(path.FuncName).Params(funcArgs...)

			var rsp, schemaOut *Schema
			if out != nil {
				schemaOut, err = doc.getSchema(out.Schema.Ref)
				if err != nil {
					return err
				}

				schemaOut.out = true
				schemaOut.init(doc, scope, path.FuncName)
				rsp = getResponse(schemaOut)
			}

			// forcePointer Required objects must be returned by a pointer for consistency
			forcePointer := rsp != nil && rsp.required && rsp.isObject()

			if rsp != nil {
				ret := jen.List(getType(rsp), jen.Error())
				if forcePointer {
					// foo() (*Foo, err)
					ret = jen.List(jen.Id("*"+rsp.CamelName), jen.Error())
				}

				typeMeth.Parens(ret)
				structMeth.Parens(ret)
			} else {
				typeMeth.Error()
				structMeth.Error()
			}

			typeMethods = append(typeMethods, path.Comment(), typeMeth.Line())

			// Crates a go formattable path, i.e.:
			// /foo/{foo}/ => /foo/%s/
			paramIndex := -1
			url := pathClean.ReplaceAllStringFunc(path.Path, func(_ string) string {
				paramIndex++
				t, ok := strFormatters[path.Parameters[paramIndex].Schema.Type]
				if !ok {
					panic(fmt.Sprintf("%s unexpected parameter type %s", path.OperationID, t))
				}
				return t
			})

			urlParams := make([]jen.Code, 0, len(funcArgs))
			urlParams = append(urlParams, jen.Lit(url))
			inObj := jen.Nil()
			for _, s := range schemas {
				if s.isObject() {
					inObj = jen.Id("in")
					continue
				}

				v := jen.Id(s.lowerCamel())
				if s.isEnum() {
					// Stringifies enums
					v = jen.String().Call(v)
				}

				// Escapes string values
				if s.Type == SchemaTypeString {
					v = jen.Qual("net/url", "PathEscape").Call(v)
				}

				urlParams = append(urlParams, v)
			}

			outObj := jen.Id("_")
			returnErr := jen.Return(jen.Err())

			// Formats "return" statement
			if rsp != nil {
				outObj = jen.Id("b")

				// In most cases, "nil" is for error return
				// But for required scalars should be zero values
				returnErr = jen.Return(jen.Nil(), jen.Err())

				if rsp.required {
					switch rsp.Type {
					case SchemaTypeString:
						returnErr = jen.Return(jen.Lit(""), jen.Err())
					case SchemaTypeInteger, SchemaTypeNumber:
						returnErr = jen.Return(jen.Lit(0), jen.Err())
					case SchemaTypeBoolean:
						returnErr = jen.Return(jen.False(), jen.Err())
					}
				}
			}

			// The Doer call
			callOpts := []jen.Code{
				jen.Id("ctx"),
				jen.Lit(path.OperationID),
				jen.Lit(path.Method),
				jen.Id("path"),
				inObj,
			}

			var block []jen.Code

			// Adds unpacking for query params
			if len(queryParams) > 1 {
				q := jen.Id(queryParamName)
				p := jen.Id("p")
				v := jen.Id("v")
				callOpts = append(callOpts, p.Clone().Op("..."))
				block = append(
					block,
					p.Clone().Op(":=").Make(jen.Index().Index(jen.Lit(queryParamArraySize)).String(), jen.Lit(0), jen.Len(q)),
					jen.For(jen.List(jen.Id("_"), v.Clone().Op(":=").Range().Add(jen.Id(queryParamName)))).
						Block(p.Clone().Op("=").Append(p, v)),
				)
			}

			// Implementation (method's) body
			block = append(
				block,
				jen.Id("path").Op(":=").Qual("fmt", "Sprintf").Call(urlParams...),
				jen.List(outObj, jen.Err()).Op(":=").Id("h.doer.Do").Call(callOpts...),
			)

			ifErr := jen.If(jen.Err().Op("!=").Nil()).Block(returnErr)
			if rsp == nil {
				block = append(block, jen.Return(jen.Err()))
			} else {
				block = append(block, ifErr)
				outReturn := jen.Id("out")

				if rsp.CamelName != schemaOut.CamelName {
					// Takes original name and turns to camel.
					// "CamelName" field might have been modified because of name collisions
					outReturn.Dot(customCamelCase(rsp.name))

					if forcePointer {
						// return &out.Foo
						outReturn = jen.Id("&").Add(outReturn)
					}
				}

				block = append(
					block,
					jen.Id("out").Op(":=").New(jen.Id(schemaOut.CamelName)),
					jen.Err().Op("=").Qual("encoding/json", "Unmarshal").Call(jen.Id("b"), jen.Id("out")),
					ifErr,
					jen.Return(outReturn, jen.Nil()),
				)
			}

			file.Add(structMeth.Block(block...))
		}

		for _, k := range sortedKeys(scope) {
			v := scope[k]
			err = writeStruct(file, v)
			if err != nil {
				return err
			}
		}

		dirPath := filepath.Join(cfg.HandlerDir, fileName)

		err = os.MkdirAll(dirPath, dirMode)
		if err != nil {
			return err
		}

		handlerType.Interface(typeMethods...)

		err = file.Save(filepath.Join(dirPath, fileName+".go"))
		if err != nil {
			return err
		}

		pkgName := filepath.Join(cfg.Module, cfg.HandlerDir, fileName)
		clientFields = append(clientFields, jen.Qual(pkgName, handlerName))
		clientValues[jen.Id(handlerName)] = jen.Qual(pkgName, newHandler).Call(jen.Id(doerName))
		clientTypeValues = append(clientTypeValues, jen.Qual(pkgName, handlerTypeName))
	}

	client := jen.NewFile(cfg.Package)
	client.HeaderComment(generatedHeader)
	client.Add(doer)
	client.Func().Id("newClient").Params(jen.Id(doerName).Id(doerName)).Id("Client").Block(
		jen.Return(jen.Id("&client").Values(clientValues)),
	)
	client.Type().Id("client").Struct(clientFields...)
	client.Type().Id("Client").Interface(clientTypeValues...)

	return client.Save(cfg.ClientFile)
}

// reMakesSense sometimes there are invalid enums, for instance, just a comma ","
var reMakesSense = regexp.MustCompile(`\w`)

//nolint:funlen,nestif // It's a generator, it's supposed to be long, and we won't expand it.
func writeStruct(f *jen.File, s *Schema) error {
	if s.isAnonymous() {
		return nil
	}

	if s.isEnum() {
		kind := getScalarType(s)
		o := f.Type().Id(s.CamelName)
		o.Add(kind)

		enums := make([]jen.Code, 0)
		values := make([]jen.Code, 0)
		for _, e := range s.Enum {
			literal := fmt.Sprint(e)
			if !reMakesSense.MatchString(literal) {
				continue
			}

			constant := s.CamelName + customCamelCase(literal)

			// KafkaMirror ReplicationPolicyClassType makes bad generated name
			if strings.HasPrefix(literal, "org.apache.kafka.connect.mirror.") {
				constant = s.CamelName + literal[32:len(literal)-17]
			}

			// OpenSearch HealthType has value "red*"
			if strings.HasSuffix(literal, "*") {
				constant += "Asterisk"
			}

			// Turns integer literals into integers
			var v any = literal
			if s.Type == SchemaTypeInteger {
				i, err := strconv.Atoi(literal)
				if err != nil {
					return err
				}
				v = i
			}

			enums = append(enums, jen.Id(constant).Op(s.CamelName).Op("=").Lit(v))
			values = append(values, jen.Lit(v))
		}

		if len(enums) == 0 {
			return nil
		}

		o.Line().Const().Defs(enums...)
		o.Line().Func().Id(s.CamelName + "Choices").Params().Index().Add(kind).Block(
			jen.Return(jen.Index().Add(kind).Values(values...)),
		)

		return nil
	}

	if s.Description != "" {
		f.Comment(fmt.Sprintf("%s %s", s.CamelName, fmtComment(s.Description)))
	}

	f.Type().Id(s.CamelName).Add(fmtStruct(s))
	return nil
}

// reInvertSnakeCase snake case name pattern reversed
var reInvertSnakeCase = regexp.MustCompile("[^a-zA-Z0-9_]+")

// fmtStruct returns anonymous struct
func fmtStruct(s *Schema) *jen.Statement {
	// Sorts field names
	// Resolves field name collision when a new more conventional JSON name replaces an existing field
	jsonNames := maps.Keys(s.Properties)
	sort.Slice(jsonNames, func(i, j int) bool {
		return reInvertSnakeCase.ReplaceAllString(jsonNames[i], "") > reInvertSnakeCase.ReplaceAllString(jsonNames[j], "")
	})

	// Resolves collisions
	uniqueNames := make(map[string]string, len(jsonNames))
	for _, jsonName := range jsonNames {
		p := s.Properties[jsonName]
		goName := customCamelCase(jsonName)
		if exist, ok := uniqueNames[goName]; ok {
			log.Warn().Msgf("Field collision: %q overrides %q", p.path(), exist)
		}
		// WARNING: This is a hack to avoid name collisions in Go fields
		// overrides duplicate fields
		uniqueNames[goName] = jsonName
	}

	fields := make([]jen.Code, 0, len(uniqueNames))
	for _, goName := range sortedKeys(uniqueNames) {
		jsonName := uniqueNames[goName]
		p := s.Properties[jsonName]

		// Adds json tags
		tag := jsonName
		if !p.required {
			tag += ",omitempty"
		}

		// Some fields have special characters escaped in the name
		tag = strings.ReplaceAll(tag, `\`, "")

		field := jen.Id(goName).Add(getType(p))
		field = field.Tag(map[string]string{"json": tag})

		// Adds a comment if it's not equal to the field name
		if p.Description != "" && p.Description != p.CamelName {
			field = field.Add(jen.Comment(fmtComment(p.Description)))
		}

		fields = append(fields, field)
	}
	return jen.Struct(fields...)
}

func getResponse(s *Schema) *Schema {
	switch len(s.Properties) {
	case 1:
		// If the schema has just one field, then uses it as out dto.
		// That makes code simpler.
		return s.Properties[s.propertyNames[0]]
	case 0:
		return nil
	}

	return s
}

func toSingle(src string) string {
	s := strings.TrimSuffix(src, "ies")
	if s != src {
		return s + "y"
	}

	return strings.TrimSuffix(src, "s")
}

const (
	yamlTabSize = 2
	writeMode   = os.FileMode(0o644)
	dirMode     = os.FileMode(0o750)
)

// readConfig reads and formats the config
func readConfig(path string) (map[string][]string, error) {
	filePath := filepath.Clean(path)

	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	c := make(map[string][]string)

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}

	// Updates the config
	for _, v := range c {
		slices.Sort(v)
	}

	var buffer bytes.Buffer
	encoder := yaml.NewEncoder(&buffer)
	encoder.SetIndent(yamlTabSize)

	err = encoder.Encode(&c)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(filePath, buffer.Bytes(), writeMode)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// reComment finds new lines and trailing period,
// which is added in fmtComment as a separater
var reComment = regexp.MustCompile(`\.?[\r\n]+\s*?`)

func fmtComment(c string) string {
	return html.UnescapeString(reComment.ReplaceAllString(c, ". "))
}

// fmtQueryParam returns a query param
func fmtQueryParam(funcName string, p *Parameter) (*jen.Statement, error) {
	keyFuncName := funcName + p.Schema.CamelName
	keyVarName := jen.Id(p.Schema.lowerCamel())

	format, ok := strFormatters[p.Schema.Type]
	if !ok {
		return nil, fmt.Errorf("query param with type %q is not supported", p.Schema.Type)
	}

	// Stringifies non-string values and enums
	value := keyVarName.Clone()
	if p.Schema.isEnum() || p.Schema.Type != SchemaTypeString {
		value = jen.Qual("fmt", "Sprintf").Call(jen.Lit(format), keyVarName.Clone())
	}

	param := jen.Comment(fmt.Sprintf("%s %s", keyFuncName, fmtComment(p.Description)))
	param.Line()
	param.Func().Id(keyFuncName).
		Params(keyVarName.Clone().Add(getType(p.Schema))).Params(fmtQueryParamType()).
		Block(
			jen.Return(fmtQueryParamType().Values(jen.Lit(p.Schema.name), value)),
		)
	return param, nil
}

// fmtQueryParamType literally returns: [2]string
func fmtQueryParamType() *jen.Statement {
	return jen.Index(jen.Lit(queryParamArraySize)).String()
}

var reNonWord = regexp.MustCompile(`\W+`)

func customCamelCase(s string) string {
	return strcase.ToCamel(reNonWord.ReplaceAllString(s, "_"))
}
