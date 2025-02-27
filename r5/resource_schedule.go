// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Schedule
type Schedule struct {
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
	// External Ids for this item
	Identifier []Identifier `json:"identifier,omitempty"`
	// Whether this schedule is in active use
	Active *bool `json:"active,omitempty"`
	// High-level category
	ServiceCategory []CodeableConcept `json:"serviceCategory,omitempty"`
	// Specific service
	ServiceType []custom.CodeableReference[HealthcareService] `json:"serviceType,omitempty"`
	// Type of specialty needed
	Specialty []CodeableConcept `json:"specialty,omitempty"`
	// Human-readable label
	Name *string `json:"name,omitempty"`
	// Resource(s) that availability information is being provided for
	Actor []custom.Reference[ScheduleActor] `json:"actor"`
	// Period of time covered by schedule
	PlanningHorizon *Period `json:"planningHorizon,omitempty"`
	// Comments on availability
	Comment *custom.Markdown `json:"comment,omitempty"`
}

type OtherSchedule Schedule

type ScheduleActor interface {
	gofhirclient.Resource

	Is_ScheduleActor()
}

func (p Patient) Is_ScheduleActor()           {}
func (p Practitioner) Is_ScheduleActor()      {}
func (p PractitionerRole) Is_ScheduleActor()  {}
func (c CareTeam) Is_ScheduleActor()          {}
func (r RelatedPerson) Is_ScheduleActor()     {}
func (d Device) Is_ScheduleActor()            {}
func (h HealthcareService) Is_ScheduleActor() {}
func (l Location) Is_ScheduleActor()          {}

func (s Schedule) ResourceType() string {
	return "Schedule"
}

func (s Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSchedule
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSchedule: OtherSchedule(s), ResourceType: s.ResourceType()})
}

func UnmarshalSchedule(b []byte) (res Schedule, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
