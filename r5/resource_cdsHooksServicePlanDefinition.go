// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/cdshooksserviceplandefinition
type CdsHooksServicePlanDefinitionActionOutput struct {
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

type CdsHooksServicePlanDefinitionActionRelatedAction struct {
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

type CdsHooksServicePlanDefinitionActionParticipant struct {
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
	TypeReference *custom.Reference[CdsHooksServicePlanDefinitionActionParticipantTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent
	Role *CodeableConcept `json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `json:"function,omitempty"`
}

type CdsHooksServicePlanDefinitionActionDynamicValue struct {
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

type CdsHooksServicePlanDefinitionActor struct {
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
	Option []CdsHooksServicePlanDefinitionActorOption `json:"option"`
}

type CdsHooksServicePlanDefinitionGoal struct {
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
	Target []CdsHooksServicePlanDefinitionGoalTarget `json:"target,omitempty"`
}

type CdsHooksServicePlanDefinitionGoalTarget struct {
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

type CdsHooksServicePlanDefinitionActorOption struct {
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
	TypeReference *custom.Reference[CdsHooksServicePlanDefinitionActorOptionTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent
	Role *CodeableConcept `json:"role,omitempty"`
}

type CdsHooksServicePlanDefinitionAction struct {
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
	Condition []CdsHooksServicePlanDefinitionActionCondition `json:"condition,omitempty"`
	// Input data requirements
	Input []CdsHooksServicePlanDefinitionActionInput `json:"input,omitempty"`
	// Output data definition
	Output []CdsHooksServicePlanDefinitionActionOutput `json:"output,omitempty"`
	// Relationship to another action
	RelatedAction []CdsHooksServicePlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
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
	Participant []CdsHooksServicePlanDefinitionActionParticipant `json:"participant,omitempty"`
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
	DefinitionCanonical *custom.Canonical[CdsHooksServicePlanDefinitionActionDefinitionCanonical] `json:"definitionCanonical,omitempty"`
	// Description of the activity to be performed
	DefinitionUri *custom.URI `json:"definitionUri,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []CdsHooksServicePlanDefinitionActionDynamicValue `json:"dynamicValue,omitempty"`
	// A sub-action
	Action []interface{} `json:"action,omitempty"`
}

type CdsHooksServicePlanDefinitionActionCondition struct {
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

type CdsHooksServicePlanDefinitionActionInput struct {
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

type CdsHooksServicePlanDefinition struct {
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
	// Extension
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
	SubjectReference *custom.Reference[CdsHooksServicePlanDefinitionSubjectReference] `json:"subjectReference,omitempty"`
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
	Goal []CdsHooksServicePlanDefinitionGoal `json:"goal,omitempty"`
	// Actors within the plan
	Actor []CdsHooksServicePlanDefinitionActor `json:"actor,omitempty"`
	// Action defined by the plan
	Action []CdsHooksServicePlanDefinitionAction `json:"action,omitempty"`
	// Preconditions for service
	AsNeededBoolean *bool `json:"asNeededBoolean,omitempty"`
	// Preconditions for service
	AsNeededCodeableConcept *CodeableConcept `json:"asNeededCodeableConcept,omitempty"`
}

type OtherCdsHooksServicePlanDefinition CdsHooksServicePlanDefinition

type CdsHooksServicePlanDefinitionSubjectReference interface {
	gofhirclient.Resource

	Is_CdsHooksServicePlanDefinitionSubjectReference()
}

func (g Group) Is_CdsHooksServicePlanDefinitionSubjectReference()                          {}
func (m MedicinalProductDefinition) Is_CdsHooksServicePlanDefinitionSubjectReference()     {}
func (s SubstanceDefinition) Is_CdsHooksServicePlanDefinitionSubjectReference()            {}
func (a AdministrableProductDefinition) Is_CdsHooksServicePlanDefinitionSubjectReference() {}
func (m ManufacturedItemDefinition) Is_CdsHooksServicePlanDefinitionSubjectReference()     {}
func (p PackagedProductDefinition) Is_CdsHooksServicePlanDefinitionSubjectReference()      {}

type CdsHooksServicePlanDefinitionActorOptionTypeReference interface {
	gofhirclient.Resource

	Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()
}

func (c CareTeam) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()          {}
func (d Device) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()            {}
func (d DeviceDefinition) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()  {}
func (e Endpoint) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()          {}
func (g Group) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()             {}
func (h HealthcareService) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference() {}
func (l Location) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()          {}
func (o Organization) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()      {}
func (p Patient) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()           {}
func (p Practitioner) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()      {}
func (p PractitionerRole) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()  {}
func (r RelatedPerson) Is_CdsHooksServicePlanDefinitionActorOptionTypeReference()     {}

type CdsHooksServicePlanDefinitionActionParticipantTypeReference interface {
	gofhirclient.Resource

	Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()
}

func (c CareTeam) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()          {}
func (d Device) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()            {}
func (d DeviceDefinition) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()  {}
func (e Endpoint) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()          {}
func (g Group) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()             {}
func (h HealthcareService) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference() {}
func (l Location) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()          {}
func (o Organization) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()      {}
func (p Patient) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()           {}
func (p Practitioner) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()      {}
func (p PractitionerRole) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()  {}
func (r RelatedPerson) Is_CdsHooksServicePlanDefinitionActionParticipantTypeReference()     {}

type CdsHooksServicePlanDefinitionActionDefinitionCanonical interface {
	gofhirclient.Resource

	Is_CdsHooksServicePlanDefinitionActionDefinitionCanonical()
}

func (a ActivityDefinition) Is_CdsHooksServicePlanDefinitionActionDefinitionCanonical()    {}
func (m MessageDefinition) Is_CdsHooksServicePlanDefinitionActionDefinitionCanonical()     {}
func (o ObservationDefinition) Is_CdsHooksServicePlanDefinitionActionDefinitionCanonical() {}
func (p PlanDefinition) Is_CdsHooksServicePlanDefinitionActionDefinitionCanonical()        {}
func (q Questionnaire) Is_CdsHooksServicePlanDefinitionActionDefinitionCanonical()         {}
func (s SpecimenDefinition) Is_CdsHooksServicePlanDefinitionActionDefinitionCanonical()    {}

func (c CdsHooksServicePlanDefinition) ResourceType() string {
	return "CDSHooksServicePlanDefinition"
}

func (c CdsHooksServicePlanDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCdsHooksServicePlanDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCdsHooksServicePlanDefinition: OtherCdsHooksServicePlanDefinition(c), ResourceType: c.ResourceType()})
}

func UnmarshalCdsHooksServicePlanDefinition(b []byte) (res CdsHooksServicePlanDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
