// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Group
type GroupMember struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to the group member
	Entity custom.Reference[GroupMemberEntity] `bson:"entity" json:"entity"`
	// Period member belonged to the group
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// If member is no longer in group
	Inactive *bool `bson:"inactive,omitempty" json:"inactive,omitempty"`
}

type Group struct {
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
	// Business Identifier for this Group
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether this group's record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// person | animal | practitioner | device | careteam | healthcareservice | location | organization | relatedperson | specimen
	Type custom.Code `bson:"type" json:"type"`
	// definitional | enumerated
	Membership custom.Code `bson:"membership" json:"membership"`
	// Kind of Group members
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Label for Group
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Natural language description of the group
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Number of members
	Quantity *uint32 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Entity that is the custodian of the Group's definition
	ManagingEntity *custom.Reference[GroupManagingEntity] `bson:"managingEntity,omitempty" json:"managingEntity,omitempty"`
	// Include / Exclude group members by Trait
	Characteristic []GroupCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	// Who or what is in group
	Member []GroupMember `bson:"member,omitempty" json:"member,omitempty"`
}

type GroupCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Kind of characteristic
	Code CodeableConcept `bson:"code" json:"code"`
	// Value held by characteristic
	Value CodeableConcept `bson:"value" json:"value"`
	// Group includes or excludes
	Exclude bool `bson:"exclude" json:"exclude"`
	// Period over which characteristic is tested
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type GroupManagingEntity interface {
	gofhirclient.Resource

	Is_GroupManagingEntity()
}

func (o Organization) Is_GroupManagingEntity()     {}
func (r RelatedPerson) Is_GroupManagingEntity()    {}
func (p Practitioner) Is_GroupManagingEntity()     {}
func (p PractitionerRole) Is_GroupManagingEntity() {}

type GroupMemberEntity interface {
	gofhirclient.Resource

	Is_GroupMemberEntity()
}

func (c CareTeam) Is_GroupMemberEntity()          {}
func (d Device) Is_GroupMemberEntity()            {}
func (g Group) Is_GroupMemberEntity()             {}
func (h HealthcareService) Is_GroupMemberEntity() {}
func (l Location) Is_GroupMemberEntity()          {}
func (o Organization) Is_GroupMemberEntity()      {}
func (p Patient) Is_GroupMemberEntity()           {}
func (p Practitioner) Is_GroupMemberEntity()      {}
func (p PractitionerRole) Is_GroupMemberEntity()  {}
func (r RelatedPerson) Is_GroupMemberEntity()     {}
func (s Specimen) Is_GroupMemberEntity()          {}

func (g Group) ResourceType() string {
	return "StructureDefinition"
}
