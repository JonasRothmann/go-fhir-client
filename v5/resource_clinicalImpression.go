// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ClinicalImpression
type ClinicalImpression struct {
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
	// Business identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// Why/how the assessment was performed
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Patient or group assessed
	Subject custom.Reference[ClinicalImpressionSubject] `bson:"subject" json:"subject"`
	// The Encounter during which this ClinicalImpression was created
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Time of assessment
	Effective *custom.DateTime `bson:"effective,omitempty" json:"effective,omitempty"`
	// When the assessment was documented
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// The clinician performing the assessment
	Performer *custom.Reference[ClinicalImpressionPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Reference to last assessment
	Previous *custom.Reference[ClinicalImpression] `bson:"previous,omitempty" json:"previous,omitempty"`
	// Relevant impressions of patient state
	Problem []custom.Reference[ClinicalImpressionProblem] `bson:"problem,omitempty" json:"problem,omitempty"`
	// Change in the status/pattern of a subject's condition since previously assessed, such as worsening, improving, or no change
	ChangePattern *CodeableConcept `bson:"changePattern,omitempty" json:"changePattern,omitempty"`
	// Clinical Protocol followed
	Protocol []custom.URI `bson:"protocol,omitempty" json:"protocol,omitempty"`
	// Summary of the assessment
	Summary *string `bson:"summary,omitempty" json:"summary,omitempty"`
	// Possible or likely findings and diagnoses
	Finding []ClinicalImpressionFinding `bson:"finding,omitempty" json:"finding,omitempty"`
	// Estimate of likely outcome
	PrognosisCodeableConcept []CodeableConcept `bson:"prognosisCodeableConcept,omitempty" json:"prognosisCodeableConcept,omitempty"`
	// RiskAssessment expressing likely outcome
	PrognosisReference []custom.Reference[RiskAssessment] `bson:"prognosisReference,omitempty" json:"prognosisReference,omitempty"`
	// Information supporting the clinical impression
	SupportingInfo []custom.Reference[Resource] `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Comments made about the ClinicalImpression
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type ClinicalImpressionFinding struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What was found
	Item *custom.CodeableReference[ClinicalImpressionFindingItem] `bson:"item,omitempty" json:"item,omitempty"`
	// Which investigations support finding
	Basis *string `bson:"basis,omitempty" json:"basis,omitempty"`
}

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
	return "StructureDefinition"
}
