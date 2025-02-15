// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/RatioRange
type RatioRange struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Low Numerator limit
	LowNumerator *Quantity `bson:"lowNumerator,omitempty" json:"lowNumerator,omitempty"`
	// High Numerator limit
	HighNumerator *Quantity `bson:"highNumerator,omitempty" json:"highNumerator,omitempty"`
	// Denominator value
	Denominator *Quantity `bson:"denominator,omitempty" json:"denominator,omitempty"`
}

func (r RatioRange) ResourceType() string {
	return "StructureDefinition"
}
