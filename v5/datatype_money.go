// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Money
type Money struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Numerical value (with implicit precision)
	Value *json.Number `bson:"value,omitempty" json:"value,omitempty"`
	// ISO 4217 Currency Code
	Currency *custom.Code `bson:"currency,omitempty" json:"currency,omitempty"`
}

func (m Money) ResourceType() string {
	return "StructureDefinition"
}
