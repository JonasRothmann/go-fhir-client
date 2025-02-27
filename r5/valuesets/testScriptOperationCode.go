// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/restful-interaction
*/type TestScriptOperationCode string

const (
	// Read the current state of the resource.
	TestScriptOperationCodeRead TestScriptOperationCode = "read"
	// Read the state of a specific version of the resource.
	TestScriptOperationCodeVread TestScriptOperationCode = "vread"
	// Update an existing resource by its id (or create it if it is new).
	TestScriptOperationCodeUpdate TestScriptOperationCode = "update"
	// Update an existing resource by posting a set of changes to it.
	TestScriptOperationCodePatch TestScriptOperationCode = "patch"
	// Delete a resource.
	TestScriptOperationCodeDelete TestScriptOperationCode = "delete"
	// Retrieve the change history for a particular resource, type of resource, or the entire system.
	TestScriptOperationCodeHistory TestScriptOperationCode = "history"
	// Create a new resource with a server assigned id.
	TestScriptOperationCodeCreate TestScriptOperationCode = "create"
	// Search a resource type or all resources based on some filter criteria.
	TestScriptOperationCodeSearch TestScriptOperationCode = "search"
	// Get a Capability Statement for the system.
	TestScriptOperationCodeCapabilities TestScriptOperationCode = "capabilities"
	// Update, create or delete a set of resources as a single transaction.
	TestScriptOperationCodeTransaction TestScriptOperationCode = "transaction"
	// perform a set of a separate interactions in a single http operation
	TestScriptOperationCodeBatch TestScriptOperationCode = "batch"
	// Perform an operation as defined by an OperationDefinition.
	TestScriptOperationCodeOperation TestScriptOperationCode = "operation"
)
