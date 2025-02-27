// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/report-status-codes
*/type TestReportStatus string

const (
	// All test operations have completed.
	TestReportStatusCompleted TestReportStatus = "completed"
	// A test operations is currently executing.
	TestReportStatusInProgress TestReportStatus = "in-progress"
	// A test operation is waiting for an external client request.
	TestReportStatusWaiting TestReportStatus = "waiting"
	// The test script execution was manually stopped.
	TestReportStatusStopped TestReportStatus = "stopped"
	// This test report was entered or created in error.
	TestReportStatusEnteredInError TestReportStatus = "entered-in-error"
)
