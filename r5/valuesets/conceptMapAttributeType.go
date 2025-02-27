// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/conceptmap-attribute-type
*/type ConceptMapAttributeType string

const (
	// The attribute value is a code defined in the code system in context.
	ConceptMapAttributeTypeCode ConceptMapAttributeType = "code"
	// The attribute value is a code defined in a code system.
	ConceptMapAttributeTypeCoding ConceptMapAttributeType = "Coding"
	// The attribute value is a string.
	ConceptMapAttributeTypeString ConceptMapAttributeType = "string"
	// The attribute value is a boolean true | false.
	ConceptMapAttributeTypeBoolean ConceptMapAttributeType = "boolean"
	// The attribute is a Quantity (may represent an integer or a decimal with no units).
	ConceptMapAttributeTypeQuantity ConceptMapAttributeType = "Quantity"
)
