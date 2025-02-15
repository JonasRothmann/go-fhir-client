// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
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
	// External Ids for this goal
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// proposed | planned | accepted | active | on-hold | completed | cancelled | entered-in-error | rejected
	LifecycleStatus custom.Code `bson:"lifecycleStatus" json:"lifecycleStatus"`
	// in-progress | improving | worsening | no-change | achieved | sustaining | not-achieved | no-progress | not-attainable
	AchievementStatus *CodeableConcept `bson:"achievementStatus,omitempty" json:"achievementStatus,omitempty"`
	// E.g. Treatment, dietary, behavioral, etc
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// After meeting the goal, ongoing activity is needed to sustain the goal objective
	Continuous *bool `bson:"continuous,omitempty" json:"continuous,omitempty"`
	// high-priority | medium-priority | low-priority
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty"`
	// Code or text describing goal
	Description CodeableConcept `bson:"description" json:"description"`
	// Who this goal is intended for
	Subject custom.Reference[GoalSubject] `bson:"subject" json:"subject"`
	// When goal pursuit begins
	Start *custom.Date `bson:"start,omitempty" json:"start,omitempty"`
	// Target outcome for the goal
	Target []GoalTarget `bson:"target,omitempty" json:"target,omitempty"`
	// When goal status took effect
	StatusDate *custom.Date `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	// Reason for current status
	StatusReason *string `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// Who's responsible for creating Goal?
	Source *custom.Reference[GoalSource] `bson:"source,omitempty" json:"source,omitempty"`
	// Issues addressed by this goal
	Addresses []custom.Reference[GoalAddresses] `bson:"addresses,omitempty" json:"addresses,omitempty"`
	// Comments about the goal
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// What result was achieved regarding the goal?
	Outcome []custom.CodeableReference[Observation] `bson:"outcome,omitempty" json:"outcome,omitempty"`
}

type GoalTarget struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The parameter whose value is being tracked
	Measure *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
	// The target value to be achieved
	Detail *Quantity `bson:"detail,omitempty" json:"detail,omitempty"`
	// Reach goal on or before
	Due *custom.Date `bson:"due,omitempty" json:"due,omitempty"`
}

type GoalSubject interface {
	gofhirclient.Resource

	Is_GoalSubject()
}

func (p Patient) Is_GoalSubject()      {}
func (g Group) Is_GoalSubject()        {}
func (o Organization) Is_GoalSubject() {}

type GoalSource interface {
	gofhirclient.Resource

	Is_GoalSource()
}

func (p Patient) Is_GoalSource()          {}
func (p Practitioner) Is_GoalSource()     {}
func (p PractitionerRole) Is_GoalSource() {}
func (r RelatedPerson) Is_GoalSource()    {}
func (c CareTeam) Is_GoalSource()         {}

type GoalAddresses interface {
	gofhirclient.Resource

	Is_GoalAddresses()
}

func (c Condition) Is_GoalAddresses()           {}
func (o Observation) Is_GoalAddresses()         {}
func (m MedicationStatement) Is_GoalAddresses() {}
func (m MedicationRequest) Is_GoalAddresses()   {}
func (n NutritionOrder) Is_GoalAddresses()      {}
func (s ServiceRequest) Is_GoalAddresses()      {}
func (r RiskAssessment) Is_GoalAddresses()      {}
func (p Procedure) Is_GoalAddresses()           {}

func (g Goal) ResourceType() string {
	return "StructureDefinition"
}
