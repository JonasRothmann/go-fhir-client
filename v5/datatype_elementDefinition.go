// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ElementDefinition
type ElementDefinitionSlicing struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Element values that are used to distinguish the slices
	Discriminator []ElementDefinitionSlicingDiscriminator `bson:"discriminator,omitempty" json:"discriminator,omitempty"`
	// Text description of how slicing works (or not)
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// If elements must be in same order as slices
	Ordered *bool `bson:"ordered,omitempty" json:"ordered,omitempty"`
	// closed | open | openAtEnd
	Rules custom.Code `bson:"rules" json:"rules"`
}

type ElementDefinitionBase struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Path that identifies the base element
	Path string `bson:"path" json:"path"`
	// Min cardinality of the base element
	Min uint32 `bson:"min" json:"min"`
	// Max cardinality of the base element
	Max string `bson:"max" json:"max"`
}

type ElementDefinitionType struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Data type or Resource (reference to definition)
	Code custom.URI `bson:"code" json:"code"`
	// Profiles (StructureDefinition or IG) - one must apply
	Profile []custom.Canonical[ElementDefinitionTypeProfile] `bson:"profile,omitempty" json:"profile,omitempty"`
	// Profile (StructureDefinition or IG) on the Reference/canonical target - one must apply
	TargetProfile []custom.Canonical[ElementDefinitionTypeTargetProfile] `bson:"targetProfile,omitempty" json:"targetProfile,omitempty"`
	// contained | referenced | bundled - how aggregated
	Aggregation []custom.Code `bson:"aggregation,omitempty" json:"aggregation,omitempty"`
	// either | independent | specific
	Versioning *custom.Code `bson:"versioning,omitempty" json:"versioning,omitempty"`
}

type ElementDefinitionBindingAdditional struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// maximum | minimum | required | extensible | candidate | current | preferred | ui | starter | component
	Purpose custom.Code `bson:"purpose" json:"purpose"`
	// The value set for the additional binding
	ValueSet custom.Canonical[ValueSet] `bson:"valueSet" json:"valueSet"`
	// Documentation of the purpose of use of the binding
	Documentation *custom.Markdown `bson:"documentation,omitempty" json:"documentation,omitempty"`
	// Concise documentation - for summary tables
	ShortDoco *string `bson:"shortDoco,omitempty" json:"shortDoco,omitempty"`
	// Qualifies the usage - jurisdiction, gender, workflow status etc.
	Usage []UsageContext `bson:"usage,omitempty" json:"usage,omitempty"`
	// Whether binding can applies to all repeats, or just one
	Any *bool `bson:"any,omitempty" json:"any,omitempty"`
}

