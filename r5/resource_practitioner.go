// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Practitioner
type Practitioner struct {
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
	// An identifier for the person as this agent
	Identifier []Identifier `json:"identifier,omitempty"`
	// Whether this practitioner's record is in active use
	Active *bool `json:"active,omitempty"`
	// The name(s) associated with the practitioner
	Name []HumanName `json:"name,omitempty"`
	// A contact detail for the practitioner (that apply to all roles)
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `json:"gender,omitempty"`
	// The date  on which the practitioner was born
	BirthDate *custom.Date `json:"birthDate,omitempty"`
	// Indicates if the practitioner is deceased or not
	DeceasedBoolean *bool `json:"deceasedBoolean,omitempty"`
	// Indicates if the practitioner is deceased or not
	DeceasedDateTime *custom.DateTime `json:"deceasedDateTime,omitempty"`
	// Address(es) of the practitioner that are not role specific (typically home address)
	Address []Address `json:"address,omitempty"`
	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`
	// Qualifications, certifications, accreditations, licenses, training, etc. pertaining to the provision of care
	Qualification []PractitionerQualification `json:"qualification,omitempty"`
	// A language which may be used to communicate with the practitioner
	Communication []PractitionerCommunication `json:"communication,omitempty"`
}

type PractitionerQualification struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An identifier for this qualification for the practitioner
	Identifier []Identifier `json:"identifier,omitempty"`
	// Coded representation of the qualification
	Code CodeableConcept `json:"code"`
	// Period during which the qualification is valid
	Period *Period `json:"period,omitempty"`
	// Organization that regulates and issues the qualification
	Issuer *custom.Reference[Organization] `json:"issuer,omitempty"`
}

type PractitionerCommunication struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The language code used to communicate with the practitioner
	Language CodeableConcept `json:"language"`
	// Language preference indicator
	Preferred *bool `json:"preferred,omitempty"`
}

type OtherPractitioner Practitioner

func (p Practitioner) ResourceType() string {
	return "Practitioner"
}

func (p Practitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPractitioner: OtherPractitioner(p), ResourceType: p.ResourceType()})
}

func UnmarshalPractitioner(b []byte) (res Practitioner, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
