// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/NutritionOrder
type NutritionOrderOralDietNutrient struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of nutrient that is being modified
	Modifier *CodeableConcept `json:"modifier,omitempty"`
	// Quantity of the specified nutrient
	Amount *Quantity `json:"amount,omitempty"`
}

type NutritionOrderOralDietTexture struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code to indicate how to alter the texture of the foods, e.g. pureed
	Modifier *CodeableConcept `json:"modifier,omitempty"`
	// Concepts that are used to identify an entity that is ingested for nutritional purposes
	FoodType *CodeableConcept `json:"foodType,omitempty"`
}

type NutritionOrderSupplement struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of supplement product requested
	Type *custom.CodeableReference[NutritionProduct] `json:"type,omitempty"`
	// Product or brand name of the nutritional supplement
	ProductName *string `json:"productName,omitempty"`
	// Scheduling information for supplements
	Schedule *NutritionOrderSupplementSchedule `json:"schedule,omitempty"`
	// Amount of the nutritional supplement
	Quantity *Quantity `json:"quantity,omitempty"`
	// Instructions or additional information about the oral supplement
	Instruction *string `json:"instruction,omitempty"`
}

type NutritionOrderEnteralFormulaAdministration struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Scheduling information for enteral formula products
	Schedule *NutritionOrderEnteralFormulaAdministrationSchedule `json:"schedule,omitempty"`
	// The volume of formula to provide
	Quantity *Quantity `json:"quantity,omitempty"`
	// Speed with which the formula is provided per period of time
	RateQuantity *Quantity `json:"rateQuantity,omitempty"`
	// Speed with which the formula is provided per period of time
	RateRatio *Ratio `json:"rateRatio,omitempty"`
}

