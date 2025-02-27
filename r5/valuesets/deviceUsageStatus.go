// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/deviceusage-status
*/type DeviceUsageStatus string

const (
	// The device is still being used.
	DeviceUsageStatusActive DeviceUsageStatus = "active"
	// The device is no longer being used.
	DeviceUsageStatusCompleted DeviceUsageStatus = "completed"
	// The device was not used.
	DeviceUsageStatusNotDone DeviceUsageStatus = "not-done"
	// The statement was recorded incorrectly.
	DeviceUsageStatusEnteredInError DeviceUsageStatus = "entered-in-error"
	// The device may be used at some time in the future.
	DeviceUsageStatusIntended DeviceUsageStatus = "intended"
	// Actions implied by the statement have been permanently halted, before all of them occurred.
	DeviceUsageStatusStopped DeviceUsageStatus = "stopped"
	// Actions implied by the statement have been temporarily halted, but are expected to continue later. May also be called "suspended".
	DeviceUsageStatusOnHold DeviceUsageStatus = "on-hold"
)
