// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
type EnrollmentRequest struct {
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
	// Creation date
	Created *custom.DateTime `bson:"created,omitempty" json:"created,omitempty"`
	// Target
	Insurer *custom.Reference[Organization] `bson:"insurer,omitempty" json:"insurer,omitempty"`
	// Responsible practitioner
	Provider *custom.Reference[EnrollmentRequestProvider] `bson:"provider,omitempty" json:"provider,omitempty"`
	// The subject to be enrolled
	Candidate *custom.Reference[Patient] `bson:"candidate,omitempty" json:"candidate,omitempty"`
	// Insurance information
	Coverage *custom.Reference[Coverage] `bson:"coverage,omitempty" json:"coverage,omitempty"`
}

type EnrollmentRequestProvider interface {
	gofhirclient.Resource

	Is_EnrollmentRequestProvider()
}

func (p Practitioner) Is_EnrollmentRequestProvider()     {}
func (p PractitionerRole) Is_EnrollmentRequestProvider() {}
func (o Organization) Is_EnrollmentRequestProvider()     {}

func (e EnrollmentRequest) ResourceType() string {
	return "StructureDefinition"
}
