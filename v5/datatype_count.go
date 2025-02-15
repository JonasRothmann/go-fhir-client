// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Count
type Count struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Numerical value (with implicit precision)
	Value *json.Number `bson:"value,omitempty" json:"value,omitempty"`
	// < | <= | >= | > | ad - how to understand the value
	Comparator *custom.Code `bson:"comparator,omitempty" json:"comparator,omitempty"`
	// Unit representation
	Unit *string `bson:"unit,omitempty" json:"unit,omitempty"`
	// System that defines coded unit form
	System *custom.URI `bson:"system,omitempty" json:"system,omitempty"`
	// Coded form of the unit
	Code *custom.Code `bson:"code,omitempty" json:"code,omitempty"`
}

func (c Count) ResourceType() string {
	return "StructureDefinition"
}
