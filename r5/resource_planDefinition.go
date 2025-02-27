// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/PlanDefinition
type PlanDefinitionActorOption struct {
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
	TypeReference *custom.Reference[PlanDefinitionActorOptionTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent
	Role *CodeableConcept `json:"role,omitempty"`
}

type PlanDefinitionActionInput struct {
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

type PlanDefinitionActionRelatedAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What action is this related to
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

type PlanDefinitionActionParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What actor
	ActorId *string `json:"actorId,omitempty"`
	// careteam | device | group | healthcareservice | location | organization | patient | practitioner | practitionerrole | relatedperson
	Type *custom.Code `json:"type,omitempty"`
	// Who or what can participate
	TypeCanonical *custom.Canonical[CapabilityStatement] `json:"typeCanonical,omitempty"`
	// Who or what can participate
	TypeReference *custom.Reference[PlanDefinitionActionParticipantTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent
	Role *CodeableConcept `json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `json:"function,omitempty"`
}

type PlanDefinitionActionDynamicValue struct {
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

type PlanDefinitionGoal struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// E.g. Treatment, dietary, behavioral
	Category *CodeableConcept `json:"category,omitempty"`
	// Code or text describing the goal
	Description CodeableConcept `json:"description"`
	// high-priority | medium-priority | low-priority
	Priority *CodeableConcept `json:"priority,omitempty"`
	// When goal pursuit begins
	Start *CodeableConcept `json:"start,omitempty"`
	// What does the goal address
	Addresses []CodeableConcept `json:"addresses,omitempty"`
	// Supporting documentation for the goal
	Documentation []RelatedArtifact `json:"documentation,omitempty"`
	// Target outcome for the goal
	Target []PlanDefinitionGoalTarget `json:"target,omitempty"`
}

type PlanDefinitionActor struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// User-visible title
	Title *string `json:"title,omitempty"`
	// Describes the actor
	Description *custom.Markdown `json:"description,omitempty"`
	// Who or what can be this actor
	Option []PlanDefinitionActorOption `json:"option"`
}

type PlanDefinitionAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique id for the action in the PlanDefinition
	LinkId *string `json:"linkId,omitempty"`
	// User-visible prefix for the action (e.g. 1. or A.)
	Prefix *string `json:"prefix,omitempty"`
	// User-visible title
	Title *string `json:"title,omitempty"`
	// Brief description of the action
	Description *custom.Markdown `json:"description,omitempty"`
	// Static text equivalent of the action, used if the dynamic aspects cannot be interpreted by the receiving system
	TextEquivalent *custom.Markdown `json:"textEquivalent,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// Code representing the meaning of the action or sub-actions
	Code *CodeableConcept `json:"code,omitempty"`
	// Why the action should be performed
	Reason []CodeableConcept `json:"reason,omitempty"`
	// Supporting documentation for the intended performer of the action
	Documentation []RelatedArtifact `json:"documentation,omitempty"`
	// What goals this action supports
	GoalId []custom.ID `json:"goalId,omitempty"`
	// Type of individual the action is focused on
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	// Type of individual the action is focused on
	SubjectReference *custom.Reference[Group] `json:"subjectReference,omitempty"`
	// Type of individual the action is focused on
	SubjectCanonical *custom.Canonical[any] `json:"subjectCanonical,omitempty"`
	// When the action should be triggered
	Trigger []TriggerDefinition `json:"trigger,omitempty"`
	// Whether or not the action is applicable
	Condition []PlanDefinitionActionCondition `json:"condition,omitempty"`
	// Input data requirements
	Input []PlanDefinitionActionInput `json:"input,omitempty"`
	// Output data definition
	Output []PlanDefinitionActionOutput `json:"output,omitempty"`
	// Relationship to another action
	RelatedAction []PlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
	// When the action should take place
	TimingAge *Age `json:"timingAge,omitempty"`
	// When the action should take place
	TimingDuration *Duration `json:"timingDuration,omitempty"`
	// When the action should take place
	TimingRange *Range `json:"timingRange,omitempty"`
	// When the action should take place
	TimingTiming *Timing `json:"timingTiming,omitempty"`
	// Where it should happen
	Location *custom.CodeableReference[Location] `json:"location,omitempty"`
	// Who should participate in the action
	Participant []PlanDefinitionActionParticipant `json:"participant,omitempty"`
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
	// Description of the activity to be performed
	DefinitionCanonical *custom.Canonical[PlanDefinitionActionDefinitionCanonical] `json:"definitionCanonical,omitempty"`
	// Description of the activity to be performed
	DefinitionUri *custom.URI `json:"definitionUri,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []PlanDefinitionActionDynamicValue `json:"dynamicValue,omitempty"`
	// A sub-action
	Action []interface{} `json:"action,omitempty"`
}

type PlanDefinitionActionCondition struct {
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

type PlanDefinitionActionOutput struct {
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

type PlanDefinition struct {
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
	// Canonical identifier for this plan definition, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the plan definition
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the plan definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this plan definition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this plan definition (human friendly)
	Title *string `json:"title,omitempty"`
	// Subordinate title of the plan definition
	Subtitle *string `json:"subtitle,omitempty"`
	// order-set | clinical-protocol | eca-rule | workflow-definition
	Type *CodeableConcept `json:"type,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Type of individual the plan definition is focused on
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	// Type of individual the plan definition is focused on
	SubjectReference *custom.Reference[PlanDefinitionSubjectReference] `json:"subjectReference,omitempty"`
	// Type of individual the plan definition is focused on
	SubjectCanonical *custom.Canonical[EvidenceVariable] `json:"subjectCanonical,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the plan definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for plan definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this plan definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Describes the clinical usage of the plan
	Usage *custom.Markdown `json:"usage,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the plan definition was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the plan definition was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the plan definition is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment
	Topic []CodeableConcept `json:"topic,omitempty"`
	// Who authored the content
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Additional documentation, citations
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Logic used by the plan definition
	Library []custom.Canonical[Library] `json:"library,omitempty"`
	// What the plan is trying to accomplish
	Goal []PlanDefinitionGoal `json:"goal,omitempty"`
	// Actors within the plan
	Actor []PlanDefinitionActor `json:"actor,omitempty"`
	// Action defined by the plan
	Action []PlanDefinitionAction `json:"action,omitempty"`
	// Preconditions for service
	AsNeededBoolean *bool `json:"asNeededBoolean,omitempty"`
	// Preconditions for service
	AsNeededCodeableConcept *CodeableConcept `json:"asNeededCodeableConcept,omitempty"`
}

type PlanDefinitionGoalTarget struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The parameter whose value is to be tracked
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
	// Reach goal within
	Due *Duration `json:"due,omitempty"`
}

type OtherPlanDefinition PlanDefinition

type PlanDefinitionSubjectReference interface {
	gofhirclient.Resource

	Is_PlanDefinitionSubjectReference()
}

func (g Group) Is_PlanDefinitionSubjectReference()                          {}
func (m MedicinalProductDefinition) Is_PlanDefinitionSubjectReference()     {}
func (s SubstanceDefinition) Is_PlanDefinitionSubjectReference()            {}
func (a AdministrableProductDefinition) Is_PlanDefinitionSubjectReference() {}
func (m ManufacturedItemDefinition) Is_PlanDefinitionSubjectReference()     {}
func (p PackagedProductDefinition) Is_PlanDefinitionSubjectReference()      {}

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

type PlanDefinitionActionDefinitionCanonical interface {
	gofhirclient.Resource

	Is_PlanDefinitionActionDefinitionCanonical()
}

func (a ActivityDefinition) Is_PlanDefinitionActionDefinitionCanonical()    {}
func (m MessageDefinition) Is_PlanDefinitionActionDefinitionCanonical()     {}
func (o ObservationDefinition) Is_PlanDefinitionActionDefinitionCanonical() {}
func (p PlanDefinition) Is_PlanDefinitionActionDefinitionCanonical()        {}
func (q Questionnaire) Is_PlanDefinitionActionDefinitionCanonical()         {}
func (s SpecimenDefinition) Is_PlanDefinitionActionDefinitionCanonical()    {}

func (p PlanDefinition) ResourceType() string {
	return "PlanDefinition"
}

func (p PlanDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPlanDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPlanDefinition: OtherPlanDefinition(p), ResourceType: p.ResourceType()})
}

func UnmarshalPlanDefinition(b []byte) (res PlanDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
