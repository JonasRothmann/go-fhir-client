// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CommunicationRequest
type CommunicationRequest struct {
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
	// Unique identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Fulfills plan or proposal
	BasedOn []custom.Reference[Resource] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Request(s) replaced by this request
	Replaces []custom.Reference[CommunicationRequest] `bson:"replaces,omitempty" json:"replaces,omitempty"`
	// Composite request this is part of
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `bson:"intent" json:"intent"`
	// Message category
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// True if request is prohibiting action
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	// A channel of communication
	Medium []CodeableConcept `bson:"medium,omitempty" json:"medium,omitempty"`
	// Focus of message
	Subject *custom.Reference[CommunicationRequestSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Resources that pertain to this communication request
	About []custom.Reference[Resource] `bson:"about,omitempty" json:"about,omitempty"`
	// The Encounter during which this CommunicationRequest was created
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Message payload
	Payload []CommunicationRequestPayload `bson:"payload,omitempty" json:"payload,omitempty"`
	// When scheduled
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// When request transitioned to being actionable
	AuthoredOn *custom.DateTime `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	// Who asks for the information to be shared
	Requester *custom.Reference[CommunicationRequestRequester] `bson:"requester,omitempty" json:"requester,omitempty"`
	// Who to share the information with
	Recipient []custom.Reference[CommunicationRequestRecipient] `bson:"recipient,omitempty" json:"recipient,omitempty"`
	// Who should share the information
	InformationProvider []custom.Reference[CommunicationRequestInformationProvider] `bson:"informationProvider,omitempty" json:"informationProvider,omitempty"`
	// Why is communication needed?
	Reason []custom.CodeableReference[Resource] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Comments made about communication request
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type CommunicationRequestPayload struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Message part content
	Content Attachment `bson:"content" json:"content"`
}

type CommunicationRequestSubject interface {
	gofhirclient.Resource

	Is_CommunicationRequestSubject()
}

func (p Patient) Is_CommunicationRequestSubject() {}
func (g Group) Is_CommunicationRequestSubject()   {}

type CommunicationRequestRequester interface {
	gofhirclient.Resource

	Is_CommunicationRequestRequester()
}

func (p Practitioner) Is_CommunicationRequestRequester()     {}
func (p PractitionerRole) Is_CommunicationRequestRequester() {}
func (o Organization) Is_CommunicationRequestRequester()     {}
func (p Patient) Is_CommunicationRequestRequester()          {}
func (r RelatedPerson) Is_CommunicationRequestRequester()    {}
func (d Device) Is_CommunicationRequestRequester()           {}

type CommunicationRequestRecipient interface {
	gofhirclient.Resource

	Is_CommunicationRequestRecipient()
}

func (d Device) Is_CommunicationRequestRecipient()            {}
func (o Organization) Is_CommunicationRequestRecipient()      {}
func (p Patient) Is_CommunicationRequestRecipient()           {}
func (p Practitioner) Is_CommunicationRequestRecipient()      {}
func (p PractitionerRole) Is_CommunicationRequestRecipient()  {}
func (r RelatedPerson) Is_CommunicationRequestRecipient()     {}
func (g Group) Is_CommunicationRequestRecipient()             {}
func (c CareTeam) Is_CommunicationRequestRecipient()          {}
func (h HealthcareService) Is_CommunicationRequestRecipient() {}
func (e Endpoint) Is_CommunicationRequestRecipient()          {}

type CommunicationRequestInformationProvider interface {
	gofhirclient.Resource

	Is_CommunicationRequestInformationProvider()
}

func (d Device) Is_CommunicationRequestInformationProvider()            {}
func (o Organization) Is_CommunicationRequestInformationProvider()      {}
func (p Patient) Is_CommunicationRequestInformationProvider()           {}
func (p Practitioner) Is_CommunicationRequestInformationProvider()      {}
func (p PractitionerRole) Is_CommunicationRequestInformationProvider()  {}
func (r RelatedPerson) Is_CommunicationRequestInformationProvider()     {}
func (h HealthcareService) Is_CommunicationRequestInformationProvider() {}
func (e Endpoint) Is_CommunicationRequestInformationProvider()          {}

func (c CommunicationRequest) ResourceType() string {
	return "StructureDefinition"
}
