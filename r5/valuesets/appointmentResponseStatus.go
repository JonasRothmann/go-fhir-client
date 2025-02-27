// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/participationstatus
- http://hl7.org/fhir/appointmentstatus
  - entered-in-error
*/type AppointmentResponseStatus string

const (
	// The participant has accepted the appointment.
	AppointmentResponseStatusAccepted AppointmentResponseStatus = "accepted"
	// The participant has declined the appointment and will not participate in the appointment.
	AppointmentResponseStatusDeclined AppointmentResponseStatus = "declined"
	// The participant has  tentatively accepted the appointment. This could be automatically created by a system and requires further processing before it can be accepted. There is no commitment that attendance will occur.
	AppointmentResponseStatusTentative AppointmentResponseStatus = "tentative"
	// The participant needs to indicate if they accept the appointment by changing this status to one of the other statuses.
	AppointmentResponseStatusNeedsAction AppointmentResponseStatus = "needs-action"
	// This instance should not have been part of this patient's medical record.
	AppointmentResponseStatusEnteredInError AppointmentResponseStatus = "entered-in-error"
)
