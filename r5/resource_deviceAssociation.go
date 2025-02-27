// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceAssociation
type DeviceAssociation struct {
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
	// Instance identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Reference to the devices associated with the patient or group
	Device custom.Reference[Device] `json:"device"`
	// Describes the relationship between the device and subject
	Category []CodeableConcept `json:"category,omitempty"`
	// implanted | explanted | attached | entered-in-error | unknown
	Status CodeableConcept `json:"status"`
	// The reasons given for the current association status
	StatusReason []CodeableConcept `json:"statusReason,omitempty"`
	// The individual, group of individuals or device that the device is on or associated with
	Subject *custom.Reference[DeviceAssociationSubject] `json:"subject,omitempty"`
	// Current anatomical location of the device in/on subject
	BodyStructure *custom.Reference[BodyStructure] `json:"bodyStructure,omitempty"`
	// Begin and end dates and times for the device association
	Period *Period `json:"period,omitempty"`
	// The details about the device when it is in use to describe its operation
	Operation []DeviceAssociationOperation `json:"operation,omitempty"`
}

type DeviceAssociationOperation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Device operational condition
	Status CodeableConcept `json:"status"`
	// The individual performing the action enabled by the device
	Operator []custom.Reference[DeviceAssociationOperationOperator] `json:"operator,omitempty"`
	// Begin and end dates and times for the device's operation
	Period *Period `json:"period,omitempty"`
}

type OtherDeviceAssociation DeviceAssociation

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
	return "DeviceAssociation"
}

func (d DeviceAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceAssociation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDeviceAssociation: OtherDeviceAssociation(d), ResourceType: d.ResourceType()})
}

func UnmarshalDeviceAssociation(b []byte) (res DeviceAssociation, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
