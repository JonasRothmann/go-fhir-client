// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EncounterHistory
type EncounterHistory struct {
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
	// The Encounter associated with this set of historic values
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Identifier(s) by which this encounter is known
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// planned | in-progress | on-hold | discharged | completed | cancelled | discontinued | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Classification of patient encounter
	Class CodeableConcept `bson:"class" json:"class"`
	// Specific type of encounter
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Specific type of service
	ServiceType []custom.CodeableReference[HealthcareService] `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	// The patient or group related to this encounter
	Subject *custom.Reference[EncounterHistorySubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// The current status of the subject in relation to the Encounter
	SubjectStatus *CodeableConcept `bson:"subjectStatus,omitempty" json:"subjectStatus,omitempty"`
	// The actual start and end time associated with this set of values associated with the encounter
	ActualPeriod *Period `bson:"actualPeriod,omitempty" json:"actualPeriod,omitempty"`
	// The planned start date/time (or admission date) of the encounter
	PlannedStartDate *custom.DateTime `bson:"plannedStartDate,omitempty" json:"plannedStartDate,omitempty"`
	// The planned end date/time (or discharge date) of the encounter
	PlannedEndDate *custom.DateTime `bson:"plannedEndDate,omitempty" json:"plannedEndDate,omitempty"`
	// Actual quantity of time the encounter lasted (less time absent)
	Length *Duration `bson:"length,omitempty" json:"length,omitempty"`
	// Location of the patient at this point in the encounter
	Location []EncounterHistoryLocation `bson:"location,omitempty" json:"location,omitempty"`
}

type EncounterHistoryLocation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Location the encounter takes place
	Location custom.Reference[Location] `bson:"location" json:"location"`
	// The physical type of the location (usually the level in the location hierarchy - bed, room, ward, virtual etc.)
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty"`
}

type EncounterHistorySubject interface {
	gofhirclient.Resource

	Is_EncounterHistorySubject()
}

func (p Patient) Is_EncounterHistorySubject() {}
func (g Group) Is_EncounterHistorySubject()   {}

func (e EncounterHistory) ResourceType() string {
	return "StructureDefinition"
}
