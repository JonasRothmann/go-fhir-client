// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
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
	// active | inactive | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Clinical, administrative, etc
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Coded or textual message to display to user
	Code CodeableConcept `bson:"code" json:"code"`
	// Who/What is flag about?
	Subject custom.Reference[FlagSubject] `bson:"subject" json:"subject"`
	// Time period when flag is active
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Alert relevant during encounter
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Flag creator
	Author *custom.Reference[FlagAuthor] `bson:"author,omitempty" json:"author,omitempty"`
}

type FlagSubject interface {
	gofhirclient.Resource

	Is_FlagSubject()
}

func (p Patient) Is_FlagSubject()          {}
func (r RelatedPerson) Is_FlagSubject()    {}
func (l Location) Is_FlagSubject()         {}
func (g Group) Is_FlagSubject()            {}
func (o Organization) Is_FlagSubject()     {}
func (p Practitioner) Is_FlagSubject()     {}
func (p PractitionerRole) Is_FlagSubject() {}
func (p PlanDefinition) Is_FlagSubject()   {}
func (m Medication) Is_FlagSubject()       {}
func (p Procedure) Is_FlagSubject()        {}

type FlagAuthor interface {
	gofhirclient.Resource

	Is_FlagAuthor()
}

func (d Device) Is_FlagAuthor()           {}
func (o Organization) Is_FlagAuthor()     {}
func (p Patient) Is_FlagAuthor()          {}
func (r RelatedPerson) Is_FlagAuthor()    {}
func (p Practitioner) Is_FlagAuthor()     {}
func (p PractitionerRole) Is_FlagAuthor() {}

func (f Flag) ResourceType() string {
	return "StructureDefinition"
}
