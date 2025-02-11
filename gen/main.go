package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/jonasrothmann/go-fhir-client/gen/overwrites"
	"github.com/jonasrothmann/go-fhir-client/gen/types"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stoewer/go-strcase"
	"github.com/valyala/fastjson"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

var (
	cwd     = must(os.Getwd())
	outDir  = path.Join(cwd, "../v5")
	jsonDir = path.Join(cwd, "fhir")
	//valuesetsDir = path.Join(outDir, "valuesets")
	dirs = []string{outDir, jsonDir}
)

// TODO: auto import from overwrites
var fhirTypes = jenTypeIdentifiersFromMap(
	overwrites.PrimitiveTypes,
)

func main() {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			log.Fatal().Err(err).Msg("Failed to create directory")
		}
	}

	// Initialize our collection.
	collections := types.ResourceCollections{
		StructureDefinitions: make(map[string]types.StructureDefinition),
	}

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

		log.Debug().Msg(meta.ResourceType)

		switch meta.ResourceType {
		case "StructureDefinition":
			var sd types.StructureDefinition
			if err := json.Unmarshal(data, &sd); err != nil {
				return fmt.Errorf("failed to unmarshal StructureDefinition in %s: %w", filePath, err)
			}
			if sd.Name != nil {
				collections.StructureDefinitions[*sd.Name] = sd
				log.Debug().Str("file", filePath).Str("name", *sd.Name).Msg("Parsed StructureDefinition")
			}
		case "Bundle":
			var bundle types.Bundle
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
				if resourceType == "StructureDefinition" {
					var sd types.StructureDefinition
					if err := json.Unmarshal(*entry.Resource, &sd); err != nil {
						return fmt.Errorf("failed to unmarshal StructureDefinition in %s: %w", filePath, err)
					}

					log.Debug().Msg(*sd.Name)

					if sd.Name != nil {
						collections.StructureDefinitions[*sd.Name] = sd
						log.Debug().Str("file", filePath).Str("name", *sd.Name).Msg("Parsed StructureDefinition")
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

	log.Debug().Msgf("%+v", fhirTypes)

	for name, sd := range collections.StructureDefinitions {
		if _, isDefined := fhirTypes[name]; isDefined {
			log.Debug().Msgf("%s is already defined - skipping", name)
		} else {
			if sd.Kind == types.StructureDefinitionKindPrimitive {
				log.Fatal().Msgf("expected primitve %s to be defined in overwrites/primivites.go", name)
			}

			statement := &jen.Statement{}
			statement.Qual(sd.Kind.Path(), name)

			fhirTypes[name] = statement
		}
	}

	// Generate Go files from the collected StructureDefinitions.
	for name, sd := range collections.StructureDefinitions {
		if err := generateStructFromStructureDefinition(sd); err != nil {
			log.Error().Err(err).Msgf("Failed to generate struct for %s", name)
		}
	}
}

func generateStructFromStructureDefinition(sd types.StructureDefinition) error {
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

	f = jen.NewFile("v5")
	dir = outDir

	f.HeaderComment("GENERATED CODE – DO NOT EDIT!")
	if sd.Url != nil {
		f.Commentf("%s", *sd.Url)
	}

	if len(sd.Snapshot.Element) < 2 {
		return fmt.Errorf("not enough elements in snapshot for %s", *sd.Name)
	}

	structFields := map[string][]jen.Code{}
	additionalStatements := []*jen.Statement{}

	for _, elem := range sd.Snapshot.Element[1:] {
		structName, fieldName := types.NamesFromPath(elem.Path)
		tagFieldName := strcase.LowerCamelCase(fieldName)

		statement := jen.Id(fieldName)

		if elem.Max != nil && *elem.Max == "*" {
			statement.Op("[]")
		} else if elem.Min == nil || *elem.Min == 0 {
			statement.Op("*")
		}

		// Append the type.
		additionalStatements = addTypeIdentifier(statement, elem, structName)

		// Set JSON and BSON struct tags.
		min := 0
		if elem.Min != nil {
			min = *elem.Min
		}
		if min == 0 {
			statement.Tag(map[string]string{
				"json": tagFieldName + ",omitempty",
				"bson": tagFieldName + ",omitempty",
			})
		} else {
			statement.Tag(map[string]string{
				"json": tagFieldName,
				"bson": tagFieldName,
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
		f.Type().Id(structName).Struct(fields...)
	}

	for _, statement := range additionalStatements {
		f.Add(*statement...)
	}

	mainStructName := strcase.UpperCamelCase(*sd.Name)
	mainStructParam := strings.ToLower(string(mainStructName[0]))
	f.Func().Params(
		jen.Id(mainStructParam).Id(mainStructName),
	).Id("ResourceType").Params().String().Block(
		jen.Return(jen.Lit(*sd.Name)),
	)

	filename := path.Join(dir, fmt.Sprintf("%s_%s.go", prefix, strcase.LowerCamelCase(fileName)))
	if err := f.Save(filename); err != nil {
		return fmt.Errorf("error saving file %s: %w", filename, err)
	}
	log.Info().Str("file", filename).Msg("Generated Go source")
	return nil
}

func addTypeIdentifier(statement *jen.Statement, elem types.ElementDefinition, structName string) (additionalStatements []*jen.Statement) {
	if len(elem.Type) == 0 {
		statement.Id("interface{}")
		return
	}

	code, found := strings.CutPrefix(elem.Type[0].Code, "http://hl7.org/fhirpath/System.")
	if found {
		code = strcase.LowerCamelCase(code)
	}

	targetProfiles := elem.Type[0].TargetProfile

	if code == "BackboneElement" || code == "Element" {
		structName, fieldName := types.NamesFromPath(elem.Path)
		statement.Id(strings.Join([]string{structName, fieldName}, ""))

		return
	}

	if datatype, ok := fhirTypes[code]; ok {
		var targets []string
		if targetProfiles != nil && len(targetProfiles) > 0 {
			targets = make([]string, len(targetProfiles))
			for i, url := range targetProfiles {
				targets[i] = strings.TrimPrefix(url, "http://hl7.org/fhir/StructureDefinition/")
			}
		}

		if len(*datatype) == 0 {
			log.Fatal().Msgf("expected datatype len > 0 on code %s", code)
		}

		statement.Add(*datatype...)

		if len(targets) > 0 {
			/*
				type PersonLinkTarget interface {
					fhirclient.Resource

					Is_PersonLinkTarget()
				}
				func (p Patient) Is_PersonLinkTarget() {}
				func (p Practitioner) Is_PersonLinkTarget() {}
				func (p RelatedPerson) Is_PersonLinkTarget() {}
				func (p Person) Is_PersonLinkTarget() {}
			*/

			enumName := structName + strcase.UpperCamelCase(code)

			for _, target := range targets {
				paramName := strings.ToLower(string(target[0]))

				additionalStatements = append(additionalStatements,
					jen.Func().Params(
						jen.Id(paramName).Id(target),
					).Id(
						fmt.Sprintf("Is_%s", enumName),
					).Params().Block(),
				)
			}

			statement.Types(
				jen.Id(enumName),
			)
		}
	} else {
		log.Debug().Msg(code)
		statement.Id(code)
	}

	return additionalStatements
}

func namesFromPath(p string) (structName string, fieldName string) {
	p = strings.ReplaceAll(p, "[x]", "")
	parts := strings.Split(p, ".")

	for _, part := range parts[:len(parts)-1] {
		structName += strcase.UpperCamelCase(part)
	}
	fieldName = parts[len(parts)-1]

	return strcase.UpperCamelCase(structName), strcase.UpperCamelCase(fieldName)
}

func ptr[T any](v T) *T {
	return &v
}

// must is a helper that panics if an error occurs.
func must[T any](val T, err error) T {
	if err != nil {
		log.Fatal().Err(err).Msg("must function error")
	}
	return val
}

func jenTypeIdentifiersFromMap(input map[string]any) map[string]*jen.Statement {
	result := make(map[string]*jen.Statement, len(input))
	for key, val := range input {
		statement := &jen.Statement{}

		t := reflect.TypeOf(val)

		if t.PkgPath() == "" {
			statement.Id(t.Name())
		} else {
			parts := strings.Split(t.Name(), "[")
			if len(parts) > 1 {
				// TODO: Anything to do if generics
				//generic, isGeneric := strings.CutSuffix(parts[1], "]")

				statement.Qual(t.PkgPath(), parts[0])

			} else {
				statement.Qual(t.PkgPath(), t.Name())
			}
		}

		result[key] = statement
	}

	return result
}
