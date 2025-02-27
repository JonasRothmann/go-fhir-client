// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicationAdministration
type MedicationAdministration struct {
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
	// Plan this is fulfilled by this administration
	BasedOn []custom.Reference[CarePlan] `json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[MedicationAdministrationPartOf] `json:"partOf,omitempty"`
	// in-progress | not-done | on-hold | completed | entered-in-error | stopped | unknown
	Status custom.Code `json:"status"`
	// Reason administration not performed
	StatusReason []CodeableConcept `json:"statusReason,omitempty"`
	// Type of medication administration
	Category []CodeableConcept `json:"category,omitempty"`
	// What was administered
	Medication custom.CodeableReference[Medication] `json:"medication"`
	// Who received medication
	Subject custom.Reference[MedicationAdministrationSubject] `json:"subject"`
	// Encounter administered as part of
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Additional information to support administration
	SupportingInformation []custom.Reference[Resource] `json:"supportingInformation,omitempty"`
	// Specific date/time or interval of time during which the administration took place (or did not take place)
	OccurenceDateTime *custom.DateTime `json:"occurenceDateTime,omitempty"`
	// Specific date/time or interval of time during which the administration took place (or did not take place)
	OccurencePeriod *Period `json:"occurencePeriod,omitempty"`
	// Specific date/time or interval of time during which the administration took place (or did not take place)
	OccurenceTiming *Timing `json:"occurenceTiming,omitempty"`
	// When the MedicationAdministration was first captured in the subject's record
	Recorded *custom.DateTime `json:"recorded,omitempty"`
	// Full dose was not administered
	IsSubPotent *bool `json:"isSubPotent,omitempty"`
	// Reason full dose was not administered
	SubPotentReason []CodeableConcept `json:"subPotentReason,omitempty"`
	// Who or what performed the medication administration and what type of performance they did
	Performer []MedicationAdministrationPerformer `json:"performer,omitempty"`
	// Concept, condition or observation that supports why the medication was administered
	Reason []custom.CodeableReference[MedicationAdministrationReason] `json:"reason,omitempty"`
	// Request administration performed against
	Request *custom.Reference[MedicationRequest] `json:"request,omitempty"`
	// Device used to administer
	Device []custom.CodeableReference[Device] `json:"device,omitempty"`
	// Information about the administration
	Note []Annotation `json:"note,omitempty"`
	// Details of how medication was taken
	Dosage *MedicationAdministrationDosage `json:"dosage,omitempty"`
	// A list of events of interest in the lifecycle
	EventHistory []custom.Reference[Provenance] `json:"eventHistory,omitempty"`
}

type MedicationAdministrationPerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of performance
	Function *CodeableConcept `json:"function,omitempty"`
	// Who or what performed the medication administration
	Actor custom.CodeableReference[MedicationAdministrationPerformerActor] `json:"actor"`
}

type MedicationAdministrationDosage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Free text dosage instructions e.g. SIG
	Text *string `json:"text,omitempty"`
	// Body site administered to
	Site *CodeableConcept `json:"site,omitempty"`
	// Path of substance into body
	Route *CodeableConcept `json:"route,omitempty"`
	// How drug was administered
	Method *CodeableConcept `json:"method,omitempty"`
	// Amount of medication per dose
	Dose *Quantity `json:"dose,omitempty"`
	// Dose quantity per unit of time
	RateRatio *Ratio `json:"rateRatio,omitempty"`
	// Dose quantity per unit of time
	RateQuantity *Quantity `json:"rateQuantity,omitempty"`
}

type OtherMedicationAdministration MedicationAdministration

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
	return "MedicationAdministration"
}

func (m MedicationAdministration) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationAdministration
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMedicationAdministration: OtherMedicationAdministration(m), ResourceType: m.ResourceType()})
}

func UnmarshalMedicationAdministration(b []byte) (res MedicationAdministration, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
