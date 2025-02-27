// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/device-operation-mode
*/type FHIRDeviceOperationMode string

const (
	// The device operates in a mode that supports the fulfillment of its clinical functions.
	FHIRDeviceOperationModeNormal FHIRDeviceOperationMode = "normal"
	// The device operates in a mode that is intended for demonstration purposes only. Arbitrary values are generated.
	FHIRDeviceOperationModeDemo FHIRDeviceOperationMode = "demo"
	// The device operates in a mode that is intended for correcting a functional problem of the device only. Arbitrary values may be generated.
	FHIRDeviceOperationModeService FHIRDeviceOperationMode = "service"
	// The device operates in a mode that is intended for preventative and/or scheduled maintenance purposes only.
	FHIRDeviceOperationModeMaintenance FHIRDeviceOperationMode = "maintenance"
	// The device operates in a test mode that is not intended to be used for production/operational purposes.
	FHIRDeviceOperationModeTest FHIRDeviceOperationMode = "test"
)
