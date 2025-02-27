// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/restful-interaction
  - read
  - vread
  - update
  - patch
  - delete
  - history-instance
  - history-type
  - create
  - search-type
*/type TypeRestfulInteraction string

const (
	// Read the current state of the resource.
	TypeRestfulInteractionRead TypeRestfulInteraction = "read"
	// Read the state of a specific version of the resource.
	TypeRestfulInteractionVread TypeRestfulInteraction = "vread"
	// Update an existing resource by its id (or create it if it is new).
	TypeRestfulInteractionUpdate TypeRestfulInteraction = "update"
	// Update an existing resource by posting a set of changes to it.
	TypeRestfulInteractionPatch TypeRestfulInteraction = "patch"
	// Delete a resource.
	TypeRestfulInteractionDelete TypeRestfulInteraction = "delete"
	// Create a new resource with a server assigned id.
	TypeRestfulInteractionCreate TypeRestfulInteraction = "create"
)
