// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Subscription
type Subscription struct {
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
	// Additional identifiers (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Human readable name for this subscription
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// requested | active | error | off | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Reference to the subscription topic being subscribed to
	Topic custom.Canonical[SubscriptionTopic] `bson:"topic" json:"topic"`
	// Contact details for source (e.g. troubleshooting)
	Contact []ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	// When to automatically delete the subscription
	End *custom.Instant `bson:"end,omitempty" json:"end,omitempty"`
	// Entity responsible for Subscription changes
	ManagingEntity *custom.Reference[SubscriptionManagingEntity] `bson:"managingEntity,omitempty" json:"managingEntity,omitempty"`
	// Description of why this subscription was created
	Reason *string `bson:"reason,omitempty" json:"reason,omitempty"`
	// Criteria for narrowing the subscription topic stream
	FilterBy []SubscriptionFilterBy `bson:"filterBy,omitempty" json:"filterBy,omitempty"`
	// Channel type for notifications
	ChannelType Coding `bson:"channelType" json:"channelType"`
	// Where the channel points to
	Endpoint *custom.URL `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// Channel type
	Parameter []SubscriptionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	// Interval in seconds to send 'heartbeat' notification
	HeartbeatPeriod *uint32 `bson:"heartbeatPeriod,omitempty" json:"heartbeatPeriod,omitempty"`
	// Timeout in seconds to attempt notification delivery
	Timeout *uint32 `bson:"timeout,omitempty" json:"timeout,omitempty"`
	// MIME type to send, or omit for no payload
	ContentType *custom.Code `bson:"contentType,omitempty" json:"contentType,omitempty"`
	// empty | id-only | full-resource
	Content *custom.Code `bson:"content,omitempty" json:"content,omitempty"`
	// Maximum number of events that can be combined in a single notification
	MaxCount *uint32 `bson:"maxCount,omitempty" json:"maxCount,omitempty"`
}

type SubscriptionFilterBy struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Allowed Resource (reference to definition) for this Subscription filter
	ResourceType *custom.URI `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	// Filter label defined in SubscriptionTopic
	FilterParameter string `bson:"filterParameter" json:"filterParameter"`
	// eq | ne | gt | lt | ge | le | sa | eb | ap
	Comparator *custom.Code `bson:"comparator,omitempty" json:"comparator,omitempty"`
	// missing | exact | contains | not | text | in | not-in | below | above | type | identifier | of-type | code-text | text-advanced | iterate
	Modifier *custom.Code `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Literal value or resource path
	Value string `bson:"value" json:"value"`
}

type SubscriptionParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name (key) of the parameter
	Name string `bson:"name" json:"name"`
	// Value of the parameter to use or pass through
	Value string `bson:"value" json:"value"`
}

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
	return "StructureDefinition"
}
