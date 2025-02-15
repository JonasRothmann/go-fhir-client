// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ManufacturedItemDefinition
type ManufacturedItemDefinition struct {
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
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// A descriptive name applied to this item
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Dose form as manufactured (before any necessary transformation)
	ManufacturedDoseForm CodeableConcept `bson:"manufacturedDoseForm" json:"manufacturedDoseForm"`
	// The “real-world” units in which the quantity of the item is described
	UnitOfPresentation *CodeableConcept `bson:"unitOfPresentation,omitempty" json:"unitOfPresentation,omitempty"`
	// Manufacturer of the item, one of several possible
	Manufacturer []custom.Reference[Organization] `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// Allows specifying that an item is on the market for sale, or that it is not available, and the dates and locations associated
	MarketingStatus []MarketingStatus `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty"`
	// The ingredients of this manufactured item. Only needed if these are not specified by incoming references from the Ingredient resource
	Ingredient []CodeableConcept `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	// General characteristics of this item
	Property []ManufacturedItemDefinitionProperty `bson:"property,omitempty" json:"property,omitempty"`
	// Physical parts of the manufactured item, that it is intrisically made from. This is distinct from the ingredients that are part of its chemical makeup
	Component []ManufacturedItemDefinitionComponent `bson:"component,omitempty" json:"component,omitempty"`
}

type ManufacturedItemDefinitionProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A code expressing the type of characteristic
	Type CodeableConcept `bson:"type" json:"type"`
	// A value for the characteristic
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

type ManufacturedItemDefinitionComponent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Defining type of the component e.g. shell, layer, ink
	Type CodeableConcept `bson:"type" json:"type"`
	// The function of this component within the item e.g. delivers active ingredient, masks taste
	Function []CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// The measurable amount of total quantity of all substances in the component, expressable in different ways (e.g. by mass or volume)
	Amount []Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	// A reference to a constituent of the manufactured item as a whole, linked here so that its component location within the item can be indicated. This not where the item's ingredient are primarily stated (for which see Ingredient.for or ManufacturedItemDefinition.ingredient)
	Constituent []ManufacturedItemDefinitionComponentConstituent `bson:"constituent,omitempty" json:"constituent,omitempty"`
	// General characteristics of this component
	Property []interface{} `bson:"property,omitempty" json:"property,omitempty"`
	// A component that this component contains or is made from
	Component []interface{} `bson:"component,omitempty" json:"component,omitempty"`
}

type ManufacturedItemDefinitionComponentConstituent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The measurable amount of the substance, expressable in different ways (e.g. by mass or volume)
	Amount []Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	// The physical location of the constituent/ingredient within the component
	Location []CodeableConcept `bson:"location,omitempty" json:"location,omitempty"`
	// The function of this constituent within the component e.g. binder
	Function []CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// The ingredient that is the constituent of the given component
	HasIngredient []custom.CodeableReference[Ingredient] `bson:"hasIngredient,omitempty" json:"hasIngredient,omitempty"`
}

func (m ManufacturedItemDefinition) ResourceType() string {
	return "StructureDefinition"
}
