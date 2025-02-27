// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefitCareTeam struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Order of care team
	Sequence uint32 `json:"sequence"`
	// Practitioner or organization
	Provider custom.Reference[ExplanationOfBenefitCareTeamProvider] `json:"provider"`
	// Indicator of the lead practitioner
	Responsible *bool `json:"responsible,omitempty"`
	// Function within the team
	Role *CodeableConcept `json:"role,omitempty"`
	// Practitioner or provider specialization
	Specialty *CodeableConcept `json:"specialty,omitempty"`
}

type ExplanationOfBenefitDiagnosis struct {
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

type ExplanationOfBenefitItemBodySite struct {
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

type ExplanationOfBenefitItemAdjudication struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of adjudication information
	Category CodeableConcept `json:"category"`
	// Explanation of adjudication outcome
	Reason *CodeableConcept `json:"reason,omitempty"`
	// Monetary amount
	Amount *Money `json:"amount,omitempty"`
	// Non-monitary value
	Quantity *Quantity `json:"quantity,omitempty"`
}

type ExplanationOfBenefitItemDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Product or service provided
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
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Detail level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Detail level adjudication details
	Adjudication []interface{} `json:"adjudication,omitempty"`
	// Additional items
	SubDetail []ExplanationOfBenefitItemDetailSubDetail `json:"subDetail,omitempty"`
}

type ExplanationOfBenefitSupportingInfo struct {
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
	Reason *Coding `json:"reason,omitempty"`
}

type ExplanationOfBenefitProcedure struct {
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

type ExplanationOfBenefitItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Item instance identifier
	Sequence uint32 `json:"sequence"`
	// Applicable care team members
	CareTeamSequence []uint32 `json:"careTeamSequence,omitempty"`
	// Applicable diagnoses
	DiagnosisSequence []uint32 `json:"diagnosisSequence,omitempty"`
	// Applicable procedures
	ProcedureSequence []uint32 `json:"procedureSequence,omitempty"`
	// Applicable exception and supporting information
	InformationSequence []uint32 `json:"informationSequence,omitempty"`
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
	// Request or Referral for Service
	Request []custom.Reference[ExplanationOfBenefitItemRequest] `json:"request,omitempty"`
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
	BodySite []ExplanationOfBenefitItemBodySite `json:"bodySite,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Adjudication results
	ReviewOutcome *ExplanationOfBenefitItemReviewOutcome `json:"reviewOutcome,omitempty"`
	// Adjudication details
	Adjudication []ExplanationOfBenefitItemAdjudication `json:"adjudication,omitempty"`
	// Additional items
	Detail []ExplanationOfBenefitItemDetail `json:"detail,omitempty"`
}

type ExplanationOfBenefitAddItemBodySite struct {
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

type ExplanationOfBenefitProcessNote struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Note instance identifier
	Number *uint32 `json:"number,omitempty"`
	// Note purpose
	Type *CodeableConcept `json:"type,omitempty"`
	// Note explanatory text
	Text *string `json:"text,omitempty"`
	// Language of the text
	Language *CodeableConcept `json:"language,omitempty"`
}

type ExplanationOfBenefitBenefitBalanceFinancial struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Benefit classification
	Type CodeableConcept `json:"type"`
	// Benefits allowed
	AllowedUnsignedInt *uint32 `json:"allowedUnsignedInt,omitempty"`
	// Benefits allowed
	AllowedString *string `json:"allowedString,omitempty"`
	// Benefits allowed
	AllowedMoney *Money `json:"allowedMoney,omitempty"`
	// Benefits used
	UsedUnsignedInt *uint32 `json:"usedUnsignedInt,omitempty"`
	// Benefits used
	UsedMoney *Money `json:"usedMoney,omitempty"`
}

type ExplanationOfBenefitAddItemDetailSubDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `json:"revenue,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `json:"productOrServiceEnd,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `json:"modifier,omitempty"`
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
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Additem subdetail level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Added items adjudication
	Adjudication []interface{} `json:"adjudication,omitempty"`
}

type ExplanationOfBenefitPayment struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Partial or complete payment
	Type *CodeableConcept `json:"type,omitempty"`
	// Payment adjustment for non-claim issues
	Adjustment *Money `json:"adjustment,omitempty"`
	// Explanation for the variance
	AdjustmentReason *CodeableConcept `json:"adjustmentReason,omitempty"`
	// Expected date of payment
	Date *custom.Date `json:"date,omitempty"`
	// Payable amount after adjustment
	Amount *Money `json:"amount,omitempty"`
	// Business identifier for the payment
	Identifier *Identifier `json:"identifier,omitempty"`
}

