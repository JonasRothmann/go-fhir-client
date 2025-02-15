// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Expression
type Expression struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Natural language description of the condition
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Short name assigned to expression for reuse
	Name *custom.Code `bson:"name,omitempty" json:"name,omitempty"`
	// text/cql | text/fhirpath | application/x-fhir-query | etc.
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Expression in specified language
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty"`
	// Where the expression is found
	Reference *custom.URI `bson:"reference,omitempty" json:"reference,omitempty"`
}

func (e Expression) ResourceType() string {
	return "StructureDefinition"
}
