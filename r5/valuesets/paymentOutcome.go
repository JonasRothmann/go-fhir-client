// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/payment-outcome
*/type PaymentOutcome string

const (
	// The Claim/Pre-authorization/Pre-determination has been received but processing has not begun.
	PaymentOutcomeQueued PaymentOutcome = "queued"
	// The processing has completed without errors
	PaymentOutcomeComplete PaymentOutcome = "complete"
	// One or more errors have been detected in the Claim
	PaymentOutcomeError PaymentOutcome = "error"
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	PaymentOutcomePartial PaymentOutcome = "partial"
)
