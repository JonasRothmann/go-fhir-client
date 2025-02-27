// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Goal
type Goal struct {
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
	// External Ids for this goal
	Identifier []Identifier `json:"identifier,omitempty"`
	// proposed | planned | accepted | active | on-hold | completed | cancelled | entered-in-error | rejected
	LifecycleStatus custom.Code `json:"lifecycleStatus"`
	// in-progress | improving | worsening | no-change | achieved | sustaining | not-achieved | no-progress | not-attainable
	AchievementStatus *CodeableConcept `json:"achievementStatus,omitempty"`
	// E.g. Treatment, dietary, behavioral, etc
	Category []CodeableConcept `json:"category,omitempty"`
	// After meeting the goal, ongoing activity is needed to sustain the goal objective
	Continuous *bool `json:"continuous,omitempty"`
	// high-priority | medium-priority | low-priority
	Priority *CodeableConcept `json:"priority,omitempty"`
	// Code or text describing goal
	Description CodeableConcept `json:"description"`
	// Who this goal is intended for
	Subject custom.Reference[GoalSubject] `json:"subject"`
	// When goal pursuit begins
	StartDate *custom.Date `json:"startDate,omitempty"`
	// When goal pursuit begins
	StartCodeableConcept *CodeableConcept `json:"startCodeableConcept,omitempty"`
	// Target outcome for the goal
	Target []GoalTarget `json:"target,omitempty"`
	// When goal status took effect
	StatusDate *custom.Date `json:"statusDate,omitempty"`
	// Reason for current status
	StatusReason *string `json:"statusReason,omitempty"`
	// Who's responsible for creating Goal?
	Source *custom.Reference[GoalSource] `json:"source,omitempty"`
	// Issues addressed by this goal
	Addresses []custom.Reference[GoalAddresses] `json:"addresses,omitempty"`
	// Comments about the goal
	Note []Annotation `json:"note,omitempty"`
	// What result was achieved regarding the goal?
	Outcome []custom.CodeableReference[Observation] `json:"outcome,omitempty"`
}

type GoalTarget struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The parameter whose value is being tracked
	Measure *CodeableConcept `json:"measure,omitempty"`
	// The target value to be achieved
	DetailQuantity *Quantity `json:"detailQuantity,omitempty"`
	// The target value to be achieved
	DetailRange *Range `json:"detailRange,omitempty"`
	// The target value to be achieved
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty"`
	// The target value to be achieved
	DetailString *string `json:"detailString,omitempty"`
	// The target value to be achieved
	DetailBoolean *bool `json:"detailBoolean,omitempty"`
	// The target value to be achieved
	DetailInteger *int32 `json:"detailInteger,omitempty"`
	// The target value to be achieved
	DetailRatio *Ratio `json:"detailRatio,omitempty"`
	// Reach goal on or before
	DueDate *custom.Date `json:"dueDate,omitempty"`
	// Reach goal on or before
	DueDuration *Duration `json:"dueDuration,omitempty"`
}

type OtherGoal Goal

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
	return "Goal"
}

func (g Goal) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGoal
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherGoal: OtherGoal(g), ResourceType: g.ResourceType()})
}

func UnmarshalGoal(b []byte) (res Goal, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
