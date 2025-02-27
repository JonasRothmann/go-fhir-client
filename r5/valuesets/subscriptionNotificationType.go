// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/subscription-notification-type
*/type SubscriptionNotificationType string

const (
	// The status was generated as part of the setup or verification of a communications channel.
	SubscriptionNotificationTypeHandshake SubscriptionNotificationType = "handshake"
	// The status was generated to perform a heartbeat notification to the subscriber.
	SubscriptionNotificationTypeHeartbeat SubscriptionNotificationType = "heartbeat"
	// The status was generated for an event to the subscriber.
	SubscriptionNotificationTypeEventNotification SubscriptionNotificationType = "event-notification"
	// The status was generated in response to a status query/request.
	SubscriptionNotificationTypeQueryStatus SubscriptionNotificationType = "query-status"
	// The status was generated in response to an event query/request.
	SubscriptionNotificationTypeQueryEvent SubscriptionNotificationType = "query-event"
)
