// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/event-status
*/type EventStatus string

const (
	// The core event has not started yet, but some staging activities have begun (e.g. surgical suite preparation).  Preparation stages may be tracked for billing purposes.
	EventStatusPreparation EventStatus = "preparation"
	// The event is currently occurring.
	EventStatusInProgress EventStatus = "in-progress"
	// The event was terminated prior to any activity beyond preparation.  I.e. The 'main' activity has not yet begun.  The boundary between preparatory and the 'main' activity is context-specific.
	EventStatusNotDone EventStatus = "not-done"
	// The event has been temporarily stopped but is expected to resume in the future.
	EventStatusOnHold EventStatus = "on-hold"
	// The event was terminated prior to the full completion of the intended activity but after at least some of the 'main' activity (beyond preparation) has occurred.
	EventStatusStopped EventStatus = "stopped"
	// The event has now concluded.
	EventStatusCompleted EventStatus = "completed"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "stopped" rather than "entered-in-error".).
	EventStatusEnteredInError EventStatus = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this event.  Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply,  but the authoring/source system does not know which.
	EventStatusUnknown EventStatus = "unknown"
)
