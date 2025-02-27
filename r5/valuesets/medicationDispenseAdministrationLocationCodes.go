// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/medicationdispense-admin-location
*/type MedicationDispenseAdministrationLocationCodes string

const (
	// Includes dispenses for medications to be administered or consumed in an inpatient or acute care setting.
	MedicationDispenseAdministrationLocationCodesInpatient MedicationDispenseAdministrationLocationCodes = "inpatient"
	// Includes dispenses for medications to be administered or consumed in an outpatient setting (for example, Emergency Department, Outpatient Clinic, Outpatient Surgery, Doctor's office).
	MedicationDispenseAdministrationLocationCodesOutpatient MedicationDispenseAdministrationLocationCodes = "outpatient"
	// Includes dispenses for medications to be administered or consumed by the patient in their home (this would include long term care or nursing homes, hospices, etc.).
	MedicationDispenseAdministrationLocationCodesCommunity MedicationDispenseAdministrationLocationCodes = "community"
)
