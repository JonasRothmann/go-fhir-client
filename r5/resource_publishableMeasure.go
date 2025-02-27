// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/publishablemeasure
type PublishableMeasureGroupPopulation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique id for population in measure
	LinkId *string `json:"linkId,omitempty"`
	// initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Code *CodeableConcept `json:"code,omitempty"`
	// The human readable description of this population criteria
	Description *custom.Markdown `json:"description,omitempty"`
	// The criteria that defines this population
	Criteria *Expression `json:"criteria,omitempty"`
	// A group resource that defines this population
	GroupDefinition *custom.Reference[Group] `json:"groupDefinition,omitempty"`
	// Which population
	InputPopulationId *string `json:"inputPopulationId,omitempty"`
	// Aggregation method for a measure score (e.g. sum, average, median, minimum, maximum, count)
	AggregateMethod *CodeableConcept `json:"aggregateMethod,omitempty"`
}

type PublishableMeasureGroupStratifier struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique id for stratifier in measure
	LinkId *string `json:"linkId,omitempty"`
	// Meaning of the stratifier
	Code *CodeableConcept `json:"code,omitempty"`
	// The human readable description of this stratifier
	Description *custom.Markdown `json:"description,omitempty"`
	// How the measure should be stratified
	Criteria *Expression `json:"criteria,omitempty"`
	// A group resource that defines this population
	GroupDefinition *custom.Reference[Group] `json:"groupDefinition,omitempty"`
	// Stratifier criteria component for the measure
	Component []PublishableMeasureGroupStratifierComponent `json:"component,omitempty"`
}

type PublishableMeasureGroupStratifierComponent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique id for stratifier component in measure
	LinkId *string `json:"linkId,omitempty"`
	// Meaning of the stratifier component
	Code *CodeableConcept `json:"code,omitempty"`
	// The human readable description of this stratifier component
	Description *custom.Markdown `json:"description,omitempty"`
	// Component of how the measure should be stratified
	Criteria *Expression `json:"criteria,omitempty"`
	// A group resource that defines this population
	GroupDefinition *custom.Reference[Group] `json:"groupDefinition,omitempty"`
}

type PublishableMeasureSupplementalData struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique id for supplementalData in measure
	LinkId *string `json:"linkId,omitempty"`
	// Meaning of the supplemental data
	Code *CodeableConcept `json:"code,omitempty"`
	// supplemental-data | risk-adjustment-factor
	Usage []CodeableConcept `json:"usage,omitempty"`
	// The human readable description of this supplemental data
	Description *custom.Markdown `json:"description,omitempty"`
	// Expression describing additional data to be reported
	Criteria Expression `json:"criteria"`
}

