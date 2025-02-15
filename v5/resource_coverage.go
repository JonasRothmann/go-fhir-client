// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Coverage
type Coverage struct {
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
	// Business identifier(s) for this coverage
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// insurance | self-pay | other
	Kind custom.Code `bson:"kind" json:"kind"`
	// Self-pay parties and responsibility
	PaymentBy []CoveragePaymentBy `bson:"paymentBy,omitempty" json:"paymentBy,omitempty"`
	// Coverage category such as medical or accident
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Owner of the policy
	PolicyHolder *custom.Reference[CoveragePolicyHolder] `bson:"policyHolder,omitempty" json:"policyHolder,omitempty"`
	// Subscriber to the policy
	Subscriber *custom.Reference[CoverageSubscriber] `bson:"subscriber,omitempty" json:"subscriber,omitempty"`
	// ID assigned to the subscriber
	SubscriberId []Identifier `bson:"subscriberId,omitempty" json:"subscriberId,omitempty"`
	// Plan beneficiary
	Beneficiary custom.Reference[Patient] `bson:"beneficiary" json:"beneficiary"`
	// Dependent number
	Dependent *string `bson:"dependent,omitempty" json:"dependent,omitempty"`
	// Beneficiary relationship to the subscriber
	Relationship *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	// Coverage start and end dates
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Issuer of the policy
	Insurer *custom.Reference[Organization] `bson:"insurer,omitempty" json:"insurer,omitempty"`
	// Additional coverage classifications
	Class []CoverageClass `bson:"class,omitempty" json:"class,omitempty"`
	// Relative order of the coverage
	Order *uint32 `bson:"order,omitempty" json:"order,omitempty"`
	// Insurer network
	Network *string `bson:"network,omitempty" json:"network,omitempty"`
	// Patient payments for services/products
	CostToBeneficiary []CoverageCostToBeneficiary `bson:"costToBeneficiary,omitempty" json:"costToBeneficiary,omitempty"`
	// Reimbursement to insurer
	Subrogation *bool `bson:"subrogation,omitempty" json:"subrogation,omitempty"`
	// Contract details
	Contract []custom.Reference[Contract] `bson:"contract,omitempty" json:"contract,omitempty"`
	// Insurance plan details
	InsurancePlan *custom.Reference[InsurancePlan] `bson:"insurancePlan,omitempty" json:"insurancePlan,omitempty"`
}

type CoveragePaymentBy struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Parties performing self-payment
	Party custom.Reference[CoveragePaymentByParty] `bson:"party" json:"party"`
	// Party's responsibility
	Responsibility *string `bson:"responsibility,omitempty" json:"responsibility,omitempty"`
}

type CoverageClass struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of class such as 'group' or 'plan'
	Type CodeableConcept `bson:"type" json:"type"`
	// Value associated with the type
	Value Identifier `bson:"value" json:"value"`
	// Human readable description of the type and value
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
}

type CoverageCostToBeneficiary struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Cost category
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Benefit classification
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// In or out of network
	Network *CodeableConcept `bson:"network,omitempty" json:"network,omitempty"`
	// Individual or family
	Unit *CodeableConcept `bson:"unit,omitempty" json:"unit,omitempty"`
	// Annual or lifetime
	Term *CodeableConcept `bson:"term,omitempty" json:"term,omitempty"`
	// The amount or percentage due from the beneficiary
	Value *Quantity `bson:"value,omitempty" json:"value,omitempty"`
	// Exceptions for patient payments
	Exception []CoverageCostToBeneficiaryException `bson:"exception,omitempty" json:"exception,omitempty"`
}

type CoverageCostToBeneficiaryException struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Exception category
	Type CodeableConcept `bson:"type" json:"type"`
	// The effective period of the exception
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type CoveragePaymentByParty interface {
	gofhirclient.Resource

	Is_CoveragePaymentByParty()
}

func (p Patient) Is_CoveragePaymentByParty()       {}
func (r RelatedPerson) Is_CoveragePaymentByParty() {}
func (o Organization) Is_CoveragePaymentByParty()  {}

type CoveragePolicyHolder interface {
	gofhirclient.Resource

	Is_CoveragePolicyHolder()
}

func (p Patient) Is_CoveragePolicyHolder()       {}
func (r RelatedPerson) Is_CoveragePolicyHolder() {}
func (o Organization) Is_CoveragePolicyHolder()  {}

type CoverageSubscriber interface {
	gofhirclient.Resource

	Is_CoverageSubscriber()
}

func (p Patient) Is_CoverageSubscriber()       {}
func (r RelatedPerson) Is_CoverageSubscriber() {}

func (c Coverage) ResourceType() string {
	return "StructureDefinition"
}
