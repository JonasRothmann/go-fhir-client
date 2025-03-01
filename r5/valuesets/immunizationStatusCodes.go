// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/event-status
  - completed
  - entered-in-error
  - not-done
*/type ImmunizationStatusCodes string

const (
	// The event was terminated prior to any activity beyond preparation.  I.e. The 'main' activity has not yet begun.  The boundary between preparatory and the 'main' activity is context-specific.
	ImmunizationStatusCodesNotDone ImmunizationStatusCodes = "not-done"
	// The event has now concluded.
	ImmunizationStatusCodesCompleted ImmunizationStatusCodes = "completed"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "stopped" rather than "entered-in-error".).
	ImmunizationStatusCodesEnteredInError ImmunizationStatusCodes = "entered-in-error"
)
