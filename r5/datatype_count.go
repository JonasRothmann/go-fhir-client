// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Count
type Count struct {
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

type OtherCount Count

func (c Count) ResourceType() string {
	return "Count"
}

func (c Count) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCount
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCount: OtherCount(c), ResourceType: c.ResourceType()})
}

func UnmarshalCount(b []byte) (res Count, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
