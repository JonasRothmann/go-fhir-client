// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Substance
type SubstanceIngredient struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Optional amount (concentration)
	Quantity *Ratio `json:"quantity,omitempty"`
	// A component of the substance
	SubstanceCodeableConcept *CodeableConcept `json:"substanceCodeableConcept,omitempty"`
	// A component of the substance
	SubstanceReference *custom.Reference[Substance] `json:"substanceReference,omitempty"`
}

type Substance struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Unique identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Is this an instance of a substance or a kind of one
	Instance bool `json:"instance"`
	// active | inactive | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// What class/type of substance this is
	Category []CodeableConcept `json:"category,omitempty"`
	// What substance this is
	Code custom.CodeableReference[SubstanceDefinition] `json:"code"`
	// Textual description of the substance, comments
	Description *custom.Markdown `json:"description,omitempty"`
	// When no longer valid to use
	Expiry *custom.DateTime `json:"expiry,omitempty"`
	// Amount of substance in the package
	Quantity *Quantity `json:"quantity,omitempty"`
	// Composition information about the substance
	Ingredient []SubstanceIngredient `json:"ingredient,omitempty"`
}

type OtherSubstance Substance

func (s Substance) ResourceType() string {
	return "Substance"
}

func (s Substance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstance
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSubstance: OtherSubstance(s), ResourceType: s.ResourceType()})
}

func UnmarshalSubstance(b []byte) (res Substance, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
