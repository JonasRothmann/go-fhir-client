// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Condition
type ConditionStage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Simple summary (disease specific)
	Summary *CodeableConcept `json:"summary,omitempty"`
	// Formal record of assessment
	Assessment []custom.Reference[ConditionStageAssessment] `json:"assessment,omitempty"`
	// Kind of staging
	Type *CodeableConcept `json:"type,omitempty"`
}

type Condition struct {
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
	// External Ids for this condition
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | recurrence | relapse | inactive | remission | resolved | unknown
	ClinicalStatus CodeableConcept `json:"clinicalStatus"`
	// unconfirmed | provisional | differential | confirmed | refuted | entered-in-error
	VerificationStatus *CodeableConcept `json:"verificationStatus,omitempty"`
	// problem-list-item | encounter-diagnosis
	Category []CodeableConcept `json:"category,omitempty"`
	// Subjective severity of condition
	Severity *CodeableConcept `json:"severity,omitempty"`
	// Identification of the condition, problem or diagnosis
	Code *CodeableConcept `json:"code,omitempty"`
	// Anatomical location, if relevant
	BodySite []CodeableConcept `json:"bodySite,omitempty"`
	// Who has the condition?
	Subject custom.Reference[ConditionSubject] `json:"subject"`
	// The Encounter during which this Condition was created
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Estimated or actual date,  date-time, or age
	OnsetDateTime *custom.DateTime `json:"onsetDateTime,omitempty"`
	// Estimated or actual date,  date-time, or age
	OnsetAge *Age `json:"onsetAge,omitempty"`
	// Estimated or actual date,  date-time, or age
	OnsetPeriod *Period `json:"onsetPeriod,omitempty"`
	// Estimated or actual date,  date-time, or age
	OnsetRange *Range `json:"onsetRange,omitempty"`
	// Estimated or actual date,  date-time, or age
	OnsetString *string `json:"onsetString,omitempty"`
	// When in resolution/remission
	AbatementDateTime *custom.DateTime `json:"abatementDateTime,omitempty"`
	// When in resolution/remission
	AbatementAge *Age `json:"abatementAge,omitempty"`
	// When in resolution/remission
	AbatementPeriod *Period `json:"abatementPeriod,omitempty"`
	// When in resolution/remission
	AbatementRange *Range `json:"abatementRange,omitempty"`
	// When in resolution/remission
	AbatementString *string `json:"abatementString,omitempty"`
	// Date condition was first recorded
	RecordedDate *custom.DateTime `json:"recordedDate,omitempty"`
	// Who or what participated in the activities related to the condition and how they were involved
	Participant []ConditionParticipant `json:"participant,omitempty"`
	// Stage/grade, usually assessed formally
	Stage []ConditionStage `json:"stage,omitempty"`
	// Supporting evidence for the verification status
	Evidence []custom.CodeableReference[Resource] `json:"evidence,omitempty"`
	// Additional information about the Condition
	Note []Annotation `json:"note,omitempty"`
}

type ConditionParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `json:"function,omitempty"`
	// Who or what participated in the activities related to the condition
	Actor custom.Reference[ConditionParticipantActor] `json:"actor"`
}

type OtherCondition Condition

type ConditionSubject interface {
	gofhirclient.Resource

	Is_ConditionSubject()
}

func (p Patient) Is_ConditionSubject() {}
func (g Group) Is_ConditionSubject()   {}

type ConditionParticipantActor interface {
	gofhirclient.Resource

	Is_ConditionParticipantActor()
}

func (p Practitioner) Is_ConditionParticipantActor()     {}
func (p PractitionerRole) Is_ConditionParticipantActor() {}
func (p Patient) Is_ConditionParticipantActor()          {}
func (r RelatedPerson) Is_ConditionParticipantActor()    {}
func (d Device) Is_ConditionParticipantActor()           {}
func (o Organization) Is_ConditionParticipantActor()     {}
func (c CareTeam) Is_ConditionParticipantActor()         {}

type ConditionStageAssessment interface {
	gofhirclient.Resource

	Is_ConditionStageAssessment()
}

func (c ClinicalImpression) Is_ConditionStageAssessment() {}
func (d DiagnosticReport) Is_ConditionStageAssessment()   {}
func (o Observation) Is_ConditionStageAssessment()        {}

func (c Condition) ResourceType() string {
	return "Condition"
}

func (c Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCondition: OtherCondition(c), ResourceType: c.ResourceType()})
}

func UnmarshalCondition(b []byte) (res Condition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
