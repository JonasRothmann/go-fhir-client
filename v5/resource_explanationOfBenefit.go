// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ExplanationOfBenefit
type ExplanationOfBenefit struct {
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
	// Business Identifier for the resource
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
	// Response creation date
	Created custom.DateTime `bson:"created" json:"created"`
	// Author of the claim
	Enterer *custom.Reference[ExplanationOfBenefitEnterer] `bson:"enterer,omitempty" json:"enterer,omitempty"`
	// Party responsible for reimbursement
	Insurer *custom.Reference[Organization] `bson:"insurer,omitempty" json:"insurer,omitempty"`
	// Party responsible for the claim
	Provider *custom.Reference[ExplanationOfBenefitProvider] `bson:"provider,omitempty" json:"provider,omitempty"`
	// Desired processing urgency
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty"`
	// For whom to reserve funds
	FundsReserveRequested *CodeableConcept `bson:"fundsReserveRequested,omitempty" json:"fundsReserveRequested,omitempty"`
	// Funds reserved status
	FundsReserve *CodeableConcept `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	// Prior or corollary claims
	Related []ExplanationOfBenefitRelated `bson:"related,omitempty" json:"related,omitempty"`
	// Prescription authorizing services or products
	Prescription *custom.Reference[ExplanationOfBenefitPrescription] `bson:"prescription,omitempty" json:"prescription,omitempty"`
	// Original prescription if superceded by fulfiller
	OriginalPrescription *custom.Reference[MedicationRequest] `bson:"originalPrescription,omitempty" json:"originalPrescription,omitempty"`
	// Event information
	Event []ExplanationOfBenefitEvent `bson:"event,omitempty" json:"event,omitempty"`
	// Recipient of benefits payable
	Payee *ExplanationOfBenefitPayee `bson:"payee,omitempty" json:"payee,omitempty"`
	// Treatment Referral
	Referral *custom.Reference[ServiceRequest] `bson:"referral,omitempty" json:"referral,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Servicing Facility
	Facility *custom.Reference[ExplanationOfBenefitFacility] `bson:"facility,omitempty" json:"facility,omitempty"`
	// Claim reference
	Claim *custom.Reference[Claim] `bson:"claim,omitempty" json:"claim,omitempty"`
	// Claim response reference
	ClaimResponse *custom.Reference[ClaimResponse] `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
	// queued | complete | error | partial
	Outcome custom.Code `bson:"outcome" json:"outcome"`
	// Result of the adjudication
	Decision *CodeableConcept `bson:"decision,omitempty" json:"decision,omitempty"`
	// Disposition Message
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty"`
	// Preauthorization reference
	PreAuthRef []string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	// Preauthorization in-effect period
	PreAuthRefPeriod []Period `bson:"preAuthRefPeriod,omitempty" json:"preAuthRefPeriod,omitempty"`
	// Package billing code
	DiagnosisRelatedGroup *CodeableConcept `bson:"diagnosisRelatedGroup,omitempty" json:"diagnosisRelatedGroup,omitempty"`
	// Care Team members
	CareTeam []ExplanationOfBenefitCareTeam `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	// Supporting information
	SupportingInfo []ExplanationOfBenefitSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Pertinent diagnosis information
	Diagnosis []ExplanationOfBenefitDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	// Clinical procedures performed
	Procedure []ExplanationOfBenefitProcedure `bson:"procedure,omitempty" json:"procedure,omitempty"`
	// Precedence (primary, secondary, etc.)
	Precedence *uint32 `bson:"precedence,omitempty" json:"precedence,omitempty"`
	// Patient insurance information
	Insurance []ExplanationOfBenefitInsurance `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Details of the event
	Accident *ExplanationOfBenefitAccident `bson:"accident,omitempty" json:"accident,omitempty"`
	// Paid by the patient
	PatientPaid *Money `bson:"patientPaid,omitempty" json:"patientPaid,omitempty"`
	// Product or service provided
	Item []ExplanationOfBenefitItem `bson:"item,omitempty" json:"item,omitempty"`
	// Insurer added line items
	AddItem []ExplanationOfBenefitAddItem `bson:"addItem,omitempty" json:"addItem,omitempty"`
	// Header-level adjudication
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Adjudication totals
	Total []ExplanationOfBenefitTotal `bson:"total,omitempty" json:"total,omitempty"`
	// Payment Details
	Payment *ExplanationOfBenefitPayment `bson:"payment,omitempty" json:"payment,omitempty"`
	// Printed form identifier
	FormCode *CodeableConcept `bson:"formCode,omitempty" json:"formCode,omitempty"`
	// Printed reference or actual form
	Form *Attachment `bson:"form,omitempty" json:"form,omitempty"`
	// Note concerning adjudication
	ProcessNote []ExplanationOfBenefitProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
	// When the benefits are applicable
	BenefitPeriod *Period `bson:"benefitPeriod,omitempty" json:"benefitPeriod,omitempty"`
	// Balance by Benefit Category
	BenefitBalance []ExplanationOfBenefitBenefitBalance `bson:"benefitBalance,omitempty" json:"benefitBalance,omitempty"`
}

