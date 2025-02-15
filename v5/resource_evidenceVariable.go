// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EvidenceVariable
type EvidenceVariable struct {
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
	// Canonical identifier for this evidence variable, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the evidence variable
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the evidence variable
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this evidence variable (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this evidence variable (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Title for use in informal contexts
	ShortTitle *string `bson:"shortTitle,omitempty" json:"shortTitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the evidence variable
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Used for footnotes or explanatory notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Why this EvidenceVariable is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the resource was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the resource was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the resource is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// Who authored the content
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Actual or conceptual
	Actual *bool `bson:"actual,omitempty" json:"actual,omitempty"`
	// A defining factor of the EvidenceVariable
	Characteristic []EvidenceVariableCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	// continuous | dichotomous | ordinal | polychotomous
	Handling *custom.Code `bson:"handling,omitempty" json:"handling,omitempty"`
	// A grouping for ordinal or polychotomous variables
	Category []EvidenceVariableCategory `bson:"category,omitempty" json:"category,omitempty"`
}

type EvidenceVariableCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for internal linking
	LinkId *custom.ID `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Natural language description of the characteristic
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Used for footnotes or explanatory notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Whether the characteristic is an inclusion criterion or exclusion criterion
	Exclude *bool `bson:"exclude,omitempty" json:"exclude,omitempty"`
	// Defines the characteristic (without using type and value) by a Reference
	DefinitionReference *custom.Reference[EvidenceVariableCharacteristicDefinitionReference] `bson:"definitionReference,omitempty" json:"definitionReference,omitempty"`
	// Defines the characteristic (without using type and value) by a Canonical
	DefinitionCanonical *custom.Canonical[EvidenceVariableCharacteristicDefinitionCanonical] `bson:"definitionCanonical,omitempty" json:"definitionCanonical,omitempty"`
	// Defines the characteristic (without using type and value) by a CodeableConcept
	DefinitionCodeableConcept *CodeableConcept `bson:"definitionCodeableConcept,omitempty" json:"definitionCodeableConcept,omitempty"`
	// Defines the characteristic (without using type and value) by an expression
	DefinitionExpression *Expression `bson:"definitionExpression,omitempty" json:"definitionExpression,omitempty"`
	// Defines the characteristic (without using type and value) by an id
	DefinitionId *custom.ID `bson:"definitionId,omitempty" json:"definitionId,omitempty"`
	// Defines the characteristic using type and value
	DefinitionByTypeAndValue *EvidenceVariableCharacteristicDefinitionByTypeAndValue `bson:"definitionByTypeAndValue,omitempty" json:"definitionByTypeAndValue,omitempty"`
	// Used to specify how two or more characteristics are combined
	DefinitionByCombination *EvidenceVariableCharacteristicDefinitionByCombination `bson:"definitionByCombination,omitempty" json:"definitionByCombination,omitempty"`
	// Number of occurrences meeting the characteristic
	Instances *Quantity `bson:"instances,omitempty" json:"instances,omitempty"`
	// Length of time in which the characteristic is met
	Duration *Quantity `bson:"duration,omitempty" json:"duration,omitempty"`
	// Timing in which the characteristic is determined
	TimeFromEvent []EvidenceVariableCharacteristicTimeFromEvent `bson:"timeFromEvent,omitempty" json:"timeFromEvent,omitempty"`
}

type EvidenceVariableCharacteristicDefinitionByTypeAndValue struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Expresses the type of characteristic
	Type CodeableConcept `bson:"type" json:"type"`
	// Method for how the characteristic value was determined
	Method []CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Device used for determining characteristic
	Device *custom.Reference[EvidenceVariableCharacteristicDefinitionByTypeAndValueDevice] `bson:"device,omitempty" json:"device,omitempty"`
	// Defines the characteristic when coupled with characteristic.type
	Value CodeableConcept `bson:"value" json:"value"`
	// Reference point for valueQuantity or valueRange
	Offset *CodeableConcept `bson:"offset,omitempty" json:"offset,omitempty"`
}

type EvidenceVariableCharacteristicDefinitionByCombination struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// all-of | any-of | at-least | at-most | statistical | net-effect | dataset
	Code custom.Code `bson:"code" json:"code"`
	// Provides the value of "n" when "at-least" or "at-most" codes are used
	Threshold *uint32 `bson:"threshold,omitempty" json:"threshold,omitempty"`
	// A defining factor of the characteristic
	Characteristic []interface{} `bson:"characteristic" json:"characteristic"`
}

type EvidenceVariableCharacteristicTimeFromEvent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Human readable description
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Used for footnotes or explanatory notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// The event used as a base point (reference point) in time
	Event *CodeableConcept `bson:"event,omitempty" json:"event,omitempty"`
	// Used to express the observation at a defined amount of time before or after the event
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Used to express the observation within a period before and/or after the event
	Range *Range `bson:"range,omitempty" json:"range,omitempty"`
}

type EvidenceVariableCategory struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Description of the grouping
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Definition of the grouping
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

type EvidenceVariableCharacteristicDefinitionReference interface {
	gofhirclient.Resource

	Is_EvidenceVariableCharacteristicDefinitionReference()
}

func (e EvidenceVariable) Is_EvidenceVariableCharacteristicDefinitionReference() {}
func (g Group) Is_EvidenceVariableCharacteristicDefinitionReference()            {}
func (e Evidence) Is_EvidenceVariableCharacteristicDefinitionReference()         {}

type EvidenceVariableCharacteristicDefinitionCanonical interface {
	gofhirclient.Resource

	Is_EvidenceVariableCharacteristicDefinitionCanonical()
}

func (e EvidenceVariable) Is_EvidenceVariableCharacteristicDefinitionCanonical() {}
func (e Evidence) Is_EvidenceVariableCharacteristicDefinitionCanonical()         {}

type EvidenceVariableCharacteristicDefinitionByTypeAndValueDevice interface {
	gofhirclient.Resource

	Is_EvidenceVariableCharacteristicDefinitionByTypeAndValueDevice()
}

func (d Device) Is_EvidenceVariableCharacteristicDefinitionByTypeAndValueDevice()       {}
func (d DeviceMetric) Is_EvidenceVariableCharacteristicDefinitionByTypeAndValueDevice() {}

func (e EvidenceVariable) ResourceType() string {
	return "StructureDefinition"
}
