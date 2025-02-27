// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ImagingStudy
type ImagingStudy struct {
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
	// Identifiers for the whole study
	Identifier []Identifier `json:"identifier,omitempty"`
	// registered | available | cancelled | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// All of the distinct values for series' modalities
	Modality []CodeableConcept `json:"modality,omitempty"`
	// Who or what is the subject of the study
	Subject custom.Reference[ImagingStudySubject] `json:"subject"`
	// Encounter with which this imaging study is associated
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When the study was started
	Started *custom.DateTime `json:"started,omitempty"`
	// Request fulfilled
	BasedOn []custom.Reference[ImagingStudyBasedOn] `json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[Procedure] `json:"partOf,omitempty"`
	// Referring physician
	Referrer *custom.Reference[ImagingStudyReferrer] `json:"referrer,omitempty"`
	// Study access endpoint
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
	// Number of Study Related Series
	NumberOfSeries *uint32 `json:"numberOfSeries,omitempty"`
	// Number of Study Related Instances
	NumberOfInstances *uint32 `json:"numberOfInstances,omitempty"`
	// The performed procedure or code
	Procedure []custom.CodeableReference[ImagingStudyProcedure] `json:"procedure,omitempty"`
	// Where ImagingStudy occurred
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Why the study was requested / performed
	Reason []custom.CodeableReference[ImagingStudyReason] `json:"reason,omitempty"`
	// User-defined comments
	Note []Annotation `json:"note,omitempty"`
	// Institution-generated description
	Description *string `json:"description,omitempty"`
	// Each study has one or more series of instances
	Series []ImagingStudySeries `json:"series,omitempty"`
}

type ImagingStudySeries struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// DICOM Series Instance UID for the series
	Uid custom.ID `json:"uid"`
	// Numeric identifier of this series
	Number *uint32 `json:"number,omitempty"`
	// The modality used for this series
	Modality CodeableConcept `json:"modality"`
	// A short human readable summary of the series
	Description *string `json:"description,omitempty"`
	// Number of Series Related Instances
	NumberOfInstances *uint32 `json:"numberOfInstances,omitempty"`
	// Series access endpoint
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
	// Body part examined
	BodySite *custom.CodeableReference[BodyStructure] `json:"bodySite,omitempty"`
	// Body part laterality
	Laterality *CodeableConcept `json:"laterality,omitempty"`
	// Specimen imaged
	Specimen []custom.Reference[Specimen] `json:"specimen,omitempty"`
	// When the series started
	Started *custom.DateTime `json:"started,omitempty"`
	// Who performed the series
	Performer []ImagingStudySeriesPerformer `json:"performer,omitempty"`
	// A single SOP instance from the series
	Instance []ImagingStudySeriesInstance `json:"instance,omitempty"`
}

type ImagingStudySeriesPerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of performance
	Function *CodeableConcept `json:"function,omitempty"`
	// Who performed the series
	Actor custom.Reference[ImagingStudySeriesPerformerActor] `json:"actor"`
}

type ImagingStudySeriesInstance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// DICOM SOP Instance UID
	Uid custom.ID `json:"uid"`
	// DICOM class type
	SopClass Coding `json:"sopClass"`
	// The number of this instance in the series
	Number *uint32 `json:"number,omitempty"`
	// Description of instance
	Title *string `json:"title,omitempty"`
}

type OtherImagingStudy ImagingStudy

type ImagingStudySubject interface {
	gofhirclient.Resource

	Is_ImagingStudySubject()
}

func (p Patient) Is_ImagingStudySubject() {}
func (d Device) Is_ImagingStudySubject()  {}
func (g Group) Is_ImagingStudySubject()   {}

type ImagingStudyBasedOn interface {
	gofhirclient.Resource

	Is_ImagingStudyBasedOn()
}

func (c CarePlan) Is_ImagingStudyBasedOn()            {}
func (s ServiceRequest) Is_ImagingStudyBasedOn()      {}
func (a Appointment) Is_ImagingStudyBasedOn()         {}
func (a AppointmentResponse) Is_ImagingStudyBasedOn() {}
func (t Task) Is_ImagingStudyBasedOn()                {}

type ImagingStudyReferrer interface {
	gofhirclient.Resource

	Is_ImagingStudyReferrer()
}

func (p Practitioner) Is_ImagingStudyReferrer()     {}
func (p PractitionerRole) Is_ImagingStudyReferrer() {}

type ImagingStudyProcedure interface {
	gofhirclient.Resource

	Is_ImagingStudyProcedure()
}

func (p PlanDefinition) Is_ImagingStudyProcedure()     {}
func (a ActivityDefinition) Is_ImagingStudyProcedure() {}

type ImagingStudyReason interface {
	gofhirclient.Resource

	Is_ImagingStudyReason()
}

func (c Condition) Is_ImagingStudyReason()         {}
func (o Observation) Is_ImagingStudyReason()       {}
func (d DiagnosticReport) Is_ImagingStudyReason()  {}
func (d DocumentReference) Is_ImagingStudyReason() {}

type ImagingStudySeriesPerformerActor interface {
	gofhirclient.Resource

	Is_ImagingStudySeriesPerformerActor()
}

func (p Practitioner) Is_ImagingStudySeriesPerformerActor()      {}
func (p PractitionerRole) Is_ImagingStudySeriesPerformerActor()  {}
func (o Organization) Is_ImagingStudySeriesPerformerActor()      {}
func (c CareTeam) Is_ImagingStudySeriesPerformerActor()          {}
func (p Patient) Is_ImagingStudySeriesPerformerActor()           {}
func (d Device) Is_ImagingStudySeriesPerformerActor()            {}
func (r RelatedPerson) Is_ImagingStudySeriesPerformerActor()     {}
func (h HealthcareService) Is_ImagingStudySeriesPerformerActor() {}

func (i ImagingStudy) ResourceType() string {
	return "ImagingStudy"
}

func (i ImagingStudy) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImagingStudy
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherImagingStudy: OtherImagingStudy(i), ResourceType: i.ResourceType()})
}

func UnmarshalImagingStudy(b []byte) (res ImagingStudy, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
