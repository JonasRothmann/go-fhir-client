// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/Ratio
type Ratio struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Numerator value
	Numerator *Quantity `json:"numerator,omitempty"`
	// Denominator value
	Denominator *Quantity `json:"denominator,omitempty"`
}

type OtherRatio Ratio

func (r Ratio) ResourceType() string {
	return "Ratio"
}

func (r Ratio) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRatio
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherRatio: OtherRatio(r), ResourceType: r.ResourceType()})
}

func UnmarshalRatio(b []byte) (res Ratio, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
