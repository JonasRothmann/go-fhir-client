package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/dave/jennifer/jen"
	fhirclient "github.com/jonasrothmann/go-fhir-client"
	"github.com/jonasrothmann/go-fhir-client/custom"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
	"github.com/valyala/fastjson"
	"golang.org/x/sync/errgroup"
)

var errNotEnoughElements = errors.New("not enough elements")

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

var (
	cwd          = must(os.Getwd())
	outDir       = path.Join(cwd, "../r5")
	jsonDir      = path.Join(cwd, "fhir")
	customDir    = path.Join(cwd, "../custom")
	valuesetsDir = path.Join(outDir, "valuesets")
	dirs         = []string{outDir, jsonDir, valuesetsDir}
)

var fhirTypes = jenTypeIdentifiersFromMap(
	custom.PrimitiveTypes,
	map[string]any{
		"CodeableReference": custom.CodeableReference[fhirclient.Resource]{},
		"Reference":         custom.Reference[fhirclient.Resource]{},
		"Canonical":         custom.Canonical[any](""),
	},
)

func main() {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal().Err(err).Msg("Failed to create directory")
		}
	}

	var (
		structureDefinitions = make(map[string]StructureDefinition)
		valuesets            = make(map[string]ValueSet)
	)

	// Walk through the JSON directory.
	err := filepath.Walk(jsonDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Only process .json files.
		if info.IsDir() || filepath.Ext(filePath) != ".json" {
			return nil
		}

		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		// Quickly check the resource type.
		var meta struct {
			ResourceType string `json:"resourceType"`
		}
		if err := json.Unmarshal(data, &meta); err != nil {
			// Skip files that do not look like FHIR resources.
			return nil
		}

		switch meta.ResourceType {
		case "StructureDefinition":
			var sd StructureDefinition
			if err := json.Unmarshal(data, &sd); err != nil {
				return fmt.Errorf("failed to unmarshal StructureDefinition in %s: %w", filePath, err)
			}
			if sd.Name != nil {
				structureDefinitions[*sd.Name] = sd
				//log.Debug().Str("file", filePath).Str("name", *sd.Name).Msg("Parsed StructureDefinition")
			}
		case "Bundle":
			var bundle Bundle
			if err := json.Unmarshal(data, &bundle); err != nil {
				return fmt.Errorf("failed to unmarshal StructureDefinition in %s: %w", filePath, err)
			}
			if len(bundle.Entry) == 0 {
				log.Fatal().Msgf("expected bundle %s to have entries", *bundle.Id)
			}
			for _, entry := range bundle.Entry {
				if entry.Resource == nil {
					log.Fatal().Msgf("expected bundle entry %s to have resource", *bundle.Id)
					continue
				}

				resourceType := fastjson.GetString(*entry.Resource, "resourceType")
				switch resourceType {
				case "StructureDefinition":
					var sd StructureDefinition
					if err := json.Unmarshal(*entry.Resource, &sd); err != nil {
						return fmt.Errorf("failed to unmarshal StructureDefinition in %s: %w", filePath, err)
					}

					if sd.Name != nil {
						structureDefinitions[*sd.Name] = sd
						//log.Debug().Str("file", filePath).Str("name", *sd.Name).Msg("Parsed StructureDefinition")
					}
				case "ValueSet":
					var vs ValueSet
					if err := json.Unmarshal(*entry.Resource, &vs); err != nil {
						return fmt.Errorf("failed to unmarshal StructureDefinition in %s: %w", filePath, err)
					}

					if vs.Name != nil {
						valuesets[*vs.Name] = vs
						//log.Debug().Str("file", filePath).Str("name", *vs.Name).Msg("Parsed ValueSet")
					}
				}
			}
		default:
			// Ignore other types for now.
		}

		return nil
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error walking files")
	}

	skippedSds := []string{}
	for name, sd := range structureDefinitions {
		if _, isDefined := fhirTypes[name]; isDefined {
			delete(structureDefinitions, name)
			skippedSds = append(skippedSds, name)
		} else {
			if sd.Kind == StructureDefinitionKindPrimitive {
				log.Fatal().Msgf("expected primitive %s to be defined in overwrites/primivites.go", name)
			}

			statement := &jen.Statement{}
			statement.Id(name)

			fhirTypes[name] = GenStatement{
				statement: statement,
			}
		}
	}
	log.Debug().Msgf("following structs are already defined - skipping: %s", strings.Join(skippedSds, ", "))

	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)
	//g.SetLimit(25)

	for name, sd := range structureDefinitions {
		g.Go(func() error {
			err := generateStructFromStructureDefinition(sd)

			if err != nil {
				if errors.Is(err, errNotEnoughElements) {
					log.Error().Err(err).Msgf("Failed to generate struct for %s", name)
				} else {
					return errors.Wrapf(err, "Failed to generate struct for %s", name)
				}
			}

			return nil
		})
	}

	for name, vs := range valuesets {
		g.Go(func() error {
			f := jen.NewFile("valuesets")
			f.HeaderComment("GENERATED CODE – DO NOT EDIT!")

			// TODO: Fix performance bottleneck
			now := time.Now()
			e, err := vs.generateEnum(ctx)
			if err != nil {
				return errors.Wrapf(err, "Failed to generate valueset for %s", name)
			}
			log.Debug().Msgf("elapsed: %s", time.Now().Sub(now).String())

			statement := e.generateStatement()
			f.Add(statement)

			filename := path.Join(valuesetsDir, fmt.Sprintf("%s.go", strcase.LowerCamelCase(name)))
			if err := f.Save(filename); err != nil {
				return errors.Wrapf(err, "error saving file %s", filename)
			}
			log.Info().Str("valueset", filename).Msg("Generated Go source")

			return nil
		})
	}

	if err = g.Wait(); err != nil {
		log.Error().Err(err).Msg("errgroup failed")
	}

	// might not be needed
	/*
		if err = filepath.Walk(customDir, func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return errors.Wrap(err, "could not walk filepath")
			}

			bytes, err := os.ReadFile(filePath)
			if err != nil {
				return errors.Wrapf(err, "could not read file %s", filePath)
			}

			str := strings.Replace(string(bytes), "package custom", "package v5", 1)

			if err = os.WriteFile(path.Join(outDir, info.Name()), []byte(str), os.ModePerm); err != nil {
				return errors.Wrapf(err, "could not write file %s", path.Join(outDir, info.Name()))
			}

			return nil
		}); err != nil {
			log.Fatal().Err(err).Msg("Failed to copy custom into v5")
			} */
}

