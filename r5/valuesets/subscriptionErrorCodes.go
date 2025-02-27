// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/subscription-error
*/type SubscriptionErrorCodes string

const (
	// DNS resolution error for client rest-hook endpoint.
	SubscriptionErrorCodesDnsResolutionError SubscriptionErrorCodes = "dns-resolution-error"
	// No response for client rest-hook endpoint.
	SubscriptionErrorCodesNoResponse SubscriptionErrorCodes = "no-response"
	// Error response from client rest-hook endpoint. Recommended practice: convey the actual HTTP error in an adjacent Coding or in text
	SubscriptionErrorCodesErrorResponse SubscriptionErrorCodes = "error-response"
)
