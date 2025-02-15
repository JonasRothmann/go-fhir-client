// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ContactPoint
type ContactPoint struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// phone | fax | email | pager | url | sms | other
	System *custom.Code `bson:"system,omitempty" json:"system,omitempty"`
	// The actual contact point details
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
	// home | work | temp | old | mobile - purpose of this contact point
	Use *custom.Code `bson:"use,omitempty" json:"use,omitempty"`
	// Specify preferred order of use (1 = highest)
	Rank *uint32 `bson:"rank,omitempty" json:"rank,omitempty"`
	// Time period when the contact point was/is in use
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

func (c ContactPoint) ResourceType() string {
	return "StructureDefinition"
}
