// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/RequestOrchestration
type RequestOrchestrationActionDynamicValue struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The path to the element to be set dynamically
	Path *string `bson:"path,omitempty" json:"path,omitempty"`
	// An expression that provides the dynamic value for the customization
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty"`
}

type RequestOrchestration struct {
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
	// Business identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Fulfills plan, proposal, or order
	BasedOn []custom.Reference[Resource] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Request(s) replaced by this request
	Replaces []custom.Reference[Resource] `bson:"replaces,omitempty" json:"replaces,omitempty"`
	// Composite request this is part of
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `bson:"intent" json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// What's being requested/ordered
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Who the request orchestration is about
	Subject *custom.Reference[RequestOrchestrationSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Created as part of
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the request orchestration was authored
	AuthoredOn *custom.DateTime `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	// Device or practitioner that authored the request orchestration
	Author *custom.Reference[RequestOrchestrationAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Why the request orchestration is needed
	Reason []custom.CodeableReference[RequestOrchestrationReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// What goals
	Goal []custom.Reference[Goal] `bson:"goal,omitempty" json:"goal,omitempty"`
	// Additional notes about the response
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Proposed actions, if any
	Action []RequestOrchestrationAction `bson:"action,omitempty" json:"action,omitempty"`
}

type RequestOrchestrationAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Pointer to specific item from the PlanDefinition
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// User-visible prefix for the action (e.g. 1. or A.)
	Prefix *string `bson:"prefix,omitempty" json:"prefix,omitempty"`
	// User-visible title
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Short description of the action
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Static text equivalent of the action, used if the dynamic aspects cannot be interpreted by the receiving system
	TextEquivalent *custom.Markdown `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// Code representing the meaning of the action or sub-actions
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Supporting documentation for the intended performer of the action
	Documentation []RelatedArtifact `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// What goals
	Goal []custom.Reference[Goal] `bson:"goal,omitempty" json:"goal,omitempty"`
	// Whether or not the action is applicable
	Condition []RequestOrchestrationActionCondition `bson:"condition,omitempty" json:"condition,omitempty"`
	// Input data requirements
	Input []RequestOrchestrationActionInput `bson:"input,omitempty" json:"input,omitempty"`
	// Output data definition
	Output []RequestOrchestrationActionOutput `bson:"output,omitempty" json:"output,omitempty"`
	// Relationship to another action
	RelatedAction []RequestOrchestrationActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	// When the action should take place
	Timing *custom.DateTime `bson:"timing,omitempty" json:"timing,omitempty"`
	// Where it should happen
	Location *custom.CodeableReference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Who should perform the action
	Participant []RequestOrchestrationActionParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// create | update | remove | fire-event
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// visual-group | logical-group | sentence-group
	GroupingBehavior *custom.Code `bson:"groupingBehavior,omitempty" json:"groupingBehavior,omitempty"`
	// any | all | all-or-none | exactly-one | at-most-one | one-or-more
	SelectionBehavior *custom.Code `bson:"selectionBehavior,omitempty" json:"selectionBehavior,omitempty"`
	// must | could | must-unless-documented
	RequiredBehavior *custom.Code `bson:"requiredBehavior,omitempty" json:"requiredBehavior,omitempty"`
	// yes | no
	PrecheckBehavior *custom.Code `bson:"precheckBehavior,omitempty" json:"precheckBehavior,omitempty"`
	// single | multiple
	CardinalityBehavior *custom.Code `bson:"cardinalityBehavior,omitempty" json:"cardinalityBehavior,omitempty"`
	// The target of the action
	Resource *custom.Reference[Resource] `bson:"resource,omitempty" json:"resource,omitempty"`
	// Description of the activity to be performed
	Definition *custom.Canonical[RequestOrchestrationActionDefinition] `bson:"definition,omitempty" json:"definition,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `bson:"transform,omitempty" json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []RequestOrchestrationActionDynamicValue `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
	// Sub action
	Action []interface{} `bson:"action,omitempty" json:"action,omitempty"`
}

type RequestOrchestrationActionCondition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// applicability | start | stop
	Kind custom.Code `bson:"kind" json:"kind"`
	// Boolean-valued expression
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty"`
}

type RequestOrchestrationActionInput struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// User-visible title
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// What data is provided
	Requirement *DataRequirement `bson:"requirement,omitempty" json:"requirement,omitempty"`
	// What data is provided
	RelatedData *custom.ID `bson:"relatedData,omitempty" json:"relatedData,omitempty"`
}

type RequestOrchestrationActionOutput struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// User-visible title
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// What data is provided
	Requirement *DataRequirement `bson:"requirement,omitempty" json:"requirement,omitempty"`
	// What data is provided
	RelatedData *string `bson:"relatedData,omitempty" json:"relatedData,omitempty"`
}

type RequestOrchestrationActionRelatedAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What action this is related to
	TargetId custom.ID `bson:"targetId" json:"targetId"`
	// before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	Relationship custom.Code `bson:"relationship" json:"relationship"`
	// before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	EndRelationship *custom.Code `bson:"endRelationship,omitempty" json:"endRelationship,omitempty"`
	// Time offset for the relationship
	Offset *Duration `bson:"offset,omitempty" json:"offset,omitempty"`
}

type RequestOrchestrationActionParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// Who or what can participate
	TypeCanonical *custom.Canonical[CapabilityStatement] `bson:"typeCanonical,omitempty" json:"typeCanonical,omitempty"`
	// Who or what can participate
	TypeReference *custom.Reference[RequestOrchestrationActionParticipantTypeReference] `bson:"typeReference,omitempty" json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent, etc
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who/what is participating?
	Actor *custom.Canonical[CapabilityStatement] `bson:"actor,omitempty" json:"actor,omitempty"`
}

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

type RequestOrchestrationActionDefinition interface {
	gofhirclient.Resource

	Is_RequestOrchestrationActionDefinition()
}

func (a ActivityDefinition) Is_RequestOrchestrationActionDefinition()    {}
func (o ObservationDefinition) Is_RequestOrchestrationActionDefinition() {}
func (p PlanDefinition) Is_RequestOrchestrationActionDefinition()        {}
func (q Questionnaire) Is_RequestOrchestrationActionDefinition()         {}
func (s SpecimenDefinition) Is_RequestOrchestrationActionDefinition()    {}

func (r RequestOrchestration) ResourceType() string {
	return "StructureDefinition"
}
