// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/GuidanceResponse
type GuidanceResponse struct {
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
	// The identifier of the request associated with this response, if any
	RequestIdentifier *Identifier `json:"requestIdentifier,omitempty"`
	// Business identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// What guidance was requested
	ModuleUri *custom.URI `json:"moduleUri,omitempty"`
	// What guidance was requested
	ModuleCanonical *custom.Canonical[any] `json:"moduleCanonical,omitempty"`
	// What guidance was requested
	ModuleCodeableConcept *CodeableConcept `json:"moduleCodeableConcept,omitempty"`
	// success | data-requested | data-required | in-progress | failure | entered-in-error
	Status custom.Code `json:"status"`
	// Patient the request was performed for
	Subject *custom.Reference[GuidanceResponseSubject] `json:"subject,omitempty"`
	// Encounter during which the response was returned
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When the guidance response was processed
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// Device returning the guidance
	Performer *custom.Reference[Device] `json:"performer,omitempty"`
	// Why guidance is needed
	Reason []custom.CodeableReference[gofhirclient.Resource] `json:"reason,omitempty"`
	// Additional notes about the response
	Note []Annotation `json:"note,omitempty"`
	// Messages resulting from the evaluation of the artifact or artifacts
	EvaluationMessage *custom.Reference[OperationOutcome] `json:"evaluationMessage,omitempty"`
	// The output parameters of the evaluation, if any
	OutputParameters *custom.Reference[Parameters] `json:"outputParameters,omitempty"`
	// Proposed actions, if any
	Result []custom.Reference[GuidanceResponseResult] `json:"result,omitempty"`
	// Additional required data
	DataRequirement []DataRequirement `json:"dataRequirement,omitempty"`
}

type OtherGuidanceResponse GuidanceResponse

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
	return "GuidanceResponse"
}

func (g GuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGuidanceResponse
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherGuidanceResponse: OtherGuidanceResponse(g), ResourceType: g.ResourceType()})
}

func UnmarshalGuidanceResponse(b []byte) (res GuidanceResponse, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
