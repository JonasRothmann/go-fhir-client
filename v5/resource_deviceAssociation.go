// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceAssociation
type DeviceAssociation struct {
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
	// Instance identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Reference to the devices associated with the patient or group
	Device custom.Reference[Device] `bson:"device" json:"device"`
	// Describes the relationship between the device and subject
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// implanted | explanted | attached | entered-in-error | unknown
	Status CodeableConcept `bson:"status" json:"status"`
	// The reasons given for the current association status
	StatusReason []CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// The individual, group of individuals or device that the device is on or associated with
	Subject *custom.Reference[DeviceAssociationSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Current anatomical location of the device in/on subject
	BodyStructure *custom.Reference[BodyStructure] `bson:"bodyStructure,omitempty" json:"bodyStructure,omitempty"`
	// Begin and end dates and times for the device association
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// The details about the device when it is in use to describe its operation
	Operation []DeviceAssociationOperation `bson:"operation,omitempty" json:"operation,omitempty"`
}

type DeviceAssociationOperation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Device operational condition
	Status CodeableConcept `bson:"status" json:"status"`
	// The individual performing the action enabled by the device
	Operator []custom.Reference[DeviceAssociationOperationOperator] `bson:"operator,omitempty" json:"operator,omitempty"`
	// Begin and end dates and times for the device's operation
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type DeviceAssociationSubject interface {
	gofhirclient.Resource

	Is_DeviceAssociationSubject()
}

func (p Patient) Is_DeviceAssociationSubject()       {}
func (g Group) Is_DeviceAssociationSubject()         {}
func (p Practitioner) Is_DeviceAssociationSubject()  {}
func (r RelatedPerson) Is_DeviceAssociationSubject() {}
func (d Device) Is_DeviceAssociationSubject()        {}

type DeviceAssociationOperationOperator interface {
	gofhirclient.Resource

	Is_DeviceAssociationOperationOperator()
}

func (p Patient) Is_DeviceAssociationOperationOperator()       {}
func (p Practitioner) Is_DeviceAssociationOperationOperator()  {}
func (r RelatedPerson) Is_DeviceAssociationOperationOperator() {}

func (d DeviceAssociation) ResourceType() string {
	return "StructureDefinition"
}
