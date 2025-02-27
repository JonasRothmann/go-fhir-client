// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Device
type DeviceProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code that specifies the property being represented
	Type CodeableConcept `json:"type"`
	// Value of the property
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Value of the property
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Value of the property
	ValueString *string `json:"valueString,omitempty"`
	// Value of the property
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Value of the property
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Value of the property
	ValueRange *Range `json:"valueRange,omitempty"`
	// Value of the property
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
}

type Device struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Instance identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// The name used to display by default when the device is referenced
	DisplayName *string `json:"displayName,omitempty"`
	// The reference to the definition for the device
	Definition *custom.CodeableReference[DeviceDefinition] `json:"definition,omitempty"`
	// Unique Device Identifier (UDI) Barcode string
	UdiCarrier []DeviceUdiCarrier `json:"udiCarrier,omitempty"`
	// active | inactive | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// lost | damaged | destroyed | available
	AvailabilityStatus *CodeableConcept `json:"availabilityStatus,omitempty"`
	// An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled
	BiologicalSourceEvent *Identifier `json:"biologicalSourceEvent,omitempty"`
	// Name of device manufacturer
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Date when the device was made
	ManufactureDate *custom.DateTime `json:"manufactureDate,omitempty"`
	// Date and time of expiry of this device (if applicable)
	ExpirationDate *custom.DateTime `json:"expirationDate,omitempty"`
	// Lot number of manufacture
	LotNumber *string `json:"lotNumber,omitempty"`
	// Serial number assigned by the manufacturer
	SerialNumber *string `json:"serialNumber,omitempty"`
	// The name or names of the device as known to the manufacturer and/or patient
	Name []DeviceName `json:"name,omitempty"`
	// The manufacturer's model number for the device
	ModelNumber *string `json:"modelNumber,omitempty"`
	// The part number or catalog number of the device
	PartNumber *string `json:"partNumber,omitempty"`
	// Indicates a high-level grouping of the device
	Category []CodeableConcept `json:"category,omitempty"`
	// The kind or type of device
	Type []CodeableConcept `json:"type,omitempty"`
	// The actual design of the device or software version running on the device
	Version []DeviceVersion `json:"version,omitempty"`
	// Identifies the standards, specifications, or formal guidances for the capabilities supported by the device
	ConformsTo []DeviceConformsTo `json:"conformsTo,omitempty"`
	// Inherent, essentially fixed, characteristics of the device.  e.g., time properties, size, material, etc.
	Property []DeviceProperty `json:"property,omitempty"`
	// The designated condition for performing a task
	Mode *CodeableConcept `json:"mode,omitempty"`
	// The series of occurrences that repeats during the operation of the device
	Cycle *Count `json:"cycle,omitempty"`
	// A measurement of time during the device's operation (e.g., days, hours, mins, etc.)
	Duration *Duration `json:"duration,omitempty"`
	// Organization responsible for device
	Owner *custom.Reference[Organization] `json:"owner,omitempty"`
	// Details for human/organization for support
	Contact []ContactPoint `json:"contact,omitempty"`
	// Where the device is found
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Network address to contact device
	Url *custom.URI `json:"url,omitempty"`
	// Technical endpoints providing access to electronic services provided by the device
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
	// Linked device acting as a communication/data collector, translator or controller
	Gateway []custom.CodeableReference[Device] `json:"gateway,omitempty"`
	// Device notes and comments
	Note []Annotation `json:"note,omitempty"`
	// Safety Characteristics of Device
	Safety []CodeableConcept `json:"safety,omitempty"`
	// The higher level or encompassing device that this device is a logical part of
	Parent *custom.Reference[Device] `json:"parent,omitempty"`
}

type DeviceUdiCarrier struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Mandatory fixed portion of UDI
	DeviceIdentifier string `json:"deviceIdentifier"`
	// UDI Issuing Organization
	Issuer custom.URI `json:"issuer"`
	// Regional UDI authority
	Jurisdiction *custom.URI `json:"jurisdiction,omitempty"`
	// UDI Machine Readable Barcode String
	CarrierAidc *custom.Base64binary `json:"carrierAidc,omitempty"`
	// UDI Human Readable Barcode String
	CarrierHrf *string `json:"carrierHrf,omitempty"`
	// barcode | rfid | manual | card | self-reported | electronic-transmission | unknown
	EntryType *custom.Code `json:"entryType,omitempty"`
}

type DeviceName struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The term that names the device
	Value string `json:"value"`
	// registered-name | user-friendly-name | patient-reported-name
	Type custom.Code `json:"type"`
	// The preferred device name
	Display *bool `json:"display,omitempty"`
}

type DeviceVersion struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of the device version, e.g. manufacturer, approved, internal
	Type *CodeableConcept `json:"type,omitempty"`
	// The hardware or software module of the device to which the version applies
	Component *Identifier `json:"component,omitempty"`
	// The date the version was installed on the device
	InstallDate *custom.DateTime `json:"installDate,omitempty"`
	// The version text
	Value string `json:"value"`
}

type DeviceConformsTo struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Describes the common type of the standard, specification, or formal guidance.  communication | performance | measurement
	Category *CodeableConcept `json:"category,omitempty"`
	// Identifies the standard, specification, or formal guidance that the device adheres to
	Specification CodeableConcept `json:"specification"`
	// Specific form or variant of the standard
	Version *string `json:"version,omitempty"`
}

type OtherDevice Device

func (d Device) ResourceType() string {
	return "Device"
}

func (d Device) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDevice
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDevice: OtherDevice(d), ResourceType: d.ResourceType()})
}

func UnmarshalDevice(b []byte) (res Device, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
