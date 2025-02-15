// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
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
	// Plan this is fulfilled by this administration
	BasedOn []custom.Reference[CarePlan] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[MedicationAdministrationPartOf] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// in-progress | not-done | on-hold | completed | entered-in-error | stopped | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Reason administration not performed
	StatusReason []CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// Type of medication administration
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// What was administered
	Medication custom.CodeableReference[Medication] `bson:"medication" json:"medication"`
	// Who received medication
	Subject custom.Reference[MedicationAdministrationSubject] `bson:"subject" json:"subject"`
	// Encounter administered as part of
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Additional information to support administration
	SupportingInformation []custom.Reference[Resource] `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	// Specific date/time or interval of time during which the administration took place (or did not take place)
	Occurence custom.DateTime `bson:"occurence" json:"occurence"`
	// When the MedicationAdministration was first captured in the subject's record
	Recorded *custom.DateTime `bson:"recorded,omitempty" json:"recorded,omitempty"`
	// Full dose was not administered
	IsSubPotent *bool `bson:"isSubPotent,omitempty" json:"isSubPotent,omitempty"`
	// Reason full dose was not administered
	SubPotentReason []CodeableConcept `bson:"subPotentReason,omitempty" json:"subPotentReason,omitempty"`
	// Who or what performed the medication administration and what type of performance they did
	Performer []MedicationAdministrationPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Concept, condition or observation that supports why the medication was administered
	Reason []custom.CodeableReference[MedicationAdministrationReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Request administration performed against
	Request *custom.Reference[MedicationRequest] `bson:"request,omitempty" json:"request,omitempty"`
	// Device used to administer
	Device []custom.CodeableReference[Device] `bson:"device,omitempty" json:"device,omitempty"`
	// Information about the administration
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Details of how medication was taken
	Dosage *MedicationAdministrationDosage `bson:"dosage,omitempty" json:"dosage,omitempty"`
	// A list of events of interest in the lifecycle
	EventHistory []custom.Reference[Provenance] `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

type MedicationAdministrationPerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of performance
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who or what performed the medication administration
	Actor custom.CodeableReference[MedicationAdministrationPerformerActor] `bson:"actor" json:"actor"`
}

type MedicationAdministrationDosage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Free text dosage instructions e.g. SIG
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// Body site administered to
	Site *CodeableConcept `bson:"site,omitempty" json:"site,omitempty"`
	// Path of substance into body
	Route *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	// How drug was administered
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Amount of medication per dose
	Dose *Quantity `bson:"dose,omitempty" json:"dose,omitempty"`
	// Dose quantity per unit of time
	Rate *Ratio `bson:"rate,omitempty" json:"rate,omitempty"`
}

type MedicationAdministrationPartOf interface {
	gofhirclient.Resource

	Is_MedicationAdministrationPartOf()
}

func (m MedicationAdministration) Is_MedicationAdministrationPartOf() {}
func (p Procedure) Is_MedicationAdministrationPartOf()                {}
func (m MedicationDispense) Is_MedicationAdministrationPartOf()       {}

type MedicationAdministrationSubject interface {
	gofhirclient.Resource

	Is_MedicationAdministrationSubject()
}

func (p Patient) Is_MedicationAdministrationSubject() {}
func (g Group) Is_MedicationAdministrationSubject()   {}

type MedicationAdministrationPerformerActor interface {
	gofhirclient.Resource

	Is_MedicationAdministrationPerformerActor()
}

func (p Practitioner) Is_MedicationAdministrationPerformerActor()     {}
func (p PractitionerRole) Is_MedicationAdministrationPerformerActor() {}
func (p Patient) Is_MedicationAdministrationPerformerActor()          {}
func (r RelatedPerson) Is_MedicationAdministrationPerformerActor()    {}
func (d Device) Is_MedicationAdministrationPerformerActor()           {}

type MedicationAdministrationReason interface {
	gofhirclient.Resource

	Is_MedicationAdministrationReason()
}

func (c Condition) Is_MedicationAdministrationReason()        {}
func (o Observation) Is_MedicationAdministrationReason()      {}
func (d DiagnosticReport) Is_MedicationAdministrationReason() {}

func (m MedicationAdministration) ResourceType() string {
	return "StructureDefinition"
}
