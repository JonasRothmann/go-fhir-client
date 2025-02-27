// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Person
type Person struct {
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
	// A human identifier for this person
	Identifier []Identifier `json:"identifier,omitempty"`
	// This person's record is in active use
	Active *bool `json:"active,omitempty"`
	// A name associated with the person
	Name []HumanName `json:"name,omitempty"`
	// A contact detail for the person
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `json:"gender,omitempty"`
	// The date on which the person was born
	BirthDate *custom.Date `json:"birthDate,omitempty"`
	// Indicates if the individual is deceased or not
	DeceasedBoolean *bool `json:"deceasedBoolean,omitempty"`
	// Indicates if the individual is deceased or not
	DeceasedDateTime *custom.DateTime `json:"deceasedDateTime,omitempty"`
	// One or more addresses for the person
	Address []Address `json:"address,omitempty"`
	// Marital (civil) status of a person
	MaritalStatus *CodeableConcept `json:"maritalStatus,omitempty"`
	// Image of the person
	Photo []Attachment `json:"photo,omitempty"`
	// A language which may be used to communicate with the person about his or her health
	Communication []PersonCommunication `json:"communication,omitempty"`
	// The organization that is the custodian of the person record
	ManagingOrganization *custom.Reference[Organization] `json:"managingOrganization,omitempty"`
	// Link to a resource that concerns the same actual person
	Link []PersonLink `json:"link,omitempty"`
}

type PersonCommunication struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The language which can be used to communicate with the person about his or her health
	Language CodeableConcept `json:"language"`
	// Language preference indicator
	Preferred *bool `json:"preferred,omitempty"`
}

type PersonLink struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The resource to which this actual person is associated
	Target custom.Reference[PersonLinkTarget] `json:"target"`
	// level1 | level2 | level3 | level4
	Assurance *custom.Code `json:"assurance,omitempty"`
}

type OtherPerson Person

type PersonLinkTarget interface {
	gofhirclient.Resource

	Is_PersonLinkTarget()
}

func (p Patient) Is_PersonLinkTarget()       {}
func (p Practitioner) Is_PersonLinkTarget()  {}
func (r RelatedPerson) Is_PersonLinkTarget() {}
func (p Person) Is_PersonLinkTarget()        {}

func (p Person) ResourceType() string {
	return "Person"
}

func (p Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPerson
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPerson: OtherPerson(p), ResourceType: p.ResourceType()})
}

func UnmarshalPerson(b []byte) (res Person, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
