// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/extension-context-type
*/type ExtensionContextType string

const (
	// The context is all elements that match the FHIRPath query found in the expression.
	ExtensionContextTypeFhirpath ExtensionContextType = "fhirpath"
	// The context is any element that has an ElementDefinition.id that matches that found in the expression. This includes ElementDefinition Ids that have slicing identifiers. The full path for the element is [url]#[elementid]. If there is no #, the Element id is one defined in the base specification.
	ExtensionContextTypeElement ExtensionContextType = "element"
	// The context is a particular extension from a particular StructureDefinition, and the expression is just a uri that identifies the extension.
	ExtensionContextTypeExtension ExtensionContextType = "extension"
)
