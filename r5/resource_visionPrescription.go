// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/VisionPrescription
type VisionPrescription struct {
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
	// Business Identifier for vision prescription
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `json:"status"`
	// Response creation date
	Created custom.DateTime `json:"created"`
	// Who prescription is for
	Patient custom.Reference[Patient] `json:"patient"`
	// Created during encounter / admission / stay
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When prescription was authorized
	DateWritten custom.DateTime `json:"dateWritten"`
	// Who authorized the vision prescription
	Prescriber custom.Reference[VisionPrescriptionPrescriber] `json:"prescriber"`
	// Vision lens authorization
	LensSpecification []VisionPrescriptionLensSpecification `json:"lensSpecification"`
}

type VisionPrescriptionLensSpecification struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Product to be supplied
	Product CodeableConcept `json:"product"`
	// right | left
	Eye custom.Code `json:"eye"`
	// Power of the lens
	Sphere *json.Number `json:"sphere,omitempty"`
	// Lens power for astigmatism
	Cylinder *json.Number `json:"cylinder,omitempty"`
	// Lens meridian which contain no power for astigmatism
	Axis *int32 `json:"axis,omitempty"`
	// Eye alignment compensation
	Prism []VisionPrescriptionLensSpecificationPrism `json:"prism,omitempty"`
	// Added power for multifocal levels
	Add *json.Number `json:"add,omitempty"`
	// Contact lens power
	Power *json.Number `json:"power,omitempty"`
	// Contact lens back curvature
	BackCurve *json.Number `json:"backCurve,omitempty"`
	// Contact lens diameter
	Diameter *json.Number `json:"diameter,omitempty"`
	// Lens wear duration
	Duration *Quantity `json:"duration,omitempty"`
	// Color required
	Color *string `json:"color,omitempty"`
	// Brand required
	Brand *string `json:"brand,omitempty"`
	// Notes for coatings
	Note []Annotation `json:"note,omitempty"`
}

type VisionPrescriptionLensSpecificationPrism struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Amount of adjustment
	Amount json.Number `json:"amount"`
	// up | down | in | out
	Base custom.Code `json:"base"`
}

type OtherVisionPrescription VisionPrescription

type VisionPrescriptionPrescriber interface {
	gofhirclient.Resource

	Is_VisionPrescriptionPrescriber()
}

func (p Practitioner) Is_VisionPrescriptionPrescriber()     {}
func (p PractitionerRole) Is_VisionPrescriptionPrescriber() {}

func (v VisionPrescription) ResourceType() string {
	return "VisionPrescription"
}

func (v VisionPrescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVisionPrescription
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherVisionPrescription: OtherVisionPrescription(v), ResourceType: v.ResourceType()})
}

func UnmarshalVisionPrescription(b []byte) (res VisionPrescription, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
