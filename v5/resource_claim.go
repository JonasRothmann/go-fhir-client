// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Claim
type ClaimItemBodySite struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Location
	Site []custom.CodeableReference[BodyStructure] `bson:"site" json:"site"`
	// Sub-location
	SubSite []CodeableConcept `bson:"subSite,omitempty" json:"subSite,omitempty"`
}

type ClaimDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Diagnosis instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Nature of illness or problem
	Diagnosis CodeableConcept `bson:"diagnosis" json:"diagnosis"`
	// Timing or nature of the diagnosis
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Present on admission
	OnAdmission *CodeableConcept `bson:"onAdmission,omitempty" json:"onAdmission,omitempty"`
}

type ClaimItemDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Item instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty"`
	// Benefit classification
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Program the product or service is provided under
	ProgramCode []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	// Paid by the patient
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty"`
	// Count of products or services
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Fee, charge or cost per item
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	// Price scaling factor
	Factor *json.Number `bson:"factor,omitempty" json:"factor,omitempty"`
	// Total tax
	Tax *Money `bson:"tax,omitempty" json:"tax,omitempty"`
	// Total item cost
	Net *Money `bson:"net,omitempty" json:"net,omitempty"`
	// Unique device identifier
	Udi []custom.Reference[Device] `bson:"udi,omitempty" json:"udi,omitempty"`
	// Product or service provided
	SubDetail []ClaimItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimPayee struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Category of recipient
	Type CodeableConcept `bson:"type" json:"type"`
	// Recipient reference
	Party *custom.Reference[ClaimPayeeParty] `bson:"party,omitempty" json:"party,omitempty"`
}

type ClaimCareTeam struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Order of care team
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Practitioner or organization
	Provider custom.Reference[ClaimCareTeamProvider] `bson:"provider" json:"provider"`
	// Indicator of the lead practitioner
	Responsible *bool `bson:"responsible,omitempty" json:"responsible,omitempty"`
	// Function within the team
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Practitioner or provider specialization
	Specialty *CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
}

type ClaimItemDetailSubDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Item instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty"`
	// Benefit classification
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Program the product or service is provided under
	ProgramCode []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	// Paid by the patient
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty"`
	// Count of products or services
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Fee, charge or cost per item
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	// Price scaling factor
	Factor *json.Number `bson:"factor,omitempty" json:"factor,omitempty"`
	// Total tax
	Tax *Money `bson:"tax,omitempty" json:"tax,omitempty"`
	// Total item cost
	Net *Money `bson:"net,omitempty" json:"net,omitempty"`
	// Unique device identifier
	Udi []custom.Reference[Device] `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ClaimEvent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Specific event
	Type CodeableConcept `bson:"type" json:"type"`
	// Occurance date or period
	When custom.DateTime `bson:"when" json:"when"`
}

type ClaimItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Item instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Applicable careTeam members
	CareTeamSequence []uint32 `bson:"careTeamSequence,omitempty" json:"careTeamSequence,omitempty"`
	// Applicable diagnoses
	DiagnosisSequence []uint32 `bson:"diagnosisSequence,omitempty" json:"diagnosisSequence,omitempty"`
	// Applicable procedures
	ProcedureSequence []uint32 `bson:"procedureSequence,omitempty" json:"procedureSequence,omitempty"`
	// Applicable exception and supporting information
	InformationSequence []uint32 `bson:"informationSequence,omitempty" json:"informationSequence,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty"`
	// Benefit classification
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty"`
	// Request or Referral for Service
	Request []custom.Reference[ClaimItemRequest] `bson:"request,omitempty" json:"request,omitempty"`
	// Product or service billing modifiers
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Program the product or service is provided under
	ProgramCode []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	// Date or dates of service or product delivery
	Serviced *custom.Date `bson:"serviced,omitempty" json:"serviced,omitempty"`
	// Place of service or where product was supplied
	Location *CodeableConcept `bson:"location,omitempty" json:"location,omitempty"`
	// Paid by the patient
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty"`
	// Count of products or services
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Fee, charge or cost per item
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	// Price scaling factor
	Factor *json.Number `bson:"factor,omitempty" json:"factor,omitempty"`
	// Total tax
	Tax *Money `bson:"tax,omitempty" json:"tax,omitempty"`
	// Total item cost
	Net *Money `bson:"net,omitempty" json:"net,omitempty"`
	// Unique device identifier
	Udi []custom.Reference[Device] `bson:"udi,omitempty" json:"udi,omitempty"`
	// Anatomical location
	BodySite []ClaimItemBodySite `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Product or service provided
	Detail []ClaimItemDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimSupportingInfo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Information instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Classification of the supplied information
	Category CodeableConcept `bson:"category" json:"category"`
	// Type of information
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// When it occurred
	Timing *custom.Date `bson:"timing,omitempty" json:"timing,omitempty"`
	// Data to be provided
	Value *bool `bson:"value,omitempty" json:"value,omitempty"`
	// Explanation for the information
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ClaimProcedure struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Procedure instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Category of Procedure
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// When the procedure was performed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Specific clinical procedure
	Procedure CodeableConcept `bson:"procedure" json:"procedure"`
	// Unique device identifier
	Udi []custom.Reference[Device] `bson:"udi,omitempty" json:"udi,omitempty"`
}

type ClaimInsurance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Insurance instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Coverage to be used for adjudication
	Focal bool `bson:"focal" json:"focal"`
	// Pre-assigned Claim number
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Insurance information
	Coverage custom.Reference[Coverage] `bson:"coverage" json:"coverage"`
	// Additional provider contract number
	BusinessArrangement *string `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	// Prior authorization reference number
	PreAuthRef []string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	// Adjudication results
	ClaimResponse *custom.Reference[ClaimResponse] `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}

type ClaimAccident struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// When the incident occurred
	Date custom.Date `bson:"date" json:"date"`
	// The nature of the accident
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Where the event occurred
	Location *Address `bson:"location,omitempty" json:"location,omitempty"`
}

type Claim struct {
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
	// Business Identifier for claim
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Category or discipline
	Type CodeableConcept `bson:"type" json:"type"`
	// More granular claim type
	SubType *CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty"`
	// claim | preauthorization | predetermination
	Use custom.Code `bson:"use" json:"use"`
	// The recipient of the products and services
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Relevant time frame for the claim
	BillablePeriod *Period `bson:"billablePeriod,omitempty" json:"billablePeriod,omitempty"`
	// Resource creation date
	Created custom.DateTime `bson:"created" json:"created"`
	// Author of the claim
	Enterer *custom.Reference[ClaimEnterer] `bson:"enterer,omitempty" json:"enterer,omitempty"`
	// Target
	Insurer *custom.Reference[Organization] `bson:"insurer,omitempty" json:"insurer,omitempty"`
	// Party responsible for the claim
	Provider *custom.Reference[ClaimProvider] `bson:"provider,omitempty" json:"provider,omitempty"`
	// Desired processing urgency
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty"`
	// For whom to reserve funds
	FundsReserve *CodeableConcept `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	// Prior or corollary claims
	Related []ClaimRelated `bson:"related,omitempty" json:"related,omitempty"`
	// Prescription authorizing services and products
	Prescription *custom.Reference[ClaimPrescription] `bson:"prescription,omitempty" json:"prescription,omitempty"`
	// Original prescription if superseded by fulfiller
	OriginalPrescription *custom.Reference[ClaimOriginalPrescription] `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	// Recipient of benefits payable
	Payee *ClaimPayee `bson:"payee,omitempty" json:"payee,omitempty"`
	// Treatment referral
	Referral *custom.Reference[ServiceRequest] `bson:"referral,omitempty" json:"referral,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Servicing facility
	Facility *custom.Reference[ClaimFacility] `bson:"facility,omitempty" json:"facility,omitempty"`
	// Package billing code
	DiagnosisRelatedGroup *CodeableConcept `bson:"diagnosisRelatedGroup,omitempty" json:"diagnosisRelatedGroup,omitempty"`
	// Event information
	Event []ClaimEvent `bson:"event,omitempty" json:"event,omitempty"`
	// Members of the care team
	CareTeam []ClaimCareTeam `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	// Supporting information
	SupportingInfo []ClaimSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Pertinent diagnosis information
	Diagnosis []ClaimDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	// Clinical procedures performed
	Procedure []ClaimProcedure `bson:"procedure,omitempty" json:"procedure,omitempty"`
	// Patient insurance information
	Insurance []ClaimInsurance `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Details of the event
	Accident *ClaimAccident `bson:"accident,omitempty" json:"accident,omitempty"`
	// Paid by the patient
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty"`
	// Product or service provided
	Item []ClaimItem `bson:"item,omitempty" json:"item,omitempty"`
	// Total claim cost
	Total *Money `bson:"total,omitempty" json:"total,omitempty"`
}

type ClaimRelated struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to the related claim
	Claim *custom.Reference[Claim] `bson:"claim,omitempty" json:"claim,omitempty"`
	// How the reference claim is related
	Relationship *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	// File or case reference
	Reference *Identifier `bson:"reference,omitempty" json:"reference,omitempty"`
}

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
	return "StructureDefinition"
}
