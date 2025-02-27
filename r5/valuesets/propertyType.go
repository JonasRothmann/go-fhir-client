// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/concept-property-type
*/type PropertyType string

const (
	// The property value is a code that identifies a concept defined in the code system.
	PropertyTypeCode PropertyType = "code"
	// The property  value is a code defined in an external code system. This may be used for translations, but is not the intent.
	PropertyTypeCoding PropertyType = "Coding"
	// The property value is a string.
	PropertyTypeString PropertyType = "string"
	// The property value is an integer (often used to assign ranking values to concepts for supporting score assessments).
	PropertyTypeInteger PropertyType = "integer"
	// The property value is a boolean true | false.
	PropertyTypeBoolean PropertyType = "boolean"
	// The property is a date or a date + time.
	PropertyTypeDateTime PropertyType = "dateTime"
	// The property value is a decimal number.
	PropertyTypeDecimal PropertyType = "decimal"
)
