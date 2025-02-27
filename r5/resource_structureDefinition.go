// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/StructureDefinition
type StructureDefinitionDifferential struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Definition of elements in the resource (if no StructureDefinition)
	Element []ElementDefinition `json:"element"`
}

type StructureDefinition struct {
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
	// Canonical identifier for this structure definition, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Additional identifier for the structure definition
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the structure definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this structure definition (computer friendly)
	Name string `json:"name"`
	// Name for this structure definition (human friendly)
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
	// Natural language description of the structure definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for structure definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this structure definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Assist with indexing and finding
	Keyword []Coding `json:"keyword,omitempty"`
	// FHIR Version this StructureDefinition targets
	FhirVersion *custom.Code `json:"fhirVersion,omitempty"`
	// External specification that the content is mapped to
	Mapping []StructureDefinitionMapping `json:"mapping,omitempty"`
	// primitive-type | complex-type | resource | logical
	Kind custom.Code `json:"kind"`
	// Whether the structure is abstract
	Abstract bool `json:"abstract"`
	// If an extension, where it can be used in instances
	Context []StructureDefinitionContext `json:"context,omitempty"`
	// FHIRPath invariants - when the extension can be used
	ContextInvariant []string `json:"contextInvariant,omitempty"`
	// Type defined or constrained by this structure
	Type custom.URI `json:"type"`
	// Definition that this type is constrained/specialized from
	BaseDefinition *custom.Canonical[StructureDefinition] `json:"baseDefinition,omitempty"`
	// specialization | constraint - How relates to base definition
	Derivation *custom.Code `json:"derivation,omitempty"`
	// Snapshot view of the structure
	Snapshot *StructureDefinitionSnapshot `json:"snapshot,omitempty"`
	// Differential view of the structure
	Differential *StructureDefinitionDifferential `json:"differential,omitempty"`
}

type StructureDefinitionMapping struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Internal id when this mapping is used
	Identity custom.ID `json:"identity"`
	// Identifies what this mapping refers to
	Uri *custom.URI `json:"uri,omitempty"`
	// Names what this mapping refers to
	Name *string `json:"name,omitempty"`
	// Versions, Issues, Scope limitations etc
	Comment *string `json:"comment,omitempty"`
}

type StructureDefinitionContext struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// fhirpath | element | extension
	Type custom.Code `json:"type"`
	// Where the extension can be used in instances
	Expression string `json:"expression"`
}

type StructureDefinitionSnapshot struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Definition of elements in the resource (if no StructureDefinition)
	Element []ElementDefinition `json:"element"`
}

type OtherStructureDefinition StructureDefinition

func (s StructureDefinition) ResourceType() string {
	return "StructureDefinition"
}

func (s StructureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherStructureDefinition: OtherStructureDefinition(s), ResourceType: s.ResourceType()})
}

func UnmarshalStructureDefinition(b []byte) (res StructureDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
