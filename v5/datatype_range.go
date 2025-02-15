// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/Range
type Range struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Low limit
	Low *Quantity `bson:"low,omitempty" json:"low,omitempty"`
	// High limit
	High *Quantity `bson:"high,omitempty" json:"high,omitempty"`
}

func (r Range) ResourceType() string {
	return "StructureDefinition"
}
