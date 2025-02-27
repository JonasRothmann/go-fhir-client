// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/HumanName
type HumanName struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// usual | official | temp | nickname | anonymous | old | maiden
	Use *custom.Code `json:"use,omitempty"`
	// Text representation of the full name
	Text *string `json:"text,omitempty"`
	// Family name (often called 'Surname')
	Family *string `json:"family,omitempty"`
	// Given names (not always 'first'). Includes middle names
	Given []string `json:"given,omitempty"`
	// Parts that come before the name
	Prefix []string `json:"prefix,omitempty"`
	// Parts that come after the name
	Suffix []string `json:"suffix,omitempty"`
	// Time period when name was/is in use
	Period *Period `json:"period,omitempty"`
}

type OtherHumanName HumanName

func (h HumanName) ResourceType() string {
	return "HumanName"
}

func (h HumanName) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherHumanName
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherHumanName: OtherHumanName(h), ResourceType: h.ResourceType()})
}

func UnmarshalHumanName(b []byte) (res HumanName, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
