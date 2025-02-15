// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MessageHeader
type MessageHeader struct {
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
	// Event code or link to EventDefinition
	Event Coding `bson:"event" json:"event"`
	// Message destination application(s)
	Destination []MessageHeaderDestination `bson:"destination,omitempty" json:"destination,omitempty"`
	// Real world sender of the message
	Sender *custom.Reference[MessageHeaderSender] `bson:"sender,omitempty" json:"sender,omitempty"`
	// The source of the decision
	Author *custom.Reference[MessageHeaderAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Message source application
	Source MessageHeaderSource `bson:"source" json:"source"`
	// Final responsibility for event
	Responsible *custom.Reference[MessageHeaderResponsible] `bson:"responsible,omitempty" json:"responsible,omitempty"`
	// Cause of event
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
	// If this is a reply to prior message
	Response *MessageHeaderResponse `bson:"response,omitempty" json:"response,omitempty"`
	// The actual content of the message
	Focus []custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// Link to the definition for this message
	Definition *custom.Canonical[MessageDefinition] `bson:"definition,omitempty" json:"definition,omitempty"`
}

type MessageHeaderDestination struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Actual destination address or Endpoint resource
	Endpoint *custom.URL `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// Name of system
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Particular delivery destination within the destination
	Target *custom.Reference[Device] `bson:"target,omitempty" json:"target,omitempty"`
	// Intended "real-world" recipient for the data
	Receiver *custom.Reference[MessageHeaderDestinationReceiver] `bson:"receiver,omitempty" json:"receiver,omitempty"`
}

type MessageHeaderSource struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Actual source address or Endpoint resource
	Endpoint *custom.URL `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// Name of system
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name of software running the system
	Software *string `bson:"software,omitempty" json:"software,omitempty"`
	// Version of software running
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// Human contact for problems
	Contact *ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
}

type MessageHeaderResponse struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Bundle.identifier of original message
	Identifier Identifier `bson:"identifier" json:"identifier"`
	// ok | transient-error | fatal-error
	Code custom.Code `bson:"code" json:"code"`
	// Specific list of hints/warnings/errors
	Details *custom.Reference[OperationOutcome] `bson:"details,omitempty" json:"details,omitempty"`
}

type MessageHeaderDestinationReceiver interface {
	gofhirclient.Resource

	Is_MessageHeaderDestinationReceiver()
}

func (p Practitioner) Is_MessageHeaderDestinationReceiver()     {}
func (p PractitionerRole) Is_MessageHeaderDestinationReceiver() {}
func (o Organization) Is_MessageHeaderDestinationReceiver()     {}

type MessageHeaderSender interface {
	gofhirclient.Resource

	Is_MessageHeaderSender()
}

func (p Practitioner) Is_MessageHeaderSender()     {}
func (p PractitionerRole) Is_MessageHeaderSender() {}
func (d Device) Is_MessageHeaderSender()           {}
func (o Organization) Is_MessageHeaderSender()     {}

type MessageHeaderAuthor interface {
	gofhirclient.Resource

	Is_MessageHeaderAuthor()
}

func (p Practitioner) Is_MessageHeaderAuthor()     {}
func (p PractitionerRole) Is_MessageHeaderAuthor() {}
func (d Device) Is_MessageHeaderAuthor()           {}
func (o Organization) Is_MessageHeaderAuthor()     {}

type MessageHeaderResponsible interface {
	gofhirclient.Resource

	Is_MessageHeaderResponsible()
}

func (p Practitioner) Is_MessageHeaderResponsible()     {}
func (p PractitionerRole) Is_MessageHeaderResponsible() {}
func (o Organization) Is_MessageHeaderResponsible()     {}

func (m MessageHeader) ResourceType() string {
	return "StructureDefinition"
}