type ExplanationOfBenefit struct {
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
	// Business Identifier for the resource
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
	// Response creation date
	Created custom.DateTime `json:"created"`
	// Author of the claim
	Enterer *custom.Reference[ExplanationOfBenefitEnterer] `json:"enterer,omitempty"`
	// Party responsible for reimbursement
	Insurer *custom.Reference[Organization] `json:"insurer,omitempty"`
	// Party responsible for the claim
	Provider *custom.Reference[ExplanationOfBenefitProvider] `json:"provider,omitempty"`
	// Desired processing urgency
	Priority *CodeableConcept `json:"priority,omitempty"`
	// For whom to reserve funds
	FundsReserveRequested *CodeableConcept `json:"fundsReserveRequested,omitempty"`
	// Funds reserved status
	FundsReserve *CodeableConcept `json:"fundsReserve,omitempty"`
	// Prior or corollary claims
	Related []ExplanationOfBenefitRelated `json:"related,omitempty"`
	// Prescription authorizing services or products
	Prescription *custom.Reference[ExplanationOfBenefitPrescription] `json:"prescription,omitempty"`
	// Original prescription if superceded by fulfiller
	OriginalPrescription *custom.Reference[MedicationRequest] `json:"originalPrescription,omitempty"`
	// Event information
	Event []ExplanationOfBenefitEvent `json:"event,omitempty"`
	// Recipient of benefits payable
	Payee *ExplanationOfBenefitPayee `json:"payee,omitempty"`
	// Treatment Referral
	Referral *custom.Reference[ServiceRequest] `json:"referral,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Servicing Facility
	Facility *custom.Reference[ExplanationOfBenefitFacility] `json:"facility,omitempty"`
	// Claim reference
	Claim *custom.Reference[Claim] `json:"claim,omitempty"`
	// Claim response reference
	ClaimResponse *custom.Reference[ClaimResponse] `json:"claimResponse,omitempty"`
	// queued | complete | error | partial
	Outcome custom.Code `json:"outcome"`
	// Result of the adjudication
	Decision *CodeableConcept `json:"decision,omitempty"`
	// Disposition Message
	Disposition *string `json:"disposition,omitempty"`
	// Preauthorization reference
	PreAuthRef []string `json:"preAuthRef,omitempty"`
	// Preauthorization in-effect period
	PreAuthRefPeriod []Period `json:"preAuthRefPeriod,omitempty"`
	// Package billing code
	DiagnosisRelatedGroup *CodeableConcept `json:"diagnosisRelatedGroup,omitempty"`
	// Care Team members
	CareTeam []ExplanationOfBenefitCareTeam `json:"careTeam,omitempty"`
	// Supporting information
	SupportingInfo []ExplanationOfBenefitSupportingInfo `json:"supportingInfo,omitempty"`
	// Pertinent diagnosis information
	Diagnosis []ExplanationOfBenefitDiagnosis `json:"diagnosis,omitempty"`
	// Clinical procedures performed
	Procedure []ExplanationOfBenefitProcedure `json:"procedure,omitempty"`
	// Precedence (primary, secondary, etc.)
	Precedence *uint32 `json:"precedence,omitempty"`
	// Patient insurance information
	Insurance []ExplanationOfBenefitInsurance `json:"insurance,omitempty"`
	// Details of the event
	Accident *ExplanationOfBenefitAccident `json:"accident,omitempty"`
	// Paid by the patient
	PatientPaid *Money `json:"patientPaid,omitempty"`
	// Product or service provided
	Item []ExplanationOfBenefitItem `json:"item,omitempty"`
	// Insurer added line items
	AddItem []ExplanationOfBenefitAddItem `json:"addItem,omitempty"`
	// Header-level adjudication
	Adjudication []interface{} `json:"adjudication,omitempty"`
	// Adjudication totals
	Total []ExplanationOfBenefitTotal `json:"total,omitempty"`
	// Payment Details
	Payment *ExplanationOfBenefitPayment `json:"payment,omitempty"`
	// Printed form identifier
	FormCode *CodeableConcept `json:"formCode,omitempty"`
	// Printed reference or actual form
	Form *Attachment `json:"form,omitempty"`
	// Note concerning adjudication
	ProcessNote []ExplanationOfBenefitProcessNote `json:"processNote,omitempty"`
	// When the benefits are applicable
	BenefitPeriod *Period `json:"benefitPeriod,omitempty"`
	// Balance by Benefit Category
	BenefitBalance []ExplanationOfBenefitBenefitBalance `json:"benefitBalance,omitempty"`
}

type ExplanationOfBenefitRelated struct {
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

type ExplanationOfBenefitPayee struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Category of recipient
	Type *CodeableConcept `json:"type,omitempty"`
	// Recipient reference
	Party *custom.Reference[ExplanationOfBenefitPayeeParty] `json:"party,omitempty"`
}

type ExplanationOfBenefitItemReviewOutcome struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Result of the adjudication
	Decision *CodeableConcept `json:"decision,omitempty"`
	// Reason for result of the adjudication
	Reason []CodeableConcept `json:"reason,omitempty"`
	// Preauthorization reference
	PreAuthRef *string `json:"preAuthRef,omitempty"`
	// Preauthorization reference effective period
	PreAuthPeriod *Period `json:"preAuthPeriod,omitempty"`
}

type ExplanationOfBenefitItemDetailSubDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Product or service provided
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
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Subdetail level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Subdetail level adjudication details
	Adjudication []interface{} `json:"adjudication,omitempty"`
}

type ExplanationOfBenefitAddItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Item sequence number
	ItemSequence []uint32 `json:"itemSequence,omitempty"`
	// Detail sequence number
	DetailSequence []uint32 `json:"detailSequence,omitempty"`
	// Subdetail sequence number
	SubDetailSequence []uint32 `json:"subDetailSequence,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Authorized providers
	Provider []custom.Reference[ExplanationOfBenefitAddItemProvider] `json:"provider,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `json:"revenue,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `json:"productOrServiceEnd,omitempty"`
	// Request or Referral for Service
	Request []custom.Reference[ExplanationOfBenefitAddItemRequest] `json:"request,omitempty"`
	// Service/Product billing modifiers
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
	// Anatomical location
	BodySite []ExplanationOfBenefitAddItemBodySite `json:"bodySite,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Additem level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Added items adjudication
	Adjudication []interface{} `json:"adjudication,omitempty"`
	// Insurer added line items
	Detail []ExplanationOfBenefitAddItemDetail `json:"detail,omitempty"`
}

type ExplanationOfBenefitEvent struct {
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

type ExplanationOfBenefitInsurance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Coverage to be used for adjudication
	Focal bool `json:"focal"`
	// Insurance information
	Coverage custom.Reference[Coverage] `json:"coverage"`
	// Prior authorization reference number
	PreAuthRef []string `json:"preAuthRef,omitempty"`
}

type ExplanationOfBenefitAccident struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// When the incident occurred
	Date *custom.Date `json:"date,omitempty"`
	// The nature of the accident
	Type *CodeableConcept `json:"type,omitempty"`
	// Where the event occurred
	LocationAddress *Address `json:"locationAddress,omitempty"`
	// Where the event occurred
	LocationReference *custom.Reference[Location] `json:"locationReference,omitempty"`
}

type ExplanationOfBenefitAddItemDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `json:"revenue,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `json:"productOrServiceEnd,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `json:"modifier,omitempty"`
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
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Additem detail level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Added items adjudication
	Adjudication []interface{} `json:"adjudication,omitempty"`
	// Insurer added line items
	SubDetail []ExplanationOfBenefitAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

type ExplanationOfBenefitTotal struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of adjudication information
	Category CodeableConcept `json:"category"`
	// Financial total for the category
	Amount Money `json:"amount"`
}

