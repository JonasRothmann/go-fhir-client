// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinition struct {
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
	// Business identifier for this product. Could be an MPID
	Identifier []Identifier `json:"identifier,omitempty"`
	// Regulatory type, e.g. Investigational or Authorized
	Type *CodeableConcept `json:"type,omitempty"`
	// If this medicine applies to human or veterinary uses
	Domain *CodeableConcept `json:"domain,omitempty"`
	// A business identifier relating to a specific version of the product
	Version *string `json:"version,omitempty"`
	// The status within the lifecycle of this product record
	Status *CodeableConcept `json:"status,omitempty"`
	// The date at which the given status became applicable
	StatusDate *custom.DateTime `json:"statusDate,omitempty"`
	// General description of this product
	Description *custom.Markdown `json:"description,omitempty"`
	// The dose form for a single part product, or combined form of a multiple part product
	CombinedPharmaceuticalDoseForm *CodeableConcept `json:"combinedPharmaceuticalDoseForm,omitempty"`
	// The path by which the product is taken into or makes contact with the body
	Route []CodeableConcept `json:"route,omitempty"`
	// Description of indication(s) for this product, used when structured indications are not required
	Indication *custom.Markdown `json:"indication,omitempty"`
	// The legal status of supply of the medicinal product as classified by the regulator
	LegalStatusOfSupply *CodeableConcept `json:"legalStatusOfSupply,omitempty"`
	// Whether the Medicinal Product is subject to additional monitoring for regulatory reasons
	AdditionalMonitoringIndicator *CodeableConcept `json:"additionalMonitoringIndicator,omitempty"`
	// Whether the Medicinal Product is subject to special measures for regulatory reasons
	SpecialMeasures []CodeableConcept `json:"specialMeasures,omitempty"`
	// If authorised for use in children
	PediatricUseIndicator *CodeableConcept `json:"pediatricUseIndicator,omitempty"`
	// Allows the product to be classified by various systems
	Classification []CodeableConcept `json:"classification,omitempty"`
	// Marketing status of the medicinal product, in contrast to marketing authorization
	MarketingStatus []MarketingStatus `json:"marketingStatus,omitempty"`
	// Package type for the product
	PackagedMedicinalProduct []CodeableConcept `json:"packagedMedicinalProduct,omitempty"`
	// Types of medicinal manufactured items and/or devices that this product consists of, such as tablets, capsule, or syringes
	ComprisedOf []custom.Reference[MedicinalProductDefinitionComprisedOf] `json:"comprisedOf,omitempty"`
	// The ingredients of this medicinal product - when not detailed in other resources
	Ingredient []CodeableConcept `json:"ingredient,omitempty"`
	// Any component of the drug product which is not the chemical entity defined as the drug substance, or an excipient in the drug product
	Impurity []custom.CodeableReference[SubstanceDefinition] `json:"impurity,omitempty"`
	// Additional documentation about the medicinal product
	AttachedDocument []custom.Reference[DocumentReference] `json:"attachedDocument,omitempty"`
	// A master file for the medicinal product (e.g. Pharmacovigilance System Master File)
	MasterFile []custom.Reference[DocumentReference] `json:"masterFile,omitempty"`
	// A product specific contact, person (in a role), or an organization
	Contact []MedicinalProductDefinitionContact `json:"contact,omitempty"`
	// Clinical trials or studies that this product is involved in
	ClinicalTrial []custom.Reference[ResearchStudy] `json:"clinicalTrial,omitempty"`
	// A code that this product is known by, within some formal terminology
	Code []Coding `json:"code,omitempty"`
	// The product's name, including full name and possibly coded parts
	Name []MedicinalProductDefinitionName `json:"name"`
	// Reference to another product, e.g. for linking authorised to investigational product
	CrossReference []MedicinalProductDefinitionCrossReference `json:"crossReference,omitempty"`
	// A manufacturing or administrative process for the medicinal product
	Operation []MedicinalProductDefinitionOperation `json:"operation,omitempty"`
	// Key product features such as "sugar free", "modified release"
	Characteristic []MedicinalProductDefinitionCharacteristic `json:"characteristic,omitempty"`
}

