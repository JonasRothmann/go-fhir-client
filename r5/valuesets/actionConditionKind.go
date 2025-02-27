// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-condition-kind
*/type ActionConditionKind string

const (
	// The condition describes whether or not a given action is applicable.
	ActionConditionKindApplicability ActionConditionKind = "applicability"
	// The condition is a starting condition for the action.
	ActionConditionKindStart ActionConditionKind = "start"
	// The condition is a stop, or exit condition for the action.
	ActionConditionKindStop ActionConditionKind = "stop"
)
