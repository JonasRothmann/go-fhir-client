// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/restful-interaction
  - create
  - update
  - delete
*/type InteractionTrigger string

const (
	// Update an existing resource by its id (or create it if it is new).
	InteractionTriggerUpdate InteractionTrigger = "update"
	// Delete a resource.
	InteractionTriggerDelete InteractionTrigger = "delete"
	// Create a new resource with a server assigned id.
	InteractionTriggerCreate InteractionTrigger = "create"
)
