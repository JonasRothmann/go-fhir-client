// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Medication
type Medication struct {
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
	// Business identifier for this medication
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Codes that identify this medication
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// active | inactive | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Organization that has authorization to market medication
	MarketingAuthorizationHolder *custom.Reference[Organization] `bson:"marketingAuthorizationHolder,omitempty" json:"marketingAuthorizationHolder,omitempty"`
	// powder | tablets | capsule +
	DoseForm *CodeableConcept `bson:"doseForm,omitempty" json:"doseForm,omitempty"`
	// When the specified product code does not infer a package size, this is the specific amount of drug in the product
	TotalVolume *Quantity `bson:"totalVolume,omitempty" json:"totalVolume,omitempty"`
	// Active or inactive ingredient
	Ingredient []MedicationIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	// Details about packaged medications
	Batch *MedicationBatch `bson:"batch,omitempty" json:"batch,omitempty"`
	// Knowledge about this medication
	Definition *custom.Reference[MedicationKnowledge] `bson:"definition,omitempty" json:"definition,omitempty"`
}

type MedicationIngredient struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The ingredient (substance or medication) that the ingredient.strength relates to
	Item custom.CodeableReference[MedicationIngredientItem] `bson:"item" json:"item"`
	// Active ingredient indicator
	IsActive *bool `bson:"isActive,omitempty" json:"isActive,omitempty"`
	// Quantity of ingredient present
	Strength *Ratio `bson:"strength,omitempty" json:"strength,omitempty"`
}

type MedicationBatch struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifier assigned to batch
	LotNumber *string `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	// When batch will expire
	ExpirationDate *custom.DateTime `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
}

type MedicationIngredientItem interface {
	gofhirclient.Resource

	Is_MedicationIngredientItem()
}

func (s Substance) Is_MedicationIngredientItem()  {}
func (m Medication) Is_MedicationIngredientItem() {}

func (m Medication) ResourceType() string {
	return "StructureDefinition"
}
