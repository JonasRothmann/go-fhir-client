// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EvidenceVariable
type EvidenceVariableCharacteristicDefinitionByCombination struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// all-of | any-of | at-least | at-most | statistical | net-effect | dataset
	Code custom.Code `json:"code"`
	// Provides the value of "n" when "at-least" or "at-most" codes are used
	Threshold *uint32 `json:"threshold,omitempty"`
	// A defining factor of the characteristic
	Characteristic []interface{} `json:"characteristic"`
}

type EvidenceVariableCharacteristicTimeFromEvent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Human readable description
	Description *custom.Markdown `json:"description,omitempty"`
	// Used for footnotes or explanatory notes
	Note []Annotation `json:"note,omitempty"`
	// The event used as a base point (reference point) in time
	EventCodeableConcept *CodeableConcept `json:"eventCodeableConcept,omitempty"`
	// The event used as a base point (reference point) in time
	EventReference *custom.Reference[gofhirclient.Resource] `json:"eventReference,omitempty"`
	// The event used as a base point (reference point) in time
	EventDateTime *custom.DateTime `json:"eventDateTime,omitempty"`
	// The event used as a base point (reference point) in time
	EventId *custom.ID `json:"eventId,omitempty"`
	// Used to express the observation at a defined amount of time before or after the event
	Quantity *Quantity `json:"quantity,omitempty"`
	// Used to express the observation within a period before and/or after the event
	Range *Range `json:"range,omitempty"`
}

type EvidenceVariableCategory struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Description of the grouping
	Name *string `json:"name,omitempty"`
	// Definition of the grouping
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Definition of the grouping
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Definition of the grouping
	ValueRange *Range `json:"valueRange,omitempty"`
}

type EvidenceVariable struct {
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
	// Canonical identifier for this evidence variable, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the evidence variable
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the evidence variable
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this evidence variable (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this evidence variable (human friendly)
	Title *string `json:"title,omitempty"`
	// Title for use in informal contexts
	ShortTitle *string `json:"shortTitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the evidence variable
	Description *custom.Markdown `json:"description,omitempty"`
	// Used for footnotes or explanatory notes
	Note []Annotation `json:"note,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Why this EvidenceVariable is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the resource was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the resource was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the resource is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// Who authored the content
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Actual or conceptual
	Actual *bool `json:"actual,omitempty"`
	// A defining factor of the EvidenceVariable
	Characteristic []EvidenceVariableCharacteristic `json:"characteristic,omitempty"`
	// continuous | dichotomous | ordinal | polychotomous
	Handling *custom.Code `json:"handling,omitempty"`
	// A grouping for ordinal or polychotomous variables
	Category []EvidenceVariableCategory `json:"category,omitempty"`
}

type EvidenceVariableCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for internal linking
	LinkId *custom.ID `json:"linkId,omitempty"`
	// Natural language description of the characteristic
	Description *custom.Markdown `json:"description,omitempty"`
	// Used for footnotes or explanatory notes
	Note []Annotation `json:"note,omitempty"`
	// Whether the characteristic is an inclusion criterion or exclusion criterion
	Exclude *bool `json:"exclude,omitempty"`
	// Defines the characteristic (without using type and value) by a Reference
	DefinitionReference *custom.Reference[EvidenceVariableCharacteristicDefinitionReference] `json:"definitionReference,omitempty"`
	// Defines the characteristic (without using type and value) by a Canonical
	DefinitionCanonical *custom.Canonical[EvidenceVariableCharacteristicDefinitionCanonical] `json:"definitionCanonical,omitempty"`
	// Defines the characteristic (without using type and value) by a CodeableConcept
	DefinitionCodeableConcept *CodeableConcept `json:"definitionCodeableConcept,omitempty"`
	// Defines the characteristic (without using type and value) by an expression
	DefinitionExpression *Expression `json:"definitionExpression,omitempty"`
	// Defines the characteristic (without using type and value) by an id
	DefinitionId *custom.ID `json:"definitionId,omitempty"`
	// Defines the characteristic using type and value
	DefinitionByTypeAndValue *EvidenceVariableCharacteristicDefinitionByTypeAndValue `json:"definitionByTypeAndValue,omitempty"`
	// Used to specify how two or more characteristics are combined
	DefinitionByCombination *EvidenceVariableCharacteristicDefinitionByCombination `json:"definitionByCombination,omitempty"`
	// Number of occurrences meeting the characteristic
	InstancesQuantity *Quantity `json:"instancesQuantity,omitempty"`
	// Number of occurrences meeting the characteristic
	InstancesRange *Range `json:"instancesRange,omitempty"`
	// Length of time in which the characteristic is met
	DurationQuantity *Quantity `json:"durationQuantity,omitempty"`
	// Length of time in which the characteristic is met
	DurationRange *Range `json:"durationRange,omitempty"`
	// Timing in which the characteristic is determined
	TimeFromEvent []EvidenceVariableCharacteristicTimeFromEvent `json:"timeFromEvent,omitempty"`
}

type EvidenceVariableCharacteristicDefinitionByTypeAndValue struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Expresses the type of characteristic
	Type CodeableConcept `json:"type"`
	// Method for how the characteristic value was determined
	Method []CodeableConcept `json:"method,omitempty"`
	// Device used for determining characteristic
	Device *custom.Reference[EvidenceVariableCharacteristicDefinitionByTypeAndValueDevice] `json:"device,omitempty"`
	// Defines the characteristic when coupled with characteristic.type
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Defines the characteristic when coupled with characteristic.type
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Defines the characteristic when coupled with characteristic.type
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Defines the characteristic when coupled with characteristic.type
	ValueRange *Range `json:"valueRange,omitempty"`
	// Defines the characteristic when coupled with characteristic.type
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// Defines the characteristic when coupled with characteristic.type
	ValueId *custom.ID `json:"valueId,omitempty"`
	// Reference point for valueQuantity or valueRange
	Offset *CodeableConcept `json:"offset,omitempty"`
}

type OtherEvidenceVariable EvidenceVariable

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
	return "EvidenceVariable"
}

func (e EvidenceVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceVariable
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherEvidenceVariable: OtherEvidenceVariable(e), ResourceType: e.ResourceType()})
}

func UnmarshalEvidenceVariable(b []byte) (res EvidenceVariable, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
