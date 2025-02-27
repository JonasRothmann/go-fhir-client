// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponse struct {
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
	// Business Identifier for coverage eligiblity request
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `json:"status"`
	// auth-requirements | benefits | discovery | validation
	Purpose []custom.Code `json:"purpose"`
	// Intended recipient of products and services
	Patient custom.Reference[Patient] `json:"patient"`
	// Event information
	Event []CoverageEligibilityResponseEvent `json:"event,omitempty"`
	// Estimated date or dates of service
	ServicedDate *custom.Date `json:"servicedDate,omitempty"`
	// Estimated date or dates of service
	ServicedPeriod *Period `json:"servicedPeriod,omitempty"`
	// Response creation date
	Created custom.DateTime `json:"created"`
	// Party responsible for the request
	Requestor *custom.Reference[CoverageEligibilityResponseRequestor] `json:"requestor,omitempty"`
	// Eligibility request reference
	Request custom.Reference[CoverageEligibilityRequest] `json:"request"`
	// queued | complete | error | partial
	Outcome custom.Code `json:"outcome"`
	// Disposition Message
	Disposition *string `json:"disposition,omitempty"`
	// Coverage issuer
	Insurer custom.Reference[Organization] `json:"insurer"`
	// Patient insurance information
	Insurance []CoverageEligibilityResponseInsurance `json:"insurance,omitempty"`
	// Preauthorization reference
	PreAuthRef *string `json:"preAuthRef,omitempty"`
	// Printed form identifier
	Form *CodeableConcept `json:"form,omitempty"`
	// Processing errors
	Error []CoverageEligibilityResponseError `json:"error,omitempty"`
}

type CoverageEligibilityResponseEvent struct {
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

type CoverageEligibilityResponseInsurance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Insurance information
	Coverage custom.Reference[Coverage] `json:"coverage"`
	// Coverage inforce indicator
	Inforce *bool `json:"inforce,omitempty"`
	// When the benefits are applicable
	BenefitPeriod *Period `json:"benefitPeriod,omitempty"`
	// Benefits and authorization details
	Item []CoverageEligibilityResponseInsuranceItem `json:"item,omitempty"`
}

type CoverageEligibilityResponseInsuranceItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Benefit classification
	Category *CodeableConcept `json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// Product or service billing modifiers
	Modifier []CodeableConcept `json:"modifier,omitempty"`
	// Performing practitioner
	Provider *custom.Reference[CoverageEligibilityResponseInsuranceItemProvider] `json:"provider,omitempty"`
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
	Benefit []CoverageEligibilityResponseInsuranceItemBenefit `json:"benefit,omitempty"`
	// Authorization required flag
	AuthorizationRequired *bool `json:"authorizationRequired,omitempty"`
	// Type of required supporting materials
	AuthorizationSupporting []CodeableConcept `json:"authorizationSupporting,omitempty"`
	// Preauthorization requirements endpoint
	AuthorizationUrl *custom.URI `json:"authorizationUrl,omitempty"`
}

type CoverageEligibilityResponseInsuranceItemBenefit struct {
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
	UsedString *string `json:"usedString,omitempty"`
	// Benefits used
	UsedMoney *Money `json:"usedMoney,omitempty"`
}

type CoverageEligibilityResponseError struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Error code detailing processing issues
	Code CodeableConcept `json:"code"`
	// FHIRPath of element(s) related to issue
	Expression []string `json:"expression,omitempty"`
}

type OtherCoverageEligibilityResponse CoverageEligibilityResponse

type CoverageEligibilityResponseRequestor interface {
	gofhirclient.Resource

	Is_CoverageEligibilityResponseRequestor()
}

func (p Practitioner) Is_CoverageEligibilityResponseRequestor()     {}
func (p PractitionerRole) Is_CoverageEligibilityResponseRequestor() {}
func (o Organization) Is_CoverageEligibilityResponseRequestor()     {}

type CoverageEligibilityResponseInsuranceItemProvider interface {
	gofhirclient.Resource

	Is_CoverageEligibilityResponseInsuranceItemProvider()
}

func (p Practitioner) Is_CoverageEligibilityResponseInsuranceItemProvider()     {}
func (p PractitionerRole) Is_CoverageEligibilityResponseInsuranceItemProvider() {}

func (c CoverageEligibilityResponse) ResourceType() string {
	return "CoverageEligibilityResponse"
}

func (c CoverageEligibilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityResponse
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCoverageEligibilityResponse: OtherCoverageEligibilityResponse(c), ResourceType: c.ResourceType()})
}

func UnmarshalCoverageEligibilityResponse(b []byte) (res CoverageEligibilityResponse, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
