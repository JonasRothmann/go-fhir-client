// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Provenance
type Provenance struct {
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
	// Target Reference(s) (usually version specific)
	Target []custom.Reference[Resource] `json:"target"`
	// When the activity occurred
	OccurredPeriod *Period `json:"occurredPeriod,omitempty"`
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
	// Activity that occurred
	Activity *CodeableConcept `json:"activity,omitempty"`
	// Workflow authorization within which this event occurred
	BasedOn []custom.Reference[ProvenanceBasedOn] `json:"basedOn,omitempty"`
	// The patient is the subject of the data created/updated (.target) by the activity
	Patient *custom.Reference[Patient] `json:"patient,omitempty"`
	// Encounter within which this event occurred or which the event is tightly associated
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Actor involved
	Agent []ProvenanceAgent `json:"agent"`
	// An entity used in this activity
	Entity []ProvenanceEntity `json:"entity,omitempty"`
	// Signature on target
	Signature []Signature `json:"signature,omitempty"`
}

type ProvenanceAgent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// How the agent participated
	Type *CodeableConcept `json:"type,omitempty"`
	// What the agents role was
	Role []CodeableConcept `json:"role,omitempty"`
	// The agent that participated in the event
	Who custom.Reference[ProvenanceAgentWho] `json:"who"`
	// The agent that delegated
	OnBehalfOf *custom.Reference[ProvenanceAgentOnBehalfOf] `json:"onBehalfOf,omitempty"`
}

type ProvenanceEntity struct {
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

type OtherProvenance Provenance

type ProvenanceBasedOn interface {
	gofhirclient.Resource

	Is_ProvenanceBasedOn()
}

func (c CarePlan) Is_ProvenanceBasedOn()                   {}
func (d DeviceRequest) Is_ProvenanceBasedOn()              {}
func (i ImmunizationRecommendation) Is_ProvenanceBasedOn() {}
func (m MedicationRequest) Is_ProvenanceBasedOn()          {}
func (n NutritionOrder) Is_ProvenanceBasedOn()             {}
func (s ServiceRequest) Is_ProvenanceBasedOn()             {}
func (t Task) Is_ProvenanceBasedOn()                       {}

type ProvenanceAgentWho interface {
	gofhirclient.Resource

	Is_ProvenanceAgentWho()
}

func (p Practitioner) Is_ProvenanceAgentWho()     {}
func (p PractitionerRole) Is_ProvenanceAgentWho() {}
func (o Organization) Is_ProvenanceAgentWho()     {}
func (c CareTeam) Is_ProvenanceAgentWho()         {}
func (p Patient) Is_ProvenanceAgentWho()          {}
func (d Device) Is_ProvenanceAgentWho()           {}
func (r RelatedPerson) Is_ProvenanceAgentWho()    {}

type ProvenanceAgentOnBehalfOf interface {
	gofhirclient.Resource

	Is_ProvenanceAgentOnBehalfOf()
}

func (p Practitioner) Is_ProvenanceAgentOnBehalfOf()     {}
func (p PractitionerRole) Is_ProvenanceAgentOnBehalfOf() {}
func (o Organization) Is_ProvenanceAgentOnBehalfOf()     {}
func (c CareTeam) Is_ProvenanceAgentOnBehalfOf()         {}
func (p Patient) Is_ProvenanceAgentOnBehalfOf()          {}

func (p Provenance) ResourceType() string {
	return "Provenance"
}

func (p Provenance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProvenance
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherProvenance: OtherProvenance(p), ResourceType: p.ResourceType()})
}

func UnmarshalProvenance(b []byte) (res Provenance, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
