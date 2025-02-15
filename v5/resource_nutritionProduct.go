// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/NutritionProduct
type NutritionProduct struct {
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
	// A code that can identify the detailed nutrients and ingredients in a specific food product
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// active | inactive | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Broad product groups or categories used to classify the product, such as Legume and Legume Products, Beverages, or Beef Products
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Manufacturer, representative or officially responsible for the product
	Manufacturer []custom.Reference[Organization] `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// The product's nutritional information expressed by the nutrients
	Nutrient []NutritionProductNutrient `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	// Ingredients contained in this product
	Ingredient []NutritionProductIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	// Known or suspected allergens that are a part of this product
	KnownAllergen []custom.CodeableReference[Substance] `bson:"knownAllergen,omitempty" json:"knownAllergen,omitempty"`
	// Specifies descriptive properties of the nutrition product
	Characteristic []NutritionProductCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	// One or several physical instances or occurrences of the nutrition product
	Instance []NutritionProductInstance `bson:"instance,omitempty" json:"instance,omitempty"`
	// Comments made about the product
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type NutritionProductNutrient struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The (relevant) nutrients in the product
	Item *custom.CodeableReference[Substance] `bson:"item,omitempty" json:"item,omitempty"`
	// The amount of nutrient expressed in one or more units: X per pack / per serving / per dose
	Amount []Ratio `bson:"amount,omitempty" json:"amount,omitempty"`
}

type NutritionProductIngredient struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The ingredient contained in the product
	Item custom.CodeableReference[NutritionProduct] `bson:"item" json:"item"`
	// The amount of ingredient that is in the product
	Amount []Ratio `bson:"amount,omitempty" json:"amount,omitempty"`
}

type NutritionProductCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code specifying the type of characteristic
	Type CodeableConcept `bson:"type" json:"type"`
	// The value of the characteristic
	Value CodeableConcept `bson:"value" json:"value"`
}

type NutritionProductInstance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The amount of items or instances
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// The identifier for the physical instance, typically a serial number or manufacturer number
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The name for the specific product
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// The identification of the batch or lot of the product
	LotNumber *string `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	// The expiry date or date and time for the product
	Expiry *custom.DateTime `bson:"expiry,omitempty" json:"expiry,omitempty"`
	// The date until which the product is expected to be good for consumption
	UseBy *custom.DateTime `bson:"useBy,omitempty" json:"useBy,omitempty"`
	// An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled
	BiologicalSourceEvent *Identifier `bson:"biologicalSourceEvent,omitempty" json:"biologicalSourceEvent,omitempty"`
}

func (n NutritionProduct) ResourceType() string {
	return "StructureDefinition"
}
