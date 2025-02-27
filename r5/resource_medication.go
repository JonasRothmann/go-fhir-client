// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Medication
type Medication struct {
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
	// Business identifier for this medication
	Identifier []Identifier `json:"identifier,omitempty"`
	// Codes that identify this medication
	Code *CodeableConcept `json:"code,omitempty"`
	// active | inactive | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// Organization that has authorization to market medication
	MarketingAuthorizationHolder *custom.Reference[Organization] `json:"marketingAuthorizationHolder,omitempty"`
	// powder | tablets | capsule +
	DoseForm *CodeableConcept `json:"doseForm,omitempty"`
	// When the specified product code does not infer a package size, this is the specific amount of drug in the product
	TotalVolume *Quantity `json:"totalVolume,omitempty"`
	// Active or inactive ingredient
	Ingredient []MedicationIngredient `json:"ingredient,omitempty"`
	// Details about packaged medications
	Batch *MedicationBatch `json:"batch,omitempty"`
	// Knowledge about this medication
	Definition *custom.Reference[MedicationKnowledge] `json:"definition,omitempty"`
}

type MedicationIngredient struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The ingredient (substance or medication) that the ingredient.strength relates to
	Item custom.CodeableReference[MedicationIngredientItem] `json:"item"`
	// Active ingredient indicator
	IsActive *bool `json:"isActive,omitempty"`
	// Quantity of ingredient present
	StrengthRatio *Ratio `json:"strengthRatio,omitempty"`
	// Quantity of ingredient present
	StrengthCodeableConcept *CodeableConcept `json:"strengthCodeableConcept,omitempty"`
	// Quantity of ingredient present
	StrengthQuantity *Quantity `json:"strengthQuantity,omitempty"`
}

type MedicationBatch struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifier assigned to batch
	LotNumber *string `json:"lotNumber,omitempty"`
	// When batch will expire
	ExpirationDate *custom.DateTime `json:"expirationDate,omitempty"`
}

type OtherMedication Medication

type MedicationIngredientItem interface {
	gofhirclient.Resource

	Is_MedicationIngredientItem()
}

func (s Substance) Is_MedicationIngredientItem()  {}
func (m Medication) Is_MedicationIngredientItem() {}

func (m Medication) ResourceType() string {
	return "Medication"
}

func (m Medication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedication
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMedication: OtherMedication(m), ResourceType: m.ResourceType()})
}

func UnmarshalMedication(b []byte) (res Medication, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
