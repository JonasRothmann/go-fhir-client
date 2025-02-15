// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Basic
type Basic struct {
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
	// Business identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Kind of Resource
	Code CodeableConcept `bson:"code" json:"code"`
	// Identifies the focus of this resource
	Subject *custom.Reference[Resource] `bson:"subject,omitempty" json:"subject,omitempty"`
	// When created
	Created *custom.DateTime `bson:"created,omitempty" json:"created,omitempty"`
	// Who created
	Author *custom.Reference[BasicAuthor] `bson:"author,omitempty" json:"author,omitempty"`
}

type BasicAuthor interface {
	gofhirclient.Resource

	Is_BasicAuthor()
}

func (p Practitioner) Is_BasicAuthor()     {}
func (p PractitionerRole) Is_BasicAuthor() {}
func (p Patient) Is_BasicAuthor()          {}
func (r RelatedPerson) Is_BasicAuthor()    {}
func (o Organization) Is_BasicAuthor()     {}
func (d Device) Is_BasicAuthor()           {}
func (c CareTeam) Is_BasicAuthor()         {}

func (b Basic) ResourceType() string {
	return "StructureDefinition"
}
