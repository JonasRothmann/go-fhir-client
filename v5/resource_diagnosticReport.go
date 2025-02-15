// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DiagnosticReport
type DiagnosticReportMedia struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Comment about the image or data (e.g. explanation)
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty"`
	// Reference to the image or data source
	Link custom.Reference[DocumentReference] `bson:"link" json:"link"`
}

type DiagnosticReport struct {
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
	// Business identifier for report
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// What was requested
	BasedOn []custom.Reference[DiagnosticReportBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// registered | partial | preliminary | modified | final | amended | corrected | appended | cancelled | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Service category
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Name/Code for this diagnostic report
	Code CodeableConcept `bson:"code" json:"code"`
	// The subject of the report - usually, but not always, the patient
	Subject *custom.Reference[DiagnosticReportSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Health care event when test ordered
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Clinically relevant time/time-period for report
	Effective *custom.DateTime `bson:"effective,omitempty" json:"effective,omitempty"`
	// DateTime this version was made
	Issued *custom.Instant `bson:"issued,omitempty" json:"issued,omitempty"`
	// Responsible Diagnostic Service
	Performer []custom.Reference[DiagnosticReportPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Primary result interpreter
	ResultsInterpreter []custom.Reference[DiagnosticReportResultsInterpreter] `bson:"resultsInterpreter,omitempty" json:"resultsInterpreter,omitempty"`
	// Specimens this report is based on
	Specimen []custom.Reference[Specimen] `bson:"specimen,omitempty" json:"specimen,omitempty"`
	// Observations
	Result []custom.Reference[Observation] `bson:"result,omitempty" json:"result,omitempty"`
	// Comments about the diagnostic report
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Reference to full details of an analysis associated with the diagnostic report
	Study []custom.Reference[DiagnosticReportStudy] `bson:"study,omitempty" json:"study,omitempty"`
	// Additional information supporting the diagnostic report
	SupportingInfo []DiagnosticReportSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Key images or data associated with this report
	Media []DiagnosticReportMedia `bson:"media,omitempty" json:"media,omitempty"`
	// Reference to a Composition resource for the DiagnosticReport structure
	Composition *custom.Reference[Composition] `bson:"composition,omitempty" json:"composition,omitempty"`
	// Clinical conclusion (interpretation) of test results
	Conclusion *custom.Markdown `bson:"conclusion,omitempty" json:"conclusion,omitempty"`
	// Codes for the clinical conclusion of test results
	ConclusionCode []CodeableConcept `bson:"conclusionCode,omitempty" json:"conclusionCode,omitempty"`
	// Entire report as issued
	PresentedForm []Attachment `bson:"presentedForm,omitempty" json:"presentedForm,omitempty"`
}

type DiagnosticReportSupportingInfo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Supporting information role code
	Type CodeableConcept `bson:"type" json:"type"`
	// Supporting information reference
	Reference custom.Reference[DiagnosticReportSupportingInfoReference] `bson:"reference" json:"reference"`
}

type DiagnosticReportBasedOn interface {
	gofhirclient.Resource

	Is_DiagnosticReportBasedOn()
}

func (c CarePlan) Is_DiagnosticReportBasedOn()                   {}
func (i ImmunizationRecommendation) Is_DiagnosticReportBasedOn() {}
func (m MedicationRequest) Is_DiagnosticReportBasedOn()          {}
func (n NutritionOrder) Is_DiagnosticReportBasedOn()             {}
func (s ServiceRequest) Is_DiagnosticReportBasedOn()             {}

type DiagnosticReportSubject interface {
	gofhirclient.Resource

	Is_DiagnosticReportSubject()
}

func (p Patient) Is_DiagnosticReportSubject()                    {}
func (g Group) Is_DiagnosticReportSubject()                      {}
func (d Device) Is_DiagnosticReportSubject()                     {}
func (l Location) Is_DiagnosticReportSubject()                   {}
func (o Organization) Is_DiagnosticReportSubject()               {}
func (p Practitioner) Is_DiagnosticReportSubject()               {}
func (m Medication) Is_DiagnosticReportSubject()                 {}
func (s Substance) Is_DiagnosticReportSubject()                  {}
func (b BiologicallyDerivedProduct) Is_DiagnosticReportSubject() {}

type DiagnosticReportPerformer interface {
	gofhirclient.Resource

	Is_DiagnosticReportPerformer()
}

func (p Practitioner) Is_DiagnosticReportPerformer()     {}
func (p PractitionerRole) Is_DiagnosticReportPerformer() {}
func (o Organization) Is_DiagnosticReportPerformer()     {}
func (c CareTeam) Is_DiagnosticReportPerformer()         {}

type DiagnosticReportResultsInterpreter interface {
	gofhirclient.Resource

	Is_DiagnosticReportResultsInterpreter()
}

func (p Practitioner) Is_DiagnosticReportResultsInterpreter()     {}
func (p PractitionerRole) Is_DiagnosticReportResultsInterpreter() {}
func (o Organization) Is_DiagnosticReportResultsInterpreter()     {}
func (c CareTeam) Is_DiagnosticReportResultsInterpreter()         {}

type DiagnosticReportStudy interface {
	gofhirclient.Resource

	Is_DiagnosticReportStudy()
}

func (g GenomicStudy) Is_DiagnosticReportStudy() {}
func (i ImagingStudy) Is_DiagnosticReportStudy() {}

type DiagnosticReportSupportingInfoReference interface {
	gofhirclient.Resource

	Is_DiagnosticReportSupportingInfoReference()
}

func (p Procedure) Is_DiagnosticReportSupportingInfoReference()        {}
func (o Observation) Is_DiagnosticReportSupportingInfoReference()      {}
func (d DiagnosticReport) Is_DiagnosticReportSupportingInfoReference() {}
func (c Citation) Is_DiagnosticReportSupportingInfoReference()         {}

func (d DiagnosticReport) ResourceType() string {
	return "StructureDefinition"
}
