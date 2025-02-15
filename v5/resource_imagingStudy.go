// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ImagingStudy
type ImagingStudy struct {
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
	// Identifiers for the whole study
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// registered | available | cancelled | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// All of the distinct values for series' modalities
	Modality []CodeableConcept `bson:"modality,omitempty" json:"modality,omitempty"`
	// Who or what is the subject of the study
	Subject custom.Reference[ImagingStudySubject] `bson:"subject" json:"subject"`
	// Encounter with which this imaging study is associated
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the study was started
	Started *custom.DateTime `bson:"started,omitempty" json:"started,omitempty"`
	// Request fulfilled
	BasedOn []custom.Reference[ImagingStudyBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[Procedure] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// Referring physician
	Referrer *custom.Reference[ImagingStudyReferrer] `bson:"referrer,omitempty" json:"referrer,omitempty"`
	// Study access endpoint
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// Number of Study Related Series
	NumberOfSeries *uint32 `bson:"numberOfSeries,omitempty" json:"numberOfSeries,omitempty"`
	// Number of Study Related Instances
	NumberOfInstances *uint32 `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	// The performed procedure or code
	Procedure []custom.CodeableReference[ImagingStudyProcedure] `bson:"procedure,omitempty" json:"procedure,omitempty"`
	// Where ImagingStudy occurred
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Why the study was requested / performed
	Reason []custom.CodeableReference[ImagingStudyReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// User-defined comments
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Institution-generated description
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Each study has one or more series of instances
	Series []ImagingStudySeries `bson:"series,omitempty" json:"series,omitempty"`
}

type ImagingStudySeries struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// DICOM Series Instance UID for the series
	Uid custom.ID `bson:"uid" json:"uid"`
	// Numeric identifier of this series
	Number *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	// The modality used for this series
	Modality CodeableConcept `bson:"modality" json:"modality"`
	// A short human readable summary of the series
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Number of Series Related Instances
	NumberOfInstances *uint32 `bson:"numberOfInstances,omitempty" json:"numberOfInstances,omitempty"`
	// Series access endpoint
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// Body part examined
	BodySite *custom.CodeableReference[BodyStructure] `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Body part laterality
	Laterality *CodeableConcept `bson:"laterality,omitempty" json:"laterality,omitempty"`
	// Specimen imaged
	Specimen []custom.Reference[Specimen] `bson:"specimen,omitempty" json:"specimen,omitempty"`
	// When the series started
	Started *custom.DateTime `bson:"started,omitempty" json:"started,omitempty"`
	// Who performed the series
	Performer []ImagingStudySeriesPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// A single SOP instance from the series
	Instance []ImagingStudySeriesInstance `bson:"instance,omitempty" json:"instance,omitempty"`
}

type ImagingStudySeriesPerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of performance
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who performed the series
	Actor custom.Reference[ImagingStudySeriesPerformerActor] `bson:"actor" json:"actor"`
}

type ImagingStudySeriesInstance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// DICOM SOP Instance UID
	Uid custom.ID `bson:"uid" json:"uid"`
	// DICOM class type
	SopClass Coding `bson:"sopClass" json:"sopClass"`
	// The number of this instance in the series
	Number *uint32 `bson:"number,omitempty" json:"number,omitempty"`
	// Description of instance
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
}

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
	return "StructureDefinition"
}
