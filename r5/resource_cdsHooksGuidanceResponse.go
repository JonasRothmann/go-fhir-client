// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/cdshooksguidanceresponse
type CdsHooksGuidanceResponse struct {
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
	// Extension
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The identifier of the request associated with this response, if any
	RequestIdentifier Identifier `json:"requestIdentifier"`
	// Business identifier
	Identifier Identifier `json:"identifier"`
	// What guidance was requested
	ModuleUri *custom.URI `json:"moduleUri,omitempty"`
	// success | data-requested | data-required | in-progress | failure | entered-in-error
	Status custom.Code `json:"status"`
	// Patient the request was performed for
	Subject *custom.Reference[CdsHooksGuidanceResponseSubject] `json:"subject,omitempty"`
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
	Result []custom.Reference[CdsHooksGuidanceResponseResult] `json:"result,omitempty"`
	// Additional required data
	DataRequirement []DataRequirement `json:"dataRequirement,omitempty"`
}

type OtherCdsHooksGuidanceResponse CdsHooksGuidanceResponse

type CdsHooksGuidanceResponseSubject interface {
	gofhirclient.Resource

	Is_CdsHooksGuidanceResponseSubject()
}

func (p Patient) Is_CdsHooksGuidanceResponseSubject() {}
func (g Group) Is_CdsHooksGuidanceResponseSubject()   {}

type CdsHooksGuidanceResponseResult interface {
	gofhirclient.Resource

	Is_CdsHooksGuidanceResponseResult()
}

func (a Appointment) Is_CdsHooksGuidanceResponseResult()                {}
func (a AppointmentResponse) Is_CdsHooksGuidanceResponseResult()        {}
func (c CarePlan) Is_CdsHooksGuidanceResponseResult()                   {}
func (c Claim) Is_CdsHooksGuidanceResponseResult()                      {}
func (c CommunicationRequest) Is_CdsHooksGuidanceResponseResult()       {}
func (c Contract) Is_CdsHooksGuidanceResponseResult()                   {}
func (c CoverageEligibilityRequest) Is_CdsHooksGuidanceResponseResult() {}
func (d DeviceRequest) Is_CdsHooksGuidanceResponseResult()              {}
func (e EnrollmentRequest) Is_CdsHooksGuidanceResponseResult()          {}
func (i ImmunizationRecommendation) Is_CdsHooksGuidanceResponseResult() {}
func (m MedicationRequest) Is_CdsHooksGuidanceResponseResult()          {}
func (n NutritionOrder) Is_CdsHooksGuidanceResponseResult()             {}
func (r RequestOrchestration) Is_CdsHooksGuidanceResponseResult()       {}
func (s ServiceRequest) Is_CdsHooksGuidanceResponseResult()             {}
func (s SupplyRequest) Is_CdsHooksGuidanceResponseResult()              {}
func (t Task) Is_CdsHooksGuidanceResponseResult()                       {}
func (v VisionPrescription) Is_CdsHooksGuidanceResponseResult()         {}

func (c CdsHooksGuidanceResponse) ResourceType() string {
	return "CDSHooksGuidanceResponse"
}

func (c CdsHooksGuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCdsHooksGuidanceResponse
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCdsHooksGuidanceResponse: OtherCdsHooksGuidanceResponse(c), ResourceType: c.ResourceType()})
}

func UnmarshalCdsHooksGuidanceResponse(b []byte) (res CdsHooksGuidanceResponse, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
