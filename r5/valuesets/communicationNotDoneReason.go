// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/communication-not-done-reason
*/type CommunicationNotDoneReason string

const (
	// The communication was not done due to an unknown reason.
	CommunicationNotDoneReasonUnknown CommunicationNotDoneReason = "unknown"
	// The communication was not done due to a system error.
	CommunicationNotDoneReasonSystemError CommunicationNotDoneReason = "system-error"
	// The communication was not done due to an invalid phone number.
	CommunicationNotDoneReasonInvalidPhoneNumber CommunicationNotDoneReason = "invalid-phone-number"
	// The communication was not done due to the recipient being unavailable.
	CommunicationNotDoneReasonRecipientUnavailable CommunicationNotDoneReason = "recipient-unavailable"
	// The communication was not done due to a family objection.
	CommunicationNotDoneReasonFamilyObjection CommunicationNotDoneReason = "family-objection"
	// The communication was not done due to a patient objection.
	CommunicationNotDoneReasonPatientObjection CommunicationNotDoneReason = "patient-objection"
)
