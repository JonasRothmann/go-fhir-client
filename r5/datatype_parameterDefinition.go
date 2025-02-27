// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ParameterDefinition
type ParameterDefinition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Name used to access the parameter value
	Name *custom.Code `json:"name,omitempty"`
	// in | out
	Use custom.Code `json:"use"`
	// Minimum cardinality
	Min *int32 `json:"min,omitempty"`
	// Maximum cardinality (a number of *)
	Max *string `json:"max,omitempty"`
	// A brief description of the parameter
	Documentation *string `json:"documentation,omitempty"`
	// What type of value
	Type custom.Code `json:"type"`
	// What profile the value is expected to be
	Profile *custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
}

type OtherParameterDefinition ParameterDefinition

func (p ParameterDefinition) ResourceType() string {
	return "ParameterDefinition"
}

func (p ParameterDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherParameterDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherParameterDefinition: OtherParameterDefinition(p), ResourceType: p.ResourceType()})
}

func UnmarshalParameterDefinition(b []byte) (res ParameterDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
