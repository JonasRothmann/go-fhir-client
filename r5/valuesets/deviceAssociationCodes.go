// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/deviceassociation-status-reason
*/type DeviceAssociationCodes string

const (
	// The device is connected to the patient.
	DeviceAssociationCodesAttached DeviceAssociationCodes = "attached"
	// The device is no longer connected.
	DeviceAssociationCodesDisconnected DeviceAssociationCodes = "disconnected"
	// The device has failed to work, or continue to function.
	DeviceAssociationCodesFailed DeviceAssociationCodes = "failed"
	// The device was placed in the patient.
	DeviceAssociationCodesPlaced DeviceAssociationCodes = "placed"
	// The device was updated with a new device or device part.
	DeviceAssociationCodesReplaced DeviceAssociationCodes = "replaced"
)
