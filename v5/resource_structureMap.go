// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/StructureMap
type StructureMapGroupRuleTarget struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Variable this rule applies to
	Context *string `bson:"context,omitempty" json:"context,omitempty"`
	// Field to create in the context
	Element *string `bson:"element,omitempty" json:"element,omitempty"`
	// Named context for field, if desired, and a field is specified
	Variable *custom.ID `bson:"variable,omitempty" json:"variable,omitempty"`
	// first | share | last | single
	ListMode []custom.Code `bson:"listMode,omitempty" json:"listMode,omitempty"`
	// Internal rule reference for shared list items
	ListRuleId *custom.ID `bson:"listRuleId,omitempty" json:"listRuleId,omitempty"`
	// create | copy +
	Transform *custom.Code `bson:"transform,omitempty" json:"transform,omitempty"`
	// Parameters to the transform
	Parameter []StructureMapGroupRuleTargetParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
}

type StructureMapStructure struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Canonical reference to structure definition
	Url custom.Canonical[StructureDefinition] `bson:"url" json:"url"`
	// source | queried | target | produced
	Mode custom.Code `bson:"mode" json:"mode"`
	// Name for type in this map
	Alias *string `bson:"alias,omitempty" json:"alias,omitempty"`
	// Documentation on use of structure
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type StructureMapGroup struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Human-readable label
	Name custom.ID `bson:"name" json:"name"`
	// Another group that this group adds rules to
	Extends *custom.ID `bson:"extends,omitempty" json:"extends,omitempty"`
	// types | type-and-types
	TypeMode *custom.Code `bson:"typeMode,omitempty" json:"typeMode,omitempty"`
	// Additional description/explanation for group
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// Named instance provided when invoking the map
	Input []StructureMapGroupInput `bson:"input" json:"input"`
	// Transform Rule from source to target
	Rule []StructureMapGroupRule `bson:"rule,omitempty" json:"rule,omitempty"`
}

type StructureMapGroupInput struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name for this instance of data
	Name custom.ID `bson:"name" json:"name"`
	// Type for this instance of data
	Type *string `bson:"type,omitempty" json:"type,omitempty"`
	// source | target
	Mode custom.Code `bson:"mode" json:"mode"`
	// Documentation for this instance of data
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type StructureMapGroupRule struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name of the rule for internal references
	Name *custom.ID `bson:"name,omitempty" json:"name,omitempty"`
	// Source inputs to the mapping
	Source []StructureMapGroupRuleSource `bson:"source" json:"source"`
	// Content to create because of this mapping rule
	Target []StructureMapGroupRuleTarget `bson:"target,omitempty" json:"target,omitempty"`
	// Rules contained in this rule
	Rule []interface{} `bson:"rule,omitempty" json:"rule,omitempty"`
	// Which other rules to apply in the context of this rule
	Dependent []StructureMapGroupRuleDependent `bson:"dependent,omitempty" json:"dependent,omitempty"`
	// Documentation for this instance of data
	Documentation *string `bson:"documentation,omitempty" json:"documentation,omitempty"`
}

type StructureMapGroupRuleSource struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type or variable this rule applies to
	Context custom.ID `bson:"context" json:"context"`
	// Specified minimum cardinality
	Min *int32 `bson:"min,omitempty" json:"min,omitempty"`
	// Specified maximum cardinality (number or *)
	Max *string `bson:"max,omitempty" json:"max,omitempty"`
	// Rule only applies if source has this type
	Type *string `bson:"type,omitempty" json:"type,omitempty"`
	// Default value if no value exists
	DefaultValue *string `bson:"defaultValue,omitempty" json:"defaultValue,omitempty"`
	// Optional field for this source
	Element *string `bson:"element,omitempty" json:"element,omitempty"`
	// first | not_first | last | not_last | only_one
	ListMode *custom.Code `bson:"listMode,omitempty" json:"listMode,omitempty"`
	// Named context for field, if a field is specified
	Variable *custom.ID `bson:"variable,omitempty" json:"variable,omitempty"`
	// FHIRPath expression  - must be true or the rule does not apply
	Condition *string `bson:"condition,omitempty" json:"condition,omitempty"`
	// FHIRPath expression  - must be true or the mapping engine throws an error instead of completing
	Check *string `bson:"check,omitempty" json:"check,omitempty"`
	// Message to put in log if source exists (FHIRPath)
	LogMessage *string `bson:"logMessage,omitempty" json:"logMessage,omitempty"`
}

type StructureMapGroupRuleTargetParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Parameter value - variable or literal
	Value custom.ID `bson:"value" json:"value"`
}

type StructureMapGroupRuleDependent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name of a rule or group to apply
	Name custom.ID `bson:"name" json:"name"`
	// Parameter to pass to the rule or group
	Parameter []interface{} `bson:"parameter" json:"parameter"`
}

type StructureMap struct {
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
	// Canonical identifier for this structure map, represented as a URI (globally unique)
	Url custom.URI `bson:"url" json:"url"`
	// Additional identifier for the structure map
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the structure map
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this structure map (computer friendly)
	Name string `bson:"name" json:"name"`
	// Name for this structure map (human friendly)
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
	// Natural language description of the structure map
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for structure map (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this structure map is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// Structure Definition used by this map
	Structure []StructureMapStructure `bson:"structure,omitempty" json:"structure,omitempty"`
	// Other maps used by this map (canonical URLs)
	Import []custom.Canonical[StructureMap] `bson:"import,omitempty" json:"import,omitempty"`
	// Definition of the constant value used in the map rules
	Const []StructureMapConst `bson:"const,omitempty" json:"const,omitempty"`
	// Named sections for reader convenience
	Group []StructureMapGroup `bson:"group" json:"group"`
}

type StructureMapConst struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Constant name
	Name *custom.ID `bson:"name,omitempty" json:"name,omitempty"`
	// FHIRPath exression - value of the constant
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
}

func (s StructureMap) ResourceType() string {
	return "StructureDefinition"
}
