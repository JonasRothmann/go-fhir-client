// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistoryProcedure struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Procedures performed on the related person
	Code CodeableConcept `json:"code"`
	// What happened following the procedure
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	// Whether the procedure contributed to the cause of death
	ContributedToDeath *bool `json:"contributedToDeath,omitempty"`
	// When the procedure was performed
	PerformedAge *Age `json:"performedAge,omitempty"`
	// When the procedure was performed
	PerformedRange *Range `json:"performedRange,omitempty"`
	// When the procedure was performed
	PerformedPeriod *Period `json:"performedPeriod,omitempty"`
	// When the procedure was performed
	PerformedString *string `json:"performedString,omitempty"`
	// When the procedure was performed
	PerformedDateTime *custom.DateTime `json:"performedDateTime,omitempty"`
	// Extra information about the procedure
	Note []Annotation `json:"note,omitempty"`
}

type FamilyMemberHistory struct {
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
	// External Id(s) for this record
	Identifier []Identifier `json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[FamilyMemberHistoryInstantiatesCanonical] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `json:"instantiatesUri,omitempty"`
	// partial | completed | entered-in-error | health-unknown
	Status custom.Code `json:"status"`
	// subject-unknown | withheld | unable-to-obtain | deferred
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	// Patient history is about
	Patient custom.Reference[Patient] `json:"patient"`
	// When history was recorded or last updated
	Date *custom.DateTime `json:"date,omitempty"`
	// Who or what participated in the activities related to the family member history and how they were involved
	Participant []FamilyMemberHistoryParticipant `json:"participant,omitempty"`
	// The family member described
	Name *string `json:"name,omitempty"`
	// Relationship to the subject
	Relationship CodeableConcept `json:"relationship"`
	// male | female | other | unknown
	Sex *CodeableConcept `json:"sex,omitempty"`
	// (approximate) date of birth
	BornPeriod *Period `json:"bornPeriod,omitempty"`
	// (approximate) date of birth
	BornDate *custom.Date `json:"bornDate,omitempty"`
	// (approximate) date of birth
	BornString *string `json:"bornString,omitempty"`
	// (approximate) age
	AgeAge *Age `json:"ageAge,omitempty"`
	// (approximate) age
	AgeRange *Range `json:"ageRange,omitempty"`
	// (approximate) age
	AgeString *string `json:"ageString,omitempty"`
	// Age is estimated?
	EstimatedAge *bool `json:"estimatedAge,omitempty"`
	// Dead? How old/when?
	DeceasedBoolean *bool `json:"deceasedBoolean,omitempty"`
	// Dead? How old/when?
	DeceasedAge *Age `json:"deceasedAge,omitempty"`
	// Dead? How old/when?
	DeceasedRange *Range `json:"deceasedRange,omitempty"`
	// Dead? How old/when?
	DeceasedDate *custom.Date `json:"deceasedDate,omitempty"`
	// Dead? How old/when?
	DeceasedString *string `json:"deceasedString,omitempty"`
	// Why was family member history performed?
	Reason []custom.CodeableReference[FamilyMemberHistoryReason] `json:"reason,omitempty"`
	// General note about related person
	Note []Annotation `json:"note,omitempty"`
	// Condition that the related person had
	Condition []FamilyMemberHistoryCondition `json:"condition,omitempty"`
	// Procedures that the related person had
	Procedure []FamilyMemberHistoryProcedure `json:"procedure,omitempty"`
}

type FamilyMemberHistoryParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `json:"function,omitempty"`
	// Who or what participated in the activities related to the family member history
	Actor custom.Reference[FamilyMemberHistoryParticipantActor] `json:"actor"`
}

type FamilyMemberHistoryCondition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Condition suffered by relation
	Code CodeableConcept `json:"code"`
	// deceased | permanent disability | etc
	Outcome *CodeableConcept `json:"outcome,omitempty"`
	// Whether the condition contributed to the cause of death
	ContributedToDeath *bool `json:"contributedToDeath,omitempty"`
	// When condition first manifested
	OnsetAge *Age `json:"onsetAge,omitempty"`
	// When condition first manifested
	OnsetRange *Range `json:"onsetRange,omitempty"`
	// When condition first manifested
	OnsetPeriod *Period `json:"onsetPeriod,omitempty"`
	// When condition first manifested
	OnsetString *string `json:"onsetString,omitempty"`
	// Extra information about condition
	Note []Annotation `json:"note,omitempty"`
}

type OtherFamilyMemberHistory FamilyMemberHistory

type FamilyMemberHistoryInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_FamilyMemberHistoryInstantiatesCanonical()
}

func (p PlanDefinition) Is_FamilyMemberHistoryInstantiatesCanonical()      {}
func (q Questionnaire) Is_FamilyMemberHistoryInstantiatesCanonical()       {}
func (a ActivityDefinition) Is_FamilyMemberHistoryInstantiatesCanonical()  {}
func (m Measure) Is_FamilyMemberHistoryInstantiatesCanonical()             {}
func (o OperationDefinition) Is_FamilyMemberHistoryInstantiatesCanonical() {}

type FamilyMemberHistoryParticipantActor interface {
	gofhirclient.Resource

	Is_FamilyMemberHistoryParticipantActor()
}

func (p Practitioner) Is_FamilyMemberHistoryParticipantActor()     {}
func (p PractitionerRole) Is_FamilyMemberHistoryParticipantActor() {}
func (p Patient) Is_FamilyMemberHistoryParticipantActor()          {}
func (r RelatedPerson) Is_FamilyMemberHistoryParticipantActor()    {}
func (d Device) Is_FamilyMemberHistoryParticipantActor()           {}
func (o Organization) Is_FamilyMemberHistoryParticipantActor()     {}
func (c CareTeam) Is_FamilyMemberHistoryParticipantActor()         {}

type FamilyMemberHistoryReason interface {
	gofhirclient.Resource

	Is_FamilyMemberHistoryReason()
}

func (c Condition) Is_FamilyMemberHistoryReason()             {}
func (o Observation) Is_FamilyMemberHistoryReason()           {}
func (a AllergyIntolerance) Is_FamilyMemberHistoryReason()    {}
func (q QuestionnaireResponse) Is_FamilyMemberHistoryReason() {}
func (d DiagnosticReport) Is_FamilyMemberHistoryReason()      {}
func (d DocumentReference) Is_FamilyMemberHistoryReason()     {}

func (f FamilyMemberHistory) ResourceType() string {
	return "FamilyMemberHistory"
}

func (f FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFamilyMemberHistory
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherFamilyMemberHistory: OtherFamilyMemberHistory(f), ResourceType: f.ResourceType()})
}

func UnmarshalFamilyMemberHistory(b []byte) (res FamilyMemberHistory, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
