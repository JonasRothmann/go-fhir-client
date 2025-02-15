// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/InsurancePlan
type InsurancePlanCoverageBenefit struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of benefit
	Type CodeableConcept `bson:"type" json:"type"`
	// Referral requirements
	Requirement *string `bson:"requirement,omitempty" json:"requirement,omitempty"`
	// Benefit limits
	Limit []InsurancePlanCoverageBenefitLimit `bson:"limit,omitempty" json:"limit,omitempty"`
}

type InsurancePlanPlanGeneralCost struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of cost
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Number of enrollees
	GroupSize *uint32 `bson:"groupSize,omitempty" json:"groupSize,omitempty"`
	// Cost value
	Cost *Money `bson:"cost,omitempty" json:"cost,omitempty"`
	// Additional cost information
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty"`
}

type InsurancePlanCoverageBenefitLimit struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Maximum value allowed
	Value *Quantity `bson:"value,omitempty" json:"value,omitempty"`
	// Benefit limit details
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}

type InsurancePlanPlan struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Business Identifier for Product
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Type of plan
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Where product applies
	CoverageArea []custom.Reference[Location] `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	// What networks provide coverage
	Network []custom.Reference[Organization] `bson:"network,omitempty" json:"network,omitempty"`
	// Overall costs
	GeneralCost []InsurancePlanPlanGeneralCost `bson:"generalCost,omitempty" json:"generalCost,omitempty"`
	// Specific costs
	SpecificCost []InsurancePlanPlanSpecificCost `bson:"specificCost,omitempty" json:"specificCost,omitempty"`
}

type InsurancePlanPlanSpecificCost struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// General category of benefit
	Category CodeableConcept `bson:"category" json:"category"`
	// Benefits list
	Benefit []InsurancePlanPlanSpecificCostBenefit `bson:"benefit,omitempty" json:"benefit,omitempty"`
}

type InsurancePlanPlanSpecificCostBenefit struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of specific benefit
	Type CodeableConcept `bson:"type" json:"type"`
	// List of the costs
	Cost []InsurancePlanPlanSpecificCostBenefitCost `bson:"cost,omitempty" json:"cost,omitempty"`
}

type InsurancePlanPlanSpecificCostBenefitCost struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of cost
	Type CodeableConcept `bson:"type" json:"type"`
	// in-network | out-of-network | other
	Applicability *CodeableConcept `bson:"applicability,omitempty" json:"applicability,omitempty"`
	// Additional information about the cost
	Qualifiers []CodeableConcept `bson:"qualifiers,omitempty" json:"qualifiers,omitempty"`
	// The actual cost value
	Value *Quantity `bson:"value,omitempty" json:"value,omitempty"`
}

type InsurancePlan struct {
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
	// Business Identifier for Product
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// draft | active | retired | unknown
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Kind of product
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Official name
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Alternate names
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty"`
	// When the product is available
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Product issuer
	OwnedBy *custom.Reference[Organization] `bson:"ownedBy,omitempty" json:"ownedBy,omitempty"`
	// Product administrator
	AdministeredBy *custom.Reference[Organization] `bson:"administeredBy,omitempty" json:"administeredBy,omitempty"`
	// Where product applies
	CoverageArea []custom.Reference[Location] `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	// Official contact details relevant to the health insurance plan/product
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Technical endpoint
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// What networks are Included
	Network []custom.Reference[Organization] `bson:"network,omitempty" json:"network,omitempty"`
	// Coverage details
	Coverage []InsurancePlanCoverage `bson:"coverage,omitempty" json:"coverage,omitempty"`
	// Plan details
	Plan []InsurancePlanPlan `bson:"plan,omitempty" json:"plan,omitempty"`
}

type InsurancePlanCoverage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of coverage
	Type CodeableConcept `bson:"type" json:"type"`
	// What networks provide coverage
	Network []custom.Reference[Organization] `bson:"network,omitempty" json:"network,omitempty"`
	// List of benefits
	Benefit []InsurancePlanCoverageBenefit `bson:"benefit" json:"benefit"`
}

func (i InsurancePlan) ResourceType() string {
	return "StructureDefinition"
}
