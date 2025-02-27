// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AuditEvent
type AuditEventEntityDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name of the property
	Type CodeableConcept `json:"type"`
	// Property value
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Property value
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Property value
	ValueString *string `json:"valueString,omitempty"`
	// Property value
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Property value
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Property value
	ValueRange *Range `json:"valueRange,omitempty"`
	// Property value
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Property value
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Property value
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Property value
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Property value
	ValueBase64binary *custom.Base64binary `json:"valueBase64binary,omitempty"`
}

type AuditEvent struct {
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
	// Type/identifier of event
	Category []CodeableConcept `json:"category,omitempty"`
	// Specific type of event
	Code CodeableConcept `json:"code"`
	// Type of action performed during the event
	Action *custom.Code `json:"action,omitempty"`
	// emergency | alert | critical | error | warning | notice | informational | debug
	Severity *custom.Code `json:"severity,omitempty"`
	// When the activity occurred
	OccurredPeriod *Period `json:"occurredPeriod,omitempty"`
	// When the activity occurred
	OccurredDateTime *custom.DateTime `json:"occurredDateTime,omitempty"`
	// Time when the event was recorded
	Recorded custom.Instant `json:"recorded"`
	// Whether the event succeeded or failed
	Outcome *AuditEventOutcome `json:"outcome,omitempty"`
	// Authorization related to the event
	Authorization []CodeableConcept `json:"authorization,omitempty"`
	// Workflow authorization within which this event occurred
	BasedOn []custom.Reference[AuditEventBasedOn] `json:"basedOn,omitempty"`
	// The patient is the subject of the data used/created/updated/deleted during the activity
	Patient *custom.Reference[Patient] `json:"patient,omitempty"`
	// Encounter within which this event occurred or which the event is tightly associated
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Actor involved in the event
	Agent []AuditEventAgent `json:"agent"`
	// Audit Event Reporter
	Source AuditEventSource `json:"source"`
	// Data or objects used
	Entity []AuditEventEntity `json:"entity,omitempty"`
}

type AuditEventOutcome struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Whether the event succeeded or failed
	Code Coding `json:"code"`
	// Additional outcome detail
	Detail []CodeableConcept `json:"detail,omitempty"`
}

type AuditEventAgent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// How agent participated
	Type *CodeableConcept `json:"type,omitempty"`
	// Agent role in the event
	Role []CodeableConcept `json:"role,omitempty"`
	// Identifier of who
	Who custom.Reference[AuditEventAgentWho] `json:"who"`
	// Whether user is initiator
	Requestor *bool `json:"requestor,omitempty"`
	// The agent location when the event occurred
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Policy that authorized the agent participation in the event
	Policy []custom.URI `json:"policy,omitempty"`
	// This agent network location for the activity
	NetworkReference *custom.Reference[Endpoint] `json:"networkReference,omitempty"`
	// This agent network location for the activity
	NetworkUri *custom.URI `json:"networkUri,omitempty"`
	// This agent network location for the activity
	NetworkString *string `json:"networkString,omitempty"`
	// Allowable authorization for this agent
	Authorization []CodeableConcept `json:"authorization,omitempty"`
}

type AuditEventSource struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Logical source location within the enterprise
	Site *custom.Reference[Location] `json:"site,omitempty"`
	// The identity of source detecting the event
	Observer custom.Reference[AuditEventSourceObserver] `json:"observer"`
	// The type of source where event originated
	Type []CodeableConcept `json:"type,omitempty"`
}

type AuditEventEntity struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Specific instance of resource
	What *custom.Reference[Resource] `json:"what,omitempty"`
	// What role the entity played
	Role *CodeableConcept `json:"role,omitempty"`
	// Security labels on the entity
	SecurityLabel []CodeableConcept `json:"securityLabel,omitempty"`
	// Query parameters
	Query *custom.Base64binary `json:"query,omitempty"`
	// Additional Information about the entity
	Detail []AuditEventEntityDetail `json:"detail,omitempty"`
	// Entity is attributed to this agent
	Agent []interface{} `json:"agent,omitempty"`
}

type OtherAuditEvent AuditEvent

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
	return "AuditEvent"
}

func (a AuditEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAuditEvent
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAuditEvent: OtherAuditEvent(a), ResourceType: a.ResourceType()})
}

func UnmarshalAuditEvent(b []byte) (res AuditEvent, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
