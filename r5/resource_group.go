// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Group
type Group struct {
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
	// Business Identifier for this Group
	Identifier []Identifier `json:"identifier,omitempty"`
	// Whether this group's record is in active use
	Active *bool `json:"active,omitempty"`
	// person | animal | practitioner | device | careteam | healthcareservice | location | organization | relatedperson | specimen
	Type custom.Code `json:"type"`
	// definitional | enumerated
	Membership custom.Code `json:"membership"`
	// Kind of Group members
	Code *CodeableConcept `json:"code,omitempty"`
	// Label for Group
	Name *string `json:"name,omitempty"`
	// Natural language description of the group
	Description *custom.Markdown `json:"description,omitempty"`
	// Number of members
	Quantity *uint32 `json:"quantity,omitempty"`
	// Entity that is the custodian of the Group's definition
	ManagingEntity *custom.Reference[GroupManagingEntity] `json:"managingEntity,omitempty"`
	// Include / Exclude group members by Trait
	Characteristic []GroupCharacteristic `json:"characteristic,omitempty"`
	// Who or what is in group
	Member []GroupMember `json:"member,omitempty"`
}

type GroupCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Kind of characteristic
	Code CodeableConcept `json:"code"`
	// Value held by characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Value held by characteristic
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Value held by characteristic
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Value held by characteristic
	ValueRange *Range `json:"valueRange,omitempty"`
	// Value held by characteristic
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// Group includes or excludes
	Exclude bool `json:"exclude"`
	// Period over which characteristic is tested
	Period *Period `json:"period,omitempty"`
}

type GroupMember struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to the group member
	Entity custom.Reference[GroupMemberEntity] `json:"entity"`
	// Period member belonged to the group
	Period *Period `json:"period,omitempty"`
	// If member is no longer in group
	Inactive *bool `json:"inactive,omitempty"`
}

type OtherGroup Group

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
	return "Group"
}

func (g Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGroup
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherGroup: OtherGroup(g), ResourceType: g.ResourceType()})
}

func UnmarshalGroup(b []byte) (res Group, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
