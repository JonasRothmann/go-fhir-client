// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/supplyrequest-status
*/type SupplyRequestStatus string

const (
	// The request has been created but is not yet complete or ready for action.
	SupplyRequestStatusDraft SupplyRequestStatus = "draft"
	// The request is ready to be acted upon.
	SupplyRequestStatusActive SupplyRequestStatus = "active"
	// The authorization/request to act has been temporarily withdrawn but is expected to resume in the future.
	SupplyRequestStatusSuspended SupplyRequestStatus = "suspended"
	// The authorization/request to act has been terminated prior to the full completion of the intended actions.  No further activity should occur.
	SupplyRequestStatusCancelled SupplyRequestStatus = "cancelled"
	// Activity against the request has been sufficiently completed to the satisfaction of the requester.
	SupplyRequestStatusCompleted SupplyRequestStatus = "completed"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	SupplyRequestStatusEnteredInError SupplyRequestStatus = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	SupplyRequestStatusUnknown SupplyRequestStatus = "unknown"
)
