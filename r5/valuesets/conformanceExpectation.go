// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/conformance-expectation
*/type ConformanceExpectation string

const (
	// Support for the specified capability is required to be considered conformant.
	ConformanceExpectationShall ConformanceExpectation = "SHALL"
	// Support for the specified capability is strongly encouraged, and failure to support it should only occur after careful consideration.
	ConformanceExpectationShould ConformanceExpectation = "SHOULD"
	// Support for the specified capability is not necessary to be considered conformant, and the requirement should be considered strictly optional.
	ConformanceExpectationMay ConformanceExpectation = "MAY"
	// Support for the specified capability is strongly discouraged and should occur only after careful consideration.
	ConformanceExpectationShouldnot ConformanceExpectation = "SHOULD-NOT"
)
