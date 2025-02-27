// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Invoice
type Invoice struct {
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
	// Business Identifier for item
	Identifier []Identifier `json:"identifier,omitempty"`
	// draft | issued | balanced | cancelled | entered-in-error
	Status custom.Code `json:"status"`
	// Reason for cancellation of this Invoice
	CancelledReason *string `json:"cancelledReason,omitempty"`
	// Type of Invoice
	Type *CodeableConcept `json:"type,omitempty"`
	// Recipient(s) of goods and services
	Subject *custom.Reference[InvoiceSubject] `json:"subject,omitempty"`
	// Recipient of this invoice
	Recipient *custom.Reference[InvoiceRecipient] `json:"recipient,omitempty"`
	// DEPRICATED
	Date *custom.DateTime `json:"date,omitempty"`
	// When posted
	Creation *custom.DateTime `json:"creation,omitempty"`
	// Billing date or period
	PeriodDate *custom.Date `json:"periodDate,omitempty"`
	// Billing date or period
	PeriodPeriod *Period `json:"periodPeriod,omitempty"`
	// Participant in creation of this Invoice
	Participant []InvoiceParticipant `json:"participant,omitempty"`
	// Issuing Organization of Invoice
	Issuer *custom.Reference[Organization] `json:"issuer,omitempty"`
	// Account that is being balanced
	Account *custom.Reference[Account] `json:"account,omitempty"`
	// Line items of this Invoice
	LineItem []InvoiceLineItem `json:"lineItem,omitempty"`
	// Components of Invoice total
	TotalPriceComponent []MonetaryComponent `json:"totalPriceComponent,omitempty"`
	// Net total of this Invoice
	TotalNet *Money `json:"totalNet,omitempty"`
	// Gross total of this Invoice
	TotalGross *Money `json:"totalGross,omitempty"`
	// Payment details
	PaymentTerms *custom.Markdown `json:"paymentTerms,omitempty"`
	// Comments made about the invoice
	Note []Annotation `json:"note,omitempty"`
}

type InvoiceParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of involvement in creation of this Invoice
	Role *CodeableConcept `json:"role,omitempty"`
	// Individual who was involved
	Actor custom.Reference[InvoiceParticipantActor] `json:"actor"`
}

type InvoiceLineItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Sequence number of line item
	Sequence *uint32 `json:"sequence,omitempty"`
	// Service data or period
	ServicedDate *custom.Date `json:"servicedDate,omitempty"`
	// Service data or period
	ServicedPeriod *Period `json:"servicedPeriod,omitempty"`
	// Reference to ChargeItem containing details of this line item or an inline billing code
	ChargeItemReference *custom.Reference[ChargeItem] `json:"chargeItemReference,omitempty"`
	// Reference to ChargeItem containing details of this line item or an inline billing code
	ChargeItemCodeableConcept *CodeableConcept `json:"chargeItemCodeableConcept,omitempty"`
	// Components of total line item price
	PriceComponent []MonetaryComponent `json:"priceComponent,omitempty"`
}

type OtherInvoice Invoice

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
	return "Invoice"
}

func (i Invoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInvoice
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherInvoice: OtherInvoice(i), ResourceType: i.ResourceType()})
}

func UnmarshalInvoice(b []byte) (res Invoice, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
