// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ImagingSelection
type ImagingSelectionInstance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// DICOM SOP Instance UID
	Uid custom.ID `json:"uid"`
	// DICOM Instance Number
	Number *uint32 `json:"number,omitempty"`
	// DICOM SOP Class UID
	SopClass *Coding `json:"sopClass,omitempty"`
	// The selected subset of the SOP Instance
	Subset []string `json:"subset,omitempty"`
	// A specific 2D region in a DICOM image / frame
	ImageRegion2d []ImagingSelectionInstanceImageRegion2d `json:"imageRegion2d,omitempty"`
	// A specific 3D region in a DICOM frame of reference
	ImageRegion3d []ImagingSelectionInstanceImageRegion3d `json:"imageRegion3d,omitempty"`
}

type ImagingSelectionInstanceImageRegion2d struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// point | polyline | interpolated | circle | ellipse
	RegionType custom.Code `json:"regionType"`
	// Specifies the coordinates that define the image region
	Coordinate []json.Number `json:"coordinate"`
}

type ImagingSelectionInstanceImageRegion3d struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// point | multipoint | polyline | polygon | ellipse | ellipsoid
	RegionType custom.Code `json:"regionType"`
	// Specifies the coordinates that define the image region
	Coordinate []json.Number `json:"coordinate"`
}

type ImagingSelection struct {
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
	// Business Identifier for Imaging Selection
	Identifier []Identifier `json:"identifier,omitempty"`
	// available | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Subject of the selected instances
	Subject *custom.Reference[ImagingSelectionSubject] `json:"subject,omitempty"`
	// Date / Time when this imaging selection was created
	Issued *custom.Instant `json:"issued,omitempty"`
	// Selector of the instances (human or machine)
	Performer []ImagingSelectionPerformer `json:"performer,omitempty"`
	// Associated request
	BasedOn []custom.Reference[ImagingSelectionBasedOn] `json:"basedOn,omitempty"`
	// Classifies the imaging selection
	Category []CodeableConcept `json:"category,omitempty"`
	// Imaging Selection purpose text or code
	Code CodeableConcept `json:"code"`
	// DICOM Study Instance UID
	StudyUid *custom.ID `json:"studyUid,omitempty"`
	// The imaging study from which the imaging selection is derived
	DerivedFrom []custom.Reference[ImagingSelectionDerivedFrom] `json:"derivedFrom,omitempty"`
	// The network service providing retrieval for the images referenced in the imaging selection
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
	// DICOM Series Instance UID
	SeriesUid *custom.ID `json:"seriesUid,omitempty"`
	// DICOM Series Number
	SeriesNumber *uint32 `json:"seriesNumber,omitempty"`
	// The Frame of Reference UID for the selected images
	FrameOfReferenceUid *custom.ID `json:"frameOfReferenceUid,omitempty"`
	// Body part examined
	BodySite *custom.CodeableReference[BodyStructure] `json:"bodySite,omitempty"`
	// Related resource that is the focus for the imaging selection
	Focus []custom.Reference[ImagingSelection] `json:"focus,omitempty"`
	// The selected instances
	Instance []ImagingSelectionInstance `json:"instance,omitempty"`
}

type ImagingSelectionPerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of performer
	Function *CodeableConcept `json:"function,omitempty"`
	// Author (human or machine)
	Actor *custom.Reference[ImagingSelectionPerformerActor] `json:"actor,omitempty"`
}

type OtherImagingSelection ImagingSelection

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
	return "ImagingSelection"
}

func (i ImagingSelection) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingSelection
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherImagingSelection: OtherImagingSelection(i), ResourceType: i.ResourceType()})
}

func UnmarshalImagingSelection(b []byte) (res ImagingSelection, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
