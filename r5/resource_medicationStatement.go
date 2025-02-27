// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicationStatement
type MedicationStatement struct {
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
	// External identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[MedicationStatementPartOf] `json:"partOf,omitempty"`
	// recorded | entered-in-error | draft
	Status custom.Code `json:"status"`
	// Type of medication statement
	Category []CodeableConcept `json:"category,omitempty"`
	// What medication was taken
	Medication custom.CodeableReference[Medication] `json:"medication"`
	// Who is/was taking  the medication
	Subject custom.Reference[MedicationStatementSubject] `json:"subject"`
	// Encounter associated with MedicationStatement
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// The date/time or interval when the medication is/was/will be taken
	EffectiveDateTime *custom.DateTime `json:"effectiveDateTime,omitempty"`
	// The date/time or interval when the medication is/was/will be taken
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// The date/time or interval when the medication is/was/will be taken
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	// When the usage was asserted?
	DateAsserted *custom.DateTime `json:"dateAsserted,omitempty"`
	// Person or organization that provided the information about the taking of this medication
	InformationSource []custom.Reference[MedicationStatementInformationSource] `json:"informationSource,omitempty"`
	// Link to information used to derive the MedicationStatement
	DerivedFrom []custom.Reference[Resource] `json:"derivedFrom,omitempty"`
	// Reason for why the medication is being/was taken
	Reason []custom.CodeableReference[MedicationStatementReason] `json:"reason,omitempty"`
	// Further information about the usage
	Note []Annotation `json:"note,omitempty"`
	// Link to information relevant to the usage of a medication
	RelatedClinicalInformation []custom.Reference[MedicationStatementRelatedClinicalInformation] `json:"relatedClinicalInformation,omitempty"`
	// Full representation of the dosage instructions
	RenderedDosageInstruction *custom.Markdown `json:"renderedDosageInstruction,omitempty"`
	// Details of how medication is/was taken or should be taken
	Dosage []Dosage `json:"dosage,omitempty"`
	// Indicates whether the medication is or is not being consumed or administered
	Adherence *MedicationStatementAdherence `json:"adherence,omitempty"`
}

type MedicationStatementAdherence struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of adherence
	Code CodeableConcept `json:"code"`
	// Details of the reason for the current use of the medication
	Reason *CodeableConcept `json:"reason,omitempty"`
}

type OtherMedicationStatement MedicationStatement

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
	return "MedicationStatement"
}

func (m MedicationStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationStatement
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMedicationStatement: OtherMedicationStatement(m), ResourceType: m.ResourceType()})
}

func UnmarshalMedicationStatement(b []byte) (res MedicationStatement, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
