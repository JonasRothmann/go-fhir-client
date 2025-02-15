// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Ingredient
type Ingredient struct {
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
	// An identifier or code by which the ingredient can be referenced
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// The product which this ingredient is a constituent part of
	For []custom.Reference[IngredientFor] `bson:"for,omitempty" json:"for,omitempty"`
	// Purpose of the ingredient within the product, e.g. active, inactive
	Role CodeableConcept `bson:"role" json:"role"`
	// Precise action within the drug product, e.g. antioxidant, alkalizing agent
	Function []CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// A classification of the ingredient according to where in the physical item it tends to be used, such the outer shell of a tablet, inner body or ink
	Group *CodeableConcept `bson:"group,omitempty" json:"group,omitempty"`
	// If the ingredient is a known or suspected allergen
	AllergenicIndicator *bool `bson:"allergenicIndicator,omitempty" json:"allergenicIndicator,omitempty"`
	// A place for providing any notes that are relevant to the component, e.g. removed during process, adjusted for loss on drying
	Comment *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
	// An organization that manufactures this ingredient
	Manufacturer []IngredientManufacturer `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// The substance that comprises this ingredient
	Substance IngredientSubstance `bson:"substance" json:"substance"`
}

type IngredientManufacturer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// allowed | possible | actual
	Role *custom.Code `bson:"role,omitempty" json:"role,omitempty"`
	// An organization that manufactures this ingredient
	Manufacturer custom.Reference[Organization] `bson:"manufacturer" json:"manufacturer"`
}

type IngredientSubstance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A code or full resource that represents the ingredient substance
	Code custom.CodeableReference[SubstanceDefinition] `bson:"code" json:"code"`
	// The quantity of substance, per presentation, or per volume or mass, and type of quantity
	Strength []IngredientSubstanceStrength `bson:"strength,omitempty" json:"strength,omitempty"`
}

type IngredientSubstanceStrength struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The quantity of substance in the unit of presentation
	Presentation *Ratio `bson:"presentation,omitempty" json:"presentation,omitempty"`
	// Text of either the whole presentation strength or a part of it (rest being in Strength.presentation as a ratio)
	TextPresentation *string `bson:"textPresentation,omitempty" json:"textPresentation,omitempty"`
	// The strength per unitary volume (or mass)
	Concentration *Ratio `bson:"concentration,omitempty" json:"concentration,omitempty"`
	// Text of either the whole concentration strength or a part of it (rest being in Strength.concentration as a ratio)
	TextConcentration *string `bson:"textConcentration,omitempty" json:"textConcentration,omitempty"`
	// A code that indicates if the strength is, for example, based on the ingredient substance as stated or on the substance base (when the ingredient is a salt)
	Basis *CodeableConcept `bson:"basis,omitempty" json:"basis,omitempty"`
	// When strength is measured at a particular point or distance
	MeasurementPoint *string `bson:"measurementPoint,omitempty" json:"measurementPoint,omitempty"`
	// Where the strength range applies
	Country []CodeableConcept `bson:"country,omitempty" json:"country,omitempty"`
	// Strength expressed in terms of a reference substance
	ReferenceStrength []IngredientSubstanceStrengthReferenceStrength `bson:"referenceStrength,omitempty" json:"referenceStrength,omitempty"`
}

type IngredientSubstanceStrengthReferenceStrength struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Relevant reference substance
	Substance custom.CodeableReference[SubstanceDefinition] `bson:"substance" json:"substance"`
	// Strength expressed in terms of a reference substance
	Strength Ratio `bson:"strength" json:"strength"`
	// When strength is measured at a particular point or distance
	MeasurementPoint *string `bson:"measurementPoint,omitempty" json:"measurementPoint,omitempty"`
	// Where the strength range applies
	Country []CodeableConcept `bson:"country,omitempty" json:"country,omitempty"`
}

type IngredientFor interface {
	gofhirclient.Resource

	Is_IngredientFor()
}

func (m MedicinalProductDefinition) Is_IngredientFor()     {}
func (a AdministrableProductDefinition) Is_IngredientFor() {}
func (m ManufacturedItemDefinition) Is_IngredientFor()     {}

func (i Ingredient) ResourceType() string {
	return "StructureDefinition"
}
