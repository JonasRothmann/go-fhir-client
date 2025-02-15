// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Measure
type MeasureGroup struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique id for group in measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Meaning of the group
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Summary description
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// process | outcome | structure | patient-reported-outcome | composite
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	Subject *CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// Population basis
	Basis *custom.Code `bson:"basis,omitempty" json:"basis,omitempty"`
	// proportion | ratio | continuous-variable | cohort
	Scoring *CodeableConcept `bson:"scoring,omitempty" json:"scoring,omitempty"`
	// What units?
	ScoringUnit *CodeableConcept `bson:"scoringUnit,omitempty" json:"scoringUnit,omitempty"`
	// How is rate aggregation performed for this measure
	RateAggregation *custom.Markdown `bson:"rateAggregation,omitempty" json:"rateAggregation,omitempty"`
	// increase | decrease
	ImprovementNotation *CodeableConcept `bson:"improvementNotation,omitempty" json:"improvementNotation,omitempty"`
	// Logic used by the measure group
	Library []custom.Canonical[Library] `bson:"library,omitempty" json:"library,omitempty"`
	// Population criteria
	Population []MeasureGroupPopulation `bson:"population,omitempty" json:"population,omitempty"`
	// Stratifier criteria for the measure
	Stratifier []MeasureGroupStratifier `bson:"stratifier,omitempty" json:"stratifier,omitempty"`
}

type MeasureGroupPopulation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique id for population in measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// The human readable description of this population criteria
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The criteria that defines this population
	Criteria *Expression `bson:"criteria,omitempty" json:"criteria,omitempty"`
	// A group resource that defines this population
	GroupDefinition *custom.Reference[Group] `bson:"groupDefinition,omitempty" json:"groupDefinition,omitempty"`
	// Which population
	InputPopulationId *string `bson:"inputPopulationId,omitempty" json:"inputPopulationId,omitempty"`
	// Aggregation method for a measure score (e.g. sum, average, median, minimum, maximum, count)
	AggregateMethod *CodeableConcept `bson:"aggregateMethod,omitempty" json:"aggregateMethod,omitempty"`
}

type MeasureGroupStratifier struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique id for stratifier in measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Meaning of the stratifier
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// The human readable description of this stratifier
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// How the measure should be stratified
	Criteria *Expression `bson:"criteria,omitempty" json:"criteria,omitempty"`
	// A group resource that defines this population
	GroupDefinition *custom.Reference[Group] `bson:"groupDefinition,omitempty" json:"groupDefinition,omitempty"`
	// Stratifier criteria component for the measure
	Component []MeasureGroupStratifierComponent `bson:"component,omitempty" json:"component,omitempty"`
}

type MeasureGroupStratifierComponent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique id for stratifier component in measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Meaning of the stratifier component
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// The human readable description of this stratifier component
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Component of how the measure should be stratified
	Criteria *Expression `bson:"criteria,omitempty" json:"criteria,omitempty"`
	// A group resource that defines this population
	GroupDefinition *custom.Reference[Group] `bson:"groupDefinition,omitempty" json:"groupDefinition,omitempty"`
}

type MeasureSupplementalData struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique id for supplementalData in measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Meaning of the supplemental data
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// supplemental-data | risk-adjustment-factor
	Usage []CodeableConcept `bson:"usage,omitempty" json:"usage,omitempty"`
	// The human readable description of this supplemental data
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Expression describing additional data to be reported
	Criteria Expression `bson:"criteria" json:"criteria"`
}

type Measure struct {
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
	// Canonical identifier for this measure, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the measure
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the measure
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this measure (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this measure (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Subordinate title of the measure
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// E.g. Patient, Practitioner, RelatedPerson, Organization, Location, Device
	Subject *CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// Population basis
	Basis *custom.Code `bson:"basis,omitempty" json:"basis,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the measure
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for measure (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this measure is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Describes the clinical usage of the measure
	Usage *custom.Markdown `bson:"usage,omitempty" json:"usage,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the measure was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the measure was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the measure is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// The category of the measure, such as Education, Treatment, Assessment, etc
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Who authored the content
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Logic used by the measure
	Library []custom.Canonical[Library] `bson:"library,omitempty" json:"library,omitempty"`
	// Disclaimer for use of the measure or its referenced content
	Disclaimer *custom.Markdown `bson:"disclaimer,omitempty" json:"disclaimer,omitempty"`
	// proportion | ratio | continuous-variable | cohort
	Scoring *CodeableConcept `bson:"scoring,omitempty" json:"scoring,omitempty"`
	// What units?
	ScoringUnit *CodeableConcept `bson:"scoringUnit,omitempty" json:"scoringUnit,omitempty"`
	// opportunity | all-or-nothing | linear | weighted
	CompositeScoring *CodeableConcept `bson:"compositeScoring,omitempty" json:"compositeScoring,omitempty"`
	// process | outcome | structure | patient-reported-outcome | composite
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// How risk adjustment is applied for this measure
	RiskAdjustment *custom.Markdown `bson:"riskAdjustment,omitempty" json:"riskAdjustment,omitempty"`
	// How is rate aggregation performed for this measure
	RateAggregation *custom.Markdown `bson:"rateAggregation,omitempty" json:"rateAggregation,omitempty"`
	// Detailed description of why the measure exists
	Rationale *custom.Markdown `bson:"rationale,omitempty" json:"rationale,omitempty"`
	// Summary of clinical guidelines
	ClinicalRecommendationStatement *custom.Markdown `bson:"clinicalRecommendationStatement,omitempty" json:"clinicalRecommendationStatement,omitempty"`
	// increase | decrease
	ImprovementNotation *CodeableConcept `bson:"improvementNotation,omitempty" json:"improvementNotation,omitempty"`
	// Defined terms used in the measure documentation
	Term []MeasureTerm `bson:"term,omitempty" json:"term,omitempty"`
	// Additional guidance for implementers (deprecated)
	Guidance *custom.Markdown `bson:"guidance,omitempty" json:"guidance,omitempty"`
	// Population criteria group
	Group []MeasureGroup `bson:"group,omitempty" json:"group,omitempty"`
	// What other data should be reported with the measure
	SupplementalData []MeasureSupplementalData `bson:"supplementalData,omitempty" json:"supplementalData,omitempty"`
}

type MeasureTerm struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What term?
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Meaning of the term
	Definition *custom.Markdown `bson:"definition,omitempty" json:"definition,omitempty"`
}

func (m Measure) ResourceType() string {
	return "StructureDefinition"
}
