// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ServiceRequest
type ServiceRequest struct {
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
	InstantiatesCanonical []custom.Canonical[ServiceRequestInstantiatesCanonical] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `json:"instantiatesUri,omitempty"`
	// What request fulfills
	BasedOn []custom.Reference[ServiceRequestBasedOn] `json:"basedOn,omitempty"`
	// What request replaces
	Replaces []custom.Reference[ServiceRequest] `json:"replaces,omitempty"`
	// Composite Request ID
	Requisition *Identifier `json:"requisition,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// proposal | plan | directive | order +
	Intent custom.Code `json:"intent"`
	// Classification of service
	Category []CodeableConcept `json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// True if service/procedure should not be performed
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	// What is being requested/ordered
	Code *custom.CodeableReference[ServiceRequestCode] `json:"code,omitempty"`
	// Additional order information
	OrderDetail []ServiceRequestOrderDetail `json:"orderDetail,omitempty"`
	// Service amount
	QuantityQuantity *Quantity `json:"quantityQuantity,omitempty"`
	// Service amount
	QuantityRatio *Ratio `json:"quantityRatio,omitempty"`
	// Service amount
	QuantityRange *Range `json:"quantityRange,omitempty"`
	// Individual or Entity the service is ordered for
	Subject custom.Reference[ServiceRequestSubject] `json:"subject"`
	// What the service request is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `json:"focus,omitempty"`
	// Encounter in which the request was created
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When service should occur
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// When service should occur
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When service should occur
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	// Preconditions for service
	AsNeededBoolean *bool `json:"asNeededBoolean,omitempty"`
	// Preconditions for service
	AsNeededCodeableConcept *CodeableConcept `json:"asNeededCodeableConcept,omitempty"`
	// Date request signed
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Who/what is requesting service
	Requester *custom.Reference[ServiceRequestRequester] `json:"requester,omitempty"`
	// Performer role
	PerformerType *CodeableConcept `json:"performerType,omitempty"`
	// Requested performer
	Performer []custom.Reference[ServiceRequestPerformer] `json:"performer,omitempty"`
	// Requested location
	Location []custom.CodeableReference[Location] `json:"location,omitempty"`
	// Explanation/Justification for procedure or service
	Reason []custom.CodeableReference[ServiceRequestReason] `json:"reason,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[ServiceRequestInsurance] `json:"insurance,omitempty"`
	// Additional clinical information
	SupportingInfo []custom.CodeableReference[Resource] `json:"supportingInfo,omitempty"`
	// Procedure Samples
	Specimen []custom.Reference[Specimen] `json:"specimen,omitempty"`
	// Coded location on Body
	BodySite []CodeableConcept `json:"bodySite,omitempty"`
	// BodyStructure-based location on the body
	BodyStructure *custom.Reference[BodyStructure] `json:"bodyStructure,omitempty"`
	// Comments
	Note []Annotation `json:"note,omitempty"`
	// Patient or consumer-oriented instructions
	PatientInstruction []ServiceRequestPatientInstruction `json:"patientInstruction,omitempty"`
	// Request provenance
	RelevantHistory []custom.Reference[Provenance] `json:"relevantHistory,omitempty"`
}

type ServiceRequestOrderDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The context of the order details by reference
	ParameterFocus *custom.CodeableReference[ServiceRequestOrderDetailParameterFocus] `json:"parameterFocus,omitempty"`
	// The parameter details for the service being requested
	Parameter []ServiceRequestOrderDetailParameter `json:"parameter"`
}

type ServiceRequestOrderDetailParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The detail of the order being requested
	Code CodeableConcept `json:"code"`
	// The value for the order detail
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The value for the order detail
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// The value for the order detail
	ValueRange *Range `json:"valueRange,omitempty"`
	// The value for the order detail
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// The value for the order detail
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// The value for the order detail
	ValueString *string `json:"valueString,omitempty"`
	// The value for the order detail
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
}

type ServiceRequestPatientInstruction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Patient or consumer-oriented instructions
	InstructionMarkdown *custom.Markdown `json:"instructionMarkdown,omitempty"`
	// Patient or consumer-oriented instructions
	InstructionReference *custom.Reference[DocumentReference] `json:"instructionReference,omitempty"`
}

type OtherServiceRequest ServiceRequest

type ServiceRequestInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_ServiceRequestInstantiatesCanonical()
}

func (a ActivityDefinition) Is_ServiceRequestInstantiatesCanonical() {}
func (p PlanDefinition) Is_ServiceRequestInstantiatesCanonical()     {}

type ServiceRequestBasedOn interface {
	gofhirclient.Resource

	Is_ServiceRequestBasedOn()
}

func (c CarePlan) Is_ServiceRequestBasedOn()          {}
func (s ServiceRequest) Is_ServiceRequestBasedOn()    {}
func (m MedicationRequest) Is_ServiceRequestBasedOn() {}

type ServiceRequestCode interface {
	gofhirclient.Resource

	Is_ServiceRequestCode()
}

func (a ActivityDefinition) Is_ServiceRequestCode() {}
func (p PlanDefinition) Is_ServiceRequestCode()     {}

type ServiceRequestOrderDetailParameterFocus interface {
	gofhirclient.Resource

	Is_ServiceRequestOrderDetailParameterFocus()
}

func (d Device) Is_ServiceRequestOrderDetailParameterFocus()                     {}
func (d DeviceDefinition) Is_ServiceRequestOrderDetailParameterFocus()           {}
func (d DeviceRequest) Is_ServiceRequestOrderDetailParameterFocus()              {}
func (s SupplyRequest) Is_ServiceRequestOrderDetailParameterFocus()              {}
func (m Medication) Is_ServiceRequestOrderDetailParameterFocus()                 {}
func (m MedicationRequest) Is_ServiceRequestOrderDetailParameterFocus()          {}
func (b BiologicallyDerivedProduct) Is_ServiceRequestOrderDetailParameterFocus() {}
func (s Substance) Is_ServiceRequestOrderDetailParameterFocus()                  {}

type ServiceRequestSubject interface {
	gofhirclient.Resource

	Is_ServiceRequestSubject()
}

func (p Patient) Is_ServiceRequestSubject()  {}
func (g Group) Is_ServiceRequestSubject()    {}
func (l Location) Is_ServiceRequestSubject() {}
func (d Device) Is_ServiceRequestSubject()   {}

type ServiceRequestRequester interface {
	gofhirclient.Resource

	Is_ServiceRequestRequester()
}

func (p Practitioner) Is_ServiceRequestRequester()     {}
func (p PractitionerRole) Is_ServiceRequestRequester() {}
func (o Organization) Is_ServiceRequestRequester()     {}
func (p Patient) Is_ServiceRequestRequester()          {}
func (r RelatedPerson) Is_ServiceRequestRequester()    {}
func (d Device) Is_ServiceRequestRequester()           {}

type ServiceRequestPerformer interface {
	gofhirclient.Resource

	Is_ServiceRequestPerformer()
}

func (p Practitioner) Is_ServiceRequestPerformer()      {}
func (p PractitionerRole) Is_ServiceRequestPerformer()  {}
func (o Organization) Is_ServiceRequestPerformer()      {}
func (c CareTeam) Is_ServiceRequestPerformer()          {}
func (h HealthcareService) Is_ServiceRequestPerformer() {}
func (p Patient) Is_ServiceRequestPerformer()           {}
func (d Device) Is_ServiceRequestPerformer()            {}
func (r RelatedPerson) Is_ServiceRequestPerformer()     {}

type ServiceRequestReason interface {
	gofhirclient.Resource

	Is_ServiceRequestReason()
}

func (c Condition) Is_ServiceRequestReason()         {}
func (o Observation) Is_ServiceRequestReason()       {}
func (d DiagnosticReport) Is_ServiceRequestReason()  {}
func (d DocumentReference) Is_ServiceRequestReason() {}
func (d DetectedIssue) Is_ServiceRequestReason()     {}

type ServiceRequestInsurance interface {
	gofhirclient.Resource

	Is_ServiceRequestInsurance()
}

func (c Coverage) Is_ServiceRequestInsurance()      {}
func (c ClaimResponse) Is_ServiceRequestInsurance() {}

func (s ServiceRequest) ResourceType() string {
	return "ServiceRequest"
}

func (s ServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherServiceRequest
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherServiceRequest: OtherServiceRequest(s), ResourceType: s.ResourceType()})
}

func UnmarshalServiceRequest(b []byte) (res ServiceRequest, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
