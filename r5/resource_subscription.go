// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
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
	// Additional identifiers (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Human readable name for this subscription
	Name *string `json:"name,omitempty"`
	// requested | active | error | off | entered-in-error
	Status custom.Code `json:"status"`
	// Reference to the subscription topic being subscribed to
	Topic custom.Canonical[SubscriptionTopic] `json:"topic"`
	// Contact details for source (e.g. troubleshooting)
	Contact []ContactPoint `json:"contact,omitempty"`
	// When to automatically delete the subscription
	End *custom.Instant `json:"end,omitempty"`
	// Entity responsible for Subscription changes
	ManagingEntity *custom.Reference[SubscriptionManagingEntity] `json:"managingEntity,omitempty"`
	// Description of why this subscription was created
	Reason *string `json:"reason,omitempty"`
	// Criteria for narrowing the subscription topic stream
	FilterBy []SubscriptionFilterBy `json:"filterBy,omitempty"`
	// Channel type for notifications
	ChannelType Coding `json:"channelType"`
	// Where the channel points to
	Endpoint *custom.URL `json:"endpoint,omitempty"`
	// Channel type
	Parameter []SubscriptionParameter `json:"parameter,omitempty"`
	// Interval in seconds to send 'heartbeat' notification
	HeartbeatPeriod *uint32 `json:"heartbeatPeriod,omitempty"`
	// Timeout in seconds to attempt notification delivery
	Timeout *uint32 `json:"timeout,omitempty"`
	// MIME type to send, or omit for no payload
	ContentType *custom.Code `json:"contentType,omitempty"`
	// empty | id-only | full-resource
	Content *custom.Code `json:"content,omitempty"`
	// Maximum number of events that can be combined in a single notification
	MaxCount *uint32 `json:"maxCount,omitempty"`
}

type SubscriptionFilterBy struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Allowed Resource (reference to definition) for this Subscription filter
	ResourceType *custom.URI `json:"resourceType,omitempty"`
	// Filter label defined in SubscriptionTopic
	FilterParameter string `json:"filterParameter"`
	// eq | ne | gt | lt | ge | le | sa | eb | ap
	Comparator *custom.Code `json:"comparator,omitempty"`
	// missing | exact | contains | not | text | in | not-in | below | above | type | identifier | of-type | code-text | text-advanced | iterate
	Modifier *custom.Code `json:"modifier,omitempty"`
	// Literal value or resource path
	Value string `json:"value"`
}

type SubscriptionParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name (key) of the parameter
	Name string `json:"name"`
	// Value of the parameter to use or pass through
	Value string `json:"value"`
}

type OtherSubscription Subscription

type SubscriptionManagingEntity interface {
	gofhirclient.Resource

	Is_SubscriptionManagingEntity()
}

func (c CareTeam) Is_SubscriptionManagingEntity()          {}
func (h HealthcareService) Is_SubscriptionManagingEntity() {}
func (o Organization) Is_SubscriptionManagingEntity()      {}
func (r RelatedPerson) Is_SubscriptionManagingEntity()     {}
func (p Patient) Is_SubscriptionManagingEntity()           {}
func (p Practitioner) Is_SubscriptionManagingEntity()      {}
func (p PractitionerRole) Is_SubscriptionManagingEntity()  {}

func (s Subscription) ResourceType() string {
	return "Subscription"
}

func (s Subscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubscription
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSubscription: OtherSubscription(s), ResourceType: s.ResourceType()})
}

func UnmarshalSubscription(b []byte) (res Subscription, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
