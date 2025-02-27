// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/eligibility-outcome
*/type EligibilityOutcome string

const (
	// The Claim/Pre-authorization/Pre-determination has been received but processing has not begun.
	EligibilityOutcomeQueued EligibilityOutcome = "queued"
	// The processing has completed without errors
	EligibilityOutcomeComplete EligibilityOutcome = "complete"
	// One or more errors have been detected in the Claim
	EligibilityOutcomeError EligibilityOutcome = "error"
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	EligibilityOutcomePartial EligibilityOutcome = "partial"
)
