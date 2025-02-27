// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/NutritionProduct
type NutritionProduct struct {
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
	// A code that can identify the detailed nutrients and ingredients in a specific food product
	Code *CodeableConcept `json:"code,omitempty"`
	// active | inactive | entered-in-error
	Status custom.Code `json:"status"`
	// Broad product groups or categories used to classify the product, such as Legume and Legume Products, Beverages, or Beef Products
	Category []CodeableConcept `json:"category,omitempty"`
	// Manufacturer, representative or officially responsible for the product
	Manufacturer []custom.Reference[Organization] `json:"manufacturer,omitempty"`
	// The product's nutritional information expressed by the nutrients
	Nutrient []NutritionProductNutrient `json:"nutrient,omitempty"`
	// Ingredients contained in this product
	Ingredient []NutritionProductIngredient `json:"ingredient,omitempty"`
	// Known or suspected allergens that are a part of this product
	KnownAllergen []custom.CodeableReference[Substance] `json:"knownAllergen,omitempty"`
	// Specifies descriptive properties of the nutrition product
	Characteristic []NutritionProductCharacteristic `json:"characteristic,omitempty"`
	// One or several physical instances or occurrences of the nutrition product
	Instance []NutritionProductInstance `json:"instance,omitempty"`
	// Comments made about the product
	Note []Annotation `json:"note,omitempty"`
}

type NutritionProductNutrient struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The (relevant) nutrients in the product
	Item *custom.CodeableReference[Substance] `json:"item,omitempty"`
	// The amount of nutrient expressed in one or more units: X per pack / per serving / per dose
	Amount []Ratio `json:"amount,omitempty"`
}

type NutritionProductIngredient struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The ingredient contained in the product
	Item custom.CodeableReference[NutritionProduct] `json:"item"`
	// The amount of ingredient that is in the product
	Amount []Ratio `json:"amount,omitempty"`
}

type NutritionProductCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code specifying the type of characteristic
	Type CodeableConcept `json:"type"`
	// The value of the characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// The value of the characteristic
	ValueString *string `json:"valueString,omitempty"`
	// The value of the characteristic
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The value of the characteristic
	ValueBase64binary *custom.Base64binary `json:"valueBase64binary,omitempty"`
	// The value of the characteristic
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// The value of the characteristic
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
}

type NutritionProductInstance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The amount of items or instances
	Quantity *Quantity `json:"quantity,omitempty"`
	// The identifier for the physical instance, typically a serial number or manufacturer number
	Identifier []Identifier `json:"identifier,omitempty"`
	// The name for the specific product
	Name *string `json:"name,omitempty"`
	// The identification of the batch or lot of the product
	LotNumber *string `json:"lotNumber,omitempty"`
	// The expiry date or date and time for the product
	Expiry *custom.DateTime `json:"expiry,omitempty"`
	// The date until which the product is expected to be good for consumption
	UseBy *custom.DateTime `json:"useBy,omitempty"`
	// An identifier that supports traceability to the event during which material in this product from one or more biological entities was obtained or pooled
	BiologicalSourceEvent *Identifier `json:"biologicalSourceEvent,omitempty"`
}

type OtherNutritionProduct NutritionProduct

func (n NutritionProduct) ResourceType() string {
	return "NutritionProduct"
}

func (n NutritionProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionProduct
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherNutritionProduct: OtherNutritionProduct(n), ResourceType: n.ResourceType()})
}

func UnmarshalNutritionProduct(b []byte) (res NutritionProduct, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
