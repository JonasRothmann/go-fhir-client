// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-participant-type
*/type ActionParticipantType string

const (
	// The participant is a care team caring for the patient under evaluation.
	ActionParticipantTypeCareteam ActionParticipantType = "careteam"
	// The participant is a system or device used in the care of the patient.
	ActionParticipantTypeDevice ActionParticipantType = "device"
	// The participant is a group of participants involved in the care of the patient.
	ActionParticipantTypeGroup ActionParticipantType = "group"
	// The participant is an institution that can provide the given healthcare service used in the care of the patient.
	ActionParticipantTypeHealthcareservice ActionParticipantType = "healthcareservice"
	// The participant is a location involved in the care of the patient.
	ActionParticipantTypeLocation ActionParticipantType = "location"
	// The participant is an organization involved in the care of the patient.
	ActionParticipantTypeOrganization ActionParticipantType = "organization"
	// The participant is the patient under evaluation.
	ActionParticipantTypePatient ActionParticipantType = "patient"
	// The participant is a practitioner involved in the patient's care.
	ActionParticipantTypePractitioner ActionParticipantType = "practitioner"
	// The participant is a particular practitioner role involved in the patient's care.
	ActionParticipantTypePractitionerrole ActionParticipantType = "practitionerrole"
	// The participant is a person related to the patient.
	ActionParticipantTypeRelatedperson ActionParticipantType = "relatedperson"
)
