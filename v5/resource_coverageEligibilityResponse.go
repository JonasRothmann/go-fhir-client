// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CoverageEligibilityResponse
type CoverageEligibilityResponseInsurance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Insurance information
	Coverage custom.Reference[Coverage] `bson:"coverage" json:"coverage"`
	// Coverage inforce indicator
	Inforce *bool `bson:"inforce,omitempty" json:"inforce,omitempty"`
	// When the benefits are applicable
	BenefitPeriod *Period `bson:"benefitPeriod,omitempty" json:"benefitPeriod,omitempty"`
	// Benefits and authorization details
	Item []CoverageEligibilityResponseInsuranceItem `bson:"item,omitempty" json:"item,omitempty"`
}

type CoverageEligibilityResponseInsuranceItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Benefit classification
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// Product or service billing modifiers
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Performing practitioner
	Provider *custom.Reference[CoverageEligibilityResponseInsuranceItemProvider] `bson:"provider,omitempty" json:"provider,omitempty"`
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
	Benefit []CoverageEligibilityResponseInsuranceItemBenefit `bson:"benefit,omitempty" json:"benefit,omitempty"`
	// Authorization required flag
	AuthorizationRequired *bool `bson:"authorizationRequired,omitempty" json:"authorizationRequired,omitempty"`
	// Type of required supporting materials
	AuthorizationSupporting []CodeableConcept `bson:"authorizationSupporting,omitempty" json:"authorizationSupporting,omitempty"`
	// Preauthorization requirements endpoint
	AuthorizationUrl *custom.URI `bson:"authorizationUrl,omitempty" json:"authorizationUrl,omitempty"`
}

type CoverageEligibilityResponseInsuranceItemBenefit struct {
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

type CoverageEligibilityResponseError struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Error code detailing processing issues
	Code CodeableConcept `bson:"code" json:"code"`
	// FHIRPath of element(s) related to issue
	Expression []string `bson:"expression,omitempty" json:"expression,omitempty"`
}

type CoverageEligibilityResponse struct {
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
	// Business Identifier for coverage eligiblity request
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// auth-requirements | benefits | discovery | validation
	Purpose []custom.Code `bson:"purpose" json:"purpose"`
	// Intended recipient of products and services
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Event information
	Event []CoverageEligibilityResponseEvent `bson:"event,omitempty" json:"event,omitempty"`
	// Estimated date or dates of service
	Serviced *custom.Date `bson:"serviced,omitempty" json:"serviced,omitempty"`
	// Response creation date
	Created custom.DateTime `bson:"created" json:"created"`
	// Party responsible for the request
	Requestor *custom.Reference[CoverageEligibilityResponseRequestor] `bson:"requestor,omitempty" json:"requestor,omitempty"`
	// Eligibility request reference
	Request custom.Reference[CoverageEligibilityRequest] `bson:"request" json:"request"`
	// queued | complete | error | partial
	Outcome custom.Code `bson:"outcome" json:"outcome"`
	// Disposition Message
	Disposition *string `bson:"disposition,omitempty" json:"disposition,omitempty"`
	// Coverage issuer
	Insurer custom.Reference[Organization] `bson:"insurer" json:"insurer"`
	// Patient insurance information
	Insurance []CoverageEligibilityResponseInsurance `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Preauthorization reference
	PreAuthRef *string `bson:"preAuthRef,omitempty" json:"preAuthRef,omitempty"`
	// Printed form identifier
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty"`
	// Processing errors
	Error []CoverageEligibilityResponseError `bson:"error,omitempty" json:"error,omitempty"`
}

type CoverageEligibilityResponseEvent struct {
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
	return "StructureDefinition"
}
