// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/CodeableConcept
type CodeableConcept struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Code defined by a terminology system
	Coding []Coding `json:"coding,omitempty"`
	// Plain text representation of the concept
	Text *string `json:"text,omitempty"`
}

type OtherCodeableConcept CodeableConcept

func (c CodeableConcept) ResourceType() string {
	return "CodeableConcept"
}

func (c CodeableConcept) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCodeableConcept
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCodeableConcept: OtherCodeableConcept(c), ResourceType: c.ResourceType()})
}

func UnmarshalCodeableConcept(b []byte) (res CodeableConcept, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
