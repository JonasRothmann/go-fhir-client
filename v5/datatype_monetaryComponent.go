// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MonetaryComponent
type MonetaryComponent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// base | surcharge | deduction | discount | tax | informational
	Type custom.Code `bson:"type" json:"type"`
	// Codes may be used to differentiate between kinds of taxes, surcharges, discounts etc.
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Factor used for calculating this component
	Factor *json.Number `bson:"factor,omitempty" json:"factor,omitempty"`
	// Explicit value amount to be used
	Amount *Money `bson:"amount,omitempty" json:"amount,omitempty"`
}

func (m MonetaryComponent) ResourceType() string {
	return "StructureDefinition"
}
