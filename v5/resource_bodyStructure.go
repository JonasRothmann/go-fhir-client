// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/BodyStructure
type BodyStructure struct {
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
	// Bodystructure identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether this record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// Kind of Structure
	Morphology *CodeableConcept `bson:"morphology,omitempty" json:"morphology,omitempty"`
	// Included anatomic location(s)
	IncludedStructure []BodyStructureIncludedStructure `bson:"includedStructure" json:"includedStructure"`
	// Excluded anatomic locations(s)
	ExcludedStructure []interface{} `bson:"excludedStructure,omitempty" json:"excludedStructure,omitempty"`
	// Text description
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Attached images
	Image []Attachment `bson:"image,omitempty" json:"image,omitempty"`
	// Who this is about
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
}

type BodyStructureIncludedStructure struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code that represents the included structure
	Structure CodeableConcept `bson:"structure" json:"structure"`
	// Code that represents the included structure laterality
	Laterality *CodeableConcept `bson:"laterality,omitempty" json:"laterality,omitempty"`
	// Landmark relative location
	BodyLandmarkOrientation []BodyStructureIncludedStructureBodyLandmarkOrientation `bson:"bodyLandmarkOrientation,omitempty" json:"bodyLandmarkOrientation,omitempty"`
	// Cartesian reference for structure
	SpatialReference []custom.Reference[ImagingSelection] `bson:"spatialReference,omitempty" json:"spatialReference,omitempty"`
	// Code that represents the included structure qualifier
	Qualifier []CodeableConcept `bson:"qualifier,omitempty" json:"qualifier,omitempty"`
}

type BodyStructureIncludedStructureBodyLandmarkOrientation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Body ]andmark description
	LandmarkDescription []CodeableConcept `bson:"landmarkDescription,omitempty" json:"landmarkDescription,omitempty"`
	// Clockface orientation
	ClockFacePosition []CodeableConcept `bson:"clockFacePosition,omitempty" json:"clockFacePosition,omitempty"`
	// Landmark relative location
	DistanceFromLandmark []BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark `bson:"distanceFromLandmark,omitempty" json:"distanceFromLandmark,omitempty"`
	// Relative landmark surface orientation
	SurfaceOrientation []CodeableConcept `bson:"surfaceOrientation,omitempty" json:"surfaceOrientation,omitempty"`
}

type BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Measurement device
	Device []custom.CodeableReference[Device] `bson:"device,omitempty" json:"device,omitempty"`
	// Measured distance from body landmark
	Value []Quantity `bson:"value,omitempty" json:"value,omitempty"`
}

func (b BodyStructure) ResourceType() string {
	return "StructureDefinition"
}
