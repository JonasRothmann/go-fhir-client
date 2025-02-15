// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Identifier
type Identifier struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// usual | official | temp | secondary | old (If known)
	Use *custom.Code `bson:"use,omitempty" json:"use,omitempty"`
	// Description of identifier
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The namespace for the identifier value
	System *custom.URI `bson:"system,omitempty" json:"system,omitempty"`
	// The value that is unique
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
	// Time period when id is/was valid for use
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Organization that issued id (may be just text)
	Assigner *custom.Reference[Organization] `bson:"assigner,omitempty" json:"assigner,omitempty"`
}

func (i Identifier) ResourceType() string {
	return "StructureDefinition"
}