type MedicinalProductDefinitionContact struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Allows the contact to be classified, for example QPPV, Pharmacovigilance Enquiry Information
	Type *CodeableConcept `json:"type,omitempty"`
	// A product specific contact, person (in a role), or an organization
	Contact custom.Reference[MedicinalProductDefinitionContactContact] `json:"contact"`
}

type MedicinalProductDefinitionName struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The full product name
	ProductName string `json:"productName"`
	// Type of product name, such as rINN, BAN, Proprietary, Non-Proprietary
	Type *CodeableConcept `json:"type,omitempty"`
	// Coding words or phrases of the name
	Part []MedicinalProductDefinitionNamePart `json:"part,omitempty"`
	// Country and jurisdiction where the name applies
	Usage []MedicinalProductDefinitionNameUsage `json:"usage,omitempty"`
}

type MedicinalProductDefinitionNamePart struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A fragment of a product name
	Part string `json:"part"`
	// Identifying type for this part of the name (e.g. strength part)
	Type CodeableConcept `json:"type"`
}

type MedicinalProductDefinitionNameUsage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Country code for where this name applies
	Country CodeableConcept `json:"country"`
	// Jurisdiction code for where this name applies
	Jurisdiction *CodeableConcept `json:"jurisdiction,omitempty"`
	// Language code for this name
	Language CodeableConcept `json:"language"`
}

type MedicinalProductDefinitionCrossReference struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to another product, e.g. for linking authorised to investigational product
	Product custom.CodeableReference[MedicinalProductDefinition] `json:"product"`
	// The type of relationship, for instance branded to generic or virtual to actual product
	Type *CodeableConcept `json:"type,omitempty"`
}

type MedicinalProductDefinitionOperation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of manufacturing operation e.g. manufacturing itself, re-packaging
	Type *custom.CodeableReference[MedicinalProductDefinitionOperationType] `json:"type,omitempty"`
	// Date range of applicability
	EffectiveDate *Period `json:"effectiveDate,omitempty"`
	// The organization responsible for the particular process, e.g. the manufacturer or importer
	Organization []custom.Reference[Organization] `json:"organization,omitempty"`
	// Specifies whether this process is considered proprietary or confidential
	ConfidentialityIndicator *CodeableConcept `json:"confidentialityIndicator,omitempty"`
}

type MedicinalProductDefinitionCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A code expressing the type of characteristic
	Type CodeableConcept `json:"type"`
	// A value for the characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// A value for the characteristic
	ValueMarkdown *custom.Markdown `json:"valueMarkdown,omitempty"`
	// A value for the characteristic
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// A value for the characteristic
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// A value for the characteristic
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// A value for the characteristic
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// A value for the characteristic
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
}

type OtherMedicinalProductDefinition MedicinalProductDefinition

type MedicinalProductDefinitionComprisedOf interface {
	gofhirclient.Resource

	Is_MedicinalProductDefinitionComprisedOf()
}

func (m ManufacturedItemDefinition) Is_MedicinalProductDefinitionComprisedOf() {}
func (d DeviceDefinition) Is_MedicinalProductDefinitionComprisedOf()           {}

type MedicinalProductDefinitionContactContact interface {
	gofhirclient.Resource

	Is_MedicinalProductDefinitionContactContact()
}

func (o Organization) Is_MedicinalProductDefinitionContactContact()     {}
func (p PractitionerRole) Is_MedicinalProductDefinitionContactContact() {}

type MedicinalProductDefinitionOperationType interface {
	gofhirclient.Resource

	Is_MedicinalProductDefinitionOperationType()
}

func (a ActivityDefinition) Is_MedicinalProductDefinitionOperationType() {}
func (p PlanDefinition) Is_MedicinalProductDefinitionOperationType()     {}

func (m MedicinalProductDefinition) ResourceType() string {
	return "MedicinalProductDefinition"
}

func (m MedicinalProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMedicinalProductDefinition: OtherMedicinalProductDefinition(m), ResourceType: m.ResourceType()})
}

func UnmarshalMedicinalProductDefinition(b []byte) (res MedicinalProductDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
