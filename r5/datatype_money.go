// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Money
type Money struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Numerical value (with implicit precision)
	Value *json.Number `json:"value,omitempty"`
	// ISO 4217 Currency Code
	Currency *custom.Code `json:"currency,omitempty"`
}

type OtherMoney Money

func (m Money) ResourceType() string {
	return "Money"
}

func (m Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMoney
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMoney: OtherMoney(m), ResourceType: m.ResourceType()})
}

func UnmarshalMoney(b []byte) (res Money, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