func generateStructFromStructureDefinition(sd StructureDefinition) error {
	var (
		dir    string
		f      *jen.File
		prefix string
	)

	if sd.ResourceType == "StructureDefinition" {
		switch sd.Kind {
		case "resource":
			prefix = "resource"
		default:
			prefix = "datatype"
		}
	} else if sd.ResourceType == "ValueSet" {
		prefix = "valueset"
	} else {
		log.Fatal().Msg("Unexpected")
	}

	f = jen.NewFile("r5")
	dir = outDir

	f.HeaderComment("GENERATED CODE – DO NOT EDIT!")
	if sd.Url != nil {
		f.Commentf("%s", *sd.Url)
	}

	if len(sd.Snapshot.Element) < 2 {
		return errors.Wrapf(errNotEnoughElements, "in element %s", *sd.Name)
	}

	structFields := map[string][]jen.Code{}
	additionalStatements := []*jen.Statement{}

	elems := []ElementDefinition{}
	for _, elem := range sd.Snapshot.Element[1:] {
		if strings.Contains(*elem.Id, ":") {
			log.Debug().Msgf("skipped id: %s because it contains colon", *elem.Id)
			continue
		}

		if strings.HasSuffix(elem.Path, "[x]") {
			path, _ := strings.CutSuffix(elem.Path, "[x]")
			for _, t := range elem.Type {
				newElem := elem
				newElem.Path = path + strcase.UpperCamelCase(string(t.Code))
				newElem.Type = []ElementDefinitionType{t}
				newElem.Min = nil

				elems = append(elems, newElem)
			}
		} else {
			elems = append(elems, elem)
		}
	}

	for _, elem := range elems {
		structName, fieldName := namesFromPath(*sd.Name, elem.Path, elem.SliceName)
		tagFieldName := strcase.LowerCamelCase(fieldName)

		statement := jen.Comment(generateElementComment(elem)).Line().Id(fieldName)

		if elem.Max != nil && *elem.Max == "*" {
			statement.Op("[]")
		} else if elem.Min == nil || *elem.Min == 0 {
			statement.Op("*")
		}

		// Append the type.
		additionalStatements = slices.Concat(additionalStatements, addTypeIdentifier(statement, elem, *sd.Name, structName, fieldName))

		// Set JSON and BSON struct tags.
		var min uint32
		if elem.Min != nil {
			min = *elem.Min
		}
		if min == 0 {
			statement.Tag(map[string]string{
				"json": tagFieldName + ",omitempty",
				//"bson": tagFieldName + ",omitempty",
			})
		} else {
			statement.Tag(map[string]string{
				"json": tagFieldName,
			})
		}

		structFields[structName] = append(structFields[structName], statement)
	}

	fileName := ""
	if sd.Name != nil {
		fileName = *sd.Name
	} else {
		return fmt.Errorf("StructureDefinition missing name")
	}

	for structName, fields := range structFields {
		f.Type().Id(structName).Struct(fields...).Line()
	}

	mainStructName := strcase.UpperCamelCase(*sd.Name)
	mainStructParam := strings.ToLower(string(mainStructName[0]))
	otherStructName := fmt.Sprintf("Other%s", mainStructName)

	f.Add(
		jen.Type().Id(otherStructName).Id(mainStructName).Line(),
	)

	for _, statement := range additionalStatements {
		f.Add(*statement...)
	}

	implementations := map[string]string{
		"ResourceType": *sd.Name,
		//"Short": *sd.Short,
		//"Definition": *sd.Short,
	}
	for key, val := range implementations {
		f.Func().Params(
			jen.Id(mainStructParam).Id(mainStructName),
		).Id(key).Params().String().Block(
			jen.Return(jen.Lit(val)),
		)
	}

	f.Add(
		jen.Line(),
		genFuncMarshalJSON(mainStructParam, mainStructName, otherStructName),
		jen.Line().Line(),
		genFuncUnmarshal(mainStructName),
	)

	filename := path.Join(dir, fmt.Sprintf("%s_%s.go", prefix, strcase.LowerCamelCase(fileName)))
	if err := f.Save(filename); err != nil {
		return fmt.Errorf("error saving file %s: %w", filename, err)
	}
	//log.Info().Str("file", filename).Msg("Generated Go source")
	return nil
}

