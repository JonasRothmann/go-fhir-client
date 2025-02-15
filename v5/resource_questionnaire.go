// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Questionnaire
type Questionnaire struct {
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
	// Canonical identifier for this questionnaire, represented as an absolute URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Business identifier for questionnaire
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the questionnaire
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this questionnaire (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this questionnaire (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Based on Questionnaire
	DerivedFrom []custom.Canonical[Questionnaire] `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Resource that can be subject of QuestionnaireResponse
	SubjectType []custom.Code `bson:"subjectType,omitempty" json:"subjectType,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the questionnaire
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for questionnaire (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this questionnaire is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the questionnaire was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the questionnaire was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the questionnaire is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// Concept that represents the overall questionnaire
	Code []Coding `bson:"code,omitempty" json:"code,omitempty"`
	// Questions and sections within the Questionnaire
	Item []QuestionnaireItem `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique id for item in questionnaire
	LinkId string `bson:"linkId" json:"linkId"`
	// ElementDefinition - details for the item
	Definition *custom.URI `bson:"definition,omitempty" json:"definition,omitempty"`
	// Corresponding concept for this item in a terminology
	Code []Coding `bson:"code,omitempty" json:"code,omitempty"`
	// E.g. "1(a)", "2.5.3"
	Prefix *string `bson:"prefix,omitempty" json:"prefix,omitempty"`
	// Primary text for the item
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// group | display | boolean | decimal | integer | date | dateTime +
	Type custom.Code `bson:"type" json:"type"`
	// Only allow data when
	EnableWhen []QuestionnaireItemEnableWhen `bson:"enableWhen,omitempty" json:"enableWhen,omitempty"`
	// all | any
	EnableBehavior *custom.Code `bson:"enableBehavior,omitempty" json:"enableBehavior,omitempty"`
	// hidden | protected
	DisabledDisplay *custom.Code `bson:"disabledDisplay,omitempty" json:"disabledDisplay,omitempty"`
	// Whether the item must be included in data results
	Required *bool `bson:"required,omitempty" json:"required,omitempty"`
	// Whether the item may repeat
	Repeats *bool `bson:"repeats,omitempty" json:"repeats,omitempty"`
	// Don't allow human editing
	ReadOnly *bool `bson:"readOnly,omitempty" json:"readOnly,omitempty"`
	// No more than these many characters
	MaxLength *int32 `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	// optionsOnly | optionsOrType | optionsOrString
	AnswerConstraint *custom.Code `bson:"answerConstraint,omitempty" json:"answerConstraint,omitempty"`
	// ValueSet containing permitted answers
	AnswerValueSet *custom.Canonical[ValueSet] `bson:"answerValueSet,omitempty" json:"answerValueSet,omitempty"`
	// Permitted answer
	AnswerOption []QuestionnaireItemAnswerOption `bson:"answerOption,omitempty" json:"answerOption,omitempty"`
	// Initial value(s) when item is first rendered
	Initial []QuestionnaireItemInitial `bson:"initial,omitempty" json:"initial,omitempty"`
	// Nested questionnaire items
	Item []interface{} `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireItemEnableWhen struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The linkId of question that determines whether item is enabled/disabled
	Question string `bson:"question" json:"question"`
	// exists | = | != | > | < | >= | <=
	Operator custom.Code `bson:"operator" json:"operator"`
	// Value for question comparison based on operator
	Answer bool `bson:"answer" json:"answer"`
}

type QuestionnaireItemAnswerOption struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Answer value
	Value int32 `bson:"value" json:"value"`
	// Whether option is selected by default
	InitialSelected *bool `bson:"initialSelected,omitempty" json:"initialSelected,omitempty"`
}

type QuestionnaireItemInitial struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Actual value for initializing the question
	Value bool `bson:"value" json:"value"`
}

func (q Questionnaire) ResourceType() string {
	return "StructureDefinition"
}
