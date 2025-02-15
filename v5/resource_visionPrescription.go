// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/VisionPrescription
type VisionPrescription struct {
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
	// Business Identifier for vision prescription
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | cancelled | draft | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Response creation date
	Created custom.DateTime `bson:"created" json:"created"`
	// Who prescription is for
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Created during encounter / admission / stay
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When prescription was authorized
	DateWritten custom.DateTime `bson:"dateWritten" json:"dateWritten"`
	// Who authorized the vision prescription
	Prescriber custom.Reference[VisionPrescriptionPrescriber] `bson:"prescriber" json:"prescriber"`
	// Vision lens authorization
	LensSpecification []VisionPrescriptionLensSpecification `bson:"lensSpecification" json:"lensSpecification"`
}

type VisionPrescriptionLensSpecification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Product to be supplied
	Product CodeableConcept `bson:"product" json:"product"`
	// right | left
	Eye custom.Code `bson:"eye" json:"eye"`
	// Power of the lens
	Sphere *json.Number `bson:"sphere,omitempty" json:"sphere,omitempty"`
	// Lens power for astigmatism
	Cylinder *json.Number `bson:"cylinder,omitempty" json:"cylinder,omitempty"`
	// Lens meridian which contain no power for astigmatism
	Axis *int32 `bson:"axis,omitempty" json:"axis,omitempty"`
	// Eye alignment compensation
	Prism []VisionPrescriptionLensSpecificationPrism `bson:"prism,omitempty" json:"prism,omitempty"`
	// Added power for multifocal levels
	Add *json.Number `bson:"add,omitempty" json:"add,omitempty"`
	// Contact lens power
	Power *json.Number `bson:"power,omitempty" json:"power,omitempty"`
	// Contact lens back curvature
	BackCurve *json.Number `bson:"backCurve,omitempty" json:"backCurve,omitempty"`
	// Contact lens diameter
	Diameter *json.Number `bson:"diameter,omitempty" json:"diameter,omitempty"`
	// Lens wear duration
	Duration *Quantity `bson:"duration,omitempty" json:"duration,omitempty"`
	// Color required
	Color *string `bson:"color,omitempty" json:"color,omitempty"`
	// Brand required
	Brand *string `bson:"brand,omitempty" json:"brand,omitempty"`
	// Notes for coatings
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type VisionPrescriptionLensSpecificationPrism struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Amount of adjustment
	Amount json.Number `bson:"amount" json:"amount"`
	// up | down | in | out
	Base custom.Code `bson:"base" json:"base"`
}

type VisionPrescriptionPrescriber interface {
	gofhirclient.Resource

	Is_VisionPrescriptionPrescriber()
}

func (p Practitioner) Is_VisionPrescriptionPrescriber()     {}
func (p PractitionerRole) Is_VisionPrescriptionPrescriber() {}

func (v VisionPrescription) ResourceType() string {
	return "StructureDefinition"
}