func addTypeIdentifier(statement *jen.Statement, elem ElementDefinition, baseName, structName, fieldName string) (additionalStatements []*jen.Statement) {
	if len(elem.Type) == 0 {
		statement.Id("interface{}")
		return
	}

	code, found := strings.CutPrefix(string(elem.Type[0].Code), "http://hl7.org/fhirpath/System.")
	if found {
		code = strcase.LowerCamelCase(code)
	}

	targetProfiles := elem.Type[0].TargetProfile

	if code == "BackboneElement" || code == "Element" {
		if strings.HasPrefix(structName, "ElementDefinition") {
			log.Debug().Msgf("%s", structName)
		}

		structName, fieldName := namesFromPath(baseName, elem.Path, elem.SliceName)
		statement.Id(strings.Join([]string{structName, fieldName}, ""))

		return
	}

	if datatype, ok := fhirTypes[code]; ok {
		var targets []string
		if targetProfiles != nil && len(targetProfiles) > 0 {
			targets = []string{}
			for _, url := range targetProfiles {
				str := strings.TrimPrefix(string(url), "http://hl7.org/fhir/StructureDefinition/")

				// a blacklist is def not the way to go
				if str == "vitalsigns" {
					continue
				}

				// TODO: Handle this more gracefully than just continuing
				if slices.Contains(targets, str) {
					log.Debug().Msgf("%s already in targets", str)
					continue
				}

				targets = append(targets, str)
			}
		}

		if len(*datatype.statement) == 0 {
			log.Fatal().Msgf("expected datatype len > 0 on code %s", code)
		}

		statement.Add(*datatype.statement...)

		if len(targets) == 0 {
			if datatype.generic != nil {
				statement.Types(
					statementFromType(*datatype.generic),
				)
			}
		} else if len(targets) == 1 {
			log.Debug().Msg(targets[0])
			statement.Types(
				jen.Id(targets[0]),
			)
		} else if len(targets) > 1 {
			enumName := structName + fieldName
			interfaceFuncName := fmt.Sprintf("Is_%s", enumName)

			additionalStatement := jen.Type().Id(enumName).Interface(
				jen.Qual("github.com/jonasrothmann/go-fhir-client", "Resource"),
				jen.Line(),
				jen.Id(interfaceFuncName).Params(),
			).Line()

			for _, target := range targets {
				paramName := strings.ToLower(string(target[0]))

				additionalStatement.Func().Params(
					jen.Id(paramName).Id(target),
				).Id(
					interfaceFuncName,
				).Params().Block().Line()
			}

			additionalStatements = append(additionalStatements, additionalStatement)

			if datatype.generic != nil {
				statement.Types(
					jen.Id(enumName),
				)
			}
		}
	} else {
		statement.Id(code)
	}

	return additionalStatements
}

