// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ExampleScenario
type ExampleScenario struct {
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
	// Canonical identifier for this example scenario, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the example scenario
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the example scenario
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// To be removed?
	Name *string `json:"name,omitempty"`
	// Name for this example scenario (human friendly)
	Title *string `json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the ExampleScenario
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for example scenario (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// The purpose of the example, e.g. to illustrate a scenario
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Individual involved in exchange
	Actor []ExampleScenarioActor `json:"actor,omitempty"`
	// Data used in the scenario
	Instance []ExampleScenarioInstance `json:"instance,omitempty"`
	// Major process within scenario
	Process []ExampleScenarioProcess `json:"process,omitempty"`
}

type ExampleScenarioProcess struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for procss
	Title string `json:"title"`
	// Human-friendly description of the process
	Description *custom.Markdown `json:"description,omitempty"`
	// Status before process starts
	PreConditions *custom.Markdown `json:"preConditions,omitempty"`
	// Status after successful completion
	PostConditions *custom.Markdown `json:"postConditions,omitempty"`
	// Event within of the process
	Step []ExampleScenarioProcessStep `json:"step,omitempty"`
}

type ExampleScenarioProcessStepAlternative struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for alternative
	Title string `json:"title"`
	// Human-readable description of option
	Description *custom.Markdown `json:"description,omitempty"`
	// Alternative action(s)
	Step []interface{} `json:"step,omitempty"`
}

type ExampleScenarioProcessStep struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Sequential number of the step
	Number *string `json:"number,omitempty"`
	// Step is nested process
	Process *interface{} `json:"process,omitempty"`
	// Step is nested workflow
	Workflow *custom.Canonical[ExampleScenario] `json:"workflow,omitempty"`
	// Step is simple action
	Operation *ExampleScenarioProcessStepOperation `json:"operation,omitempty"`
	// Alternate non-typical step action
	Alternative []ExampleScenarioProcessStepAlternative `json:"alternative,omitempty"`
	// Pause in the flow?
	Pause *bool `json:"pause,omitempty"`
}

type ExampleScenarioProcessStepOperation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Kind of action
	Type *Coding `json:"type,omitempty"`
	// Label for step
	Title string `json:"title"`
	// Who starts the operation
	Initiator *string `json:"initiator,omitempty"`
	// Who receives the operation
	Receiver *string `json:"receiver,omitempty"`
	// Human-friendly description of the operation
	Description *custom.Markdown `json:"description,omitempty"`
	// Initiator stays active?
	InitiatorActive *bool `json:"initiatorActive,omitempty"`
	// Receiver stays active?
	ReceiverActive *bool `json:"receiverActive,omitempty"`
	// Instance transmitted on invocation
	Request *interface{} `json:"request,omitempty"`
	// Instance transmitted on invocation response
	Response *interface{} `json:"response,omitempty"`
}

type ExampleScenarioActor struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// ID or acronym of the actor
	Key string `json:"key"`
	// person | system
	Type custom.Code `json:"type"`
	// Label for actor when rendering
	Title string `json:"title"`
	// Details about actor
	Description *custom.Markdown `json:"description,omitempty"`
}

type ExampleScenarioInstance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// ID or acronym of the instance
	Key string `json:"key"`
	// Data structure for example
	StructureType Coding `json:"structureType"`
	// E.g. 4.0.1
	StructureVersion *string `json:"structureVersion,omitempty"`
	// Rules instance adheres to
	StructureProfileCanonical *custom.Canonical[any] `json:"structureProfileCanonical,omitempty"`
	// Rules instance adheres to
	StructureProfileUri *custom.URI `json:"structureProfileUri,omitempty"`
	// Label for instance
	Title string `json:"title"`
	// Human-friendly description of the instance
	Description *custom.Markdown `json:"description,omitempty"`
	// Example instance data
	Content *custom.Reference[gofhirclient.Resource] `json:"content,omitempty"`
	// Snapshot of instance that changes
	Version []ExampleScenarioInstanceVersion `json:"version,omitempty"`
	// Resources contained in the instance
	ContainedInstance []ExampleScenarioInstanceContainedInstance `json:"containedInstance,omitempty"`
}

type ExampleScenarioInstanceVersion struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// ID or acronym of the version
	Key string `json:"key"`
	// Label for instance version
	Title string `json:"title"`
	// Details about version
	Description *custom.Markdown `json:"description,omitempty"`
	// Example instance version data
	Content *custom.Reference[gofhirclient.Resource] `json:"content,omitempty"`
}

type ExampleScenarioInstanceContainedInstance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Key of contained instance
	InstanceReference string `json:"instanceReference"`
	// Key of contained instance version
	VersionReference *string `json:"versionReference,omitempty"`
}

type OtherExampleScenario ExampleScenario

func (e ExampleScenario) ResourceType() string {
	return "ExampleScenario"
}

func (e ExampleScenario) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExampleScenario
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherExampleScenario: OtherExampleScenario(e), ResourceType: e.ResourceType()})
}

func UnmarshalExampleScenario(b []byte) (res ExampleScenario, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
