// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
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
	// Business identifier for this set of answers
	Identifier []Identifier `json:"identifier,omitempty"`
	// Request fulfilled by this QuestionnaireResponse
	BasedOn []custom.Reference[QuestionnaireResponseBasedOn] `json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[QuestionnaireResponsePartOf] `json:"partOf,omitempty"`
	// Canonical URL of Questionnaire being answered
	Questionnaire custom.Canonical[Questionnaire] `json:"questionnaire"`
	// in-progress | completed | amended | entered-in-error | stopped
	Status custom.Code `json:"status"`
	// The subject of the questions
	Subject *custom.Reference[Resource] `json:"subject,omitempty"`
	// Encounter the questionnaire response is part of
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Date the answers were gathered
	Authored *custom.DateTime `json:"authored,omitempty"`
	// The individual or device that received and recorded the answers
	Author *custom.Reference[QuestionnaireResponseAuthor] `json:"author,omitempty"`
	// The individual or device that answered the questions
	Source *custom.Reference[QuestionnaireResponseSource] `json:"source,omitempty"`
	// Groups and questions
	Item []QuestionnaireResponseItem `json:"item,omitempty"`
}

type QuestionnaireResponseItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Pointer to specific item from Questionnaire
	LinkId string `json:"linkId"`
	// ElementDefinition - details for the item
	Definition *custom.URI `json:"definition,omitempty"`
	// Name for group or question text
	Text *string `json:"text,omitempty"`
	// The response(s) to the question
	Answer []QuestionnaireResponseItemAnswer `json:"answer,omitempty"`
	// Child items of group item
	Item []interface{} `json:"item,omitempty"`
}

type QuestionnaireResponseItemAnswer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Single-valued answer to the question
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Single-valued answer to the question
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Single-valued answer to the question
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Single-valued answer to the question
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Single-valued answer to the question
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Single-valued answer to the question
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Single-valued answer to the question
	ValueString *string `json:"valueString,omitempty"`
	// Single-valued answer to the question
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Single-valued answer to the question
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Single-valued answer to the question
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Single-valued answer to the question
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Single-valued answer to the question
	ValueReference *custom.Reference[Resource] `json:"valueReference,omitempty"`
	// Child items of question
	Item []interface{} `json:"item,omitempty"`
}

type OtherQuestionnaireResponse QuestionnaireResponse

type QuestionnaireResponseBasedOn interface {
	gofhirclient.Resource

	Is_QuestionnaireResponseBasedOn()
}

func (c CarePlan) Is_QuestionnaireResponseBasedOn()       {}
func (s ServiceRequest) Is_QuestionnaireResponseBasedOn() {}

type QuestionnaireResponsePartOf interface {
	gofhirclient.Resource

	Is_QuestionnaireResponsePartOf()
}

func (o Observation) Is_QuestionnaireResponsePartOf() {}
func (p Procedure) Is_QuestionnaireResponsePartOf()   {}

type QuestionnaireResponseAuthor interface {
	gofhirclient.Resource

	Is_QuestionnaireResponseAuthor()
}

func (d Device) Is_QuestionnaireResponseAuthor()           {}
func (p Practitioner) Is_QuestionnaireResponseAuthor()     {}
func (p PractitionerRole) Is_QuestionnaireResponseAuthor() {}
func (p Patient) Is_QuestionnaireResponseAuthor()          {}
func (r RelatedPerson) Is_QuestionnaireResponseAuthor()    {}
func (o Organization) Is_QuestionnaireResponseAuthor()     {}

type QuestionnaireResponseSource interface {
	gofhirclient.Resource

	Is_QuestionnaireResponseSource()
}

func (d Device) Is_QuestionnaireResponseSource()           {}
func (o Organization) Is_QuestionnaireResponseSource()     {}
func (p Patient) Is_QuestionnaireResponseSource()          {}
func (p Practitioner) Is_QuestionnaireResponseSource()     {}
func (p PractitionerRole) Is_QuestionnaireResponseSource() {}
func (r RelatedPerson) Is_QuestionnaireResponseSource()    {}

func (q QuestionnaireResponse) ResourceType() string {
	return "QuestionnaireResponse"
}

func (q QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaireResponse
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherQuestionnaireResponse: OtherQuestionnaireResponse(q), ResourceType: q.ResourceType()})
}

func UnmarshalQuestionnaireResponse(b []byte) (res QuestionnaireResponse, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
