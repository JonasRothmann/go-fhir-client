// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/endpoint-status
*/type EndpointStatus string

const (
	// This endpoint is expected to be active and can be used.
	EndpointStatusActive EndpointStatus = "active"
	// This endpoint is temporarily unavailable.
	EndpointStatusSuspended EndpointStatus = "suspended"
	// This endpoint has exceeded connectivity thresholds and is considered in an error state and should no longer be attempted to connect to until corrective action is taken.
	EndpointStatusError EndpointStatus = "error"
	// This endpoint is no longer to be used.
	EndpointStatusOff EndpointStatus = "off"
	// This instance should not have been part of this patient's medical record.
	EndpointStatusEnteredInError EndpointStatus = "entered-in-error"
)
