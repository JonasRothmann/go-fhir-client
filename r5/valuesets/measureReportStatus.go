// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/measure-report-status
*/type MeasureReportStatus string

const (
	// The report is complete and ready for use.
	MeasureReportStatusComplete MeasureReportStatus = "complete"
	// The report is currently being generated.
	MeasureReportStatusPending MeasureReportStatus = "pending"
	// An error occurred attempting to generate the report.
	MeasureReportStatusError MeasureReportStatus = "error"
)
