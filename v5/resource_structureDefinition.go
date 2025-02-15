// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/StructureDefinition
type StructureDefinitionContext struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// fhirpath | element | extension
	Type custom.Code `bson:"type" json:"type"`
	// Where the extension can be used in instances
	Expression string `bson:"expression" json:"expression"`
}

type StructureDefinitionSnapshot struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Definition of elements in the resource (if no StructureDefinition)
	Element []ElementDefinition `bson:"element" json:"element"`
}

type StructureDefinitionDifferential struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Definition of elements in the resource (if no StructureDefinition)
	Element []ElementDefinition `bson:"element" json:"element"`
}

type StructureDefinition struct {
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
	// Canonical identifier for this structure definition, represented as a URI (globally unique)
	Url custom.URI `bson:"url" json:"url"`
	// Additional identifier for the structure definition
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the structure definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this structure definition (computer friendly)
	Name string `bson:"name" json:"name"`
	// Name for this structure definition (human friendly)
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
	// Natural language description of the structure definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for structure definition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this structure definition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Assist with indexing and finding
	Keyword []Coding `bson:"keyword,omitempty" json:"keyword,omitempty"`
	// FHIR Version this StructureDefinition targets
	FhirVersion *custom.Code `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	// External specification that the content is mapped to
	Mapping []StructureDefinitionMapping `bson:"mapping,omitempty" json:"mapping,omitempty"`
	// primitive-type | complex-type | resource | logical
	Kind custom.Code `bson:"kind" json:"kind"`
	// Whether the structure is abstract
	Abstract bool `bson:"abstract" json:"abstract"`
	// If an extension, where it can be used in instances
	Context []StructureDefinitionContext `bson:"context,omitempty" json:"context,omitempty"`
	// FHIRPath invariants - when the extension can be used
	ContextInvariant []string `bson:"contextInvariant,omitempty" json:"contextInvariant,omitempty"`
	// Type defined or constrained by this structure
	Type custom.URI `bson:"type" json:"type"`
	// Definition that this type is constrained/specialized from
	BaseDefinition *custom.Canonical[StructureDefinition] `bson:"baseDefinition,omitempty" json:"baseDefinition,omitempty"`
	// specialization | constraint - How relates to base definition
	Derivation *custom.Code `bson:"derivation,omitempty" json:"derivation,omitempty"`
	// Snapshot view of the structure
	Snapshot *StructureDefinitionSnapshot `bson:"snapshot,omitempty" json:"snapshot,omitempty"`
	// Differential view of the structure
	Differential *StructureDefinitionDifferential `bson:"differential,omitempty" json:"differential,omitempty"`
}

type StructureDefinitionMapping struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Internal id when this mapping is used
	Identity custom.ID `bson:"identity" json:"identity"`
	// Identifies what this mapping refers to
	Uri *custom.URI `bson:"uri,omitempty" json:"uri,omitempty"`
	// Names what this mapping refers to
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Versions, Issues, Scope limitations etc
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty"`
}

func (s StructureDefinition) ResourceType() string {
	return "StructureDefinition"
}
