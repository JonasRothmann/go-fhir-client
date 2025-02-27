// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/Range
type Range struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Low limit
	Low *Quantity `json:"low,omitempty"`
	// High limit
	High *Quantity `json:"high,omitempty"`
}

type OtherRange Range

func (r Range) ResourceType() string {
	return "Range"
}

func (r Range) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRange
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherRange: OtherRange(r), ResourceType: r.ResourceType()})
}

func UnmarshalRange(b []byte) (res Range, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
