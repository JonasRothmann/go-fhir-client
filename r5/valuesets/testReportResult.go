// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/report-result-codes
*/type TestReportResult string

const (
	// All test operations successfully passed all asserts.
	TestReportResultPass TestReportResult = "pass"
	// One or more test operations failed one or more asserts.
	TestReportResultFail TestReportResult = "fail"
	// One or more test operations is pending execution completion.
	TestReportResultPending TestReportResult = "pending"
)
