// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/BackboneType
type BackboneType struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
}

func (b BackboneType) ResourceType() string {
	return "StructureDefinition"
}
