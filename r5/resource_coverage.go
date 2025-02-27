// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Coverage
type CoverageCostToBeneficiaryException struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Exception category
	Type CodeableConcept `json:"type"`
	// The effective period of the exception
	Period *Period `json:"period,omitempty"`
}

type Coverage struct {
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
	// Business identifier(s) for this coverage
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `json:"status"`
	// insurance | self-pay | other
	Kind custom.Code `json:"kind"`
	// Self-pay parties and responsibility
	PaymentBy []CoveragePaymentBy `json:"paymentBy,omitempty"`
	// Coverage category such as medical or accident
	Type *CodeableConcept `json:"type,omitempty"`
	// Owner of the policy
	PolicyHolder *custom.Reference[CoveragePolicyHolder] `json:"policyHolder,omitempty"`
	// Subscriber to the policy
	Subscriber *custom.Reference[CoverageSubscriber] `json:"subscriber,omitempty"`
	// ID assigned to the subscriber
	SubscriberId []Identifier `json:"subscriberId,omitempty"`
	// Plan beneficiary
	Beneficiary custom.Reference[Patient] `json:"beneficiary"`
	// Dependent number
	Dependent *string `json:"dependent,omitempty"`
	// Beneficiary relationship to the subscriber
	Relationship *CodeableConcept `json:"relationship,omitempty"`
	// Coverage start and end dates
	Period *Period `json:"period,omitempty"`
	// Issuer of the policy
	Insurer *custom.Reference[Organization] `json:"insurer,omitempty"`
	// Additional coverage classifications
	Class []CoverageClass `json:"class,omitempty"`
	// Relative order of the coverage
	Order *uint32 `json:"order,omitempty"`
	// Insurer network
	Network *string `json:"network,omitempty"`
	// Patient payments for services/products
	CostToBeneficiary []CoverageCostToBeneficiary `json:"costToBeneficiary,omitempty"`
	// Reimbursement to insurer
	Subrogation *bool `json:"subrogation,omitempty"`
	// Contract details
	Contract []custom.Reference[Contract] `json:"contract,omitempty"`
	// Insurance plan details
	InsurancePlan *custom.Reference[InsurancePlan] `json:"insurancePlan,omitempty"`
}

type CoveragePaymentBy struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Parties performing self-payment
	Party custom.Reference[CoveragePaymentByParty] `json:"party"`
	// Party's responsibility
	Responsibility *string `json:"responsibility,omitempty"`
}

type CoverageClass struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of class such as 'group' or 'plan'
	Type CodeableConcept `json:"type"`
	// Value associated with the type
	Value Identifier `json:"value"`
	// Human readable description of the type and value
	Name *string `json:"name,omitempty"`
}

type CoverageCostToBeneficiary struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Cost category
	Type *CodeableConcept `json:"type,omitempty"`
	// Benefit classification
	Category *CodeableConcept `json:"category,omitempty"`
	// In or out of network
	Network *CodeableConcept `json:"network,omitempty"`
	// Individual or family
	Unit *CodeableConcept `json:"unit,omitempty"`
	// Annual or lifetime
	Term *CodeableConcept `json:"term,omitempty"`
	// The amount or percentage due from the beneficiary
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The amount or percentage due from the beneficiary
	ValueMoney *Money `json:"valueMoney,omitempty"`
	// Exceptions for patient payments
	Exception []CoverageCostToBeneficiaryException `json:"exception,omitempty"`
}

type OtherCoverage Coverage

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
	return "Coverage"
}

func (c Coverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverage
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCoverage: OtherCoverage(c), ResourceType: c.ResourceType()})
}

func UnmarshalCoverage(b []byte) (res Coverage, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
