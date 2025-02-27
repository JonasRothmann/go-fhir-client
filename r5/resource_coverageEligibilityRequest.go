// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequest struct {
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
	// Desired processing priority
	Priority *CodeableConcept `json:"priority,omitempty"`
	// auth-requirements | benefits | discovery | validation
	Purpose []custom.Code `json:"purpose"`
	// Intended recipient of products and services
	Patient custom.Reference[Patient] `json:"patient"`
	// Event information
	Event []CoverageEligibilityRequestEvent `json:"event,omitempty"`
	// Estimated date or dates of service
	ServicedDate *custom.Date `json:"servicedDate,omitempty"`
	// Estimated date or dates of service
	ServicedPeriod *Period `json:"servicedPeriod,omitempty"`
	// Creation date
	Created custom.DateTime `json:"created"`
	// Author
	Enterer *custom.Reference[CoverageEligibilityRequestEnterer] `json:"enterer,omitempty"`
	// Party responsible for the request
	Provider *custom.Reference[CoverageEligibilityRequestProvider] `json:"provider,omitempty"`
	// Coverage issuer
	Insurer custom.Reference[Organization] `json:"insurer"`
	// Servicing facility
	Facility *custom.Reference[Location] `json:"facility,omitempty"`
	// Supporting information
	SupportingInfo []CoverageEligibilityRequestSupportingInfo `json:"supportingInfo,omitempty"`
	// Patient insurance information
	Insurance []CoverageEligibilityRequestInsurance `json:"insurance,omitempty"`
	// Item to be evaluated for eligibiity
	Item []CoverageEligibilityRequestItem `json:"item,omitempty"`
}

type CoverageEligibilityRequestEvent struct {
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

type CoverageEligibilityRequestSupportingInfo struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Information instance identifier
	Sequence uint32 `json:"sequence"`
	// Data to be provided
	Information custom.Reference[Resource] `json:"information"`
	// Applies to all items
	AppliesToAll *bool `json:"appliesToAll,omitempty"`
}

type CoverageEligibilityRequestInsurance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Applicable coverage
	Focal *bool `json:"focal,omitempty"`
	// Insurance information
	Coverage custom.Reference[Coverage] `json:"coverage"`
	// Additional provider contract number
	BusinessArrangement *string `json:"businessArrangement,omitempty"`
}

type CoverageEligibilityRequestItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Applicable exception or supporting information
	SupportingInfoSequence []uint32 `json:"supportingInfoSequence,omitempty"`
	// Benefit classification
	Category *CodeableConcept `json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `json:"productOrService,omitempty"`
	// Product or service billing modifiers
	Modifier []CodeableConcept `json:"modifier,omitempty"`
	// Perfoming practitioner
	Provider *custom.Reference[CoverageEligibilityRequestItemProvider] `json:"provider,omitempty"`
	// Count of products or services
	Quantity *Quantity `json:"quantity,omitempty"`
	// Fee, charge or cost per item
	UnitPrice *Money `json:"unitPrice,omitempty"`
	// Servicing facility
	Facility *custom.Reference[CoverageEligibilityRequestItemFacility] `json:"facility,omitempty"`
	// Applicable diagnosis
	Diagnosis []CoverageEligibilityRequestItemDiagnosis `json:"diagnosis,omitempty"`
	// Product or service details
	Detail []custom.Reference[Resource] `json:"detail,omitempty"`
}

type CoverageEligibilityRequestItemDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Nature of illness or problem
	DiagnosisCodeableConcept *CodeableConcept `json:"diagnosisCodeableConcept,omitempty"`
	// Nature of illness or problem
	DiagnosisReference *custom.Reference[Condition] `json:"diagnosisReference,omitempty"`
}

type OtherCoverageEligibilityRequest CoverageEligibilityRequest

type CoverageEligibilityRequestEnterer interface {
	gofhirclient.Resource

	Is_CoverageEligibilityRequestEnterer()
}

func (p Practitioner) Is_CoverageEligibilityRequestEnterer()     {}
func (p PractitionerRole) Is_CoverageEligibilityRequestEnterer() {}

type CoverageEligibilityRequestProvider interface {
	gofhirclient.Resource

	Is_CoverageEligibilityRequestProvider()
}

func (p Practitioner) Is_CoverageEligibilityRequestProvider()     {}
func (p PractitionerRole) Is_CoverageEligibilityRequestProvider() {}
func (o Organization) Is_CoverageEligibilityRequestProvider()     {}

type CoverageEligibilityRequestItemProvider interface {
	gofhirclient.Resource

	Is_CoverageEligibilityRequestItemProvider()
}

func (p Practitioner) Is_CoverageEligibilityRequestItemProvider()     {}
func (p PractitionerRole) Is_CoverageEligibilityRequestItemProvider() {}

type CoverageEligibilityRequestItemFacility interface {
	gofhirclient.Resource

	Is_CoverageEligibilityRequestItemFacility()
}

func (l Location) Is_CoverageEligibilityRequestItemFacility()     {}
func (o Organization) Is_CoverageEligibilityRequestItemFacility() {}

func (c CoverageEligibilityRequest) ResourceType() string {
	return "CoverageEligibilityRequest"
}

func (c CoverageEligibilityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverageEligibilityRequest
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCoverageEligibilityRequest: OtherCoverageEligibilityRequest(c), ResourceType: c.ResourceType()})
}

func UnmarshalCoverageEligibilityRequest(b []byte) (res CoverageEligibilityRequest, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
