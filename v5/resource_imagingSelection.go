// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ImagingSelection
type ImagingSelectionInstance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// DICOM SOP Instance UID
	Uid custom.ID `bson:"uid" json:"uid"`
	// DICOM Instance Number
	Number *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	// DICOM SOP Class UID
	SopClass *Coding `bson:"sopClass,omitempty" json:"sopClass,omitempty"`
	// The selected subset of the SOP Instance
	Subset []string `bson:"subset,omitempty" json:"subset,omitempty"`
	// A specific 2D region in a DICOM image / frame
	ImageRegion2d []ImagingSelectionInstanceImageRegion2d `bson:"imageRegion2d,omitempty" json:"imageRegion2d,omitempty"`
	// A specific 3D region in a DICOM frame of reference
	ImageRegion3d []ImagingSelectionInstanceImageRegion3d `bson:"imageRegion3d,omitempty" json:"imageRegion3d,omitempty"`
}

type ImagingSelectionInstanceImageRegion2d struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// point | polyline | interpolated | circle | ellipse
	RegionType custom.Code `bson:"regionType" json:"regionType"`
	// Specifies the coordinates that define the image region
	Coordinate []json.Number `bson:"coordinate" json:"coordinate"`
}

type ImagingSelectionInstanceImageRegion3d struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// point | multipoint | polyline | polygon | ellipse | ellipsoid
	RegionType custom.Code `bson:"regionType" json:"regionType"`
	// Specifies the coordinates that define the image region
	Coordinate []json.Number `bson:"coordinate" json:"coordinate"`
}

type ImagingSelection struct {
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
	// Business Identifier for Imaging Selection
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// available | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Subject of the selected instances
	Subject *custom.Reference[ImagingSelectionSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Date / Time when this imaging selection was created
	Issued *custom.Instant `bson:"issued,omitempty" json:"issued,omitempty"`
	// Selector of the instances (human or machine)
	Performer []ImagingSelectionPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Associated request
	BasedOn []custom.Reference[ImagingSelectionBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Classifies the imaging selection
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Imaging Selection purpose text or code
	Code CodeableConcept `bson:"code" json:"code"`
	// DICOM Study Instance UID
	StudyUid *custom.ID `bson:"studyUid,omitempty" json:"studyUid,omitempty"`
	// The imaging study from which the imaging selection is derived
	DerivedFrom []custom.Reference[ImagingSelectionDerivedFrom] `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	// The network service providing retrieval for the images referenced in the imaging selection
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// DICOM Series Instance UID
	SeriesUid *custom.ID `bson:"seriesUid,omitempty" json:"seriesUid,omitempty"`
	// DICOM Series Number
	SeriesNumber *uint32 `bson:"seriesNumber,omitempty" json:"seriesNumber,omitempty"`
	// The Frame of Reference UID for the selected images
	FrameOfReferenceUid *custom.ID `bson:"frameOfReferenceUid,omitempty" json:"frameOfReferenceUid,omitempty"`
	// Body part examined
	BodySite *custom.CodeableReference[BodyStructure] `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Related resource that is the focus for the imaging selection
	Focus []custom.Reference[ImagingSelection] `bson:"focus,omitempty" json:"focus,omitempty"`
	// The selected instances
	Instance []ImagingSelectionInstance `bson:"instance,omitempty" json:"instance,omitempty"`
}

type ImagingSelectionPerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of performer
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Author (human or machine)
	Actor *custom.Reference[ImagingSelectionPerformerActor] `bson:"actor,omitempty" json:"actor,omitempty"`
}

type ImagingSelectionSubject interface {
	gofhirclient.Resource

	Is_ImagingSelectionSubject()
}

func (p Patient) Is_ImagingSelectionSubject()      {}
func (g Group) Is_ImagingSelectionSubject()        {}
func (d Device) Is_ImagingSelectionSubject()       {}
func (l Location) Is_ImagingSelectionSubject()     {}
func (o Organization) Is_ImagingSelectionSubject() {}
func (p Procedure) Is_ImagingSelectionSubject()    {}
func (p Practitioner) Is_ImagingSelectionSubject() {}
func (m Medication) Is_ImagingSelectionSubject()   {}
func (s Substance) Is_ImagingSelectionSubject()    {}
func (s Specimen) Is_ImagingSelectionSubject()     {}

type ImagingSelectionPerformerActor interface {
	gofhirclient.Resource

	Is_ImagingSelectionPerformerActor()
}

func (p Practitioner) Is_ImagingSelectionPerformerActor()      {}
func (p PractitionerRole) Is_ImagingSelectionPerformerActor()  {}
func (d Device) Is_ImagingSelectionPerformerActor()            {}
func (o Organization) Is_ImagingSelectionPerformerActor()      {}
func (c CareTeam) Is_ImagingSelectionPerformerActor()          {}
func (p Patient) Is_ImagingSelectionPerformerActor()           {}
func (r RelatedPerson) Is_ImagingSelectionPerformerActor()     {}
func (h HealthcareService) Is_ImagingSelectionPerformerActor() {}

type ImagingSelectionBasedOn interface {
	gofhirclient.Resource

	Is_ImagingSelectionBasedOn()
}

func (c CarePlan) Is_ImagingSelectionBasedOn()            {}
func (s ServiceRequest) Is_ImagingSelectionBasedOn()      {}
func (a Appointment) Is_ImagingSelectionBasedOn()         {}
func (a AppointmentResponse) Is_ImagingSelectionBasedOn() {}
func (t Task) Is_ImagingSelectionBasedOn()                {}

type ImagingSelectionDerivedFrom interface {
	gofhirclient.Resource

	Is_ImagingSelectionDerivedFrom()
}

func (i ImagingStudy) Is_ImagingSelectionDerivedFrom()      {}
func (d DocumentReference) Is_ImagingSelectionDerivedFrom() {}

func (i ImagingSelection) ResourceType() string {
	return "StructureDefinition"
}
