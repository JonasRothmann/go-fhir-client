// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/RequestOrchestration
type RequestOrchestrationActionDynamicValue struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The path to the element to be set dynamically
	Path *string `json:"path,omitempty"`
	// An expression that provides the dynamic value for the customization
	Expression *Expression `json:"expression,omitempty"`
}

type RequestOrchestration struct {
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
	// Business identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[any] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `json:"instantiatesUri,omitempty"`
	// Fulfills plan, proposal, or order
	BasedOn []custom.Reference[Resource] `json:"basedOn,omitempty"`
	// Request(s) replaced by this request
	Replaces []custom.Reference[Resource] `json:"replaces,omitempty"`
	// Composite request this is part of
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// What's being requested/ordered
	Code *CodeableConcept `json:"code,omitempty"`
	// Who the request orchestration is about
	Subject *custom.Reference[RequestOrchestrationSubject] `json:"subject,omitempty"`
	// Created as part of
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When the request orchestration was authored
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Device or practitioner that authored the request orchestration
	Author *custom.Reference[RequestOrchestrationAuthor] `json:"author,omitempty"`
	// Why the request orchestration is needed
	Reason []custom.CodeableReference[RequestOrchestrationReason] `json:"reason,omitempty"`
	// What goals
	Goal []custom.Reference[Goal] `json:"goal,omitempty"`
	// Additional notes about the response
	Note []Annotation `json:"note,omitempty"`
	// Proposed actions, if any
	Action []RequestOrchestrationAction `json:"action,omitempty"`
}

type RequestOrchestrationAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Pointer to specific item from the PlanDefinition
	LinkId *string `json:"linkId,omitempty"`
	// User-visible prefix for the action (e.g. 1. or A.)
	Prefix *string `json:"prefix,omitempty"`
	// User-visible title
	Title *string `json:"title,omitempty"`
	// Short description of the action
	Description *custom.Markdown `json:"description,omitempty"`
	// Static text equivalent of the action, used if the dynamic aspects cannot be interpreted by the receiving system
	TextEquivalent *custom.Markdown `json:"textEquivalent,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// Code representing the meaning of the action or sub-actions
	Code []CodeableConcept `json:"code,omitempty"`
	// Supporting documentation for the intended performer of the action
	Documentation []RelatedArtifact `json:"documentation,omitempty"`
	// What goals
	Goal []custom.Reference[Goal] `json:"goal,omitempty"`
	// Whether or not the action is applicable
	Condition []RequestOrchestrationActionCondition `json:"condition,omitempty"`
	// Input data requirements
	Input []RequestOrchestrationActionInput `json:"input,omitempty"`
	// Output data definition
	Output []RequestOrchestrationActionOutput `json:"output,omitempty"`
	// Relationship to another action
	RelatedAction []RequestOrchestrationActionRelatedAction `json:"relatedAction,omitempty"`
	// When the action should take place
	TimingDateTime *custom.DateTime `json:"timingDateTime,omitempty"`
	// When the action should take place
	TimingAge *Age `json:"timingAge,omitempty"`
	// When the action should take place
	TimingPeriod *Period `json:"timingPeriod,omitempty"`
	// When the action should take place
	TimingDuration *Duration `json:"timingDuration,omitempty"`
	// When the action should take place
	TimingRange *Range `json:"timingRange,omitempty"`
	// When the action should take place
	TimingTiming *Timing `json:"timingTiming,omitempty"`
	// Where it should happen
	Location *custom.CodeableReference[Location] `json:"location,omitempty"`
	// Who should perform the action
	Participant []RequestOrchestrationActionParticipant `json:"participant,omitempty"`
	// create | update | remove | fire-event
	Type *CodeableConcept `json:"type,omitempty"`
	// visual-group | logical-group | sentence-group
	GroupingBehavior *custom.Code `json:"groupingBehavior,omitempty"`
	// any | all | all-or-none | exactly-one | at-most-one | one-or-more
	SelectionBehavior *custom.Code `json:"selectionBehavior,omitempty"`
	// must | could | must-unless-documented
	RequiredBehavior *custom.Code `json:"requiredBehavior,omitempty"`
	// yes | no
	PrecheckBehavior *custom.Code `json:"precheckBehavior,omitempty"`
	// single | multiple
	CardinalityBehavior *custom.Code `json:"cardinalityBehavior,omitempty"`
	// The target of the action
	Resource *custom.Reference[Resource] `json:"resource,omitempty"`
	// Description of the activity to be performed
	DefinitionCanonical *custom.Canonical[RequestOrchestrationActionDefinitionCanonical] `json:"definitionCanonical,omitempty"`
	// Description of the activity to be performed
	DefinitionUri *custom.URI `json:"definitionUri,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []RequestOrchestrationActionDynamicValue `json:"dynamicValue,omitempty"`
	// Sub action
	Action []interface{} `json:"action,omitempty"`
}

