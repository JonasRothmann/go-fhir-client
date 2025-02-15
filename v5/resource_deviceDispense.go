// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceDispense
type DeviceDispense struct {
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
	// Business identifier for this dispensation
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The order or request that this dispense is fulfilling
	BasedOn []custom.Reference[DeviceDispenseBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// The bigger event that this dispense is a part of
	PartOf []custom.Reference[Procedure] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// preparation | in-progress | cancelled | on-hold | completed | entered-in-error | stopped | declined | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Why a dispense was or was not performed
	StatusReason *custom.CodeableReference[DetectedIssue] `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// Type of device dispense
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// What device was supplied
	Device custom.CodeableReference[DeviceDispenseDevice] `bson:"device" json:"device"`
	// Who the dispense is for
	Subject custom.Reference[DeviceDispenseSubject] `bson:"subject" json:"subject"`
	// Who collected the device or where the medication was delivered
	Receiver *custom.Reference[DeviceDispenseReceiver] `bson:"receiver,omitempty" json:"receiver,omitempty"`
	// Encounter associated with event
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Information that supports the dispensing of the device
	SupportingInformation []custom.Reference[Resource] `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	// Who performed event
	Performer []DeviceDispensePerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Where the dispense occurred
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Trial fill, partial fill, emergency fill, etc
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Amount dispensed
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// When product was packaged and reviewed
	PreparedDate *custom.DateTime `bson:"preparedDate,omitempty" json:"preparedDate,omitempty"`
	// When product was given out
	WhenHandedOver *custom.DateTime `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	// Where the device was sent or should be sent
	Destination *custom.Reference[Location] `bson:"destination,omitempty" json:"destination,omitempty"`
	// Information about the dispense
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Full representation of the usage instructions
	UsageInstruction *custom.Markdown `bson:"usageInstruction,omitempty" json:"usageInstruction,omitempty"`
	// A list of relevant lifecycle events
	EventHistory []custom.Reference[Provenance] `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

type DeviceDispensePerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Who performed the dispense and what they did
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Individual who was performing
	Actor custom.Reference[DeviceDispensePerformerActor] `bson:"actor" json:"actor"`
}

type DeviceDispenseBasedOn interface {
	gofhirclient.Resource

	Is_DeviceDispenseBasedOn()
}

func (c CarePlan) Is_DeviceDispenseBasedOn()      {}
func (d DeviceRequest) Is_DeviceDispenseBasedOn() {}

type DeviceDispenseDevice interface {
	gofhirclient.Resource

	Is_DeviceDispenseDevice()
}

func (d Device) Is_DeviceDispenseDevice()           {}
func (d DeviceDefinition) Is_DeviceDispenseDevice() {}

type DeviceDispenseSubject interface {
	gofhirclient.Resource

	Is_DeviceDispenseSubject()
}

func (p Patient) Is_DeviceDispenseSubject()      {}
func (p Practitioner) Is_DeviceDispenseSubject() {}

type DeviceDispenseReceiver interface {
	gofhirclient.Resource

	Is_DeviceDispenseReceiver()
}

func (p Patient) Is_DeviceDispenseReceiver()          {}
func (p Practitioner) Is_DeviceDispenseReceiver()     {}
func (r RelatedPerson) Is_DeviceDispenseReceiver()    {}
func (l Location) Is_DeviceDispenseReceiver()         {}
func (p PractitionerRole) Is_DeviceDispenseReceiver() {}

type DeviceDispensePerformerActor interface {
	gofhirclient.Resource

	Is_DeviceDispensePerformerActor()
}

func (p Practitioner) Is_DeviceDispensePerformerActor()     {}
func (p PractitionerRole) Is_DeviceDispensePerformerActor() {}
func (o Organization) Is_DeviceDispensePerformerActor()     {}
func (p Patient) Is_DeviceDispensePerformerActor()          {}
func (d Device) Is_DeviceDispensePerformerActor()           {}
func (r RelatedPerson) Is_DeviceDispensePerformerActor()    {}
func (c CareTeam) Is_DeviceDispensePerformerActor()         {}

func (d DeviceDispense) ResourceType() string {
	return "StructureDefinition"
}
