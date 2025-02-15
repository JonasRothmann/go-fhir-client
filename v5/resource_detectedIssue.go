// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DetectedIssue
type DetectedIssue struct {
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
	// Unique id for the detected issue
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// preliminary | final | entered-in-error | mitigated
	Status custom.Code `bson:"status" json:"status"`
	// Type of detected issue, e.g. drug-drug, duplicate therapy, etc
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Specific type of detected issue, e.g. drug-drug, duplicate therapy, etc
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// high | moderate | low
	Severity *custom.Code `bson:"severity,omitempty" json:"severity,omitempty"`
	// Associated subject
	Subject *custom.Reference[DetectedIssueSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Encounter detected issue is part of
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When identified
	Identified *custom.DateTime `bson:"identified,omitempty" json:"identified,omitempty"`
	// The provider or device that identified the issue
	Author *custom.Reference[DetectedIssueAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Problem resource
	Implicated []custom.Reference[Resource] `bson:"implicated,omitempty" json:"implicated,omitempty"`
	// Supporting evidence
	Evidence []DetectedIssueEvidence `bson:"evidence,omitempty" json:"evidence,omitempty"`
	// Description and context
	Detail *custom.Markdown `bson:"detail,omitempty" json:"detail,omitempty"`
	// Authority for issue
	Reference *custom.URI `bson:"reference,omitempty" json:"reference,omitempty"`
	// Step taken to address
	Mitigation []DetectedIssueMitigation `bson:"mitigation,omitempty" json:"mitigation,omitempty"`
}

type DetectedIssueEvidence struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Manifestation
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Supporting information
	Detail []custom.Reference[Resource] `bson:"detail,omitempty" json:"detail,omitempty"`
}

type DetectedIssueMitigation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What mitigation?
	Action CodeableConcept `bson:"action" json:"action"`
	// Date committed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Who is committing?
	Author *custom.Reference[DetectedIssueMitigationAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Additional notes about the mitigation
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

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
	return "StructureDefinition"
}
