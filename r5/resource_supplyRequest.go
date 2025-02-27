// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SupplyRequest
type SupplyRequest struct {
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
	// Business Identifier for SupplyRequest
	Identifier []Identifier `json:"identifier,omitempty"`
	// draft | active | suspended +
	Status *custom.Code `json:"status,omitempty"`
	// What other request is fulfilled by this supply request
	BasedOn []custom.Reference[Resource] `json:"basedOn,omitempty"`
	// The kind of supply (central, non-stock, etc.)
	Category *CodeableConcept `json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// The patient for who the supply request is for
	DeliverFor *custom.Reference[Patient] `json:"deliverFor,omitempty"`
	// Medication, Substance, or Device requested to be supplied
	Item custom.CodeableReference[SupplyRequestItem] `json:"item"`
	// The requested amount of the item indicated
	Quantity Quantity `json:"quantity"`
	// Ordered item details
	Parameter []SupplyRequestParameter `json:"parameter,omitempty"`
	// When the request should be fulfilled
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// When the request should be fulfilled
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When the request should be fulfilled
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	// When the request was made
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Individual making the request
	Requester *custom.Reference[SupplyRequestRequester] `json:"requester,omitempty"`
	// Who is intended to fulfill the request
	Supplier []custom.Reference[SupplyRequestSupplier] `json:"supplier,omitempty"`
	// The reason why the supply item was requested
	Reason []custom.CodeableReference[SupplyRequestReason] `json:"reason,omitempty"`
	// The origin of the supply
	DeliverFrom *custom.Reference[SupplyRequestDeliverFrom] `json:"deliverFrom,omitempty"`
	// The destination of the supply
	DeliverTo *custom.Reference[SupplyRequestDeliverTo] `json:"deliverTo,omitempty"`
}

type SupplyRequestParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Item detail
	Code *CodeableConcept `json:"code,omitempty"`
	// Value of detail
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Value of detail
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Value of detail
	ValueRange *Range `json:"valueRange,omitempty"`
	// Value of detail
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
}

type OtherSupplyRequest SupplyRequest

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
	return "SupplyRequest"
}

func (s SupplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyRequest
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSupplyRequest: OtherSupplyRequest(s), ResourceType: s.ResourceType()})
}

func UnmarshalSupplyRequest(b []byte) (res SupplyRequest, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
