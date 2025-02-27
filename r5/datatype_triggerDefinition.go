// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/TriggerDefinition
type TriggerDefinition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// named-event | periodic | data-changed | data-added | data-modified | data-removed | data-accessed | data-access-ended
	Type custom.Code `json:"type"`
	// Name or URI that identifies the event
	Name *string `json:"name,omitempty"`
	// Coded definition of the event
	Code *CodeableConcept `json:"code,omitempty"`
	// What event
	SubscriptionTopic *custom.Canonical[SubscriptionTopic] `json:"subscriptionTopic,omitempty"`
	// Timing of the event
	TimingTiming *Timing `json:"timingTiming,omitempty"`
	// Timing of the event
	TimingReference *custom.Reference[Schedule] `json:"timingReference,omitempty"`
	// Timing of the event
	TimingDate *custom.Date `json:"timingDate,omitempty"`
	// Timing of the event
	TimingDateTime *custom.DateTime `json:"timingDateTime,omitempty"`
	// Triggering data of the event (multiple = 'and')
	Data []DataRequirement `json:"data,omitempty"`
	// Whether the event triggers (boolean expression)
	Condition *Expression `json:"condition,omitempty"`
}

type OtherTriggerDefinition TriggerDefinition

func (t TriggerDefinition) ResourceType() string {
	return "TriggerDefinition"
}

func (t TriggerDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTriggerDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherTriggerDefinition: OtherTriggerDefinition(t), ResourceType: t.ResourceType()})
}

func UnmarshalTriggerDefinition(b []byte) (res TriggerDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
