// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/executablevalueset
type ExecutableValueSetComposeInclude struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The system the codes come from
	System *custom.URI `json:"system,omitempty"`
	// Specific version of the code system referred to
	Version *string `json:"version,omitempty"`
	// A concept defined in the system
	Concept []ExecutableValueSetComposeIncludeConcept `json:"concept,omitempty"`
	// Select codes/concepts by their properties (including relationships)
	Filter []ExecutableValueSetComposeIncludeFilter `json:"filter,omitempty"`
	// Select the contents included in this value set
	ValueSet []custom.Canonical[ValueSet] `json:"valueSet,omitempty"`
	// A copyright statement for the specific code system included in the value set
	Copyright *string `json:"copyright,omitempty"`
}

type ExecutableValueSetComposeIncludeConcept struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code or expression from system
	Code custom.Code `json:"code"`
	// Text to display for this code for this value set in this valueset
	Display *string `json:"display,omitempty"`
	// Additional representations for this concept
	Designation []ExecutableValueSetComposeIncludeConceptDesignation `json:"designation,omitempty"`
}

type ExecutableValueSetComposeIncludeConceptDesignation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Human language of the designation
	Language *custom.Code `json:"language,omitempty"`
	// Types of uses of designations
	Use *Coding `json:"use,omitempty"`
	// Additional ways how this designation would be used
	AdditionalUse []Coding `json:"additionalUse,omitempty"`
	// The text value for this designation
	Value string `json:"value"`
}

type ExecutableValueSetExpansion struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies the value set expansion (business identifier)
	Identifier *custom.URI `json:"identifier,omitempty"`
	// Opaque urls for paging through expansion results
	Next *custom.URI `json:"next,omitempty"`
	// Time ValueSet expansion happened
	Timestamp custom.DateTime `json:"timestamp"`
	// Total number of codes in the expansion
	Total *int32 `json:"total,omitempty"`
	// Offset at which this resource starts
	Offset *int32 `json:"offset,omitempty"`
	// Parameter that controlled the expansion process
	Parameter []ExecutableValueSetExpansionParameter `json:"parameter,omitempty"`
	// Additional information supplied about each concept
	Property []ExecutableValueSetExpansionProperty `json:"property,omitempty"`
	// Codes in the value set
	Contains []ExecutableValueSetExpansionContains `json:"contains,omitempty"`
}

type ExecutableValueSetExpansionProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies the property on the concepts, and when referred to in operations
	Code custom.Code `json:"code"`
	// Formal identifier for the property
	Uri *custom.URI `json:"uri,omitempty"`
}

type ExecutableValueSetExpansionContainsProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to ValueSet.expansion.property.code
	Code custom.Code `json:"code"`
	// Value of the property for this concept
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Value of the property for this concept
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Value of the property for this concept
	ValueString *string `json:"valueString,omitempty"`
	// Value of the property for this concept
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Value of the property for this concept
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Value of the property for this concept
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Value of the property for this concept
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// SubProperty value for the concept
	SubProperty []ExecutableValueSetExpansionContainsPropertySubProperty `json:"subProperty,omitempty"`
}

type ExecutableValueSetExpansionContainsPropertySubProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to ValueSet.expansion.property.code
	Code custom.Code `json:"code"`
	// Value of the subproperty for this concept
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Value of the subproperty for this concept
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Value of the subproperty for this concept
	ValueString *string `json:"valueString,omitempty"`
	// Value of the subproperty for this concept
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Value of the subproperty for this concept
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Value of the subproperty for this concept
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Value of the subproperty for this concept
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
}

type ExecutableValueSetCompose struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Fixed date for references with no specified version (transitive)
	LockedDate *custom.Date `json:"lockedDate,omitempty"`
	// Whether inactive codes are in the value set
	Inactive *bool `json:"inactive,omitempty"`
	// Include one or more codes from a code system or other value set(s)
	Include []ExecutableValueSetComposeInclude `json:"include"`
	// Explicitly exclude codes from a code system or other value sets
	Exclude []interface{} `json:"exclude,omitempty"`
	// Property to return if client doesn't override
	Property []string `json:"property,omitempty"`
}

type ExecutableValueSetComposeIncludeFilter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A property/filter defined by the code system
	Property custom.Code `json:"property"`
	// = | is-a | descendent-of | is-not-a | regex | in | not-in | generalizes | child-of | descendent-leaf | exists
	Op custom.Code `json:"op"`
	// Code from the system, or regex criteria, or boolean value for exists
	Value string `json:"value"`
}

type ExecutableValueSetExpansionParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name as assigned by the client or server
	Name string `json:"name"`
	// Value of the named parameter
	ValueString *string `json:"valueString,omitempty"`
	// Value of the named parameter
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Value of the named parameter
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Value of the named parameter
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Value of the named parameter
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Value of the named parameter
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Value of the named parameter
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
}

type ExecutableValueSetExpansionContains struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// System value for the code
	System *custom.URI `json:"system,omitempty"`
	// If user cannot select this entry
	Abstract *bool `json:"abstract,omitempty"`
	// If concept is inactive in the code system
	Inactive *bool `json:"inactive,omitempty"`
	// Version in which this code/display is defined
	Version *string `json:"version,omitempty"`
	// Code - if blank, this is not a selectable code
	Code custom.Code `json:"code"`
	// User display for the concept
	Display *string `json:"display,omitempty"`
	// Additional representations for this item
	Designation []interface{} `json:"designation,omitempty"`
	// Property value for the concept
	Property []ExecutableValueSetExpansionContainsProperty `json:"property,omitempty"`
	// Codes contained under this entry
	Contains []interface{} `json:"contains,omitempty"`
}

type ExecutableValueSetScope struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Criteria describing which concepts or codes should be included and why
	InclusionCriteria *string `json:"inclusionCriteria,omitempty"`
	// Criteria describing which concepts or codes should be excluded and why
	ExclusionCriteria *string `json:"exclusionCriteria,omitempty"`
}

type ExecutableValueSet struct {
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
	// Extension
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Canonical identifier for this value set, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Additional identifier for the value set (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the value set
	Version string `json:"version"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this value set (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this value set (human friendly)
	Title string `json:"title"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental bool `json:"experimental"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the value set
	Description custom.Markdown `json:"description"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for value set (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Indicates whether or not any change to the content logical definition may occur
	Immutable *bool `json:"immutable,omitempty"`
	// Why this value set is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the ValueSet was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the ValueSet was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the ValueSet is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `json:"topic,omitempty"`
	// Who authored the ValueSet
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the ValueSet
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the ValueSet
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the ValueSet
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Content logical definition of the value set (CLD)
	Compose *ExecutableValueSetCompose `json:"compose,omitempty"`
	// Used when the value set is "expanded"
	Expansion ExecutableValueSetExpansion `json:"expansion"`
	// Description of the semantic space the Value Set Expansion is intended to cover and should further clarify the text in ValueSet.description
	Scope *ExecutableValueSetScope `json:"scope,omitempty"`
}

type OtherExecutableValueSet ExecutableValueSet

func (e ExecutableValueSet) ResourceType() string {
	return "ExecutableValueSet"
}

func (e ExecutableValueSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExecutableValueSet
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherExecutableValueSet: OtherExecutableValueSet(e), ResourceType: e.ResourceType()})
}

func UnmarshalExecutableValueSet(b []byte) (res ExecutableValueSet, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