type NutritionOrder struct {
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
	// Identifiers assigned to this order
	Identifier []Identifier `json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[NutritionOrderInstantiatesCanonical] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `json:"instantiatesUri,omitempty"`
	// Instantiates protocol or definition
	Instantiates []custom.URI `json:"instantiates,omitempty"`
	// What this order fulfills
	BasedOn []custom.Reference[NutritionOrderBasedOn] `json:"basedOn,omitempty"`
	// Composite Request ID
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// Who requires the diet, formula or nutritional supplement
	Subject custom.Reference[NutritionOrderSubject] `json:"subject"`
	// The encounter associated with this nutrition order
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Information to support fulfilling of the nutrition order
	SupportingInformation []custom.Reference[Resource] `json:"supportingInformation,omitempty"`
	// Date and time the nutrition order was requested
	DateTime custom.DateTime `json:"dateTime"`
	// Who ordered the diet, formula or nutritional supplement
	Orderer *custom.Reference[NutritionOrderOrderer] `json:"orderer,omitempty"`
	// Who is desired to perform the administration of what is being ordered
	Performer []custom.CodeableReference[NutritionOrderPerformer] `json:"performer,omitempty"`
	// List of the patient's food and nutrition-related allergies and intolerances
	AllergyIntolerance []custom.Reference[AllergyIntolerance] `json:"allergyIntolerance,omitempty"`
	// Order-specific modifier about the type of food that should be given
	FoodPreferenceModifier []CodeableConcept `json:"foodPreferenceModifier,omitempty"`
	// Order-specific modifier about the type of food that should not be given
	ExcludeFoodModifier []CodeableConcept `json:"excludeFoodModifier,omitempty"`
	// Capture when a food item is brought in by the patient and/or family
	OutsideFoodAllowed *bool `json:"outsideFoodAllowed,omitempty"`
	// Oral diet components
	OralDiet *NutritionOrderOralDiet `json:"oralDiet,omitempty"`
	// Supplement components
	Supplement []NutritionOrderSupplement `json:"supplement,omitempty"`
	// Enteral formula components
	EnteralFormula *NutritionOrderEnteralFormula `json:"enteralFormula,omitempty"`
	// Comments
	Note []Annotation `json:"note,omitempty"`
}

type NutritionOrderOralDietSchedule struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Scheduled frequency of diet
	Timing []Timing `json:"timing,omitempty"`
	// Take 'as needed'
	AsNeeded *bool `json:"asNeeded,omitempty"`
	// Take 'as needed' for x
	AsNeededFor *CodeableConcept `json:"asNeededFor,omitempty"`
}

type NutritionOrderEnteralFormula struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of enteral or infant formula
	BaseFormulaType *custom.CodeableReference[NutritionProduct] `json:"baseFormulaType,omitempty"`
	// Product or brand name of the enteral or infant formula
	BaseFormulaProductName *string `json:"baseFormulaProductName,omitempty"`
	// Intended type of device for the administration
	DeliveryDevice []custom.CodeableReference[DeviceDefinition] `json:"deliveryDevice,omitempty"`
	// Components to add to the feeding
	Additive []NutritionOrderEnteralFormulaAdditive `json:"additive,omitempty"`
	// Amount of energy per specified volume that is required
	CaloricDensity *Quantity `json:"caloricDensity,omitempty"`
	// How the formula should enter the patient's gastrointestinal tract
	RouteOfAdministration *CodeableConcept `json:"routeOfAdministration,omitempty"`
	// Formula feeding instruction as structured data
	Administration []NutritionOrderEnteralFormulaAdministration `json:"administration,omitempty"`
	// Upper limit on formula volume per unit of time
	MaxVolumeToDeliver *Quantity `json:"maxVolumeToDeliver,omitempty"`
	// Formula feeding instructions expressed as text
	AdministrationInstruction *custom.Markdown `json:"administrationInstruction,omitempty"`
}

type NutritionOrderEnteralFormulaAdditive struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of modular component to add to the feeding
	Type *custom.CodeableReference[NutritionProduct] `json:"type,omitempty"`
	// Product or brand name of the modular additive
	ProductName *string `json:"productName,omitempty"`
	// Amount of additive to be given or mixed in
	Quantity *Quantity `json:"quantity,omitempty"`
}

type NutritionOrderEnteralFormulaAdministrationSchedule struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Scheduled frequency of enteral formula
	Timing []Timing `json:"timing,omitempty"`
	// Take 'as needed'
	AsNeeded *bool `json:"asNeeded,omitempty"`
	// Take 'as needed' for x
	AsNeededFor *CodeableConcept `json:"asNeededFor,omitempty"`
}

type NutritionOrderOralDiet struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of oral diet or diet restrictions that describe what can be consumed orally
	Type []CodeableConcept `json:"type,omitempty"`
	// Scheduling information for oral diets
	Schedule *NutritionOrderOralDietSchedule `json:"schedule,omitempty"`
	// Required  nutrient modifications
	Nutrient []NutritionOrderOralDietNutrient `json:"nutrient,omitempty"`
	// Required  texture modifications
	Texture []NutritionOrderOralDietTexture `json:"texture,omitempty"`
	// The required consistency of fluids and liquids provided to the patient
	FluidConsistencyType []CodeableConcept `json:"fluidConsistencyType,omitempty"`
	// Instructions or additional information about the oral diet
	Instruction *string `json:"instruction,omitempty"`
}

type NutritionOrderSupplementSchedule struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Scheduled frequency of diet
	Timing []Timing `json:"timing,omitempty"`
	// Take 'as needed'
	AsNeeded *bool `json:"asNeeded,omitempty"`
	// Take 'as needed' for x
	AsNeededFor *CodeableConcept `json:"asNeededFor,omitempty"`
}

type OtherNutritionOrder NutritionOrder

type NutritionOrderInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_NutritionOrderInstantiatesCanonical()
}

func (a ActivityDefinition) Is_NutritionOrderInstantiatesCanonical() {}
func (p PlanDefinition) Is_NutritionOrderInstantiatesCanonical()     {}

type NutritionOrderBasedOn interface {
	gofhirclient.Resource

	Is_NutritionOrderBasedOn()
}

func (c CarePlan) Is_NutritionOrderBasedOn()       {}
func (n NutritionOrder) Is_NutritionOrderBasedOn() {}
func (s ServiceRequest) Is_NutritionOrderBasedOn() {}

type NutritionOrderSubject interface {
	gofhirclient.Resource

	Is_NutritionOrderSubject()
}

func (p Patient) Is_NutritionOrderSubject() {}
func (g Group) Is_NutritionOrderSubject()   {}

type NutritionOrderOrderer interface {
	gofhirclient.Resource

	Is_NutritionOrderOrderer()
}

func (p Practitioner) Is_NutritionOrderOrderer()     {}
func (p PractitionerRole) Is_NutritionOrderOrderer() {}

type NutritionOrderPerformer interface {
	gofhirclient.Resource

	Is_NutritionOrderPerformer()
}

func (c CareTeam) Is_NutritionOrderPerformer()         {}
func (p Practitioner) Is_NutritionOrderPerformer()     {}
func (p PractitionerRole) Is_NutritionOrderPerformer() {}
func (r RelatedPerson) Is_NutritionOrderPerformer()    {}
func (p Patient) Is_NutritionOrderPerformer()          {}
func (o Organization) Is_NutritionOrderPerformer()     {}

func (n NutritionOrder) ResourceType() string {
	return "NutritionOrder"
}

func (n NutritionOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionOrder
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherNutritionOrder: OtherNutritionOrder(n), ResourceType: n.ResourceType()})
}

func UnmarshalNutritionOrder(b []byte) (res NutritionOrder, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
