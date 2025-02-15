// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Schedule
type Schedule struct {
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
	// External Ids for this item
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether this schedule is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// High-level category
	ServiceCategory []CodeableConcept `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	// Specific service
	ServiceType []custom.CodeableReference[HealthcareService] `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	// Type of specialty needed
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	// Human-readable label
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Resource(s) that availability information is being provided for
	Actor []custom.Reference[ScheduleActor] `bson:"actor" json:"actor"`
	// Period of time covered by schedule
	PlanningHorizon *Period `bson:"planningHorizon,omitempty" json:"planningHorizon,omitempty"`
	// Comments on availability
	Comment *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
}

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
	return "StructureDefinition"
}
