// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/request-status
*/type RequestStatus string

const (
	// The request has been created but is not yet complete or ready for action.
	RequestStatusDraft RequestStatus = "draft"
	// The request is in force and ready to be acted upon.
	RequestStatusActive RequestStatus = "active"
	// The request (and any implicit authorization to act) has been temporarily withdrawn but is expected to resume in the future.
	RequestStatusOnHold RequestStatus = "on-hold"
	// The request (and any implicit authorization to act) has been terminated prior to the known full completion of the intended actions.  No further activity should occur.
	RequestStatusRevoked RequestStatus = "revoked"
	// The activity described by the request has been fully performed.  No further activity will occur.
	RequestStatusCompleted RequestStatus = "completed"
	// This request should never have existed and should be considered 'void'.  (It is possible that real-world decisions were based on it.  If real-world activity has occurred, the status should be "revoked" rather than "entered-in-error".).
	RequestStatusEnteredInError RequestStatus = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this request.  Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which.
	RequestStatusUnknown RequestStatus = "unknown"
)