type ExplanationOfBenefitBenefitBalance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Benefit classification
	Category CodeableConcept `json:"category"`
	// Excluded from the plan
	Excluded *bool `json:"excluded,omitempty"`
	// Short name for the benefit
	Name *string `json:"name,omitempty"`
	// Description of the benefit or services covered
	Description *string `json:"description,omitempty"`
	// In or out of network
	Network *CodeableConcept `json:"network,omitempty"`
	// Individual or family
	Unit *CodeableConcept `json:"unit,omitempty"`
	// Annual or lifetime
	Term *CodeableConcept `json:"term,omitempty"`
	// Benefit Summary
	Financial []ExplanationOfBenefitBenefitBalanceFinancial `json:"financial,omitempty"`
}

type OtherExplanationOfBenefit ExplanationOfBenefit

type ExplanationOfBenefitEnterer interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitEnterer()
}

func (p Practitioner) Is_ExplanationOfBenefitEnterer()     {}
func (p PractitionerRole) Is_ExplanationOfBenefitEnterer() {}
func (p Patient) Is_ExplanationOfBenefitEnterer()          {}
func (r RelatedPerson) Is_ExplanationOfBenefitEnterer()    {}

type ExplanationOfBenefitProvider interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitProvider()
}

func (p Practitioner) Is_ExplanationOfBenefitProvider()     {}
func (p PractitionerRole) Is_ExplanationOfBenefitProvider() {}
func (o Organization) Is_ExplanationOfBenefitProvider()     {}

