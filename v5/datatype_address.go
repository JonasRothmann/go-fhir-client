// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Address
type Address struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// home | work | temp | old | billing - purpose of this address
	Use *custom.Code `bson:"use,omitempty" json:"use,omitempty"`
	// postal | physical | both
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// Text representation of the address
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// Street name, number, direction & P.O. Box etc.
	Line []string `bson:"line,omitempty" json:"line,omitempty"`
	// Name of city, town etc.
	City *string `bson:"city,omitempty" json:"city,omitempty"`
	// District name (aka county)
	District *string `bson:"district,omitempty" json:"district,omitempty"`
	// Sub-unit of country (abbreviations ok)
	State *string `bson:"state,omitempty" json:"state,omitempty"`
	// Postal code for area
	PostalCode *string `bson:"postalCode,omitempty" json:"postalCode,omitempty"`
	// Country (e.g. may be ISO 3166 2 or 3 letter code)
	Country *string `bson:"country,omitempty" json:"country,omitempty"`
	// Time period when address was/is in use
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

func (a Address) ResourceType() string {
	return "StructureDefinition"
}
