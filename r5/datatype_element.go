// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/Element
type Element struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
}

type OtherElement Element

func (e Element) ResourceType() string {
	return "Element"
}

func (e Element) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherElement
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherElement: OtherElement(e), ResourceType: e.ResourceType()})
}

func UnmarshalElement(b []byte) (res Element, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
