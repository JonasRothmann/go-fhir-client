// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/cdshooksrequestorchestration
type CdsHooksRequestOrchestrationActionCondition struct {
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

type CdsHooksRequestOrchestrationActionInput struct {
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

type CdsHooksRequestOrchestrationActionOutput struct {
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

type CdsHooksRequestOrchestrationActionRelatedAction struct {
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

type CdsHooksRequestOrchestrationActionParticipant struct {
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
	TypeReference *custom.Reference[CdsHooksRequestOrchestrationActionParticipantTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent, etc
	Role *CodeableConcept `json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `json:"function,omitempty"`
	// Who/what is participating?
	ActorCanonical *custom.Canonical[CapabilityStatement] `json:"actorCanonical,omitempty"`
	// Who/what is participating?
	ActorReference *custom.Reference[CdsHooksRequestOrchestrationActionParticipantActorReference] `json:"actorReference,omitempty"`
}

type CdsHooksRequestOrchestrationActionDynamicValue struct {
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

type CdsHooksRequestOrchestration struct {
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
	Identifier Identifier `json:"identifier"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[any] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri custom.URI `json:"instantiatesUri"`
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
	Subject *custom.Reference[CdsHooksRequestOrchestrationSubject] `json:"subject,omitempty"`
	// Created as part of
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When the request orchestration was authored
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Device or practitioner that authored the request orchestration
	Author *custom.Reference[CdsHooksRequestOrchestrationAuthor] `json:"author,omitempty"`
	// Why the request orchestration is needed
	Reason []custom.CodeableReference[CdsHooksRequestOrchestrationReason] `json:"reason,omitempty"`
	// What goals
	Goal []custom.Reference[Goal] `json:"goal,omitempty"`
	// Additional notes about the response
	Note []Annotation `json:"note,omitempty"`
	// Proposed actions, if any
	Action []CdsHooksRequestOrchestrationAction `json:"action,omitempty"`
}

type CdsHooksRequestOrchestrationAction struct {
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
	Condition []CdsHooksRequestOrchestrationActionCondition `json:"condition,omitempty"`
	// Input data requirements
	Input []CdsHooksRequestOrchestrationActionInput `json:"input,omitempty"`
	// Output data definition
	Output []CdsHooksRequestOrchestrationActionOutput `json:"output,omitempty"`
	// Relationship to another action
	RelatedAction []CdsHooksRequestOrchestrationActionRelatedAction `json:"relatedAction,omitempty"`
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
	Participant []CdsHooksRequestOrchestrationActionParticipant `json:"participant,omitempty"`
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
	DefinitionCanonical *custom.Canonical[CdsHooksRequestOrchestrationActionDefinitionCanonical] `json:"definitionCanonical,omitempty"`
	// Description of the activity to be performed
	DefinitionUri *custom.URI `json:"definitionUri,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []CdsHooksRequestOrchestrationActionDynamicValue `json:"dynamicValue,omitempty"`
	// Sub action
	Action []interface{} `json:"action,omitempty"`
}

type OtherCdsHooksRequestOrchestration CdsHooksRequestOrchestration

type CdsHooksRequestOrchestrationSubject interface {
	gofhirclient.Resource

	Is_CdsHooksRequestOrchestrationSubject()
}

func (c CareTeam) Is_CdsHooksRequestOrchestrationSubject()          {}
func (d Device) Is_CdsHooksRequestOrchestrationSubject()            {}
func (g Group) Is_CdsHooksRequestOrchestrationSubject()             {}
func (h HealthcareService) Is_CdsHooksRequestOrchestrationSubject() {}
func (l Location) Is_CdsHooksRequestOrchestrationSubject()          {}
func (o Organization) Is_CdsHooksRequestOrchestrationSubject()      {}
func (p Patient) Is_CdsHooksRequestOrchestrationSubject()           {}
func (p Practitioner) Is_CdsHooksRequestOrchestrationSubject()      {}
func (p PractitionerRole) Is_CdsHooksRequestOrchestrationSubject()  {}
func (r RelatedPerson) Is_CdsHooksRequestOrchestrationSubject()     {}

type CdsHooksRequestOrchestrationAuthor interface {
	gofhirclient.Resource

	Is_CdsHooksRequestOrchestrationAuthor()
}

func (d Device) Is_CdsHooksRequestOrchestrationAuthor()           {}
func (p Practitioner) Is_CdsHooksRequestOrchestrationAuthor()     {}
func (p PractitionerRole) Is_CdsHooksRequestOrchestrationAuthor() {}

type CdsHooksRequestOrchestrationReason interface {
	gofhirclient.Resource

	Is_CdsHooksRequestOrchestrationReason()
}

func (c Condition) Is_CdsHooksRequestOrchestrationReason()         {}
func (o Observation) Is_CdsHooksRequestOrchestrationReason()       {}
func (d DiagnosticReport) Is_CdsHooksRequestOrchestrationReason()  {}
func (d DocumentReference) Is_CdsHooksRequestOrchestrationReason() {}

type CdsHooksRequestOrchestrationActionParticipantTypeReference interface {
	gofhirclient.Resource

	Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()
}

func (c CareTeam) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()          {}
func (d Device) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()            {}
func (d DeviceDefinition) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()  {}
func (e Endpoint) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()          {}
func (g Group) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()             {}
func (h HealthcareService) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference() {}
func (l Location) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()          {}
func (o Organization) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()      {}
func (p Patient) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()           {}
func (p Practitioner) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()      {}
func (p PractitionerRole) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()  {}
func (r RelatedPerson) Is_CdsHooksRequestOrchestrationActionParticipantTypeReference()     {}

type CdsHooksRequestOrchestrationActionParticipantActorReference interface {
	gofhirclient.Resource

	Is_CdsHooksRequestOrchestrationActionParticipantActorReference()
}

func (c CareTeam) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()          {}
func (d Device) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()            {}
func (d DeviceDefinition) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()  {}
func (e Endpoint) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()          {}
func (g Group) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()             {}
func (h HealthcareService) Is_CdsHooksRequestOrchestrationActionParticipantActorReference() {}
func (l Location) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()          {}
func (o Organization) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()      {}
func (p Patient) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()           {}
func (p Practitioner) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()      {}
func (p PractitionerRole) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()  {}
func (r RelatedPerson) Is_CdsHooksRequestOrchestrationActionParticipantActorReference()     {}

type CdsHooksRequestOrchestrationActionDefinitionCanonical interface {
	gofhirclient.Resource

	Is_CdsHooksRequestOrchestrationActionDefinitionCanonical()
}

func (a ActivityDefinition) Is_CdsHooksRequestOrchestrationActionDefinitionCanonical()    {}
func (o ObservationDefinition) Is_CdsHooksRequestOrchestrationActionDefinitionCanonical() {}
func (p PlanDefinition) Is_CdsHooksRequestOrchestrationActionDefinitionCanonical()        {}
func (q Questionnaire) Is_CdsHooksRequestOrchestrationActionDefinitionCanonical()         {}
func (s SpecimenDefinition) Is_CdsHooksRequestOrchestrationActionDefinitionCanonical()    {}

func (c CdsHooksRequestOrchestration) ResourceType() string {
	return "CDSHooksRequestOrchestration"
}

func (c CdsHooksRequestOrchestration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCdsHooksRequestOrchestration
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCdsHooksRequestOrchestration: OtherCdsHooksRequestOrchestration(c), ResourceType: c.ResourceType()})
}

func UnmarshalCdsHooksRequestOrchestration(b []byte) (res CdsHooksRequestOrchestration, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
