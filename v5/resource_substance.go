// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Substance
type SubstanceIngredient struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Optional amount (concentration)
	Quantity *Ratio `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// A component of the substance
	Substance CodeableConcept `bson:"substance" json:"substance"`
}

type Substance struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Unique identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Is this an instance of a substance or a kind of one
	Instance bool `bson:"instance" json:"instance"`
	// active | inactive | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// What class/type of substance this is
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// What substance this is
	Code custom.CodeableReference[SubstanceDefinition] `bson:"code" json:"code"`
	// Textual description of the substance, comments
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// When no longer valid to use
	Expiry *custom.DateTime `bson:"expiry,omitempty" json:"expiry,omitempty"`
	// Amount of substance in the package
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Composition information about the substance
	Ingredient []SubstanceIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
}

func (s Substance) ResourceType() string {
	return "StructureDefinition"
}
