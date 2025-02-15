// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicationStatement
type MedicationStatement struct {
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
	// External identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[MedicationStatementPartOf] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// recorded | entered-in-error | draft
	Status custom.Code `bson:"status" json:"status"`
	// Type of medication statement
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// What medication was taken
	Medication custom.CodeableReference[Medication] `bson:"medication" json:"medication"`
	// Who is/was taking  the medication
	Subject custom.Reference[MedicationStatementSubject] `bson:"subject" json:"subject"`
	// Encounter associated with MedicationStatement
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// The date/time or interval when the medication is/was/will be taken
	Effective *custom.DateTime `bson:"effective,omitempty" json:"effective,omitempty"`
	// When the usage was asserted?
	DateAsserted *custom.DateTime `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	// Person or organization that provided the information about the taking of this medication
	InformationSource []custom.Reference[MedicationStatementInformationSource] `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	// Link to information used to derive the MedicationStatement
	DerivedFrom []custom.Reference[Resource] `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	// Reason for why the medication is being/was taken
	Reason []custom.CodeableReference[MedicationStatementReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Further information about the usage
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Link to information relevant to the usage of a medication
	RelatedClinicalInformation []custom.Reference[MedicationStatementRelatedClinicalInformation] `bson:"relatedClinicalInformation,omitempty" json:"relatedClinicalInformation,omitempty"`
	// Full representation of the dosage instructions
	RenderedDosageInstruction *custom.Markdown `bson:"renderedDosageInstruction,omitempty" json:"renderedDosageInstruction,omitempty"`
	// Details of how medication is/was taken or should be taken
	Dosage []Dosage `bson:"dosage,omitempty" json:"dosage,omitempty"`
	// Indicates whether the medication is or is not being consumed or administered
	Adherence *MedicationStatementAdherence `bson:"adherence,omitempty" json:"adherence,omitempty"`
}

type MedicationStatementAdherence struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of adherence
	Code CodeableConcept `bson:"code" json:"code"`
	// Details of the reason for the current use of the medication
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

type MedicationStatementPartOf interface {
	gofhirclient.Resource

	Is_MedicationStatementPartOf()
}

func (p Procedure) Is_MedicationStatementPartOf()           {}
func (m MedicationStatement) Is_MedicationStatementPartOf() {}

type MedicationStatementSubject interface {
	gofhirclient.Resource

	Is_MedicationStatementSubject()
}

func (p Patient) Is_MedicationStatementSubject() {}
func (g Group) Is_MedicationStatementSubject()   {}

type MedicationStatementInformationSource interface {
	gofhirclient.Resource

	Is_MedicationStatementInformationSource()
}

func (p Patient) Is_MedicationStatementInformationSource()          {}
func (p Practitioner) Is_MedicationStatementInformationSource()     {}
func (p PractitionerRole) Is_MedicationStatementInformationSource() {}
func (r RelatedPerson) Is_MedicationStatementInformationSource()    {}
func (o Organization) Is_MedicationStatementInformationSource()     {}

type MedicationStatementReason interface {
	gofhirclient.Resource

	Is_MedicationStatementReason()
}

func (c Condition) Is_MedicationStatementReason()        {}
func (o Observation) Is_MedicationStatementReason()      {}
func (d DiagnosticReport) Is_MedicationStatementReason() {}

type MedicationStatementRelatedClinicalInformation interface {
	gofhirclient.Resource

	Is_MedicationStatementRelatedClinicalInformation()
}

func (o Observation) Is_MedicationStatementRelatedClinicalInformation() {}
func (c Condition) Is_MedicationStatementRelatedClinicalInformation()   {}

func (m MedicationStatement) ResourceType() string {
	return "StructureDefinition"
}
