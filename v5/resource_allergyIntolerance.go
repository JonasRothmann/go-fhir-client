// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
type AllergyIntoleranceReaction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Specific substance or pharmaceutical product considered to be responsible for event
	Substance *CodeableConcept `bson:"substance,omitempty" json:"substance,omitempty"`
	// Clinical symptoms/signs associated with the Event
	Manifestation []custom.CodeableReference[Observation] `bson:"manifestation" json:"manifestation"`
	// Description of the event as a whole
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Date(/time) when manifestations showed
	Onset *custom.DateTime `bson:"onset,omitempty" json:"onset,omitempty"`
	// mild | moderate | severe (of event as a whole)
	Severity *custom.Code `bson:"severity,omitempty" json:"severity,omitempty"`
	// How the subject was exposed to the substance
	ExposureRoute *CodeableConcept `bson:"exposureRoute,omitempty" json:"exposureRoute,omitempty"`
	// Text about event not captured in other fields
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type AllergyIntolerance struct {
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
	// External ids for this item
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | inactive | resolved
	ClinicalStatus *CodeableConcept `bson:"clinicalStatus,omitempty" json:"clinicalStatus,omitempty"`
	// unconfirmed | presumed | confirmed | refuted | entered-in-error
	VerificationStatus *CodeableConcept `bson:"verificationStatus,omitempty" json:"verificationStatus,omitempty"`
	// allergy | intolerance - Underlying mechanism (if known)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// food | medication | environment | biologic
	Category []custom.Code `bson:"category,omitempty" json:"category,omitempty"`
	// low | high | unable-to-assess
	Criticality *custom.Code `bson:"criticality,omitempty" json:"criticality,omitempty"`
	// Code that identifies the allergy or intolerance
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Who the allergy or intolerance is for
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Encounter when the allergy or intolerance was asserted
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When allergy or intolerance was identified
	Onset *custom.DateTime `bson:"onset,omitempty" json:"onset,omitempty"`
	// Date allergy or intolerance was first recorded
	RecordedDate *custom.DateTime `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	// Who or what participated in the activities related to the allergy or intolerance and how they were involved
	Participant []AllergyIntoleranceParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// Date(/time) of last known occurrence of a reaction
	LastOccurrence *custom.DateTime `bson:"lastOccurrence,omitempty" json:"lastOccurrence,omitempty"`
	// Additional text not captured in other fields
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Adverse Reaction Events linked to exposure to substance
	Reaction []AllergyIntoleranceReaction `bson:"reaction,omitempty" json:"reaction,omitempty"`
}

type AllergyIntoleranceParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who or what participated in the activities related to the allergy or intolerance
	Actor custom.Reference[AllergyIntoleranceParticipantActor] `bson:"actor" json:"actor"`
}

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
	return "StructureDefinition"
}
