// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/event-status
  - in-progress
  - completed
  - entered-in-error
  - unknown
*/type AdverseEventStatus string

const (
	// The event is currently occurring.
	AdverseEventStatusInProgress AdverseEventStatus = "in-progress"
	// The event has now concluded.
	AdverseEventStatusCompleted AdverseEventStatus = "completed"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "stopped" rather than "entered-in-error".).
	AdverseEventStatusEnteredInError AdverseEventStatus = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this event.  Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which.
	AdverseEventStatusUnknown AdverseEventStatus = "unknown"
)
