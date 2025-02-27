// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CareTeam
type CareTeam struct {
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
	// External Ids for this team
	Identifier []Identifier `json:"identifier,omitempty"`
	// proposed | active | suspended | inactive | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// Type of team
	Category []CodeableConcept `json:"category,omitempty"`
	// Name of the team, such as crisis assessment team
	Name *string `json:"name,omitempty"`
	// Who care team is for
	Subject *custom.Reference[CareTeamSubject] `json:"subject,omitempty"`
	// Time period team covers
	Period *Period `json:"period,omitempty"`
	// Members of the team
	Participant []CareTeamParticipant `json:"participant,omitempty"`
	// Why the care team exists
	Reason []custom.CodeableReference[Condition] `json:"reason,omitempty"`
	// Organization responsible for the care team
	ManagingOrganization []custom.Reference[Organization] `json:"managingOrganization,omitempty"`
	// A contact detail for the care team (that applies to all members)
	Telecom []ContactPoint `json:"telecom,omitempty"`
	// Comments made about the CareTeam
	Note []Annotation `json:"note,omitempty"`
}

type CareTeamParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of involvement
	Role *CodeableConcept `json:"role,omitempty"`
	// Who is involved
	Member *custom.Reference[CareTeamParticipantMember] `json:"member,omitempty"`
	// Organization of the practitioner
	OnBehalfOf *custom.Reference[Organization] `json:"onBehalfOf,omitempty"`
	// When the member is generally available within this care team
	CoveragePeriod *Period `json:"coveragePeriod,omitempty"`
	// When the member is generally available within this care team
	CoverageTiming *Timing `json:"coverageTiming,omitempty"`
}

type OtherCareTeam CareTeam

type CareTeamSubject interface {
	gofhirclient.Resource

	Is_CareTeamSubject()
}

func (p Patient) Is_CareTeamSubject() {}
func (g Group) Is_CareTeamSubject()   {}

type CareTeamParticipantMember interface {
	gofhirclient.Resource

	Is_CareTeamParticipantMember()
}

func (p Practitioner) Is_CareTeamParticipantMember()     {}
func (p PractitionerRole) Is_CareTeamParticipantMember() {}
func (r RelatedPerson) Is_CareTeamParticipantMember()    {}
func (p Patient) Is_CareTeamParticipantMember()          {}
func (o Organization) Is_CareTeamParticipantMember()     {}
func (c CareTeam) Is_CareTeamParticipantMember()         {}

func (c CareTeam) ResourceType() string {
	return "CareTeam"
}

func (c CareTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCareTeam
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCareTeam: OtherCareTeam(c), ResourceType: c.ResourceType()})
}

func UnmarshalCareTeam(b []byte) (res CareTeam, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
