// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Evidence
type Evidence struct {
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
	// Canonical identifier for this evidence, represented as a globally unique URI
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the summary
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of this summary
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this summary (machine friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this summary (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Citation for this evidence
	CiteAs *custom.Reference[Citation] `bson:"citeAs,omitempty" json:"citeAs,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// When the summary was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the summary was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Who authored the content
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Why this Evidence is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Link or citation to artifact associated with the summary
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Description of the particular summary
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Declarative description of the Evidence
	Assertion *custom.Markdown `bson:"assertion,omitempty" json:"assertion,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Evidence variable such as population, exposure, or outcome
	VariableDefinition []EvidenceVariableDefinition `bson:"variableDefinition" json:"variableDefinition"`
	// The method to combine studies
	SynthesisType *CodeableConcept `bson:"synthesisType,omitempty" json:"synthesisType,omitempty"`
	// The design of the study that produced this evidence
	StudyDesign []CodeableConcept `bson:"studyDesign,omitempty" json:"studyDesign,omitempty"`
	// Values and parameters for a single statistic
	Statistic []EvidenceStatistic `bson:"statistic,omitempty" json:"statistic,omitempty"`
	// Certainty or quality of the evidence
	Certainty []EvidenceCertainty `bson:"certainty,omitempty" json:"certainty,omitempty"`
}

type EvidenceVariableDefinition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A text description or summary of the variable
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// population | subpopulation | exposure | referenceExposure | measuredVariable | confounder
	VariableRole CodeableConcept `bson:"variableRole" json:"variableRole"`
	// Definition of the actual variable related to the statistic(s)
	Observed *custom.Reference[EvidenceVariableDefinitionObserved] `bson:"observed,omitempty" json:"observed,omitempty"`
	// Definition of the intended variable related to the Evidence
	Intended *custom.Reference[EvidenceVariableDefinitionIntended] `bson:"intended,omitempty" json:"intended,omitempty"`
	// low | moderate | high | exact
	DirectnessMatch *CodeableConcept `bson:"directnessMatch,omitempty" json:"directnessMatch,omitempty"`
}

type EvidenceStatistic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Description of content
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Type of statistic, e.g., relative risk
	StatisticType *CodeableConcept `bson:"statisticType,omitempty" json:"statisticType,omitempty"`
	// Associated category for categorical variable
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Statistic value
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// The number of events associated with the statistic
	NumberOfEvents *uint32 `bson:"numberOfEvents,omitempty" json:"numberOfEvents,omitempty"`
	// The number of participants affected
	NumberAffected *uint32 `bson:"numberAffected,omitempty" json:"numberAffected,omitempty"`
	// Number of samples in the statistic
	SampleSize *EvidenceStatisticSampleSize `bson:"sampleSize,omitempty" json:"sampleSize,omitempty"`
	// An attribute of the Statistic
	AttributeEstimate []EvidenceStatisticAttributeEstimate `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty"`
	// An aspect of the statistical model
	ModelCharacteristic []EvidenceStatisticModelCharacteristic `bson:"modelCharacteristic,omitempty" json:"modelCharacteristic,omitempty"`
}

type EvidenceStatisticSampleSize struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Textual description of sample size for statistic
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Footnote or explanatory note about the sample size
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Number of contributing studies
	NumberOfStudies *uint32 `bson:"numberOfStudies,omitempty" json:"numberOfStudies,omitempty"`
	// Cumulative number of participants
	NumberOfParticipants *uint32 `bson:"numberOfParticipants,omitempty" json:"numberOfParticipants,omitempty"`
	// Number of participants with known results for measured variables
	KnownDataCount *uint32 `bson:"knownDataCount,omitempty" json:"knownDataCount,omitempty"`
}

type EvidenceStatisticAttributeEstimate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Textual description of the attribute estimate
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Footnote or explanatory note about the estimate
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// The type of attribute estimate, e.g., confidence interval or p value
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The singular quantity of the attribute estimate, for attribute estimates represented as single values; also used to report unit of measure
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Level of confidence interval, e.g., 0.95 for 95% confidence interval
	Level *json.Number `bson:"level,omitempty" json:"level,omitempty"`
	// Lower and upper bound values of the attribute estimate
	Range *Range `bson:"range,omitempty" json:"range,omitempty"`
	// A nested attribute estimate; which is the attribute estimate of an attribute estimate
	AttributeEstimate []interface{} `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty"`
}

type EvidenceStatisticModelCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Model specification
	Code CodeableConcept `bson:"code" json:"code"`
	// Numerical value to complete model specification
	Value *Quantity `bson:"value,omitempty" json:"value,omitempty"`
	// A variable adjusted for in the adjusted analysis
	Variable []EvidenceStatisticModelCharacteristicVariable `bson:"variable,omitempty" json:"variable,omitempty"`
	// An attribute of the statistic used as a model characteristic
	AttributeEstimate []interface{} `bson:"attributeEstimate,omitempty" json:"attributeEstimate,omitempty"`
}

type EvidenceStatisticModelCharacteristicVariable struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Description of the variable
	VariableDefinition custom.Reference[EvidenceStatisticModelCharacteristicVariableVariableDefinition] `bson:"variableDefinition" json:"variableDefinition"`
	// continuous | dichotomous | ordinal | polychotomous
	Handling *custom.Code `bson:"handling,omitempty" json:"handling,omitempty"`
	// Description for grouping of ordinal or polychotomous variables
	ValueCategory []CodeableConcept `bson:"valueCategory,omitempty" json:"valueCategory,omitempty"`
	// Discrete value for grouping of ordinal or polychotomous variables
	ValueQuantity []Quantity `bson:"valueQuantity,omitempty" json:"valueQuantity,omitempty"`
	// Range of values for grouping of ordinal or polychotomous variables
	ValueRange []Range `bson:"valueRange,omitempty" json:"valueRange,omitempty"`
}

type EvidenceCertainty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Textual description of certainty
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Aspect of certainty being rated
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Assessment or judgement of the aspect
	Rating *CodeableConcept `bson:"rating,omitempty" json:"rating,omitempty"`
	// Individual or group who did the rating
	Rater *string `bson:"rater,omitempty" json:"rater,omitempty"`
	// A domain or subdomain of certainty
	Subcomponent []interface{} `bson:"subcomponent,omitempty" json:"subcomponent,omitempty"`
}

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
	return "StructureDefinition"
}
