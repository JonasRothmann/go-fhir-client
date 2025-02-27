// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SearchParameter
type SearchParameter struct {
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
	// Canonical identifier for this search parameter, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Additional identifier for the search parameter (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the search parameter
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this search parameter (computer friendly)
	Name string `json:"name"`
	// Name for this search parameter (human friendly)
	Title *string `json:"title,omitempty"`
	// Original definition for the search parameter
	DerivedFrom *custom.Canonical[SearchParameter] `json:"derivedFrom,omitempty"`
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
	// Natural language description of the search parameter
	Description custom.Markdown `json:"description"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for search parameter (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this search parameter is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Recommended name for parameter in search url
	Code custom.Code `json:"code"`
	// The resource type(s) this search parameter applies to
	Base []custom.Code `json:"base"`
	// number | date | string | token | reference | composite | quantity | uri | special
	Type custom.Code `json:"type"`
	// FHIRPath expression that extracts the values
	Expression *string `json:"expression,omitempty"`
	// normal | phonetic | other
	ProcessingMode *custom.Code `json:"processingMode,omitempty"`
	// FHIRPath expression that constraints the usage of this SearchParamete
	Constraint *string `json:"constraint,omitempty"`
	// Types of resource (if a resource reference)
	Target []custom.Code `json:"target,omitempty"`
	// Allow multiple values per parameter (or)
	MultipleOr *bool `json:"multipleOr,omitempty"`
	// Allow multiple parameters (and)
	MultipleAnd *bool `json:"multipleAnd,omitempty"`
	// eq | ne | gt | lt | ge | le | sa | eb | ap
	Comparator []custom.Code `json:"comparator,omitempty"`
	// missing | exact | contains | not | text | in | not-in | below | above | type | identifier | of-type | code-text | text-advanced | iterate
	Modifier []custom.Code `json:"modifier,omitempty"`
	// Chained names supported
	Chain []string `json:"chain,omitempty"`
	// For Composite resources to define the parts
	Component []SearchParameterComponent `json:"component,omitempty"`
}

type SearchParameterComponent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Defines how the part works
	Definition custom.Canonical[SearchParameter] `json:"definition"`
	// Subexpression relative to main expression
	Expression string `json:"expression"`
}

type OtherSearchParameter SearchParameter

func (s SearchParameter) ResourceType() string {
	return "SearchParameter"
}

func (s SearchParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSearchParameter
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSearchParameter: OtherSearchParameter(s), ResourceType: s.ResourceType()})
}

func UnmarshalSearchParameter(b []byte) (res SearchParameter, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
