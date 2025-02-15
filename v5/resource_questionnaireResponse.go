// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponseItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Pointer to specific item from Questionnaire
	LinkId string `bson:"linkId" json:"linkId"`
	// ElementDefinition - details for the item
	Definition *custom.URI `bson:"definition,omitempty" json:"definition,omitempty"`
	// Name for group or question text
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// The response(s) to the question
	Answer []QuestionnaireResponseItemAnswer `bson:"answer,omitempty" json:"answer,omitempty"`
	// Child items of group item
	Item []interface{} `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireResponseItemAnswer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Single-valued answer to the question
	Value bool `bson:"value" json:"value"`
	// Child items of question
	Item []interface{} `bson:"item,omitempty" json:"item,omitempty"`
}

type QuestionnaireResponse struct {
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
	// Business identifier for this set of answers
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Request fulfilled by this QuestionnaireResponse
	BasedOn []custom.Reference[QuestionnaireResponseBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[QuestionnaireResponsePartOf] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// Canonical URL of Questionnaire being answered
	Questionnaire custom.Canonical[Questionnaire] `bson:"questionnaire" json:"questionnaire"`
	// in-progress | completed | amended | entered-in-error | stopped
	Status custom.Code `bson:"status" json:"status"`
	// The subject of the questions
	Subject *custom.Reference[Resource] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Encounter the questionnaire response is part of
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Date the answers were gathered
	Authored *custom.DateTime `bson:"authored,omitempty" json:"authored,omitempty"`
	// The individual or device that received and recorded the answers
	Author *custom.Reference[QuestionnaireResponseAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// The individual or device that answered the questions
	Source *custom.Reference[QuestionnaireResponseSource] `bson:"source,omitempty" json:"source,omitempty"`
	// Groups and questions
	Item []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
}

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
	return "StructureDefinition"
}
