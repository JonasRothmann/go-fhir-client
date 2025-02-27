// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/conceptmap-property-type
*/type ConceptMapPropertyType string

const (
	// The property  value is a code defined in an external code system. This may be used for translations, but is not the intent.
	ConceptMapPropertyTypeCoding ConceptMapPropertyType = "Coding"
	// The property value is a string.
	ConceptMapPropertyTypeString ConceptMapPropertyType = "string"
	// The property value is an integer (often used to assign ranking values to concepts for supporting score assessments).
	ConceptMapPropertyTypeInteger ConceptMapPropertyType = "integer"
	// The property value is a boolean true | false.
	ConceptMapPropertyTypeBoolean ConceptMapPropertyType = "boolean"
	// The property is a date or a date + time.
	ConceptMapPropertyTypeDateTime ConceptMapPropertyType = "dateTime"
	// The property value is a decimal number.
	ConceptMapPropertyTypeDecimal ConceptMapPropertyType = "decimal"
	// The property value is a code as defined in the CodeSystem in ConceptMap.property.system.
	ConceptMapPropertyTypeCode ConceptMapPropertyType = "code"
)
