// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Practitioner
type Practitioner struct {
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
	// An identifier for the person as this agent
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether this practitioner's record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// The name(s) associated with the practitioner
	Name []HumanName `bson:"name,omitempty" json:"name,omitempty"`
	// A contact detail for the practitioner (that apply to all roles)
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `bson:"gender,omitempty" json:"gender,omitempty"`
	// The date  on which the practitioner was born
	BirthDate *custom.Date `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	// Indicates if the practitioner is deceased or not
	Deceased *bool `bson:"deceased,omitempty" json:"deceased,omitempty"`
	// Address(es) of the practitioner that are not role specific (typically home address)
	Address []Address `bson:"address,omitempty" json:"address,omitempty"`
	// Image of the person
	Photo []Attachment `bson:"photo,omitempty" json:"photo,omitempty"`
	// Qualifications, certifications, accreditations, licenses, training, etc. pertaining to the provision of care
	Qualification []PractitionerQualification `bson:"qualification,omitempty" json:"qualification,omitempty"`
	// A language which may be used to communicate with the practitioner
	Communication []PractitionerCommunication `bson:"communication,omitempty" json:"communication,omitempty"`
}

type PractitionerQualification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// An identifier for this qualification for the practitioner
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Coded representation of the qualification
	Code CodeableConcept `bson:"code" json:"code"`
	// Period during which the qualification is valid
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Organization that regulates and issues the qualification
	Issuer *custom.Reference[Organization] `bson:"issuer,omitempty" json:"issuer,omitempty"`
}

type PractitionerCommunication struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The language code used to communicate with the practitioner
	Language CodeableConcept `bson:"language" json:"language"`
	// Language preference indicator
	Preferred *bool `bson:"preferred,omitempty" json:"preferred,omitempty"`
}

func (p Practitioner) ResourceType() string {
	return "StructureDefinition"
}
