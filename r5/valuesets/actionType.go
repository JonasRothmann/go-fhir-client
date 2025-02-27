// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/action-type
*/type ActionType string

const (
	// The action is to create a new resource.
	ActionTypeCreate ActionType = "create"
	// The action is to update an existing resource.
	ActionTypeUpdate ActionType = "update"
	// The action is to remove an existing resource.
	ActionTypeRemove ActionType = "remove"
	// The action is to fire a specific event.
	ActionTypeFireEvent ActionType = "fire-event"
)
