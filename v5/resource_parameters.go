// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Parameters
type Parameters struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Operation Parameter
	Parameter []ParametersParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}

type ParametersParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name from the definition
	Name string `bson:"name" json:"name"`
	// If parameter is a data type
	Value *custom.Base64binary `bson:"value,omitempty" json:"value,omitempty"`
	// If parameter is a whole resource
	Resource *Resource `bson:"resource,omitempty" json:"resource,omitempty"`
	// Named part of a multi-part parameter
	Part []interface{} `bson:"part,omitempty" json:"part,omitempty"`
}

func (p Parameters) ResourceType() string {
	return "StructureDefinition"
}
