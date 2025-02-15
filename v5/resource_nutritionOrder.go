// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/NutritionOrder
type NutritionOrderOralDiet struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of oral diet or diet restrictions that describe what can be consumed orally
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Scheduling information for oral diets
	Schedule *NutritionOrderOralDietSchedule `bson:"schedule,omitempty" json:"schedule,omitempty"`
	// Required  nutrient modifications
	Nutrient []NutritionOrderOralDietNutrient `bson:"nutrient,omitempty" json:"nutrient,omitempty"`
	// Required  texture modifications
	Texture []NutritionOrderOralDietTexture `bson:"texture,omitempty" json:"texture,omitempty"`
	// The required consistency of fluids and liquids provided to the patient
	FluidConsistencyType []CodeableConcept `bson:"fluidConsistencyType,omitempty" json:"fluidConsistencyType,omitempty"`
	// Instructions or additional information about the oral diet
	Instruction *string `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

type NutritionOrderOralDietNutrient struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of nutrient that is being modified
	Modifier *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Quantity of the specified nutrient
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
}

type NutritionOrderOralDietTexture struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code to indicate how to alter the texture of the foods, e.g. pureed
	Modifier *CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Concepts that are used to identify an entity that is ingested for nutritional purposes
	FoodType *CodeableConcept `bson:"foodType,omitempty" json:"foodType,omitempty"`
}

type NutritionOrderSupplement struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of supplement product requested
	Type *custom.CodeableReference[NutritionProduct] `bson:"type,omitempty" json:"type,omitempty"`
	// Product or brand name of the nutritional supplement
	ProductName *string `bson:"productName,omitempty" json:"productName,omitempty"`
	// Scheduling information for supplements
	Schedule *NutritionOrderSupplementSchedule `bson:"schedule,omitempty" json:"schedule,omitempty"`
	// Amount of the nutritional supplement
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Instructions or additional information about the oral supplement
	Instruction *string `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

type NutritionOrderEnteralFormulaAdministration struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Scheduling information for enteral formula products
	Schedule *NutritionOrderEnteralFormulaAdministrationSchedule `bson:"schedule,omitempty" json:"schedule,omitempty"`
	// The volume of formula to provide
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Speed with which the formula is provided per period of time
	Rate *Quantity `bson:"rate,omitempty" json:"rate,omitempty"`
}

type NutritionOrder struct {
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
	// Identifiers assigned to this order
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[NutritionOrderInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Instantiates protocol or definition
	Instantiates []custom.URI `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	// What this order fulfills
	BasedOn []custom.Reference[NutritionOrderBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Composite Request ID
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `bson:"intent" json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// Who requires the diet, formula or nutritional supplement
	Subject custom.Reference[NutritionOrderSubject] `bson:"subject" json:"subject"`
	// The encounter associated with this nutrition order
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Information to support fulfilling of the nutrition order
	SupportingInformation []custom.Reference[Resource] `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	// Date and time the nutrition order was requested
	DateTime custom.DateTime `bson:"dateTime" json:"dateTime"`
	// Who ordered the diet, formula or nutritional supplement
	Orderer *custom.Reference[NutritionOrderOrderer] `bson:"orderer,omitempty" json:"orderer,omitempty"`
	// Who is desired to perform the administration of what is being ordered
	Performer []custom.CodeableReference[NutritionOrderPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// List of the patient's food and nutrition-related allergies and intolerances
	AllergyIntolerance []custom.Reference[AllergyIntolerance] `bson:"allergyIntolerance,omitempty" json:"allergyIntolerance,omitempty"`
	// Order-specific modifier about the type of food that should be given
	FoodPreferenceModifier []CodeableConcept `bson:"foodPreferenceModifier,omitempty" json:"foodPreferenceModifier,omitempty"`
	// Order-specific modifier about the type of food that should not be given
	ExcludeFoodModifier []CodeableConcept `bson:"excludeFoodModifier,omitempty" json:"excludeFoodModifier,omitempty"`
	// Capture when a food item is brought in by the patient and/or family
	OutsideFoodAllowed *bool `bson:"outsideFoodAllowed,omitempty" json:"outsideFoodAllowed,omitempty"`
	// Oral diet components
	OralDiet *NutritionOrderOralDiet `bson:"oralDiet,omitempty" json:"oralDiet,omitempty"`
	// Supplement components
	Supplement []NutritionOrderSupplement `bson:"supplement,omitempty" json:"supplement,omitempty"`
	// Enteral formula components
	EnteralFormula *NutritionOrderEnteralFormula `bson:"enteralFormula,omitempty" json:"enteralFormula,omitempty"`
	// Comments
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type NutritionOrderOralDietSchedule struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Scheduled frequency of diet
	Timing []Timing `bson:"timing,omitempty" json:"timing,omitempty"`
	// Take 'as needed'
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
	// Take 'as needed' for x
	AsNeededFor *CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty"`
}

type NutritionOrderSupplementSchedule struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Scheduled frequency of diet
	Timing []Timing `bson:"timing,omitempty" json:"timing,omitempty"`
	// Take 'as needed'
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
	// Take 'as needed' for x
	AsNeededFor *CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty"`
}

type NutritionOrderEnteralFormula struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of enteral or infant formula
	BaseFormulaType *custom.CodeableReference[NutritionProduct] `bson:"baseFormulaType,omitempty" json:"baseFormulaType,omitempty"`
	// Product or brand name of the enteral or infant formula
	BaseFormulaProductName *string `bson:"baseFormulaProductName,omitempty" json:"baseFormulaProductName,omitempty"`
	// Intended type of device for the administration
	DeliveryDevice []custom.CodeableReference[DeviceDefinition] `bson:"deliveryDevice,omitempty" json:"deliveryDevice,omitempty"`
	// Components to add to the feeding
	Additive []NutritionOrderEnteralFormulaAdditive `bson:"additive,omitempty" json:"additive,omitempty"`
	// Amount of energy per specified volume that is required
	CaloricDensity *Quantity `bson:"caloricDensity,omitempty" json:"caloricDensity,omitempty"`
	// How the formula should enter the patient's gastrointestinal tract
	RouteOfAdministration *CodeableConcept `bson:"routeOfAdministration,omitempty" json:"routeOfAdministration,omitempty"`
	// Formula feeding instruction as structured data
	Administration []NutritionOrderEnteralFormulaAdministration `bson:"administration,omitempty" json:"administration,omitempty"`
	// Upper limit on formula volume per unit of time
	MaxVolumeToDeliver *Quantity `bson:"maxVolumeToDeliver,omitempty" json:"maxVolumeToDeliver,omitempty"`
	// Formula feeding instructions expressed as text
	AdministrationInstruction *custom.Markdown `bson:"administrationInstruction,omitempty" json:"administrationInstruction,omitempty"`
}

type NutritionOrderEnteralFormulaAdditive struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of modular component to add to the feeding
	Type *custom.CodeableReference[NutritionProduct] `bson:"type,omitempty" json:"type,omitempty"`
	// Product or brand name of the modular additive
	ProductName *string `bson:"productName,omitempty" json:"productName,omitempty"`
	// Amount of additive to be given or mixed in
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
}

type NutritionOrderEnteralFormulaAdministrationSchedule struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Scheduled frequency of enteral formula
	Timing []Timing `bson:"timing,omitempty" json:"timing,omitempty"`
	// Take 'as needed'
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
	// Take 'as needed' for x
	AsNeededFor *CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty"`
}

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
	return "StructureDefinition"
}
