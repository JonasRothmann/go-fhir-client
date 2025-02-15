// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
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
	// The identifier of the request associated with this response, if any
	RequestIdentifier *Identifier `bson:"requestIdentifier,omitempty" json:"requestIdentifier,omitempty"`
	// Business identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// What guidance was requested
	Module custom.URI `bson:"module" json:"module"`
	// success | data-requested | data-required | in-progress | failure | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Patient the request was performed for
	Subject *custom.Reference[GuidanceResponseSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Encounter during which the response was returned
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the guidance response was processed
	OccurrenceDateTime *custom.DateTime `bson:"occurrenceDateTime,omitempty" json:"occurrenceDateTime,omitempty"`
	// Device returning the guidance
	Performer *custom.Reference[Device] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Why guidance is needed
	Reason []custom.CodeableReference `bson:"reason,omitempty" json:"reason,omitempty"`
	// Additional notes about the response
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Messages resulting from the evaluation of the artifact or artifacts
	EvaluationMessage *custom.Reference[OperationOutcome] `bson:"evaluationMessage,omitempty" json:"evaluationMessage,omitempty"`
	// The output parameters of the evaluation, if any
	OutputParameters *custom.Reference[Parameters] `bson:"outputParameters,omitempty" json:"outputParameters,omitempty"`
	// Proposed actions, if any
	Result []custom.Reference[GuidanceResponseResult] `bson:"result,omitempty" json:"result,omitempty"`
	// Additional required data
	DataRequirement []DataRequirement `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty"`
}

type GuidanceResponseSubject interface {
	gofhirclient.Resource

	Is_GuidanceResponseSubject()
}

func (p Patient) Is_GuidanceResponseSubject() {}
func (g Group) Is_GuidanceResponseSubject()   {}

type GuidanceResponseResult interface {
	gofhirclient.Resource

	Is_GuidanceResponseResult()
}

func (a Appointment) Is_GuidanceResponseResult()                {}
func (a AppointmentResponse) Is_GuidanceResponseResult()        {}
func (c CarePlan) Is_GuidanceResponseResult()                   {}
func (c Claim) Is_GuidanceResponseResult()                      {}
func (c CommunicationRequest) Is_GuidanceResponseResult()       {}
func (c Contract) Is_GuidanceResponseResult()                   {}
func (c CoverageEligibilityRequest) Is_GuidanceResponseResult() {}
func (d DeviceRequest) Is_GuidanceResponseResult()              {}
func (e EnrollmentRequest) Is_GuidanceResponseResult()          {}
func (i ImmunizationRecommendation) Is_GuidanceResponseResult() {}
func (m MedicationRequest) Is_GuidanceResponseResult()          {}
func (n NutritionOrder) Is_GuidanceResponseResult()             {}
func (r RequestOrchestration) Is_GuidanceResponseResult()       {}
func (s ServiceRequest) Is_GuidanceResponseResult()             {}
func (s SupplyRequest) Is_GuidanceResponseResult()              {}
func (t Task) Is_GuidanceResponseResult()                       {}
func (v VisionPrescription) Is_GuidanceResponseResult()         {}

func (g GuidanceResponse) ResourceType() string {
	return "StructureDefinition"
}
