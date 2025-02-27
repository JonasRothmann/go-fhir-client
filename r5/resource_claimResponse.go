// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ClaimResponse
type ClaimResponseError struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Item sequence number
	ItemSequence *uint32 `json:"itemSequence,omitempty"`
	// Detail sequence number
	DetailSequence *uint32 `json:"detailSequence,omitempty"`
	// Subdetail sequence number
	SubDetailSequence *uint32 `json:"subDetailSequence,omitempty"`
	// Error code detailing processing issues
	Code CodeableConcept `json:"code"`
	// FHIRPath of element(s) related to issue
	Expression []string `json:"expression,omitempty"`
}

type ClaimResponse struct {
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
	// Business Identifier for a claim response
	Identifier []Identifier `json:"identifier,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `json:"status"`
	// More granular claim type
	Type CodeableConcept `json:"type"`
	// More granular claim type
	SubType *CodeableConcept `json:"subType,omitempty"`
	// claim | preauthorization | predetermination
	Use custom.Code `json:"use"`
	// The recipient of the products and services
	Patient custom.Reference[Patient] `json:"patient"`
	// Response creation date
	Created custom.DateTime `json:"created"`
	// Party responsible for reimbursement
	Insurer *custom.Reference[Organization] `json:"insurer,omitempty"`
	// Party responsible for the claim
	Requestor *custom.Reference[ClaimResponseRequestor] `json:"requestor,omitempty"`
	// Id of resource triggering adjudication
	Request *custom.Reference[Claim] `json:"request,omitempty"`
	// queued | complete | error | partial
	Outcome custom.Code `json:"outcome"`
	// Result of the adjudication
	Decision *CodeableConcept `json:"decision,omitempty"`
	// Disposition Message
	Disposition *string `json:"disposition,omitempty"`
	// Preauthorization reference
	PreAuthRef *string `json:"preAuthRef,omitempty"`
	// Preauthorization reference effective period
	PreAuthPeriod *Period `json:"preAuthPeriod,omitempty"`
	// Event information
	Event []ClaimResponseEvent `json:"event,omitempty"`
	// Party to be paid any benefits payable
	PayeeType *CodeableConcept `json:"payeeType,omitempty"`
	// Encounters associated with the listed treatments
	Encounter []custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Package billing code
	DiagnosisRelatedGroup *CodeableConcept `json:"diagnosisRelatedGroup,omitempty"`
	// Adjudication for claim line items
	Item []ClaimResponseItem `json:"item,omitempty"`
	// Insurer added line items
	AddItem []ClaimResponseAddItem `json:"addItem,omitempty"`
	// Header-level adjudication
	Adjudication []interface{} `json:"adjudication,omitempty"`
	// Adjudication totals
	Total []ClaimResponseTotal `json:"total,omitempty"`
	// Payment Details
	Payment *ClaimResponsePayment `json:"payment,omitempty"`
	// Funds reserved status
	FundsReserve *CodeableConcept `json:"fundsReserve,omitempty"`
	// Printed form identifier
	FormCode *CodeableConcept `json:"formCode,omitempty"`
	// Printed reference or actual form
	Form *Attachment `json:"form,omitempty"`
	// Note concerning adjudication
	ProcessNote []ClaimResponseProcessNote `json:"processNote,omitempty"`
	// Request for additional information
	CommunicationRequest []custom.Reference[CommunicationRequest] `json:"communicationRequest,omitempty"`
	// Patient insurance information
	Insurance []ClaimResponseInsurance `json:"insurance,omitempty"`
	// Processing errors
	Error []ClaimResponseError `json:"error,omitempty"`
}

type ClaimResponseItemAdjudication struct {
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
	// Non-monetary value
	Quantity *Quantity `json:"quantity,omitempty"`
}

type ClaimResponseAddItemBodySite struct {
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

type ClaimResponsePayment struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Partial or complete payment
	Type CodeableConcept `json:"type"`
	// Payment adjustment for non-claim issues
	Adjustment *Money `json:"adjustment,omitempty"`
	// Explanation for the adjustment
	AdjustmentReason *CodeableConcept `json:"adjustmentReason,omitempty"`
	// Expected date of payment
	Date *custom.Date `json:"date,omitempty"`
	// Payable amount after adjustment
	Amount Money `json:"amount"`
	// Business identifier for the payment
	Identifier *Identifier `json:"identifier,omitempty"`
}

type ClaimResponseItemDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Claim detail instance identifier
	DetailSequence uint32 `json:"detailSequence"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Detail level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Detail level adjudication details
	Adjudication []interface{} `json:"adjudication,omitempty"`
	// Adjudication for claim sub-details
	SubDetail []ClaimResponseItemDetailSubDetail `json:"subDetail,omitempty"`
}

