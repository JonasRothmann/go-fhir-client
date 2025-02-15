// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilities struct {
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
	// Canonical identifier for this terminology capabilities, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the terminology capabilities
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the terminology capabilities
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this terminology capabilities (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this terminology capabilities (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date custom.DateTime `bson:"date" json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the terminology capabilities
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for terminology capabilities (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this terminology capabilities is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// instance | capability | requirements
	Kind custom.Code `bson:"kind" json:"kind"`
	// Software that is covered by this terminology capability statement
	Software *TerminologyCapabilitiesSoftware `bson:"software,omitempty" json:"software,omitempty"`
	// If this describes a specific instance
	Implementation *TerminologyCapabilitiesImplementation `bson:"implementation,omitempty" json:"implementation,omitempty"`
	// Whether lockedDate is supported
	LockedDate *bool `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	// A code system supported by the server
	CodeSystem []TerminologyCapabilitiesCodeSystem `bson:"codeSystem,omitempty" json:"codeSystem,omitempty"`
	// Information about the [ValueSet/$expand](valueset-operation-expand.html) operation
	Expansion *TerminologyCapabilitiesExpansion `bson:"expansion,omitempty" json:"expansion,omitempty"`
	// in-compose | in-expansion | in-compose-or-expansion
	CodeSearch *custom.Code `bson:"codeSearch,omitempty" json:"codeSearch,omitempty"`
	// Information about the [ValueSet/$validate-code](valueset-operation-validate-code.html) operation
	ValidateCode *TerminologyCapabilitiesValidateCode `bson:"validateCode,omitempty" json:"validateCode,omitempty"`
	// Information about the [ConceptMap/$translate](conceptmap-operation-translate.html) operation
	Translation *TerminologyCapabilitiesTranslation `bson:"translation,omitempty" json:"translation,omitempty"`
	// Information about the [ConceptMap/$closure](conceptmap-operation-closure.html) operation
	Closure *TerminologyCapabilitiesClosure `bson:"closure,omitempty" json:"closure,omitempty"`
}

type TerminologyCapabilitiesImplementation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Describes this specific instance
	Description string `bson:"description" json:"description"`
	// Base URL for the implementation
	Url *custom.URL `bson:"url,omitempty" json:"url,omitempty"`
}

type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code of the property supported
	Code custom.Code `bson:"code" json:"code"`
	// Operations supported for the property
	Op []custom.Code `bson:"op" json:"op"`
}

type TerminologyCapabilitiesValidateCode struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Whether translations are validated
	Translations bool `bson:"translations" json:"translations"`
}

type TerminologyCapabilitiesTranslation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Whether the client must identify the map
	NeedsMap bool `bson:"needsMap" json:"needsMap"`
}

type TerminologyCapabilitiesClosure struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// If cross-system closure is supported
	Translation *bool `bson:"translation,omitempty" json:"translation,omitempty"`
}

type TerminologyCapabilitiesSoftware struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A name the software is known by
	Name string `bson:"name" json:"name"`
	// Version covered by this statement
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
}

type TerminologyCapabilitiesCodeSystem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Canonical identifier for the code system, represented as a URI
	Uri *custom.Canonical[CodeSystem] `bson:"uri,omitempty" json:"uri,omitempty"`
	// Version of Code System supported
	Version []TerminologyCapabilitiesCodeSystemVersion `bson:"version,omitempty" json:"version,omitempty"`
	// not-present | example | fragment | complete | supplement
	Content custom.Code `bson:"content" json:"content"`
	// Whether subsumption is supported
	Subsumption *bool `bson:"subsumption,omitempty" json:"subsumption,omitempty"`
}

type TerminologyCapabilitiesCodeSystemVersion struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Version identifier for this version
	Code *string `bson:"code,omitempty" json:"code,omitempty"`
	// If this is the default version for this code system
	IsDefault *bool `bson:"isDefault,omitempty" json:"isDefault,omitempty"`
	// If compositional grammar is supported
	Compositional *bool `bson:"compositional,omitempty" json:"compositional,omitempty"`
	// Language Displays supported
	Language []custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Filter Properties supported
	Filter []TerminologyCapabilitiesCodeSystemVersionFilter `bson:"filter,omitempty" json:"filter,omitempty"`
	// Properties supported for $lookup
	Property []custom.Code `bson:"property,omitempty" json:"property,omitempty"`
}

type TerminologyCapabilitiesExpansion struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Whether the server can return nested value sets
	Hierarchical *bool `bson:"hierarchical,omitempty" json:"hierarchical,omitempty"`
	// Whether the server supports paging on expansion
	Paging *bool `bson:"paging,omitempty" json:"paging,omitempty"`
	// Allow request for incomplete expansions?
	Incomplete *bool `bson:"incomplete,omitempty" json:"incomplete,omitempty"`
	// Supported expansion parameter
	Parameter []TerminologyCapabilitiesExpansionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	// Documentation about text searching works
	TextFilter *custom.Markdown `bson:"textFilter,omitempty" json:"textFilter,omitempty"`
}

type TerminologyCapabilitiesExpansionParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name of the supported expansion parameter
	Name custom.Code `bson:"name" json:"name"`
	// Description of support for parameter
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

func (t TerminologyCapabilities) ResourceType() string {
	return "StructureDefinition"
}
