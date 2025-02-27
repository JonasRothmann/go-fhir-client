// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
type PaymentReconciliation struct {
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
	// Business Identifier for a payment reconciliation
	Identifier []Identifier `json:"identifier,omitempty"`
	// Category of payment
	Type CodeableConcept `json:"type"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `json:"status"`
	// Workflow originating payment
	Kind *CodeableConcept `json:"kind,omitempty"`
	// Period covered
	Period *Period `json:"period,omitempty"`
	// Creation date
	Created custom.DateTime `json:"created"`
	// Who entered the payment
	Enterer *custom.Reference[PaymentReconciliationEnterer] `json:"enterer,omitempty"`
	// Nature of the source
	IssuerType *CodeableConcept `json:"issuerType,omitempty"`
	// Party generating payment
	PaymentIssuer *custom.Reference[PaymentReconciliationPaymentIssuer] `json:"paymentIssuer,omitempty"`
	// Reference to requesting resource
	Request *custom.Reference[Task] `json:"request,omitempty"`
	// Responsible practitioner
	Requestor *custom.Reference[PaymentReconciliationRequestor] `json:"requestor,omitempty"`
	// queued | complete | error | partial
	Outcome *custom.Code `json:"outcome,omitempty"`
	// Disposition message
	Disposition *string `json:"disposition,omitempty"`
	// When payment issued
	Date custom.Date `json:"date"`
	// Where payment collected
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Payment instrument
	Method *CodeableConcept `json:"method,omitempty"`
	// Type of card
	CardBrand *string `json:"cardBrand,omitempty"`
	// Digits for verification
	AccountNumber *string `json:"accountNumber,omitempty"`
	// Expiration year-month
	ExpirationDate *custom.Date `json:"expirationDate,omitempty"`
	// Processor name
	Processor *string `json:"processor,omitempty"`
	// Check number or payment reference
	ReferenceNumber *string `json:"referenceNumber,omitempty"`
	// Authorization number
	Authorization *string `json:"authorization,omitempty"`
	// Amount offered by the issuer
	TenderedAmount *Money `json:"tenderedAmount,omitempty"`
	// Amount returned by the receiver
	ReturnedAmount *Money `json:"returnedAmount,omitempty"`
	// Total amount of Payment
	Amount Money `json:"amount"`
	// Business identifier for the payment
	PaymentIdentifier *Identifier `json:"paymentIdentifier,omitempty"`
	// Settlement particulars
	Allocation []PaymentReconciliationAllocation `json:"allocation,omitempty"`
	// Printed form identifier
	FormCode *CodeableConcept `json:"formCode,omitempty"`
	// Note concerning processing
	ProcessNote []PaymentReconciliationProcessNote `json:"processNote,omitempty"`
}

type PaymentReconciliationAllocation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Business identifier of the payment detail
	Identifier *Identifier `json:"identifier,omitempty"`
	// Business identifier of the prior payment detail
	Predecessor *Identifier `json:"predecessor,omitempty"`
	// Subject of the payment
	Target *custom.Reference[PaymentReconciliationAllocationTarget] `json:"target,omitempty"`
	// Sub-element of the subject
	TargetItemString *string `json:"targetItemString,omitempty"`
	// Sub-element of the subject
	TargetItemIdentifier *Identifier `json:"targetItemIdentifier,omitempty"`
	// Sub-element of the subject
	TargetItemPositiveInt *uint32 `json:"targetItemPositiveInt,omitempty"`
	// Applied-to encounter
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Applied-to account
	Account *custom.Reference[Account] `json:"account,omitempty"`
	// Category of payment
	Type *CodeableConcept `json:"type,omitempty"`
	// Submitter of the request
	Submitter *custom.Reference[PaymentReconciliationAllocationSubmitter] `json:"submitter,omitempty"`
	// Response committing to a payment
	Response *custom.Reference[ClaimResponse] `json:"response,omitempty"`
	// Date of commitment to pay
	Date *custom.Date `json:"date,omitempty"`
	// Contact for the response
	Responsible *custom.Reference[PractitionerRole] `json:"responsible,omitempty"`
	// Recipient of the payment
	Payee *custom.Reference[PaymentReconciliationAllocationPayee] `json:"payee,omitempty"`
	// Amount allocated to this payable
	Amount *Money `json:"amount,omitempty"`
}

type PaymentReconciliationProcessNote struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// display | print | printoper
	Type *custom.Code `json:"type,omitempty"`
	// Note explanatory text
	Text *string `json:"text,omitempty"`
}

type OtherPaymentReconciliation PaymentReconciliation

type PaymentReconciliationEnterer interface {
	gofhirclient.Resource

	Is_PaymentReconciliationEnterer()
}

func (p Practitioner) Is_PaymentReconciliationEnterer()     {}
func (p PractitionerRole) Is_PaymentReconciliationEnterer() {}
func (o Organization) Is_PaymentReconciliationEnterer()     {}

type PaymentReconciliationPaymentIssuer interface {
	gofhirclient.Resource

	Is_PaymentReconciliationPaymentIssuer()
}

func (o Organization) Is_PaymentReconciliationPaymentIssuer()  {}
func (p Patient) Is_PaymentReconciliationPaymentIssuer()       {}
func (r RelatedPerson) Is_PaymentReconciliationPaymentIssuer() {}

type PaymentReconciliationRequestor interface {
	gofhirclient.Resource

	Is_PaymentReconciliationRequestor()
}

func (p Practitioner) Is_PaymentReconciliationRequestor()     {}
func (p PractitionerRole) Is_PaymentReconciliationRequestor() {}
func (o Organization) Is_PaymentReconciliationRequestor()     {}

type PaymentReconciliationAllocationTarget interface {
	gofhirclient.Resource

	Is_PaymentReconciliationAllocationTarget()
}

func (c Claim) Is_PaymentReconciliationAllocationTarget()      {}
func (a Account) Is_PaymentReconciliationAllocationTarget()    {}
func (i Invoice) Is_PaymentReconciliationAllocationTarget()    {}
func (c ChargeItem) Is_PaymentReconciliationAllocationTarget() {}
func (e Encounter) Is_PaymentReconciliationAllocationTarget()  {}
func (c Contract) Is_PaymentReconciliationAllocationTarget()   {}

type PaymentReconciliationAllocationSubmitter interface {
	gofhirclient.Resource

	Is_PaymentReconciliationAllocationSubmitter()
}

func (p Practitioner) Is_PaymentReconciliationAllocationSubmitter()     {}
func (p PractitionerRole) Is_PaymentReconciliationAllocationSubmitter() {}
func (o Organization) Is_PaymentReconciliationAllocationSubmitter()     {}

type PaymentReconciliationAllocationPayee interface {
	gofhirclient.Resource

	Is_PaymentReconciliationAllocationPayee()
}

func (p Practitioner) Is_PaymentReconciliationAllocationPayee()     {}
func (p PractitionerRole) Is_PaymentReconciliationAllocationPayee() {}
func (o Organization) Is_PaymentReconciliationAllocationPayee()     {}

func (p PaymentReconciliation) ResourceType() string {
	return "PaymentReconciliation"
}

func (p PaymentReconciliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentReconciliation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPaymentReconciliation: OtherPaymentReconciliation(p), ResourceType: p.ResourceType()})
}

func UnmarshalPaymentReconciliation(b []byte) (res PaymentReconciliation, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
