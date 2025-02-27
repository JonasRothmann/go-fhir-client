// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/endpoint-payload-type
*/type EndpointPayloadType string

const (
	// Any payload type can be used with this endpoint, it is either a payload agnostic infrastructure (such as a storage repository), or some other type of endpoint where payload considerations are internally handled, and not available
	EndpointPayloadTypeAny EndpointPayloadType = "any"
	// This endpoint does not require any content to be sent; simply connecting to the endpoint is enough notification. This can be used as a 'ping' to wakeup a service to retrieve content, which could be to ensure security considerations are correctly handled
	EndpointPayloadTypeNone EndpointPayloadType = "none"
)
