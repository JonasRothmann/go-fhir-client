// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Extension
type Extension struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// identifies the meaning of the extension
	Url string `bson:"url" json:"url"`
	// Value of extension
	Value *custom.Base64binary `bson:"value,omitempty" json:"value,omitempty"`
}

func (e Extension) ResourceType() string {
	return "StructureDefinition"
}