type ExplanationOfBenefitPrescription interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitPrescription()
}

func (m MedicationRequest) Is_ExplanationOfBenefitPrescription()  {}
func (v VisionPrescription) Is_ExplanationOfBenefitPrescription() {}

type ExplanationOfBenefitPayeeParty interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitPayeeParty()
}

func (p Practitioner) Is_ExplanationOfBenefitPayeeParty()     {}
func (p PractitionerRole) Is_ExplanationOfBenefitPayeeParty() {}
func (o Organization) Is_ExplanationOfBenefitPayeeParty()     {}
func (p Patient) Is_ExplanationOfBenefitPayeeParty()          {}
func (r RelatedPerson) Is_ExplanationOfBenefitPayeeParty()    {}

type ExplanationOfBenefitFacility interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitFacility()
}

func (l Location) Is_ExplanationOfBenefitFacility()     {}
func (o Organization) Is_ExplanationOfBenefitFacility() {}

type ExplanationOfBenefitCareTeamProvider interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitCareTeamProvider()
}

func (p Practitioner) Is_ExplanationOfBenefitCareTeamProvider()     {}
func (p PractitionerRole) Is_ExplanationOfBenefitCareTeamProvider() {}
func (o Organization) Is_ExplanationOfBenefitCareTeamProvider()     {}

type ExplanationOfBenefitItemRequest interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitItemRequest()
}

func (d DeviceRequest) Is_ExplanationOfBenefitItemRequest()      {}
func (m MedicationRequest) Is_ExplanationOfBenefitItemRequest()  {}
func (n NutritionOrder) Is_ExplanationOfBenefitItemRequest()     {}
func (s ServiceRequest) Is_ExplanationOfBenefitItemRequest()     {}
func (s SupplyRequest) Is_ExplanationOfBenefitItemRequest()      {}
func (v VisionPrescription) Is_ExplanationOfBenefitItemRequest() {}

type ExplanationOfBenefitAddItemProvider interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitAddItemProvider()
}

func (p Practitioner) Is_ExplanationOfBenefitAddItemProvider()     {}
func (p PractitionerRole) Is_ExplanationOfBenefitAddItemProvider() {}
func (o Organization) Is_ExplanationOfBenefitAddItemProvider()     {}

type ExplanationOfBenefitAddItemRequest interface {
	gofhirclient.Resource

	Is_ExplanationOfBenefitAddItemRequest()
}

func (d DeviceRequest) Is_ExplanationOfBenefitAddItemRequest()      {}
func (m MedicationRequest) Is_ExplanationOfBenefitAddItemRequest()  {}
func (n NutritionOrder) Is_ExplanationOfBenefitAddItemRequest()     {}
func (s ServiceRequest) Is_ExplanationOfBenefitAddItemRequest()     {}
func (s SupplyRequest) Is_ExplanationOfBenefitAddItemRequest()      {}
func (v VisionPrescription) Is_ExplanationOfBenefitAddItemRequest() {}

func (e ExplanationOfBenefit) ResourceType() string {
	return "ExplanationOfBenefit"
}

func (e ExplanationOfBenefit) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExplanationOfBenefit
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherExplanationOfBenefit: OtherExplanationOfBenefit(e), ResourceType: e.ResourceType()})
}

func UnmarshalExplanationOfBenefit(b []byte) (res ExplanationOfBenefit, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
