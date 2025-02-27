// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AdverseEvent
type AdverseEventSuspectEntityCausality struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Method of evaluating the relatedness of the suspected entity to the event
	AssessmentMethod *CodeableConcept `json:"assessmentMethod,omitempty"`
	// Result of the assessment regarding the relatedness of the suspected entity to the event
	EntityRelatedness *CodeableConcept `json:"entityRelatedness,omitempty"`
	// Author of the information on the possible cause of the event
	Author *custom.Reference[AdverseEventSuspectEntityCausalityAuthor] `json:"author,omitempty"`
}

type AdverseEventContributingFactor struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Item suspected to have increased the probability or severity of the adverse event
	ItemReference *custom.Reference[AdverseEventContributingFactorItemReference] `json:"itemReference,omitempty"`
	// Item suspected to have increased the probability or severity of the adverse event
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

type AdverseEventPreventiveAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Action that contributed to avoiding the adverse event
	ItemReference *custom.Reference[AdverseEventPreventiveActionItemReference] `json:"itemReference,omitempty"`
	// Action that contributed to avoiding the adverse event
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

type AdverseEventMitigatingAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Ameliorating action taken after the adverse event occured in order to reduce the extent of harm
	ItemReference *custom.Reference[AdverseEventMitigatingActionItemReference] `json:"itemReference,omitempty"`
	// Ameliorating action taken after the adverse event occured in order to reduce the extent of harm
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

type AdverseEventSupportingInfo struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Subject medical history or document relevant to this adverse event
	ItemReference *custom.Reference[AdverseEventSupportingInfoItemReference] `json:"itemReference,omitempty"`
	// Subject medical history or document relevant to this adverse event
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

type AdverseEvent struct {
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
	// Business identifier for the event
	Identifier []Identifier `json:"identifier,omitempty"`
	// in-progress | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// actual | potential
	Actuality custom.Code `json:"actuality"`
	// wrong-patient | procedure-mishap | medication-mishap | device | unsafe-physical-environment | hospital-aquired-infection | wrong-body-site
	Category []CodeableConcept `json:"category,omitempty"`
	// Event or incident that occurred or was averted
	Code *CodeableConcept `json:"code,omitempty"`
	// Subject impacted by event
	Subject custom.Reference[AdverseEventSubject] `json:"subject"`
	// The Encounter associated with the start of the AdverseEvent
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When the event occurred
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// When the event occurred
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When the event occurred
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	// When the event was detected
	Detected *custom.DateTime `json:"detected,omitempty"`
	// When the event was recorded
	RecordedDate *custom.DateTime `json:"recordedDate,omitempty"`
	// Effect on the subject due to this event
	ResultingEffect []custom.Reference[AdverseEventResultingEffect] `json:"resultingEffect,omitempty"`
	// Location where adverse event occurred
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Seriousness or gravity of the event
	Seriousness *CodeableConcept `json:"seriousness,omitempty"`
	// Type of outcome from the adverse event
	Outcome []CodeableConcept `json:"outcome,omitempty"`
	// Who recorded the adverse event
	Recorder *custom.Reference[AdverseEventRecorder] `json:"recorder,omitempty"`
	// Who was involved in the adverse event or the potential adverse event and what they did
	Participant []AdverseEventParticipant `json:"participant,omitempty"`
	// Research study that the subject is enrolled in
	Study []custom.Reference[ResearchStudy] `json:"study,omitempty"`
	// Considered likely or probable or anticipated in the research study
	ExpectedInResearchStudy *bool `json:"expectedInResearchStudy,omitempty"`
	// The suspected agent causing the adverse event
	SuspectEntity []AdverseEventSuspectEntity `json:"suspectEntity,omitempty"`
	// Contributing factors suspected to have increased the probability or severity of the adverse event
	ContributingFactor []AdverseEventContributingFactor `json:"contributingFactor,omitempty"`
	// Preventive actions that contributed to avoiding the adverse event
	PreventiveAction []AdverseEventPreventiveAction `json:"preventiveAction,omitempty"`
	// Ameliorating actions taken after the adverse event occured in order to reduce the extent of harm
	MitigatingAction []AdverseEventMitigatingAction `json:"mitigatingAction,omitempty"`
	// Supporting information relevant to the event
	SupportingInfo []AdverseEventSupportingInfo `json:"supportingInfo,omitempty"`
	// Comment on adverse event
	Note []Annotation `json:"note,omitempty"`
}

type AdverseEventParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of involvement
	Function *CodeableConcept `json:"function,omitempty"`
	// Who was involved in the adverse event or the potential adverse event
	Actor custom.Reference[AdverseEventParticipantActor] `json:"actor"`
}

type AdverseEventSuspectEntity struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Refers to the specific entity that caused the adverse event
	InstanceCodeableConcept *CodeableConcept `json:"instanceCodeableConcept,omitempty"`
	// Refers to the specific entity that caused the adverse event
	InstanceReference *custom.Reference[AdverseEventSuspectEntityInstanceReference] `json:"instanceReference,omitempty"`
	// Information on the possible cause of the event
	Causality *AdverseEventSuspectEntityCausality `json:"causality,omitempty"`
}