type ClaimResponseAddItem struct {
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
	SubdetailSequence []uint32 `json:"subdetailSequence,omitempty"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Authorized providers
	Provider []custom.Reference[ClaimResponseAddItemProvider] `json:"provider,omitempty"`
	// Revenue or cost center code
	Revenue *CodeableConcept `json:"revenue,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// End of a range of codes
	ProductOrServiceEnd *CodeableConcept `json:"productOrServiceEnd,omitempty"`
	// Request or Referral for Service
	Request []custom.Reference[ClaimResponseAddItemRequest] `json:"request,omitempty"`
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
	BodySite []ClaimResponseAddItemBodySite `json:"bodySite,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Added items adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Added items adjudication
	Adjudication []interface{} `json:"adjudication,omitempty"`
	// Insurer added line details
	Detail []ClaimResponseAddItemDetail `json:"detail,omitempty"`
}

type ClaimResponseEvent struct {
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

type ClaimResponseItemReviewOutcome struct {
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

type ClaimResponseItemDetailSubDetail struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Claim sub-detail instance identifier
	SubDetailSequence uint32 `json:"subDetailSequence"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Subdetail level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Subdetail level adjudication details
	Adjudication []interface{} `json:"adjudication,omitempty"`
}

type ClaimResponseTotal struct {
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

type ClaimResponseInsurance struct {
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
	// Insurance information
	Coverage custom.Reference[Coverage] `json:"coverage"`
	// Additional provider contract number
	BusinessArrangement *string `json:"businessArrangement,omitempty"`
	// Adjudication results
	ClaimResponse *custom.Reference[ClaimResponse] `json:"claimResponse,omitempty"`
}

type ClaimResponseItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Claim item instance identifier
	ItemSequence uint32 `json:"itemSequence"`
	// Number for tracking
	TraceNumber []Identifier `json:"traceNumber,omitempty"`
	// Applicable note numbers
	NoteNumber []uint32 `json:"noteNumber,omitempty"`
	// Adjudication results
	ReviewOutcome *ClaimResponseItemReviewOutcome `json:"reviewOutcome,omitempty"`
	// Adjudication details
	Adjudication []ClaimResponseItemAdjudication `json:"adjudication,omitempty"`
	// Adjudication for claim details
	Detail []ClaimResponseItemDetail `json:"detail,omitempty"`
}

type ClaimResponseAddItemDetail struct {
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
	// Added items detail level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Added items detail adjudication
	Adjudication []interface{} `json:"adjudication,omitempty"`
	// Insurer added line items
	SubDetail []ClaimResponseAddItemDetailSubDetail `json:"subDetail,omitempty"`
}

type ClaimResponseAddItemDetailSubDetail struct {
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
	// Added items subdetail level adjudication results
	ReviewOutcome *interface{} `json:"reviewOutcome,omitempty"`
	// Added items subdetail adjudication
	Adjudication []interface{} `json:"adjudication,omitempty"`
}

type ClaimResponseProcessNote struct {
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
	Text string `json:"text"`
	// Language of the text
	Language *CodeableConcept `json:"language,omitempty"`
}

type OtherClaimResponse ClaimResponse

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
	return "ClaimResponse"
}

func (c ClaimResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClaimResponse
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherClaimResponse: OtherClaimResponse(c), ResourceType: c.ResourceType()})
}

func UnmarshalClaimResponse(b []byte) (res ClaimResponse, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
