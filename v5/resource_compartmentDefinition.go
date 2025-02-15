// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/CompartmentDefinition
type CompartmentDefinition struct {
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
	// Canonical identifier for this compartment definition, represented as a URI (globally unique)
	Url custom.URI `bson:"url" json:"url"`
	// Business version of the compartment definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this compartment definition (computer friendly)
	Name string `bson:"name" json:"name"`
	// Name for this compartment definition (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the compartment definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Why this compartment definition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Patient | Encounter | RelatedPerson | Practitioner | Device | EpisodeOfCare
	Code custom.Code `bson:"code" json:"code"`
	// Whether the search syntax is supported
	Search bool `bson:"search" json:"search"`
	// How a resource is related to the compartment
	Resource []CompartmentDefinitionResource `bson:"resource,omitempty" json:"resource,omitempty"`
}

type CompartmentDefinitionResource struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name of resource type
	Code custom.Code `bson:"code" json:"code"`
	// Search Parameter Name, or chained parameters
	Param []string `bson:"param,omitempty" json:"param,omitempty"`
	// Additional documentation about the resource and compartment
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// Search Param for interpreting $everything.start
	StartParam *custom.URI `bson:"startParam,omitempty" json:"startParam,omitempty"`
	// Search Param for interpreting $everything.end
	EndParam *custom.URI `bson:"endParam,omitempty" json:"endParam,omitempty"`
}

func (c CompartmentDefinition) ResourceType() string {
	return "StructureDefinition"
}
