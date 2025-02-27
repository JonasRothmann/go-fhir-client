// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/condition-precondition-type
*/type ConditionPreconditionType string

const (
	// The observation is very sensitive for the condition, but may also indicate other conditions.
	ConditionPreconditionTypeSensitive ConditionPreconditionType = "sensitive"
	// The observation is very specific for this condition, but not particularly sensitive.
	ConditionPreconditionTypeSpecific ConditionPreconditionType = "specific"
)
