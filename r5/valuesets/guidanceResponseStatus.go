// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/guidance-response-status
*/type GuidanceResponseStatus string

const (
	// The request was processed successfully.
	GuidanceResponseStatusSuccess GuidanceResponseStatus = "success"
	// The request was processed successfully, but more data may result in a more complete evaluation.
	GuidanceResponseStatusDataRequested GuidanceResponseStatus = "data-requested"
	// The request was processed, but more data is required to complete the evaluation.
	GuidanceResponseStatusDataRequired GuidanceResponseStatus = "data-required"
	// The request is currently being processed.
	GuidanceResponseStatusInProgress GuidanceResponseStatus = "in-progress"
	// The request was not processed successfully.
	GuidanceResponseStatusFailure GuidanceResponseStatus = "failure"
	// The response was entered in error.
	GuidanceResponseStatusEnteredInError GuidanceResponseStatus = "entered-in-error"
)
