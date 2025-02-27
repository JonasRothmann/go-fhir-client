// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/search-param-type
*/type SearchParamType string

const (
	// Search parameter SHALL be a number (a whole number, or a decimal).
	SearchParamTypeNumber SearchParamType = "number"
	// Search parameter is on a date/time. The date format is the standard XML format, though other formats may be supported.
	SearchParamTypeDate SearchParamType = "date"
	// Search parameter is a simple string, like a name part. Search is case-insensitive and accent-insensitive. May match just the start of a string. String parameters may contain spaces.
	SearchParamTypeString SearchParamType = "string"
	// Search parameter on a coded element or identifier. May be used to search through the text, display, code and code/codesystem (for codes) and label, system and key (for identifier). Its value is either a string or a pair of namespace and value, separated by a "|", depending on the modifier used.
	SearchParamTypeToken SearchParamType = "token"
	// A reference to another resource (Reference or canonical).
	SearchParamTypeReference SearchParamType = "reference"
	// A composite search parameter that combines a search on two values together.
	SearchParamTypeComposite SearchParamType = "composite"
	// A search parameter that searches on a quantity.
	SearchParamTypeQuantity SearchParamType = "quantity"
	// A search parameter that searches on a URI (RFC 3986).
	SearchParamTypeUri SearchParamType = "uri"
	// Special logic applies to this parameter per the description of the search parameter.
	SearchParamTypeSpecial SearchParamType = "special"
)
