// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Claim
type ClaimPayee struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Category of recipient
	Type CodeableConcept `json:"type"`
	// Recipient reference
	Party *custom.Reference[ClaimPayeeParty] `json:"party,omitempty"`
}

type ClaimEvent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Specific event
	Type CodeableConcept `json:"type"`
	// Occurance date or period
	WhenDateTime *custom.DateTime `json:"whenDateTime,omitempty"`
	// Occurance date or period
	WhenPeriod *Period `json:"whenPeriod,omitempty"`
}

type ClaimCareTeam struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Order of care team
	Sequence uint32 `json:"sequence"`
	// Practitioner or organization
	Provider custom.Reference[ClaimCareTeamProvider] `json:"provider"`
	// Indicator of the lead practitioner
	Responsible *bool `json:"responsible,omitempty"`
	// Function within the team
	Role *CodeableConcept `json:"role,omitempty"`
	// Practitioner or provider specialization
	Specialty *CodeableConcept `json:"specialty,omitempty"`
}

type ClaimSupportingInfo struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Information instance identifier
	Sequence uint32 `json:"sequence"`
	// Classification of the supplied information
	Category CodeableConcept `json:"category"`
	// Type of information
	Code *CodeableConcept `json:"code,omitempty"`
	// When it occurred
	TimingDate *custom.Date `json:"timingDate,omitempty"`
	// When it occurred
	TimingPeriod *Period `json:"timingPeriod,omitempty"`
	// Data to be provided
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Data to be provided
	ValueString *string `json:"valueString,omitempty"`
	// Data to be provided
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Data to be provided
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Data to be provided
	ValueReference *custom.Reference[Resource] `json:"valueReference,omitempty"`
	// Data to be provided
	ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`
	// Explanation for the information
	Reason *CodeableConcept `json:"reason,omitempty"`
}

type Claim struct {
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
	// Business Identifier for claim
	Identifier []Identifier `json:"identifier,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `json:"status"`
	// Category or discipline
	Type CodeableConcept `json:"type"`
	// More granular claim type
	SubType *CodeableConcept `json:"subType,omitempty"`
	// claim | preauthorization | predetermination
	Use custom.Code `json:"use"`
	// The recipient of the products and services
	Patient custom.Reference[Patient] `json:"patient"`
	// Relevant time frame for the claim
	BillablePeriod *Period `json:"billablePeriod,omitempty"`
	// Resource creation date
	Created custom.DateTime `json:"created"`
	// Author of the claim
	Enterer *custom.Reference[ClaimEnterer] `json:"enterer,omitempty"`
	// Target
	Insurer *custom.Reference[Organization] `json:"insurer,omitempty"`
	// Party responsible for the claim
	Provider *custom.Reference[ClaimProvider] `json:"provider,omitempty"`
	// Desired processing urgency
	Priority *CodeableConcept `json:"priority,omitempty"`
	// For whom to reserve funds
	FundsReserve *CodeableConcept `json:"fundsReserve,omitempty"`
	// Prior or corollary claims
	Related []ClaimRelated `json:"related,omitempty"`
	// Prescription authorizing services and products
	Prescription *custom.Reference[ClaimPrescription] `json:"prescription,omitempty"`
	// Original prescription if superseded by fulfiller
	OriginalPrescription *custom.Reference[ClaimOriginalPrescription] `json:"originalPrescription,omitempty"`
	// Recipient of benefits payable
	Payee *ClaimPayee `json:"payee,omitempty"`
	// Treatment referral
	Referral *custom.Reference[ServiceRequest] `json:"referral,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Servicing facility
	Facility *custom.Reference[ClaimFacility] `json:"facility,omitempty"`
	// Package billing code
	DiagnosisRelatedGroup *CodeableConcept `json:"diagnosisRelatedGroup,omitempty"`
	// Event information
	Event []ClaimEvent `json:"event,omitempty"`
	// Members of the care team
	CareTeam []ClaimCareTeam `json:"careTeam,omitempty"`
	// Supporting information
	SupportingInfo []ClaimSupportingInfo `json:"supportingInfo,omitempty"`
	// Pertinent diagnosis information
	Diagnosis []ClaimDiagnosis `json:"diagnosis,omitempty"`
	// Clinical procedures performed
	Procedure []ClaimProcedure `json:"procedure,omitempty"`
	// Patient insurance information
	Insurance []ClaimInsurance `json:"insurance,omitempty"`
	// Details of the event
	Accident *ClaimAccident `json:"accident,omitempty"`
	// Paid by the patient
	PatientPaid *Money `json:"patientPaid,omitempty"`
	// Product or service provided
	Item []ClaimItem `json:"item,omitempty"`
	// Total claim cost
	Total *Money `json:"total,omitempty"`
}

type ClaimProcedure struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Procedure instance identifier
	Sequence uint32 `json:"sequence"`
	// Category of Procedure
	Type []CodeableConcept `json:"type,omitempty"`
	// When the procedure was performed
	Date *custom.DateTime `json:"date,omitempty"`
	// Specific clinical procedure
	ProcedureCodeableConcept *CodeableConcept `json:"procedureCodeableConcept,omitempty"`
	// Specific clinical procedure
	ProcedureReference *custom.Reference[Procedure] `json:"procedureReference,omitempty"`
	// Unique device identifier
	Udi []custom.Reference[Device] `json:"udi,omitempty"`
}

type ClaimInsurance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Insurance instance identifier
	Sequence uint32 `json:"sequence"`
	// Coverage to be used for adjudication
	Focal bool `json:"focal"`
	// Pre-assigned Claim number
	Identifier *Identifier `json:"identifier,omitempty"`
	// Insurance information
	Coverage custom.Reference[Coverage] `json:"coverage"`
	// Additional provider contract number
	BusinessArrangement *string `json:"businessArrangement,omitempty"`
	// Prior authorization reference number
	PreAuthRef []string `json:"preAuthRef,omitempty"`
	// Adjudication results
	ClaimResponse *custom.Reference[ClaimResponse] `json:"claimResponse,omitempty"`
}

type ClaimItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Item instance identifier
	Sequence uint32 `json:"sequence"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Applicable careTeam members
	CareTeamSequence []uint32 `json:"careTeamSequence,omitempty"`
	// Applicable diagnoses
	DiagnosisSequence []uint32 `json:"diagnosisSequence,omitempty"`
	// Applicable procedures
	ProcedureSequence []uint32 `json:"procedureSequence,omitempty"`
	// Applicable exception and supporting information
	InformationSequence []uint32 `json:"informationSequence,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `json:"revenue,omitempty"`
	// Benefit classification
	Category *CodeableConcept `json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `json:"productOrServiceEnd,omitempty"`
	// Request or Referral for Service
	Request []custom.Reference[ClaimItemRequest] `json:"request,omitempty"`
	// Product or service billing modifiers
	Modifier []CodeableConcept `json:"modifier,omitempty"`
	// Program the product or service is provided under
	ProgramCode []CodeableConcept `json:"programCode,omitempty"`
	// Date or dates of service or product delivery
	ServicedDate *custom.Date `json:"servicedDate,omitempty"`
	// Date or dates of service or product delivery
	ServicedPeriod *Period `json:"servicedPeriod,omitempty"`
	// Place of service or where product was supplied
	LocationCodeableConcept *CodeableConcept `json:"locationCodeableConcept,omitempty"`
	// Place of service or where product was supplied
	LocationAddress *Address `json:"locationAddress,omitempty"`
	// Place of service or where product was supplied
	LocationReference *custom.Reference[Location] `json:"locationReference,omitempty"`
	// Paid by the patient
	PatientPaid *Money `json:"patientPaid,omitempty"`
	// Count of products or services
	Quantity *Quantity `json:"quantity,omitempty"`
	// Fee, charge or cost per item
	UnitPrice *Money `json:"unitPrice,omitempty"`
	// Price scaling factor
	Factor *json.Number `json:"factor,omitempty"`
	// Total tax
	Tax *Money `json:"tax,omitempty"`
	// Total item cost
	Net *Money `json:"net,omitempty"`
	// Unique device identifier
	Udi []custom.Reference[Device] `json:"udi,omitempty"`
	// Anatomical location
	BodySite []ClaimItemBodySite `json:"bodySite,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Product or service provided
	Detail []ClaimItemDetail `json:"detail,omitempty"`
}

type ClaimItemBodySite struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Location
	Site []custom.CodeableReference[BodyStructure] `json:"site"`
	// Sub-location
	SubSite []CodeableConcept `json:"subSite,omitempty"`
}

type ClaimItemDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Item instance identifier
	Sequence uint32 `json:"sequence"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `json:"revenue,omitempty"`
	// Benefit classification
	Category *CodeableConcept `json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `json:"productOrServiceEnd,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `json:"modifier,omitempty"`
	// Program the product or service is provided under
	ProgramCode []CodeableConcept `json:"programCode,omitempty"`
	// Paid by the patient
	PatientPaid *Money `json:"patientPaid,omitempty"`
	// Count of products or services
	Quantity *Quantity `json:"quantity,omitempty"`
	// Fee, charge or cost per item
	UnitPrice *Money `json:"unitPrice,omitempty"`
	// Price scaling factor
	Factor *json.Number `json:"factor,omitempty"`
	// Total tax
	Tax *Money `json:"tax,omitempty"`
	// Total item cost
	Net *Money `json:"net,omitempty"`
	// Unique device identifier
	Udi []custom.Reference[Device] `json:"udi,omitempty"`
	// Product or service provided
	SubDetail []ClaimItemDetailSubDetail `json:"subDetail,omitempty"`
}

type ClaimItemDetailSubDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Item instance identifier
	Sequence uint32 `json:"sequence"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `json:"revenue,omitempty"`
	// Benefit classification
	Category *CodeableConcept `json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `json:"productOrServiceEnd,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `json:"modifier,omitempty"`
	// Program the product or service is provided under
	ProgramCode []CodeableConcept `json:"programCode,omitempty"`
	// Paid by the patient
	PatientPaid *Money `json:"patientPaid,omitempty"`
	// Count of products or services
	Quantity *Quantity `json:"quantity,omitempty"`
	// Fee, charge or cost per item
	UnitPrice *Money `json:"unitPrice,omitempty"`
	// Price scaling factor
	Factor *json.Number `json:"factor,omitempty"`
	// Total tax
	Tax *Money `json:"tax,omitempty"`
	// Total item cost
	Net *Money `json:"net,omitempty"`
	// Unique device identifier
	Udi []custom.Reference[Device] `json:"udi,omitempty"`
}

type ClaimDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Diagnosis instance identifier
	Sequence uint32 `json:"sequence"`
	// Nature of illness or problem
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	// Nature of illness or problem
	DiagnosisReference *custom.Reference[Condition] `json:"diagnosisReference,omitempty"`
	// Timing or nature of the diagnosis
	Type []CodeableConcept `json:"type,omitempty"`
	// Present on admission
	OnAdmission *CodeableConcept `json:"onAdmission,omitempty"`
}

type ClaimAccident struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// When the incident occurred
	Date custom.Date `json:"date"`
	// The nature of the accident
	Type *CodeableConcept `json:"type,omitempty"`
	// Where the event occurred
	LocationAddress *Address `json:"locationAddress,omitempty"`
	// Where the event occurred
	LocationReference *custom.Reference[Location] `json:"locationReference,omitempty"`
}

type ClaimRelated struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to the related claim
	Claim *custom.Reference[Claim] `json:"claim,omitempty"`
	// How the reference claim is related
	Relationship *CodeableConcept `json:"relationship,omitempty"`
	// File or case reference
	Reference *Identifier `json:"reference,omitempty"`
}

type OtherClaim Claim

type ClaimEnterer interface {
	gofhirclient.Resource

	Is_ClaimEnterer()
}

func (p Practitioner) Is_ClaimEnterer()     {}
func (p PractitionerRole) Is_ClaimEnterer() {}
func (p Patient) Is_ClaimEnterer()          {}
func (r RelatedPerson) Is_ClaimEnterer()    {}

type ClaimProvider interface {
	gofhirclient.Resource

	Is_ClaimProvider()
}

func (p Practitioner) Is_ClaimProvider()     {}
func (p PractitionerRole) Is_ClaimProvider() {}
func (o Organization) Is_ClaimProvider()     {}

type ClaimPrescription interface {
	gofhirclient.Resource

	Is_ClaimPrescription()
}

func (d DeviceRequest) Is_ClaimPrescription()      {}
func (m MedicationRequest) Is_ClaimPrescription()  {}
func (v VisionPrescription) Is_ClaimPrescription() {}

type ClaimOriginalPrescription interface {
	gofhirclient.Resource

	Is_ClaimOriginalPrescription()
}

func (d DeviceRequest) Is_ClaimOriginalPrescription()      {}
func (m MedicationRequest) Is_ClaimOriginalPrescription()  {}
func (v VisionPrescription) Is_ClaimOriginalPrescription() {}

type ClaimPayeeParty interface {
	gofhirclient.Resource

	Is_ClaimPayeeParty()
}

func (p Practitioner) Is_ClaimPayeeParty()     {}
func (p PractitionerRole) Is_ClaimPayeeParty() {}
func (o Organization) Is_ClaimPayeeParty()     {}
func (p Patient) Is_ClaimPayeeParty()          {}
func (r RelatedPerson) Is_ClaimPayeeParty()    {}

type ClaimFacility interface {
	gofhirclient.Resource

	Is_ClaimFacility()
}

func (l Location) Is_ClaimFacility()     {}
func (o Organization) Is_ClaimFacility() {}

type ClaimCareTeamProvider interface {
	gofhirclient.Resource

	Is_ClaimCareTeamProvider()
}

func (p Practitioner) Is_ClaimCareTeamProvider()     {}
func (p PractitionerRole) Is_ClaimCareTeamProvider() {}
func (o Organization) Is_ClaimCareTeamProvider()     {}

type ClaimItemRequest interface {
	gofhirclient.Resource

	Is_ClaimItemRequest()
}

func (d DeviceRequest) Is_ClaimItemRequest()      {}
func (m MedicationRequest) Is_ClaimItemRequest()  {}
func (n NutritionOrder) Is_ClaimItemRequest()     {}
func (s ServiceRequest) Is_ClaimItemRequest()     {}
func (s SupplyRequest) Is_ClaimItemRequest()      {}
func (v VisionPrescription) Is_ClaimItemRequest() {}

func (c Claim) ResourceType() string {
	return "Claim"
}

func (c Claim) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaim
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherClaim: OtherClaim(c), ResourceType: c.ResourceType()})
}

func UnmarshalClaim(b []byte) (res Claim, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
