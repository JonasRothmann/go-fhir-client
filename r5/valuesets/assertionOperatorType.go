// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/assert-operator-codes
*/type AssertionOperatorType string

const (
	// Default value. Equals comparison.
	AssertionOperatorTypeEquals AssertionOperatorType = "equals"
	// Not equals comparison.
	AssertionOperatorTypeNotEquals AssertionOperatorType = "notEquals"
	// Compare value within a known set of values.
	AssertionOperatorTypeIn AssertionOperatorType = "in"
	// Compare value not within a known set of values.
	AssertionOperatorTypeNotIn AssertionOperatorType = "notIn"
	// Compare value to be greater than a known value.
	AssertionOperatorTypeGreaterThan AssertionOperatorType = "greaterThan"
	// Compare value to be less than a known value.
	AssertionOperatorTypeLessThan AssertionOperatorType = "lessThan"
	// Compare value is empty.
	AssertionOperatorTypeEmpty AssertionOperatorType = "empty"
	// Compare value is not empty.
	AssertionOperatorTypeNotEmpty AssertionOperatorType = "notEmpty"
	// Compare value string contains a known value.
	AssertionOperatorTypeContains AssertionOperatorType = "contains"
	// Compare value string does not contain a known value.
	AssertionOperatorTypeNotContains AssertionOperatorType = "notContains"
	// Evaluate the FHIRPath expression as a boolean condition.
	AssertionOperatorTypeEval AssertionOperatorType = "eval"
	// Manually evaluate the condition described by this assert. The test engine SHALL pause and provide an input mechanism to set the outcome of this assert to 'pass', 'fail', 'skip' or 'stop'.
	AssertionOperatorTypeManualEval AssertionOperatorType = "manualEval"
)
