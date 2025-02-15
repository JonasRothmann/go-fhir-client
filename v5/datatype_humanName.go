// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/HumanName
type HumanName struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// usual | official | temp | nickname | anonymous | old | maiden
	Use *custom.Code `bson:"use,omitempty" json:"use,omitempty"`
	// Text representation of the full name
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// Family name (often called 'Surname')
	Family *string `bson:"family,omitempty" json:"family,omitempty"`
	// Given names (not always 'first'). Includes middle names
	Given []string `bson:"given,omitempty" json:"given,omitempty"`
	// Parts that come before the name
	Prefix []string `bson:"prefix,omitempty" json:"prefix,omitempty"`
	// Parts that come after the name
	Suffix []string `bson:"suffix,omitempty" json:"suffix,omitempty"`
	// Time period when name was/is in use
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

func (h HumanName) ResourceType() string {
	return "StructureDefinition"
}
