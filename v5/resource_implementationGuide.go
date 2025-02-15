// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ImplementationGuide
type ImplementationGuideManifestResource struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Location of the resource
	Reference custom.Reference[Resource] `bson:"reference" json:"reference"`
	// Is this an example
	IsExample *bool `bson:"isExample,omitempty" json:"isExample,omitempty"`
	// Profile(s) this is an example of
	Profile []custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	// Relative path for page in IG
	RelativePath *custom.URL `bson:"relativePath,omitempty" json:"relativePath,omitempty"`
}

type ImplementationGuideManifestPage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// HTML page name
	Name string `bson:"name" json:"name"`
	// Title of the page, for references
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Anchor available on the page
	Anchor []string `bson:"anchor,omitempty" json:"anchor,omitempty"`
}

type ImplementationGuideDependsOn struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identity of the IG that this depends on
	Uri custom.Canonical[ImplementationGuide] `bson:"uri" json:"uri"`
	// NPM Package name for IG this depends on
	PackageId *custom.ID `bson:"packageId,omitempty" json:"packageId,omitempty"`
	// Version of the IG
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// Why dependency exists
	Reason *custom.Markdown `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ImplementationGuideDefinition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Grouping used to present related resources in the IG
	Grouping []ImplementationGuideDefinitionGrouping `bson:"grouping,omitempty" json:"grouping,omitempty"`
	// Resource in the implementation guide
	Resource []ImplementationGuideDefinitionResource `bson:"resource,omitempty" json:"resource,omitempty"`
	// Page/Section in the Guide
	Page *ImplementationGuideDefinitionPage `bson:"page,omitempty" json:"page,omitempty"`
	// Defines how IG is built by tools
	Parameter []ImplementationGuideDefinitionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	// A template for building resources
	Template []ImplementationGuideDefinitionTemplate `bson:"template,omitempty" json:"template,omitempty"`
}

type ImplementationGuideDefinitionGrouping struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Descriptive name for the package
	Name string `bson:"name" json:"name"`
	// Human readable text describing the package
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
}

type ImplementationGuideDefinitionPage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Source for page
	Source *custom.URL `bson:"source,omitempty" json:"source,omitempty"`
	// Name of the page when published
	Name custom.URL `bson:"name" json:"name"`
	// Short title shown for navigational assistance
	Title string `bson:"title" json:"title"`
	// html | markdown | xml | generated
	Generation custom.Code `bson:"generation" json:"generation"`
	// Nested Pages / Sections
	Page []interface{} `bson:"page,omitempty" json:"page,omitempty"`
}

type ImplementationGuideDefinitionParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code that identifies parameter
	Code Coding `bson:"code" json:"code"`
	// Value for named type
	Value string `bson:"value" json:"value"`
}

type ImplementationGuideDefinitionTemplate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of template specified
	Code custom.Code `bson:"code" json:"code"`
	// The source location for the template
	Source string `bson:"source" json:"source"`
	// The scope in which the template applies
	Scope *string `bson:"scope,omitempty" json:"scope,omitempty"`
}

type ImplementationGuideManifest struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Location of rendered implementation guide
	Rendering *custom.URL `bson:"rendering,omitempty" json:"rendering,omitempty"`
	// Resource in the implementation guide
	Resource []ImplementationGuideManifestResource `bson:"resource" json:"resource"`
	// HTML page within the parent IG
	Page []ImplementationGuideManifestPage `bson:"page,omitempty" json:"page,omitempty"`
	// Image within the IG
	Image []string `bson:"image,omitempty" json:"image,omitempty"`
	// Additional linkable file in IG
	Other []string `bson:"other,omitempty" json:"other,omitempty"`
}

type ImplementationGuide struct {
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
	// Canonical identifier for this implementation guide, represented as a URI (globally unique)
	Url custom.URI `bson:"url" json:"url"`
	// Additional identifier for the implementation guide (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the implementation guide
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this implementation guide (computer friendly)
	Name string `bson:"name" json:"name"`
	// Name for this implementation guide (human friendly)
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
	// Natural language description of the implementation guide
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for implementation guide (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this implementation guide is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// NPM Package name for IG
	PackageId custom.ID `bson:"packageId" json:"packageId"`
	// SPDX license code for this IG (or not-open-source)
	License *custom.Code `bson:"license,omitempty" json:"license,omitempty"`
	// FHIR Version(s) this Implementation Guide targets
	FhirVersion []custom.Code `bson:"fhirVersion" json:"fhirVersion"`
	// Another Implementation guide this depends on
	DependsOn []ImplementationGuideDependsOn `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	// Profiles that apply globally
	Global []ImplementationGuideGlobal `bson:"global,omitempty" json:"global,omitempty"`
	// Information needed to build the IG
	Definition *ImplementationGuideDefinition `bson:"definition,omitempty" json:"definition,omitempty"`
	// Information about an assembled IG
	Manifest *ImplementationGuideManifest `bson:"manifest,omitempty" json:"manifest,omitempty"`
}

type ImplementationGuideGlobal struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type this profile applies to
	Type custom.Code `bson:"type" json:"type"`
	// Profile that all resources must conform to
	Profile custom.Canonical[StructureDefinition] `bson:"profile" json:"profile"`
}

type ImplementationGuideDefinitionResource struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Location of the resource
	Reference custom.Reference[Resource] `bson:"reference" json:"reference"`
	// Versions this applies to (if different to IG)
	FhirVersion []custom.Code `bson:"fhirVersion,omitempty" json:"fhirVersion,omitempty"`
	// Human readable name for the resource
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Reason why included in guide
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Is this an example
	IsExample *bool `bson:"isExample,omitempty" json:"isExample,omitempty"`
	// Profile(s) this is an example of
	Profile []custom.Canonical[StructureDefinition] `bson:"profile,omitempty" json:"profile,omitempty"`
	// Grouping this is part of
	GroupingId *custom.ID `bson:"groupingId,omitempty" json:"groupingId,omitempty"`
}

func (i ImplementationGuide) ResourceType() string {
	return "StructureDefinition"
}
