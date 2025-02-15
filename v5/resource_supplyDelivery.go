// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
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
	// External identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Fulfills plan, proposal or order
	BasedOn []custom.Reference[SupplyRequest] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[SupplyDeliveryPartOf] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// in-progress | completed | abandoned | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Patient for whom the item is supplied
	Patient *custom.Reference[Patient] `bson:"patient,omitempty" json:"patient,omitempty"`
	// Category of supply event
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The item that is delivered or supplied
	SuppliedItem []SupplyDeliverySuppliedItem `bson:"suppliedItem,omitempty" json:"suppliedItem,omitempty"`
	// When event occurred
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// The item supplier
	Supplier *custom.Reference[SupplyDeliverySupplier] `bson:"supplier,omitempty" json:"supplier,omitempty"`
	// Where the delivery was sent
	Destination *custom.Reference[Location] `bson:"destination,omitempty" json:"destination,omitempty"`
	// Who received the delivery
	Receiver []custom.Reference[SupplyDeliveryReceiver] `bson:"receiver,omitempty" json:"receiver,omitempty"`
}

type SupplyDeliverySuppliedItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Amount supplied
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Medication, Substance, Device or Biologically Derived Product supplied
	Item *CodeableConcept `bson:"item,omitempty" json:"item,omitempty"`
}

type SupplyDeliveryPartOf interface {
	gofhirclient.Resource

	Is_SupplyDeliveryPartOf()
}

func (s SupplyDelivery) Is_SupplyDeliveryPartOf() {}
func (c Contract) Is_SupplyDeliveryPartOf()       {}

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
	return "StructureDefinition"
}
