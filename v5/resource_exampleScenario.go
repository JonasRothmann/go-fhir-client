// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ExampleScenario
type ExampleScenarioActor struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// ID or acronym of the actor
	Key string `bson:"key" json:"key"`
	// person | system
	Type custom.Code `bson:"type" json:"type"`
	// Label for actor when rendering
	Title string `bson:"title" json:"title"`
	// Details about actor
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
}

type ExampleScenarioInstanceContainedInstance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Key of contained instance
	InstanceReference string `bson:"instanceReference" json:"instanceReference"`
	// Key of contained instance version
	VersionReference *string `bson:"versionReference,omitempty" json:"versionReference,omitempty"`
}

type ExampleScenarioProcessStep struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Sequential number of the step
	Number *string `bson:"number,omitempty" json:"number,omitempty"`
	// Step is nested process
	Process *interface{} `bson:"process,omitempty" json:"process,omitempty"`
	// Step is nested workflow
	Workflow *custom.Canonical[ExampleScenario] `bson:"workflow,omitempty" json:"workflow,omitempty"`
	// Step is simple action
	Operation *ExampleScenarioProcessStepOperation `bson:"operation,omitempty" json:"operation,omitempty"`
	// Alternate non-typical step action
	Alternative []ExampleScenarioProcessStepAlternative `bson:"alternative,omitempty" json:"alternative,omitempty"`
	// Pause in the flow?
	Pause *bool `bson:"pause,omitempty" json:"pause,omitempty"`
}

type ExampleScenarioProcessStepAlternative struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for alternative
	Title string `bson:"title" json:"title"`
	// Human-readable description of option
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Alternative action(s)
	Step []interface{} `bson:"step,omitempty" json:"step,omitempty"`
}

type ExampleScenario struct {
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
	// Canonical identifier for this example scenario, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the example scenario
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the example scenario
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// To be removed?
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this example scenario (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the ExampleScenario
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for example scenario (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// The purpose of the example, e.g. to illustrate a scenario
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Individual involved in exchange
	Actor []ExampleScenarioActor `bson:"actor,omitempty" json:"actor,omitempty"`
	// Data used in the scenario
	Instance []ExampleScenarioInstance `bson:"instance,omitempty" json:"instance,omitempty"`
	// Major process within scenario
	Process []ExampleScenarioProcess `bson:"process,omitempty" json:"process,omitempty"`
}

type ExampleScenarioInstance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// ID or acronym of the instance
	Key string `bson:"key" json:"key"`
	// Data structure for example
	StructureType Coding `bson:"structureType" json:"structureType"`
	// E.g. 4.0.1
	StructureVersion *string `bson:"structureVersion,omitempty" json:"structureVersion,omitempty"`
	// Rules instance adheres to
	StructureProfile *custom.Canonical `bson:"structureProfile,omitempty" json:"structureProfile,omitempty"`
	// Label for instance
	Title string `bson:"title" json:"title"`
	// Human-friendly description of the instance
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Example instance data
	Content *custom.Reference `bson:"content,omitempty" json:"content,omitempty"`
	// Snapshot of instance that changes
	Version []ExampleScenarioInstanceVersion `bson:"version,omitempty" json:"version,omitempty"`
	// Resources contained in the instance
	ContainedInstance []ExampleScenarioInstanceContainedInstance `bson:"containedInstance,omitempty" json:"containedInstance,omitempty"`
}

type ExampleScenarioInstanceVersion struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// ID or acronym of the version
	Key string `bson:"key" json:"key"`
	// Label for instance version
	Title string `bson:"title" json:"title"`
	// Details about version
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Example instance version data
	Content *custom.Reference `bson:"content,omitempty" json:"content,omitempty"`
}

type ExampleScenarioProcess struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for procss
	Title string `bson:"title" json:"title"`
	// Human-friendly description of the process
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Status before process starts
	PreConditions *custom.Markdown `bson:"preConditions,omitempty" json:"preConditions,omitempty"`
	// Status after successful completion
	PostConditions *custom.Markdown `bson:"postConditions,omitempty" json:"postConditions,omitempty"`
	// Event within of the process
	Step []ExampleScenarioProcessStep `bson:"step,omitempty" json:"step,omitempty"`
}

type ExampleScenarioProcessStepOperation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Kind of action
	Type *Coding `bson:"type,omitempty" json:"type,omitempty"`
	// Label for step
	Title string `bson:"title" json:"title"`
	// Who starts the operation
	Initiator *string `bson:"initiator,omitempty" json:"initiator,omitempty"`
	// Who receives the operation
	Receiver *string `bson:"receiver,omitempty" json:"receiver,omitempty"`
	// Human-friendly description of the operation
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Initiator stays active?
	InitiatorActive *bool `bson:"initiatorActive,omitempty" json:"initiatorActive,omitempty"`
	// Receiver stays active?
	ReceiverActive *bool `bson:"receiverActive,omitempty" json:"receiverActive,omitempty"`
	// Instance transmitted on invocation
	Request *interface{} `bson:"request,omitempty" json:"request,omitempty"`
	// Instance transmitted on invocation response
	Response *interface{} `bson:"response,omitempty" json:"response,omitempty"`
}

func (e ExampleScenario) ResourceType() string {
	return "StructureDefinition"
}
