// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Narrative
type Narrative struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// generated | extensions | additional | empty
	Status custom.Code `bson:"status" json:"status"`
	// Limited xhtml content
	Div custom.XHTML `bson:"div" json:"div"`
}

func (n Narrative) ResourceType() string {
	return "StructureDefinition"
}
