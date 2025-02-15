// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ClaimResponse
type ClaimResponseItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Claim item instance identifier
	ItemSequence uint32 `bson:"itemSequence" json:"itemSequence"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Adjudication results
	ReviewOutcome *ClaimResponseItemReviewOutcome `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Adjudication details
	Adjudication []ClaimResponseItemAdjudication `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Adjudication for claim details
	Detail []ClaimResponseItemDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimResponseItemDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Claim detail instance identifier
	DetailSequence uint32 `bson:"detailSequence" json:"detailSequence"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Detail level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Detail level adjudication details
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Adjudication for claim sub-details
	SubDetail []ClaimResponseItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimResponseAddItemDetail struct {
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
	// Added items detail level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Added items detail adjudication
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Insurer added line items
	SubDetail []ClaimResponseAddItemDetailSubDetail `bson:"subDetail,omitempty" json:"subDetail,omitempty"`
}

type ClaimResponsePayment struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Partial or complete payment
	Type CodeableConcept `bson:"type" json:"type"`
	// Payment adjustment for non-claim issues
	Adjustment *Money `bson:"adjustment,omitempty" json:"adjustment,omitempty"`
	// Explanation for the adjustment
	AdjustmentReason *CodeableConcept `bson:"adjustmentReason,omitempty" json:"adjustmentReason,omitempty"`
	// Expected date of payment
	Date *custom.Date `bson:"date,omitempty" json:"date,omitempty"`
	// Payable amount after adjustment
	Amount Money `bson:"amount" json:"amount"`
	// Business identifier for the payment
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
}

type ClaimResponseError struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Item sequence number
	ItemSequence *uint32 `bson:"itemSequence,omitempty" json:"itemSequence,omitempty"`
	// Detail sequence number
	DetailSequence *uint32 `bson:"detailSequence,omitempty" json:"detailSequence,omitempty"`
	// Subdetail sequence number
	SubDetailSequence *uint32 `bson:"subDetailSequence,omitempty" json:"subDetailSequence,omitempty"`
	// Error code detailing processing issues
	Code CodeableConcept `bson:"code" json:"code"`
	// FHIRPath of element(s) related to issue
	Expression []string `bson:"expression,omitempty" json:"expression,omitempty"`
}

type ClaimResponseEvent struct {
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

type ClaimResponseItemAdjudication struct {
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
	// Non-monetary value
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
}

type ClaimResponseItemDetailSubDetail struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Claim sub-detail instance identifier
	SubDetailSequence uint32 `bson:"subDetailSequence" json:"subDetailSequence"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Subdetail level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Subdetail level adjudication details
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseAddItemBodySite struct {
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

type ClaimResponseProcessNote struct {
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
	Text string `bson:"text" json:"text"`
	// Language of the text
	Language *CodeableConcept `bson:"language,omitempty" json:"language,omitempty"`
}

type ClaimResponseInsurance struct {
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
	// Insurance information
	Coverage custom.Reference[Coverage] `bson:"coverage" json:"coverage"`
	// Additional provider contract number
	BusinessArrangement *string `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
	// Adjudication results
	ClaimResponse *custom.Reference[ClaimResponse] `bson:"claimResponse,omitempty" json:"claimResponse,omitempty"`
}

