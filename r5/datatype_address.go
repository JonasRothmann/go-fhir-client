// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Address
type Address struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// home | work | temp | old | billing - purpose of this address
	Use *custom.Code `json:"use,omitempty"`
	// postal | physical | both
	Type *custom.Code `json:"type,omitempty"`
	// Text representation of the address
	Text *string `json:"text,omitempty"`
	// Street name, number, direction & P.O. Box etc.
	Line []string `json:"line,omitempty"`
	// Name of city, town etc.
	City *string `json:"city,omitempty"`
	// District name (aka county)
	District *string `json:"district,omitempty"`
	// Sub-unit of country (abbreviations ok)
	State *string `json:"state,omitempty"`
	// Postal code for area
	PostalCode *string `json:"postalCode,omitempty"`
	// Country (e.g. may be ISO 3166 2 or 3 letter code)
	Country *string `json:"country,omitempty"`
	// Time period when address was/is in use
	Period *Period `json:"period,omitempty"`
}

type OtherAddress Address

func (a Address) ResourceType() string {
	return "Address"
}

func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAddress
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAddress: OtherAddress(a), ResourceType: a.ResourceType()})
}

func UnmarshalAddress(b []byte) (res Address, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
