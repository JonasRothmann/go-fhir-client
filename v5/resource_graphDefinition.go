// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/GraphDefinition
type GraphDefinition struct {
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
	// Canonical identifier for this graph definition, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the GraphDefinition (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the graph definition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this graph definition (computer friendly)
	Name string `bson:"name" json:"name"`
	// Name for this graph definition (human friendly)
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
	// Natural language description of the graph definition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for graph definition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this graph definition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Starting Node
	Start *custom.ID `bson:"start,omitempty" json:"start,omitempty"`
	// Potential target for the link
	Node []GraphDefinitionNode `bson:"node,omitempty" json:"node,omitempty"`
	// Links this graph makes rules about
	Link []GraphDefinitionLink `bson:"link,omitempty" json:"link,omitempty"`
}

type GraphDefinitionNode struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Internal ID - target for link references
	NodeId custom.ID `bson:"nodeId" json:"nodeId"`
	// Why this node is specified
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Type of resource this link refers to
	Type custom.Code `bson:"type" json:"type"`
	// Profile for the target resource
	Profile *custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
}

type GraphDefinitionLink struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Why this link is specified
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Minimum occurrences for this link
	Min *int32 `bson:"min,omitempty" json:"min,omitempty"`
	// Maximum occurrences for this link
	Max *string `bson:"max,omitempty" json:"max,omitempty"`
	// Source Node for this link
	SourceId custom.ID `bson:"sourceId" json:"sourceId"`
	// Path in the resource that contains the link
	Path *string `bson:"path,omitempty" json:"path,omitempty"`
	// Which slice (if profiled)
	SliceName *string `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	// Target Node for this link
	TargetId custom.ID `bson:"targetId" json:"targetId"`
	// Criteria for reverse lookup
	Params *string `bson:"params,omitempty" json:"params,omitempty"`
	// Compartment Consistency Rules
	Compartment []GraphDefinitionLinkCompartment `bson:"compartment,omitempty" json:"compartment,omitempty"`
}

type GraphDefinitionLinkCompartment struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// where | requires
	Use custom.Code `bson:"use" json:"use"`
	// identical | matching | different | custom
	Rule custom.Code `bson:"rule" json:"rule"`
	// Patient | Encounter | RelatedPerson | Practitioner | Device | EpisodeOfCare
	Code custom.Code `bson:"code" json:"code"`
	// Custom rule, as a FHIRPath expression
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty"`
	// Documentation for FHIRPath expression
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
}

func (g GraphDefinition) ResourceType() string {
	return "StructureDefinition"
}