type ExplanationOfBenefitEvent struct {
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

type ExplanationOfBenefitInsurance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Coverage to be used for adjudication
	Focal bool `bson:"focal" json:"focal"`
	// Insurance information
	Coverage custom.Reference[Coverage] `bson:"coverage" json:"coverage"`
	// Prior authorization reference number
	PreAuthRef []string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
}

type ExplanationOfBenefitItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Item instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Applicable care team members
	CareTeamSequence []uint32 `bson:"careTeamSequence,omitempty" json:"careTeamSequence,omitempty"`
	// Applicable diagnoses
	DiagnosisSequence []uint32 `bson:"diagnosisSequence,omitempty" json:"diagnosisSequence,omitempty"`
	// Applicable procedures
	ProcedureSequence []uint32 `bson:"procedureSequence,omitempty" json:"procedureSequence,omitempty"`
	// Applicable exception and supporting information
	InformationSequence []uint32 `bson:"informationSequence,omitempty" json:"informationSequence,omitempty"`
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
	// Request or Referral for Service
	Request []custom.Reference[ExplanationOfBenefitItemRequest] `bson:"request,omitempty" json:"request,omitempty"`
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
	BodySite []ExplanationOfBenefitItemBodySite `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Adjudication results
	ReviewOutcome *ExplanationOfBenefitItemReviewOutcome `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Adjudication details
	Adjudication []ExplanationOfBenefitItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Additional items
	Detail []ExplanationOfBenefitItemDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ExplanationOfBenefitAddItemDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
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
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Additem detail level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Added items adjudication
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Insurer added line items
	SubDetail []ExplanationOfBenefitAddItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ExplanationOfBenefitTotal struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of adjudication information
	Category CodeableConcept `bson:"category" json:"category"`
	// Financial total for the category
	Amount Money `bson:"amount" json:"amount"`
}

type ExplanationOfBenefitDiagnosis struct {
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

type ExplanationOfBenefitItemReviewOutcome struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Result of the adjudication
	Decision *CodeableConcept `bson:"decision,omitempty" json:"decision,omitempty"`
	// Reason for result of the adjudication
	Reason []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	// Preauthorization reference
	PreAuthRef *string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	// Preauthorization reference effective period
	PreAuthPeriod *Period `bson:"preAuthPeriod,omitempty" json:"preAuthPeriod,omitempty"`
}

type ExplanationOfBenefitItemAdjudication struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of adjudication information
	Category CodeableConcept `bson:"category" json:"category"`
	// Explanation of adjudication outcome
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	// Monetary amount
	Amount *Money `bson:"amount,omitempty" json:"amount,omitempty"`
	// Non-monitary value
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
}

