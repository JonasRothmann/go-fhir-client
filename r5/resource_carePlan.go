// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CarePlan
type CarePlan struct {
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
	// External Ids for this plan
	Identifier []Identifier `json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[CarePlanInstantiatesCanonical] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `json:"instantiatesUri,omitempty"`
	// Fulfills plan, proposal or order
	BasedOn []custom.Reference[CarePlanBasedOn] `json:"basedOn,omitempty"`
	// CarePlan replaced by this CarePlan
	Replaces []custom.Reference[CarePlan] `json:"replaces,omitempty"`
	// Part of referenced CarePlan
	PartOf []custom.Reference[CarePlan] `json:"partOf,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// proposal | plan | order | option | directive
	Intent custom.Code `json:"intent"`
	// Type of plan
	Category []CodeableConcept `json:"category,omitempty"`
	// Human-friendly name for the care plan
	Title *string `json:"title,omitempty"`
	// Summary of nature of plan
	Description *string `json:"description,omitempty"`
	// Who the care plan is for
	Subject custom.Reference[CarePlanSubject] `json:"subject"`
	// The Encounter during which this CarePlan was created
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Time period plan covers
	Period *Period `json:"period,omitempty"`
	// Date record was first recorded
	Created *custom.DateTime `json:"created,omitempty"`
	// Who is the designated responsible party
	Custodian *custom.Reference[CarePlanCustodian] `json:"custodian,omitempty"`
	// Who provided the content of the care plan
	Contributor []custom.Reference[CarePlanContributor] `json:"contributor,omitempty"`
	// Who's involved in plan?
	CareTeam []custom.Reference[CareTeam] `json:"careTeam,omitempty"`
	// Health issues this plan addresses
	Addresses []custom.CodeableReference[Condition] `json:"addresses,omitempty"`
	// Information considered as part of plan
	SupportingInfo []custom.Reference[Resource] `json:"supportingInfo,omitempty"`
	// Desired outcome of plan
	Goal []custom.Reference[Goal] `json:"goal,omitempty"`
	// Action to occur or has occurred as part of plan
	Activity []CarePlanActivity `json:"activity,omitempty"`
	// Comments about the plan
	Note []Annotation `json:"note,omitempty"`
}

type CarePlanActivity struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Results of the activity (concept, or Appointment, Encounter, Procedure, etc.)
	PerformedActivity []custom.CodeableReference[Resource] `json:"performedActivity,omitempty"`
	// Comments about the activity status/progress
	Progress []Annotation `json:"progress,omitempty"`
	// Activity that is intended to be part of the care plan
	PlannedActivityReference *custom.Reference[CarePlanActivityPlannedActivityReference] `json:"plannedActivityReference,omitempty"`
}

type OtherCarePlan CarePlan

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
	return "CarePlan"
}

func (c CarePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCarePlan
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCarePlan: OtherCarePlan(c), ResourceType: c.ResourceType()})
}

func UnmarshalCarePlan(b []byte) (res CarePlan, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
