// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/transport-status-reason
*/type TransportStatusReason string

const (
	// Patient declined transport.
	TransportStatusReasonDeclinedByPatient TransportStatusReason = "declined-by-patient"
	// Transport was refused by recipient.
	TransportStatusReasonRefusedByRecipient TransportStatusReason = "refused-by-recipient"
	// The patient was not available for transport.
	TransportStatusReasonPatientNotAvailable TransportStatusReason = "patient-not-available"
	// The specimen was not available for transport.
	TransportStatusReasonSpecimenNotAvailable TransportStatusReason = "specimen-not-available"
	// Wrong data in the transport request.
	TransportStatusReasonWrongRequestData TransportStatusReason = "wrong-request-data"
	// An accident has occurred in route.
	TransportStatusReasonInRouteAccident TransportStatusReason = "in-route-accident"
	// The transport request was not acknowledged.
	TransportStatusReasonRequestNotAcknowledged TransportStatusReason = "request-not-acknowledged"
)
