// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Flag
type Flag struct {
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
	// active | inactive | entered-in-error
	Status custom.Code `json:"status"`
	// Clinical, administrative, etc
	Category []CodeableConcept `json:"category,omitempty"`
	// Coded or textual message to display to user
	Code CodeableConcept `json:"code"`
	// Who/What is flag about?
	Subject custom.Reference[FlagSubject] `json:"subject"`
	// Time period when flag is active
	Period *Period `json:"period,omitempty"`
	// Alert relevant during encounter
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Flag creator
	Author *custom.Reference[FlagAuthor] `json:"author,omitempty"`
}

type OtherFlag Flag

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
	return "Flag"
}

func (f Flag) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFlag
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherFlag: OtherFlag(f), ResourceType: f.ResourceType()})
}

func UnmarshalFlag(b []byte) (res Flag, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
