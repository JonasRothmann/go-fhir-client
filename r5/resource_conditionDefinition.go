// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ConditionDefinition
type ConditionDefinition struct {
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
	// Canonical identifier for this condition definition, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the condition definition
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the condition definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this condition definition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this condition definition (human friendly)
	Title *string `json:"title,omitempty"`
	// Subordinate title of the event definition
	Subtitle *string `json:"subtitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the condition definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for condition definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Identification of the condition, problem or diagnosis
	Code CodeableConcept `json:"code"`
	// Subjective severity of condition
	Severity *CodeableConcept `json:"severity,omitempty"`
	// Anatomical location, if relevant
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	// Stage/grade, usually assessed formally
	Stage *CodeableConcept `json:"stage,omitempty"`
	// Whether Severity is appropriate
	HasSeverity *bool `json:"hasSeverity,omitempty"`
	// Whether bodySite is appropriate
	HasBodySite *bool `json:"hasBodySite,omitempty"`
	// Whether stage is appropriate
	HasStage *bool `json:"hasStage,omitempty"`
	// Formal Definition for the condition
	Definition []custom.URI `json:"definition,omitempty"`
	// Observations particularly relevant to this condition
	Observation []ConditionDefinitionObservation `json:"observation,omitempty"`
	// Medications particularly relevant for this condition
	Medication []ConditionDefinitionMedication `json:"medication,omitempty"`
	// Observation that suggets this condition
	Precondition []ConditionDefinitionPrecondition `json:"precondition,omitempty"`
	// Appropriate team for this condition
	Team []custom.Reference[CareTeam] `json:"team,omitempty"`
	// Questionnaire for this condition
	Questionnaire []ConditionDefinitionQuestionnaire `json:"questionnaire,omitempty"`
	// Plan that is appropriate
	Plan []ConditionDefinitionPlan `json:"plan,omitempty"`
}

type ConditionDefinitionObservation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Category that is relevant
	Category *CodeableConcept `json:"category,omitempty"`
	// Code for relevant Observation
	Code *CodeableConcept `json:"code,omitempty"`
}

type ConditionDefinitionMedication struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Category that is relevant
	Category *CodeableConcept `json:"category,omitempty"`
	// Code for relevant Medication
	Code *CodeableConcept `json:"code,omitempty"`
}

type ConditionDefinitionPrecondition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// sensitive | specific
	Type custom.Code `json:"type"`
	// Code for relevant Observation
	Code CodeableConcept `json:"code"`
	// Value of Observation
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Value of Observation
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
}

type ConditionDefinitionQuestionnaire struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// preadmit | diff-diagnosis | outcome
	Purpose custom.Code `json:"purpose"`
	// Specific Questionnaire
	Reference custom.Reference[Questionnaire] `json:"reference"`
}

type ConditionDefinitionPlan struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Use for the plan
	Role *CodeableConcept `json:"role,omitempty"`
	// The actual plan
	Reference custom.Reference[PlanDefinition] `json:"reference"`
}

type OtherConditionDefinition ConditionDefinition

func (c ConditionDefinition) ResourceType() string {
	return "ConditionDefinition"
}

func (c ConditionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConditionDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherConditionDefinition: OtherConditionDefinition(c), ResourceType: c.ResourceType()})
}

func UnmarshalConditionDefinition(b []byte) (res ConditionDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
