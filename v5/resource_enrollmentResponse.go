// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
type EnrollmentResponse struct {
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
	// Business Identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Claim reference
	Request *custom.Reference[EnrollmentRequest] `bson:"request,omitempty" json:"request,omitempty"`
	// queued | complete | error | partial
	Outcome *custom.Code `bson:"outcome,omitempty" json:"outcome,omitempty"`
	// Disposition Message
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty"`
	// Creation date
	Created *custom.DateTime `bson:"created,omitempty" json:"created,omitempty"`
	// Insurer
	Organization *custom.Reference[Organization] `bson:"organization,omitempty" json:"organization,omitempty"`
	// Responsible practitioner
	RequestProvider *custom.Reference[EnrollmentResponseRequestProvider] `bson:"requestProvider,omitempty" json:"requestProvider,omitempty"`
}

type EnrollmentResponseRequestProvider interface {
	gofhirclient.Resource

	Is_EnrollmentResponseRequestProvider()
}

func (p Practitioner) Is_EnrollmentResponseRequestProvider()     {}
func (p PractitionerRole) Is_EnrollmentResponseRequestProvider() {}
func (o Organization) Is_EnrollmentResponseRequestProvider()     {}

func (e EnrollmentResponse) ResourceType() string {
	return "StructureDefinition"
}
