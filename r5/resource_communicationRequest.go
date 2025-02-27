// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CommunicationRequest
type CommunicationRequest struct {
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
	// Unique identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Fulfills plan or proposal
	BasedOn []custom.Reference[Resource] `json:"basedOn,omitempty"`
	// Request(s) replaced by this request
	Replaces []custom.Reference[CommunicationRequest] `json:"replaces,omitempty"`
	// Composite request this is part of
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `json:"intent"`
	// Message category
	Category []CodeableConcept `json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// True if request is prohibiting action
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	// A channel of communication
	Medium []CodeableConcept `json:"medium,omitempty"`
	// Focus of message
	Subject *custom.Reference[CommunicationRequestSubject] `json:"subject,omitempty"`
	// Resources that pertain to this communication request
	About []custom.Reference[Resource] `json:"about,omitempty"`
	// The Encounter during which this CommunicationRequest was created
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Message payload
	Payload []CommunicationRequestPayload `json:"payload,omitempty"`
	// When scheduled
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// When scheduled
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When request transitioned to being actionable
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Who asks for the information to be shared
	Requester *custom.Reference[CommunicationRequestRequester] `json:"requester,omitempty"`
	// Who to share the information with
	Recipient []custom.Reference[CommunicationRequestRecipient] `json:"recipient,omitempty"`
	// Who should share the information
	InformationProvider []custom.Reference[CommunicationRequestInformationProvider] `json:"informationProvider,omitempty"`
	// Why is communication needed?
	Reason []custom.CodeableReference[Resource] `json:"reason,omitempty"`
	// Comments made about communication request
	Note []Annotation `json:"note,omitempty"`
}

type CommunicationRequestPayload struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Message part content
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`
	// Message part content
	ContentReference *custom.Reference[Resource] `json:"contentReference,omitempty"`
	// Message part content
	ContentCodeableConcept *CodeableConcept `json:"contentCodeableConcept,omitempty"`
}

type OtherCommunicationRequest CommunicationRequest

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
	return "CommunicationRequest"
}

func (c CommunicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunicationRequest
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCommunicationRequest: OtherCommunicationRequest(c), ResourceType: c.ResourceType()})
}

func UnmarshalCommunicationRequest(b []byte) (res CommunicationRequest, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
