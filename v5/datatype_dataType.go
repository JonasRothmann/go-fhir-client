// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/DataType
type DataType struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
}

func (d DataType) ResourceType() string {
	return "StructureDefinition"
}
