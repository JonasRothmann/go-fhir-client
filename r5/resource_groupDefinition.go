// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/groupdefinition
type GroupDefinition struct {
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
	ManagingEntity *custom.Reference[GroupDefinitionManagingEntity] `json:"managingEntity,omitempty"`
	// Include / Exclude group members by Trait
	Characteristic []GroupDefinitionCharacteristic `json:"characteristic"`
	// Who or what is in group
	Member []GroupDefinitionMember `json:"member,omitempty"`
}

type GroupDefinitionCharacteristic struct {
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

type GroupDefinitionMember struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to the group member
	Entity custom.Reference[GroupDefinitionMemberEntity] `json:"entity"`
	// Period member belonged to the group
	Period *Period `json:"period,omitempty"`
	// If member is no longer in group
	Inactive *bool `json:"inactive,omitempty"`
}

type OtherGroupDefinition GroupDefinition

type GroupDefinitionManagingEntity interface {
	gofhirclient.Resource

	Is_GroupDefinitionManagingEntity()
}

func (o Organization) Is_GroupDefinitionManagingEntity()     {}
func (r RelatedPerson) Is_GroupDefinitionManagingEntity()    {}
func (p Practitioner) Is_GroupDefinitionManagingEntity()     {}
func (p PractitionerRole) Is_GroupDefinitionManagingEntity() {}

type GroupDefinitionMemberEntity interface {
	gofhirclient.Resource

	Is_GroupDefinitionMemberEntity()
}

func (c CareTeam) Is_GroupDefinitionMemberEntity()          {}
func (d Device) Is_GroupDefinitionMemberEntity()            {}
func (g Group) Is_GroupDefinitionMemberEntity()             {}
func (h HealthcareService) Is_GroupDefinitionMemberEntity() {}
func (l Location) Is_GroupDefinitionMemberEntity()          {}
func (o Organization) Is_GroupDefinitionMemberEntity()      {}
func (p Patient) Is_GroupDefinitionMemberEntity()           {}
func (p Practitioner) Is_GroupDefinitionMemberEntity()      {}
func (p PractitionerRole) Is_GroupDefinitionMemberEntity()  {}
func (r RelatedPerson) Is_GroupDefinitionMemberEntity()     {}
func (s Specimen) Is_GroupDefinitionMemberEntity()          {}

func (g GroupDefinition) ResourceType() string {
	return "GroupDefinition"
}

func (g GroupDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGroupDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherGroupDefinition: OtherGroupDefinition(g), ResourceType: g.ResourceType()})
}

func UnmarshalGroupDefinition(b []byte) (res GroupDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
