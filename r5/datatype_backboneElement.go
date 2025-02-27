// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/BackboneElement
type BackboneElement struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
}

type OtherBackboneElement BackboneElement

func (b BackboneElement) ResourceType() string {
	return "BackboneElement"
}

func (b BackboneElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBackboneElement
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherBackboneElement: OtherBackboneElement(b), ResourceType: b.ResourceType()})
}

func UnmarshalBackboneElement(b []byte) (res BackboneElement, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
