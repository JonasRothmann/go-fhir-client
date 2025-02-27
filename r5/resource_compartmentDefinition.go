// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
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
	// Canonical identifier for this compartment definition, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Business version of the compartment definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this compartment definition (computer friendly)
	Name string `json:"name"`
	// Name for this compartment definition (human friendly)
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
	// Natural language description of the compartment definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Why this compartment definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Patient | Encounter | RelatedPerson | Practitioner | Device | EpisodeOfCare
	Code custom.Code `json:"code"`
	// Whether the search syntax is supported
	Search bool `json:"search"`
	// How a resource is related to the compartment
	Resource []CompartmentDefinitionResource `json:"resource,omitempty"`
}

type CompartmentDefinitionResource struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name of resource type
	Code custom.Code `json:"code"`
	// Search Parameter Name, or chained parameters
	Param []string `json:"param,omitempty"`
	// Additional documentation about the resource and compartment
	Documentation *string `json:"documentation,omitempty"`
	// Search Param for interpreting $everything.start
	StartParam *custom.URI `json:"startParam,omitempty"`
	// Search Param for interpreting $everything.end
	EndParam *custom.URI `json:"endParam,omitempty"`
}

type OtherCompartmentDefinition CompartmentDefinition

func (c CompartmentDefinition) ResourceType() string {
	return "CompartmentDefinition"
}

func (c CompartmentDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCompartmentDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherCompartmentDefinition: OtherCompartmentDefinition(c), ResourceType: c.ResourceType()})
}

func UnmarshalCompartmentDefinition(b []byte) (res CompartmentDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
