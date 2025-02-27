// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/transport-status
*/type TransportStatus string

const (
	// Transport has started but not completed.
	TransportStatusInProgress TransportStatus = "in-progress"
	// Transport has been completed.
	TransportStatusCompleted TransportStatus = "completed"
	// Transport was started but not completed.
	TransportStatusAbandoned TransportStatus = "abandoned"
	// Transport was cancelled before started.
	TransportStatusCancelled TransportStatus = "cancelled"
	// Planned transport that is not yet requested.
	TransportStatusPlanned TransportStatus = "planned"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "abandoned" rather than "entered-in-error".).
	TransportStatusEnteredInError TransportStatus = "entered-in-error"
)