type ExplanationOfBenefitAddItemBodySite struct {
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

type ExplanationOfBenefitPayment struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Partial or complete payment
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Payment adjustment for non-claim issues
	Adjustment *Money `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	// Explanation for the variance
	AdjustmentReason *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	// Expected date of payment
	Date *custom.Date `bson:"date,omitempty" json:"date,omitempty"`
	// Payable amount after adjustment
	Amount *Money `bson:"amount,omitempty" json:"amount,omitempty"`
	// Business identifier for the payment
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
}

type ExplanationOfBenefitRelated struct {
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

type ExplanationOfBenefitProcedure struct {
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

type ExplanationOfBenefitItemDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Product or service provided
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
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Detail level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Detail level adjudication details
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Additional items
	SubDetail []ExplanationOfBenefitItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ExplanationOfBenefitItemDetailSubDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Product or service provided
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
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Subdetail level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Subdetail level adjudication details
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitAddItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Item sequence number
	ItemSequence []uint32 `bson:"itemSequence,omitempty" json:"itemSequence,omitempty"`
	// Detail sequence number
	DetailSequence []uint32 `bson:"detailSequence,omitempty" json:"detailSequence,omitempty"`
	// Subdetail sequence number
	SubDetailSequence []uint32 `bson:"subDetailSequence,omitempty" json:"subDetailSequence,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Authorized providers
	Provider []custom.Reference[ExplanationOfBenefitAddItemProvider] `bson:"provider,omitempty" json:"provider,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty"`
	// Request or Referral for Service
	Request []custom.Reference[ExplanationOfBenefitAddItemRequest] `bson:"request,omitempty" json:"request,omitempty"`
	// Service/Product billing modifiers
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
	// Anatomical location
	BodySite []ExplanationOfBenefitAddItemBodySite `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Additem level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Added items adjudication
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Insurer added line items
	Detail []ExplanationOfBenefitAddItemDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ExplanationOfBenefitBenefitBalance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Benefit classification
	Category CodeableConcept `bson:"category" json:"category"`
	// Excluded from the plan
	Excluded *bool `bson:"excluded,omitempty" json:"excluded,omitempty"`
	// Short name for the benefit
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Description of the benefit or services covered
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// In or out of network
	Network *CodeableConcept `bson:"network,omitempty" json:"network,omitempty"`
	// Individual or family
	Unit *CodeableConcept `bson:"unit,omitempty" json:"unit,omitempty"`
	// Annual or lifetime
	Term *CodeableConcept `bson:"term,omitempty" json:"term,omitempty"`
	// Benefit Summary
	Financial []ExplanationOfBenefitBenefitBalanceFinancial `bson:"financial,omitempty" json:"financial,omitempty"`
}

type ExplanationOfBenefitBenefitBalanceFinancial struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Benefit classification
	Type CodeableConcept `bson:"type" json:"type"`
	// Benefits allowed
	Allowed *uint32 `bson:"allowed,omitempty" json:"allowed,omitempty"`
	// Benefits used
	Used *uint32 `bson:"used,omitempty" json:"used,omitempty"`
}

type ExplanationOfBenefitPayee struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Category of recipient
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Recipient reference
	Party *custom.Reference[ExplanationOfBenefitPayeeParty] `bson:"party,omitempty" json:"party,omitempty"`
}

type ExplanationOfBenefitCareTeam struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Order of care team
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Practitioner or organization
	Provider custom.Reference[ExplanationOfBenefitCareTeamProvider] `bson:"provider" json:"provider"`
	// Indicator of the lead practitioner
	Responsible *bool `bson:"responsible,omitempty" json:"responsible,omitempty"`
	// Function within the team
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Practitioner or provider specialization
	Specialty *CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
}

type ExplanationOfBenefitSupportingInfo struct {
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
	Reason *Coding `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ExplanationOfBenefitAccident struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// When the incident occurred
	Date *custom.Date `bson:"date,omitempty" json:"date,omitempty"`
	// The nature of the accident
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Where the event occurred
	Location *Address `bson:"location,omitempty" json:"location,omitempty"`
}

type ExplanationOfBenefitItemBodySite struct {
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

type ExplanationOfBenefitAddItemDetailSubDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
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
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Additem subdetail level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Added items adjudication
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ExplanationOfBenefitProcessNote struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Note instance identifier
	Number *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	// Note purpose
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Note explanatory text
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// Language of the text
	Language *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}

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
	return "StructureDefinition"
}
