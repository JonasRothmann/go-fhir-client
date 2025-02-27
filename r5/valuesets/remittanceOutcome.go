// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/remittance-outcome
*/type RemittanceOutcome string

const (
	// The processing completed without errors.
	RemittanceOutcomeComplete RemittanceOutcome = "complete"
	// The processing identified errors.
	RemittanceOutcomeError RemittanceOutcome = "error"
	// No errors have been detected and some of the adjudication has been performed.
	RemittanceOutcomePartial RemittanceOutcome = "partial"
)
