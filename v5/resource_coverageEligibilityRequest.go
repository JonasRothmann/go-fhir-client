// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CoverageEligibilityRequest
type CoverageEligibilityRequest struct {
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
	// Desired processing priority
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty"`
	// auth-requirements | benefits | discovery | validation
	Purpose []custom.Code `bson:"purpose" json:"purpose"`
	// Intended recipient of products and services
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Event information
	Event []CoverageEligibilityRequestEvent `bson:"event,omitempty" json:"event,omitempty"`
	// Estimated date or dates of service
	Serviced *custom.Date `bson:"serviced,omitempty" json:"serviced,omitempty"`
	// Creation date
	Created custom.DateTime `bson:"created" json:"created"`
	// Author
	Enterer *custom.Reference[CoverageEligibilityRequestEnterer] `bson:"enterer,omitempty" json:"enterer,omitempty"`
	// Party responsible for the request
	Provider *custom.Reference[CoverageEligibilityRequestProvider] `bson:"provider,omitempty" json:"provider,omitempty"`
	// Coverage issuer
	Insurer custom.Reference[Organization] `bson:"insurer" json:"insurer"`
	// Servicing facility
	Facility *custom.Reference[Location] `bson:"facility,omitempty" json:"facility,omitempty"`
	// Supporting information
	SupportingInfo []CoverageEligibilityRequestSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Patient insurance information
	Insurance []CoverageEligibilityRequestInsurance `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Item to be evaluated for eligibiity
	Item []CoverageEligibilityRequestItem `bson:"item,omitempty" json:"item,omitempty"`
}

type CoverageEligibilityRequestEvent struct {
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

type CoverageEligibilityRequestSupportingInfo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Information instance identifier
	Sequence uint32 `bson:"sequence" json:"sequence"`
	// Data to be provided
	Information custom.Reference[Resource] `bson:"information" json:"information"`
	// Applies to all items
	AppliesToAll *bool `bson:"appliesToAll,omitempty" json:"appliesToAll,omitempty"`
}

type CoverageEligibilityRequestInsurance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Applicable coverage
	Focal *bool `bson:"focal,omitempty" json:"focal,omitempty"`
	// Insurance information
	Coverage custom.Reference[Coverage] `bson:"coverage" json:"coverage"`
	// Additional provider contract number
	BusinessArrangement *string `bson:"businessArrangement,omitempty" json:"businessArrangement,omitempty"`
}

type CoverageEligibilityRequestItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Applicable exception or supporting information
	SupportingInfoSequence []uint32 `bson:"supportingInfoSequence,omitempty" json:"supportingInfoSequence,omitempty"`
	// Benefit classification
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Billing, service, product, or drug code
	ProductOrService *CodeableConcept `bson:"productOrService,omitempty" json:"productOrService,omitempty"`
	// Product or service billing modifiers
	Modifier []CodeableConcept `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Perfoming practitioner
	Provider *custom.Reference[CoverageEligibilityRequestItemProvider] `bson:"provider,omitempty" json:"provider,omitempty"`
	// Count of products or services
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Fee, charge or cost per item
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	// Servicing facility
	Facility *custom.Reference[CoverageEligibilityRequestItemFacility] `bson:"facility,omitempty" json:"facility,omitempty"`
	// Applicable diagnosis
	Diagnosis []CoverageEligibilityRequestItemDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	// Product or service details
	Detail []custom.Reference[Resource] `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CoverageEligibilityRequestItemDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Nature of illness or problem
	Diagnosis *CodeableConcept `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
}

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
	return "StructureDefinition"
}
