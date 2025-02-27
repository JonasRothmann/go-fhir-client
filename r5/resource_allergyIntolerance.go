// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
type AllergyIntolerance struct {
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
	// External ids for this item
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | inactive | resolved
	ClinicalStatus *CodeableConcept `json:"clinicalStatus,omitempty"`
	// unconfirmed | presumed | confirmed | refuted | entered-in-error
	VerificationStatus *CodeableConcept `json:"verificationStatus,omitempty"`
	// allergy | intolerance - Underlying mechanism (if known)
	Type *CodeableConcept `json:"type,omitempty"`
	// food | medication | environment | biologic
	Category []custom.Code `json:"category,omitempty"`
	// low | high | unable-to-assess
	Criticality *custom.Code `json:"criticality,omitempty"`
	// Code that identifies the allergy or intolerance
	Code *CodeableConcept `json:"code,omitempty"`
	// Who the allergy or intolerance is for
	Patient custom.Reference[Patient] `json:"patient"`
	// Encounter when the allergy or intolerance was asserted
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When allergy or intolerance was identified
	OnsetDateTime *custom.DateTime `json:"onsetDateTime,omitempty"`
	// When allergy or intolerance was identified
	OnsetAge *Age `json:"onsetAge,omitempty"`
	// When allergy or intolerance was identified
	OnsetPeriod *Period `json:"onsetPeriod,omitempty"`
	// When allergy or intolerance was identified
	OnsetRange *Range `json:"onsetRange,omitempty"`
	// When allergy or intolerance was identified
	OnsetString *string `json:"onsetString,omitempty"`
	// Date allergy or intolerance was first recorded
	RecordedDate *custom.DateTime `json:"recordedDate,omitempty"`
	// Who or what participated in the activities related to the allergy or intolerance and how they were involved
	Participant []AllergyIntoleranceParticipant `json:"participant,omitempty"`
	// Date(/time) of last known occurrence of a reaction
	LastOccurrence *custom.DateTime `json:"lastOccurrence,omitempty"`
	// Additional text not captured in other fields
	Note []Annotation `json:"note,omitempty"`
	// Adverse Reaction Events linked to exposure to substance
	Reaction []AllergyIntoleranceReaction `json:"reaction,omitempty"`
}

type AllergyIntoleranceParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `json:"function,omitempty"`
	// Who or what participated in the activities related to the allergy or intolerance
	Actor custom.Reference[AllergyIntoleranceParticipantActor] `json:"actor"`
}

type AllergyIntoleranceReaction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Specific substance or pharmaceutical product considered to be responsible for event
	Substance *CodeableConcept `json:"substance,omitempty"`
	// Clinical symptoms/signs associated with the Event
	Manifestation []custom.CodeableReference[Observation] `json:"manifestation"`
	// Description of the event as a whole
	Description *string `json:"description,omitempty"`
	// Date(/time) when manifestations showed
	Onset *custom.DateTime `json:"onset,omitempty"`
	// mild | moderate | severe (of event as a whole)
	Severity *custom.Code `json:"severity,omitempty"`
	// How the subject was exposed to the substance
	ExposureRoute *CodeableConcept `json:"exposureRoute,omitempty"`
	// Text about event not captured in other fields
	Note []Annotation `json:"note,omitempty"`
}

type OtherAllergyIntolerance AllergyIntolerance

type AllergyIntoleranceParticipantActor interface {
	gofhirclient.Resource

	Is_AllergyIntoleranceParticipantActor()
}

func (p Practitioner) Is_AllergyIntoleranceParticipantActor()     {}
func (p PractitionerRole) Is_AllergyIntoleranceParticipantActor() {}
func (p Patient) Is_AllergyIntoleranceParticipantActor()          {}
func (r RelatedPerson) Is_AllergyIntoleranceParticipantActor()    {}
func (d Device) Is_AllergyIntoleranceParticipantActor()           {}
func (o Organization) Is_AllergyIntoleranceParticipantActor()     {}
func (c CareTeam) Is_AllergyIntoleranceParticipantActor()         {}

func (a AllergyIntolerance) ResourceType() string {
	return "AllergyIntolerance"
}

func (a AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAllergyIntolerance
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAllergyIntolerance: OtherAllergyIntolerance(a), ResourceType: a.ResourceType()})
}

func UnmarshalAllergyIntolerance(b []byte) (res AllergyIntolerance, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
