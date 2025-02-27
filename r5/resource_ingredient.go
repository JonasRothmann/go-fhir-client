// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Ingredient
type IngredientSubstance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A code or full resource that represents the ingredient substance
	Code custom.CodeableReference[SubstanceDefinition] `json:"code"`
	// The quantity of substance, per presentation, or per volume or mass, and type of quantity
	Strength []IngredientSubstanceStrength `json:"strength,omitempty"`
}

type IngredientSubstanceStrength struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The quantity of substance in the unit of presentation
	PresentationRatio *Ratio `json:"presentationRatio,omitempty"`
	// The quantity of substance in the unit of presentation
	PresentationRatioRange *RatioRange `json:"presentationRatioRange,omitempty"`
	// The quantity of substance in the unit of presentation
	PresentationCodeableConcept *CodeableConcept `json:"presentationCodeableConcept,omitempty"`
	// The quantity of substance in the unit of presentation
	PresentationQuantity *Quantity `json:"presentationQuantity,omitempty"`
	// Text of either the whole presentation strength or a part of it (rest being in Strength.presentation as a ratio)
	TextPresentation *string `json:"textPresentation,omitempty"`
	// The strength per unitary volume (or mass)
	ConcentrationRatio *Ratio `json:"concentrationRatio,omitempty"`
	// The strength per unitary volume (or mass)
	ConcentrationRatioRange *RatioRange `json:"concentrationRatioRange,omitempty"`
	// The strength per unitary volume (or mass)
	ConcentrationCodeableConcept *CodeableConcept `json:"concentrationCodeableConcept,omitempty"`
	// The strength per unitary volume (or mass)
	ConcentrationQuantity *Quantity `json:"concentrationQuantity,omitempty"`
	// Text of either the whole concentration strength or a part of it (rest being in Strength.concentration as a ratio)
	TextConcentration *string `json:"textConcentration,omitempty"`
	// A code that indicates if the strength is, for example, based on the ingredient substance as stated or on the substance base (when the ingredient is a salt)
	Basis *CodeableConcept `json:"basis,omitempty"`
	// When strength is measured at a particular point or distance
	MeasurementPoint *string `json:"measurementPoint,omitempty"`
	// Where the strength range applies
	Country []CodeableConcept `json:"country,omitempty"`
	// Strength expressed in terms of a reference substance
	ReferenceStrength []IngredientSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`
}

type IngredientSubstanceStrengthReferenceStrength struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Relevant reference substance
	Substance custom.CodeableReference[SubstanceDefinition] `json:"substance"`
	// Strength expressed in terms of a reference substance
	StrengthRatio *Ratio `json:"strengthRatio,omitempty"`
	// Strength expressed in terms of a reference substance
	StrengthRatioRange *RatioRange `json:"strengthRatioRange,omitempty"`
	// Strength expressed in terms of a reference substance
	StrengthQuantity *Quantity `json:"strengthQuantity,omitempty"`
	// When strength is measured at a particular point or distance
	MeasurementPoint *string `json:"measurementPoint,omitempty"`
	// Where the strength range applies
	Country []CodeableConcept `json:"country,omitempty"`
}

type Ingredient struct {
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
	// An identifier or code by which the ingredient can be referenced
	Identifier *Identifier `json:"identifier,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// The product which this ingredient is a constituent part of
	For []custom.Reference[IngredientFor] `json:"for,omitempty"`
	// Purpose of the ingredient within the product, e.g. active, inactive
	Role CodeableConcept `json:"role"`
	// Precise action within the drug product, e.g. antioxidant, alkalizing agent
	Function []CodeableConcept `json:"function,omitempty"`
	// A classification of the ingredient according to where in the physical item it tends to be used, such the outer shell of a tablet, inner body or ink
	Group *CodeableConcept `json:"group,omitempty"`
	// If the ingredient is a known or suspected allergen
	AllergenicIndicator *bool `json:"allergenicIndicator,omitempty"`
	// A place for providing any notes that are relevant to the component, e.g. removed during process, adjusted for loss on drying
	Comment *custom.Markdown `json:"comment,omitempty"`
	// An organization that manufactures this ingredient
	Manufacturer []IngredientManufacturer `json:"manufacturer,omitempty"`
	// The substance that comprises this ingredient
	Substance IngredientSubstance `json:"substance"`
}

type IngredientManufacturer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// allowed | possible | actual
	Role *custom.Code `json:"role,omitempty"`
	// An organization that manufactures this ingredient
	Manufacturer custom.Reference[Organization] `json:"manufacturer"`
}

type OtherIngredient Ingredient

type IngredientFor interface {
	gofhirclient.Resource

	Is_IngredientFor()
}

func (m MedicinalProductDefinition) Is_IngredientFor()     {}
func (a AdministrableProductDefinition) Is_IngredientFor() {}
func (m ManufacturedItemDefinition) Is_IngredientFor()     {}

func (i Ingredient) ResourceType() string {
	return "Ingredient"
}

func (i Ingredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherIngredient
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherIngredient: OtherIngredient(i), ResourceType: i.ResourceType()})
}

func UnmarshalIngredient(b []byte) (res Ingredient, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
