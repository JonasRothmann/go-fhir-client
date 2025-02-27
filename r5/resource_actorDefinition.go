// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ActorDefinition
type ActorDefinition struct {
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
	// Canonical identifier for this actor definition, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the actor definition (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the actor definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this actor definition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this actor definition (human friendly)
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
	// Natural language description of the actor
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for actor definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this actor definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// person | system
	Type custom.Code `json:"type"`
	// Functionality associated with the actor
	Documentation *custom.Markdown `json:"documentation,omitempty"`
	// Reference to more information about the actor
	Reference []custom.URL `json:"reference,omitempty"`
	// CapabilityStatement for the actor (if applicable)
	Capabilities *custom.Canonical[CapabilityStatement] `json:"capabilities,omitempty"`
	// Definition of this actor in another context / IG
	DerivedFrom []custom.Canonical[ActorDefinition] `json:"derivedFrom,omitempty"`
}

type OtherActorDefinition ActorDefinition

func (a ActorDefinition) ResourceType() string {
	return "ActorDefinition"
}

func (a ActorDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherActorDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherActorDefinition: OtherActorDefinition(a), ResourceType: a.ResourceType()})
}

func UnmarshalActorDefinition(b []byte) (res ActorDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
