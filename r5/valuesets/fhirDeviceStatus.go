// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/device-status
*/type FHIRDeviceStatus string

const (
	// The device record is current and is appropriate for reference in new instances.
	FHIRDeviceStatusActive FHIRDeviceStatus = "active"
	// The device record is not current and is not appropriate for reference in new instances.
	FHIRDeviceStatusInactive FHIRDeviceStatus = "inactive"
	// The device record is not current and is not appropriate for reference in new instances.
	FHIRDeviceStatusEnteredInError FHIRDeviceStatus = "entered-in-error"
)
