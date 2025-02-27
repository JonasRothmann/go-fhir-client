// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceDispense
type DeviceDispense struct {
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
	// Business identifier for this dispensation
	Identifier []Identifier `json:"identifier,omitempty"`
	// The order or request that this dispense is fulfilling
	BasedOn []custom.Reference[DeviceDispenseBasedOn] `json:"basedOn,omitempty"`
	// The bigger event that this dispense is a part of
	PartOf []custom.Reference[Procedure] `json:"partOf,omitempty"`
	// preparation | in-progress | cancelled | on-hold | completed | entered-in-error | stopped | declined | unknown
	Status custom.Code `json:"status"`
	// Why a dispense was or was not performed
	StatusReason *custom.CodeableReference[DetectedIssue] `json:"statusReason,omitempty"`
	// Type of device dispense
	Category []CodeableConcept `json:"category,omitempty"`
	// What device was supplied
	Device custom.CodeableReference[DeviceDispenseDevice] `json:"device"`
	// Who the dispense is for
	Subject custom.Reference[DeviceDispenseSubject] `json:"subject"`
	// Who collected the device or where the medication was delivered
	Receiver *custom.Reference[DeviceDispenseReceiver] `json:"receiver,omitempty"`
	// Encounter associated with event
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Information that supports the dispensing of the device
	SupportingInformation []custom.Reference[Resource] `json:"supportingInformation,omitempty"`
	// Who performed event
	Performer []DeviceDispensePerformer `json:"performer,omitempty"`
	// Where the dispense occurred
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Trial fill, partial fill, emergency fill, etc
	Type *CodeableConcept `json:"type,omitempty"`
	// Amount dispensed
	Quantity *Quantity `json:"quantity,omitempty"`
	// When product was packaged and reviewed
	PreparedDate *custom.DateTime `json:"preparedDate,omitempty"`
	// When product was given out
	WhenHandedOver *custom.DateTime `json:"whenHandedOver,omitempty"`
	// Where the device was sent or should be sent
	Destination *custom.Reference[Location] `json:"destination,omitempty"`
	// Information about the dispense
	Note []Annotation `json:"note,omitempty"`
	// Full representation of the usage instructions
	UsageInstruction *custom.Markdown `json:"usageInstruction,omitempty"`
	// A list of relevant lifecycle events
	EventHistory []custom.Reference[Provenance] `json:"eventHistory,omitempty"`
}

type DeviceDispensePerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Who performed the dispense and what they did
	Function *CodeableConcept `json:"function,omitempty"`
	// Individual who was performing
	Actor custom.Reference[DeviceDispensePerformerActor] `json:"actor"`
}

type OtherDeviceDispense DeviceDispense

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
	return "DeviceDispense"
}

func (d DeviceDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceDispense
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDeviceDispense: OtherDeviceDispense(d), ResourceType: d.ResourceType()})
}

func UnmarshalDeviceDispense(b []byte) (res DeviceDispense, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
