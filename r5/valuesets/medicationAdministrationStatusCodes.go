// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/medication-admin-status
*/type MedicationAdministrationStatusCodes string

const (
	// The administration has started but has not yet completed.
	MedicationAdministrationStatusCodesInProgress MedicationAdministrationStatusCodes = "in-progress"
	// The administration was terminated prior to any impact on the subject (though preparatory actions may have been taken)
	MedicationAdministrationStatusCodesNotDone MedicationAdministrationStatusCodes = "not-done"
	// Actions implied by the administration have been temporarily halted, but are expected to continue later. May also be called 'suspended'.
	MedicationAdministrationStatusCodesOnHold MedicationAdministrationStatusCodes = "on-hold"
	// All actions that are implied by the administration have occurred.
	MedicationAdministrationStatusCodesCompleted MedicationAdministrationStatusCodes = "completed"
	// The administration was entered in error and therefore nullified.
	MedicationAdministrationStatusCodesEnteredInError MedicationAdministrationStatusCodes = "entered-in-error"
	// Actions implied by the administration have been permanently halted, before all of them occurred.
	MedicationAdministrationStatusCodesStopped MedicationAdministrationStatusCodes = "stopped"
	// The authoring system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, it's just not known which one.
	MedicationAdministrationStatusCodesUnknown MedicationAdministrationStatusCodes = "unknown"
)
