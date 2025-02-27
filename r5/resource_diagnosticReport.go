// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DiagnosticReport
type DiagnosticReportMedia struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Comment about the image or data (e.g. explanation)
	Comment *string `json:"comment,omitempty"`
	// Reference to the image or data source
	Link custom.Reference[DocumentReference] `json:"link"`
}

type DiagnosticReport struct {
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
	// Business identifier for report
	Identifier []Identifier `json:"identifier,omitempty"`
	// What was requested
	BasedOn []custom.Reference[DiagnosticReportBasedOn] `json:"basedOn,omitempty"`
	// registered | partial | preliminary | modified | final | amended | corrected | appended | cancelled | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Service category
	Category []CodeableConcept `json:"category,omitempty"`
	// Name/Code for this diagnostic report
	Code CodeableConcept `json:"code"`
	// The subject of the report - usually, but not always, the patient
	Subject *custom.Reference[DiagnosticReportSubject] `json:"subject,omitempty"`
	// Health care event when test ordered
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Clinically relevant time/time-period for report
	EffectiveDateTime *custom.DateTime `json:"effectiveDateTime,omitempty"`
	// Clinically relevant time/time-period for report
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// DateTime this version was made
	Issued *custom.Instant `json:"issued,omitempty"`
	// Responsible Diagnostic Service
	Performer []custom.Reference[DiagnosticReportPerformer] `json:"performer,omitempty"`
	// Primary result interpreter
	ResultsInterpreter []custom.Reference[DiagnosticReportResultsInterpreter] `json:"resultsInterpreter,omitempty"`
	// Specimens this report is based on
	Specimen []custom.Reference[Specimen] `json:"specimen,omitempty"`
	// Observations
	Result []custom.Reference[Observation] `json:"result,omitempty"`
	// Comments about the diagnostic report
	Note []Annotation `json:"note,omitempty"`
	// Reference to full details of an analysis associated with the diagnostic report
	Study []custom.Reference[DiagnosticReportStudy] `json:"study,omitempty"`
	// Additional information supporting the diagnostic report
	SupportingInfo []DiagnosticReportSupportingInfo `json:"supportingInfo,omitempty"`
	// Key images or data associated with this report
	Media []DiagnosticReportMedia `json:"media,omitempty"`
	// Reference to a Composition resource for the DiagnosticReport structure
	Composition *custom.Reference[Composition] `json:"composition,omitempty"`
	// Clinical conclusion (interpretation) of test results
	Conclusion *custom.Markdown `json:"conclusion,omitempty"`
	// Codes for the clinical conclusion of test results
	ConclusionCode []CodeableConcept `json:"conclusionCode,omitempty"`
	// Entire report as issued
	PresentedForm []Attachment `json:"presentedForm,omitempty"`
}

type DiagnosticReportSupportingInfo struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Supporting information role code
	Type CodeableConcept `json:"type"`
	// Supporting information reference
	Reference custom.Reference[DiagnosticReportSupportingInfoReference] `json:"reference"`
}

type OtherDiagnosticReport DiagnosticReport

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
	return "DiagnosticReport"
}

func (d DiagnosticReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDiagnosticReport
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDiagnosticReport: OtherDiagnosticReport(d), ResourceType: d.ResourceType()})
}

func UnmarshalDiagnosticReport(b []byte) (res DiagnosticReport, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
