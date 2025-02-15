// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicationDispense
type MedicationDispense struct {
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
	// Plan that is fulfilled by this dispense
	BasedOn []custom.Reference[CarePlan] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Event that dispense is part of
	PartOf []custom.Reference[MedicationDispensePartOf] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// preparation | in-progress | cancelled | on-hold | completed | entered-in-error | stopped | declined | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Why a dispense was not performed
	NotPerformedReason *custom.CodeableReference[DetectedIssue] `bson:"notPerformedReason,omitempty" json:"notPerformedReason,omitempty"`
	// When the status changed
	StatusChanged *custom.DateTime `bson:"statusChanged,omitempty" json:"statusChanged,omitempty"`
	// Type of medication dispense
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// What medication was supplied
	Medication custom.CodeableReference[Medication] `bson:"medication" json:"medication"`
	// Who the dispense is for
	Subject custom.Reference[MedicationDispenseSubject] `bson:"subject" json:"subject"`
	// Encounter associated with event
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Information that supports the dispensing of the medication
	SupportingInformation []custom.Reference[Resource] `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	// Who performed event
	Performer []MedicationDispensePerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Where the dispense occurred
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Medication order that authorizes the dispense
	AuthorizingPrescription []custom.Reference[MedicationRequest] `bson:"authorizingPrescription,omitempty" json:"authorizingPrescription,omitempty"`
	// Trial fill, partial fill, emergency fill, etc
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Amount dispensed
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Amount of medication expressed as a timing amount
	DaysSupply *Quantity `bson:"daysSupply,omitempty" json:"daysSupply,omitempty"`
	// When the recording of the dispense started
	Recorded *custom.DateTime `bson:"recorded,omitempty" json:"recorded,omitempty"`
	// When product was packaged and reviewed
	WhenPrepared *custom.DateTime `bson:"whenPrepared,omitempty" json:"whenPrepared,omitempty"`
	// When product was given out
	WhenHandedOver *custom.DateTime `bson:"whenHandedOver,omitempty" json:"whenHandedOver,omitempty"`
	// Where the medication was/will be sent
	Destination *custom.Reference[Location] `bson:"destination,omitempty" json:"destination,omitempty"`
	// Who collected the medication or where the medication was delivered
	Receiver []custom.Reference[MedicationDispenseReceiver] `bson:"receiver,omitempty" json:"receiver,omitempty"`
	// Information about the dispense
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Full representation of the dosage instructions
	RenderedDosageInstruction *custom.Markdown `bson:"renderedDosageInstruction,omitempty" json:"renderedDosageInstruction,omitempty"`
	// How the medication is to be used by the patient or administered by the caregiver
	DosageInstruction []Dosage `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	// Whether a substitution was performed on the dispense
	Substitution *MedicationDispenseSubstitution `bson:"substitution,omitempty" json:"substitution,omitempty"`
	// A list of relevant lifecycle events
	EventHistory []custom.Reference[Provenance] `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

type MedicationDispensePerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Who performed the dispense and what they did
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Individual who was performing
	Actor custom.Reference[MedicationDispensePerformerActor] `bson:"actor" json:"actor"`
}

type MedicationDispenseSubstitution struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Whether a substitution was or was not performed on the dispense
	WasSubstituted bool `bson:"wasSubstituted" json:"wasSubstituted"`
	// Code signifying whether a different drug was dispensed from what was prescribed
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Why was substitution made
	Reason []CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	// Who is responsible for the substitution
	ResponsibleParty *custom.Reference[MedicationDispenseSubstitutionResponsibleParty] `bson:"responsibleParty,omitempty" json:"responsibleParty,omitempty"`
}

type MedicationDispensePartOf interface {
	gofhirclient.Resource

	Is_MedicationDispensePartOf()
}

func (p Procedure) Is_MedicationDispensePartOf()                {}
func (m MedicationAdministration) Is_MedicationDispensePartOf() {}

type MedicationDispenseSubject interface {
	gofhirclient.Resource

	Is_MedicationDispenseSubject()
}

func (p Patient) Is_MedicationDispenseSubject() {}
func (g Group) Is_MedicationDispenseSubject()   {}

type MedicationDispensePerformerActor interface {
	gofhirclient.Resource

	Is_MedicationDispensePerformerActor()
}

func (p Practitioner) Is_MedicationDispensePerformerActor()     {}
func (p PractitionerRole) Is_MedicationDispensePerformerActor() {}
func (o Organization) Is_MedicationDispensePerformerActor()     {}
func (p Patient) Is_MedicationDispensePerformerActor()          {}
func (d Device) Is_MedicationDispensePerformerActor()           {}
func (r RelatedPerson) Is_MedicationDispensePerformerActor()    {}
func (c CareTeam) Is_MedicationDispensePerformerActor()         {}

type MedicationDispenseReceiver interface {
	gofhirclient.Resource

	Is_MedicationDispenseReceiver()
}

func (p Patient) Is_MedicationDispenseReceiver()          {}
func (p Practitioner) Is_MedicationDispenseReceiver()     {}
func (r RelatedPerson) Is_MedicationDispenseReceiver()    {}
func (l Location) Is_MedicationDispenseReceiver()         {}
func (p PractitionerRole) Is_MedicationDispenseReceiver() {}

type MedicationDispenseSubstitutionResponsibleParty interface {
	gofhirclient.Resource

	Is_MedicationDispenseSubstitutionResponsibleParty()
}

func (p Practitioner) Is_MedicationDispenseSubstitutionResponsibleParty()     {}
func (p PractitionerRole) Is_MedicationDispenseSubstitutionResponsibleParty() {}
func (o Organization) Is_MedicationDispenseSubstitutionResponsibleParty()     {}

func (m MedicationDispense) ResourceType() string {
	return "StructureDefinition"
}
