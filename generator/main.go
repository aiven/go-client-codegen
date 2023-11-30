package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"github.com/samber/lo"
)

const (
	genImport       = "github.com/aiven/aiven-go-client-v2"
	openapiFile     = "openapi.json"
	clientDir       = "aiven"
	handlerDir      = "handler"
	generatedHeader = "Code generated by Aiven. DO NOT EDIT."
	versionIDParam  = `/{version_id:latest|\d+}`
)

var flinkParseID = regexp.MustCompile(`(Flink)([A-Z][a-z]+)(\w+)`)
var pathClean = regexp.MustCompile(`\{[^{]+}`)

func main() {
	err := exec()
	if err != nil {
		log.Fatal(err)
	}
}

func exec() error {
	b, err := os.ReadFile(openapiFile)
	if err != nil {
		return err
	}

	var doc Doc
	err = json.Unmarshal(b, &doc)
	if err != nil {
		return err
	}

	pkgs := make(map[string][]*Path)
	for path, v := range doc.Paths {
		for meth, p := range v {
			if p.Deprecated {
				continue
			}

			p.Path = path
			p.Method = strings.ToUpper(meth)
			p.ID = p.OperationID

			// Service:_Kafka -> ServiceKafka
			pkg := strings.ReplaceAll(p.Tags[0], "Service:_", "")
			if p.Tags[0] != pkg {
				p.ID = strings.TrimPrefix(p.ID, "Service")
			}

			pkg = strings.TrimSuffix(strcase.ToCamel(pkg), "s")

			if strings.HasPrefix(p.ID, "ServiceIntegrationEndpoint") {
				pkg = "ServiceIntegrationEndpoint"
			}

			if pkg == "Group" {
				pkg = "UserGroup"
			}

			if pkg == "Kafka" {
				if strings.HasPrefix(p.ID, "KafkaTopic") {
					pkg = "KafkaTopic"
				}

				if strings.HasPrefix(p.ID, "KafkaConnect") {
					pkg = "KafkaConnect"
				}

				if strings.HasPrefix(p.ID, "KafkaTieredStorage") {
					// Double "Storage"
					p.ID = strings.ReplaceAll(p.ID, "StorageStorage", "Storage")
				}

				if strings.HasPrefix(p.ID, "SchemaRegistry") {
					pkg = "KafkaSchemaRegistry"
					p.ID = strings.TrimPrefix(p.ID, "SchemaRegistry")
				}
			}

			if pkg == "Flink" {
				if strings.Contains(p.ID, "Job") {
					pkg = "FlinkJob"
				} else if strings.Contains(p.ID, "ApplicationDeployment") {
					pkg = "FlinkApplicationDeployment"
					p.ID = flinkParseID.ReplaceAllString(p.ID, "$1$3$2")
				} else if strings.Contains(p.ID, "ApplicationVersion") {
					pkg = "FlinkApplicationVersion"
					p.ID = flinkParseID.ReplaceAllString(p.ID, "$1$3$2")
				} else if strings.Contains(p.ID, "Application") {
					pkg = "FlinkApplication"
					p.ID = flinkParseID.ReplaceAllString(p.ID, "$1$3$2")
				}
			}

			if strings.Contains(p.ID, "Privatelink") {
				pkg = "Privatelink"
			}

			// fixme: This is bad
			if pkg == "" {
				pkg = "Misc"
			}

			pkgs[pkg] = append(pkgs[pkg], p)
			params := make([]*Parameter, 0, len(p.Parameters))
			for _, ref := range p.Parameters {
				name, _ := lo.Last(strings.Split(ref.Ref, "/"))
				param, ok := doc.Components.Parameters[name]
				if !ok {
					return fmt.Errorf("param %q not found", name)
				}

				if param.In != ParameterInPath {
					log.Printf("%q param %s in %q. Skipping", p.OperationID, param.Name, param.In)
					continue
				}

				param.Ref = ref.Ref
				params = append(params, param)
			}

			if strings.HasSuffix(p.Path, versionIDParam) {
				params = append(params, &Parameter{
					Name:   "version_id",
					Schema: &Schema{Type: SchemaTypeInteger},
				})
			}

			p.Parameters = params
		}
	}

	const doerName = "doer"

	ctx := jen.Id("ctx").Qual("context", "Context")
	doer := jen.Type().Id(doerName).Interface(
		jen.Id("Do").Params(
			ctx,
			jen.List(jen.Id("operationID"), jen.Id("method"), jen.Id("path")).String(),
			jen.Id("v").Any(),
		).Parens(jen.List(jen.Index().Byte(), jen.Error())),
	).Line()
	clientFields := make([]jen.Code, 0, len(pkgs))
	clientValues := jen.Dict{}

	pkgKeys := lo.Keys(pkgs)
	slices.Sort(pkgKeys)

	for _, pkg := range pkgKeys {
		paths := pkgs[pkg]
		fileName := strings.ToLower(pkg)
		scope := make(map[string]*Schema)
		for _, p := range paths {
			//p.FuncName = p.OperationID
			p.FuncName = trimPrefix(pkg, p.ID)
		}

		sort.SliceStable(paths, func(i, j int) bool {
			return paths[i].FuncName < paths[j].FuncName
		})

		file := jen.NewFile(fileName)
		file.HeaderComment(generatedHeader)
		handler := file.Type().Id("Handler")
		file.Func().Id("NewHandler").Params(jen.Id("doer").Id(doerName)).Id("Handler").Block(
			jen.Return(jen.Id("&handler").Values(jen.Id("doer"))),
		)
		file.Add(doer)
		file.Type().Id("handler").Struct(jen.Id("doer").Id(doerName))
		typeMethods := make([]jen.Code, len(paths))
		for _, path := range paths {
			schemas := make([]*Schema, 0)
			params := make([]jen.Code, 0, len(path.Parameters))
			params = append(params, ctx)
			for _, p := range path.Parameters {
				p.Schema.required = true
				p.Schema.init(scope, pkg, p.Name)
				schemas = append(schemas, p.Schema)
				param := jen.Id(strcase.ToLowerCamel(p.Schema.camelName)).Add(getType(p.Schema))
				params = append(params, param)
			}

			// todo: support 204
			out := path.Out.OK.Content["application/json"]
			if out == nil && path.Out.NoContent.Content == nil {
				log.Printf("%q has no json response. Skipping", path.OperationID)
				continue
			}

			in := path.In.Content["application/json"]
			if in != nil {
				schemaIn, err := doc.getSchema(in.Schema.Ref)
				if err != nil {
					return err
				}
				schemaIn.in = true
				schemaIn.init(scope, "", path.FuncName+"In")
				schemas = append(schemas, schemaIn)
				params = append(params, jen.Id("in").Id("*"+schemaIn.camelName))
			}

			typeMeth := jen.Id(path.FuncName).Params(params...)
			structMeth := jen.Func().Params(jen.Id("h").Id("*handler")).Id(path.FuncName).Params(params...)

			var rsp, schemaOut *Schema
			if out != nil {
				schemaOut, err = doc.getSchema(out.Schema.Ref)
				if err != nil {
					return err
				}
				schemaOut.out = true
				schemaOut.init(scope, "", path.FuncName+"Out")
				rsp = getResponse(schemaOut)
			}

			if rsp != nil {
				ret := jen.List(getType(rsp), jen.Error())
				typeMeth.Parens(ret)
				structMeth.Parens(ret)
			} else {
				typeMeth.Error()
				structMeth.Error()
			}

			typeMethods = append(typeMethods, path.Comment(), typeMeth.Line())

			paramIndex := -1
			url := pathClean.ReplaceAllStringFunc(path.Path, func(_ string) string {
				paramIndex++
				switch t := path.Parameters[paramIndex].Schema.Type; t {
				case SchemaTypeInteger:
					return "%d"
				case SchemaTypeString:
					return "%s"
				default:
					panic(fmt.Sprintf("%s unexpected parameter type %s", path.OperationID, t))
				}
			})
			urlParams := make([]jen.Code, 0, len(params))
			urlParams = append(urlParams, jen.Lit(url))
			inObj := jen.Nil()
			for _, s := range schemas {
				if s.isObject() {
					inObj = jen.Id("in")
					continue
				}
				urlParams = append(urlParams, jen.Id(strcase.ToLowerCamel(s.camelName)))
			}

			outObj := jen.Id("_")
			returnErr := jen.Return(jen.Err())
			if rsp != nil {
				outObj = jen.Id("b")

				// In most cases "nil" is for error return
				// But for required scalars should be zero values
				returnErr = jen.Return(jen.Nil(), jen.Err())
				if rsp.required || rsp.Type == SchemaTypeString {
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

			block := []jen.Code{
				jen.Id("path").Op(":=").Qual("fmt", "Sprintf").Call(urlParams...),
				jen.List(outObj, jen.Err()).Op(":=").Id("h.doer.Do").Call(
					jen.Id("ctx"),
					jen.Lit(path.OperationID),
					jen.Lit(path.Method),
					jen.Id("path"),
					inObj,
				),
			}

			if rsp == nil {
				block = append(block, jen.Return(jen.Err()))
			} else {
				out := jen.Id("out")
				if rsp.camelName != schemaOut.camelName {
					out.Dot(strcase.ToCamel(rsp.name))
				}

				block = append(
					block,
					jen.Id("out").Op(":=").New(jen.Id(schemaOut.camelName)),
					jen.Err().Op("=").Qual("encoding/json", "Unmarshal").Call(jen.Id("b"), jen.Id("out")),
					jen.If(jen.Err().Op("!=").Nil()).Block(returnErr),
					jen.Return(out, jen.Nil()),
				)
			}

			file.Add(structMeth.Block(block...))
		}

		keys := lo.Keys(scope)
		slices.Sort(keys)
		for _, k := range keys {
			err = writeStruct(file, scope[k])
			if err != nil {
				return err
			}
		}

		dirPath := filepath.Join(handlerDir, fileName)
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return err
		}

		handler.Interface(typeMethods...)
		err = file.Save(filepath.Join(dirPath, fileName+".go"))
		if err != nil {
			return err
		}

		pkgName := filepath.Join(genImport, handlerDir, fileName)
		clientFields = append(clientFields, jen.Id(pkg).Qual(pkgName, "Handler"))
		clientValues[jen.Id(pkg)] = jen.Qual(pkgName, "NewHandler").Call(jen.Id("doer"))
	}

	client := jen.NewFile(clientDir)
	client.HeaderComment(generatedHeader)
	client.Add(doer)
	client.Func().Id("newClient").Params(jen.Id("doer").Id(doerName)).Id("*Client").Block(
		jen.Return(jen.Id("&Client").Values(clientValues)),
	)
	client.Type().Id("Client").Struct(clientFields...)
	return client.Save(filepath.Join(clientDir, "client.go"))
}

var reMakesSense = regexp.MustCompile(`\w`)

func writeStruct(f *jen.File, s *Schema) error {
	if s.isMap() {
		return nil
	}

	if s.isEnum() {
		kind := getScalarType(s)
		o := f.Type().Id(s.camelName)
		o.Add(kind)
		enums := make([]jen.Code, len(s.Enum))
		values := make([]jen.Code, len(s.Enum))
		for _, e := range s.Enum {
			literal := fmt.Sprint(e)
			if !reMakesSense.MatchString(literal) {
				continue
			}

			constant := s.camelName + strcase.ToCamel(literal)

			// KafkaMirror ReplicationPolicyClassType makes bad generated name
			if strings.HasPrefix(literal, "org.apache.kafka.connect.mirror.") {
				constant = s.camelName + literal[32:len(literal)-17]
			}

			// OpenSearch HealthType has value "red*"
			if strings.HasSuffix(literal, "*") {
				constant += "Asterisk"
			}
			enums = append(enums, jen.Id(constant).Op(s.camelName).Op("=").Lit(literal))
			values = append(values, jen.Lit(literal))
		}

		if len(enums) == 0 {
			return nil
		}

		o.Line().Const().Defs(enums...)

		if !s.isOut() {
			o.Line().Func().Id(s.camelName + "Choices").Params().Index().Add(kind).Block(
				jen.Return(jen.Index().Add(kind).Values(values...)),
			)
		}
		return nil
	}

	if !s.isObject() {
		return nil
	}

	fields := make([]jen.Code, 0, len(s.properties))
	for _, p := range s.properties {
		field := jen.Id(strcase.ToCamel(p.name)).Add(getType(p))
		js := p.name
		if !p.required {
			js += ",omitempty"
		}
		fields = append(fields, field.Tag(map[string]string{"json": js}))
	}

	f.Type().Id(s.camelName).Struct(fields...)
	return nil
}

func getResponse(s *Schema) *Schema {
	switch len(s.Properties) {
	case 1:
		for _, p := range s.Properties {
			return p
		}
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

var reLower = regexp.MustCompile(`^[a-z]+`)

func trimPrefix(p, s string) string {
	l := strings.TrimPrefix(strings.ToLower(s), strings.ToLower(p))
	return reLower.ReplaceAllString(s[len(s)-len(l):], "")
}
