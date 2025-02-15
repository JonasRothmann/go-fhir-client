// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/RelatedPerson
type RelatedPerson struct {
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
	// Whether this related person's record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// The patient this person is related to
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// The relationship of the related person to the patient
	Relationship []CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	// A name associated with the person
	Name []HumanName `bson:"name,omitempty" json:"name,omitempty"`
	// A contact detail for the person
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `bson:"gender,omitempty" json:"gender,omitempty"`
	// The date on which the related person was born
	BirthDate *custom.Date `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	// Address where the related person can be contacted or visited
	Address []Address `bson:"address,omitempty" json:"address,omitempty"`
	// Image of the person
	Photo []Attachment `bson:"photo,omitempty" json:"photo,omitempty"`
	// Period of time that this relationship is considered valid
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// A language which may be used to communicate with the related person about the patient's health
	Communication []RelatedPersonCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
}

type RelatedPersonCommunication struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The language which can be used to communicate with the related person about the patient's health
	Language CodeableConcept `bson:"language" json:"language"`
	// Language preference indicator
	Preferred *bool `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

func (r RelatedPerson) ResourceType() string {
	return "StructureDefinition"
}
