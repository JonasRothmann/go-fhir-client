// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/BodyStructure
type BodyStructureIncludedStructure struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code that represents the included structure
	Structure CodeableConcept `json:"structure"`
	// Code that represents the included structure laterality
	Laterality *CodeableConcept `json:"laterality,omitempty"`
	// Landmark relative location
	BodyLandmarkOrientation []BodyStructureIncludedStructureBodyLandmarkOrientation `json:"bodyLandmarkOrientation,omitempty"`
	// Cartesian reference for structure
	SpatialReference []custom.Reference[ImagingSelection] `json:"spatialReference,omitempty"`
	// Code that represents the included structure qualifier
	Qualifier []CodeableConcept `json:"qualifier,omitempty"`
}

type BodyStructureIncludedStructureBodyLandmarkOrientation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Body ]andmark description
	LandmarkDescription []CodeableConcept `json:"landmarkDescription,omitempty"`
	// Clockface orientation
	ClockFacePosition []CodeableConcept `json:"clockFacePosition,omitempty"`
	// Landmark relative location
	DistanceFromLandmark []BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark `json:"distanceFromLandmark,omitempty"`
	// Relative landmark surface orientation
	SurfaceOrientation []CodeableConcept `json:"surfaceOrientation,omitempty"`
}

type BodyStructureIncludedStructureBodyLandmarkOrientationDistanceFromLandmark struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Measurement device
	Device []custom.CodeableReference[Device] `json:"device,omitempty"`
	// Measured distance from body landmark
	Value []Quantity `json:"value,omitempty"`
}

type BodyStructure struct {
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
	// Bodystructure identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Whether this record is in active use
	Active *bool `json:"active,omitempty"`
	// Kind of Structure
	Morphology *CodeableConcept `json:"morphology,omitempty"`
	// Included anatomic location(s)
	IncludedStructure []BodyStructureIncludedStructure `json:"includedStructure"`
	// Excluded anatomic locations(s)
	ExcludedStructure []interface{} `json:"excludedStructure,omitempty"`
	// Text description
	Description *custom.Markdown `json:"description,omitempty"`
	// Attached images
	Image []Attachment `json:"image,omitempty"`
	// Who this is about
	Patient custom.Reference[Patient] `json:"patient"`
}

type OtherBodyStructure BodyStructure

func (b BodyStructure) ResourceType() string {
	return "BodyStructure"
}

func (b BodyStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBodyStructure
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherBodyStructure: OtherBodyStructure(b), ResourceType: b.ResourceType()})
}

func UnmarshalBodyStructure(b []byte) (res BodyStructure, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
