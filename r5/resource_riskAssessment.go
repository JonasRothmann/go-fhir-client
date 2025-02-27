// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/RiskAssessment
type RiskAssessment struct {
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
	// Unique identifier for the assessment
	Identifier []Identifier `json:"identifier,omitempty"`
	// Request fulfilled by this assessment
	BasedOn *custom.Reference[Resource] `json:"basedOn,omitempty"`
	// Part of this occurrence
	Parent *custom.Reference[Resource] `json:"parent,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `json:"status"`
	// Evaluation mechanism
	Method *CodeableConcept `json:"method,omitempty"`
	// Type of assessment
	Code *CodeableConcept `json:"code,omitempty"`
	// Who/what does assessment apply to?
	Subject custom.Reference[RiskAssessmentSubject] `json:"subject"`
	// Where was assessment performed?
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When was assessment made?
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// When was assessment made?
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// Condition assessed
	Condition *custom.Reference[Condition] `json:"condition,omitempty"`
	// Who did assessment?
	Performer *custom.Reference[RiskAssessmentPerformer] `json:"performer,omitempty"`
	// Why the assessment was necessary?
	Reason []custom.CodeableReference[RiskAssessmentReason] `json:"reason,omitempty"`
	// Information used in assessment
	Basis []custom.Reference[Resource] `json:"basis,omitempty"`
	// Outcome predicted
	Prediction []RiskAssessmentPrediction `json:"prediction,omitempty"`
	// How to reduce risk
	Mitigation *string `json:"mitigation,omitempty"`
	// Comments on the risk assessment
	Note []Annotation `json:"note,omitempty"`
}

type RiskAssessmentPrediction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Possible outcome for the subject
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	// Likelihood of specified outcome
	ProbabilityDecimal *json.Number `json:"probabilityDecimal,omitempty"`
	// Likelihood of specified outcome
	ProbabilityRange *Range `json:"probabilityRange,omitempty"`
	// Likelihood of specified outcome as a qualitative value
	QualitativeRisk *CodeableConcept `json:"qualitativeRisk,omitempty"`
	// Relative likelihood
	RelativeRisk *json.Number `json:"relativeRisk,omitempty"`
	// Timeframe or age range
	WhenPeriod *Period `json:"whenPeriod,omitempty"`
	// Timeframe or age range
	WhenRange *Range `json:"whenRange,omitempty"`
	// Explanation of prediction
	Rationale *string `json:"rationale,omitempty"`
}

type OtherRiskAssessment RiskAssessment

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
	return "RiskAssessment"
}

func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskAssessment
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherRiskAssessment: OtherRiskAssessment(r), ResourceType: r.ResourceType()})
}

func UnmarshalRiskAssessment(b []byte) (res RiskAssessment, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
