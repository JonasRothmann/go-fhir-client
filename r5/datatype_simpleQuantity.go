// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SimpleQuantity
type SimpleQuantity struct {
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

type OtherSimpleQuantity SimpleQuantity

func (s SimpleQuantity) ResourceType() string {
	return "SimpleQuantity"
}

func (s SimpleQuantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSimpleQuantity
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSimpleQuantity: OtherSimpleQuantity(s), ResourceType: s.ResourceType()})
}

func UnmarshalSimpleQuantity(b []byte) (res SimpleQuantity, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
