// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Provenance
type Provenance struct {
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
	// Target Reference(s) (usually version specific)
	Target []custom.Reference[Resource] `bson:"target" json:"target"`
	// When the activity occurred
	Occurred *Period `bson:"occurred,omitempty" json:"occurred,omitempty"`
	// When the activity was recorded / updated
	Recorded *custom.Instant `bson:"recorded,omitempty" json:"recorded,omitempty"`
	// Policy or plan the activity was defined by
	Policy []custom.URI `bson:"policy,omitempty" json:"policy,omitempty"`
	// Where the activity occurred, if relevant
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Authorization (purposeOfUse) related to the event
	Authorization []custom.CodeableReference `bson:"authorization,omitempty" json:"authorization,omitempty"`
	// Activity that occurred
	Activity *CodeableConcept `bson:"activity,omitempty" json:"activity,omitempty"`
	// Workflow authorization within which this event occurred
	BasedOn []custom.Reference[ProvenanceBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// The patient is the subject of the data created/updated (.target) by the activity
	Patient *custom.Reference[Patient] `bson:"patient,omitempty" json:"patient,omitempty"`
	// Encounter within which this event occurred or which the event is tightly associated
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Actor involved
	Agent []ProvenanceAgent `bson:"agent" json:"agent"`
	// An entity used in this activity
	Entity []ProvenanceEntity `bson:"entity,omitempty" json:"entity,omitempty"`
	// Signature on target
	Signature []Signature `bson:"signature,omitempty" json:"signature,omitempty"`
}

type ProvenanceAgent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// How the agent participated
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// What the agents role was
	Role []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// The agent that participated in the event
	Who custom.Reference[ProvenanceAgentWho] `bson:"who" json:"who"`
	// The agent that delegated
	OnBehalfOf *custom.Reference[ProvenanceAgentOnBehalfOf] `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
}

type ProvenanceEntity struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// revision | quotation | source | instantiates | removal
	Role custom.Code `bson:"role" json:"role"`
	// Identity of entity
	What custom.Reference[Resource] `bson:"what" json:"what"`
	// Entity is attributed to this agent
	Agent []interface{} `bson:"agent,omitempty" json:"agent,omitempty"`
}

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
	return "StructureDefinition"
}
