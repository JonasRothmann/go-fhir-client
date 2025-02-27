// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/TerminologyCapabilities
type TerminologyCapabilitiesCodeSystem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Canonical identifier for the code system, represented as a URI
	Uri *custom.Canonical[CodeSystem] `json:"uri,omitempty"`
	// Version of Code System supported
	Version []TerminologyCapabilitiesCodeSystemVersion `json:"version,omitempty"`
	// not-present | example | fragment | complete | supplement
	Content custom.Code `json:"content"`
	// Whether subsumption is supported
	Subsumption *bool `json:"subsumption,omitempty"`
}

type TerminologyCapabilitiesCodeSystemVersionFilter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code of the property supported
	Code custom.Code `json:"code"`
	// Operations supported for the property
	Op []custom.Code `json:"op"`
}

type TerminologyCapabilitiesExpansionParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name of the supported expansion parameter
	Name custom.Code `json:"name"`
	// Description of support for parameter
	Documentation *string `json:"documentation,omitempty"`
}

type TerminologyCapabilitiesTranslation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Whether the client must identify the map
	NeedsMap bool `json:"needsMap"`
}

type TerminologyCapabilities struct {
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
	// Canonical identifier for this terminology capabilities, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the terminology capabilities
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the terminology capabilities
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this terminology capabilities (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this terminology capabilities (human friendly)
	Title *string `json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date custom.DateTime `json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the terminology capabilities
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for terminology capabilities (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this terminology capabilities is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// instance | capability | requirements
	Kind custom.Code `json:"kind"`
	// Software that is covered by this terminology capability statement
	Software *TerminologyCapabilitiesSoftware `json:"software,omitempty"`
	// If this describes a specific instance
	Implementation *TerminologyCapabilitiesImplementation `json:"implementation,omitempty"`
	// Whether lockedDate is supported
	LockedDate *bool `json:"lockedDate,omitempty"`
	// A code system supported by the server
	CodeSystem []TerminologyCapabilitiesCodeSystem `json:"codeSystem,omitempty"`
	// Information about the [ValueSet/$expand](valueset-operation-expand.html) operation
	Expansion *TerminologyCapabilitiesExpansion `json:"expansion,omitempty"`
	// in-compose | in-expansion | in-compose-or-expansion
	CodeSearch *custom.Code `json:"codeSearch,omitempty"`
	// Information about the [ValueSet/$validate-code](valueset-operation-validate-code.html) operation
	ValidateCode *TerminologyCapabilitiesValidateCode `json:"validateCode,omitempty"`
	// Information about the [ConceptMap/$translate](conceptmap-operation-translate.html) operation
	Translation *TerminologyCapabilitiesTranslation `json:"translation,omitempty"`
	// Information about the [ConceptMap/$closure](conceptmap-operation-closure.html) operation
	Closure *TerminologyCapabilitiesClosure `json:"closure,omitempty"`
}

type TerminologyCapabilitiesImplementation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Describes this specific instance
	Description string `json:"description"`
	// Base URL for the implementation
	Url *custom.URL `json:"url,omitempty"`
}

type TerminologyCapabilitiesCodeSystemVersion struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Version identifier for this version
	Code *string `json:"code,omitempty"`
	// If this is the default version for this code system
	IsDefault *bool `json:"isDefault,omitempty"`
	// If compositional grammar is supported
	Compositional *bool `json:"compositional,omitempty"`
	// Language Displays supported
	Language []custom.Code `json:"language,omitempty"`
	// Filter Properties supported
	Filter []TerminologyCapabilitiesCodeSystemVersionFilter `json:"filter,omitempty"`
	// Properties supported for $lookup
	Property []custom.Code `json:"property,omitempty"`
}

type TerminologyCapabilitiesExpansion struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Whether the server can return nested value sets
	Hierarchical *bool `json:"hierarchical,omitempty"`
	// Whether the server supports paging on expansion
	Paging *bool `json:"paging,omitempty"`
	// Allow request for incomplete expansions?
	Incomplete *bool `json:"incomplete,omitempty"`
	// Supported expansion parameter
	Parameter []TerminologyCapabilitiesExpansionParameter `json:"parameter,omitempty"`
	// Documentation about text searching works
	TextFilter *custom.Markdown `json:"textFilter,omitempty"`
}

type TerminologyCapabilitiesValidateCode struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Whether translations are validated
	Translations bool `json:"translations"`
}

type TerminologyCapabilitiesClosure struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// If cross-system closure is supported
	Translation *bool `json:"translation,omitempty"`
}

type TerminologyCapabilitiesSoftware struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A name the software is known by
	Name string `json:"name"`
	// Version covered by this statement
	Version *string `json:"version,omitempty"`
}

type OtherTerminologyCapabilities TerminologyCapabilities

func (t TerminologyCapabilities) ResourceType() string {
	return "TerminologyCapabilities"
}

func (t TerminologyCapabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTerminologyCapabilities
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherTerminologyCapabilities: OtherTerminologyCapabilities(t), ResourceType: t.ResourceType()})
}

func UnmarshalTerminologyCapabilities(b []byte) (res TerminologyCapabilities, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
