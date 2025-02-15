// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SearchParameter
type SearchParameter struct {
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
	// Canonical identifier for this search parameter, represented as a URI (globally unique)
	Url custom.URI `bson:"url" json:"url"`
	// Additional identifier for the search parameter (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the search parameter
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this search parameter (computer friendly)
	Name string `bson:"name" json:"name"`
	// Name for this search parameter (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Original definition for the search parameter
	DerivedFrom *custom.Canonical[SearchParameter] `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
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
	// Natural language description of the search parameter
	Description custom.Markdown `bson:"description" json:"description"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for search parameter (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this search parameter is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Recommended name for parameter in search url
	Code custom.Code `bson:"code" json:"code"`
	// The resource type(s) this search parameter applies to
	Base []custom.Code `bson:"base" json:"base"`
	// number | date | string | token | reference | composite | quantity | uri | special
	Type custom.Code `bson:"type" json:"type"`
	// FHIRPath expression that extracts the values
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty"`
	// normal | phonetic | other
	ProcessingMode *custom.Code `bson:"processingMode,omitempty" json:"processingMode,omitempty"`
	// FHIRPath expression that constraints the usage of this SearchParamete
	Constraint *string `bson:"constraint,omitempty" json:"constraint,omitempty"`
	// Types of resource (if a resource reference)
	Target []custom.Code `bson:"target,omitempty" json:"target,omitempty"`
	// Allow multiple values per parameter (or)
	MultipleOr *bool `bson:"multipleOr,omitempty" json:"multipleOr,omitempty"`
	// Allow multiple parameters (and)
	MultipleAnd *bool `bson:"multipleAnd,omitempty" json:"multipleAnd,omitempty"`
	// eq | ne | gt | lt | ge | le | sa | eb | ap
	Comparator []custom.Code `bson:"comparator,omitempty" json:"comparator,omitempty"`
	// missing | exact | contains | not | text | in | not-in | below | above | type | identifier | of-type | code-text | text-advanced | iterate
	Modifier []custom.Code `bson:"modifier,omitempty" json:"modifier,omitempty"`
	// Chained names supported
	Chain []string `bson:"chain,omitempty" json:"chain,omitempty"`
	// For Composite resources to define the parts
	Component []SearchParameterComponent `bson:"component,omitempty" json:"component,omitempty"`
}

type SearchParameterComponent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Defines how the part works
	Definition custom.Canonical[SearchParameter] `bson:"definition" json:"definition"`
	// Subexpression relative to main expression
	Expression string `bson:"expression" json:"expression"`
}

func (s SearchParameter) ResourceType() string {
	return "StructureDefinition"
}
