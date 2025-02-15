// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SubscriptionStatus
type SubscriptionStatus struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// requested | active | error | off | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// handshake | heartbeat | event-notification | query-status | query-event
	Type custom.Code `bson:"type" json:"type"`
	// Events since the Subscription was created
	EventsSinceSubscriptionStart *int64 `bson:"eventsSinceSubscriptionStart,omitempty" json:"eventsSinceSubscriptionStart,omitempty"`
	// Detailed information about any events relevant to this notification
	NotificationEvent []SubscriptionStatusNotificationEvent `bson:"notificationEvent,omitempty" json:"notificationEvent,omitempty"`
	// Reference to the Subscription responsible for this notification
	Subscription custom.Reference[Subscription] `bson:"subscription" json:"subscription"`
	// Reference to the SubscriptionTopic this notification relates to
	Topic *custom.Canonical[SubscriptionTopic] `bson:"topic,omitempty" json:"topic,omitempty"`
	// List of errors on the subscription
	Error []CodeableConcept `bson:"error,omitempty" json:"error,omitempty"`
}

type SubscriptionStatusNotificationEvent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Sequencing index of this event
	EventNumber int64 `bson:"eventNumber" json:"eventNumber"`
	// The instant this event occurred
	Timestamp *custom.Instant `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	// Reference to the primary resource or information of this event
	Focus *custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// References related to the focus resource and/or context of this event
	AdditionalContext []custom.Reference[Resource] `bson:"additionalContext,omitempty" json:"additionalContext,omitempty"`
}

func (s SubscriptionStatus) ResourceType() string {
	return "StructureDefinition"
}
