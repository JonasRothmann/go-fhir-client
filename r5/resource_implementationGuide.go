// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ImplementationGuide
type ImplementationGuideDefinitionPage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Source for page
	SourceUrl *custom.URL `json:"sourceUrl,omitempty"`
	// Source for page
	SourceString *string `json:"sourceString,omitempty"`
	// Source for page
	SourceMarkdown *custom.Markdown `json:"sourceMarkdown,omitempty"`
	// Name of the page when published
	Name custom.URL `json:"name"`
	// Short title shown for navigational assistance
	Title string `json:"title"`
	// html | markdown | xml | generated
	Generation custom.Code `json:"generation"`
	// Nested Pages / Sections
	Page []interface{} `json:"page,omitempty"`
}

type ImplementationGuideDefinitionTemplate struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of template specified
	Code custom.Code `json:"code"`
	// The source location for the template
	Source string `json:"source"`
	// The scope in which the template applies
	Scope *string `json:"scope,omitempty"`
}

type ImplementationGuideManifest struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Location of rendered implementation guide
	Rendering *custom.URL `json:"rendering,omitempty"`
	// Resource in the implementation guide
	Resource []ImplementationGuideManifestResource `json:"resource"`
	// HTML page within the parent IG
	Page []ImplementationGuideManifestPage `json:"page,omitempty"`
	// Image within the IG
	Image []string `json:"image,omitempty"`
	// Additional linkable file in IG
	Other []string `json:"other,omitempty"`
}

type ImplementationGuideManifestResource struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Location of the resource
	Reference custom.Reference[Resource] `json:"reference"`
	// Is this an example
	IsExample *bool `json:"isExample,omitempty"`
	// Profile(s) this is an example of
	Profile []custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
	// Relative path for page in IG
	RelativePath *custom.URL `json:"relativePath,omitempty"`
}

type ImplementationGuide struct {
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
	// Canonical identifier for this implementation guide, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Additional identifier for the implementation guide (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the implementation guide
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this implementation guide (computer friendly)
	Name string `json:"name"`
	// Name for this implementation guide (human friendly)
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
	// Natural language description of the implementation guide
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for implementation guide (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this implementation guide is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// NPM Package name for IG
	PackageId custom.ID `json:"packageId"`
	// SPDX license code for this IG (or not-open-source)
	License *custom.Code `json:"license,omitempty"`
	// FHIR Version(s) this Implementation Guide targets
	FhirVersion []custom.Code `json:"fhirVersion"`
	// Another Implementation guide this depends on
	DependsOn []ImplementationGuideDependsOn `json:"dependsOn,omitempty"`
	// Profiles that apply globally
	Global []ImplementationGuideGlobal `json:"global,omitempty"`
	// Information needed to build the IG
	Definition *ImplementationGuideDefinition `json:"definition,omitempty"`
	// Information about an assembled IG
	Manifest *ImplementationGuideManifest `json:"manifest,omitempty"`
}

type ImplementationGuideDependsOn struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identity of the IG that this depends on
	Uri custom.Canonical[ImplementationGuide] `json:"uri"`
	// NPM Package name for IG this depends on
	PackageId *custom.ID `json:"packageId,omitempty"`
	// Version of the IG
	Version *string `json:"version,omitempty"`
	// Why dependency exists
	Reason *custom.Markdown `json:"reason,omitempty"`
}

type ImplementationGuideGlobal struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type this profile applies to
	Type custom.Code `json:"type"`
	// Profile that all resources must conform to
	Profile custom.Canonical[StructureDefinition] `json:"profile"`
}

type ImplementationGuideDefinitionResource struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Location of the resource
	Reference custom.Reference[Resource] `json:"reference"`
	// Versions this applies to (if different to IG)
	FhirVersion []custom.Code `json:"fhirVersion,omitempty"`
	// Human readable name for the resource
	Name *string `json:"name,omitempty"`
	// Reason why included in guide
	Description *custom.Markdown `json:"description,omitempty"`
	// Is this an example
	IsExample *bool `json:"isExample,omitempty"`
	// Profile(s) this is an example of
	Profile []custom.Canonical[StructureDefinition] `json:"profile,omitempty"`
	// Grouping this is part of
	GroupingId *custom.ID `json:"groupingId,omitempty"`
}

type ImplementationGuideDefinition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Grouping used to present related resources in the IG
	Grouping []ImplementationGuideDefinitionGrouping `json:"grouping,omitempty"`
	// Resource in the implementation guide
	Resource []ImplementationGuideDefinitionResource `json:"resource,omitempty"`
	// Page/Section in the Guide
	Page *ImplementationGuideDefinitionPage `json:"page,omitempty"`
	// Defines how IG is built by tools
	Parameter []ImplementationGuideDefinitionParameter `json:"parameter,omitempty"`
	// A template for building resources
	Template []ImplementationGuideDefinitionTemplate `json:"template,omitempty"`
}

type ImplementationGuideDefinitionGrouping struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Descriptive name for the package
	Name string `json:"name"`
	// Human readable text describing the package
	Description *custom.Markdown `json:"description,omitempty"`
}

type ImplementationGuideDefinitionParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code that identifies parameter
	Code Coding `json:"code"`
	// Value for named type
	Value string `json:"value"`
}

type ImplementationGuideManifestPage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// HTML page name
	Name string `json:"name"`
	// Title of the page, for references
	Title *string `json:"title,omitempty"`
	// Anchor available on the page
	Anchor []string `json:"anchor,omitempty"`
}

type OtherImplementationGuide ImplementationGuide

func (i ImplementationGuide) ResourceType() string {
	return "ImplementationGuide"
}

func (i ImplementationGuide) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImplementationGuide
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherImplementationGuide: OtherImplementationGuide(i), ResourceType: i.ResourceType()})
}

func UnmarshalImplementationGuide(b []byte) (res ImplementationGuide, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
