// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/BackboneType
type BackboneType struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

type OtherBackboneType BackboneType

func (b BackboneType) ResourceType() string {
	return "BackboneType"
}

func (b BackboneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBackboneType
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherBackboneType: OtherBackboneType(b), ResourceType: b.ResourceType()})
}

func UnmarshalBackboneType(b []byte) (res BackboneType, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
