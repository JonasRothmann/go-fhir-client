// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/CodeableConcept
type CodeableConcept struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Code defined by a terminology system
	Coding []Coding `bson:"coding,omitempty" json:"coding,omitempty"`
	// Plain text representation of the concept
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
}

func (c CodeableConcept) ResourceType() string {
	return "StructureDefinition"
}