type PublishableMeasure struct {
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
	// Extension
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Canonical identifier for this measure, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Additional identifier for the measure
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the measure
	Version string `json:"version"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this measure (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this measure (human friendly)
	Title string `json:"title"`
	// Subordinate title of the measure
	Subtitle *string `json:"subtitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectReference *custom.Reference[Group] `json:"subjectReference,omitempty"`
	// Population basis
	Basis *custom.Code `json:"basis,omitempty"`
	// Date last changed
	Date custom.DateTime `json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the measure
	Description custom.Markdown `json:"description"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for measure (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this measure is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Describes the clinical usage of the measure
	Usage *custom.Markdown `json:"usage,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the measure was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the measure was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the measure is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// The category of the measure, such as Education, Treatment, Assessment, etc
	Topic []CodeableConcept `json:"topic,omitempty"`
	// Who authored the content
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Logic used by the measure
	Library []custom.Canonical[Library] `json:"library,omitempty"`
	// Disclaimer for use of the measure or its referenced content
	Disclaimer *custom.Markdown `json:"disclaimer,omitempty"`
	// proportion | ratio | continuous-variable | cohort
	Scoring *CodeableConcept `json:"scoring,omitempty"`
	// What units?
	ScoringUnit *CodeableConcept `json:"scoringUnit,omitempty"`
	// opportunity | all-or-nothing | linear | weighted
	CompositeScoring *CodeableConcept `json:"compositeScoring,omitempty"`
	// process | outcome | structure | patient-reported-outcome | composite
	Type []CodeableConcept `json:"type,omitempty"`
	// How risk adjustment is applied for this measure
	RiskAdjustment *custom.Markdown `json:"riskAdjustment,omitempty"`
	// How is rate aggregation performed for this measure
	RateAggregation *custom.Markdown `json:"rateAggregation,omitempty"`
	// Detailed description of why the measure exists
	Rationale *custom.Markdown `json:"rationale,omitempty"`
	// Summary of clinical guidelines
	ClinicalRecommendationStatement *custom.Markdown `json:"clinicalRecommendationStatement,omitempty"`
	// increase | decrease
	ImprovementNotation *CodeableConcept `json:"improvementNotation,omitempty"`
	// Defined terms used in the measure documentation
	Term []PublishableMeasureTerm `json:"term,omitempty"`
	// Additional guidance for implementers (deprecated)
	Guidance *custom.Markdown `json:"guidance,omitempty"`
	// Population criteria group
	Group []PublishableMeasureGroup `json:"group,omitempty"`
	// What other data should be reported with the measure
	SupplementalData []PublishableMeasureSupplementalData `json:"supplementalData,omitempty"`
}

type PublishableMeasureRelatedArtifact struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// documentation | justification | citation | predecessor | successor | derived-from | depends-on | composed-of | part-of | amends | amended-with | appends | appended-with | cites | cited-by | comments-on | comment-in | contains | contained-in | corrects | correction-in | replaces | replaced-with | retracts | retracted-by | signs | similar-to | supports | supported-with | transforms | transformed-into | transformed-with | documents | specification-of | created-with | cite-as
	Type custom.Code `json:"type"`
	// Additional classifiers
	Classifier []CodeableConcept `json:"classifier,omitempty"`
	// Short label
	Label *string `json:"label,omitempty"`
	// Brief description of the related artifact
	Display *string `json:"display,omitempty"`
	// Bibliographic citation for the artifact
	Citation *custom.Markdown `json:"citation,omitempty"`
	// What document is being referenced
	Document *Attachment `json:"document,omitempty"`
	// What artifact is being referenced
	Resource *custom.Canonical[Resource] `json:"resource,omitempty"`
	// What artifact, if not a conformance resource
	ResourceReference *custom.Reference[Resource] `json:"resourceReference,omitempty"`
	// draft | active | retired | unknown
	PublicationStatus *custom.Code `json:"publicationStatus,omitempty"`
	// Date of publication of the artifact being referred to
	PublicationDate *custom.Date `json:"publicationDate,omitempty"`
}

type PublishableMeasureTerm struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What term?
	Code *CodeableConcept `json:"code,omitempty"`
	// Meaning of the term
	Definition *custom.Markdown `json:"definition,omitempty"`
}

type PublishableMeasureGroup struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique id for group in measure
	LinkId *string `json:"linkId,omitempty"`
	// Meaning of the group
	Code *CodeableConcept `json:"code,omitempty"`
	// Summary description
	Description *custom.Markdown `json:"description,omitempty"`
	// process | outcome | structure | patient-reported-outcome | composite
	Type []CodeableConcept `json:"type,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	SubjectReference *custom.Reference[Group] `json:"subjectReference,omitempty"`
	// Population basis
	Basis *custom.Code `json:"basis,omitempty"`
	// proportion | ratio | continuous-variable | cohort
	Scoring *CodeableConcept `json:"scoring,omitempty"`
	// What units?
	ScoringUnit *CodeableConcept `json:"scoringUnit,omitempty"`
	// How is rate aggregation performed for this measure
	RateAggregation *custom.Markdown `json:"rateAggregation,omitempty"`
	// increase | decrease
	ImprovementNotation *CodeableConcept `json:"improvementNotation,omitempty"`
	// Logic used by the measure group
	Library []custom.Canonical[Library] `json:"library,omitempty"`
	// Population criteria
	Population []PublishableMeasureGroupPopulation `json:"population,omitempty"`
	// Stratifier criteria for the measure
	Stratifier []PublishableMeasureGroupStratifier `json:"stratifier,omitempty"`
}

type OtherPublishableMeasure PublishableMeasure

func (p PublishableMeasure) ResourceType() string {
	return "PublishableMeasure"
}

func (p PublishableMeasure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPublishableMeasure
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPublishableMeasure: OtherPublishableMeasure(p), ResourceType: p.ResourceType()})
}

func UnmarshalPublishableMeasure(b []byte) (res PublishableMeasure, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
