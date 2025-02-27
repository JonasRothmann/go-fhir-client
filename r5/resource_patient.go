// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Patient
type PatientContact struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The kind of relationship
	Relationship []CodeableConcept `json:"relationship,omitempty"`
	// A name associated with the contact person
	Name *HumanName `json:"name,omitempty"`
	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// Address for the contact person
	Address *Address `json:"address,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `json:"gender,omitempty"`
	// Organization that is associated with the contact
	Organization *custom.Reference[Organization] `json:"organization,omitempty"`
	// The period during which this contact person or organization is valid to be contacted relating to this patient
	Period *Period `json:"period,omitempty"`
}

type PatientCommunication struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The language which can be used to communicate with the patient about his or her health
	Language CodeableConcept `json:"language"`
	// Language preference indicator
	Preferred *bool `json:"preferred,omitempty"`
}

type PatientLink struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The other patient or related person resource that the link refers to
	Other custom.Reference[PatientLinkOther] `json:"other"`
	// replaced-by | replaces | refer | seealso
	Type custom.Code `json:"type"`
}

type Patient struct {
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
	// An identifier for this patient
	Identifier []Identifier `json:"identifier,omitempty"`
	// Whether this patient's record is in active use
	Active *bool `json:"active,omitempty"`
	// A name associated with the patient
	Name []HumanName `json:"name,omitempty"`
	// A contact detail for the individual
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `json:"gender,omitempty"`
	// The date of birth for the individual
	BirthDate *custom.Date `json:"birthDate,omitempty"`
	// Indicates if the individual is deceased or not
	DeceasedBoolean *bool `json:"deceasedBoolean,omitempty"`
	// Indicates if the individual is deceased or not
	DeceasedDateTime *custom.DateTime `json:"deceasedDateTime,omitempty"`
	// An address for the individual
	Address []Address `json:"address,omitempty"`
	// Marital (civil) status of a patient
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`
	// Whether patient is part of a multiple birth
	MultipleBirthBoolean *bool `json:"multipleBirthBoolean,omitempty"`
	// Whether patient is part of a multiple birth
	MultipleBirthInteger *int32 `json:"multipleBirthInteger,omitempty"`
	// Image of the patient
	Photo []Attachment `json:"photo,omitempty"`
	// A contact party (e.g. guardian, partner, friend) for the patient
	Contact []PatientContact `json:"contact,omitempty"`
	// A language which may be used to communicate with the patient about his or her health
	Communication []PatientCommunication `json:"communication,omitempty"`
	// Patient's nominated primary care provider
	GeneralPractitioner []custom.Reference[PatientGeneralPractitioner] `json:"generalPractitioner,omitempty"`
	// Organization that is the custodian of the patient record
	ManagingOrganization *custom.Reference[Organization] `json:"managingOrganization,omitempty"`
	// Link to a Patient or RelatedPerson resource that concerns the same actual individual
	Link []PatientLink `json:"link,omitempty"`
}

type OtherPatient Patient

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
	return "Patient"
}

func (p Patient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPatient
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPatient: OtherPatient(p), ResourceType: p.ResourceType()})
}

func UnmarshalPatient(b []byte) (res Patient, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
