// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/medication-statement-status
*/type MedicationStatementStatusCodes string

const (
	// The action of recording the medication statement is finished.
	MedicationStatementStatusCodesRecorded MedicationStatementStatusCodes = "recorded"
	// Some of the actions that are implied by the medication usage may have occurred.  For example, the patient may have taken some of the medication.  Clinical decision support systems should take this status into account.
	MedicationStatementStatusCodesEnteredInError MedicationStatementStatusCodes = "entered-in-error"
	// The medication usage is draft or preliminary.
	MedicationStatementStatusCodesDraft MedicationStatementStatusCodes = "draft"
)
