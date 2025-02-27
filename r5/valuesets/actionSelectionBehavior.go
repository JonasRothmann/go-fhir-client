// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-selection-behavior
*/type ActionSelectionBehavior string

const (
	// Any number of the actions in the group may be chosen, from zero to all.
	ActionSelectionBehaviorAny ActionSelectionBehavior = "any"
	// All the actions in the group must be selected as a single unit.
	ActionSelectionBehaviorAll ActionSelectionBehavior = "all"
	// All the actions in the group are meant to be chosen as a single unit: either all must be selected by the end user, or none may be selected.
	ActionSelectionBehaviorAllOrNone ActionSelectionBehavior = "all-or-none"
	// The end user must choose one and only one of the selectable actions in the group. The user SHALL NOT choose none of the actions in the group.
	ActionSelectionBehaviorExactlyOne ActionSelectionBehavior = "exactly-one"
	// The end user may choose zero or at most one of the actions in the group.
	ActionSelectionBehaviorAtMostOne ActionSelectionBehavior = "at-most-one"
	// The end user must choose a minimum of one, and as many additional as desired.
	ActionSelectionBehaviorOneOrMore ActionSelectionBehavior = "one-or-more"
)
