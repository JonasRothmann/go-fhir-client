// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MonetaryComponent
type MonetaryComponent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// base | surcharge | deduction | discount | tax | informational
	Type custom.Code `json:"type"`
	// Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Code *CodeableConcept `json:"code,omitempty"`
	// Factor used for calculating this component
	Factor *json.Number `json:"factor,omitempty"`
	// Explicit value amount to be used
	Amount *Money `json:"amount,omitempty"`
}

type OtherMonetaryComponent MonetaryComponent

func (m MonetaryComponent) ResourceType() string {
	return "MonetaryComponent"
}

func (m MonetaryComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMonetaryComponent
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMonetaryComponent: OtherMonetaryComponent(m), ResourceType: m.ResourceType()})
}

func UnmarshalMonetaryComponent(b []byte) (res MonetaryComponent, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
