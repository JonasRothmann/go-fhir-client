// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicinalProductDefinition
type MedicinalProductDefinitionName struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The full product name
	ProductName string `bson:"productName" json:"productName"`
	// Type of product name, such as rINN, BAN, Proprietary, Non-Proprietary
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Coding words or phrases of the name
	Part []MedicinalProductDefinitionNamePart `bson:"part,omitempty" json:"part,omitempty"`
	// Country and jurisdiction where the name applies
	Usage []MedicinalProductDefinitionNameUsage `bson:"usage,omitempty" json:"usage,omitempty"`
}

type MedicinalProductDefinitionNamePart struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A fragment of a product name
	Part string `bson:"part" json:"part"`
	// Identifying type for this part of the name (e.g. strength part)
	Type CodeableConcept `bson:"type" json:"type"`
}

type MedicinalProductDefinitionNameUsage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Country code for where this name applies
	Country CodeableConcept `bson:"country" json:"country"`
	// Jurisdiction code for where this name applies
	Jurisdiction *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Language code for this name
	Language CodeableConcept `bson:"language" json:"language"`
}

type MedicinalProductDefinitionCrossReference struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to another product, e.g. for linking authorised to investigational product
	Product custom.CodeableReference[MedicinalProductDefinition] `bson:"product" json:"product"`
	// The type of relationship, for instance branded to generic or virtual to actual product
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}

type MedicinalProductDefinitionOperation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of manufacturing operation e.g. manufacturing itself, re-packaging
	Type *custom.CodeableReference[MedicinalProductDefinitionOperationType] `bson:"type,omitempty" json:"type,omitempty"`
	// Date range of applicability
	EffectiveDate *Period `bson:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
	// The organization responsible for the particular process, e.g. the manufacturer or importer
	Organization []custom.Reference[Organization] `bson:"organization,omitempty" json:"organization,omitempty"`
	// Specifies whether this process is considered proprietary or confidential
	ConfidentialityIndicator *CodeableConcept `bson:"confidentialityIndicator,omitempty" json:"confidentialityIndicator,omitempty"`
}

type MedicinalProductDefinitionCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A code expressing the type of characteristic
	Type CodeableConcept `bson:"type" json:"type"`
	// A value for the characteristic
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

type MedicinalProductDefinition struct {
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
	// Business identifier for this product. Could be an MPID
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Regulatory type, e.g. Investigational or Authorized
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// If this medicine applies to human or veterinary uses
	Domain *CodeableConcept `bson:"domain,omitempty" json:"domain,omitempty"`
	// A business identifier relating to a specific version of the product
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// The status within the lifecycle of this product record
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// The date at which the given status became applicable
	StatusDate *custom.DateTime `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	// General description of this product
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The dose form for a single part product, or combined form of a multiple part product
	CombinedPharmaceuticalDoseForm *CodeableConcept `bson:"combinedPharmaceuticalDoseForm,omitempty" json:"combinedPharmaceuticalDoseForm,omitempty"`
	// The path by which the product is taken into or makes contact with the body
	Route []CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	// Description of indication(s) for this product, used when structured indications are not required
	Indication *custom.Markdown `bson:"indication,omitempty" json:"indication,omitempty"`
	// The legal status of supply of the medicinal product as classified by the regulator
	LegalStatusOfSupply *CodeableConcept `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty"`
	// Whether the Medicinal Product is subject to additional monitoring for regulatory reasons
	AdditionalMonitoringIndicator *CodeableConcept `bson:"additionalMonitoringIndicator,omitempty" json:"additionalMonitoringIndicator,omitempty"`
	// Whether the Medicinal Product is subject to special measures for regulatory reasons
	SpecialMeasures []CodeableConcept `bson:"specialMeasures,omitempty" json:"specialMeasures,omitempty"`
	// If authorised for use in children
	PediatricUseIndicator *CodeableConcept `bson:"pediatricUseIndicator,omitempty" json:"pediatricUseIndicator,omitempty"`
	// Allows the product to be classified by various systems
	Classification []CodeableConcept `bson:"classification,omitempty" json:"classification,omitempty"`
	// Marketing status of the medicinal product, in contrast to marketing authorization
	MarketingStatus []MarketingStatus `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty"`
	// Package type for the product
	PackagedMedicinalProduct []CodeableConcept `bson:"packagedMedicinalProduct,omitempty" json:"packagedMedicinalProduct,omitempty"`
	// Types of medicinal manufactured items and/or devices that this product consists of, such as tablets, capsule, or syringes
	ComprisedOf []custom.Reference[MedicinalProductDefinitionComprisedOf] `bson:"comprisedOf,omitempty" json:"comprisedOf,omitempty"`
	// The ingredients of this medicinal product - when not detailed in other resources
	Ingredient []CodeableConcept `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	// Any component of the drug product which is not the chemical entity defined as the drug substance, or an excipient in the drug product
	Impurity []custom.CodeableReference[SubstanceDefinition] `bson:"impurity,omitempty" json:"impurity,omitempty"`
	// Additional documentation about the medicinal product
	AttachedDocument []custom.Reference[DocumentReference] `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty"`
	// A master file for the medicinal product (e.g. Pharmacovigilance System Master File)
	MasterFile []custom.Reference[DocumentReference] `bson:"masterFile,omitempty" json:"masterFile,omitempty"`
	// A product specific contact, person (in a role), or an organization
	Contact []MedicinalProductDefinitionContact `bson:"contact,omitempty" json:"contact,omitempty"`
	// Clinical trials or studies that this product is involved in
	ClinicalTrial []custom.Reference[ResearchStudy] `bson:"clinicalTrial,omitempty" json:"clinicalTrial,omitempty"`
	// A code that this product is known by, within some formal terminology
	Code []Coding `bson:"code,omitempty" json:"code,omitempty"`
	// The product's name, including full name and possibly coded parts
	Name []MedicinalProductDefinitionName `bson:"name" json:"name"`
	// Reference to another product, e.g. for linking authorised to investigational product
	CrossReference []MedicinalProductDefinitionCrossReference `bson:"crossReference,omitempty" json:"crossReference,omitempty"`
	// A manufacturing or administrative process for the medicinal product
	Operation []MedicinalProductDefinitionOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	// Key product features such as "sugar free", "modified release"
	Characteristic []MedicinalProductDefinitionCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
}

type MedicinalProductDefinitionContact struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Allows the contact to be classified, for example QPPV, Pharmacovigilance Enquiry Information
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// A product specific contact, person (in a role), or an organization
	Contact custom.Reference[MedicinalProductDefinitionContactContact] `bson:"contact" json:"contact"`
}

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
	return "StructureDefinition"
}
