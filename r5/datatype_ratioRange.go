// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/RatioRange
type RatioRange struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Low Numerator limit
	LowNumerator *Quantity `json:"lowNumerator,omitempty"`
	// High Numerator limit
	HighNumerator *Quantity `json:"highNumerator,omitempty"`
	// Denominator value
	Denominator *Quantity `json:"denominator,omitempty"`
}

type OtherRatioRange RatioRange

func (r RatioRange) ResourceType() string {
	return "RatioRange"
}

func (r RatioRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRatioRange
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherRatioRange: OtherRatioRange(r), ResourceType: r.ResourceType()})
}

func UnmarshalRatioRange(b []byte) (res RatioRange, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
