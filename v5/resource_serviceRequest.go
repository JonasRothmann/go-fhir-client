// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ServiceRequest
type ServiceRequest struct {
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
	InstantiatesCanonical []custom.Canonical[ServiceRequestInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// What request fulfills
	BasedOn []custom.Reference[ServiceRequestBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// What request replaces
	Replaces []custom.Reference[ServiceRequest] `bson:"replaces,omitempty" json:"replaces,omitempty"`
	// Composite Request ID
	Requisition *Identifier `bson:"requisition,omitempty" json:"requisition,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// proposal | plan | directive | order +
	Intent custom.Code `bson:"intent" json:"intent"`
	// Classification of service
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// True if service/procedure should not be performed
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	// What is being requested/ordered
	Code *custom.CodeableReference[ServiceRequestCode] `bson:"code,omitempty" json:"code,omitempty"`
	// Additional order information
	OrderDetail []ServiceRequestOrderDetail `bson:"orderDetail,omitempty" json:"orderDetail,omitempty"`
	// Service amount
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Individual or Entity the service is ordered for
	Subject custom.Reference[ServiceRequestSubject] `bson:"subject" json:"subject"`
	// What the service request is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// Encounter in which the request was created
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When service should occur
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// Preconditions for service
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
	// Date request signed
	AuthoredOn *custom.DateTime `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	// Who/what is requesting service
	Requester *custom.Reference[ServiceRequestRequester] `bson:"requester,omitempty" json:"requester,omitempty"`
	// Performer role
	PerformerType *CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty"`
	// Requested performer
	Performer []custom.Reference[ServiceRequestPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Requested location
	Location []custom.CodeableReference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Explanation/Justification for procedure or service
	Reason []custom.CodeableReference[ServiceRequestReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[ServiceRequestInsurance] `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Additional clinical information
	SupportingInfo []custom.CodeableReference[Resource] `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Procedure Samples
	Specimen []custom.Reference[Specimen] `bson:"specimen,omitempty" json:"specimen,omitempty"`
	// Coded location on Body
	BodySite []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// BodyStructure-based location on the body
	BodyStructure *custom.Reference[BodyStructure] `bson:"bodyStructure,omitempty" json:"bodyStructure,omitempty"`
	// Comments
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Patient or consumer-oriented instructions
	PatientInstruction []ServiceRequestPatientInstruction `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	// Request provenance
	RelevantHistory []custom.Reference[Provenance] `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}

type ServiceRequestOrderDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The context of the order details by reference
	ParameterFocus *custom.CodeableReference[ServiceRequestOrderDetailParameterFocus] `bson:"parameterFocus,omitempty" json:"parameterFocus,omitempty"`
	// The parameter details for the service being requested
	Parameter []ServiceRequestOrderDetailParameter `bson:"parameter" json:"parameter"`
}

type ServiceRequestOrderDetailParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The detail of the order being requested
	Code CodeableConcept `bson:"code" json:"code"`
	// The value for the order detail
	Value Quantity `bson:"value" json:"value"`
}

type ServiceRequestPatientInstruction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Patient or consumer-oriented instructions
	Instruction *custom.Markdown `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

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
	return "StructureDefinition"
}
