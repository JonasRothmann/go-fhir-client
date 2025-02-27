// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SubscriptionStatus
type SubscriptionStatus struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// requested | active | error | off | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// handshake | heartbeat | event-notification | query-status | query-event
	Type custom.Code `json:"type"`
	// Events since the Subscription was created
	EventsSinceSubscriptionStart *int64 `json:"eventsSinceSubscriptionStart,omitempty"`
	// Detailed information about any events relevant to this notification
	NotificationEvent []SubscriptionStatusNotificationEvent `json:"notificationEvent,omitempty"`
	// Reference to the Subscription responsible for this notification
	Subscription custom.Reference[Subscription] `json:"subscription"`
	// Reference to the SubscriptionTopic this notification relates to
	Topic *custom.Canonical[SubscriptionTopic] `json:"topic,omitempty"`
	// List of errors on the subscription
	Error []CodeableConcept `json:"error,omitempty"`
}

type SubscriptionStatusNotificationEvent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Sequencing index of this event
	EventNumber int64 `json:"eventNumber"`
	// The instant this event occurred
	Timestamp *custom.Instant `json:"timestamp,omitempty"`
	// Reference to the primary resource or information of this event
	Focus *custom.Reference[Resource] `json:"focus,omitempty"`
	// References related to the focus resource and/or context of this event
	AdditionalContext []custom.Reference[Resource] `json:"additionalContext,omitempty"`
}

type OtherSubscriptionStatus SubscriptionStatus

func (s SubscriptionStatus) ResourceType() string {
	return "SubscriptionStatus"
}

func (s SubscriptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscriptionStatus
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSubscriptionStatus: OtherSubscriptionStatus(s), ResourceType: s.ResourceType()})
}

func UnmarshalSubscriptionStatus(b []byte) (res SubscriptionStatus, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
