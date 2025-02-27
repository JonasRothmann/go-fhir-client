// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MessageHeader
type MessageHeader struct {
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
	// Event code or link to EventDefinition
	EventCoding *Coding `json:"eventCoding,omitempty"`
	// Event code or link to EventDefinition
	EventCanonical *custom.Canonical[EventDefinition] `json:"eventCanonical,omitempty"`
	// Message destination application(s)
	Destination []MessageHeaderDestination `json:"destination,omitempty"`
	// Real world sender of the message
	Sender *custom.Reference[MessageHeaderSender] `json:"sender,omitempty"`
	// The source of the decision
	Author *custom.Reference[MessageHeaderAuthor] `json:"author,omitempty"`
	// Message source application
	Source MessageHeaderSource `json:"source"`
	// Final responsibility for event
	Responsible *custom.Reference[MessageHeaderResponsible] `json:"responsible,omitempty"`
	// Cause of event
	Reason *CodeableConcept `json:"reason,omitempty"`
	// If this is a reply to prior message
	Response *MessageHeaderResponse `json:"response,omitempty"`
	// The actual content of the message
	Focus []custom.Reference[Resource] `json:"focus,omitempty"`
	// Link to the definition for this message
	Definition *custom.Canonical[MessageDefinition] `json:"definition,omitempty"`
}

type MessageHeaderDestination struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Actual destination address or Endpoint resource
	EndpointUrl *custom.URL `json:"endpointUrl,omitempty"`
	// Actual destination address or Endpoint resource
	EndpointReference *custom.Reference[Endpoint] `json:"endpointReference,omitempty"`
	// Name of system
	Name *string `json:"name,omitempty"`
	// Particular delivery destination within the destination
	Target *custom.Reference[Device] `json:"target,omitempty"`
	// Intended "real-world" recipient for the data
	Receiver *custom.Reference[MessageHeaderDestinationReceiver] `json:"receiver,omitempty"`
}

type MessageHeaderSource struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Actual source address or Endpoint resource
	EndpointUrl *custom.URL `json:"endpointUrl,omitempty"`
	// Actual source address or Endpoint resource
	EndpointReference *custom.Reference[Endpoint] `json:"endpointReference,omitempty"`
	// Name of system
	Name *string `json:"name,omitempty"`
	// Name of software running the system
	Software *string `json:"software,omitempty"`
	// Version of software running
	Version *string `json:"version,omitempty"`
	// Human contact for problems
	Contact *ContactPoint `json:"contact,omitempty"`
}

type MessageHeaderResponse struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Bundle.identifier of original message
	Identifier Identifier `json:"identifier"`
	// ok | transient-error | fatal-error
	Code custom.Code `json:"code"`
	// Specific list of hints/warnings/errors
	Details *custom.Reference[OperationOutcome] `json:"details,omitempty"`
}

type OtherMessageHeader MessageHeader

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
	return "MessageHeader"
}

func (m MessageHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageHeader
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMessageHeader: OtherMessageHeader(m), ResourceType: m.ResourceType()})
}

func UnmarshalMessageHeader(b []byte) (res MessageHeader, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
