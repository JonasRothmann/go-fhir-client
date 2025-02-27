// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/diagnostic-report-status
*/type DiagnosticReportStatus string

const (
	// The existence of the report is registered, but there is nothing yet available.
	DiagnosticReportStatusRegistered DiagnosticReportStatus = "registered"
	// This is a partial (e.g. initial, interim or preliminary) report: data in the report may be incomplete or unverified.
	DiagnosticReportStatusPartial DiagnosticReportStatus = "partial"
	// The report is complete and verified by an authorized person.
	DiagnosticReportStatusFinal DiagnosticReportStatus = "final"
	// Subsequent to being final, the report has been modified.  This includes any change in the results, diagnosis, narrative text, or other content of a report that has been issued.
	DiagnosticReportStatusAmended DiagnosticReportStatus = "amended"
	// The report is unavailable because the measurement was not started or not completed (also sometimes called "aborted").
	DiagnosticReportStatusCancelled DiagnosticReportStatus = "cancelled"
	// The report has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	DiagnosticReportStatusEnteredInError DiagnosticReportStatus = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	DiagnosticReportStatusUnknown DiagnosticReportStatus = "unknown"
)
