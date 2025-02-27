// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/consent-state-codes
*/type ConsentState string

const (
	// The consent is in development or awaiting use but is not yet intended to be acted upon.
	ConsentStateDraft ConsentState = "draft"
	// The consent is to be followed and enforced.
	ConsentStateActive ConsentState = "active"
	// The consent is terminated or replaced.
	ConsentStateInactive ConsentState = "inactive"
	// The consent development has been terminated prior to completion.
	ConsentStateAbandoned ConsentState = "not-done"
	// The consent was created wrongly (e.g. wrong patient) and should be ignored.
	ConsentStateEnteredInError ConsentState = "entered-in-error"
	// The resource is in an indeterminate state.
	ConsentStateUnknown ConsentState = "unknown"
)
