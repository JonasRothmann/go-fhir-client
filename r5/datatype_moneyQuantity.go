// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MoneyQuantity
type MoneyQuantity struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Numerical value (with implicit precision)
	Value *json.Number `json:"value,omitempty"`
	// < | <= | >= | > | ad - how to understand the value
	Comparator *custom.Code `json:"comparator,omitempty"`
	// Unit representation
	Unit *string `json:"unit,omitempty"`
	// System that defines coded unit form
	System *custom.URI `json:"system,omitempty"`
	// Coded form of the unit
	Code *custom.Code `json:"code,omitempty"`
}

type OtherMoneyQuantity MoneyQuantity

func (m MoneyQuantity) ResourceType() string {
	return "MoneyQuantity"
}

func (m MoneyQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMoneyQuantity
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMoneyQuantity: OtherMoneyQuantity(m), ResourceType: m.ResourceType()})
}

func UnmarshalMoneyQuantity(b []byte) (res MoneyQuantity, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
