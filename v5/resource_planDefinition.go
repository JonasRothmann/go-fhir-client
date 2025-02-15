// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/PlanDefinition
type PlanDefinition struct {
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
	// Canonical identifier for this plan definition, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the plan definition
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the plan definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this plan definition (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this plan definition (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Subordinate title of the plan definition
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	// order-set | clinical-protocol | eca-rule | workflow-definition
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Type of individual the plan definition is focused on
	Subject *CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the plan definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for plan definition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this plan definition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Describes the clinical usage of the plan
	Usage *custom.Markdown `bson:"usage,omitempty" json:"usage,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the plan definition was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the plan definition was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the plan definition is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Who authored the content
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Logic used by the plan definition
	Library []custom.Canonical[Library] `bson:"library,omitempty" json:"library,omitempty"`
	// What the plan is trying to accomplish
	Goal []PlanDefinitionGoal `bson:"goal,omitempty" json:"goal,omitempty"`
	// Actors within the plan
	Actor []PlanDefinitionActor `bson:"actor,omitempty" json:"actor,omitempty"`
	// Action defined by the plan
	Action []PlanDefinitionAction `bson:"action,omitempty" json:"action,omitempty"`
	// Preconditions for service
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
}

type PlanDefinitionGoal struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// E.g. Treatment, dietary, behavioral
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Code or text describing the goal
	Description CodeableConcept `bson:"description" json:"description"`
	// high-priority | medium-priority | low-priority
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty"`
	// When goal pursuit begins
	Start *CodeableConcept `bson:"start,omitempty" json:"start,omitempty"`
	// What does the goal address
	Addresses []CodeableConcept `bson:"addresses,omitempty" json:"addresses,omitempty"`
	// Supporting documentation for the goal
	Documentation []RelatedArtifact `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// Target outcome for the goal
	Target []PlanDefinitionGoalTarget `bson:"target,omitempty" json:"target,omitempty"`
}

type PlanDefinitionGoalTarget struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The parameter whose value is to be tracked
	Measure *CodeableConcept `bson:"measure,omitempty" json:"measure,omitempty"`
	// The target value to be achieved
	Detail *Quantity `bson:"detail,omitempty" json:"detail,omitempty"`
	// Reach goal within
	Due *Duration `bson:"due,omitempty" json:"due,omitempty"`
}

type PlanDefinitionActor struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// User-visible title
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Describes the actor
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Who or what can be this actor
	Option []PlanDefinitionActorOption `bson:"option" json:"option"`
}

type PlanDefinitionActionCondition struct {
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

type PlanDefinitionActionDynamicValue struct {
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

type PlanDefinitionActorOption struct {
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
	TypeReference *custom.Reference[PlanDefinitionActorOptionTypeReference] `bson:"typeReference,omitempty" json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type PlanDefinitionAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique id for the action in the PlanDefinition
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// User-visible prefix for the action (e.g. 1. or A.)
	Prefix *string `bson:"prefix,omitempty" json:"prefix,omitempty"`
	// User-visible title
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Brief description of the action
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Static text equivalent of the action, used if the dynamic aspects cannot be interpreted by the receiving system
	TextEquivalent *custom.Markdown `bson:"textEquivalent,omitempty" json:"textEquivalent,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// Code representing the meaning of the action or sub-actions
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Why the action should be performed
	Reason []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	// Supporting documentation for the intended performer of the action
	Documentation []RelatedArtifact `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// What goals this action supports
	GoalId []custom.ID `bson:"goalId,omitempty" json:"goalId,omitempty"`
	// Type of individual the action is focused on
	Subject *CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// When the action should be triggered
	Trigger []TriggerDefinition `bson:"trigger,omitempty" json:"trigger,omitempty"`
	// Whether or not the action is applicable
	Condition []PlanDefinitionActionCondition `bson:"condition,omitempty" json:"condition,omitempty"`
	// Input data requirements
	Input []PlanDefinitionActionInput `bson:"input,omitempty" json:"input,omitempty"`
	// Output data definition
	Output []PlanDefinitionActionOutput `bson:"output,omitempty" json:"output,omitempty"`
	// Relationship to another action
	RelatedAction []PlanDefinitionActionRelatedAction `bson:"relatedAction,omitempty" json:"relatedAction,omitempty"`
	// When the action should take place
	Timing *Age `bson:"timing,omitempty" json:"timing,omitempty"`
	// Where it should happen
	Location *custom.CodeableReference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Who should participate in the action
	Participant []PlanDefinitionActionParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
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
	// Description of the activity to be performed
	Definition *custom.Canonical[PlanDefinitionActionDefinition] `bson:"definition,omitempty" json:"definition,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `bson:"transform,omitempty" json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []PlanDefinitionActionDynamicValue `bson:"dynamicValue,omitempty" json:"dynamicValue,omitempty"`
	// A sub-action
	Action []interface{} `bson:"action,omitempty" json:"action,omitempty"`
}

type PlanDefinitionActionInput struct {
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

type PlanDefinitionActionOutput struct {
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

type PlanDefinitionActionRelatedAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What action is this related to
	TargetId custom.ID `bson:"targetId" json:"targetId"`
	// before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	Relationship custom.Code `bson:"relationship" json:"relationship"`
	// before | before-start | before-end | concurrent | concurrent-with-start | concurrent-with-end | after | after-start | after-end
	EndRelationship *custom.Code `bson:"endRelationship,omitempty" json:"endRelationship,omitempty"`
	// Time offset for the relationship
	Offset *Duration `bson:"offset,omitempty" json:"offset,omitempty"`
}

type PlanDefinitionActionParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What actor
	ActorId *string `bson:"actorId,omitempty" json:"actorId,omitempty"`
	// careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// Who or what can participate
	TypeCanonical *custom.Canonical[CapabilityStatement] `bson:"typeCanonical,omitempty" json:"typeCanonical,omitempty"`
	// Who or what can participate
	TypeReference *custom.Reference[PlanDefinitionActionParticipantTypeReference] `bson:"typeReference,omitempty" json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
}

type PlanDefinitionActorOptionTypeReference interface {
	gofhirclient.Resource

	Is_PlanDefinitionActorOptionTypeReference()
}

func (c CareTeam) Is_PlanDefinitionActorOptionTypeReference()          {}
func (d Device) Is_PlanDefinitionActorOptionTypeReference()            {}
func (d DeviceDefinition) Is_PlanDefinitionActorOptionTypeReference()  {}
func (e Endpoint) Is_PlanDefinitionActorOptionTypeReference()          {}
func (g Group) Is_PlanDefinitionActorOptionTypeReference()             {}
func (h HealthcareService) Is_PlanDefinitionActorOptionTypeReference() {}
func (l Location) Is_PlanDefinitionActorOptionTypeReference()          {}
func (o Organization) Is_PlanDefinitionActorOptionTypeReference()      {}
func (p Patient) Is_PlanDefinitionActorOptionTypeReference()           {}
func (p Practitioner) Is_PlanDefinitionActorOptionTypeReference()      {}
func (p PractitionerRole) Is_PlanDefinitionActorOptionTypeReference()  {}
func (r RelatedPerson) Is_PlanDefinitionActorOptionTypeReference()     {}

type PlanDefinitionActionParticipantTypeReference interface {
	gofhirclient.Resource

	Is_PlanDefinitionActionParticipantTypeReference()
}

func (c CareTeam) Is_PlanDefinitionActionParticipantTypeReference()          {}
func (d Device) Is_PlanDefinitionActionParticipantTypeReference()            {}
func (d DeviceDefinition) Is_PlanDefinitionActionParticipantTypeReference()  {}
func (e Endpoint) Is_PlanDefinitionActionParticipantTypeReference()          {}
func (g Group) Is_PlanDefinitionActionParticipantTypeReference()             {}
func (h HealthcareService) Is_PlanDefinitionActionParticipantTypeReference() {}
func (l Location) Is_PlanDefinitionActionParticipantTypeReference()          {}
func (o Organization) Is_PlanDefinitionActionParticipantTypeReference()      {}
func (p Patient) Is_PlanDefinitionActionParticipantTypeReference()           {}
func (p Practitioner) Is_PlanDefinitionActionParticipantTypeReference()      {}
func (p PractitionerRole) Is_PlanDefinitionActionParticipantTypeReference()  {}
func (r RelatedPerson) Is_PlanDefinitionActionParticipantTypeReference()     {}

type PlanDefinitionActionDefinition interface {
	gofhirclient.Resource

	Is_PlanDefinitionActionDefinition()
}

func (a ActivityDefinition) Is_PlanDefinitionActionDefinition()    {}
func (m MessageDefinition) Is_PlanDefinitionActionDefinition()     {}
func (o ObservationDefinition) Is_PlanDefinitionActionDefinition() {}
func (p PlanDefinition) Is_PlanDefinitionActionDefinition()        {}
func (q Questionnaire) Is_PlanDefinitionActionDefinition()         {}
func (s SpecimenDefinition) Is_PlanDefinitionActionDefinition()    {}

func (p PlanDefinition) ResourceType() string {
	return "StructureDefinition"
}
