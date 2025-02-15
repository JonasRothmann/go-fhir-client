// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ParameterDefinition
type ParameterDefinition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Name used to access the parameter value
	Name *custom.Code `bson:"name,omitempty" json:"name,omitempty"`
	// in | out
	Use custom.Code `bson:"use" json:"use"`
	// Minimum cardinality
	Min *int32 `bson:"min,omitempty" json:"min,omitempty"`
	// Maximum cardinality (a number of *)
	Max *string `bson:"max,omitempty" json:"max,omitempty"`
	// A brief description of the parameter
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// What type of value
	Type custom.Code `bson:"type" json:"type"`
	// What profile the value is expected to be
	Profile *custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
}

func (p ParameterDefinition) ResourceType() string {
	return "StructureDefinition"
}
