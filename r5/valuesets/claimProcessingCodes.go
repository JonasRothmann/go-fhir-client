// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/claim-outcome
*/type ClaimProcessingCodes string

const (
	// The Claim/Pre-authorization/Pre-determination has been received but processing has not begun.
	ClaimProcessingCodesQueued ClaimProcessingCodes = "queued"
	// The processing has completed without errors
	ClaimProcessingCodesComplete ClaimProcessingCodes = "complete"
	// One or more errors have been detected in the Claim
	ClaimProcessingCodesError ClaimProcessingCodes = "error"
	// No errors have been detected in the Claim and some of the adjudication has been performed.
	ClaimProcessingCodesPartial ClaimProcessingCodes = "partial"
)
