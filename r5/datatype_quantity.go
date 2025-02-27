// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Quantity
type Quantity struct {
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

type OtherQuantity Quantity

func (q Quantity) ResourceType() string {
	return "Quantity"
}

func (q Quantity) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuantity
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherQuantity: OtherQuantity(q), ResourceType: q.ResourceType()})
}

func UnmarshalQuantity(b []byte) (res Quantity, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
