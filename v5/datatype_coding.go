// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Coding
type Coding struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Identity of the terminology system
	System *custom.URI `bson:"system,omitempty" json:"system,omitempty"`
	// Version of the system - if relevant
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// Symbol in syntax defined by the system
	Code *custom.Code `bson:"code,omitempty" json:"code,omitempty"`
	// Representation defined by the system
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// If this coding was chosen directly by the user
	UserSelected *bool `bson:"userSelected,omitempty" json:"userSelected,omitempty"`
}

func (c Coding) ResourceType() string {
	return "StructureDefinition"
}
