// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Coding
type Coding struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Identity of the terminology system
	System *custom.URI `json:"system,omitempty"`
	// Version of the system - if relevant
	Version *string `json:"version,omitempty"`
	// Symbol in syntax defined by the system
	Code *custom.Code `json:"code,omitempty"`
	// Representation defined by the system
	Display *string `json:"display,omitempty"`
	// If this coding was chosen directly by the user
	UserSelected *bool `json:"userSelected,omitempty"`
}

type OtherCoding Coding

func (c Coding) ResourceType() string {
	return "Coding"
}

func (c Coding) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoding
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCoding: OtherCoding(c), ResourceType: c.ResourceType()})
}

func UnmarshalCoding(b []byte) (res Coding, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
