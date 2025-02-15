// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Period
type Period struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Starting time with inclusive boundary
	Start *custom.DateTime `bson:"start,omitempty" json:"start,omitempty"`
	// End time with inclusive boundary, if not ongoing
	End *custom.DateTime `bson:"end,omitempty" json:"end,omitempty"`
}

func (p Period) ResourceType() string {
	return "StructureDefinition"
}
