// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/InsurancePlan
type InsurancePlanPlanGeneralCost struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of cost
	Type *CodeableConcept `json:"type,omitempty"`
	// Number of enrollees
	GroupSize *uint32 `json:"groupSize,omitempty"`
	// Cost value
	Cost *Money `json:"cost,omitempty"`
	// Additional cost information
	Comment *string `json:"comment,omitempty"`
}

type InsurancePlanPlanSpecificCost struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// General category of benefit
	Category CodeableConcept `json:"category"`
	// Benefits list
	Benefit []InsurancePlanPlanSpecificCostBenefit `json:"benefit,omitempty"`
}

type InsurancePlanPlanSpecificCostBenefit struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of specific benefit
	Type CodeableConcept `json:"type"`
	// List of the costs
	Cost []InsurancePlanPlanSpecificCostBenefitCost `json:"cost,omitempty"`
}

type InsurancePlanPlanSpecificCostBenefitCost struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of cost
	Type CodeableConcept `json:"type"`
	// in-network | out-of-network | other
	Applicability *CodeableConcept `json:"applicability,omitempty"`
	// Additional information about the cost
	Qualifiers []CodeableConcept `json:"qualifiers,omitempty"`
	// The actual cost value
	Value *Quantity `json:"value,omitempty"`
}

type InsurancePlan struct {
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
	// Business Identifier for Product
	Identifier []Identifier `json:"identifier,omitempty"`
	// draft | active | retired | unknown
	Status *custom.Code `json:"status,omitempty"`
	// Kind of product
	Type []CodeableConcept `json:"type,omitempty"`
	// Official name
	Name *string `json:"name,omitempty"`
	// Alternate names
	Alias []string `json:"alias,omitempty"`
	// When the product is available
	Period *Period `json:"period,omitempty"`
	// Product issuer
	OwnedBy *custom.Reference[Organization] `json:"ownedBy,omitempty"`
	// Product administrator
	AdministeredBy *custom.Reference[Organization] `json:"administeredBy,omitempty"`
	// Where product applies
	CoverageArea []custom.Reference[Location] `json:"coverageArea,omitempty"`
	// Official contact details relevant to the health insurance plan/product
	Contact []ExtendedContactDetail `json:"contact,omitempty"`
	// Technical endpoint
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
	// What networks are Included
	Network []custom.Reference[Organization] `json:"network,omitempty"`
	// Coverage details
	Coverage []InsurancePlanCoverage `json:"coverage,omitempty"`
	// Plan details
	Plan []InsurancePlanPlan `json:"plan,omitempty"`
}

type InsurancePlanCoverage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of coverage
	Type CodeableConcept `json:"type"`
	// What networks provide coverage
	Network []custom.Reference[Organization] `json:"network,omitempty"`
	// List of benefits
	Benefit []InsurancePlanCoverageBenefit `json:"benefit"`
}

type InsurancePlanCoverageBenefit struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of benefit
	Type CodeableConcept `json:"type"`
	// Referral requirements
	Requirement *string `json:"requirement,omitempty"`
	// Benefit limits
	Limit []InsurancePlanCoverageBenefitLimit `json:"limit,omitempty"`
}

type InsurancePlanCoverageBenefitLimit struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Maximum value allowed
	Value *Quantity `json:"value,omitempty"`
	// Benefit limit details
	Code *CodeableConcept `json:"code,omitempty"`
}

type InsurancePlanPlan struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Business Identifier for Product
	Identifier []Identifier `json:"identifier,omitempty"`
	// Type of plan
	Type *CodeableConcept `json:"type,omitempty"`
	// Where product applies
	CoverageArea []custom.Reference[Location] `json:"coverageArea,omitempty"`
	// What networks provide coverage
	Network []custom.Reference[Organization] `json:"network,omitempty"`
	// Overall costs
	GeneralCost []InsurancePlanPlanGeneralCost `json:"generalCost,omitempty"`
	// Specific costs
	SpecificCost []InsurancePlanPlanSpecificCost `json:"specificCost,omitempty"`
}

type OtherInsurancePlan InsurancePlan

func (i InsurancePlan) ResourceType() string {
	return "InsurancePlan"
}

func (i InsurancePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInsurancePlan
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherInsurancePlan: OtherInsurancePlan(i), ResourceType: i.ResourceType()})
}

func UnmarshalInsurancePlan(b []byte) (res InsurancePlan, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
