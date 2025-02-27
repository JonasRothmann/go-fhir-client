// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/medication-admin-status
  - completed
  - entered-in-error
*/type ImmunizationEvaluationStatusCodes string

const (
	// All actions that are implied by the administration have occurred.
	ImmunizationEvaluationStatusCodesCompleted ImmunizationEvaluationStatusCodes = "completed"
	// The administration was entered in error and therefore nullified.
	ImmunizationEvaluationStatusCodesEnteredInError ImmunizationEvaluationStatusCodes = "entered-in-error"
)
