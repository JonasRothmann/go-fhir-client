// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/structure-definition-kind
*/type StructureDefinitionKind string

const (
	// A primitive type that has a value and an extension. These can be used throughout complex datatype, Resource and extension definitions. Only the base specification can define primitive types.
	StructureDefinitionKindPrimitiveDataType StructureDefinitionKind = "primitive-type"
	// A  complex structure that defines a set of data elements that is suitable for use in 'resources'. The base specification defines a number of complex types, and other specifications can define additional types. These structures do not have a maintained identity.
	StructureDefinitionKindComplexDataType StructureDefinitionKind = "complex-type"
	// A 'resource' - a directed acyclic graph of elements that aggregrates other types into an identifiable entity. The base FHIR resources are defined by the FHIR specification itself but other 'resources' can be defined in additional specifications (though these will not be recognized as 'resources' by the FHIR specification; i.e. they do not get end-points etc., or act as the targets of references in FHIR defined resources - though other specifications can treat them this way).
	StructureDefinitionKindResource StructureDefinitionKind = "resource"
	// A pattern or a template that is not intended to be a real resource or complex type.
	StructureDefinitionKindLogical StructureDefinitionKind = "logical"
)
