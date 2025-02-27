// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DetectedIssue
type DetectedIssue struct {
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
	// Unique id for the detected issue
	Identifier []Identifier `json:"identifier,omitempty"`
	// preliminary | final | entered-in-error | mitigated
	Status custom.Code `json:"status"`
	// Type of detected issue, e.g. drug-drug, duplicate therapy, etc
	Category []CodeableConcept `json:"category,omitempty"`
	// Specific type of detected issue, e.g. drug-drug, duplicate therapy, etc
	Code *CodeableConcept `json:"code,omitempty"`
	// high | moderate | low
	Severity *custom.Code `json:"severity,omitempty"`
	// Associated subject
	Subject *custom.Reference[DetectedIssueSubject] `json:"subject,omitempty"`
	// Encounter detected issue is part of
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When identified
	IdentifiedDateTime *custom.DateTime `json:"identifiedDateTime,omitempty"`
	// When identified
	IdentifiedPeriod *Period `json:"identifiedPeriod,omitempty"`
	// The provider or device that identified the issue
	Author *custom.Reference[DetectedIssueAuthor] `json:"author,omitempty"`
	// Problem resource
	Implicated []custom.Reference[Resource] `json:"implicated,omitempty"`
	// Supporting evidence
	Evidence []DetectedIssueEvidence `json:"evidence,omitempty"`
	// Description and context
	Detail *custom.Markdown `json:"detail,omitempty"`
	// Authority for issue
	Reference *custom.URI `json:"reference,omitempty"`
	// Step taken to address
	Mitigation []DetectedIssueMitigation `json:"mitigation,omitempty"`
}

type DetectedIssueEvidence struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Manifestation
	Code []CodeableConcept `json:"code,omitempty"`
	// Supporting information
	Detail []custom.Reference[Resource] `json:"detail,omitempty"`
}

type DetectedIssueMitigation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What mitigation?
	Action CodeableConcept `json:"action"`
	// Date committed
	Date *custom.DateTime `json:"date,omitempty"`
	// Who is committing?
	Author *custom.Reference[DetectedIssueMitigationAuthor] `json:"author,omitempty"`
	// Additional notes about the mitigation
	Note []Annotation `json:"note,omitempty"`
}

type OtherDetectedIssue DetectedIssue

type DetectedIssueSubject interface {
	gofhirclient.Resource

	Is_DetectedIssueSubject()
}

func (p Patient) Is_DetectedIssueSubject()                    {}
func (g Group) Is_DetectedIssueSubject()                      {}
func (d Device) Is_DetectedIssueSubject()                     {}
func (l Location) Is_DetectedIssueSubject()                   {}
func (o Organization) Is_DetectedIssueSubject()               {}
func (p Procedure) Is_DetectedIssueSubject()                  {}
func (p Practitioner) Is_DetectedIssueSubject()               {}
func (m Medication) Is_DetectedIssueSubject()                 {}
func (s Substance) Is_DetectedIssueSubject()                  {}
func (b BiologicallyDerivedProduct) Is_DetectedIssueSubject() {}
func (n NutritionProduct) Is_DetectedIssueSubject()           {}

type DetectedIssueAuthor interface {
	gofhirclient.Resource

	Is_DetectedIssueAuthor()
}

func (p Patient) Is_DetectedIssueAuthor()          {}
func (r RelatedPerson) Is_DetectedIssueAuthor()    {}
func (p Practitioner) Is_DetectedIssueAuthor()     {}
func (p PractitionerRole) Is_DetectedIssueAuthor() {}
func (d Device) Is_DetectedIssueAuthor()           {}

type DetectedIssueMitigationAuthor interface {
	gofhirclient.Resource

	Is_DetectedIssueMitigationAuthor()
}

func (p Practitioner) Is_DetectedIssueMitigationAuthor()     {}
func (p PractitionerRole) Is_DetectedIssueMitigationAuthor() {}

func (d DetectedIssue) ResourceType() string {
	return "DetectedIssue"
}

func (d DetectedIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDetectedIssue
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDetectedIssue: OtherDetectedIssue(d), ResourceType: d.ResourceType()})
}

func UnmarshalDetectedIssue(b []byte) (res DetectedIssue, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
