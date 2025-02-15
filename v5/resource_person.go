// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Person
type Person struct {
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
	// A human identifier for this person
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// This person's record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// A name associated with the person
	Name []HumanName `bson:"name,omitempty" json:"name,omitempty"`
	// A contact detail for the person
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `bson:"gender,omitempty" json:"gender,omitempty"`
	// The date on which the person was born
	BirthDate *custom.Date `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	// Indicates if the individual is deceased or not
	Deceased *bool `bson:"deceased,omitempty" json:"deceased,omitempty"`
	// One or more addresses for the person
	Address []Address `bson:"address,omitempty" json:"address,omitempty"`
	// Marital (civil) status of a person
	MaritalStatus *CodeableConcept `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	// Image of the person
	Photo []Attachment `bson:"photo,omitempty" json:"photo,omitempty"`
	// A language which may be used to communicate with the person about his or her health
	Communication []PersonCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
	// The organization that is the custodian of the person record
	ManagingOrganization *custom.Reference[Organization] `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	// Link to a resource that concerns the same actual person
	Link []PersonLink `bson:"link,omitempty" json:"link,omitempty"`
}

type PersonCommunication struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The language which can be used to communicate with the person about his or her health
	Language CodeableConcept `bson:"language" json:"language"`
	// Language preference indicator
	Preferred *bool `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

type PersonLink struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The resource to which this actual person is associated
	Target custom.Reference[PersonLinkTarget] `bson:"target" json:"target"`
	// level1 | level2 | level3 | level4
	Assurance *custom.Code `bson:"assurance,omitempty" json:"assurance,omitempty"`
}

type PersonLinkTarget interface {
	gofhirclient.Resource

	Is_PersonLinkTarget()
}

func (p Patient) Is_PersonLinkTarget()       {}
func (p Practitioner) Is_PersonLinkTarget()  {}
func (r RelatedPerson) Is_PersonLinkTarget() {}
func (p Person) Is_PersonLinkTarget()        {}

func (p Person) ResourceType() string {
	return "StructureDefinition"
}
