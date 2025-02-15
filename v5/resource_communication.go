// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Communication
type Communication struct {
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
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[CommunicationInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Request fulfilled by this communication
	BasedOn []custom.Reference[Resource] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Part of referenced event (e.g. Communication, Procedure)
	PartOf []custom.Reference[Resource] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// Reply to
	InResponseTo []custom.Reference[Communication] `bson:"inResponseTo,omitempty" json:"inResponseTo,omitempty"`
	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// Message category
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// A channel of communication
	Medium []CodeableConcept `bson:"medium,omitempty" json:"medium,omitempty"`
	// Focus of message
	Subject *custom.Reference[CommunicationSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Description of the purpose/content
	Topic *CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Resources that pertain to this communication
	About []custom.Reference[Resource] `bson:"about,omitempty" json:"about,omitempty"`
	// The Encounter during which this Communication was created
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When sent
	Sent *custom.DateTime `bson:"sent,omitempty" json:"sent,omitempty"`
	// When received
	Received *custom.DateTime `bson:"received,omitempty" json:"received,omitempty"`
	// Who the information is shared with
	Recipient []custom.Reference[CommunicationRecipient] `bson:"recipient,omitempty" json:"recipient,omitempty"`
	// Who shares the information
	Sender *custom.Reference[CommunicationSender] `bson:"sender,omitempty" json:"sender,omitempty"`
	// Indication for message
	Reason []custom.CodeableReference[Resource] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Message payload
	Payload []CommunicationPayload `bson:"payload,omitempty" json:"payload,omitempty"`
	// Comments made about the communication
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type CommunicationPayload struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Message part content
	Content Attachment `bson:"content" json:"content"`
}

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
	return "StructureDefinition"
}
