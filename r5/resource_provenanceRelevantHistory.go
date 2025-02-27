// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/provenance-relevant-history
type ProvenanceRelevantHistory struct {
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
	// Resource version
	Target []custom.Reference[Resource] `json:"target"`
	// When the activity occurred
	OccurredDateTime *custom.DateTime `json:"occurredDateTime,omitempty"`
	// When the activity was recorded / updated
	Recorded *custom.Instant `json:"recorded,omitempty"`
	// Policy or plan the activity was defined by
	Policy []custom.URI `json:"policy,omitempty"`
	// Where the activity occurred, if relevant
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Authorization (purposeOfUse) related to the event
	Authorization []custom.CodeableReference[gofhirclient.Resource] `json:"authorization,omitempty"`
	// Record activity
	Activity CodeableConcept `json:"activity"`
	// Workflow authorization within which this event occurred
	BasedOn []custom.Reference[ProvenanceRelevantHistoryBasedOn] `json:"basedOn,omitempty"`
	// The patient is the subject of the data created/updated (.target) by the activity
	Patient *custom.Reference[Patient] `json:"patient,omitempty"`
	// Encounter within which this event occurred or which the event is tightly associated
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Who was involved with change
	Agent []ProvenanceRelevantHistoryAgent `json:"agent"`
	// An entity used in this activity
	Entity []ProvenanceRelevantHistoryEntity `json:"entity,omitempty"`
	// Signature on target
	Signature []Signature `json:"signature,omitempty"`
}

type ProvenanceRelevantHistoryAgent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// How the agent participated
	Type CodeableConcept `json:"type"`
	// What the agents role was
	Role []CodeableConcept `json:"role,omitempty"`
	// The agent that participated in the event
	Who custom.Reference[ProvenanceRelevantHistoryAgentWho] `json:"who"`
	// The agent that delegated
	OnBehalfOf *custom.Reference[ProvenanceRelevantHistoryAgentOnBehalfOf] `json:"onBehalfOf,omitempty"`
}

type ProvenanceRelevantHistoryEntity struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// revision | quotation | source | instantiates | removal
	Role custom.Code `json:"role"`
	// Identity of entity
	What custom.Reference[Resource] `json:"what"`
	// Entity is attributed to this agent
	Agent []interface{} `json:"agent,omitempty"`
}

type OtherProvenanceRelevantHistory ProvenanceRelevantHistory

type ProvenanceRelevantHistoryBasedOn interface {
	gofhirclient.Resource

	Is_ProvenanceRelevantHistoryBasedOn()
}

func (c CarePlan) Is_ProvenanceRelevantHistoryBasedOn()                   {}
func (d DeviceRequest) Is_ProvenanceRelevantHistoryBasedOn()              {}
func (i ImmunizationRecommendation) Is_ProvenanceRelevantHistoryBasedOn() {}
func (m MedicationRequest) Is_ProvenanceRelevantHistoryBasedOn()          {}
func (n NutritionOrder) Is_ProvenanceRelevantHistoryBasedOn()             {}
func (s ServiceRequest) Is_ProvenanceRelevantHistoryBasedOn()             {}
func (t Task) Is_ProvenanceRelevantHistoryBasedOn()                       {}

type ProvenanceRelevantHistoryAgentWho interface {
	gofhirclient.Resource

	Is_ProvenanceRelevantHistoryAgentWho()
}

func (p Practitioner) Is_ProvenanceRelevantHistoryAgentWho()     {}
func (p PractitionerRole) Is_ProvenanceRelevantHistoryAgentWho() {}
func (o Organization) Is_ProvenanceRelevantHistoryAgentWho()     {}
func (c CareTeam) Is_ProvenanceRelevantHistoryAgentWho()         {}
func (p Patient) Is_ProvenanceRelevantHistoryAgentWho()          {}
func (d Device) Is_ProvenanceRelevantHistoryAgentWho()           {}
func (r RelatedPerson) Is_ProvenanceRelevantHistoryAgentWho()    {}

type ProvenanceRelevantHistoryAgentOnBehalfOf interface {
	gofhirclient.Resource

	Is_ProvenanceRelevantHistoryAgentOnBehalfOf()
}

func (p Practitioner) Is_ProvenanceRelevantHistoryAgentOnBehalfOf()     {}
func (p PractitionerRole) Is_ProvenanceRelevantHistoryAgentOnBehalfOf() {}
func (o Organization) Is_ProvenanceRelevantHistoryAgentOnBehalfOf()     {}
func (c CareTeam) Is_ProvenanceRelevantHistoryAgentOnBehalfOf()         {}
func (p Patient) Is_ProvenanceRelevantHistoryAgentOnBehalfOf()          {}

func (p ProvenanceRelevantHistory) ResourceType() string {
	return "ProvenanceRelevantHistory"
}

func (p ProvenanceRelevantHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProvenanceRelevantHistory
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherProvenanceRelevantHistory: OtherProvenanceRelevantHistory(p), ResourceType: p.ResourceType()})
}

func UnmarshalProvenanceRelevantHistory(b []byte) (res ProvenanceRelevantHistory, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