type ClaimResponseItemReviewOutcome struct {
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

type ClaimResponse struct {
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
	// Business Identifier for a claim response
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// More granular claim type
	Type CodeableConcept `bson:"type" json:"type"`
	// More granular claim type
	SubType *CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty"`
	// claim | preauthorization | predetermination
	Use custom.Code `bson:"use" json:"use"`
	// The recipient of the products and services
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Response creation date
	Created custom.DateTime `bson:"created" json:"created"`
	// Party responsible for reimbursement
	Insurer *custom.Reference[Organization] `bson:"insurer,omitempty" json:"insurer,omitempty"`
	// Party responsible for the claim
	Requestor *custom.Reference[ClaimResponseRequestor] `bson:"requestor,omitempty" json:"requestor,omitempty"`
	// Id of resource triggering adjudication
	Request *custom.Reference[Claim] `bson:"request,omitempty" json:"request,omitempty"`
	// queued | complete | error | partial
	Outcome custom.Code `bson:"outcome" json:"outcome"`
	// Result of the adjudication
	Decision *CodeableConcept `bson:"decision,omitempty" json:"decision,omitempty"`
	// Disposition Message
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty"`
	// Preauthorization reference
	PreAuthRef *string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	// Preauthorization reference effective period
	PreAuthPeriod *Period `bson:"preAuthPeriod,omitempty" json:"preAuthPeriod,omitempty"`
	// Event information
	Event []ClaimResponseEvent `bson:"event,omitempty" json:"event,omitempty"`
	// Party to be paid any benefits payable
	PayeeType *CodeableConcept `bson:"payeeType,omitempty" json:"payeeType,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Package billing code
	DiagnosisRelatedGroup *CodeableConcept `bson:"diagnosisRelatedGroup,omitempty" json:"diagnosisRelatedGroup,omitempty"`
	// Adjudication for claim line items
	Item []ClaimResponseItem `bson:"item,omitempty" json:"item,omitempty"`
	// Insurer added line items
	AddItem []ClaimResponseAddItem `bson:"addItem,omitempty" json:"addItem,omitempty"`
	// Header-level adjudication
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Adjudication totals
	Total []ClaimResponseTotal `bson:"total,omitempty" json:"total,omitempty"`
	// Payment Details
	Payment *ClaimResponsePayment `bson:"payment,omitempty" json:"payment,omitempty"`
	// Funds reserved status
	FundsReserve *CodeableConcept `bson:"fundsReserve,omitempty" json:"fundsReserve,omitempty"`
	// Printed form identifier
	FormCode *CodeableConcept `bson:"formCode,omitempty" json:"formCode,omitempty"`
	// Printed reference or actual form
	Form *Attachment `bson:"form,omitempty" json:"form,omitempty"`
	// Note concerning adjudication
	ProcessNote []ClaimResponseProcessNote `bson:"processNote,omitempty" json:"processNote,omitempty"`
	// Request for additional information
	CommunicationRequest []custom.Reference[CommunicationRequest] `bson:"communicationRequest,omitempty" json:"communicationRequest,omitempty"`
	// Patient insurance information
	Insurance []ClaimResponseInsurance `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Processing errors
	Error []ClaimResponseError `bson:"error,omitempty" json:"error,omitempty"`
}

type ClaimResponseAddItem struct {
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
	SubdetailSequence []uint32 `bson:"subdetailSequence,omitempty" json:"subdetailSequence,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `bson:"traceNumber,omitempty" json:"traceNumber,omitempty"`
	// Authorized providers
	Provider []custom.Reference[ClaimResponseAddItemProvider] `bson:"provider,omitempty" json:"provider,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `bson:"revenue,omitempty" json:"revenue,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `bson:"productOrServiceEnd,omitempty" json:"productOrServiceEnd,omitempty"`
	// Request or Referral for Service
	Request []custom.Reference[ClaimResponseAddItemRequest] `bson:"request,omitempty" json:"request,omitempty"`
	// Service/Product billing modifiers
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Program the product or service is provided under
	ProgramCode []CodeableConcept `bson:"programCode,omitempty" json:"programCode,omitempty"`
	// Date or dates of service or product delivery
	Serviced *custom.Date `bson:"serviced,omitempty" json:"serviced,omitempty"`
	// Place of service or where product was supplied
	Location *CodeableConcept `bson:"location,omitempty" json:"location,omitempty"`
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
	BodySite []ClaimResponseAddItemBodySite `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `bson:"noteNumber,omitempty" json:"noteNumber,omitempty"`
	// Added items adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Added items adjudication
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
	// Insurer added line details
	Detail []ClaimResponseAddItemDetail `bson:"detail,omitempty" json:"detail,omitempty"`
}

type ClaimResponseAddItemDetailSubDetail struct {
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
	// Added items subdetail level adjudication results
	ReviewOutcome *interface{} `bson:"reviewOutcome,omitempty" json:"reviewOutcome,omitempty"`
	// Added items subdetail adjudication
	Adjudication []interface{} `bson:"adjudication,omitempty" json:"adjudication,omitempty"`
}

type ClaimResponseTotal struct {
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

type ClaimResponseRequestor interface {
	gofhirclient.Resource

	Is_ClaimResponseRequestor()
}

func (p Practitioner) Is_ClaimResponseRequestor()     {}
func (p PractitionerRole) Is_ClaimResponseRequestor() {}
func (o Organization) Is_ClaimResponseRequestor()     {}

type ClaimResponseAddItemProvider interface {
	gofhirclient.Resource

	Is_ClaimResponseAddItemProvider()
}

func (p Practitioner) Is_ClaimResponseAddItemProvider()     {}
func (p PractitionerRole) Is_ClaimResponseAddItemProvider() {}
func (o Organization) Is_ClaimResponseAddItemProvider()     {}

type ClaimResponseAddItemRequest interface {
	gofhirclient.Resource

	Is_ClaimResponseAddItemRequest()
}

func (d DeviceRequest) Is_ClaimResponseAddItemRequest()      {}
func (m MedicationRequest) Is_ClaimResponseAddItemRequest()  {}
func (n NutritionOrder) Is_ClaimResponseAddItemRequest()     {}
func (s ServiceRequest) Is_ClaimResponseAddItemRequest()     {}
func (s SupplyRequest) Is_ClaimResponseAddItemRequest()      {}
func (v VisionPrescription) Is_ClaimResponseAddItemRequest() {}

func (c ClaimResponse) ResourceType() string {
	return "StructureDefinition"
}
