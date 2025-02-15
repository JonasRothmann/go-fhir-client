// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CareTeam
type CareTeam struct {
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
	// External Ids for this team
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// proposed | active | suspended | inactive | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Type of team
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Name of the team, such as crisis assessment team
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Who care team is for
	Subject *custom.Reference[CareTeamSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Time period team covers
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Members of the team
	Participant []CareTeamParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// Why the care team exists
	Reason []custom.CodeableReference[Condition] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Organization responsible for the care team
	ManagingOrganization []custom.Reference[Organization] `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	// A contact detail for the care team (that applies to all members)
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
	// Comments made about the CareTeam
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type CareTeamParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of involvement
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Who is involved
	Member *custom.Reference[CareTeamParticipantMember] `bson:"member,omitempty" json:"member,omitempty"`
	// Organization of the practitioner
	OnBehalfOf *custom.Reference[Organization] `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	// When the member is generally available within this care team
	Coverage *Period `bson:"coverage,omitempty" json:"coverage,omitempty"`
}

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
	return "StructureDefinition"
}
