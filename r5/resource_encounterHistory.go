// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EncounterHistory
type EncounterHistory struct {
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
	// The Encounter associated with this set of historic values
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Identifier(s) by which this encounter is known
	Identifier []Identifier `json:"identifier,omitempty"`
	// planned | in-progress | on-hold | discharged | completed | cancelled | discontinued | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Classification of patient encounter
	Class CodeableConcept `json:"class"`
	// Specific type of encounter
	Type []CodeableConcept `json:"type,omitempty"`
	// Specific type of service
	ServiceType []custom.CodeableReference[HealthcareService] `json:"serviceType,omitempty"`
	// The patient or group related to this encounter
	Subject *custom.Reference[EncounterHistorySubject] `json:"subject,omitempty"`
	// The current status of the subject in relation to the Encounter
	SubjectStatus *CodeableConcept `json:"subjectStatus,omitempty"`
	// The actual start and end time associated with this set of values associated with the encounter
	ActualPeriod *Period `json:"actualPeriod,omitempty"`
	// The planned start date/time (or admission date) of the encounter
	PlannedStartDate *custom.DateTime `json:"plannedStartDate,omitempty"`
	// The planned end date/time (or discharge date) of the encounter
	PlannedEndDate *custom.DateTime `json:"plannedEndDate,omitempty"`
	// Actual quantity of time the encounter lasted (less time absent)
	Length *Duration `json:"length,omitempty"`
	// Location of the patient at this point in the encounter
	Location []EncounterHistoryLocation `json:"location,omitempty"`
}

type EncounterHistoryLocation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Location the encounter takes place
	Location custom.Reference[Location] `json:"location"`
	// The physical type of the location (usually the level in the location hierarchy - bed, room, ward, virtual etc.)
	Form *CodeableConcept `json:"form,omitempty"`
}

type OtherEncounterHistory EncounterHistory

type EncounterHistorySubject interface {
	gofhirclient.Resource

	Is_EncounterHistorySubject()
}

func (p Patient) Is_EncounterHistorySubject() {}
func (g Group) Is_EncounterHistorySubject()   {}

func (e EncounterHistory) ResourceType() string {
	return "EncounterHistory"
}

func (e EncounterHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEncounterHistory
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherEncounterHistory: OtherEncounterHistory(e), ResourceType: e.ResourceType()})
}

func UnmarshalEncounterHistory(b []byte) (res EncounterHistory, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
