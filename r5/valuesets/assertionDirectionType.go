// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/assert-direction-codes
*/type AssertionDirectionType string

const (
	// The assertion is evaluated on the response. This is the default value.
	AssertionDirectionTypeResponse AssertionDirectionType = "response"
	// The assertion is evaluated on the request.
	AssertionDirectionTypeRequest AssertionDirectionType = "request"
)
