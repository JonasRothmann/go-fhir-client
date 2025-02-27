// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/devicedispense-status
*/type DeviceDispenseStatusCodes string

const (
	// The core event has not started yet, but some staging activities have begun (e.g. initial preparing of the device. Preparation stages may be tracked e.g. for planning, supply or billing purposes.
	DeviceDispenseStatusCodesPreparation DeviceDispenseStatusCodes = "preparation"
	// The dispensed product is ready for pickup.
	DeviceDispenseStatusCodesInProgress DeviceDispenseStatusCodes = "in-progress"
	// The dispensed product was not and will never be picked up by the patient.
	DeviceDispenseStatusCodesCancelled DeviceDispenseStatusCodes = "cancelled"
	// The dispense process is paused while waiting for an external event to reactivate the dispense.  For example, new stock has arrived or the prescriber has called.
	DeviceDispenseStatusCodesOnHold DeviceDispenseStatusCodes = "on-hold"
	// The dispensed product has been picked up.
	DeviceDispenseStatusCodesCompleted DeviceDispenseStatusCodes = "completed"
	// The dispense was entered in error and therefore nullified.
	DeviceDispenseStatusCodesEnteredInError DeviceDispenseStatusCodes = "entered-in-error"
	// Actions implied by the dispense have been permanently halted, before all of them occurred.
	DeviceDispenseStatusCodesStopped DeviceDispenseStatusCodes = "stopped"
	// The dispense was declined and not performed.
	DeviceDispenseStatusCodesDeclined DeviceDispenseStatusCodes = "declined"
	// The authoring system does not know which of the status values applies for this dispense.  Note: this concept is not to be used for other - one of the listed statuses is presumed to apply, it's just now known which one.
	DeviceDispenseStatusCodesUnknown DeviceDispenseStatusCodes = "unknown"
)
