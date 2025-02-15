// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Transport
type Transport struct {
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
	// External identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Formal definition of transport
	InstantiatesCanonical *custom.Canonical[ActivityDefinition] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Formal definition of transport
	InstantiatesUri *custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Request fulfilled by this transport
	BasedOn []custom.Reference[Resource] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Requisition or grouper id
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[Transport] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// in-progress | completed | abandoned | cancelled | planned | entered-in-error
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Reason for current status
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// unknown | proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `bson:"intent" json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// Transport Type
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Human-readable explanation of transport
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// What transport is acting on
	Focus *custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// Beneficiary of the Transport
	For *custom.Reference[Resource] `bson:"for,omitempty" json:"for,omitempty"`
	// Healthcare event during which this transport originated
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Completion time of the event (the occurrence)
	CompletionTime *custom.DateTime `bson:"completionTime,omitempty" json:"completionTime,omitempty"`
	// Transport Creation Date
	AuthoredOn *custom.DateTime `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	// Transport Last Modified Date
	LastModified *custom.DateTime `bson:"lastModified,omitempty" json:"lastModified,omitempty"`
	// Who is asking for transport to be done
	Requester *custom.Reference[TransportRequester] `bson:"requester,omitempty" json:"requester,omitempty"`
	// Requested performer
	PerformerType []CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty"`
	// Responsible individual
	Owner *custom.Reference[TransportOwner] `bson:"owner,omitempty" json:"owner,omitempty"`
	// Where transport occurs
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[TransportInsurance] `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Comments made about the transport
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Key events in history of the Transport
	RelevantHistory []custom.Reference[Provenance] `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	// Constraints on fulfillment transports
	Restriction *TransportRestriction `bson:"restriction,omitempty" json:"restriction,omitempty"`
	// Information used to perform transport
	Input []TransportInput `bson:"input,omitempty" json:"input,omitempty"`
	// Information produced as part of transport
	Output []TransportOutput `bson:"output,omitempty" json:"output,omitempty"`
	// The desired location
	RequestedLocation custom.Reference[Location] `bson:"requestedLocation" json:"requestedLocation"`
	// The entity current location
	CurrentLocation custom.Reference[Location] `bson:"currentLocation" json:"currentLocation"`
	// Why transport is needed
	Reason *custom.CodeableReference[Resource] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Parent (or preceding) transport
	History *custom.Reference[Transport] `bson:"history,omitempty" json:"history,omitempty"`
}

type TransportRestriction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// How many times to repeat
	Repetitions *uint32 `bson:"repetitions,omitempty" json:"repetitions,omitempty"`
	// When fulfillment sought
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// For whom is fulfillment sought?
	Recipient []custom.Reference[TransportRestrictionRecipient] `bson:"recipient,omitempty" json:"recipient,omitempty"`
}

type TransportInput struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for the input
	Type CodeableConcept `bson:"type" json:"type"`
	// Content to use in performing the transport
	Value custom.Base64binary `bson:"value" json:"value"`
}

type TransportOutput struct {
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

type TransportRequester interface {
	gofhirclient.Resource

	Is_TransportRequester()
}

func (d Device) Is_TransportRequester()           {}
func (o Organization) Is_TransportRequester()     {}
func (p Patient) Is_TransportRequester()          {}
func (p Practitioner) Is_TransportRequester()     {}
func (p PractitionerRole) Is_TransportRequester() {}
func (r RelatedPerson) Is_TransportRequester()    {}

type TransportOwner interface {
	gofhirclient.Resource

	Is_TransportOwner()
}

func (p Practitioner) Is_TransportOwner()      {}
func (p PractitionerRole) Is_TransportOwner()  {}
func (o Organization) Is_TransportOwner()      {}
func (c CareTeam) Is_TransportOwner()          {}
func (h HealthcareService) Is_TransportOwner() {}
func (p Patient) Is_TransportOwner()           {}
func (d Device) Is_TransportOwner()            {}
func (r RelatedPerson) Is_TransportOwner()     {}

type TransportInsurance interface {
	gofhirclient.Resource

	Is_TransportInsurance()
}

func (c Coverage) Is_TransportInsurance()      {}
func (c ClaimResponse) Is_TransportInsurance() {}

type TransportRestrictionRecipient interface {
	gofhirclient.Resource

	Is_TransportRestrictionRecipient()
}

func (p Patient) Is_TransportRestrictionRecipient()          {}
func (p Practitioner) Is_TransportRestrictionRecipient()     {}
func (p PractitionerRole) Is_TransportRestrictionRecipient() {}
func (r RelatedPerson) Is_TransportRestrictionRecipient()    {}
func (g Group) Is_TransportRestrictionRecipient()            {}
func (o Organization) Is_TransportRestrictionRecipient()     {}

func (t Transport) ResourceType() string {
	return "StructureDefinition"
}
