// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/UsageContext
type UsageContext struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Type of context being specified
	Code Coding `bson:"code" json:"code"`
	// Value that defines the context
	Value CodeableConcept `bson:"value" json:"value"`
}

func (u UsageContext) ResourceType() string {
	return "StructureDefinition"
}
