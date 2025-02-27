// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/audit-event-action
*/type AuditEventAction string

const (
	// Create a new database object, such as placing an order.
	AuditEventActionCreate AuditEventAction = "C"
	// Read data, such as to print or display to a doctor.
	AuditEventActionRead AuditEventAction = "R"
	// Update data, such as revise patient information.
	AuditEventActionUpdate AuditEventAction = "U"
	// Delete items, such as a doctor master file record.
	AuditEventActionDelete AuditEventAction = "D"
	// Perform a system or application function such as log-on, program execution or use of an object's method, or perform a query/search operation.
	AuditEventActionExecute AuditEventAction = "E"
)
