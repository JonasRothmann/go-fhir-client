// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/NutritionIntake
type NutritionIntake struct {
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
	// External identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[NutritionIntakeInstantiatesCanonical] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `json:"instantiatesUri,omitempty"`
	// Fulfils plan, proposal or order
	BasedOn []custom.Reference[NutritionIntakeBasedOn] `json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[NutritionIntakePartOf] `json:"partOf,omitempty"`
	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Reason for current status
	StatusReason []CodeableConcept `json:"statusReason,omitempty"`
	// Code representing an overall type of nutrition intake
	Code *CodeableConcept `json:"code,omitempty"`
	// Who is/was consuming the food or fluid
	Subject custom.Reference[NutritionIntakeSubject] `json:"subject"`
	// Encounter associated with NutritionIntake
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// The date/time or interval when the food or fluid is/was consumed
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// The date/time or interval when the food or fluid is/was consumed
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When the intake was recorded
	Recorded *custom.DateTime `json:"recorded,omitempty"`
	// Person or organization that provided the information about the consumption of this food or fluid
	ReportedBoolean *bool `json:"reportedBoolean,omitempty"`
	// Person or organization that provided the information about the consumption of this food or fluid
	ReportedReference *custom.Reference[NutritionIntakeReportedReference] `json:"reportedReference,omitempty"`
	// What food or fluid product or item was consumed
	ConsumedItem []NutritionIntakeConsumedItem `json:"consumedItem"`
	// Total nutrient for the whole meal, product, serving
	IngredientLabel []NutritionIntakeIngredientLabel `json:"ingredientLabel,omitempty"`
	// Who was performed in the intake
	Performer []NutritionIntakePerformer `json:"performer,omitempty"`
	// Where the intake occurred
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Additional supporting information
	DerivedFrom []custom.Reference[Resource] `json:"derivedFrom,omitempty"`
	// Reason for why the food or fluid is /was consumed
	Reason []custom.CodeableReference[NutritionIntakeReason] `json:"reason,omitempty"`
	// Further information about the consumption
	Note []Annotation `json:"note,omitempty"`
}

type NutritionIntakeConsumedItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of food or fluid product
	Type CodeableConcept `json:"type"`
	// Code that identifies the food or fluid product that was consumed
	NutritionProduct custom.CodeableReference[NutritionProduct] `json:"nutritionProduct"`
	// Scheduled frequency of consumption
	Schedule *Timing `json:"schedule,omitempty"`
	// Quantity of the specified food
	Amount *Quantity `json:"amount,omitempty"`
	// Rate at which enteral feeding was administered
	Rate *Quantity `json:"rate,omitempty"`
	// Flag to indicate if the food or fluid item was refused or otherwise not consumed
	NotConsumed *bool `json:"notConsumed,omitempty"`
	// Reason food or fluid was not consumed
	NotConsumedReason *CodeableConcept `json:"notConsumedReason,omitempty"`
}

type NutritionIntakeIngredientLabel struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Total nutrient consumed
	Nutrient custom.CodeableReference[Substance] `json:"nutrient"`
	// Total amount of nutrient consumed
	Amount Quantity `json:"amount"`
}

type NutritionIntakePerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of performer
	Function *CodeableConcept `json:"function,omitempty"`
	// Who performed the intake
	Actor custom.Reference[NutritionIntakePerformerActor] `json:"actor"`
}

type OtherNutritionIntake NutritionIntake

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

type NutritionIntakeReportedReference interface {
	gofhirclient.Resource

	Is_NutritionIntakeReportedReference()
}

func (p Patient) Is_NutritionIntakeReportedReference()          {}
func (r RelatedPerson) Is_NutritionIntakeReportedReference()    {}
func (p Practitioner) Is_NutritionIntakeReportedReference()     {}
func (p PractitionerRole) Is_NutritionIntakeReportedReference() {}
func (o Organization) Is_NutritionIntakeReportedReference()     {}

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
	return "NutritionIntake"
}

func (n NutritionIntake) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionIntake
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherNutritionIntake: OtherNutritionIntake(n), ResourceType: n.ResourceType()})
}

func UnmarshalNutritionIntake(b []byte) (res NutritionIntake, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
