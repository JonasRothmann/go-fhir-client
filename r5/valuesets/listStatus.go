// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/list-status
*/type ListStatus string

const (
	// The list is considered to be an active part of the patient's record.
	ListStatusCurrent ListStatus = "current"
	// The list is "old" and should no longer be considered accurate or relevant.
	ListStatusRetired ListStatus = "retired"
	// The list was never accurate.  It is retained for medico-legal purposes only.
	ListStatusEnteredInError ListStatus = "entered-in-error"
)
