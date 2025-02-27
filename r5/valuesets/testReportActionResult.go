// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/report-action-result-codes
*/type TestReportActionResult string

const (
	// The action was successful.
	TestReportActionResultPass TestReportActionResult = "pass"
	// The action was skipped.
	TestReportActionResultSkip TestReportActionResult = "skip"
	// The action failed.
	TestReportActionResultFail TestReportActionResult = "fail"
	// The action passed but with warnings.
	TestReportActionResultWarning TestReportActionResult = "warning"
	// The action encountered a fatal error and the engine was unable to process.
	TestReportActionResultError TestReportActionResult = "error"
)
