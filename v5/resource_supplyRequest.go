// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SupplyRequest
type SupplyRequest struct {
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
	// Business Identifier for SupplyRequest
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// draft | active | suspended +
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// What other request is fulfilled by this supply request
	BasedOn []custom.Reference[Resource] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// The kind of supply (central, non-stock, etc.)
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// The patient for who the supply request is for
	DeliverFor *custom.Reference[Patient] `bson:"deliverFor,omitempty" json:"deliverFor,omitempty"`
	// Medication, Substance, or Device requested to be supplied
	Item custom.CodeableReference[SupplyRequestItem] `bson:"item" json:"item"`
	// The requested amount of the item indicated
	Quantity Quantity `bson:"quantity" json:"quantity"`
	// Ordered item details
	Parameter []SupplyRequestParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	// When the request should be fulfilled
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// When the request was made
	AuthoredOn *custom.DateTime `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	// Individual making the request
	Requester *custom.Reference[SupplyRequestRequester] `bson:"requester,omitempty" json:"requester,omitempty"`
	// Who is intended to fulfill the request
	Supplier []custom.Reference[SupplyRequestSupplier] `bson:"supplier,omitempty" json:"supplier,omitempty"`
	// The reason why the supply item was requested
	Reason []custom.CodeableReference[SupplyRequestReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// The origin of the supply
	DeliverFrom *custom.Reference[SupplyRequestDeliverFrom] `bson:"deliverFrom,omitempty" json:"deliverFrom,omitempty"`
	// The destination of the supply
	DeliverTo *custom.Reference[SupplyRequestDeliverTo] `bson:"deliverTo,omitempty" json:"deliverTo,omitempty"`
}

type SupplyRequestParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Item detail
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Value of detail
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

type SupplyRequestItem interface {
	gofhirclient.Resource

	Is_SupplyRequestItem()
}

func (m Medication) Is_SupplyRequestItem()                 {}
func (s Substance) Is_SupplyRequestItem()                  {}
func (d Device) Is_SupplyRequestItem()                     {}
func (d DeviceDefinition) Is_SupplyRequestItem()           {}
func (b BiologicallyDerivedProduct) Is_SupplyRequestItem() {}
func (n NutritionProduct) Is_SupplyRequestItem()           {}
func (i InventoryItem) Is_SupplyRequestItem()              {}

type SupplyRequestRequester interface {
	gofhirclient.Resource

	Is_SupplyRequestRequester()
}

func (p Practitioner) Is_SupplyRequestRequester()     {}
func (p PractitionerRole) Is_SupplyRequestRequester() {}
func (o Organization) Is_SupplyRequestRequester()     {}
func (p Patient) Is_SupplyRequestRequester()          {}
func (r RelatedPerson) Is_SupplyRequestRequester()    {}
func (d Device) Is_SupplyRequestRequester()           {}
func (c CareTeam) Is_SupplyRequestRequester()         {}

type SupplyRequestSupplier interface {
	gofhirclient.Resource

	Is_SupplyRequestSupplier()
}

func (o Organization) Is_SupplyRequestSupplier()      {}
func (h HealthcareService) Is_SupplyRequestSupplier() {}

type SupplyRequestReason interface {
	gofhirclient.Resource

	Is_SupplyRequestReason()
}

func (c Condition) Is_SupplyRequestReason()         {}
func (o Observation) Is_SupplyRequestReason()       {}
func (d DiagnosticReport) Is_SupplyRequestReason()  {}
func (d DocumentReference) Is_SupplyRequestReason() {}

type SupplyRequestDeliverFrom interface {
	gofhirclient.Resource

	Is_SupplyRequestDeliverFrom()
}

func (o Organization) Is_SupplyRequestDeliverFrom() {}
func (l Location) Is_SupplyRequestDeliverFrom()     {}

type SupplyRequestDeliverTo interface {
	gofhirclient.Resource

	Is_SupplyRequestDeliverTo()
}

func (o Organization) Is_SupplyRequestDeliverTo()  {}
func (l Location) Is_SupplyRequestDeliverTo()      {}
func (p Patient) Is_SupplyRequestDeliverTo()       {}
func (r RelatedPerson) Is_SupplyRequestDeliverTo() {}

func (s SupplyRequest) ResourceType() string {
	return "StructureDefinition"
}
