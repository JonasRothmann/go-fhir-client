// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CarePlan
type CarePlan struct {
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
	// External Ids for this plan
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[CarePlanInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Fulfills plan, proposal or order
	BasedOn []custom.Reference[CarePlanBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// CarePlan replaced by this CarePlan
	Replaces []custom.Reference[CarePlan] `bson:"replaces,omitempty" json:"replaces,omitempty"`
	// Part of referenced CarePlan
	PartOf []custom.Reference[CarePlan] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// proposal | plan | order | option | directive
	Intent custom.Code `bson:"intent" json:"intent"`
	// Type of plan
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Human-friendly name for the care plan
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Summary of nature of plan
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Who the care plan is for
	Subject custom.Reference[CarePlanSubject] `bson:"subject" json:"subject"`
	// The Encounter during which this CarePlan was created
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Time period plan covers
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Date record was first recorded
	Created *custom.DateTime `bson:"created,omitempty" json:"created,omitempty"`
	// Who is the designated responsible party
	Custodian *custom.Reference[CarePlanCustodian] `bson:"custodian,omitempty" json:"custodian,omitempty"`
	// Who provided the content of the care plan
	Contributor []custom.Reference[CarePlanContributor] `bson:"contributor,omitempty" json:"contributor,omitempty"`
	// Who's involved in plan?
	CareTeam []custom.Reference[CareTeam] `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	// Health issues this plan addresses
	Addresses []custom.CodeableReference[Condition] `bson:"addresses,omitempty" json:"addresses,omitempty"`
	// Information considered as part of plan
	SupportingInfo []custom.Reference[Resource] `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Desired outcome of plan
	Goal []custom.Reference[Goal] `bson:"goal,omitempty" json:"goal,omitempty"`
	// Action to occur or has occurred as part of plan
	Activity []CarePlanActivity `bson:"activity,omitempty" json:"activity,omitempty"`
	// Comments about the plan
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type CarePlanActivity struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Results of the activity (concept, or Appointment, Encounter, Procedure, etc.)
	PerformedActivity []custom.CodeableReference[Resource] `bson:"performedActivity,omitempty" json:"performedActivity,omitempty"`
	// Comments about the activity status/progress
	Progress []Annotation `bson:"progress,omitempty" json:"progress,omitempty"`
	// Activity that is intended to be part of the care plan
	PlannedActivityReference *custom.Reference[CarePlanActivityPlannedActivityReference] `bson:"plannedActivityReference,omitempty" json:"plannedActivityReference,omitempty"`
}

type CarePlanInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_CarePlanInstantiatesCanonical()
}

func (p PlanDefinition) Is_CarePlanInstantiatesCanonical()      {}
func (q Questionnaire) Is_CarePlanInstantiatesCanonical()       {}
func (m Measure) Is_CarePlanInstantiatesCanonical()             {}
func (a ActivityDefinition) Is_CarePlanInstantiatesCanonical()  {}
func (o OperationDefinition) Is_CarePlanInstantiatesCanonical() {}

type CarePlanBasedOn interface {
	gofhirclient.Resource

	Is_CarePlanBasedOn()
}

func (c CarePlan) Is_CarePlanBasedOn()             {}
func (s ServiceRequest) Is_CarePlanBasedOn()       {}
func (r RequestOrchestration) Is_CarePlanBasedOn() {}
func (n NutritionOrder) Is_CarePlanBasedOn()       {}

type CarePlanSubject interface {
	gofhirclient.Resource

	Is_CarePlanSubject()
}

func (p Patient) Is_CarePlanSubject() {}
func (g Group) Is_CarePlanSubject()   {}

type CarePlanCustodian interface {
	gofhirclient.Resource

	Is_CarePlanCustodian()
}

func (p Patient) Is_CarePlanCustodian()          {}
func (p Practitioner) Is_CarePlanCustodian()     {}
func (p PractitionerRole) Is_CarePlanCustodian() {}
func (d Device) Is_CarePlanCustodian()           {}
func (r RelatedPerson) Is_CarePlanCustodian()    {}
func (o Organization) Is_CarePlanCustodian()     {}
func (c CareTeam) Is_CarePlanCustodian()         {}

type CarePlanContributor interface {
	gofhirclient.Resource

	Is_CarePlanContributor()
}

func (p Patient) Is_CarePlanContributor()          {}
func (p Practitioner) Is_CarePlanContributor()     {}
func (p PractitionerRole) Is_CarePlanContributor() {}
func (d Device) Is_CarePlanContributor()           {}
func (r RelatedPerson) Is_CarePlanContributor()    {}
func (o Organization) Is_CarePlanContributor()     {}
func (c CareTeam) Is_CarePlanContributor()         {}

type CarePlanActivityPlannedActivityReference interface {
	gofhirclient.Resource

	Is_CarePlanActivityPlannedActivityReference()
}

func (a Appointment) Is_CarePlanActivityPlannedActivityReference()                {}
func (c CommunicationRequest) Is_CarePlanActivityPlannedActivityReference()       {}
func (d DeviceRequest) Is_CarePlanActivityPlannedActivityReference()              {}
func (m MedicationRequest) Is_CarePlanActivityPlannedActivityReference()          {}
func (n NutritionOrder) Is_CarePlanActivityPlannedActivityReference()             {}
func (t Task) Is_CarePlanActivityPlannedActivityReference()                       {}
func (s ServiceRequest) Is_CarePlanActivityPlannedActivityReference()             {}
func (v VisionPrescription) Is_CarePlanActivityPlannedActivityReference()         {}
func (r RequestOrchestration) Is_CarePlanActivityPlannedActivityReference()       {}
func (i ImmunizationRecommendation) Is_CarePlanActivityPlannedActivityReference() {}
func (s SupplyRequest) Is_CarePlanActivityPlannedActivityReference()              {}

func (c CarePlan) ResourceType() string {
	return "StructureDefinition"
}
