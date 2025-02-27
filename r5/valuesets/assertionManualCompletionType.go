// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/assert-manual-completion-codes
*/type AssertionManualCompletionType string

const (
	// Mark the currently waiting test failed and proceed with the next assert if the stopTestOnFail is false or the next test in the TestScript if the stopTestOnFail is true.
	AssertionManualCompletionTypeFail AssertionManualCompletionType = "fail"
	// Mark the currently waiting test passed (if the test is not failed already) and proceed with the next action in the TestScript.
	AssertionManualCompletionTypePass AssertionManualCompletionType = "pass"
	// Mark this assert as skipped and proceed with the next action in the TestScript.
	AssertionManualCompletionTypeSkip AssertionManualCompletionType = "skip"
	// Stop execution of this TestScript. The overall status of this TestScript is evaluated based on the status of the completed tests.
	AssertionManualCompletionTypeStop AssertionManualCompletionType = "stop"
)
