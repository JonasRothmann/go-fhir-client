// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MessageDefinition
type MessageDefinition struct {
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
	// The cannonical URL for a given MessageDefinition
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Business Identifier for a given MessageDefinition
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the message definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this message definition (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this message definition (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Takes the place of
	Replaces []custom.Canonical[MessageDefinition] `bson:"replaces,omitempty" json:"replaces,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date custom.DateTime `bson:"date" json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the message definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for message definition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this message definition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Definition this one is based on
	Base *custom.Canonical[MessageDefinition] `bson:"base,omitempty" json:"base,omitempty"`
	// Protocol/workflow this is part of
	Parent []custom.Canonical[MessageDefinitionParent] `bson:"parent,omitempty" json:"parent,omitempty"`
	// Event code  or link to the EventDefinition
	Event Coding `bson:"event" json:"event"`
	// consequence | currency | notification
	Category *custom.Code `bson:"category,omitempty" json:"category,omitempty"`
	// Resource(s) that are the subject of the event
	Focus []MessageDefinitionFocus `bson:"focus,omitempty" json:"focus,omitempty"`
	// always | on-error | never | on-success
	ResponseRequired *custom.Code `bson:"responseRequired,omitempty" json:"responseRequired,omitempty"`
	// Responses to this message
	AllowedResponse []MessageDefinitionAllowedResponse `bson:"allowedResponse,omitempty" json:"allowedResponse,omitempty"`
	// Canonical reference to a GraphDefinition
	Graph *custom.Canonical[GraphDefinition] `bson:"graph,omitempty" json:"graph,omitempty"`
}

type MessageDefinitionFocus struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of resource
	Code custom.Code `bson:"code" json:"code"`
	// Profile that must be adhered to by focus
	Profile *custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	// Minimum number of focuses of this type
	Min uint32 `bson:"min" json:"min"`
	// Maximum number of focuses of this type
	Max *string `bson:"max,omitempty" json:"max,omitempty"`
}

type MessageDefinitionAllowedResponse struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to allowed message definition response
	Message custom.Canonical[MessageDefinition] `bson:"message" json:"message"`
	// When should this response be used
	Situation *custom.Markdown `bson:"situation,omitempty" json:"situation,omitempty"`
}

type MessageDefinitionParent interface {
	gofhirclient.Resource

	Is_MessageDefinitionParent()
}

func (a ActivityDefinition) Is_MessageDefinitionParent() {}
func (p PlanDefinition) Is_MessageDefinitionParent()     {}

func (m MessageDefinition) ResourceType() string {
	return "StructureDefinition"
}