func isGeneric(group *jen.Group) bool {
	// TODO: this is such a hacky way to do this lol
	t := reflect.ValueOf(*group)
	return strings.HasPrefix(t.String(), "{name:types")
}
func namesFromPath(id, p string, sliceName *string) (structName string, fieldName string) {
	p = strings.ReplaceAll(p, "[x]", "")
	parts := strings.Split(p, ".")

	parts[0] = id

	for _, part := range parts[:len(parts)-1] {
		structName += strcase.UpperCamelCase(part)
	}
	fieldName = strcase.UpperCamelCase(parts[len(parts)-1])

	if sliceName != nil {
		structName += strcase.UpperCamelCase(*sliceName)
	}

	return strcase.UpperCamelCase(structName), fieldName
}

func generateElementComment(elem ElementDefinition) (result string) {
	result = *elem.Short // + "\n"

	return result

	// TODO: Decide if we return constraints
	/*for i, constraint := range elem.Constraint {
		result += fmt.Sprintf("+ %s: %s", constraint.Severity, constraint.Human)
		if i < len(elem.Constraint)-1 {
			result += "\n"
		}
	}*/

	//return fmt.Sprintf("/*%s */", result)
}

// must is a helper that panics if an error occurs.
func must[T any](val T, err error) T {
	if err != nil {
		log.Fatal().Err(err).Msg("must function error")
	}
	return val
}

type empty struct{}

type GenStatement struct {
	statement *jen.Statement
	generic   *string
}

func jenTypeIdentifiersFromMap(inputs ...map[string]any) map[string]GenStatement {
	input := merge(inputs...)

	result := make(map[string]GenStatement, len(input))
	for key, val := range input {
		g := GenStatement{
			statement: &jen.Statement{},
		}

		t := reflect.TypeOf(val)

		if t.PkgPath() == "" {
			g.statement.Id(t.Name())
		} else {
			parts := strings.Split(t.Name(), "[")
			name := t.Name()

			if len(parts) > 1 {
				generic, isGeneric := strings.CutSuffix(parts[1], "]")
				if !isGeneric {
					log.Fatal().Str("name", name).Msg("unexpected missing generic")
				}
				g.generic = &generic

				name = parts[0]
			}

			if t.PkgPath() == reflect.TypeOf(fhirclient.Empty{}).PkgPath()+"/r5" {
				g.statement.Id(name)
			} else {
				g.statement.Qual(t.PkgPath(), name)
			}
		}

		result[key] = g
	}

	return result
}

func statementFromType(t string) *jen.Statement {
	parts := strings.Split(t, ".")
	path := strings.Join(parts[:len(parts)-1], ".")
	name := parts[len(parts)-1]

	if name == "interface {}" {
		name = "any"
	}

	return jen.Qual(path, name)
}

// TODO: move following stuff diff place:
func keys[T map[K]V, K comparable, V any](hashmap T) []K {
	keys := make([]K, len(hashmap))
	for key := range hashmap {
		keys = append(keys, key)
	}
	return keys
}

func merge[T map[K]V, K comparable, V any](hashmaps ...T) T {
	result := T{}

	for _, hashmap := range hashmaps {
		for k, v := range hashmap {
			result[k] = v
		}
	}

	return result
}

func ptr[T any](v T) *T {
	return &v
}

func genFuncMarshalJSON(paramName, structName, otherStructName string) jen.Code {
	return jen.Func().Params(jen.Id(paramName).Id(structName)).Id("MarshalJSON").Params().Params(jen.Index().Id("byte"), jen.Id("error")).Block(jen.Return().Qual("encoding/json", "Marshal").Call(
		jen.Struct(jen.Id(otherStructName), jen.Id("ResourceType").String().Tag(map[string]string{
			"json": "resourceType,omitempty",
		},
		)).Values(jen.Id(otherStructName).Op(":").Id(otherStructName).Call(jen.Id(paramName)), jen.Id("ResourceType").Op(":").Id(paramName).Dot("ResourceType").Call())))
}
func genFuncUnmarshal(structName string) jen.Code {
	return jen.Func().Id(fmt.Sprintf("Unmarshal%s", structName)).Params(jen.Id("b").Index().Byte()).Params(jen.Id("res").Id(structName), jen.Id("err").Error()).Block(jen.If(jen.Id("err").Op(":=").Qual("encoding/json", "Unmarshal").Call(jen.Id("b"), jen.Op("&").Id("res")), jen.Id("err").Op("!=").Id("nil")).Block(jen.Return().List(jen.Id("res"), jen.Id("err"))), jen.Return().List(jen.Id("res"), jen.Id("nil")))
}
