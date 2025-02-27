// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/computableplandefinition
type ComputablePlanDefinition struct {
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
	SubjectReference *custom.Reference[ComputablePlanDefinitionSubjectReference] `json:"subjectReference,omitempty"`
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
	Library custom.Canonical[Library] `json:"library"`
	// What the plan is trying to accomplish
	Goal []ComputablePlanDefinitionGoal `json:"goal,omitempty"`
	// Actors within the plan
	Actor []ComputablePlanDefinitionActor `json:"actor,omitempty"`
	// Action defined by the plan
	Action []ComputablePlanDefinitionAction `json:"action,omitempty"`
	// Preconditions for service
	AsNeededBoolean *bool `json:"asNeededBoolean,omitempty"`
	// Preconditions for service
	AsNeededCodeableConcept *CodeableConcept `json:"asNeededCodeableConcept,omitempty"`
}

type ComputablePlanDefinitionActor struct {
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
	Option []ComputablePlanDefinitionActorOption `json:"option"`
}

type ComputablePlanDefinitionActionInput struct {
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

type ComputablePlanDefinitionActionOutput struct {
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

type ComputablePlanDefinitionActionParticipant struct {
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
	TypeReference *custom.Reference[ComputablePlanDefinitionActionParticipantTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent
	Role *CodeableConcept `json:"role,omitempty"`
	// E.g. Author, Reviewer, Witness, etc
	Function *CodeableConcept `json:"function,omitempty"`
}

type ComputablePlanDefinitionActionDynamicValue struct {
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

type ComputablePlanDefinitionGoal struct {
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
	Target []ComputablePlanDefinitionGoalTarget `json:"target,omitempty"`
}

type ComputablePlanDefinitionGoalTarget struct {
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

type ComputablePlanDefinitionActorOption struct {
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
	TypeReference *custom.Reference[ComputablePlanDefinitionActorOptionTypeReference] `json:"typeReference,omitempty"`
	// E.g. Nurse, Surgeon, Parent
	Role *CodeableConcept `json:"role,omitempty"`
}

type ComputablePlanDefinitionAction struct {
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
	Condition []ComputablePlanDefinitionActionCondition `json:"condition,omitempty"`
	// Input data requirements
	Input []ComputablePlanDefinitionActionInput `json:"input,omitempty"`
	// Output data definition
	Output []ComputablePlanDefinitionActionOutput `json:"output,omitempty"`
	// Relationship to another action
	RelatedAction []ComputablePlanDefinitionActionRelatedAction `json:"relatedAction,omitempty"`
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
	Participant []ComputablePlanDefinitionActionParticipant `json:"participant,omitempty"`
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
	DefinitionCanonical *custom.Canonical[ComputablePlanDefinitionActionDefinitionCanonical] `json:"definitionCanonical,omitempty"`
	// Description of the activity to be performed
	DefinitionUri *custom.URI `json:"definitionUri,omitempty"`
	// Transform to apply the template
	Transform *custom.Canonical[StructureMap] `json:"transform,omitempty"`
	// Dynamic aspects of the definition
	DynamicValue []ComputablePlanDefinitionActionDynamicValue `json:"dynamicValue,omitempty"`
	// A sub-action
	Action []interface{} `json:"action,omitempty"`
}

type ComputablePlanDefinitionActionCondition struct {
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

type ComputablePlanDefinitionActionRelatedAction struct {
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

type OtherComputablePlanDefinition ComputablePlanDefinition

type ComputablePlanDefinitionSubjectReference interface {
	gofhirclient.Resource

	Is_ComputablePlanDefinitionSubjectReference()
}

func (g Group) Is_ComputablePlanDefinitionSubjectReference()                          {}
func (m MedicinalProductDefinition) Is_ComputablePlanDefinitionSubjectReference()     {}
func (s SubstanceDefinition) Is_ComputablePlanDefinitionSubjectReference()            {}
func (a AdministrableProductDefinition) Is_ComputablePlanDefinitionSubjectReference() {}
func (m ManufacturedItemDefinition) Is_ComputablePlanDefinitionSubjectReference()     {}
func (p PackagedProductDefinition) Is_ComputablePlanDefinitionSubjectReference()      {}

type ComputablePlanDefinitionActorOptionTypeReference interface {
	gofhirclient.Resource

	Is_ComputablePlanDefinitionActorOptionTypeReference()
}

func (c CareTeam) Is_ComputablePlanDefinitionActorOptionTypeReference()          {}
func (d Device) Is_ComputablePlanDefinitionActorOptionTypeReference()            {}
func (d DeviceDefinition) Is_ComputablePlanDefinitionActorOptionTypeReference()  {}
func (e Endpoint) Is_ComputablePlanDefinitionActorOptionTypeReference()          {}
func (g Group) Is_ComputablePlanDefinitionActorOptionTypeReference()             {}
func (h HealthcareService) Is_ComputablePlanDefinitionActorOptionTypeReference() {}
func (l Location) Is_ComputablePlanDefinitionActorOptionTypeReference()          {}
func (o Organization) Is_ComputablePlanDefinitionActorOptionTypeReference()      {}
func (p Patient) Is_ComputablePlanDefinitionActorOptionTypeReference()           {}
func (p Practitioner) Is_ComputablePlanDefinitionActorOptionTypeReference()      {}
func (p PractitionerRole) Is_ComputablePlanDefinitionActorOptionTypeReference()  {}
func (r RelatedPerson) Is_ComputablePlanDefinitionActorOptionTypeReference()     {}

type ComputablePlanDefinitionActionParticipantTypeReference interface {
	gofhirclient.Resource

	Is_ComputablePlanDefinitionActionParticipantTypeReference()
}

func (c CareTeam) Is_ComputablePlanDefinitionActionParticipantTypeReference()          {}
func (d Device) Is_ComputablePlanDefinitionActionParticipantTypeReference()            {}
func (d DeviceDefinition) Is_ComputablePlanDefinitionActionParticipantTypeReference()  {}
func (e Endpoint) Is_ComputablePlanDefinitionActionParticipantTypeReference()          {}
func (g Group) Is_ComputablePlanDefinitionActionParticipantTypeReference()             {}
func (h HealthcareService) Is_ComputablePlanDefinitionActionParticipantTypeReference() {}
func (l Location) Is_ComputablePlanDefinitionActionParticipantTypeReference()          {}
func (o Organization) Is_ComputablePlanDefinitionActionParticipantTypeReference()      {}
func (p Patient) Is_ComputablePlanDefinitionActionParticipantTypeReference()           {}
func (p Practitioner) Is_ComputablePlanDefinitionActionParticipantTypeReference()      {}
func (p PractitionerRole) Is_ComputablePlanDefinitionActionParticipantTypeReference()  {}
func (r RelatedPerson) Is_ComputablePlanDefinitionActionParticipantTypeReference()     {}

type ComputablePlanDefinitionActionDefinitionCanonical interface {
	gofhirclient.Resource

	Is_ComputablePlanDefinitionActionDefinitionCanonical()
}

func (a ActivityDefinition) Is_ComputablePlanDefinitionActionDefinitionCanonical()    {}
func (m MessageDefinition) Is_ComputablePlanDefinitionActionDefinitionCanonical()     {}
func (o ObservationDefinition) Is_ComputablePlanDefinitionActionDefinitionCanonical() {}
func (p PlanDefinition) Is_ComputablePlanDefinitionActionDefinitionCanonical()        {}
func (q Questionnaire) Is_ComputablePlanDefinitionActionDefinitionCanonical()         {}
func (s SpecimenDefinition) Is_ComputablePlanDefinitionActionDefinitionCanonical()    {}

func (c ComputablePlanDefinition) ResourceType() string {
	return "ComputablePlanDefinition"
}

func (c ComputablePlanDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherComputablePlanDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherComputablePlanDefinition: OtherComputablePlanDefinition(c), ResourceType: c.ResourceType()})
}

func UnmarshalComputablePlanDefinition(b []byte) (res ComputablePlanDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
