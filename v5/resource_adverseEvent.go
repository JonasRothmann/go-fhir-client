// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AdverseEvent
type AdverseEventMitigatingAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Ameliorating action taken after the adverse event occured in order to reduce the extent of harm
	Item custom.Reference[AdverseEventMitigatingActionItem] `bson:"item" json:"item"`
}

type AdverseEventSupportingInfo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Subject medical history or document relevant to this adverse event
	Item custom.Reference[AdverseEventSupportingInfoItem] `bson:"item" json:"item"`
}

type AdverseEvent struct {
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
	// Business identifier for the event
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// in-progress | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// actual | potential
	Actuality custom.Code `bson:"actuality" json:"actuality"`
	// wrong-patient | procedure-mishap | medication-mishap | device | unsafe-physical-environment | hospital-aquired-infection | wrong-body-site
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Event or incident that occurred or was averted
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Subject impacted by event
	Subject custom.Reference[AdverseEventSubject] `bson:"subject" json:"subject"`
	// The Encounter associated with the start of the AdverseEvent
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the event occurred
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// When the event was detected
	Detected *custom.DateTime `bson:"detected,omitempty" json:"detected,omitempty"`
	// When the event was recorded
	RecordedDate *custom.DateTime `bson:"recordedDate,omitempty" json:"recordedDate,omitempty"`
	// Effect on the subject due to this event
	ResultingEffect []custom.Reference[AdverseEventResultingEffect] `bson:"resultingEffect,omitempty" json:"resultingEffect,omitempty"`
	// Location where adverse event occurred
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Seriousness or gravity of the event
	Seriousness *CodeableConcept `bson:"seriousness,omitempty" json:"seriousness,omitempty"`
	// Type of outcome from the adverse event
	Outcome []CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	// Who recorded the adverse event
	Recorder *custom.Reference[AdverseEventRecorder] `bson:"recorder,omitempty" json:"recorder,omitempty"`
	// Who was involved in the adverse event or the potential adverse event and what they did
	Participant []AdverseEventParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// Research study that the subject is enrolled in
	Study []custom.Reference[ResearchStudy] `bson:"study,omitempty" json:"study,omitempty"`
	// Considered likely or probable or anticipated in the research study
	ExpectedInResearchStudy *bool `bson:"expectedInResearchStudy,omitempty" json:"expectedInResearchStudy,omitempty"`
	// The suspected agent causing the adverse event
	SuspectEntity []AdverseEventSuspectEntity `bson:"suspectEntity,omitempty" json:"suspectEntity,omitempty"`
	// Contributing factors suspected to have increased the probability or severity of the adverse event
	ContributingFactor []AdverseEventContributingFactor `bson:"contributingFactor,omitempty" json:"contributingFactor,omitempty"`
	// Preventive actions that contributed to avoiding the adverse event
	PreventiveAction []AdverseEventPreventiveAction `bson:"preventiveAction,omitempty" json:"preventiveAction,omitempty"`
	// Ameliorating actions taken after the adverse event occured in order to reduce the extent of harm
	MitigatingAction []AdverseEventMitigatingAction `bson:"mitigatingAction,omitempty" json:"mitigatingAction,omitempty"`
	// Supporting information relevant to the event
	SupportingInfo []AdverseEventSupportingInfo `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Comment on adverse event
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type AdverseEventParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who was involved in the adverse event or the potential adverse event
	Actor custom.Reference[AdverseEventParticipantActor] `bson:"actor" json:"actor"`
}

type AdverseEventSuspectEntity struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Refers to the specific entity that caused the adverse event
	Instance CodeableConcept `bson:"instance" json:"instance"`
	// Information on the possible cause of the event
	Causality *AdverseEventSuspectEntityCausality `bson:"causality,omitempty" json:"causality,omitempty"`
}

type AdverseEventSuspectEntityCausality struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Method of evaluating the relatedness of the suspected entity to the event
	AssessmentMethod *CodeableConcept `bson:"assessmentMethod,omitempty" json:"assessmentMethod,omitempty"`
	// Result of the assessment regarding the relatedness of the suspected entity to the event
	EntityRelatedness *CodeableConcept `bson:"entityRelatedness,omitempty" json:"entityRelatedness,omitempty"`
	// Author of the information on the possible cause of the event
	Author *custom.Reference[AdverseEventSuspectEntityCausalityAuthor] `bson:"author,omitempty" json:"author,omitempty"`
}

type AdverseEventContributingFactor struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Item suspected to have increased the probability or severity of the adverse event
	Item custom.Reference[AdverseEventContributingFactorItem] `bson:"item" json:"item"`
}

type AdverseEventPreventiveAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Action that contributed to avoiding the adverse event
	Item custom.Reference[AdverseEventPreventiveActionItem] `bson:"item" json:"item"`
}

type AdverseEventSubject interface {
	gofhirclient.Resource

	Is_AdverseEventSubject()
}

func (p Patient) Is_AdverseEventSubject()         {}
func (g Group) Is_AdverseEventSubject()           {}
func (p Practitioner) Is_AdverseEventSubject()    {}
func (r RelatedPerson) Is_AdverseEventSubject()   {}
func (r ResearchSubject) Is_AdverseEventSubject() {}

type AdverseEventResultingEffect interface {
	gofhirclient.Resource

	Is_AdverseEventResultingEffect()
}

func (c Condition) Is_AdverseEventResultingEffect()   {}
func (o Observation) Is_AdverseEventResultingEffect() {}

type AdverseEventRecorder interface {
	gofhirclient.Resource

	Is_AdverseEventRecorder()
}

func (p Patient) Is_AdverseEventRecorder()          {}
func (p Practitioner) Is_AdverseEventRecorder()     {}
func (p PractitionerRole) Is_AdverseEventRecorder() {}
func (r RelatedPerson) Is_AdverseEventRecorder()    {}
func (r ResearchSubject) Is_AdverseEventRecorder()  {}

type AdverseEventParticipantActor interface {
	gofhirclient.Resource

	Is_AdverseEventParticipantActor()
}

func (p Practitioner) Is_AdverseEventParticipantActor()     {}
func (p PractitionerRole) Is_AdverseEventParticipantActor() {}
func (o Organization) Is_AdverseEventParticipantActor()     {}
func (c CareTeam) Is_AdverseEventParticipantActor()         {}
func (p Patient) Is_AdverseEventParticipantActor()          {}
func (d Device) Is_AdverseEventParticipantActor()           {}
func (r RelatedPerson) Is_AdverseEventParticipantActor()    {}
func (r ResearchSubject) Is_AdverseEventParticipantActor()  {}

type AdverseEventSuspectEntityCausalityAuthor interface {
	gofhirclient.Resource

	Is_AdverseEventSuspectEntityCausalityAuthor()
}

func (p Practitioner) Is_AdverseEventSuspectEntityCausalityAuthor()     {}
func (p PractitionerRole) Is_AdverseEventSuspectEntityCausalityAuthor() {}
func (p Patient) Is_AdverseEventSuspectEntityCausalityAuthor()          {}
func (r RelatedPerson) Is_AdverseEventSuspectEntityCausalityAuthor()    {}
func (r ResearchSubject) Is_AdverseEventSuspectEntityCausalityAuthor()  {}

type AdverseEventContributingFactorItem interface {
	gofhirclient.Resource

	Is_AdverseEventContributingFactorItem()
}

func (c Condition) Is_AdverseEventContributingFactorItem()                {}
func (o Observation) Is_AdverseEventContributingFactorItem()              {}
func (a AllergyIntolerance) Is_AdverseEventContributingFactorItem()       {}
func (f FamilyMemberHistory) Is_AdverseEventContributingFactorItem()      {}
func (i Immunization) Is_AdverseEventContributingFactorItem()             {}
func (p Procedure) Is_AdverseEventContributingFactorItem()                {}
func (d Device) Is_AdverseEventContributingFactorItem()                   {}
func (d DeviceUsage) Is_AdverseEventContributingFactorItem()              {}
func (d DocumentReference) Is_AdverseEventContributingFactorItem()        {}
func (m MedicationAdministration) Is_AdverseEventContributingFactorItem() {}
func (m MedicationStatement) Is_AdverseEventContributingFactorItem()      {}

type AdverseEventPreventiveActionItem interface {
	gofhirclient.Resource

	Is_AdverseEventPreventiveActionItem()
}

func (i Immunization) Is_AdverseEventPreventiveActionItem()             {}
func (p Procedure) Is_AdverseEventPreventiveActionItem()                {}
func (d DocumentReference) Is_AdverseEventPreventiveActionItem()        {}
func (m MedicationAdministration) Is_AdverseEventPreventiveActionItem() {}
func (m MedicationRequest) Is_AdverseEventPreventiveActionItem()        {}

type AdverseEventMitigatingActionItem interface {
	gofhirclient.Resource

	Is_AdverseEventMitigatingActionItem()
}

func (p Procedure) Is_AdverseEventMitigatingActionItem()                {}
func (d DocumentReference) Is_AdverseEventMitigatingActionItem()        {}
func (m MedicationAdministration) Is_AdverseEventMitigatingActionItem() {}
func (m MedicationRequest) Is_AdverseEventMitigatingActionItem()        {}

type AdverseEventSupportingInfoItem interface {
	gofhirclient.Resource

	Is_AdverseEventSupportingInfoItem()
}

func (c Condition) Is_AdverseEventSupportingInfoItem()                {}
func (o Observation) Is_AdverseEventSupportingInfoItem()              {}
func (a AllergyIntolerance) Is_AdverseEventSupportingInfoItem()       {}
func (f FamilyMemberHistory) Is_AdverseEventSupportingInfoItem()      {}
func (i Immunization) Is_AdverseEventSupportingInfoItem()             {}
func (p Procedure) Is_AdverseEventSupportingInfoItem()                {}
func (d DocumentReference) Is_AdverseEventSupportingInfoItem()        {}
func (m MedicationAdministration) Is_AdverseEventSupportingInfoItem() {}
func (m MedicationStatement) Is_AdverseEventSupportingInfoItem()      {}
func (q QuestionnaireResponse) Is_AdverseEventSupportingInfoItem()    {}

func (a AdverseEvent) ResourceType() string {
	return "StructureDefinition"
}
