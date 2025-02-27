// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MessageDefinition
type MessageDefinitionFocus struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of resource
	Code custom.Code `json:"code"`
	// Profile that must be adhered to by focus
	Profile *custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
	// Minimum number of focuses of this type
	Min uint32 `json:"min"`
	// Maximum number of focuses of this type
	Max *string `json:"max,omitempty"`
}

type MessageDefinitionAllowedResponse struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to allowed message definition response
	Message custom.Canonical[MessageDefinition] `json:"message"`
	// When should this response be used
	Situation *custom.Markdown `json:"situation,omitempty"`
}

type MessageDefinition struct {
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
	// The cannonical URL for a given MessageDefinition
	Url *custom.URI `json:"url,omitempty"`
	// Business Identifier for a given MessageDefinition
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the message definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this message definition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this message definition (human friendly)
	Title *string `json:"title,omitempty"`
	// Takes the place of
	Replaces []custom.Canonical[MessageDefinition] `json:"replaces,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date custom.DateTime `json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the message definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for message definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this message definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Definition this one is based on
	Base *custom.Canonical[MessageDefinition] `json:"base,omitempty"`
	// Protocol/workflow this is part of
	Parent []custom.Canonical[MessageDefinitionParent] `json:"parent,omitempty"`
	// Event code  or link to the EventDefinition
	EventCoding *Coding `json:"eventCoding,omitempty"`
	// Event code  or link to the EventDefinition
	EventUri *custom.URI `json:"eventUri,omitempty"`
	// consequence | currency | notification
	Category *custom.Code `json:"category,omitempty"`
	// Resource(s) that are the subject of the event
	Focus []MessageDefinitionFocus `json:"focus,omitempty"`
	// always | on-error | never | on-success
	ResponseRequired *custom.Code `json:"responseRequired,omitempty"`
	// Responses to this message
	AllowedResponse []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty"`
	// Canonical reference to a GraphDefinition
	Graph *custom.Canonical[GraphDefinition] `json:"graph,omitempty"`
}

type OtherMessageDefinition MessageDefinition

type MessageDefinitionParent interface {
	gofhirclient.Resource

	Is_MessageDefinitionParent()
}

func (a ActivityDefinition) Is_MessageDefinitionParent() {}
func (p PlanDefinition) Is_MessageDefinitionParent()     {}

func (m MessageDefinition) ResourceType() string {
	return "MessageDefinition"
}

func (m MessageDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMessageDefinition: OtherMessageDefinition(m), ResourceType: m.ResourceType()})
}

func UnmarshalMessageDefinition(b []byte) (res MessageDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
