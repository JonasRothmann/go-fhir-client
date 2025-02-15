// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/NutritionIntake
type NutritionIntakeIngredientLabel struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Total nutrient consumed
	Nutrient custom.CodeableReference[Substance] `bson:"nutrient" json:"nutrient"`
	// Total amount of nutrient consumed
	Amount Quantity `bson:"amount" json:"amount"`
}

type NutritionIntakePerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of performer
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who performed the intake
	Actor custom.Reference[NutritionIntakePerformerActor] `bson:"actor" json:"actor"`
}

type NutritionIntake struct {
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
	// External identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[NutritionIntakeInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Fulfils plan, proposal or order
	BasedOn []custom.Reference[NutritionIntakeBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[NutritionIntakePartOf] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Reason for current status
	StatusReason []CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// Code representing an overall type of nutrition intake
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Who is/was consuming the food or fluid
	Subject custom.Reference[NutritionIntakeSubject] `bson:"subject" json:"subject"`
	// Encounter associated with NutritionIntake
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// The date/time or interval when the food or fluid is/was consumed
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// When the intake was recorded
	Recorded *custom.DateTime `bson:"recorded,omitempty" json:"recorded,omitempty"`
	// Person or organization that provided the information about the consumption of this food or fluid
	Reported *bool `bson:"reported,omitempty" json:"reported,omitempty"`
	// What food or fluid product or item was consumed
	ConsumedItem []NutritionIntakeConsumedItem `bson:"consumedItem" json:"consumedItem"`
	// Total nutrient for the whole meal, product, serving
	IngredientLabel []NutritionIntakeIngredientLabel `bson:"ingredientLabel,omitempty" json:"ingredientLabel,omitempty"`
	// Who was performed in the intake
	Performer []NutritionIntakePerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Where the intake occurred
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Additional supporting information
	DerivedFrom []custom.Reference[Resource] `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	// Reason for why the food or fluid is /was consumed
	Reason []custom.CodeableReference[NutritionIntakeReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Further information about the consumption
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type NutritionIntakeConsumedItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of food or fluid product
	Type CodeableConcept `bson:"type" json:"type"`
	// Code that identifies the food or fluid product that was consumed
	NutritionProduct custom.CodeableReference[NutritionProduct] `bson:"nutritionProduct" json:"nutritionProduct"`
	// Scheduled frequency of consumption
	Schedule *Timing `bson:"schedule,omitempty" json:"schedule,omitempty"`
	// Quantity of the specified food
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
	// Rate at which enteral feeding was administered
	Rate *Quantity `bson:"rate,omitempty" json:"rate,omitempty"`
	// Flag to indicate if the food or fluid item was refused or otherwise not consumed
	NotConsumed *bool `bson:"notConsumed,omitempty" json:"notConsumed,omitempty"`
	// Reason food or fluid was not consumed
	NotConsumedReason *CodeableConcept `bson:"notConsumedReason,omitempty" json:"notConsumedReason,omitempty"`
}

type NutritionIntakeInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_NutritionIntakeInstantiatesCanonical()
}

func (a ActivityDefinition) Is_NutritionIntakeInstantiatesCanonical()    {}
func (c ChargeItemDefinition) Is_NutritionIntakeInstantiatesCanonical()  {}
func (c ClinicalUseDefinition) Is_NutritionIntakeInstantiatesCanonical() {}
func (e EventDefinition) Is_NutritionIntakeInstantiatesCanonical()       {}
func (m Measure) Is_NutritionIntakeInstantiatesCanonical()               {}
func (m MessageDefinition) Is_NutritionIntakeInstantiatesCanonical()     {}
func (o ObservationDefinition) Is_NutritionIntakeInstantiatesCanonical() {}
func (o OperationDefinition) Is_NutritionIntakeInstantiatesCanonical()   {}
func (p PlanDefinition) Is_NutritionIntakeInstantiatesCanonical()        {}
func (q Questionnaire) Is_NutritionIntakeInstantiatesCanonical()         {}
func (r Requirements) Is_NutritionIntakeInstantiatesCanonical()          {}
func (s SubscriptionTopic) Is_NutritionIntakeInstantiatesCanonical()     {}
func (t TestPlan) Is_NutritionIntakeInstantiatesCanonical()              {}
func (t TestScript) Is_NutritionIntakeInstantiatesCanonical()            {}

type NutritionIntakeBasedOn interface {
	gofhirclient.Resource

	Is_NutritionIntakeBasedOn()
}

func (n NutritionOrder) Is_NutritionIntakeBasedOn() {}
func (c CarePlan) Is_NutritionIntakeBasedOn()       {}
func (s ServiceRequest) Is_NutritionIntakeBasedOn() {}

type NutritionIntakePartOf interface {
	gofhirclient.Resource

	Is_NutritionIntakePartOf()
}

func (n NutritionIntake) Is_NutritionIntakePartOf() {}
func (p Procedure) Is_NutritionIntakePartOf()       {}
func (o Observation) Is_NutritionIntakePartOf()     {}

type NutritionIntakeSubject interface {
	gofhirclient.Resource

	Is_NutritionIntakeSubject()
}

func (p Patient) Is_NutritionIntakeSubject() {}
func (g Group) Is_NutritionIntakeSubject()   {}

type NutritionIntakePerformerActor interface {
	gofhirclient.Resource

	Is_NutritionIntakePerformerActor()
}

func (p Practitioner) Is_NutritionIntakePerformerActor()     {}
func (p PractitionerRole) Is_NutritionIntakePerformerActor() {}
func (o Organization) Is_NutritionIntakePerformerActor()     {}
func (c CareTeam) Is_NutritionIntakePerformerActor()         {}
func (p Patient) Is_NutritionIntakePerformerActor()          {}
func (d Device) Is_NutritionIntakePerformerActor()           {}
func (r RelatedPerson) Is_NutritionIntakePerformerActor()    {}

type NutritionIntakeReason interface {
	gofhirclient.Resource

	Is_NutritionIntakeReason()
}

func (c Condition) Is_NutritionIntakeReason()         {}
func (o Observation) Is_NutritionIntakeReason()       {}
func (d DiagnosticReport) Is_NutritionIntakeReason()  {}
func (d DocumentReference) Is_NutritionIntakeReason() {}

func (n NutritionIntake) ResourceType() string {
	return "StructureDefinition"
}