type RequestOrchestrationActionCondition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// applicability | start | stop
	Kind custom.Code `json:"kind"`
	// Boolean-valued expression
	Expression *Expression `json:"expression,omitempty"`
}

type RequestOrchestrationActionInput struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// User-visible title
	Title *string `json:"title,omitempty"`
	// What data is provided
	Requirement *DataRequirement `json:"requirement,omitempty"`
	// What data is provided
	RelatedData *custom.ID `json:"relatedData,omitempty"`
}

type RequestOrchestrationActionOutput struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// User-visible title
	Title *string `json:"title,omitempty"`
	// What data is provided
	Requirement *DataRequirement `json:"requirement,omitempty"`
	// What data is provided
	RelatedData *string `json:"relatedData,omitempty"`
}

type RequestOrchestrationActionRelatedAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What action this is related to
	TargetId custom.ID `json:"targetId"`
	// before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	Relationship custom.Code `json:"relationship"`
	// before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	EndRelationship *custom.Code `json:"endRelationship,omitempty"`
	// Time offset for the relationship
	OffsetDuration *Duration `json:"offsetDuration,omitempty"`
	// Time offset for the relationship
	OffsetRange *Range `json:"offsetRange,omitempty"`
}

type RequestOrchestrationActionParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	Type *custom.Code `json:"type,omitempty"`
	// Who or what can participate
	TypeCanonical *custom.Canonical[CapabilityStatement] `json:"typeCanonical,omitempty"`
	// Who or what can participate
	TypeReference *custom.Reference[RequestOrchestrationActionParticipantTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent, etc
	Role *CodeableConcept `json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `json:"function,omitempty"`
	// Who/what is participating?
	ActorCanonical *custom.Canonical[CapabilityStatement] `json:"actorCanonical,omitempty"`
	// Who/what is participating?
	ActorReference *custom.Reference[RequestOrchestrationActionParticipantActorReference] `json:"actorReference,omitempty"`
}

type OtherRequestOrchestration RequestOrchestration

type RequestOrchestrationSubject interface {
	gofhirclient.Resource

	Is_RequestOrchestrationSubject()
}

func (c CareTeam) Is_RequestOrchestrationSubject()          {}
func (d Device) Is_RequestOrchestrationSubject()            {}
func (g Group) Is_RequestOrchestrationSubject()             {}
func (h HealthcareService) Is_RequestOrchestrationSubject() {}
func (l Location) Is_RequestOrchestrationSubject()          {}
func (o Organization) Is_RequestOrchestrationSubject()      {}
func (p Patient) Is_RequestOrchestrationSubject()           {}
func (p Practitioner) Is_RequestOrchestrationSubject()      {}
func (p PractitionerRole) Is_RequestOrchestrationSubject()  {}
func (r RelatedPerson) Is_RequestOrchestrationSubject()     {}

type RequestOrchestrationAuthor interface {
	gofhirclient.Resource

	Is_RequestOrchestrationAuthor()
}

func (d Device) Is_RequestOrchestrationAuthor()           {}
func (p Practitioner) Is_RequestOrchestrationAuthor()     {}
func (p PractitionerRole) Is_RequestOrchestrationAuthor() {}

type RequestOrchestrationReason interface {
	gofhirclient.Resource

	Is_RequestOrchestrationReason()
}

func (c Condition) Is_RequestOrchestrationReason()         {}
func (o Observation) Is_RequestOrchestrationReason()       {}
func (d DiagnosticReport) Is_RequestOrchestrationReason()  {}
func (d DocumentReference) Is_RequestOrchestrationReason() {}

type RequestOrchestrationActionParticipantTypeReference interface {
	gofhirclient.Resource

	Is_RequestOrchestrationActionParticipantTypeReference()
}

func (c CareTeam) Is_RequestOrchestrationActionParticipantTypeReference()          {}
func (d Device) Is_RequestOrchestrationActionParticipantTypeReference()            {}
func (d DeviceDefinition) Is_RequestOrchestrationActionParticipantTypeReference()  {}
func (e Endpoint) Is_RequestOrchestrationActionParticipantTypeReference()          {}
func (g Group) Is_RequestOrchestrationActionParticipantTypeReference()             {}
func (h HealthcareService) Is_RequestOrchestrationActionParticipantTypeReference() {}
func (l Location) Is_RequestOrchestrationActionParticipantTypeReference()          {}
func (o Organization) Is_RequestOrchestrationActionParticipantTypeReference()      {}
func (p Patient) Is_RequestOrchestrationActionParticipantTypeReference()           {}
func (p Practitioner) Is_RequestOrchestrationActionParticipantTypeReference()      {}
func (p PractitionerRole) Is_RequestOrchestrationActionParticipantTypeReference()  {}
func (r RelatedPerson) Is_RequestOrchestrationActionParticipantTypeReference()     {}

type RequestOrchestrationActionParticipantActorReference interface {
	gofhirclient.Resource

	Is_RequestOrchestrationActionParticipantActorReference()
}

func (c CareTeam) Is_RequestOrchestrationActionParticipantActorReference()          {}
func (d Device) Is_RequestOrchestrationActionParticipantActorReference()            {}
func (d DeviceDefinition) Is_RequestOrchestrationActionParticipantActorReference()  {}
func (e Endpoint) Is_RequestOrchestrationActionParticipantActorReference()          {}
func (g Group) Is_RequestOrchestrationActionParticipantActorReference()             {}
func (h HealthcareService) Is_RequestOrchestrationActionParticipantActorReference() {}
func (l Location) Is_RequestOrchestrationActionParticipantActorReference()          {}
func (o Organization) Is_RequestOrchestrationActionParticipantActorReference()      {}
func (p Patient) Is_RequestOrchestrationActionParticipantActorReference()           {}
func (p Practitioner) Is_RequestOrchestrationActionParticipantActorReference()      {}
func (p PractitionerRole) Is_RequestOrchestrationActionParticipantActorReference()  {}
func (r RelatedPerson) Is_RequestOrchestrationActionParticipantActorReference()     {}

type RequestOrchestrationActionDefinitionCanonical interface {
	gofhirclient.Resource

	Is_RequestOrchestrationActionDefinitionCanonical()
}

func (a ActivityDefinition) Is_RequestOrchestrationActionDefinitionCanonical()    {}
func (o ObservationDefinition) Is_RequestOrchestrationActionDefinitionCanonical() {}
func (p PlanDefinition) Is_RequestOrchestrationActionDefinitionCanonical()        {}
func (q Questionnaire) Is_RequestOrchestrationActionDefinitionCanonical()         {}
func (s SpecimenDefinition) Is_RequestOrchestrationActionDefinitionCanonical()    {}

func (r RequestOrchestration) ResourceType() string {
	return "RequestOrchestration"
}

func (r RequestOrchestration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRequestOrchestration
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherRequestOrchestration: OtherRequestOrchestration(r), ResourceType: r.ResourceType()})
}

func UnmarshalRequestOrchestration(b []byte) (res RequestOrchestration, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
