// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/DataType
type DataType struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
}

type OtherDataType DataType

func (d DataType) ResourceType() string {
	return "DataType"
}

func (d DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDataType
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDataType: OtherDataType(d), ResourceType: d.ResourceType()})
}

func UnmarshalDataType(b []byte) (res DataType, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
