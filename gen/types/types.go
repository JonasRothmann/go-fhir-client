package types

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/stoewer/go-strcase"
)

var datatypesToJen = Merge(
	otherDatatypesToJen,
	baseDataTypesToJen,
	complexDatatypesToJen,
	metadataTypesToJen,
	primitiveTypesToJen,
	specialDatatypesToJen,
)

func AddTypeIdentifier(statement *jen.Statement, elem ElementDefinition) {
	if len(elem.Type) == 0 {
		statement.Id("interface{}")
		return
	}

	code := elem.Type[0].Code
	targetProfiles := elem.Type[0].TargetProfile

	if code == "BackboneElement" || code == "Element" {
		structName, fieldName := NamesFromPath(elem.Path)
		statement.Id(strings.Join([]string{structName, fieldName}, ""))

		return
	}

	if datatype, ok := datatypesToJen[code]; ok {
		var target *string
		if targetProfiles != nil && len(targetProfiles) > 0 {
			// TODO: support multiple generics
			targetURL := targetProfiles[0]
			target = Ptr(strings.TrimPrefix(targetURL, "http://hl7.org/fhir/StructureDefinition/"))
		}

		err := datatype.Jen(statement, target)
		if err != nil {

		}
	} else {
		statement.Id(code)
	}
}

func NamesFromPath(p string) (structName string, fieldName string) {
	p = strings.ReplaceAll(p, "[x]", "")
	parts := strings.Split(p, ".")

	for _, part := range parts[:len(parts)-1] {
		structName += strcase.UpperCamelCase(part)
	}
	fieldName = parts[len(parts)-1]

	return strcase.UpperCamelCase(structName), strcase.UpperCamelCase(fieldName)
}

type Jenable interface {
	Jen(statement *jen.Statement, generic *string) error
}

const (
	dataTypesPath = "github.com/jonasrothmann/go-fhir-client/generated/datatypes"
)

type GenericQualType struct {
	Path string
	Name string
}

func NewGenericQualType(path string, name string) GenericQualType {
	return GenericQualType{Path: path, Name: name}
}

func (q GenericQualType) Jen(statement *jen.Statement, generic *string) error {
	if generic == nil {
		return fmt.Errorf("Expected generic from: path: %s, name: %s", q.Path, q.Name)
	}

	statement.Qual(q.Path, q.Name).Types(jen.Id(*generic))

	return nil
}

type QualType struct {
	Path string
	Name string
}

func NewQualType(path string, name string) QualType {
	return QualType{Path: path, Name: name}
}

func (q QualType) Jen(statement *jen.Statement, _ *string) error {
	statement.Qual(q.Path, q.Name)

	return nil
}

type BuiltInType string

func (b BuiltInType) Jen(statement *jen.Statement, _ *string) error {
	statement.Id(string(b))

	return nil
}

// TODO: move following stuff diff place:
func Keys[T map[K]V, K comparable, V any](hashmap T) []K {
	keys := make([]K, len(hashmap))
	for key := range hashmap {
		keys = append(keys, key)
	}
	return keys
}

func Merge[T map[K]V, K comparable, V any](hashmaps ...T) T {
	result := T{}

	for _, hashmap := range hashmaps {
		for k, v := range hashmap {
			result[k] = v
		}
	}

	return result
}

func Ptr[T any](v T) *T {
	return &v
}
