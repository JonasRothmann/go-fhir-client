// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/actualgroup
type ActualGroup struct {
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
	ManagingEntity *custom.Reference[ActualGroupManagingEntity] `json:"managingEntity,omitempty"`
	// Include / Exclude group members by Trait
	Characteristic []ActualGroupCharacteristic `json:"characteristic,omitempty"`
	// Who or what is in group
	Member []ActualGroupMember `json:"member,omitempty"`
}

type ActualGroupCharacteristic struct {
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

type ActualGroupMember struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to the group member
	Entity custom.Reference[ActualGroupMemberEntity] `json:"entity"`
	// Period member belonged to the group
	Period *Period `json:"period,omitempty"`
	// If member is no longer in group
	Inactive *bool `json:"inactive,omitempty"`
}

type OtherActualGroup ActualGroup

type ActualGroupManagingEntity interface {
	gofhirclient.Resource

	Is_ActualGroupManagingEntity()
}

func (o Organization) Is_ActualGroupManagingEntity()     {}
func (r RelatedPerson) Is_ActualGroupManagingEntity()    {}
func (p Practitioner) Is_ActualGroupManagingEntity()     {}
func (p PractitionerRole) Is_ActualGroupManagingEntity() {}

type ActualGroupMemberEntity interface {
	gofhirclient.Resource

	Is_ActualGroupMemberEntity()
}

func (c CareTeam) Is_ActualGroupMemberEntity()          {}
func (d Device) Is_ActualGroupMemberEntity()            {}
func (g Group) Is_ActualGroupMemberEntity()             {}
func (h HealthcareService) Is_ActualGroupMemberEntity() {}
func (l Location) Is_ActualGroupMemberEntity()          {}
func (o Organization) Is_ActualGroupMemberEntity()      {}
func (p Patient) Is_ActualGroupMemberEntity()           {}
func (p Practitioner) Is_ActualGroupMemberEntity()      {}
func (p PractitionerRole) Is_ActualGroupMemberEntity()  {}
func (r RelatedPerson) Is_ActualGroupMemberEntity()     {}
func (s Specimen) Is_ActualGroupMemberEntity()          {}

func (a ActualGroup) ResourceType() string {
	return "ActualGroup"
}

func (a ActualGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActualGroup
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherActualGroup: OtherActualGroup(a), ResourceType: a.ResourceType()})
}

func UnmarshalActualGroup(b []byte) (res ActualGroup, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
