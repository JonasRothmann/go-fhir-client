// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicationDispense
type MedicationDispensePerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Who performed the dispense and what they did
	Function *CodeableConcept `json:"function,omitempty"`
	// Individual who was performing
	Actor custom.Reference[MedicationDispensePerformerActor] `json:"actor"`
}

type MedicationDispenseSubstitution struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Whether a substitution was or was not performed on the dispense
	WasSubstituted bool `json:"wasSubstituted"`
	// Code signifying whether a different drug was dispensed from what was prescribed
	Type *CodeableConcept `json:"type,omitempty"`
	// Why was substitution made
	Reason []CodeableConcept `json:"reason,omitempty"`
	// Who is responsible for the substitution
	ResponsibleParty *custom.Reference[MedicationDispenseSubstitutionResponsibleParty] `json:"responsibleParty,omitempty"`
}

type MedicationDispense struct {
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
	// Plan that is fulfilled by this dispense
	BasedOn []custom.Reference[CarePlan] `json:"basedOn,omitempty"`
	// Event that dispense is part of
	PartOf []custom.Reference[MedicationDispensePartOf] `json:"partOf,omitempty"`
	// preparation | in-progress | cancelled | on-hold | completed | entered-in-error | stopped | declined | unknown
	Status custom.Code `json:"status"`
	// Why a dispense was not performed
	NotPerformedReason *custom.CodeableReference[DetectedIssue] `json:"notPerformedReason,omitempty"`
	// When the status changed
	StatusChanged *custom.DateTime `json:"statusChanged,omitempty"`
	// Type of medication dispense
	Category []CodeableConcept `json:"category,omitempty"`
	// What medication was supplied
	Medication custom.CodeableReference[Medication] `json:"medication"`
	// Who the dispense is for
	Subject custom.Reference[MedicationDispenseSubject] `json:"subject"`
	// Encounter associated with event
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Information that supports the dispensing of the medication
	SupportingInformation []custom.Reference[Resource] `json:"supportingInformation,omitempty"`
	// Who performed event
	Performer []MedicationDispensePerformer `json:"performer,omitempty"`
	// Where the dispense occurred
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Medication order that authorizes the dispense
	AuthorizingPrescription []custom.Reference[MedicationRequest] `json:"authorizingPrescription,omitempty"`
	// Trial fill, partial fill, emergency fill, etc
	Type *CodeableConcept `json:"type,omitempty"`
	// Amount dispensed
	Quantity *Quantity `json:"quantity,omitempty"`
	// Amount of medication expressed as a timing amount
	DaysSupply *Quantity `json:"daysSupply,omitempty"`
	// When the recording of the dispense started
	Recorded *custom.DateTime `json:"recorded,omitempty"`
	// When product was packaged and reviewed
	WhenPrepared *custom.DateTime `json:"whenPrepared,omitempty"`
	// When product was given out
	WhenHandedOver *custom.DateTime `json:"whenHandedOver,omitempty"`
	// Where the medication was/will be sent
	Destination *custom.Reference[Location] `json:"destination,omitempty"`
	// Who collected the medication or where the medication was delivered
	Receiver []custom.Reference[MedicationDispenseReceiver] `json:"receiver,omitempty"`
	// Information about the dispense
	Note []Annotation `json:"note,omitempty"`
	// Full representation of the dosage instructions
	RenderedDosageInstruction *custom.Markdown `json:"renderedDosageInstruction,omitempty"`
	// How the medication is to be used by the patient or administered by the caregiver
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`
	// Whether a substitution was performed on the dispense
	Substitution *MedicationDispenseSubstitution `json:"substitution,omitempty"`
	// A list of relevant lifecycle events
	EventHistory []custom.Reference[Provenance] `json:"eventHistory,omitempty"`
}

type OtherMedicationDispense MedicationDispense

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
	return "MedicationDispense"
}

func (m MedicationDispense) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationDispense
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMedicationDispense: OtherMedicationDispense(m), ResourceType: m.ResourceType()})
}

func UnmarshalMedicationDispense(b []byte) (res MedicationDispense, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
