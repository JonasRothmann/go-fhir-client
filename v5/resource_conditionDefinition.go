// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ConditionDefinition
type ConditionDefinitionQuestionnaire struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// preadmit | diff-diagnosis | outcome
	Purpose custom.Code `bson:"purpose" json:"purpose"`
	// Specific Questionnaire
	Reference custom.Reference[Questionnaire] `bson:"reference" json:"reference"`
}

type ConditionDefinitionPlan struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Use for the plan
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// The actual plan
	Reference custom.Reference[PlanDefinition] `bson:"reference" json:"reference"`
}

type ConditionDefinition struct {
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
	// Canonical identifier for this condition definition, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the condition definition
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the condition definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this condition definition (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this condition definition (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Subordinate title of the event definition
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the condition definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for condition definition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Identification of the condition, problem or diagnosis
	Code CodeableConcept `bson:"code" json:"code"`
	// Subjective severity of condition
	Severity *CodeableConcept `bson:"severity,omitempty" json:"severity,omitempty"`
	// Anatomical location, if relevant
	BodySite *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Stage/grade, usually assessed formally
	Stage *CodeableConcept `bson:"stage,omitempty" json:"stage,omitempty"`
	// Whether Severity is appropriate
	HasSeverity *bool `bson:"hasSeverity,omitempty" json:"hasSeverity,omitempty"`
	// Whether bodySite is appropriate
	HasBodySite *bool `bson:"hasBodySite,omitempty" json:"hasBodySite,omitempty"`
	// Whether stage is appropriate
	HasStage *bool `bson:"hasStage,omitempty" json:"hasStage,omitempty"`
	// Formal Definition for the condition
	Definition []custom.URI `bson:"definition,omitempty" json:"definition,omitempty"`
	// Observations particularly relevant to this condition
	Observation []ConditionDefinitionObservation `bson:"observation,omitempty" json:"observation,omitempty"`
	// Medications particularly relevant for this condition
	Medication []ConditionDefinitionMedication `bson:"medication,omitempty" json:"medication,omitempty"`
	// Observation that suggets this condition
	Precondition []ConditionDefinitionPrecondition `bson:"precondition,omitempty" json:"precondition,omitempty"`
	// Appropriate team for this condition
	Team []custom.Reference[CareTeam] `bson:"team,omitempty" json:"team,omitempty"`
	// Questionnaire for this condition
	Questionnaire []ConditionDefinitionQuestionnaire `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	// Plan that is appropriate
	Plan []ConditionDefinitionPlan `bson:"plan,omitempty" json:"plan,omitempty"`
}

type ConditionDefinitionObservation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Category that is relevant
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Code for relevant Observation
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}

type ConditionDefinitionMedication struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Category that is relevant
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Code for relevant Medication
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}

type ConditionDefinitionPrecondition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// sensitive | specific
	Type custom.Code `bson:"type" json:"type"`
	// Code for relevant Observation
	Code CodeableConcept `bson:"code" json:"code"`
	// Value of Observation
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

func (c ConditionDefinition) ResourceType() string {
	return "StructureDefinition"
}
