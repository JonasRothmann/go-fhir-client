// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/medicationdispense-status
*/type MedicationDispenseStatusCodes string

const (
	// The core event has not started yet, but some staging activities have begun (e.g. initial compounding or packaging of medication). Preparation stages may be tracked for billing purposes.
	MedicationDispenseStatusCodesPreparation MedicationDispenseStatusCodes = "preparation"
	// The dispensed product is ready for pickup.
	MedicationDispenseStatusCodesInProgress MedicationDispenseStatusCodes = "in-progress"
	// The dispensed product was not and will never be picked up by the patient.
	MedicationDispenseStatusCodesCancelled MedicationDispenseStatusCodes = "cancelled"
	// The dispense process is paused while waiting for an external event to reactivate the dispense.  For example, new stock has arrived or the prescriber has called.
	MedicationDispenseStatusCodesOnHold MedicationDispenseStatusCodes = "on-hold"
	// The dispensed product has been picked up.
	MedicationDispenseStatusCodesCompleted MedicationDispenseStatusCodes = "completed"
	// The dispense was entered in error and therefore nullified.
	MedicationDispenseStatusCodesEnteredInError MedicationDispenseStatusCodes = "entered-in-error"
	// Actions implied by the dispense have been permanently halted, before all of them occurred.
	MedicationDispenseStatusCodesStopped MedicationDispenseStatusCodes = "stopped"
	// The dispense was declined and not performed.
	MedicationDispenseStatusCodesDeclined MedicationDispenseStatusCodes = "declined"
	// The authoring system does not know which of the status values applies for this medication dispense.  Note: this concept is not to be used for other - one of the listed statuses is presumed to apply, it's just now known which one.
	MedicationDispenseStatusCodesUnknown MedicationDispenseStatusCodes = "unknown"
)
