// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/StructureMap
type StructureMapGroupRule struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name of the rule for internal references
	Name *custom.ID `json:"name,omitempty"`
	// Source inputs to the mapping
	Source []StructureMapGroupRuleSource `json:"source"`
	// Content to create because of this mapping rule
	Target []StructureMapGroupRuleTarget `json:"target,omitempty"`
	// Rules contained in this rule
	Rule []interface{} `json:"rule,omitempty"`
	// Which other rules to apply in the context of this rule
	Dependent []StructureMapGroupRuleDependent `json:"dependent,omitempty"`
	// Documentation for this instance of data
	Documentation *string `json:"documentation,omitempty"`
}

type StructureMapGroupRuleSource struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type or variable this rule applies to
	Context custom.ID `json:"context"`
	// Specified minimum cardinality
	Min *int32 `json:"min,omitempty"`
	// Specified maximum cardinality (number or *)
	Max *string `json:"max,omitempty"`
	// Rule only applies if source has this type
	Type *string `json:"type,omitempty"`
	// Default value if no value exists
	DefaultValue *string `json:"defaultValue,omitempty"`
	// Optional field for this source
	Element *string `json:"element,omitempty"`
	// first | not_first | last | not_last | only_one
	ListMode *custom.Code `json:"listMode,omitempty"`
	// Named context for field, if a field is specified
	Variable *custom.ID `json:"variable,omitempty"`
	// FHIRPath expression  - must be true or the rule does not apply
	Condition *string `json:"condition,omitempty"`
	// FHIRPath expression  - must be true or the mapping engine throws an error instead of completing
	Check *string `json:"check,omitempty"`
	// Message to put in log if source exists (FHIRPath)
	LogMessage *string `json:"logMessage,omitempty"`
}

type StructureMapGroupRuleDependent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name of a rule or group to apply
	Name custom.ID `json:"name"`
	// Parameter to pass to the rule or group
	Parameter []interface{} `json:"parameter"`
}

type StructureMap struct {
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
	// Canonical identifier for this structure map, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Additional identifier for the structure map
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the structure map
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this structure map (computer friendly)
	Name string `json:"name"`
	// Name for this structure map (human friendly)
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
	// Natural language description of the structure map
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for structure map (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this structure map is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// Structure Definition used by this map
	Structure []StructureMapStructure `json:"structure,omitempty"`
	// Other maps used by this map (canonical URLs)
	Import []custom.Canonical[StructureMap] `json:"import,omitempty"`
	// Definition of the constant value used in the map rules
	Const []StructureMapConst `json:"const,omitempty"`
	// Named sections for reader convenience
	Group []StructureMapGroup `json:"group"`
}

type StructureMapStructure struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Canonical reference to structure definition
	Url custom.Canonical[StructureDefinition] `json:"url"`
	// source | queried | target | produced
	Mode custom.Code `json:"mode"`
	// Name for type in this map
	Alias *string `json:"alias,omitempty"`
	// Documentation on use of structure
	Documentation *string `json:"documentation,omitempty"`
}

type StructureMapConst struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Constant name
	Name *custom.ID `json:"name,omitempty"`
	// FHIRPath exression - value of the constant
	Value *string `json:"value,omitempty"`
}

type StructureMapGroupRuleTargetParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Parameter value - variable or literal
	ValueId *custom.ID `json:"valueId,omitempty"`
	// Parameter value - variable or literal
	ValueString *string `json:"valueString,omitempty"`
	// Parameter value - variable or literal
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Parameter value - variable or literal
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Parameter value - variable or literal
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Parameter value - variable or literal
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Parameter value - variable or literal
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Parameter value - variable or literal
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
}

type StructureMapGroup struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Human-readable label
	Name custom.ID `json:"name"`
	// Another group that this group adds rules to
	Extends *custom.ID `json:"extends,omitempty"`
	// types | type-and-types
	TypeMode *custom.Code `json:"typeMode,omitempty"`
	// Additional description/explanation for group
	Documentation *string `json:"documentation,omitempty"`
	// Named instance provided when invoking the map
	Input []StructureMapGroupInput `json:"input"`
	// Transform Rule from source to target
	Rule []StructureMapGroupRule `json:"rule,omitempty"`
}

type StructureMapGroupInput struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name for this instance of data
	Name custom.ID `json:"name"`
	// Type for this instance of data
	Type *string `json:"type,omitempty"`
	// source | target
	Mode custom.Code `json:"mode"`
	// Documentation for this instance of data
	Documentation *string `json:"documentation,omitempty"`
}

type StructureMapGroupRuleTarget struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Variable this rule applies to
	Context *string `json:"context,omitempty"`
	// Field to create in the context
	Element *string `json:"element,omitempty"`
	// Named context for field, if desired, and a field is specified
	Variable *custom.ID `json:"variable,omitempty"`
	// first | share | last | single
	ListMode []custom.Code `json:"listMode,omitempty"`
	// Internal rule reference for shared list items
	ListRuleId *custom.ID `json:"listRuleId,omitempty"`
	// create | copy +
	Transform *custom.Code `json:"transform,omitempty"`
	// Parameters to the transform
	Parameter []StructureMapGroupRuleTargetParameter `json:"parameter,omitempty"`
}

type OtherStructureMap StructureMap

func (s StructureMap) ResourceType() string {
	return "StructureMap"
}

func (s StructureMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherStructureMap
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherStructureMap: OtherStructureMap(s), ResourceType: s.ResourceType()})
}

func UnmarshalStructureMap(b []byte) (res StructureMap, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
