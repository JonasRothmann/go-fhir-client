// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AuditEvent
type AuditEventEntity struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Specific instance of resource
	What *custom.Reference[Resource] `bson:"what,omitempty" json:"what,omitempty"`
	// What role the entity played
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Security labels on the entity
	SecurityLabel []CodeableConcept `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	// Query parameters
	Query *custom.Base64binary `bson:"query,omitempty" json:"query,omitempty"`
	// Additional Information about the entity
	Detail []AuditEventEntityDetail `bson:"detail,omitempty" json:"detail,omitempty"`
	// Entity is attributed to this agent
	Agent []interface{} `bson:"agent,omitempty" json:"agent,omitempty"`
}

type AuditEventEntityDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name of the property
	Type CodeableConcept `bson:"type" json:"type"`
	// Property value
	Value Quantity `bson:"value" json:"value"`
}

type AuditEvent struct {
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
	// Type/identifier of event
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Specific type of event
	Code CodeableConcept `bson:"code" json:"code"`
	// Type of action performed during the event
	Action *custom.Code `bson:"action,omitempty" json:"action,omitempty"`
	// emergency | alert | critical | error | warning | notice | informational | debug
	Severity *custom.Code `bson:"severity,omitempty" json:"severity,omitempty"`
	// When the activity occurred
	Occurred *Period `bson:"occurred,omitempty" json:"occurred,omitempty"`
	// Time when the event was recorded
	Recorded custom.Instant `bson:"recorded" json:"recorded"`
	// Whether the event succeeded or failed
	Outcome *AuditEventOutcome `bson:"outcome,omitempty" json:"outcome,omitempty"`
	// Authorization related to the event
	Authorization []CodeableConcept `bson:"authorization,omitempty" json:"authorization,omitempty"`
	// Workflow authorization within which this event occurred
	BasedOn []custom.Reference[AuditEventBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// The patient is the subject of the data used/created/updated/deleted during the activity
	Patient *custom.Reference[Patient] `bson:"patient,omitempty" json:"patient,omitempty"`
	// Encounter within which this event occurred or which the event is tightly associated
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Actor involved in the event
	Agent []AuditEventAgent `bson:"agent" json:"agent"`
	// Audit Event Reporter
	Source AuditEventSource `bson:"source" json:"source"`
	// Data or objects used
	Entity []AuditEventEntity `bson:"entity,omitempty" json:"entity,omitempty"`
}

type AuditEventOutcome struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Whether the event succeeded or failed
	Code Coding `bson:"code" json:"code"`
	// Additional outcome detail
	Detail []CodeableConcept `bson:"detail,omitempty" json:"detail,omitempty"`
}

type AuditEventAgent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// How agent participated
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Agent role in the event
	Role []CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Identifier of who
	Who custom.Reference[AuditEventAgentWho] `bson:"who" json:"who"`
	// Whether user is initiator
	Requestor *bool `bson:"requestor,omitempty" json:"requestor,omitempty"`
	// The agent location when the event occurred
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Policy that authorized the agent participation in the event
	Policy []custom.URI `bson:"policy,omitempty" json:"policy,omitempty"`
	// This agent network location for the activity
	Network *custom.Reference[Endpoint] `bson:"network,omitempty" json:"network,omitempty"`
	// Allowable authorization for this agent
	Authorization []CodeableConcept `bson:"authorization,omitempty" json:"authorization,omitempty"`
}

type AuditEventSource struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Logical source location within the enterprise
	Site *custom.Reference[Location] `bson:"site,omitempty" json:"site,omitempty"`
	// The identity of source detecting the event
	Observer custom.Reference[AuditEventSourceObserver] `bson:"observer" json:"observer"`
	// The type of source where event originated
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}

type AuditEventBasedOn interface {
	gofhirclient.Resource

	Is_AuditEventBasedOn()
}

func (c CarePlan) Is_AuditEventBasedOn()                   {}
func (d DeviceRequest) Is_AuditEventBasedOn()              {}
func (i ImmunizationRecommendation) Is_AuditEventBasedOn() {}
func (m MedicationRequest) Is_AuditEventBasedOn()          {}
func (n NutritionOrder) Is_AuditEventBasedOn()             {}
func (s ServiceRequest) Is_AuditEventBasedOn()             {}
func (t Task) Is_AuditEventBasedOn()                       {}

type AuditEventAgentWho interface {
	gofhirclient.Resource

	Is_AuditEventAgentWho()
}

func (p Practitioner) Is_AuditEventAgentWho()     {}
func (p PractitionerRole) Is_AuditEventAgentWho() {}
func (o Organization) Is_AuditEventAgentWho()     {}
func (c CareTeam) Is_AuditEventAgentWho()         {}
func (p Patient) Is_AuditEventAgentWho()          {}
func (d Device) Is_AuditEventAgentWho()           {}
func (r RelatedPerson) Is_AuditEventAgentWho()    {}

type AuditEventSourceObserver interface {
	gofhirclient.Resource

	Is_AuditEventSourceObserver()
}

func (p Practitioner) Is_AuditEventSourceObserver()     {}
func (p PractitionerRole) Is_AuditEventSourceObserver() {}
func (o Organization) Is_AuditEventSourceObserver()     {}
func (c CareTeam) Is_AuditEventSourceObserver()         {}
func (p Patient) Is_AuditEventSourceObserver()          {}
func (d Device) Is_AuditEventSourceObserver()           {}
func (r RelatedPerson) Is_AuditEventSourceObserver()    {}

func (a AuditEvent) ResourceType() string {
	return "StructureDefinition"
}
