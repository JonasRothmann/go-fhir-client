// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
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
	// Additional identifier for the MeasureReport
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// complete | pending | error
	Status custom.Code `bson:"status" json:"status"`
	// individual | subject-list | summary | data-exchange
	Type custom.Code `bson:"type" json:"type"`
	// incremental | snapshot
	DataUpdateType *custom.Code `bson:"dataUpdateType,omitempty" json:"dataUpdateType,omitempty"`
	// What measure was calculated
	Measure *custom.Canonical[Measure] `bson:"measure,omitempty" json:"measure,omitempty"`
	// What individual(s) the report is for
	Subject *custom.Reference[MeasureReportSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// When the measure was calculated
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Who is reporting the data
	Reporter *custom.Reference[MeasureReportReporter] `bson:"reporter,omitempty" json:"reporter,omitempty"`
	// What vendor prepared the data
	ReportingVendor *custom.Reference[Organization] `bson:"reportingVendor,omitempty" json:"reportingVendor,omitempty"`
	// Where the reported data is from
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// What period the report covers
	Period Period `bson:"period" json:"period"`
	// What parameters were provided to the report
	InputParameters *custom.Reference[Parameters] `bson:"inputParameters,omitempty" json:"inputParameters,omitempty"`
	// What scoring method (e.g. proportion, ratio, continuous-variable)
	Scoring *CodeableConcept `bson:"scoring,omitempty" json:"scoring,omitempty"`
	// increase | decrease
	ImprovementNotation *CodeableConcept `bson:"improvementNotation,omitempty" json:"improvementNotation,omitempty"`
	// Measure results for each group
	Group []MeasureReportGroup `bson:"group,omitempty" json:"group,omitempty"`
	// Additional information collected for the report
	SupplementalData []custom.Reference[Resource] `bson:"supplementalData,omitempty" json:"supplementalData,omitempty"`
	// What data was used to calculate the measure score
	EvaluatedResource []custom.Reference[Resource] `bson:"evaluatedResource,omitempty" json:"evaluatedResource,omitempty"`
}

type MeasureReportGroup struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Pointer to specific group from Measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Meaning of the group
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// What individual(s) the report is for
	Subject *custom.Reference[MeasureReportGroupSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// The populations in the group
	Population []MeasureReportGroupPopulation `bson:"population,omitempty" json:"population,omitempty"`
	// What score this group achieved
	MeasureScore *Quantity `bson:"measureScore,omitempty" json:"measureScore,omitempty"`
	// Stratification results
	Stratifier []MeasureReportGroupStratifier `bson:"stratifier,omitempty" json:"stratifier,omitempty"`
}

type MeasureReportGroupPopulation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Pointer to specific population from Measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Size of the population
	Count *int32 `bson:"count,omitempty" json:"count,omitempty"`
	// For subject-list reports, the subject results in this population
	SubjectResults *custom.Reference[List] `bson:"subjectResults,omitempty" json:"subjectResults,omitempty"`
	// For subject-list reports, a subject result in this population
	SubjectReport []custom.Reference[MeasureReport] `bson:"subjectReport,omitempty" json:"subjectReport,omitempty"`
	// What individual(s) in the population
	Subjects *custom.Reference[Group] `bson:"subjects,omitempty" json:"subjects,omitempty"`
}

type MeasureReportGroupStratifier struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Pointer to specific stratifier from Measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// What stratifier of the group
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Stratum results, one for each unique value, or set of values, in the stratifier, or stratifier components
	Stratum []MeasureReportGroupStratifierStratum `bson:"stratum,omitempty" json:"stratum,omitempty"`
}

type MeasureReportGroupStratifierStratum struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The stratum value, e.g. male
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
	// Stratifier component values
	Component []MeasureReportGroupStratifierStratumComponent `bson:"component,omitempty" json:"component,omitempty"`
	// Population results in this stratum
	Population []MeasureReportGroupStratifierStratumPopulation `bson:"population,omitempty" json:"population,omitempty"`
	// What score this stratum achieved
	MeasureScore *Quantity `bson:"measureScore,omitempty" json:"measureScore,omitempty"`
}

type MeasureReportGroupStratifierStratumComponent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Pointer to specific stratifier component from Measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// What stratifier component of the group
	Code CodeableConcept `bson:"code" json:"code"`
	// The stratum component value, e.g. male
	Value CodeableConcept `bson:"value" json:"value"`
}

type MeasureReportGroupStratifierStratumPopulation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Pointer to specific population from Measure
	LinkId *string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Size of the population
	Count *int32 `bson:"count,omitempty" json:"count,omitempty"`
	// For subject-list reports, the subject results in this population
	SubjectResults *custom.Reference[List] `bson:"subjectResults,omitempty" json:"subjectResults,omitempty"`
	// For subject-list reports, a subject result in this population
	SubjectReport []custom.Reference[MeasureReport] `bson:"subjectReport,omitempty" json:"subjectReport,omitempty"`
	// What individual(s) in the population
	Subjects *custom.Reference[Group] `bson:"subjects,omitempty" json:"subjects,omitempty"`
}

type MeasureReportSubject interface {
	gofhirclient.Resource

	Is_MeasureReportSubject()
}

func (c CareTeam) Is_MeasureReportSubject()          {}
func (d Device) Is_MeasureReportSubject()            {}
func (g Group) Is_MeasureReportSubject()             {}
func (h HealthcareService) Is_MeasureReportSubject() {}
func (l Location) Is_MeasureReportSubject()          {}
func (o Organization) Is_MeasureReportSubject()      {}
func (p Patient) Is_MeasureReportSubject()           {}
func (p Practitioner) Is_MeasureReportSubject()      {}
func (p PractitionerRole) Is_MeasureReportSubject()  {}
func (r RelatedPerson) Is_MeasureReportSubject()     {}

type MeasureReportReporter interface {
	gofhirclient.Resource

	Is_MeasureReportReporter()
}

func (p Practitioner) Is_MeasureReportReporter()     {}
func (p PractitionerRole) Is_MeasureReportReporter() {}
func (o Organization) Is_MeasureReportReporter()     {}
func (g Group) Is_MeasureReportReporter()            {}

type MeasureReportGroupSubject interface {
	gofhirclient.Resource

	Is_MeasureReportGroupSubject()
}

func (c CareTeam) Is_MeasureReportGroupSubject()          {}
func (d Device) Is_MeasureReportGroupSubject()            {}
func (g Group) Is_MeasureReportGroupSubject()             {}
func (h HealthcareService) Is_MeasureReportGroupSubject() {}
func (l Location) Is_MeasureReportGroupSubject()          {}
func (o Organization) Is_MeasureReportGroupSubject()      {}
func (p Patient) Is_MeasureReportGroupSubject()           {}
func (p Practitioner) Is_MeasureReportGroupSubject()      {}
func (p PractitionerRole) Is_MeasureReportGroupSubject()  {}
func (r RelatedPerson) Is_MeasureReportGroupSubject()     {}

func (m MeasureReport) ResourceType() string {
	return "StructureDefinition"
}
