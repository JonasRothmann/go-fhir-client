// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/familymemberhistory-genetic
type FamilyMemberHistoryForGeneticsAnalysis struct {
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
	// Extension
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// External Id(s) for this record
	Identifier []Identifier `json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[FamilyMemberHistoryForGeneticsAnalysisInstantiatesCanonical] `json:"instantiatesCanonical,omitempty"`
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
	Participant []FamilyMemberHistoryForGeneticsAnalysisParticipant `json:"participant,omitempty"`
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
	Reason []custom.CodeableReference[FamilyMemberHistoryForGeneticsAnalysisReason] `json:"reason,omitempty"`
	// General note about related person
	Note []Annotation `json:"note,omitempty"`
	// Condition that the related person had
	Condition []FamilyMemberHistoryForGeneticsAnalysisCondition `json:"condition,omitempty"`
	// Procedures that the related person had
	Procedure []FamilyMemberHistoryForGeneticsAnalysisProcedure `json:"procedure,omitempty"`
}

type FamilyMemberHistoryForGeneticsAnalysisParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `json:"function,omitempty"`
	// Who or what participated in the activities related to the family member history
	Actor custom.Reference[FamilyMemberHistoryForGeneticsAnalysisParticipantActor] `json:"actor"`
}

type FamilyMemberHistoryForGeneticsAnalysisCondition struct {
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

type FamilyMemberHistoryForGeneticsAnalysisProcedure struct {
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

type OtherFamilyMemberHistoryForGeneticsAnalysis FamilyMemberHistoryForGeneticsAnalysis

type FamilyMemberHistoryForGeneticsAnalysisInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_FamilyMemberHistoryForGeneticsAnalysisInstantiatesCanonical()
}

func (p PlanDefinition) Is_FamilyMemberHistoryForGeneticsAnalysisInstantiatesCanonical()      {}
func (q Questionnaire) Is_FamilyMemberHistoryForGeneticsAnalysisInstantiatesCanonical()       {}
func (a ActivityDefinition) Is_FamilyMemberHistoryForGeneticsAnalysisInstantiatesCanonical()  {}
func (m Measure) Is_FamilyMemberHistoryForGeneticsAnalysisInstantiatesCanonical()             {}
func (o OperationDefinition) Is_FamilyMemberHistoryForGeneticsAnalysisInstantiatesCanonical() {}

type FamilyMemberHistoryForGeneticsAnalysisParticipantActor interface {
	gofhirclient.Resource

	Is_FamilyMemberHistoryForGeneticsAnalysisParticipantActor()
}

func (p Practitioner) Is_FamilyMemberHistoryForGeneticsAnalysisParticipantActor()     {}
func (p PractitionerRole) Is_FamilyMemberHistoryForGeneticsAnalysisParticipantActor() {}
func (p Patient) Is_FamilyMemberHistoryForGeneticsAnalysisParticipantActor()          {}
func (r RelatedPerson) Is_FamilyMemberHistoryForGeneticsAnalysisParticipantActor()    {}
func (d Device) Is_FamilyMemberHistoryForGeneticsAnalysisParticipantActor()           {}
func (o Organization) Is_FamilyMemberHistoryForGeneticsAnalysisParticipantActor()     {}
func (c CareTeam) Is_FamilyMemberHistoryForGeneticsAnalysisParticipantActor()         {}

type FamilyMemberHistoryForGeneticsAnalysisReason interface {
	gofhirclient.Resource

	Is_FamilyMemberHistoryForGeneticsAnalysisReason()
}

func (c Condition) Is_FamilyMemberHistoryForGeneticsAnalysisReason()             {}
func (o Observation) Is_FamilyMemberHistoryForGeneticsAnalysisReason()           {}
func (a AllergyIntolerance) Is_FamilyMemberHistoryForGeneticsAnalysisReason()    {}
func (q QuestionnaireResponse) Is_FamilyMemberHistoryForGeneticsAnalysisReason() {}
func (d DiagnosticReport) Is_FamilyMemberHistoryForGeneticsAnalysisReason()      {}
func (d DocumentReference) Is_FamilyMemberHistoryForGeneticsAnalysisReason()     {}

func (f FamilyMemberHistoryForGeneticsAnalysis) ResourceType() string {
	return "FamilyMemberHistoryForGeneticsAnalysis"
}

func (f FamilyMemberHistoryForGeneticsAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFamilyMemberHistoryForGeneticsAnalysis
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherFamilyMemberHistoryForGeneticsAnalysis: OtherFamilyMemberHistoryForGeneticsAnalysis(f), ResourceType: f.ResourceType()})
}

func UnmarshalFamilyMemberHistoryForGeneticsAnalysis(b []byte) (res FamilyMemberHistoryForGeneticsAnalysis, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
