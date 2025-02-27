// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/subscription-channel-type
*/type SubscriptionChannelType string

const (
	// The channel is executed by making a POST to the URI with the selected payload and MIME type.
	SubscriptionChannelTypeRestHook SubscriptionChannelType = "rest-hook"
	// The channel is executed by sending a packet across a web socket connection maintained by the client. The URL identifies the websocket, and the client binds to this URL.
	SubscriptionChannelTypeWebsocket SubscriptionChannelType = "websocket"
	// The channel is executed by sending an email to the email addressed in the URI (which must be a mailto:).
	SubscriptionChannelTypeEmail SubscriptionChannelType = "email"
	// The channel is executed by sending a message (e.g. a Bundle with a MessageHeader resource etc.) to the application identified in the URI.
	SubscriptionChannelTypeMessage SubscriptionChannelType = "message"
)
