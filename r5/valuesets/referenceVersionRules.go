// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/reference-version-rules
*/type ReferenceVersionRules string

const (
	// The reference may be either version independent or version specific.
	ReferenceVersionRulesEither ReferenceVersionRules = "either"
	// The reference must be version independent.
	ReferenceVersionRulesIndependent ReferenceVersionRules = "independent"
	// The reference must be version specific.
	ReferenceVersionRulesSpecific ReferenceVersionRules = "specific"
)
