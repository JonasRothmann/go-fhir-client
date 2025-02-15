// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
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
	// Business Identifier for the payment notice
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Request reference
	Request *custom.Reference[Resource] `bson:"request,omitempty" json:"request,omitempty"`
	// Response reference
	Response *custom.Reference[Resource] `bson:"response,omitempty" json:"response,omitempty"`
	// Creation date
	Created custom.DateTime `bson:"created" json:"created"`
	// Responsible practitioner
	Reporter *custom.Reference[PaymentNoticeReporter] `bson:"reporter,omitempty" json:"reporter,omitempty"`
	// Payment reference
	Payment *custom.Reference[PaymentReconciliation] `bson:"payment,omitempty" json:"payment,omitempty"`
	// Payment or clearing date
	PaymentDate *custom.Date `bson:"paymentDate,omitempty" json:"paymentDate,omitempty"`
	// Party being paid
	Payee *custom.Reference[PaymentNoticePayee] `bson:"payee,omitempty" json:"payee,omitempty"`
	// Party being notified
	Recipient custom.Reference[Organization] `bson:"recipient" json:"recipient"`
	// Monetary amount of the payment
	Amount Money `bson:"amount" json:"amount"`
	// Issued or cleared Status of the payment
	PaymentStatus *CodeableConcept `bson:"paymentStatus,omitempty" json:"paymentStatus,omitempty"`
}

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
	return "StructureDefinition"
}
