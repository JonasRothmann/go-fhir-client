// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/PrimitiveType
type PrimitiveType struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
}

type OtherPrimitiveType PrimitiveType

func (p PrimitiveType) ResourceType() string {
	return "PrimitiveType"
}

func (p PrimitiveType) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPrimitiveType
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPrimitiveType: OtherPrimitiveType(p), ResourceType: p.ResourceType()})
}

func UnmarshalPrimitiveType(b []byte) (res PrimitiveType, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
