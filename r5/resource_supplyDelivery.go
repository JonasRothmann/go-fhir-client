// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
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
	// External identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Fulfills plan, proposal or order
	BasedOn []custom.Reference[SupplyRequest] `json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[SupplyDeliveryPartOf] `json:"partOf,omitempty"`
	// in-progress | completed | abandoned | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// Patient for whom the item is supplied
	Patient *custom.Reference[Patient] `json:"patient,omitempty"`
	// Category of supply event
	Type *CodeableConcept `json:"type,omitempty"`
	// The item that is delivered or supplied
	SuppliedItem []SupplyDeliverySuppliedItem `json:"suppliedItem,omitempty"`
	// When event occurred
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// When event occurred
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When event occurred
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	// The item supplier
	Supplier *custom.Reference[SupplyDeliverySupplier] `json:"supplier,omitempty"`
	// Where the delivery was sent
	Destination *custom.Reference[Location] `json:"destination,omitempty"`
	// Who received the delivery
	Receiver []custom.Reference[SupplyDeliveryReceiver] `json:"receiver,omitempty"`
}

type SupplyDeliverySuppliedItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Amount supplied
	Quantity *Quantity `json:"quantity,omitempty"`
	// Medication, Substance, Device or Biologically Derived Product supplied
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	// Medication, Substance, Device or Biologically Derived Product supplied
	ItemReference *custom.Reference[SupplyDeliverySuppliedItemItemReference] `json:"itemReference,omitempty"`
}

type OtherSupplyDelivery SupplyDelivery

type SupplyDeliveryPartOf interface {
	gofhirclient.Resource

	Is_SupplyDeliveryPartOf()
}

func (s SupplyDelivery) Is_SupplyDeliveryPartOf() {}
func (c Contract) Is_SupplyDeliveryPartOf()       {}

type SupplyDeliverySuppliedItemItemReference interface {
	gofhirclient.Resource

	Is_SupplyDeliverySuppliedItemItemReference()
}

func (m Medication) Is_SupplyDeliverySuppliedItemItemReference()                 {}
func (s Substance) Is_SupplyDeliverySuppliedItemItemReference()                  {}
func (d Device) Is_SupplyDeliverySuppliedItemItemReference()                     {}
func (b BiologicallyDerivedProduct) Is_SupplyDeliverySuppliedItemItemReference() {}
func (n NutritionProduct) Is_SupplyDeliverySuppliedItemItemReference()           {}
func (i InventoryItem) Is_SupplyDeliverySuppliedItemItemReference()              {}

type SupplyDeliverySupplier interface {
	gofhirclient.Resource

	Is_SupplyDeliverySupplier()
}

func (p Practitioner) Is_SupplyDeliverySupplier()     {}
func (p PractitionerRole) Is_SupplyDeliverySupplier() {}
func (o Organization) Is_SupplyDeliverySupplier()     {}

type SupplyDeliveryReceiver interface {
	gofhirclient.Resource

	Is_SupplyDeliveryReceiver()
}

func (p Practitioner) Is_SupplyDeliveryReceiver()     {}
func (p PractitionerRole) Is_SupplyDeliveryReceiver() {}
func (o Organization) Is_SupplyDeliveryReceiver()     {}

func (s SupplyDelivery) ResourceType() string {
	return "SupplyDelivery"
}

func (s SupplyDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyDelivery
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSupplyDelivery: OtherSupplyDelivery(s), ResourceType: s.ResourceType()})
}

func UnmarshalSupplyDelivery(b []byte) (res SupplyDelivery, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
