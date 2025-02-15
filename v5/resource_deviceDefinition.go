// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/DeviceDefinition
type DeviceDefinitionCorrectiveAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Whether the corrective action was a recall
	Recall bool `bson:"recall" json:"recall"`
	// model | lot-numbers | serial-numbers
	Scope *custom.Code `bson:"scope,omitempty" json:"scope,omitempty"`
	// Start and end dates of the  corrective action
	Period Period `bson:"period" json:"period"`
}

type DeviceDefinitionUdiDeviceIdentifier struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The identifier that is to be associated with every Device that references this DeviceDefintiion for the issuer and jurisdiction provided in the DeviceDefinition.udiDeviceIdentifier
	DeviceIdentifier string `bson:"deviceIdentifier" json:"deviceIdentifier"`
	// The organization that assigns the identifier algorithm
	Issuer custom.URI `bson:"issuer" json:"issuer"`
	// The jurisdiction to which the deviceIdentifier applies
	Jurisdiction custom.URI `bson:"jurisdiction" json:"jurisdiction"`
	// Indicates whether and when the device is available on the market
	MarketDistribution []DeviceDefinitionUdiDeviceIdentifierMarketDistribution `bson:"marketDistribution,omitempty" json:"marketDistribution,omitempty"`
}

type DeviceDefinitionRegulatoryIdentifier struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// basic | master | license
	Type custom.Code `bson:"type" json:"type"`
	// The identifier itself
	DeviceIdentifier string `bson:"deviceIdentifier" json:"deviceIdentifier"`
	// The organization that issued this identifier
	Issuer custom.URI `bson:"issuer" json:"issuer"`
	// The jurisdiction to which the deviceIdentifier applies
	Jurisdiction custom.URI `bson:"jurisdiction" json:"jurisdiction"`
}

type DeviceDefinitionDeviceName struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A name that is used to refer to the device
	Name string `bson:"name" json:"name"`
	// registered-name | user-friendly-name | patient-reported-name
	Type custom.Code `bson:"type" json:"type"`
}

type DeviceDefinitionPackagingDistributor struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Distributor's human-readable name
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Distributor as an Organization resource
	OrganizationReference []custom.Reference[Organization] `bson:"organizationReference,omitempty" json:"organizationReference,omitempty"`
}

type DeviceDefinitionProperty struct {
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

type DeviceDefinitionLink struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type indicates the relationship of the related device to the device instance
	Relation Coding `bson:"relation" json:"relation"`
	// A reference to the linked device
	RelatedDevice custom.CodeableReference[DeviceDefinition] `bson:"relatedDevice" json:"relatedDevice"`
}

type DeviceDefinitionGuideline struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The circumstances that form the setting for using the device
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Detailed written and visual directions for the user on how to use the device
	UsageInstruction *custom.Markdown `bson:"usageInstruction,omitempty" json:"usageInstruction,omitempty"`
	// A source of information or reference for this guideline
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// A clinical condition for which the device was designed to be used
	Indication []CodeableConcept `bson:"indication,omitempty" json:"indication,omitempty"`
	// A specific situation when a device should not be used because it may cause harm
	Contraindication []CodeableConcept `bson:"contraindication,omitempty" json:"contraindication,omitempty"`
	// Specific hazard alert information that a user needs to know before using the device
	Warning []CodeableConcept `bson:"warning,omitempty" json:"warning,omitempty"`
	// A description of the general purpose or medical use of the device or its function
	IntendedUse *string `bson:"intendedUse,omitempty" json:"intendedUse,omitempty"`
}

