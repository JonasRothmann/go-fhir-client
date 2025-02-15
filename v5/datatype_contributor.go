// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Contributor
type Contributor struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// author | editor | reviewer | endorser
	Type custom.Code `bson:"type" json:"type"`
	// Who contributed the content
	Name string `bson:"name" json:"name"`
	// Contact details of the contributor
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
}

func (c Contributor) ResourceType() string {
	return "StructureDefinition"
}
