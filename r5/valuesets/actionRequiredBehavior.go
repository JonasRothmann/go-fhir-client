// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-required-behavior
*/type ActionRequiredBehavior string

const (
	// An action with this behavior must be included in the actions processed by the end user; the end user SHALL NOT choose not to include this action.
	ActionRequiredBehaviorMust ActionRequiredBehavior = "must"
	// An action with this behavior may be included in the set of actions processed by the end user.
	ActionRequiredBehaviorCould ActionRequiredBehavior = "could"
	// An action with this behavior must be included in the set of actions processed by the end user, unless the end user provides documentation as to why the action was not included.
	ActionRequiredBehaviorMustUnlessDocumented ActionRequiredBehavior = "must-unless-documented"
)
