// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Account
type AccountCoverage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The party(s), such as insurances, that may contribute to the payment of this account
	Coverage custom.Reference[Coverage] `bson:"coverage" json:"coverage"`
	// The priority of the coverage in the context of this account
	Priority *uint32 `bson:"priority,omitempty" json:"priority,omitempty"`
}

type AccountGuarantor struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Responsible entity
	Party custom.Reference[AccountGuarantorParty] `bson:"party" json:"party"`
	// Credit or other hold applied
	OnHold *bool `bson:"onHold,omitempty" json:"onHold,omitempty"`
	// Guarantee account during
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type AccountDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Ranking of the diagnosis (for each type)
	Sequence *uint32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	// The diagnosis relevant to the account
	Condition custom.CodeableReference[Condition] `bson:"condition" json:"condition"`
	// Date of the diagnosis (when coded diagnosis)
	DateOfDiagnosis *custom.DateTime `bson:"dateOfDiagnosis,omitempty" json:"dateOfDiagnosis,omitempty"`
	// Type that this diagnosis has relevant to the account (e.g. admission, billing, discharge …)
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Diagnosis present on Admission
	OnAdmission *bool `bson:"onAdmission,omitempty" json:"onAdmission,omitempty"`
	// Package Code specific for billing
	PackageCode []CodeableConcept `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
}

type AccountProcedure struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Ranking of the procedure (for each type)
	Sequence *uint32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	// The procedure relevant to the account
	Code custom.CodeableReference[Procedure] `bson:"code" json:"code"`
	// Date of the procedure (when coded procedure)
	DateOfService *custom.DateTime `bson:"dateOfService,omitempty" json:"dateOfService,omitempty"`
	// How this procedure value should be used in charging the account
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Package Code specific for billing
	PackageCode []CodeableConcept `bson:"packageCode,omitempty" json:"packageCode,omitempty"`
	// Any devices that were associated with the procedure
	Device []custom.Reference[Device] `bson:"device,omitempty" json:"device,omitempty"`
}

type AccountRelatedAccount struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Relationship of the associated Account
	Relationship *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	// Reference to an associated Account
	Account custom.Reference[Account] `bson:"account" json:"account"`
}

type AccountBalance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Who is expected to pay this part of the balance
	Aggregate *CodeableConcept `bson:"aggregate,omitempty" json:"aggregate,omitempty"`
	// current | 30 | 60 | 90 | 120
	Term *CodeableConcept `bson:"term,omitempty" json:"term,omitempty"`
	// Estimated balance
	Estimate *bool `bson:"estimate,omitempty" json:"estimate,omitempty"`
	// Calculated amount
	Amount Money `bson:"amount" json:"amount"`
}

type Account struct {
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
	// Account number
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | inactive | entered-in-error | on-hold | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Tracks the lifecycle of the account through the billing process
	BillingStatus *CodeableConcept `bson:"billingStatus,omitempty" json:"billingStatus,omitempty"`
	// E.g. patient, expense, depreciation
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Human-readable label
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// The entity that caused the expenses
	Subject []custom.Reference[AccountSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Transaction window
	ServicePeriod *Period `bson:"servicePeriod,omitempty" json:"servicePeriod,omitempty"`
	// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account
	Coverage []AccountCoverage `bson:"coverage,omitempty" json:"coverage,omitempty"`
	// Entity managing the Account
	Owner *custom.Reference[Organization] `bson:"owner,omitempty" json:"owner,omitempty"`
	// Explanation of purpose/use
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The parties ultimately responsible for balancing the Account
	Guarantor []AccountGuarantor `bson:"guarantor,omitempty" json:"guarantor,omitempty"`
	// The list of diagnoses relevant to this account
	Diagnosis []AccountDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	// The list of procedures relevant to this account
	Procedure []AccountProcedure `bson:"procedure,omitempty" json:"procedure,omitempty"`
	// Other associated accounts related to this account
	RelatedAccount []AccountRelatedAccount `bson:"relatedAccount,omitempty" json:"relatedAccount,omitempty"`
	// The base or default currency
	Currency *CodeableConcept `bson:"currency,omitempty" json:"currency,omitempty"`
	// Calculated account balance(s)
	Balance []AccountBalance `bson:"balance,omitempty" json:"balance,omitempty"`
	// Time the balance amount was calculated
	CalculatedAt *custom.Instant `bson:"calculatedAt,omitempty" json:"calculatedAt,omitempty"`
}

type AccountSubject interface {
	gofhirclient.Resource

	Is_AccountSubject()
}

func (p Patient) Is_AccountSubject()           {}
func (d Device) Is_AccountSubject()            {}
func (p Practitioner) Is_AccountSubject()      {}
func (p PractitionerRole) Is_AccountSubject()  {}
func (l Location) Is_AccountSubject()          {}
func (h HealthcareService) Is_AccountSubject() {}
func (o Organization) Is_AccountSubject()      {}

type AccountGuarantorParty interface {
	gofhirclient.Resource

	Is_AccountGuarantorParty()
}

func (p Patient) Is_AccountGuarantorParty()       {}
func (r RelatedPerson) Is_AccountGuarantorParty() {}
func (o Organization) Is_AccountGuarantorParty()  {}

func (a Account) ResourceType() string {
	return "StructureDefinition"
}
