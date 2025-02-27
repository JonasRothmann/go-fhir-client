// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/measure-report-type
*/type MeasureReportType string

const (
	// An individual report that provides information on the performance for a given measure with respect to a single subject.
	MeasureReportTypeIndividual MeasureReportType = "individual"
	// A subject list report that includes a listing of subjects that satisfied each population criteria in the measure.
	MeasureReportTypeSubjectList MeasureReportType = "subject-list"
	// A summary report that returns the number of members in each population criteria for the measure.
	MeasureReportTypeSummary MeasureReportType = "summary"
	// A data exchange report that contains data-of-interest for the measure (i.e. data that is needed to calculate the measure)
	MeasureReportTypeDataExchange MeasureReportType = "data-exchange"
)
