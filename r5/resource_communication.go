// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Communication
type CommunicationPayload struct {
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

type Communication struct {
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
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[CommunicationInstantiatesCanonical] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `json:"instantiatesUri,omitempty"`
	// Request fulfilled by this communication
	BasedOn []custom.Reference[Resource] `json:"basedOn,omitempty"`
	// Part of referenced event (e.g. Communication, Procedure)
	PartOf []custom.Reference[Resource] `json:"partOf,omitempty"`
	// Reply to
	InResponseTo []custom.Reference[Communication] `json:"inResponseTo,omitempty"`
	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	// Message category
	Category []CodeableConcept `json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// A channel of communication
	Medium []CodeableConcept `json:"medium,omitempty"`
	// Focus of message
	Subject *custom.Reference[CommunicationSubject] `json:"subject,omitempty"`
	// Description of the purpose/content
	Topic *CodeableConcept `json:"topic,omitempty"`
	// Resources that pertain to this communication
	About []custom.Reference[Resource] `json:"about,omitempty"`
	// The Encounter during which this Communication was created
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When sent
	Sent *custom.DateTime `json:"sent,omitempty"`
	// When received
	Received *custom.DateTime `json:"received,omitempty"`
	// Who the information is shared with
	Recipient []custom.Reference[CommunicationRecipient] `json:"recipient,omitempty"`
	// Who shares the information
	Sender *custom.Reference[CommunicationSender] `json:"sender,omitempty"`
	// Indication for message
	Reason []custom.CodeableReference[Resource] `json:"reason,omitempty"`
	// Message payload
	Payload []CommunicationPayload `json:"payload,omitempty"`
	// Comments made about the communication
	Note []Annotation `json:"note,omitempty"`
}

type OtherCommunication Communication

type CommunicationInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_CommunicationInstantiatesCanonical()
}

func (p PlanDefinition) Is_CommunicationInstantiatesCanonical()      {}
func (a ActivityDefinition) Is_CommunicationInstantiatesCanonical()  {}
func (m Measure) Is_CommunicationInstantiatesCanonical()             {}
func (o OperationDefinition) Is_CommunicationInstantiatesCanonical() {}
func (q Questionnaire) Is_CommunicationInstantiatesCanonical()       {}

type CommunicationSubject interface {
	gofhirclient.Resource

	Is_CommunicationSubject()
}

func (p Patient) Is_CommunicationSubject() {}
func (g Group) Is_CommunicationSubject()   {}

type CommunicationRecipient interface {
	gofhirclient.Resource

	Is_CommunicationRecipient()
}

func (c CareTeam) Is_CommunicationRecipient()          {}
func (d Device) Is_CommunicationRecipient()            {}
func (g Group) Is_CommunicationRecipient()             {}
func (h HealthcareService) Is_CommunicationRecipient() {}
func (l Location) Is_CommunicationRecipient()          {}
func (o Organization) Is_CommunicationRecipient()      {}
func (p Patient) Is_CommunicationRecipient()           {}
func (p Practitioner) Is_CommunicationRecipient()      {}
func (p PractitionerRole) Is_CommunicationRecipient()  {}
func (r RelatedPerson) Is_CommunicationRecipient()     {}
func (e Endpoint) Is_CommunicationRecipient()          {}

type CommunicationSender interface {
	gofhirclient.Resource

	Is_CommunicationSender()
}

func (d Device) Is_CommunicationSender()            {}
func (o Organization) Is_CommunicationSender()      {}
func (p Patient) Is_CommunicationSender()           {}
func (p Practitioner) Is_CommunicationSender()      {}
func (p PractitionerRole) Is_CommunicationSender()  {}
func (r RelatedPerson) Is_CommunicationSender()     {}
func (h HealthcareService) Is_CommunicationSender() {}
func (e Endpoint) Is_CommunicationSender()          {}
func (c CareTeam) Is_CommunicationSender()          {}

func (c Communication) ResourceType() string {
	return "Communication"
}

func (c Communication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunication
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCommunication: OtherCommunication(c), ResourceType: c.ResourceType()})
}

func UnmarshalCommunication(b []byte) (res Communication, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
