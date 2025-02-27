// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MeasureReport
type MeasureReport struct {
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
	// Additional identifier for the MeasureReport
	Identifier []Identifier `json:"identifier,omitempty"`
	// complete | pending | error
	Status custom.Code `json:"status"`
	// individual | subject-list | summary | data-exchange
	Type custom.Code `json:"type"`
	// incremental | snapshot
	DataUpdateType *custom.Code `json:"dataUpdateType,omitempty"`
	// What measure was calculated
	Measure *custom.Canonical[Measure] `json:"measure,omitempty"`
	// What individual(s) the report is for
	Subject *custom.Reference[MeasureReportSubject] `json:"subject,omitempty"`
	// When the measure was calculated
	Date *custom.DateTime `json:"date,omitempty"`
	// Who is reporting the data
	Reporter *custom.Reference[MeasureReportReporter] `json:"reporter,omitempty"`
	// What vendor prepared the data
	ReportingVendor *custom.Reference[Organization] `json:"reportingVendor,omitempty"`
	// Where the reported data is from
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// What period the report covers
	Period Period `json:"period"`
	// What parameters were provided to the report
	InputParameters *custom.Reference[Parameters] `json:"inputParameters,omitempty"`
	// What scoring method (e.g. proportion, ratio, continuous-variable)
	Scoring *CodeableConcept `json:"scoring,omitempty"`
	// increase | decrease
	ImprovementNotation *CodeableConcept `json:"improvementNotation,omitempty"`
	// Measure results for each group
	Group []MeasureReportGroup `json:"group,omitempty"`
	// Additional information collected for the report
	SupplementalData []custom.Reference[Resource] `json:"supplementalData,omitempty"`
	// What data was used to calculate the measure score
	EvaluatedResource []custom.Reference[Resource] `json:"evaluatedResource,omitempty"`
}

type MeasureReportGroup struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Pointer to specific group from Measure
	LinkId *string `json:"linkId,omitempty"`
	// Meaning of the group
	Code *CodeableConcept `json:"code,omitempty"`
	// What individual(s) the report is for
	Subject *custom.Reference[MeasureReportGroupSubject] `json:"subject,omitempty"`
	// The populations in the group
	Population []MeasureReportGroupPopulation `json:"population,omitempty"`
	// What score this group achieved
	MeasureScoreQuantity *Quantity `json:"measureScoreQuantity,omitempty"`
	// What score this group achieved
	MeasureScoreDateTime *custom.DateTime `json:"measureScoreDateTime,omitempty"`
	// What score this group achieved
	MeasureScoreCodeableConcept *CodeableConcept `json:"measureScoreCodeableConcept,omitempty"`
	// What score this group achieved
	MeasureScorePeriod *Period `json:"measureScorePeriod,omitempty"`
	// What score this group achieved
	MeasureScoreRange *Range `json:"measureScoreRange,omitempty"`
	// What score this group achieved
	MeasureScoreDuration *Duration `json:"measureScoreDuration,omitempty"`
	// Stratification results
	Stratifier []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

type MeasureReportGroupPopulation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Pointer to specific population from Measure
	LinkId *string `json:"linkId,omitempty"`
	// initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Code *CodeableConcept `json:"code,omitempty"`
	// Size of the population
	Count *int32 `json:"count,omitempty"`
	// For subject-list reports, the subject results in this population
	SubjectResults *custom.Reference[List] `json:"subjectResults,omitempty"`
	// For subject-list reports, a subject result in this population
	SubjectReport []custom.Reference[MeasureReport] `json:"subjectReport,omitempty"`
	// What individual(s) in the population
	Subjects *custom.Reference[Group] `json:"subjects,omitempty"`
}

type MeasureReportGroupStratifier struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Pointer to specific stratifier from Measure
	LinkId *string `json:"linkId,omitempty"`
	// What stratifier of the group
	Code *CodeableConcept `json:"code,omitempty"`
	// Stratum results, one for each unique value, or set of values, in the stratifier, or stratifier components
	Stratum []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

type MeasureReportGroupStratifierStratum struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The stratum value, e.g. male
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// The stratum value, e.g. male
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// The stratum value, e.g. male
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The stratum value, e.g. male
	ValueRange *Range `json:"valueRange,omitempty"`
	// The stratum value, e.g. male
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// Stratifier component values
	Component []MeasureReportGroupStratifierStratumComponent `json:"component,omitempty"`
	// Population results in this stratum
	Population []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty"`
	// What score this stratum achieved
	MeasureScoreQuantity *Quantity `json:"measureScoreQuantity,omitempty"`
	// What score this stratum achieved
	MeasureScoreDateTime *custom.DateTime `json:"measureScoreDateTime,omitempty"`
	// What score this stratum achieved
	MeasureScoreCodeableConcept *CodeableConcept `json:"measureScoreCodeableConcept,omitempty"`
	// What score this stratum achieved
	MeasureScorePeriod *Period `json:"measureScorePeriod,omitempty"`
	// What score this stratum achieved
	MeasureScoreRange *Range `json:"measureScoreRange,omitempty"`
	// What score this stratum achieved
	MeasureScoreDuration *Duration `json:"measureScoreDuration,omitempty"`
}

type MeasureReportGroupStratifierStratumComponent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Pointer to specific stratifier component from Measure
	LinkId *string `json:"linkId,omitempty"`
	// What stratifier component of the group
	Code CodeableConcept `json:"code"`
	// The stratum component value, e.g. male
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// The stratum component value, e.g. male
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// The stratum component value, e.g. male
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The stratum component value, e.g. male
	ValueRange *Range `json:"valueRange,omitempty"`
	// The stratum component value, e.g. male
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
}

type MeasureReportGroupStratifierStratumPopulation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Pointer to specific population from Measure
	LinkId *string `json:"linkId,omitempty"`
	// initial-population | numerator | numerator-exclusion | denominator | denominator-exclusion | denominator-exception | measure-population | measure-population-exclusion | measure-observation
	Code *CodeableConcept `json:"code,omitempty"`
	// Size of the population
	Count *int32 `json:"count,omitempty"`
	// For subject-list reports, the subject results in this population
	SubjectResults *custom.Reference[List] `json:"subjectResults,omitempty"`
	// For subject-list reports, a subject result in this population
	SubjectReport []custom.Reference[MeasureReport] `json:"subjectReport,omitempty"`
	// What individual(s) in the population
	Subjects *custom.Reference[Group] `json:"subjects,omitempty"`
}

type OtherMeasureReport MeasureReport

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
	return "MeasureReport"
}

func (m MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasureReport
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMeasureReport: OtherMeasureReport(m), ResourceType: m.ResourceType()})
}

func UnmarshalMeasureReport(b []byte) (res MeasureReport, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
