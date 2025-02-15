// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/PaymentReconciliation
type PaymentReconciliation struct {
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
	// Business Identifier for a payment reconciliation
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Category of payment
	Type CodeableConcept `bson:"type" json:"type"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Workflow originating payment
	Kind *CodeableConcept `bson:"kind,omitempty" json:"kind,omitempty"`
	// Period covered
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Creation date
	Created custom.DateTime `bson:"created" json:"created"`
	// Who entered the payment
	Enterer *custom.Reference[PaymentReconciliationEnterer] `bson:"enterer,omitempty" json:"enterer,omitempty"`
	// Nature of the source
	IssuerType *CodeableConcept `bson:"issuerType,omitempty" json:"issuerType,omitempty"`
	// Party generating payment
	PaymentIssuer *custom.Reference[PaymentReconciliationPaymentIssuer] `bson:"paymentIssuer,omitempty" json:"paymentIssuer,omitempty"`
	// Reference to requesting resource
	Request *custom.Reference[Task] `bson:"request,omitempty" json:"request,omitempty"`
	// Responsible practitioner
	Requestor *custom.Reference[PaymentReconciliationRequestor] `bson:"requestor,omitempty" json:"requestor,omitempty"`
	// queued | complete | error | partial
	Outcome *custom.Code `bson:"outcome,omitempty" json:"outcome,omitempty"`
	// Disposition message
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty"`
	// When payment issued
	Date custom.Date `bson:"date" json:"date"`
	// Where payment collected
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Payment instrument
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Type of card
	CardBrand *string `bson:"cardBrand,omitempty" json:"cardBrand,omitempty"`
	// Digits for verification
	AccountNumber *string `bson:"accountNumber,omitempty" json:"accountNumber,omitempty"`
	// Expiration year-month
	ExpirationDate *custom.Date `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	// Processor name
	Processor *string `bson:"processor,omitempty" json:"processor,omitempty"`
	// Check number or payment reference
	ReferenceNumber *string `bson:"referenceNumber,omitempty" json:"referenceNumber,omitempty"`
	// Authorization number
	Authorization *string `bson:"authorization,omitempty" json:"authorization,omitempty"`
	// Amount offered by the issuer
	TenderedAmount *Money `bson:"tenderedAmount,omitempty" json:"tenderedAmount,omitempty"`
	// Amount returned by the receiver
	ReturnedAmount *Money `bson:"returnedAmount,omitempty" json:"returnedAmount,omitempty"`
	// Total amount of Payment
	Amount Money `bson:"amount" json:"amount"`
	// Business identifier for the payment
	PaymentIdentifier *Identifier `bson:"paymentIdentifier,omitempty" json:"paymentIdentifier,omitempty"`
	// Settlement particulars
	Allocation []PaymentReconciliationAllocation `bson:"allocation,omitempty" json:"allocation,omitempty"`
	// Printed form identifier
	FormCode *CodeableConcept `bson:"formCode,omitempty" json:"formCode,omitempty"`
	// Note concerning processing
	ProcessNote []PaymentReconciliationProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
}

type PaymentReconciliationAllocation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Business identifier of the payment detail
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business identifier of the prior payment detail
	Predecessor *Identifier `bson:"predecessor,omitempty" json:"predecessor,omitempty"`
	// Subject of the payment
	Target *custom.Reference[PaymentReconciliationAllocationTarget] `bson:"target,omitempty" json:"target,omitempty"`
	// Sub-element of the subject
	TargetItem *string `bson:"targetItem,omitempty" json:"targetItem,omitempty"`
	// Applied-to encounter
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Applied-to account
	Account *custom.Reference[Account] `bson:"account,omitempty" json:"account,omitempty"`
	// Category of payment
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Submitter of the request
	Submitter *custom.Reference[PaymentReconciliationAllocationSubmitter] `bson:"submitter,omitempty" json:"submitter,omitempty"`
	// Response committing to a payment
	Response *custom.Reference[ClaimResponse] `bson:"response,omitempty" json:"response,omitempty"`
	// Date of commitment to pay
	Date *custom.Date `bson:"date,omitempty" json:"date,omitempty"`
	// Contact for the response
	Responsible *custom.Reference[PractitionerRole] `bson:"responsible,omitempty" json:"responsible,omitempty"`
	// Recipient of the payment
	Payee *custom.Reference[PaymentReconciliationAllocationPayee] `bson:"payee,omitempty" json:"payee,omitempty"`
	// Amount allocated to this payable
	Amount *Money `bson:"amount,omitempty" json:"amount,omitempty"`
}

type PaymentReconciliationProcessNote struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// display | print | printoper
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// Note explanatory text
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
}

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
	return "StructureDefinition"
}
