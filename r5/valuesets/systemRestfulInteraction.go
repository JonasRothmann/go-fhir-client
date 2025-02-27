// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/restful-interaction
  - transaction
  - batch
  - search-system
  - history-system
*/type SystemRestfulInteraction string

const (
	// Update, create or delete a set of resources as a single transaction.
	SystemRestfulInteractionTransaction SystemRestfulInteraction = "transaction"
	// perform a set of a separate interactions in a single http operation
	SystemRestfulInteractionBatch SystemRestfulInteraction = "batch"
)
