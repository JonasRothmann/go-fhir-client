// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinition struct {
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
	// Canonical identifier for this charge item definition, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the charge item definition
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the charge item definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this charge item definition (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this charge item definition (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Underlying externally-defined charge item definition
	DerivedFromUri []custom.URI `bson:"derivedFromUri,omitempty" json:"derivedFromUri,omitempty"`
	// A larger definition of which this particular definition is a component or step
	PartOf []custom.Canonical[ChargeItemDefinition] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// Completed or terminated request(s) whose function is taken by this new request
	Replaces []custom.Canonical[ChargeItemDefinition] `bson:"replaces,omitempty" json:"replaces,omitempty"`
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
	// Natural language description of the charge item definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for charge item definition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this charge item definition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the charge item definition was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the charge item definition was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// Billing code or product type this definition applies to
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Instances this definition applies to
	Instance []custom.Reference[ChargeItemDefinitionInstance] `bson:"instance,omitempty" json:"instance,omitempty"`
	// Whether or not the billing code is applicable
	Applicability []ChargeItemDefinitionApplicability `bson:"applicability,omitempty" json:"applicability,omitempty"`
	// Group of properties which are applicable under the same conditions
	PropertyGroup []ChargeItemDefinitionPropertyGroup `bson:"propertyGroup,omitempty" json:"propertyGroup,omitempty"`
}

type ChargeItemDefinitionApplicability struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Boolean-valued expression
	Condition *Expression `bson:"condition,omitempty" json:"condition,omitempty"`
	// When the charge item definition is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// Reference to / quotation of the external source of the group of properties
	RelatedArtifact *RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
}

type ChargeItemDefinitionPropertyGroup struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Conditions under which the priceComponent is applicable
	Applicability []interface{} `bson:"applicability,omitempty" json:"applicability,omitempty"`
	// Components of total line item price
	PriceComponent []MonetaryComponent `bson:"priceComponent,omitempty" json:"priceComponent,omitempty"`
}

type ChargeItemDefinitionInstance interface {
	gofhirclient.Resource

	Is_ChargeItemDefinitionInstance()
}

func (m Medication) Is_ChargeItemDefinitionInstance()         {}
func (s Substance) Is_ChargeItemDefinitionInstance()          {}
func (d Device) Is_ChargeItemDefinitionInstance()             {}
func (d DeviceDefinition) Is_ChargeItemDefinitionInstance()   {}
func (a ActivityDefinition) Is_ChargeItemDefinitionInstance() {}
func (p PlanDefinition) Is_ChargeItemDefinitionInstance()     {}
func (h HealthcareService) Is_ChargeItemDefinitionInstance()  {}

func (c ChargeItemDefinition) ResourceType() string {
	return "StructureDefinition"
}
