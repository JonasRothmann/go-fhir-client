// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Basic
type Basic struct {
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
	// Business identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Kind of Resource
	Code CodeableConcept `json:"code"`
	// Identifies the focus of this resource
	Subject *custom.Reference[Resource] `json:"subject,omitempty"`
	// When created
	Created *custom.DateTime `json:"created,omitempty"`
	// Who created
	Author *custom.Reference[BasicAuthor] `json:"author,omitempty"`
}

type OtherBasic Basic

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
	return "Basic"
}

func (b Basic) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBasic
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherBasic: OtherBasic(b), ResourceType: b.ResourceType()})
}

func UnmarshalBasic(b []byte) (res Basic, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
