// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Questionnaire
type Questionnaire struct {
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
	// Canonical identifier for this questionnaire, represented as an absolute URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Business identifier for questionnaire
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the questionnaire
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this questionnaire (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this questionnaire (human friendly)
	Title *string `json:"title,omitempty"`
	// Based on Questionnaire
	DerivedFrom []custom.Canonical[Questionnaire] `json:"derivedFrom,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Resource that can be subject of QuestionnaireResponse
	SubjectType []custom.Code `json:"subjectType,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the questionnaire
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for questionnaire (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this questionnaire is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the questionnaire was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the questionnaire was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the questionnaire is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// Concept that represents the overall questionnaire
	Code []Coding `json:"code,omitempty"`
	// Questions and sections within the Questionnaire
	Item []QuestionnaireItem `json:"item,omitempty"`
}

type QuestionnaireItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique id for item in questionnaire
	LinkId string `json:"linkId"`
	// ElementDefinition - details for the item
	Definition *custom.URI `json:"definition,omitempty"`
	// Corresponding concept for this item in a terminology
	Code []Coding `json:"code,omitempty"`
	// E.g. "1(a)", "2.5.3"
	Prefix *string `json:"prefix,omitempty"`
	// Primary text for the item
	Text *string `json:"text,omitempty"`
	// group | display | boolean | decimal | integer | date | dateTime +
	Type custom.Code `json:"type"`
	// Only allow data when
	EnableWhen []QuestionnaireItemEnableWhen `json:"enableWhen,omitempty"`
	// all | any
	EnableBehavior *custom.Code `json:"enableBehavior,omitempty"`
	// hidden | protected
	DisabledDisplay *custom.Code `json:"disabledDisplay,omitempty"`
	// Whether the item must be included in data results
	Required *bool `json:"required,omitempty"`
	// Whether the item may repeat
	Repeats *bool `json:"repeats,omitempty"`
	// Don't allow human editing
	ReadOnly *bool `json:"readOnly,omitempty"`
	// No more than these many characters
	MaxLength *int32 `json:"maxLength,omitempty"`
	// optionsOnly | optionsOrType | optionsOrString
	AnswerConstraint *custom.Code `json:"answerConstraint,omitempty"`
	// ValueSet containing permitted answers
	AnswerValueSet *custom.Canonical[ValueSet] `json:"answerValueSet,omitempty"`
	// Permitted answer
	AnswerOption []QuestionnaireItemAnswerOption `json:"answerOption,omitempty"`
	// Initial value(s) when item is first rendered
	Initial []QuestionnaireItemInitial `json:"initial,omitempty"`
	// Nested questionnaire items
	Item []interface{} `json:"item,omitempty"`
}

type QuestionnaireItemEnableWhen struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The linkId of question that determines whether item is enabled/disabled
	Question string `json:"question"`
	// exists | = | != | > | < | >= | <=
	Operator custom.Code `json:"operator"`
	// Value for question comparison based on operator
	AnswerBoolean *bool `json:"answerBoolean,omitempty"`
	// Value for question comparison based on operator
	AnswerDecimal *json.Number `json:"answerDecimal,omitempty"`
	// Value for question comparison based on operator
	AnswerInteger *int32 `json:"answerInteger,omitempty"`
	// Value for question comparison based on operator
	AnswerDate *custom.Date `json:"answerDate,omitempty"`
	// Value for question comparison based on operator
	AnswerDateTime *custom.DateTime `json:"answerDateTime,omitempty"`
	// Value for question comparison based on operator
	AnswerTime *custom.Time `json:"answerTime,omitempty"`
	// Value for question comparison based on operator
	AnswerString *string `json:"answerString,omitempty"`
	// Value for question comparison based on operator
	AnswerCoding *Coding `json:"answerCoding,omitempty"`
	// Value for question comparison based on operator
	AnswerQuantity *Quantity `json:"answerQuantity,omitempty"`
	// Value for question comparison based on operator
	AnswerReference *custom.Reference[Resource] `json:"answerReference,omitempty"`
}

type QuestionnaireItemAnswerOption struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Answer value
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Answer value
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Answer value
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Answer value
	ValueString *string `json:"valueString,omitempty"`
	// Answer value
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Answer value
	ValueReference *custom.Reference[Resource] `json:"valueReference,omitempty"`
	// Whether option is selected by default
	InitialSelected *bool `json:"initialSelected,omitempty"`
}

type QuestionnaireItemInitial struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Actual value for initializing the question
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Actual value for initializing the question
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Actual value for initializing the question
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Actual value for initializing the question
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Actual value for initializing the question
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Actual value for initializing the question
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Actual value for initializing the question
	ValueString *string `json:"valueString,omitempty"`
	// Actual value for initializing the question
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Actual value for initializing the question
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Actual value for initializing the question
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Actual value for initializing the question
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Actual value for initializing the question
	ValueReference *custom.Reference[Resource] `json:"valueReference,omitempty"`
}

type OtherQuestionnaire Questionnaire

func (q Questionnaire) ResourceType() string {
	return "Questionnaire"
}

func (q Questionnaire) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaire
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherQuestionnaire: OtherQuestionnaire(q), ResourceType: q.ResourceType()})
}

func UnmarshalQuestionnaire(b []byte) (res Questionnaire, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