type ElementDefinition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Path of the element in the hierarchy of elements
	Path string `bson:"path" json:"path"`
	// xmlAttr | xmlText | typeAttr | cdaText | xhtml
	Representation []custom.Code `bson:"representation,omitempty" json:"representation,omitempty"`
	// Name for this particular element (in a set of slices)
	SliceName *string `bson:"sliceName,omitempty" json:"sliceName,omitempty"`
	// If this slice definition constrains an inherited slice definition (or not)
	SliceIsConstraining *bool `bson:"sliceIsConstraining,omitempty" json:"sliceIsConstraining,omitempty"`
	// Name for element to display with or prompt for element
	Label *string `bson:"label,omitempty" json:"label,omitempty"`
	// Corresponding codes in terminologies
	Code []Coding `bson:"code,omitempty" json:"code,omitempty"`
	// This element is sliced - slices follow
	Slicing *ElementDefinitionSlicing `bson:"slicing,omitempty" json:"slicing,omitempty"`
	// Concise definition for space-constrained presentation
	Short *string `bson:"short,omitempty" json:"short,omitempty"`
	// Full formal definition as narrative text
	Definition *custom.Markdown `bson:"definition,omitempty" json:"definition,omitempty"`
	// Comments about the use of this element
	Comment *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
	// Why this resource has been created
	Requirements *custom.Markdown `bson:"requirements,omitempty" json:"requirements,omitempty"`
	// Other names
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty"`
	// Minimum Cardinality
	Min *uint32 `bson:"min,omitempty" json:"min,omitempty"`
	// Maximum Cardinality (a number or *)
	Max *string `bson:"max,omitempty" json:"max,omitempty"`
	// Base definition information for tools
	Base *ElementDefinitionBase `bson:"base,omitempty" json:"base,omitempty"`
	// Reference to definition of content for the element
	ContentReference *custom.URI `bson:"contentReference,omitempty" json:"contentReference,omitempty"`
	// Data type and Profile for this element
	Type []ElementDefinitionType `bson:"type,omitempty" json:"type,omitempty"`
	// Specified value if missing from instance
	DefaultValue *custom.Base64binary `bson:"defaultValue,omitempty" json:"defaultValue,omitempty"`
	// Implicit meaning when this element is missing
	MeaningWhenMissing *custom.Markdown `bson:"meaningWhenMissing,omitempty" json:"meaningWhenMissing,omitempty"`
	// What the order of the elements means
	OrderMeaning *string `bson:"orderMeaning,omitempty" json:"orderMeaning,omitempty"`
	// Value must be exactly this
	Fixed *custom.Base64binary `bson:"fixed,omitempty" json:"fixed,omitempty"`
	// Value must have at least these property values
	Pattern *custom.Base64binary `bson:"pattern,omitempty" json:"pattern,omitempty"`
	// Example value (as defined for type)
	Example []ElementDefinitionExample `bson:"example,omitempty" json:"example,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValue *custom.Date `bson:"minValue,omitempty" json:"minValue,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValue *custom.Date `bson:"maxValue,omitempty" json:"maxValue,omitempty"`
	// Max length for string type data
	MaxLength *int32 `bson:"maxLength,omitempty" json:"maxLength,omitempty"`
	// Reference to invariant about presence
	Condition []custom.ID `bson:"condition,omitempty" json:"condition,omitempty"`
	// Condition that must evaluate to true
	Constraint []ElementDefinitionConstraint `bson:"constraint,omitempty" json:"constraint,omitempty"`
	// For primitives, that a value must be present - not replaced by an extension
	MustHaveValue *bool `bson:"mustHaveValue,omitempty" json:"mustHaveValue,omitempty"`
	// Extensions that are allowed to replace a primitive value
	ValueAlternatives []custom.Canonical[StructureDefinition] `bson:"valueAlternatives,omitempty" json:"valueAlternatives,omitempty"`
	// If the element must be supported (discouraged - see obligations)
	MustSupport *bool `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	// If this modifies the meaning of other elements
	IsModifier *bool `bson:"isModifier,omitempty" json:"isModifier,omitempty"`
	// Reason that this element is marked as a modifier
	IsModifierReason *string `bson:"isModifierReason,omitempty" json:"isModifierReason,omitempty"`
	// Include when _summary = true?
	IsSummary *bool `bson:"isSummary,omitempty" json:"isSummary,omitempty"`
	// ValueSet details if this is coded
	Binding *ElementDefinitionBinding `bson:"binding,omitempty" json:"binding,omitempty"`
	// Map element to another set of definitions
	Mapping []ElementDefinitionMapping `bson:"mapping,omitempty" json:"mapping,omitempty"`
}

type ElementDefinitionSlicingDiscriminator struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// value | exists | type | profile | position
	Type custom.Code `bson:"type" json:"type"`
	// Path to element value
	Path string `bson:"path" json:"path"`
}

type ElementDefinitionExample struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Describes the purpose of this example
	Label string `bson:"label" json:"label"`
	// Value of Example (one of allowed types)
	Value custom.Base64binary `bson:"value" json:"value"`
}

type ElementDefinitionConstraint struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Target of 'condition' reference above
	Key custom.ID `bson:"key" json:"key"`
	// Why this constraint is necessary or appropriate
	Requirements *custom.Markdown `bson:"requirements,omitempty" json:"requirements,omitempty"`
	// error | warning
	Severity custom.Code `bson:"severity" json:"severity"`
	// Suppress warning or hint in profile
	Suppress *bool `bson:"suppress,omitempty" json:"suppress,omitempty"`
	// Human description of constraint
	Human string `bson:"human" json:"human"`
	// FHIRPath expression of constraint
	Expression *string `bson:"expression,omitempty" json:"expression,omitempty"`
	// Reference to original source of constraint
	Source *custom.Canonical[StructureDefinition] `bson:"source,omitempty" json:"source,omitempty"`
}

type ElementDefinitionBinding struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// required | extensible | preferred | example
	Strength custom.Code `bson:"strength" json:"strength"`
	// Intended use of codes in the bound value set
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Source of value set
	ValueSet *custom.Canonical[ValueSet] `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	// Additional Bindings - more rules about the binding
	Additional []ElementDefinitionBindingAdditional `bson:"additional,omitempty" json:"additional,omitempty"`
}

type ElementDefinitionMapping struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Reference to mapping declaration
	Identity custom.ID `bson:"identity" json:"identity"`
	// Computable language of mapping
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Details of the mapping
	Map string `bson:"map" json:"map"`
	// Comments about the mapping or its use
	Comment *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
}

type ElementDefinitionTypeProfile interface {
	gofhirclient.Resource

	Is_ElementDefinitionTypeProfile()
}

func (s StructureDefinition) Is_ElementDefinitionTypeProfile() {}
func (i ImplementationGuide) Is_ElementDefinitionTypeProfile() {}

type ElementDefinitionTypeTargetProfile interface {
	gofhirclient.Resource

	Is_ElementDefinitionTypeTargetProfile()
}

func (s StructureDefinition) Is_ElementDefinitionTypeTargetProfile() {}
func (i ImplementationGuide) Is_ElementDefinitionTypeTargetProfile() {}

func (e ElementDefinition) ResourceType() string {
	return "StructureDefinition"
}
