// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://dicom.nema.org/resources/ontology/DCM
- http://terminology.hl7.org/CodeSystem/audit-event-type
*/type AuditEventID string

const (
	// Audit Event: Execution of a RESTful operation as defined by FHIR.
	AuditEventIDRest AuditEventID = "rest"
	// Audit Event: Execution of an HL7 v2 operation.
	AuditEventIDHl7v2operation AuditEventID = "hl7-v2"
	// Audit Event: Execution of an HL7 v3 operation as defined by FHIR.
	AuditEventIDHl7v3operation AuditEventID = "hl7-v3"
	// Audit Event: Execution of an operation on a Document (e.g. XDS, CDA, etc).
	AuditEventIDDocument AuditEventID = "document"
	// Audit Event: Execution of an operation on an Object. For use when a more specific event type is not available.
	AuditEventIDObject AuditEventID = "object"
)
