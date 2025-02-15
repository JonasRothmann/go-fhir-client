// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistoryProcedure struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Procedures performed on the related person
	Code CodeableConcept `bson:"code" json:"code"`
	// What happened following the procedure
	Outcome *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	// Whether the procedure contributed to the cause of death
	ContributedToDeath *bool `bson:"contributedToDeath,omitempty" json:"contributedToDeath,omitempty"`
	// When the procedure was performed
	Performed *Age `bson:"performed,omitempty" json:"performed,omitempty"`
	// Extra information about the procedure
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type FamilyMemberHistory struct {
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
	// External Id(s) for this record
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[FamilyMemberHistoryInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// partial | completed | entered-in-error | health-unknown
	Status custom.Code `bson:"status" json:"status"`
	// subject-unknown | withheld | unable-to-obtain | deferred
	DataAbsentReason *CodeableConcept `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	// Patient history is about
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// When history was recorded or last updated
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Who or what participated in the activities related to the family member history and how they were involved
	Participant []FamilyMemberHistoryParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// The family member described
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Relationship to the subject
	Relationship CodeableConcept `bson:"relationship" json:"relationship"`
	// male | female | other | unknown
	Sex *CodeableConcept `bson:"sex,omitempty" json:"sex,omitempty"`
	// (approximate) date of birth
	Born *Period `bson:"born,omitempty" json:"born,omitempty"`
	// (approximate) age
	Age *Age `bson:"age,omitempty" json:"age,omitempty"`
	// Age is estimated?
	EstimatedAge *bool `bson:"estimatedAge,omitempty" json:"estimatedAge,omitempty"`
	// Dead? How old/when?
	Deceased *bool `bson:"deceased,omitempty" json:"deceased,omitempty"`
	// Why was family member history performed?
	Reason []custom.CodeableReference[FamilyMemberHistoryReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// General note about related person
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Condition that the related person had
	Condition []FamilyMemberHistoryCondition `bson:"condition,omitempty" json:"condition,omitempty"`
	// Procedures that the related person had
	Procedure []FamilyMemberHistoryProcedure `bson:"procedure,omitempty" json:"procedure,omitempty"`
}

type FamilyMemberHistoryParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who or what participated in the activities related to the family member history
	Actor custom.Reference[FamilyMemberHistoryParticipantActor] `bson:"actor" json:"actor"`
}

type FamilyMemberHistoryCondition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Condition suffered by relation
	Code CodeableConcept `bson:"code" json:"code"`
	// deceased | permanent disability | etc
	Outcome *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	// Whether the condition contributed to the cause of death
	ContributedToDeath *bool `bson:"contributedToDeath,omitempty" json:"contributedToDeath,omitempty"`
	// When condition first manifested
	Onset *Age `bson:"onset,omitempty" json:"onset,omitempty"`
	// Extra information about condition
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

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
	return "StructureDefinition"
}
