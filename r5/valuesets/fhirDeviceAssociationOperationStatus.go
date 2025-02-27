// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/deviceassociation-operationstatus
*/type FHIRDeviceAssociationOperationStatus string

const (
	// The device is working or switched on, i.e. active.
	FHIRDeviceAssociationOperationStatusOn FHIRDeviceAssociationOperationStatus = "on"
	// The device is inactive, switched off, or deactivated.
	FHIRDeviceAssociationOperationStatusOff FHIRDeviceAssociationOperationStatus = "off"
	// The device is in stand-by mode i.e. not actively working but not powered off.
	FHIRDeviceAssociationOperationStatusStandby FHIRDeviceAssociationOperationStatus = "standby"
	// The device is defective or for maintenance and is not available or working.
	FHIRDeviceAssociationOperationStatusDefective FHIRDeviceAssociationOperationStatus = "defective"
	// The operational status of the device has not been determined.
	FHIRDeviceAssociationOperationStatusUnknown FHIRDeviceAssociationOperationStatus = "unknown"
)
