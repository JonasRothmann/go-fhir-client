// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// named-event | periodic | data-changed | data-added | data-modified | data-removed | data-accessed | data-access-ended
	Type custom.Code `bson:"type" json:"type"`
	// Name or URI that identifies the event
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Coded definition of the event
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// What event
	SubscriptionTopic *custom.Canonical[SubscriptionTopic] `bson:"subscriptionTopic,omitempty" json:"subscriptionTopic,omitempty"`
	// Timing of the event
	Timing *Timing `bson:"timing,omitempty" json:"timing,omitempty"`
	// Triggering data of the event (multiple = 'and')
	Data []DataRequirement `bson:"data,omitempty" json:"data,omitempty"`
	// Whether the event triggers (boolean expression)
	Condition *Expression `bson:"condition,omitempty" json:"condition,omitempty"`
}

func (t TriggerDefinition) ResourceType() string {
	return "StructureDefinition"
}
