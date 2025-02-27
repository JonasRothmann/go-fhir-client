// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/verificationresult-status
*/type VerificationResultStatus string

const (
	VerificationResultStatusAttested             VerificationResultStatus = "attested"
	VerificationResultStatusValidated            VerificationResultStatus = "validated"
	VerificationResultStatusInProcess            VerificationResultStatus = "in-process"
	VerificationResultStatusRequiresRevalidation VerificationResultStatus = "req-revalid"
	VerificationResultStatusValidationFailed     VerificationResultStatus = "val-fail"
	VerificationResultStatusReValidationFailed   VerificationResultStatus = "reval-fail"
	// The VerificationResult record was created erroneously and is not appropriated for use.
	VerificationResultStatusEnteredInError VerificationResultStatus = "entered-in-error"
)
