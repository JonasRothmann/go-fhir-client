// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
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
	// Business identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	// Why/how the assessment was performed
	Description *string `json:"description,omitempty"`
	// Patient or group assessed
	Subject custom.Reference[ClinicalImpressionSubject] `json:"subject"`
	// The Encounter during which this ClinicalImpression was created
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Time of assessment
	EffectiveDateTime *custom.DateTime `json:"effectiveDateTime,omitempty"`
	// Time of assessment
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// When the assessment was documented
	Date *custom.DateTime `json:"date,omitempty"`
	// The clinician performing the assessment
	Performer *custom.Reference[ClinicalImpressionPerformer] `json:"performer,omitempty"`
	// Reference to last assessment
	Previous *custom.Reference[ClinicalImpression] `json:"previous,omitempty"`
	// Relevant impressions of patient state
	Problem []custom.Reference[ClinicalImpressionProblem] `json:"problem,omitempty"`
	// Change in the status/pattern of a subject's condition since previously assessed, such as worsening, improving, or no change
	ChangePattern *CodeableConcept `json:"changePattern,omitempty"`
	// Clinical Protocol followed
	Protocol []custom.URI `json:"protocol,omitempty"`
	// Summary of the assessment
	Summary *string `json:"summary,omitempty"`
	// Possible or likely findings and diagnoses
	Finding []ClinicalImpressionFinding `json:"finding,omitempty"`
	// Estimate of likely outcome
	PrognosisCodeableConcept []CodeableConcept `json:"prognosisCodeableConcept,omitempty"`
	// RiskAssessment expressing likely outcome
	PrognosisReference []custom.Reference[RiskAssessment] `json:"prognosisReference,omitempty"`
	// Information supporting the clinical impression
	SupportingInfo []custom.Reference[Resource] `json:"supportingInfo,omitempty"`
	// Comments made about the ClinicalImpression
	Note []Annotation `json:"note,omitempty"`
}

type ClinicalImpressionFinding struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What was found
	Item *custom.CodeableReference[ClinicalImpressionFindingItem] `json:"item,omitempty"`
	// Which investigations support finding
	Basis *string `json:"basis,omitempty"`
}

type OtherClinicalImpression ClinicalImpression

type ClinicalImpressionSubject interface {
	gofhirclient.Resource

	Is_ClinicalImpressionSubject()
}

func (p Patient) Is_ClinicalImpressionSubject() {}
func (g Group) Is_ClinicalImpressionSubject()   {}

type ClinicalImpressionPerformer interface {
	gofhirclient.Resource

	Is_ClinicalImpressionPerformer()
}

func (p Practitioner) Is_ClinicalImpressionPerformer()     {}
func (p PractitionerRole) Is_ClinicalImpressionPerformer() {}

type ClinicalImpressionProblem interface {
	gofhirclient.Resource

	Is_ClinicalImpressionProblem()
}

func (c Condition) Is_ClinicalImpressionProblem()          {}
func (a AllergyIntolerance) Is_ClinicalImpressionProblem() {}

type ClinicalImpressionFindingItem interface {
	gofhirclient.Resource

	Is_ClinicalImpressionFindingItem()
}

func (c Condition) Is_ClinicalImpressionFindingItem()         {}
func (o Observation) Is_ClinicalImpressionFindingItem()       {}
func (d DocumentReference) Is_ClinicalImpressionFindingItem() {}

func (c ClinicalImpression) ResourceType() string {
	return "ClinicalImpression"
}

func (c ClinicalImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalImpression
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherClinicalImpression: OtherClinicalImpression(c), ResourceType: c.ResourceType()})
}

func UnmarshalClinicalImpression(b []byte) (res ClinicalImpression, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
