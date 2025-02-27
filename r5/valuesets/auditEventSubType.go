// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://dicom.nema.org/resources/ontology/DCM
- http://dicom.nema.org/resources/ontology/DCM
- http://hl7.org/fhir/restful-interaction
*/type AuditEventSubType string

const (
	// Read the current state of the resource.
	AuditEventSubTypeRead AuditEventSubType = "read"
	// Read the state of a specific version of the resource.
	AuditEventSubTypeVread AuditEventSubType = "vread"
	// Update an existing resource by its id (or create it if it is new).
	AuditEventSubTypeUpdate AuditEventSubType = "update"
	// Update an existing resource by posting a set of changes to it.
	AuditEventSubTypePatch AuditEventSubType = "patch"
	// Delete a resource.
	AuditEventSubTypeDelete AuditEventSubType = "delete"
	// Retrieve the change history for a particular resource, type of resource, or the entire system.
	AuditEventSubTypeHistory AuditEventSubType = "history"
	// Create a new resource with a server assigned id.
	AuditEventSubTypeCreate AuditEventSubType = "create"
	// Search a resource type or all resources based on some filter criteria.
	AuditEventSubTypeSearch AuditEventSubType = "search"
	// Get a Capability Statement for the system.
	AuditEventSubTypeCapabilities AuditEventSubType = "capabilities"
	// Update, create or delete a set of resources as a single transaction.
	AuditEventSubTypeTransaction AuditEventSubType = "transaction"
	// perform a set of a separate interactions in a single http operation
	AuditEventSubTypeBatch AuditEventSubType = "batch"
	// Perform an operation as defined by an OperationDefinition.
	AuditEventSubTypeOperation AuditEventSubType = "operation"
)
