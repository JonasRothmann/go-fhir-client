// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/ContactDetail
type ContactDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Name of an individual to contact
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Contact details for individual or organization
	Telecom []ContactPoint `bson:"telecom,omitempty" json:"telecom,omitempty"`
}

func (c ContactDetail) ResourceType() string {
	return "StructureDefinition"
}
