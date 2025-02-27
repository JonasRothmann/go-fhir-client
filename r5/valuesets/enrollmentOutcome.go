// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/enrollment-outcome
*/type EnrollmentOutcome string

const (
	// The Claim/Pre-authorization/Pre-determination has been received but processing has not begun.
	EnrollmentOutcomeQueued EnrollmentOutcome = "queued"
	// The processing has completed without errors
	EnrollmentOutcomeComplete EnrollmentOutcome = "complete"
	// One or more errors have been detected in the Claim
	EnrollmentOutcomeError EnrollmentOutcome = "error"
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	EnrollmentOutcomePartial EnrollmentOutcome = "partial"
)
