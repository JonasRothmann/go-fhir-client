// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
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
	// Unique identifier for the assessment
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Request fulfilled by this assessment
	BasedOn *custom.Reference[Resource] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Part of this occurrence
	Parent *custom.Reference[Resource] `bson:"parent,omitempty" json:"parent,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `bson:"status" json:"status"`
	// Evaluation mechanism
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Type of assessment
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Who/what does assessment apply to?
	Subject custom.Reference[RiskAssessmentSubject] `bson:"subject" json:"subject"`
	// Where was assessment performed?
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When was assessment made?
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// Condition assessed
	Condition *custom.Reference[Condition] `bson:"condition,omitempty" json:"condition,omitempty"`
	// Who did assessment?
	Performer *custom.Reference[RiskAssessmentPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Why the assessment was necessary?
	Reason []custom.CodeableReference[RiskAssessmentReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Information used in assessment
	Basis []custom.Reference[Resource] `bson:"basis,omitempty" json:"basis,omitempty"`
	// Outcome predicted
	Prediction []RiskAssessmentPrediction `bson:"prediction,omitempty" json:"prediction,omitempty"`
	// How to reduce risk
	Mitigation *string `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
	// Comments on the risk assessment
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type RiskAssessmentPrediction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Possible outcome for the subject
	Outcome *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	// Likelihood of specified outcome
	Probability *json.Number `bson:"probability,omitempty" json:"probability,omitempty"`
	// Likelihood of specified outcome as a qualitative value
	QualitativeRisk *CodeableConcept `bson:"qualitativeRisk,omitempty" json:"qualitativeRisk,omitempty"`
	// Relative likelihood
	RelativeRisk *json.Number `bson:"relativeRisk,omitempty" json:"relativeRisk,omitempty"`
	// Timeframe or age range
	When *Period `bson:"when,omitempty" json:"when,omitempty"`
	// Explanation of prediction
	Rationale *string `bson:"rationale,omitempty" json:"rationale,omitempty"`
}

type RiskAssessmentSubject interface {
	gofhirclient.Resource

	Is_RiskAssessmentSubject()
}

func (p Patient) Is_RiskAssessmentSubject() {}
func (g Group) Is_RiskAssessmentSubject()   {}

type RiskAssessmentPerformer interface {
	gofhirclient.Resource

	Is_RiskAssessmentPerformer()
}

func (p Patient) Is_RiskAssessmentPerformer()          {}
func (p Practitioner) Is_RiskAssessmentPerformer()     {}
func (p PractitionerRole) Is_RiskAssessmentPerformer() {}
func (r RelatedPerson) Is_RiskAssessmentPerformer()    {}
func (d Device) Is_RiskAssessmentPerformer()           {}

type RiskAssessmentReason interface {
	gofhirclient.Resource

	Is_RiskAssessmentReason()
}

func (c Condition) Is_RiskAssessmentReason()         {}
func (o Observation) Is_RiskAssessmentReason()       {}
func (d DiagnosticReport) Is_RiskAssessmentReason()  {}
func (d DocumentReference) Is_RiskAssessmentReason() {}

func (r RiskAssessment) ResourceType() string {
	return "StructureDefinition"
}
