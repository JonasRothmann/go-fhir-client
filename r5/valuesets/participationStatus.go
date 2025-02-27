// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/participationstatus
*/type ParticipationStatus string

const (
	// The participant has accepted the appointment.
	ParticipationStatusAccepted ParticipationStatus = "accepted"
	// The participant has declined the appointment and will not participate in the appointment.
	ParticipationStatusDeclined ParticipationStatus = "declined"
	// The participant has  tentatively accepted the appointment. This could be automatically created by a system and requires further processing before it can be accepted. There is no commitment that attendance will occur.
	ParticipationStatusTentative ParticipationStatus = "tentative"
	// The participant needs to indicate if they accept the appointment by changing this status to one of the other statuses.
	ParticipationStatusNeedsAction ParticipationStatus = "needs-action"
)
