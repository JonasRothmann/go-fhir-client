// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-precheck-behavior
*/type ActionPrecheckBehavior string

const (
	// An action with this behavior is one of the most frequent action that is, or should be, included by an end user, for the particular context in which the action occurs. The system displaying the action to the end user should consider "pre-checking" such an action as a convenience for the user.
	ActionPrecheckBehaviorYes ActionPrecheckBehavior = "yes"
	// An action with this behavior is one of the less frequent actions included by the end user, for the particular context in which the action occurs. The system displaying the actions to the end user would typically not "pre-check" such an action.
	ActionPrecheckBehaviorNo ActionPrecheckBehavior = "no"
)