type DeviceDefinition struct {
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
	// Additional information to describe the device
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Instance identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Unique Device Identifier (UDI) Barcode string
	UdiDeviceIdentifier []DeviceDefinitionUdiDeviceIdentifier `bson:"udiDeviceIdentifier,omitempty" json:"udiDeviceIdentifier,omitempty"`
	// Regulatory identifier(s) associated with this device
	RegulatoryIdentifier []DeviceDefinitionRegulatoryIdentifier `bson:"regulatoryIdentifier,omitempty" json:"regulatoryIdentifier,omitempty"`
	// The part number or catalog number of the device
	PartNumber *string `bson:"partNumber,omitempty" json:"partNumber,omitempty"`
	// Name of device manufacturer
	Manufacturer *custom.Reference[Organization] `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// The name or names of the device as given by the manufacturer
	DeviceName []DeviceDefinitionDeviceName `bson:"deviceName,omitempty" json:"deviceName,omitempty"`
	// The catalog or model number for the device for example as defined by the manufacturer
	ModelNumber *string `bson:"modelNumber,omitempty" json:"modelNumber,omitempty"`
	// What kind of device or device system this is
	Classification []DeviceDefinitionClassification `bson:"classification,omitempty" json:"classification,omitempty"`
	// Identifies the standards, specifications, or formal guidances for the capabilities supported by the device
	ConformsTo []DeviceDefinitionConformsTo `bson:"conformsTo,omitempty" json:"conformsTo,omitempty"`
	// A device, part of the current one
	HasPart []DeviceDefinitionHasPart `bson:"hasPart,omitempty" json:"hasPart,omitempty"`
	// Information about the packaging of the device, i.e. how the device is packaged
	Packaging []DeviceDefinitionPackaging `bson:"packaging,omitempty" json:"packaging,omitempty"`
	// The version of the device or software
	Version []DeviceDefinitionVersion `bson:"version,omitempty" json:"version,omitempty"`
	// Safety characteristics of the device
	Safety []CodeableConcept `bson:"safety,omitempty" json:"safety,omitempty"`
	// Shelf Life and storage information
	ShelfLifeStorage []ProductShelfLife `bson:"shelfLifeStorage,omitempty" json:"shelfLifeStorage,omitempty"`
	// Language code for the human-readable text strings produced by the device (all supported)
	LanguageCode []CodeableConcept `bson:"languageCode,omitempty" json:"languageCode,omitempty"`
	// Inherent, essentially fixed, characteristics of this kind of device, e.g., time properties, size, etc
	Property []DeviceDefinitionProperty `bson:"property,omitempty" json:"property,omitempty"`
	// Organization responsible for device
	Owner *custom.Reference[Organization] `bson:"owner,omitempty" json:"owner,omitempty"`
	// Details for human/organization for support
	Contact []ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	// An associated device, attached to, used with, communicating with or linking a previous or new device model to the focal device
	Link []DeviceDefinitionLink `bson:"link,omitempty" json:"link,omitempty"`
	// Device notes and comments
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// A substance used to create the material(s) of which the device is made
	Material []DeviceDefinitionMaterial `bson:"material,omitempty" json:"material,omitempty"`
	// lot-number | manufactured-date | serial-number | expiration-date | biological-source | software-version
	ProductionIdentifierInUdi []custom.Code `bson:"productionIdentifierInUdi,omitempty" json:"productionIdentifierInUdi,omitempty"`
	// Information aimed at providing directions for the usage of this model of device
	Guideline *DeviceDefinitionGuideline `bson:"guideline,omitempty" json:"guideline,omitempty"`
	// Tracking of latest field safety corrective action
	CorrectiveAction *DeviceDefinitionCorrectiveAction `bson:"correctiveAction,omitempty" json:"correctiveAction,omitempty"`
	// Billing code or reference associated with the device
	ChargeItem []DeviceDefinitionChargeItem `bson:"chargeItem,omitempty" json:"chargeItem,omitempty"`
}

type DeviceDefinitionClassification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A classification or risk class of the device model
	Type CodeableConcept `bson:"type" json:"type"`
	// Further information qualifying this classification of the device model
	Justification []RelatedArtifact `bson:"justification,omitempty" json:"justification,omitempty"`
}

type DeviceDefinitionConformsTo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Describes the common type of the standard, specification, or formal guidance
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Identifies the standard, specification, or formal guidance that the device adheres to the Device Specification type
	Specification CodeableConcept `bson:"specification" json:"specification"`
	// The specific form or variant of the standard, specification or formal guidance
	Version []string `bson:"version,omitempty" json:"version,omitempty"`
	// Standard, regulation, certification, or guidance website, document, or other publication, or similar, supporting the conformance
	Source []RelatedArtifact `bson:"source,omitempty" json:"source,omitempty"`
}

type DeviceDefinitionHasPart struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to the part
	Reference custom.Reference[DeviceDefinition] `bson:"reference" json:"reference"`
	// Number of occurrences of the part
	Count *int32 `bson:"count,omitempty" json:"count,omitempty"`
}

type DeviceDefinitionMaterial struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A relevant substance that the device contains, may contain, or is made of
	Substance CodeableConcept `bson:"substance" json:"substance"`
	// Indicates an alternative material of the device
	Alternate *bool `bson:"alternate,omitempty" json:"alternate,omitempty"`
	// Whether the substance is a known or suspected allergen
	AllergenicIndicator *bool `bson:"allergenicIndicator,omitempty" json:"allergenicIndicator,omitempty"`
}

type DeviceDefinitionPackaging struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Business identifier of the packaged medication
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// A code that defines the specific type of packaging
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The number of items contained in the package (devices or sub-packages)
	Count *int32 `bson:"count,omitempty" json:"count,omitempty"`
	// An organization that distributes the packaged device
	Distributor []DeviceDefinitionPackagingDistributor `bson:"distributor,omitempty" json:"distributor,omitempty"`
	// Unique Device Identifier (UDI) Barcode string on the packaging
	UdiDeviceIdentifier []interface{} `bson:"udiDeviceIdentifier,omitempty" json:"udiDeviceIdentifier,omitempty"`
	// Allows packages within packages
	Packaging []interface{} `bson:"packaging,omitempty" json:"packaging,omitempty"`
}

type DeviceDefinitionVersion struct {
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
	// The version text
	Value string `bson:"value" json:"value"`
}

type DeviceDefinitionChargeItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The code or reference for the charge item
	ChargeItemCode custom.CodeableReference[ChargeItemDefinition] `bson:"chargeItemCode" json:"chargeItemCode"`
	// Coefficient applicable to the billing code
	Count Quantity `bson:"count" json:"count"`
	// A specific time period in which this charge item applies
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// The context to which this charge item applies
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
}

type DeviceDefinitionUdiDeviceIdentifierMarketDistribution struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Begin and end dates for the commercial distribution of the device
	MarketPeriod Period `bson:"marketPeriod" json:"marketPeriod"`
	// National state or territory where the device is commercialized
	SubJurisdiction custom.URI `bson:"subJurisdiction" json:"subJurisdiction"`
}

func (d DeviceDefinition) ResourceType() string {
	return "StructureDefinition"
}
