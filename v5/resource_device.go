// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Device
type Device struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Instance identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The name used to display by default when the device is referenced
	DisplayName *string `bson:"displayName,omitempty" json:"displayName,omitempty"`
	// The reference to the definition for the device
	Definition *custom.CodeableReference[DeviceDefinition] `bson:"definition,omitempty" json:"definition,omitempty"`
	// Unique Device Identifier (UDI) Barcode string
	UdiCarrier []DeviceUdiCarrier `bson:"udiCarrier,omitempty" json:"udiCarrier,omitempty"`
	// active | inactive | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// lost | damaged | destroyed | available
	AvailabilityStatus *CodeableConcept `bson:"availabilityStatus,omitempty" json:"availabilityStatus,omitempty"`
	// An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled
	BiologicalSourceEvent *Identifier `bson:"biologicalSourceEvent,omitempty" json:"biologicalSourceEvent,omitempty"`
	// Name of device manufacturer
	Manufacturer *string `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// Date when the device was made
	ManufactureDate *custom.DateTime `bson:"manufactureDate,omitempty" json:"manufactureDate,omitempty"`
	// Date and time of expiry of this device (if applicable)
	ExpirationDate *custom.DateTime `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	// Lot number of manufacture
	LotNumber *string `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	// Serial number assigned by the manufacturer
	SerialNumber *string `bson:"serialNumber,omitempty" json:"serialNumber,omitempty"`
	// The name or names of the device as known to the manufacturer and/or patient
	Name []DeviceName `bson:"name,omitempty" json:"name,omitempty"`
	// The manufacturer's model number for the device
	ModelNumber *string `bson:"modelNumber,omitempty" json:"modelNumber,omitempty"`
	// The part number or catalog number of the device
	PartNumber *string `bson:"partNumber,omitempty" json:"partNumber,omitempty"`
	// Indicates a high-level grouping of the device
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// The kind or type of device
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The actual design of the device or software version running on the device
	Version []DeviceVersion `bson:"version,omitempty" json:"version,omitempty"`
	// Identifies the standards, specifications, or formal guidances for the capabilities supported by the device
	ConformsTo []DeviceConformsTo `bson:"conformsTo,omitempty" json:"conformsTo,omitempty"`
	// Inherent, essentially fixed, characteristics of the device.  e.g., time properties, size, material, etc.
	Property []DeviceProperty `bson:"property,omitempty" json:"property,omitempty"`
	// The designated condition for performing a task
	Mode *CodeableConcept `bson:"mode,omitempty" json:"mode,omitempty"`
	// The series of occurrences that repeats during the operation of the device
	Cycle *Count `bson:"cycle,omitempty" json:"cycle,omitempty"`
	// A measurement of time during the device's operation (e.g., days, hours, mins, etc.)
	Duration *Duration `bson:"duration,omitempty" json:"duration,omitempty"`
	// Organization responsible for device
	Owner *custom.Reference[Organization] `bson:"owner,omitempty" json:"owner,omitempty"`
	// Details for human/organization for support
	Contact []ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	// Where the device is found
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Network address to contact device
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Technical endpoints providing access to electronic services provided by the device
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// Linked device acting as a communication/data collector, translator or controller
	Gateway []custom.CodeableReference[Device] `bson:"gateway,omitempty" json:"gateway,omitempty"`
	// Device notes and comments
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Safety Characteristics of Device
	Safety []CodeableConcept `bson:"safety,omitempty" json:"safety,omitempty"`
	// The higher level or encompassing device that this device is a logical part of
	Parent *custom.Reference[Device] `bson:"parent,omitempty" json:"parent,omitempty"`
}

type DeviceUdiCarrier struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Mandatory fixed portion of UDI
	DeviceIdentifier string `bson:"deviceIdentifier" json:"deviceIdentifier"`
	// UDI Issuing Organization
	Issuer custom.URI `bson:"issuer" json:"issuer"`
	// Regional UDI authority
	Jurisdiction *custom.URI `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// UDI Machine Readable Barcode String
	CarrierAidc *custom.Base64binary `bson:"carrierAidc,omitempty" json:"carrierAidc,omitempty"`
	// UDI Human Readable Barcode String
	CarrierHrf *string `bson:"carrierHrf,omitempty" json:"carrierHrf,omitempty"`
	// barcode | rfid | manual | card | self-reported | electronic-transmission | unknown
	EntryType *custom.Code `bson:"entryType,omitempty" json:"entryType,omitempty"`
}

type DeviceName struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The term that names the device
	Value string `bson:"value" json:"value"`
	// registered-name | user-friendly-name | patient-reported-name
	Type custom.Code `bson:"type" json:"type"`
	// The preferred device name
	Display *bool `bson:"display,omitempty" json:"display,omitempty"`
}

type DeviceVersion struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of the device version, e.g. manufacturer, approved, internal
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The hardware or software module of the device to which the version applies
	Component *Identifier `bson:"component,omitempty" json:"component,omitempty"`
	// The date the version was installed on the device
	InstallDate *custom.DateTime `bson:"installDate,omitempty" json:"installDate,omitempty"`
	// The version text
	Value string `bson:"value" json:"value"`
}

type DeviceConformsTo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Describes the common type of the standard, specification, or formal guidance.  communication | performance | measurement
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Identifies the standard, specification, or formal guidance that the device adheres to
	Specification CodeableConcept `bson:"specification" json:"specification"`
	// Specific form or variant of the standard
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
}

type DeviceProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code that specifies the property being represented
	Type CodeableConcept `bson:"type" json:"type"`
	// Value of the property
	Value Quantity `bson:"value" json:"value"`
}

func (d Device) ResourceType() string {
	return "StructureDefinition"
}
