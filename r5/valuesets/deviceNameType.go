// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/device-nametype
*/type DeviceNameType string

const (
	// The term assigned to a medical device by the entity who registers or submits information about it to a jurisdiction or its databases. This may be considered the manufacturer assigned name (e.g., brand name assigned by the labeler or manufacturer in US, or device name assigned by the manufacturer in EU) and may also be synonymous with proprietary name or trade name of the device.
	DeviceNameTypeRegisteredName DeviceNameType = "registered-name"
	// The term that generically describes the device by a name as assigned by the manufacturer that is recognized by lay person.  This common or generic name may be printed on the package it came in or some combination of that name with the model number, serial number, or other attribute that makes the name easy to understand for the user of that device. It is often exposed in communicating devices transport protocols. It is provided to help users identify the device when reported in discovery operations.
	DeviceNameTypeUserFriendlyName DeviceNameType = "user-friendly-name"
	// the term used by the patient associated with the device when describing the device, for example 'knee implant', when documented as a self-reported device.
	DeviceNameTypePatientReportedName DeviceNameType = "patient-reported-name"
)
