// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/explanationofbenefit-status
*/type ExplanationOfBenefitStatus string

const (
	// The resource instance is currently in-force.
	ExplanationOfBenefitStatusActive ExplanationOfBenefitStatus = "active"
	// The resource instance is withdrawn, rescinded or reversed.
	ExplanationOfBenefitStatusCancelled ExplanationOfBenefitStatus = "cancelled"
	// A new resource instance the contents of which is not complete.
	ExplanationOfBenefitStatusDraft ExplanationOfBenefitStatus = "draft"
	// The resource instance was entered in error.
	ExplanationOfBenefitStatusEnteredInError ExplanationOfBenefitStatus = "entered-in-error"
)
