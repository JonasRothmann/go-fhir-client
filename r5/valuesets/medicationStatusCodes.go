// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/medication-status
*/type MedicationStatusCodes string

const (
	// The medication record is current and is appropriate for reference in new instances.
	MedicationStatusCodesActive MedicationStatusCodes = "active"
	// The medication record is not current and is not is appropriate for reference in new instances.
	MedicationStatusCodesInactive MedicationStatusCodes = "inactive"
	// The medication record was created erroneously and is not appropriated for reference in new instances.
	MedicationStatusCodesEnteredInError MedicationStatusCodes = "entered-in-error"
)
