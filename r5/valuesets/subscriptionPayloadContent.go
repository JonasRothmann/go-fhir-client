// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/subscription-payload-content
*/type SubscriptionPayloadContent string

const (
	// No resource content is transacted in the notification payload.
	SubscriptionPayloadContentEmpty SubscriptionPayloadContent = "empty"
	// Only the resource id is transacted in the notification payload.
	SubscriptionPayloadContentIdOnly SubscriptionPayloadContent = "id-only"
	// The entire resource is transacted in the notification payload.
	SubscriptionPayloadContentFullResource SubscriptionPayloadContent = "full-resource"
)
