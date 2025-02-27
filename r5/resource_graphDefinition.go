// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/GraphDefinition
type GraphDefinition struct {
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
	// Canonical identifier for this graph definition, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the GraphDefinition (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the graph definition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this graph definition (computer friendly)
	Name string `json:"name"`
	// Name for this graph definition (human friendly)
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
	// Natural language description of the graph definition
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for graph definition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this graph definition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Starting Node
	Start *custom.ID `json:"start,omitempty"`
	// Potential target for the link
	Node []GraphDefinitionNode `json:"node,omitempty"`
	// Links this graph makes rules about
	Link []GraphDefinitionLink `json:"link,omitempty"`
}

type GraphDefinitionNode struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Internal ID - target for link references
	NodeId custom.ID `json:"nodeId"`
	// Why this node is specified
	Description *string `json:"description,omitempty"`
	// Type of resource this link refers to
	Type custom.Code `json:"type"`
	// Profile for the target resource
	Profile *custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
}

type GraphDefinitionLink struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Why this link is specified
	Description *string `json:"description,omitempty"`
	// Minimum occurrences for this link
	Min *int32 `json:"min,omitempty"`
	// Maximum occurrences for this link
	Max *string `json:"max,omitempty"`
	// Source Node for this link
	SourceId custom.ID `json:"sourceId"`
	// Path in the resource that contains the link
	Path *string `json:"path,omitempty"`
	// Which slice (if profiled)
	SliceName *string `json:"sliceName,omitempty"`
	// Target Node for this link
	TargetId custom.ID `json:"targetId"`
	// Criteria for reverse lookup
	Params *string `json:"params,omitempty"`
	// Compartment Consistency Rules
	Compartment []GraphDefinitionLinkCompartment `json:"compartment,omitempty"`
}

type GraphDefinitionLinkCompartment struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// where | requires
	Use custom.Code `json:"use"`
	// identical | matching | different | custom
	Rule custom.Code `json:"rule"`
	// Patient | Encounter | RelatedPerson | Practitioner | Device | EpisodeOfCare
	Code custom.Code `json:"code"`
	// Custom rule, as a FHIRPath expression
	Expression *string `json:"expression,omitempty"`
	// Documentation for FHIRPath expression
	Description *string `json:"description,omitempty"`
}

type OtherGraphDefinition GraphDefinition

func (g GraphDefinition) ResourceType() string {
	return "GraphDefinition"
}

func (g GraphDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGraphDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherGraphDefinition: OtherGraphDefinition(g), ResourceType: g.ResourceType()})
}

func UnmarshalGraphDefinition(b []byte) (res GraphDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
