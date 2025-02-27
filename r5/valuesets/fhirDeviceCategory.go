// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/device-category
*/type FHIRDeviceCategory string

const (
	// Device where the operation depends on a source of energy.
	FHIRDeviceCategoryActive FHIRDeviceCategory = "active"
	// Device communicates electronically to peer information systems or possibly another device.
	FHIRDeviceCategoryCommunicating FHIRDeviceCategory = "communicating"
	// Equipment and supplies that provides therapeutic benefits to a patient.
	FHIRDeviceCategoryDme FHIRDeviceCategory = "dme"
	// Medical device intended for users in a non-medical setting.
	FHIRDeviceCategoryMaintenance FHIRDeviceCategory = "home-use"
	// A device that is placed into a surgically or naturally formed cavity of the human body.
	FHIRDeviceCategoryImplantable FHIRDeviceCategory = "implantable"
	// Tests done on samples such as blood or tissue that have been taken from the human body.
	FHIRDeviceCategoryInVitro FHIRDeviceCategory = "in-vitro"
	// a class of communicating devices that are used by medical providers for various purposes (e.g., monitoring, delivering or measuring).
	FHIRDeviceCategoryPointOfCare FHIRDeviceCategory = "point-of-care"
	// A device use on one individual during a single procedure.
	FHIRDeviceCategorySingleUse FHIRDeviceCategory = "single-use"
	// A device that healthcare providers can use to diagnose and treat one or more patients.
	FHIRDeviceCategoryReusable FHIRDeviceCategory = "reusable"
	// A device that may include a software component or consist exclusively of software e.g. data transformer or converter, clinical support algorithms, clinical apps
	FHIRDeviceCategorySoftware FHIRDeviceCategory = "software"
)