type OtherAdverseEvent AdverseEvent

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

type AdverseEventSuspectEntityInstanceReference interface {
	gofhirclient.Resource

	Is_AdverseEventSuspectEntityInstanceReference()
}

func (i Immunization) Is_AdverseEventSuspectEntityInstanceReference()               {}
func (p Procedure) Is_AdverseEventSuspectEntityInstanceReference()                  {}
func (s Substance) Is_AdverseEventSuspectEntityInstanceReference()                  {}
func (m Medication) Is_AdverseEventSuspectEntityInstanceReference()                 {}
func (m MedicationAdministration) Is_AdverseEventSuspectEntityInstanceReference()   {}
func (m MedicationStatement) Is_AdverseEventSuspectEntityInstanceReference()        {}
func (d Device) Is_AdverseEventSuspectEntityInstanceReference()                     {}
func (b BiologicallyDerivedProduct) Is_AdverseEventSuspectEntityInstanceReference() {}
func (r ResearchStudy) Is_AdverseEventSuspectEntityInstanceReference()              {}

type AdverseEventSuspectEntityCausalityAuthor interface {
	gofhirclient.Resource

	Is_AdverseEventSuspectEntityCausalityAuthor()
}

func (p Practitioner) Is_AdverseEventSuspectEntityCausalityAuthor()     {}
func (p PractitionerRole) Is_AdverseEventSuspectEntityCausalityAuthor() {}
func (p Patient) Is_AdverseEventSuspectEntityCausalityAuthor()          {}
func (r RelatedPerson) Is_AdverseEventSuspectEntityCausalityAuthor()    {}
func (r ResearchSubject) Is_AdverseEventSuspectEntityCausalityAuthor()  {}

type AdverseEventContributingFactorItemReference interface {
	gofhirclient.Resource

	Is_AdverseEventContributingFactorItemReference()
}

func (c Condition) Is_AdverseEventContributingFactorItemReference()                {}
func (o Observation) Is_AdverseEventContributingFactorItemReference()              {}
func (a AllergyIntolerance) Is_AdverseEventContributingFactorItemReference()       {}
func (f FamilyMemberHistory) Is_AdverseEventContributingFactorItemReference()      {}
func (i Immunization) Is_AdverseEventContributingFactorItemReference()             {}
func (p Procedure) Is_AdverseEventContributingFactorItemReference()                {}
func (d Device) Is_AdverseEventContributingFactorItemReference()                   {}
func (d DeviceUsage) Is_AdverseEventContributingFactorItemReference()              {}
func (d DocumentReference) Is_AdverseEventContributingFactorItemReference()        {}
func (m MedicationAdministration) Is_AdverseEventContributingFactorItemReference() {}
func (m MedicationStatement) Is_AdverseEventContributingFactorItemReference()      {}

type AdverseEventPreventiveActionItemReference interface {
	gofhirclient.Resource

	Is_AdverseEventPreventiveActionItemReference()
}

func (i Immunization) Is_AdverseEventPreventiveActionItemReference()             {}
func (p Procedure) Is_AdverseEventPreventiveActionItemReference()                {}
func (d DocumentReference) Is_AdverseEventPreventiveActionItemReference()        {}
func (m MedicationAdministration) Is_AdverseEventPreventiveActionItemReference() {}
func (m MedicationRequest) Is_AdverseEventPreventiveActionItemReference()        {}

type AdverseEventMitigatingActionItemReference interface {
	gofhirclient.Resource

	Is_AdverseEventMitigatingActionItemReference()
}

func (p Procedure) Is_AdverseEventMitigatingActionItemReference()                {}
func (d DocumentReference) Is_AdverseEventMitigatingActionItemReference()        {}
func (m MedicationAdministration) Is_AdverseEventMitigatingActionItemReference() {}
func (m MedicationRequest) Is_AdverseEventMitigatingActionItemReference()        {}

type AdverseEventSupportingInfoItemReference interface {
	gofhirclient.Resource

	Is_AdverseEventSupportingInfoItemReference()
}

func (c Condition) Is_AdverseEventSupportingInfoItemReference()                {}
func (o Observation) Is_AdverseEventSupportingInfoItemReference()              {}
func (a AllergyIntolerance) Is_AdverseEventSupportingInfoItemReference()       {}
func (f FamilyMemberHistory) Is_AdverseEventSupportingInfoItemReference()      {}
func (i Immunization) Is_AdverseEventSupportingInfoItemReference()             {}
func (p Procedure) Is_AdverseEventSupportingInfoItemReference()                {}
func (d DocumentReference) Is_AdverseEventSupportingInfoItemReference()        {}
func (m MedicationAdministration) Is_AdverseEventSupportingInfoItemReference() {}
func (m MedicationStatement) Is_AdverseEventSupportingInfoItemReference()      {}
func (q QuestionnaireResponse) Is_AdverseEventSupportingInfoItemReference()    {}

func (a AdverseEvent) ResourceType() string {
	return "AdverseEvent"
}

func (a AdverseEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdverseEvent
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAdverseEvent: OtherAdverseEvent(a), ResourceType: a.ResourceType()})
}

func UnmarshalAdverseEvent(b []byte) (res AdverseEvent, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
