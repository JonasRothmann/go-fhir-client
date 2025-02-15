// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Condition
type Condition struct {
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
	// External Ids for this condition
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | recurrence | relapse | inactive | remission | resolved | unknown
	ClinicalStatus CodeableConcept `bson:"clinicalStatus" json:"clinicalStatus"`
	// unconfirmed | provisional | differential | confirmed | refuted | entered-in-error
	VerificationStatus *CodeableConcept `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	// problem-list-item | encounter-diagnosis
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Subjective severity of condition
	Severity *CodeableConcept `bson:"severity,omitempty" json:"severity,omitempty"`
	// Identification of the condition, problem or diagnosis
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Anatomical location, if relevant
	BodySite []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Who has the condition?
	Subject custom.Reference[ConditionSubject] `bson:"subject" json:"subject"`
	// The Encounter during which this Condition was created
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Estimated or actual date,  date-time, or age
	Onset *custom.DateTime `bson:"onset,omitempty" json:"onset,omitempty"`
	// When in resolution/remission
	Abatement *custom.DateTime `bson:"abatement,omitempty" json:"abatement,omitempty"`
	// Date condition was first recorded
	RecordedDate *custom.DateTime `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	// Who or what participated in the activities related to the condition and how they were involved
	Participant []ConditionParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// Stage/grade, usually assessed formally
	Stage []ConditionStage `bson:"stage,omitempty" json:"stage,omitempty"`
	// Supporting evidence for the verification status
	Evidence []custom.CodeableReference[Resource] `bson:"evidence,omitempty" json:"evidence,omitempty"`
	// Additional information about the Condition
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type ConditionParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who or what participated in the activities related to the condition
	Actor custom.Reference[ConditionParticipantActor] `bson:"actor" json:"actor"`
}

type ConditionStage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Simple summary (disease specific)
	Summary *CodeableConcept `bson:"summary,omitempty" json:"summary,omitempty"`
	// Formal record of assessment
	Assessment []custom.Reference[ConditionStageAssessment] `bson:"assessment,omitempty" json:"assessment,omitempty"`
	// Kind of staging
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
}

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
	return "StructureDefinition"
}
