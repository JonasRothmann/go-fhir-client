// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Evidence
type EvidenceStatisticSampleSize struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Textual description of sample size for statistic
	Description *custom.Markdown `json:"description,omitempty"`
	// Footnote or explanatory note about the sample size
	Note []Annotation `json:"note,omitempty"`
	// Number of contributing studies
	NumberOfStudies *uint32 `json:"numberOfStudies,omitempty"`
	// Cumulative number of participants
	NumberOfParticipants *uint32 `json:"numberOfParticipants,omitempty"`
	// Number of participants with known results for measured variables
	KnownDataCount *uint32 `json:"knownDataCount,omitempty"`
}

type EvidenceStatisticAttributeEstimate struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Textual description of the attribute estimate
	Description *custom.Markdown `json:"description,omitempty"`
	// Footnote or explanatory note about the estimate
	Note []Annotation `json:"note,omitempty"`
	// The type of attribute estimate, e.g., confidence interval or p value
	Type *CodeableConcept `json:"type,omitempty"`
	// The singular quantity of the attribute estimate, for attribute estimates represented as single values; also used to report unit of measure
	Quantity *Quantity `json:"quantity,omitempty"`
	// Level of confidence interval, e.g., 0.95 for 95% confidence interval
	Level *json.Number `json:"level,omitempty"`
	// Lower and upper bound values of the attribute estimate
	Range *Range `json:"range,omitempty"`
	// A nested attribute estimate; which is the attribute estimate of an attribute estimate
	AttributeEstimate []interface{} `json:"attributeEstimate,omitempty"`
}

type EvidenceStatisticModelCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Model specification
	Code CodeableConcept `json:"code"`
	// Numerical value to complete model specification
	Value *Quantity `json:"value,omitempty"`
	// A variable adjusted for in the adjusted analysis
	Variable []EvidenceStatisticModelCharacteristicVariable `json:"variable,omitempty"`
	// An attribute of the statistic used as a model characteristic
	AttributeEstimate []interface{} `json:"attributeEstimate,omitempty"`
}

type EvidenceStatisticModelCharacteristicVariable struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Description of the variable
	VariableDefinition custom.Reference[EvidenceStatisticModelCharacteristicVariableVariableDefinition] `json:"variableDefinition"`
	// continuous | dichotomous | ordinal | polychotomous
	Handling *custom.Code `json:"handling,omitempty"`
	// Description for grouping of ordinal or polychotomous variables
	ValueCategory []CodeableConcept `json:"valueCategory,omitempty"`
	// Discrete value for grouping of ordinal or polychotomous variables
	ValueQuantity []Quantity `json:"valueQuantity,omitempty"`
	// Range of values for grouping of ordinal or polychotomous variables
	ValueRange []Range `json:"valueRange,omitempty"`
}

type EvidenceCertainty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Textual description of certainty
	Description *custom.Markdown `json:"description,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`
	// Aspect of certainty being rated
	Type *CodeableConcept `json:"type,omitempty"`
	// Assessment or judgement of the aspect
	Rating *CodeableConcept `json:"rating,omitempty"`
	// Individual or group who did the rating
	Rater *string `json:"rater,omitempty"`
	// A domain or subdomain of certainty
	Subcomponent []interface{} `json:"subcomponent,omitempty"`
}

type Evidence struct {
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
	// Canonical identifier for this evidence, represented as a globally unique URI
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the summary
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of this summary
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this summary (machine friendly)
	Name *string `json:"name,omitempty"`
	// Name for this summary (human friendly)
	Title *string `json:"title,omitempty"`
	// Citation for this evidence
	CiteAsReference *custom.Reference[Citation] `json:"citeAsReference,omitempty"`
	// Citation for this evidence
	CiteAsMarkdown *custom.Markdown `json:"citeAsMarkdown,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// When the summary was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the summary was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Who authored the content
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Why this Evidence is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Link or citation to artifact associated with the summary
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Description of the particular summary
	Description *custom.Markdown `json:"description,omitempty"`
	// Declarative description of the Evidence
	Assertion *custom.Markdown `json:"assertion,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`
	// Evidence variable such as population, exposure, or outcome
	VariableDefinition []EvidenceVariableDefinition `json:"variableDefinition"`
	// The method to combine studies
	SynthesisType *CodeableConcept `json:"synthesisType,omitempty"`
	// The design of the study that produced this evidence
	StudyDesign []CodeableConcept `json:"studyDesign,omitempty"`
	// Values and parameters for a single statistic
	Statistic []EvidenceStatistic `json:"statistic,omitempty"`
	// Certainty or quality of the evidence
	Certainty []EvidenceCertainty `json:"certainty,omitempty"`
}

type EvidenceVariableDefinition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A text description or summary of the variable
	Description *custom.Markdown `json:"description,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`
	// population | subpopulation | exposure | referenceExposure | measuredVariable | confounder
	VariableRole CodeableConcept `json:"variableRole"`
	// Definition of the actual variable related to the statistic(s)
	Observed *custom.Reference[EvidenceVariableDefinitionObserved] `json:"observed,omitempty"`
	// Definition of the intended variable related to the Evidence
	Intended *custom.Reference[EvidenceVariableDefinitionIntended] `json:"intended,omitempty"`
	// low | moderate | high | exact
	DirectnessMatch *CodeableConcept `json:"directnessMatch,omitempty"`
}

type EvidenceStatistic struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Description of content
	Description *custom.Markdown `json:"description,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`
	// Type of statistic, e.g., relative risk
	StatisticType *CodeableConcept `json:"statisticType,omitempty"`
	// Associated category for categorical variable
	Category *CodeableConcept `json:"category,omitempty"`
	// Statistic value
	Quantity *Quantity `json:"quantity,omitempty"`
	// The number of events associated with the statistic
	NumberOfEvents *uint32 `json:"numberOfEvents,omitempty"`
	// The number of participants affected
	NumberAffected *uint32 `json:"numberAffected,omitempty"`
	// Number of samples in the statistic
	SampleSize *EvidenceStatisticSampleSize `json:"sampleSize,omitempty"`
	// An attribute of the Statistic
	AttributeEstimate []EvidenceStatisticAttributeEstimate `json:"attributeEstimate,omitempty"`
	// An aspect of the statistical model
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `json:"modelCharacteristic,omitempty"`
}

type OtherEvidence Evidence

type EvidenceVariableDefinitionObserved interface {
	gofhirclient.Resource

	Is_EvidenceVariableDefinitionObserved()
}

func (g Group) Is_EvidenceVariableDefinitionObserved()            {}
func (e EvidenceVariable) Is_EvidenceVariableDefinitionObserved() {}

type EvidenceVariableDefinitionIntended interface {
	gofhirclient.Resource

	Is_EvidenceVariableDefinitionIntended()
}

func (g Group) Is_EvidenceVariableDefinitionIntended()            {}
func (e EvidenceVariable) Is_EvidenceVariableDefinitionIntended() {}

type EvidenceStatisticModelCharacteristicVariableVariableDefinition interface {
	gofhirclient.Resource

	Is_EvidenceStatisticModelCharacteristicVariableVariableDefinition()
}

func (g Group) Is_EvidenceStatisticModelCharacteristicVariableVariableDefinition()            {}
func (e EvidenceVariable) Is_EvidenceStatisticModelCharacteristicVariableVariableDefinition() {}

func (e Evidence) ResourceType() string {
	return "Evidence"
}

func (e Evidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidence
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherEvidence: OtherEvidence(e), ResourceType: e.ResourceType()})
}

func UnmarshalEvidence(b []byte) (res Evidence, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
