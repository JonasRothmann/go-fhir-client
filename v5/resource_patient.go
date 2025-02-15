// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Patient
type PatientLink struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The other patient or related person resource that the link refers to
	Other custom.Reference[PatientLinkOther] `bson:"other" json:"other"`
	// replaced-by | replaces | refer | seealso
	Type custom.Code `bson:"type" json:"type"`
}

type Patient struct {
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
	// An identifier for this patient
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether this patient's record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// A name associated with the patient
	Name []HumanName `bson:"name,omitempty" json:"name,omitempty"`
	// A contact detail for the individual
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `bson:"gender,omitempty" json:"gender,omitempty"`
	// The date of birth for the individual
	BirthDate *custom.Date `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	// Indicates if the individual is deceased or not
	Deceased *bool `bson:"deceased,omitempty" json:"deceased,omitempty"`
	// An address for the individual
	Address []Address `bson:"address,omitempty" json:"address,omitempty"`
	// Marital (civil) status of a patient
	MaritalStatus *CodeableConcept `bson:"maritalStatus,omitempty" json:"maritalStatus,omitempty"`
	// Whether patient is part of a multiple birth
	MultipleBirth *bool `bson:"multipleBirth,omitempty" json:"multipleBirth,omitempty"`
	// Image of the patient
	Photo []Attachment `bson:"photo,omitempty" json:"photo,omitempty"`
	// A contact party (e.g. guardian, partner, friend) for the patient
	Contact []PatientContact `bson:"contact,omitempty" json:"contact,omitempty"`
	// A language which may be used to communicate with the patient about his or her health
	Communication []PatientCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
	// Patient's nominated primary care provider
	GeneralPractitioner []custom.Reference[PatientGeneralPractitioner] `bson:"generalPractitioner,omitempty" json:"generalPractitioner,omitempty"`
	// Organization that is the custodian of the patient record
	ManagingOrganization *custom.Reference[Organization] `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	// Link to a Patient or RelatedPerson resource that concerns the same actual individual
	Link []PatientLink `bson:"link,omitempty" json:"link,omitempty"`
}

type PatientContact struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The kind of relationship
	Relationship []CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	// A name associated with the contact person
	Name *HumanName `bson:"name,omitempty" json:"name,omitempty"`
	// A contact detail for the person
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
	// Address for the contact person
	Address *Address `bson:"address,omitempty" json:"address,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `bson:"gender,omitempty" json:"gender,omitempty"`
	// Organization that is associated with the contact
	Organization *custom.Reference[Organization] `bson:"organization,omitempty" json:"organization,omitempty"`
	// The period during which this contact person or organization is valid to be contacted relating to this patient
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type PatientCommunication struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The language which can be used to communicate with the patient about his or her health
	Language CodeableConcept `bson:"language" json:"language"`
	// Language preference indicator
	Preferred *bool `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

type PatientGeneralPractitioner interface {
	gofhirclient.Resource

	Is_PatientGeneralPractitioner()
}

func (o Organization) Is_PatientGeneralPractitioner()     {}
func (p Practitioner) Is_PatientGeneralPractitioner()     {}
func (p PractitionerRole) Is_PatientGeneralPractitioner() {}

type PatientLinkOther interface {
	gofhirclient.Resource

	Is_PatientLinkOther()
}

func (p Patient) Is_PatientLinkOther()       {}
func (r RelatedPerson) Is_PatientLinkOther() {}

func (p Patient) ResourceType() string {
	return "StructureDefinition"
}
