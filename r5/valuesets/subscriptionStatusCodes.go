// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/subscription-status
*/type SubscriptionStatusCodes string

const (
	// The client has requested the subscription, and the server has not yet set it up.
	SubscriptionStatusCodesRequested SubscriptionStatusCodes = "requested"
	// The subscription is active.
	SubscriptionStatusCodesActive SubscriptionStatusCodes = "active"
	// The server has an error executing the notification.
	SubscriptionStatusCodesError SubscriptionStatusCodes = "error"
	// Too many errors have occurred or the subscription has expired.
	SubscriptionStatusCodesOff SubscriptionStatusCodes = "off"
	// This subscription has been flagged as incorrect.
	SubscriptionStatusCodesEnteredInError SubscriptionStatusCodes = "entered-in-error"
)
