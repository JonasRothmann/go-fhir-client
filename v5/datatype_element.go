// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/Element
type Element struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
}

func (e Element) ResourceType() string {
	return "StructureDefinition"
}
