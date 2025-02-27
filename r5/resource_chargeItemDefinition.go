// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ChargeItemDefinition
type ChargeItemDefinition struct {
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
	// Canonical identifier for this charge item definition, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the charge item definition
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the charge item definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this charge item definition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this charge item definition (human friendly)
	Title *string `json:"title,omitempty"`
	// Underlying externally-defined charge item definition
	DerivedFromUri []custom.URI `json:"derivedFromUri,omitempty"`
	// A larger definition of which this particular definition is a component or step
	PartOf []custom.Canonical[ChargeItemDefinition] `json:"partOf,omitempty"`
	// Completed or terminated request(s) whose function is taken by this new request
	Replaces []custom.Canonical[ChargeItemDefinition] `json:"replaces,omitempty"`
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
	// Natural language description of the charge item definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for charge item definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this charge item definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the charge item definition was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the charge item definition was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// Billing code or product type this definition applies to
	Code *CodeableConcept `json:"code,omitempty"`
	// Instances this definition applies to
	Instance []custom.Reference[ChargeItemDefinitionInstance] `json:"instance,omitempty"`
	// Whether or not the billing code is applicable
	Applicability []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`
	// Group of properties which are applicable under the same conditions
	PropertyGroup []ChargeItemDefinitionPropertyGroup `json:"propertyGroup,omitempty"`
}

type ChargeItemDefinitionApplicability struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Boolean-valued expression
	Condition *Expression `json:"condition,omitempty"`
	// When the charge item definition is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// Reference to / quotation of the external source of the group of properties
	RelatedArtifact *RelatedArtifact `json:"relatedArtifact,omitempty"`
}

type ChargeItemDefinitionPropertyGroup struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Conditions under which the priceComponent is applicable
	Applicability []interface{} `json:"applicability,omitempty"`
	// Components of total line item price
	PriceComponent []MonetaryComponent `json:"priceComponent,omitempty"`
}

type OtherChargeItemDefinition ChargeItemDefinition

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
	return "ChargeItemDefinition"
}

func (c ChargeItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItemDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherChargeItemDefinition: OtherChargeItemDefinition(c), ResourceType: c.ResourceType()})
}

func UnmarshalChargeItemDefinition(b []byte) (res ChargeItemDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
