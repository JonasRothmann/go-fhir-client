// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/Ratio
type Ratio struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Numerator value
	Numerator *Quantity `bson:"numerator,omitempty" json:"numerator,omitempty"`
	// Denominator value
	Denominator *Quantity `bson:"denominator,omitempty" json:"denominator,omitempty"`
}

func (r Ratio) ResourceType() string {
	return "StructureDefinition"
}
