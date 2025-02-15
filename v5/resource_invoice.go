// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Invoice
type Invoice struct {
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
	// Business Identifier for item
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// draft | issued | balanced | cancelled | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Reason for cancellation of this Invoice
	CancelledReason *string `bson:"cancelledReason,omitempty" json:"cancelledReason,omitempty"`
	// Type of Invoice
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Recipient(s) of goods and services
	Subject *custom.Reference[InvoiceSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Recipient of this invoice
	Recipient *custom.Reference[InvoiceRecipient] `bson:"recipient,omitempty" json:"recipient,omitempty"`
	// DEPRICATED
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// When posted
	Creation *custom.DateTime `bson:"creation,omitempty" json:"creation,omitempty"`
	// Billing date or period
	Period *custom.Date `bson:"period,omitempty" json:"period,omitempty"`
	// Participant in creation of this Invoice
	Participant []InvoiceParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// Issuing Organization of Invoice
	Issuer *custom.Reference[Organization] `bson:"issuer,omitempty" json:"issuer,omitempty"`
	// Account that is being balanced
	Account *custom.Reference[Account] `bson:"account,omitempty" json:"account,omitempty"`
	// Line items of this Invoice
	LineItem []InvoiceLineItem `bson:"lineItem,omitempty" json:"lineItem,omitempty"`
	// Components of Invoice total
	TotalPriceComponent []MonetaryComponent `bson:"totalPriceComponent,omitempty" json:"totalPriceComponent,omitempty"`
	// Net total of this Invoice
	TotalNet *Money `bson:"totalNet,omitempty" json:"totalNet,omitempty"`
	// Gross total of this Invoice
	TotalGross *Money `bson:"totalGross,omitempty" json:"totalGross,omitempty"`
	// Payment details
	PaymentTerms *custom.Markdown `bson:"paymentTerms,omitempty" json:"paymentTerms,omitempty"`
	// Comments made about the invoice
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type InvoiceParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of involvement in creation of this Invoice
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Individual who was involved
	Actor custom.Reference[InvoiceParticipantActor] `bson:"actor" json:"actor"`
}

type InvoiceLineItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Sequence number of line item
	Sequence *uint32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	// Service data or period
	Serviced *custom.Date `bson:"serviced,omitempty" json:"serviced,omitempty"`
	// Reference to ChargeItem containing details of this line item or an inline billing code
	ChargeItem custom.Reference[ChargeItem] `bson:"chargeItem" json:"chargeItem"`
	// Components of total line item price
	PriceComponent []MonetaryComponent `bson:"priceComponent,omitempty" json:"priceComponent,omitempty"`
}

type InvoiceSubject interface {
	gofhirclient.Resource

	Is_InvoiceSubject()
}

func (p Patient) Is_InvoiceSubject() {}
func (g Group) Is_InvoiceSubject()   {}

type InvoiceRecipient interface {
	gofhirclient.Resource

	Is_InvoiceRecipient()
}

func (o Organization) Is_InvoiceRecipient()  {}
func (p Patient) Is_InvoiceRecipient()       {}
func (r RelatedPerson) Is_InvoiceRecipient() {}

type InvoiceParticipantActor interface {
	gofhirclient.Resource

	Is_InvoiceParticipantActor()
}

func (p Practitioner) Is_InvoiceParticipantActor()     {}
func (o Organization) Is_InvoiceParticipantActor()     {}
func (p Patient) Is_InvoiceParticipantActor()          {}
func (p PractitionerRole) Is_InvoiceParticipantActor() {}
func (d Device) Is_InvoiceParticipantActor()           {}
func (r RelatedPerson) Is_InvoiceParticipantActor()    {}

func (i Invoice) ResourceType() string {
	return "StructureDefinition"
}
