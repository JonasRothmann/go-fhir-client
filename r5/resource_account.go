// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Account
type AccountRelatedAccount struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Relationship of the associated Account
	Relationship *CodeableConcept `json:"relationship,omitempty"`
	// Reference to an associated Account
	Account custom.Reference[Account] `json:"account"`
}

type AccountBalance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Who is expected to pay this part of the balance
	Aggregate *CodeableConcept `json:"aggregate,omitempty"`
	// current | 30 | 60 | 90 | 120
	Term *CodeableConcept `json:"term,omitempty"`
	// Estimated balance
	Estimate *bool `json:"estimate,omitempty"`
	// Calculated amount
	Amount Money `json:"amount"`
}

type Account struct {
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
	// Account number
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | inactive | entered-in-error | on-hold | unknown
	Status custom.Code `json:"status"`
	// Tracks the lifecycle of the account through the billing process
	BillingStatus *CodeableConcept `json:"billingStatus,omitempty"`
	// E.g. patient, expense, depreciation
	Type *CodeableConcept `json:"type,omitempty"`
	// Human-readable label
	Name *string `json:"name,omitempty"`
	// The entity that caused the expenses
	Subject []custom.Reference[AccountSubject] `json:"subject,omitempty"`
	// Transaction window
	ServicePeriod *Period `json:"servicePeriod,omitempty"`
	// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account
	Coverage []AccountCoverage `json:"coverage,omitempty"`
	// Entity managing the Account
	Owner *custom.Reference[Organization] `json:"owner,omitempty"`
	// Explanation of purpose/use
	Description *custom.Markdown `json:"description,omitempty"`
	// The parties ultimately responsible for balancing the Account
	Guarantor []AccountGuarantor `json:"guarantor,omitempty"`
	// The list of diagnoses relevant to this account
	Diagnosis []AccountDiagnosis `json:"diagnosis,omitempty"`
	// The list of procedures relevant to this account
	Procedure []AccountProcedure `json:"procedure,omitempty"`
	// Other associated accounts related to this account
	RelatedAccount []AccountRelatedAccount `json:"relatedAccount,omitempty"`
	// The base or default currency
	Currency *CodeableConcept `json:"currency,omitempty"`
	// Calculated account balance(s)
	Balance []AccountBalance `json:"balance,omitempty"`
	// Time the balance amount was calculated
	CalculatedAt *custom.Instant `json:"calculatedAt,omitempty"`
}

type AccountCoverage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The party(s), such as insurances, that may contribute to the payment of this account
	Coverage custom.Reference[Coverage] `json:"coverage"`
	// The priority of the coverage in the context of this account
	Priority *uint32 `json:"priority,omitempty"`
}

type AccountGuarantor struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Responsible entity
	Party custom.Reference[AccountGuarantorParty] `json:"party"`
	// Credit or other hold applied
	OnHold *bool `json:"onHold,omitempty"`
	// Guarantee account during
	Period *Period `json:"period,omitempty"`
}

type AccountDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Ranking of the diagnosis (for each type)
	Sequence *uint32 `json:"sequence,omitempty"`
	// The diagnosis relevant to the account
	Condition custom.CodeableReference[Condition] `json:"condition"`
	// Date of the diagnosis (when coded diagnosis)
	DateOfDiagnosis *custom.DateTime `json:"dateOfDiagnosis,omitempty"`
	// Type that this diagnosis has relevant to the account (e.g. admission, billing, discharge …)
	Type []CodeableConcept `json:"type,omitempty"`
	// Diagnosis present on Admission
	OnAdmission *bool `json:"onAdmission,omitempty"`
	// Package Code specific for billing
	PackageCode []CodeableConcept `json:"packageCode,omitempty"`
}

type AccountProcedure struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Ranking of the procedure (for each type)
	Sequence *uint32 `json:"sequence,omitempty"`
	// The procedure relevant to the account
	Code custom.CodeableReference[Procedure] `json:"code"`
	// Date of the procedure (when coded procedure)
	DateOfService *custom.DateTime `json:"dateOfService,omitempty"`
	// How this procedure value should be used in charging the account
	Type []CodeableConcept `json:"type,omitempty"`
	// Package Code specific for billing
	PackageCode []CodeableConcept `json:"packageCode,omitempty"`
	// Any devices that were associated with the procedure
	Device []custom.Reference[Device] `json:"device,omitempty"`
}

type OtherAccount Account

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
	return "Account"
}

func (a Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAccount
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAccount: OtherAccount(a), ResourceType: a.ResourceType()})
}

func UnmarshalAccount(b []byte) (res Account, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
