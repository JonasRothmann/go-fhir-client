// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
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
	// Business Identifier for the payment notice
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `json:"status"`
	// Request reference
	Request *custom.Reference[Resource] `json:"request,omitempty"`
	// Response reference
	Response *custom.Reference[Resource] `json:"response,omitempty"`
	// Creation date
	Created custom.DateTime `json:"created"`
	// Responsible practitioner
	Reporter *custom.Reference[PaymentNoticeReporter] `json:"reporter,omitempty"`
	// Payment reference
	Payment *custom.Reference[PaymentReconciliation] `json:"payment,omitempty"`
	// Payment or clearing date
	PaymentDate *custom.Date `json:"paymentDate,omitempty"`
	// Party being paid
	Payee *custom.Reference[PaymentNoticePayee] `json:"payee,omitempty"`
	// Party being notified
	Recipient custom.Reference[Organization] `json:"recipient"`
	// Monetary amount of the payment
	Amount Money `json:"amount"`
	// Issued or cleared Status of the payment
	PaymentStatus *CodeableConcept `json:"paymentStatus,omitempty"`
}

type OtherPaymentNotice PaymentNotice

type PaymentNoticeReporter interface {
	gofhirclient.Resource

	Is_PaymentNoticeReporter()
}

func (p Practitioner) Is_PaymentNoticeReporter()     {}
func (p PractitionerRole) Is_PaymentNoticeReporter() {}
func (o Organization) Is_PaymentNoticeReporter()     {}

type PaymentNoticePayee interface {
	gofhirclient.Resource

	Is_PaymentNoticePayee()
}

func (p Practitioner) Is_PaymentNoticePayee()     {}
func (p PractitionerRole) Is_PaymentNoticePayee() {}
func (o Organization) Is_PaymentNoticePayee()     {}

func (p PaymentNotice) ResourceType() string {
	return "PaymentNotice"
}

func (p PaymentNotice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentNotice
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPaymentNotice: OtherPaymentNotice(p), ResourceType: p.ResourceType()})
}

func UnmarshalPaymentNotice(b []byte) (res PaymentNotice, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
