// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Task
type TaskOutput struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for output
	Type CodeableConcept `bson:"type" json:"type"`
	// Result of output
	Value custom.Base64binary `bson:"value" json:"value"`
}

type Task struct {
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
	// Task Instance Identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Formal definition of task
	InstantiatesCanonical *custom.Canonical[ActivityDefinition] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Formal definition of task
	InstantiatesUri *custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Request fulfilled by this task
	BasedOn []custom.Reference[Resource] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Requisition or grouper id
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	// Composite task
	PartOf []custom.Reference[Task] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// draft | requested | received | accepted | +
	Status custom.Code `bson:"status" json:"status"`
	// Reason for current status
	StatusReason *custom.CodeableReference `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// E.g. "Specimen collected", "IV prepped"
	BusinessStatus *CodeableConcept `bson:"businessStatus,omitempty" json:"businessStatus,omitempty"`
	// unknown | proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `bson:"intent" json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// True if Task is prohibiting action
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	// Task Type
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Human-readable explanation of task
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// What task is acting on
	Focus *custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// Beneficiary of the Task
	For *custom.Reference[Resource] `bson:"for,omitempty" json:"for,omitempty"`
	// Healthcare event during which this task originated
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the task should be performed
	RequestedPeriod *Period `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
	// Start and end time of execution
	ExecutionPeriod *Period `bson:"executionPeriod,omitempty" json:"executionPeriod,omitempty"`
	// Task Creation Date
	AuthoredOn *custom.DateTime `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	// Task Last Modified Date
	LastModified *custom.DateTime `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	// Who is asking for task to be done
	Requester *custom.Reference[TaskRequester] `bson:"requester,omitempty" json:"requester,omitempty"`
	// Who should perform Task
	RequestedPerformer []custom.CodeableReference[TaskRequestedPerformer] `bson:"requestedPerformer,omitempty" json:"requestedPerformer,omitempty"`
	// Responsible individual
	Owner *custom.Reference[TaskOwner] `bson:"owner,omitempty" json:"owner,omitempty"`
	// Who or what performed the task
	Performer []TaskPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Where task occurs
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Why task is needed
	Reason []custom.CodeableReference `bson:"reason,omitempty" json:"reason,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[TaskInsurance] `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Comments made about the task
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Key events in history of the Task
	RelevantHistory []custom.Reference[Provenance] `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	// Constraints on fulfillment tasks
	Restriction *TaskRestriction `bson:"restriction,omitempty" json:"restriction,omitempty"`
	// Information used to perform task
	Input []TaskInput `bson:"input,omitempty" json:"input,omitempty"`
	// Information produced as part of task
	Output []TaskOutput `bson:"output,omitempty" json:"output,omitempty"`
}

type TaskPerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of performance
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who performed the task
	Actor custom.Reference[TaskPerformerActor] `bson:"actor" json:"actor"`
}

type TaskRestriction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// How many times to repeat
	Repetitions *uint32 `bson:"repetitions,omitempty" json:"repetitions,omitempty"`
	// When fulfillment is sought
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// For whom is fulfillment sought?
	Recipient []custom.Reference[TaskRestrictionRecipient] `bson:"recipient,omitempty" json:"recipient,omitempty"`
}

type TaskInput struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for the input
	Type CodeableConcept `bson:"type" json:"type"`
	// Content to use in performing the task
	Value custom.Base64binary `bson:"value" json:"value"`
}

type TaskRequester interface {
	gofhirclient.Resource

	Is_TaskRequester()
}

func (d Device) Is_TaskRequester()           {}
func (o Organization) Is_TaskRequester()     {}
func (p Patient) Is_TaskRequester()          {}
func (p Practitioner) Is_TaskRequester()     {}
func (p PractitionerRole) Is_TaskRequester() {}
func (r RelatedPerson) Is_TaskRequester()    {}

type TaskRequestedPerformer interface {
	gofhirclient.Resource

	Is_TaskRequestedPerformer()
}

func (p Practitioner) Is_TaskRequestedPerformer()      {}
func (p PractitionerRole) Is_TaskRequestedPerformer()  {}
func (o Organization) Is_TaskRequestedPerformer()      {}
func (c CareTeam) Is_TaskRequestedPerformer()          {}
func (h HealthcareService) Is_TaskRequestedPerformer() {}
func (p Patient) Is_TaskRequestedPerformer()           {}
func (d Device) Is_TaskRequestedPerformer()            {}
func (r RelatedPerson) Is_TaskRequestedPerformer()     {}

type TaskOwner interface {
	gofhirclient.Resource

	Is_TaskOwner()
}

func (p Practitioner) Is_TaskOwner()     {}
func (p PractitionerRole) Is_TaskOwner() {}
func (o Organization) Is_TaskOwner()     {}
func (c CareTeam) Is_TaskOwner()         {}
func (p Patient) Is_TaskOwner()          {}
func (r RelatedPerson) Is_TaskOwner()    {}

type TaskPerformerActor interface {
	gofhirclient.Resource

	Is_TaskPerformerActor()
}

func (p Practitioner) Is_TaskPerformerActor()     {}
func (p PractitionerRole) Is_TaskPerformerActor() {}
func (o Organization) Is_TaskPerformerActor()     {}
func (c CareTeam) Is_TaskPerformerActor()         {}
func (p Patient) Is_TaskPerformerActor()          {}
func (r RelatedPerson) Is_TaskPerformerActor()    {}

type TaskInsurance interface {
	gofhirclient.Resource

	Is_TaskInsurance()
}

func (c Coverage) Is_TaskInsurance()      {}
func (c ClaimResponse) Is_TaskInsurance() {}

type TaskRestrictionRecipient interface {
	gofhirclient.Resource

	Is_TaskRestrictionRecipient()
}

func (p Patient) Is_TaskRestrictionRecipient()          {}
func (p Practitioner) Is_TaskRestrictionRecipient()     {}
func (p PractitionerRole) Is_TaskRestrictionRecipient() {}
func (r RelatedPerson) Is_TaskRestrictionRecipient()    {}
func (g Group) Is_TaskRestrictionRecipient()            {}
func (o Organization) Is_TaskRestrictionRecipient()     {}

func (t Task) ResourceType() string {
	return "StructureDefinition"
}
