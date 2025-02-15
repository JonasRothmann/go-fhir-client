// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ValueSet
type ValueSetComposeInclude struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The system the codes come from
	System *custom.URI `bson:"system,omitempty" json:"system,omitempty"`
	// Specific version of the code system referred to
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// A concept defined in the system
	Concept []ValueSetComposeIncludeConcept `bson:"concept,omitempty" json:"concept,omitempty"`
	// Select codes/concepts by their properties (including relationships)
	Filter []ValueSetComposeIncludeFilter `bson:"filter,omitempty" json:"filter,omitempty"`
	// Select the contents included in this value set
	ValueSet []custom.Canonical[ValueSet] `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	// A copyright statement for the specific code system included in the value set
	Copyright *string `bson:"copyright,omitempty" json:"copyright,omitempty"`
}

type ValueSetComposeIncludeConcept struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code or expression from system
	Code custom.Code `bson:"code" json:"code"`
	// Text to display for this code for this value set in this valueset
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// Additional representations for this concept
	Designation []ValueSetComposeIncludeConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
}

type ValueSetComposeIncludeFilter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A property/filter defined by the code system
	Property custom.Code `bson:"property" json:"property"`
	// = | is-a | descendent-of | is-not-a | regex | in | not-in | generalizes | child-of | descendent-leaf | exists
	Op custom.Code `bson:"op" json:"op"`
	// Code from the system, or regex criteria, or boolean value for exists
	Value string `bson:"value" json:"value"`
}

type ValueSetExpansion struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifies the value set expansion (business identifier)
	Identifier *custom.URI `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Opaque urls for paging through expansion results
	Next *custom.URI `bson:"next,omitempty" json:"next,omitempty"`
	// Time ValueSet expansion happened
	Timestamp custom.DateTime `bson:"timestamp" json:"timestamp"`
	// Total number of codes in the expansion
	Total *int32 `bson:"total,omitempty" json:"total,omitempty"`
	// Offset at which this resource starts
	Offset *int32 `bson:"offset,omitempty" json:"offset,omitempty"`
	// Parameter that controlled the expansion process
	Parameter []ValueSetExpansionParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	// Additional information supplied about each concept
	Property []ValueSetExpansionProperty `bson:"property,omitempty" json:"property,omitempty"`
	// Codes in the value set
	Contains []ValueSetExpansionContains `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetExpansionParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name as assigned by the client or server
	Name string `bson:"name" json:"name"`
	// Value of the named parameter
	Value *string `bson:"value,omitempty" json:"value,omitempty"`
}

type ValueSetExpansionProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifies the property on the concepts, and when referred to in operations
	Code custom.Code `bson:"code" json:"code"`
	// Formal identifier for the property
	Uri *custom.URI `bson:"uri,omitempty" json:"uri,omitempty"`
}

type ValueSetExpansionContains struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// System value for the code
	System *custom.URI `bson:"system,omitempty" json:"system,omitempty"`
	// If user cannot select this entry
	Abstract *bool `bson:"abstract,omitempty" json:"abstract,omitempty"`
	// If concept is inactive in the code system
	Inactive *bool `bson:"inactive,omitempty" json:"inactive,omitempty"`
	// Version in which this code/display is defined
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// Code - if blank, this is not a selectable code
	Code *custom.Code `bson:"code,omitempty" json:"code,omitempty"`
	// User display for the concept
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// Additional representations for this item
	Designation []interface{} `bson:"designation,omitempty" json:"designation,omitempty"`
	// Property value for the concept
	Property []ValueSetExpansionContainsProperty `bson:"property,omitempty" json:"property,omitempty"`
	// Codes contained under this entry
	Contains []interface{} `bson:"contains,omitempty" json:"contains,omitempty"`
}

type ValueSetCompose struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Fixed date for references with no specified version (transitive)
	LockedDate *custom.Date `bson:"lockedDate,omitempty" json:"lockedDate,omitempty"`
	// Whether inactive codes are in the value set
	Inactive *bool `bson:"inactive,omitempty" json:"inactive,omitempty"`
	// Include one or more codes from a code system or other value set(s)
	Include []ValueSetComposeInclude `bson:"include" json:"include"`
	// Explicitly exclude codes from a code system or other value sets
	Exclude []interface{} `bson:"exclude,omitempty" json:"exclude,omitempty"`
	// Property to return if client doesn't override
	Property []string `bson:"property,omitempty" json:"property,omitempty"`
}

type ValueSetScope struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Criteria describing which concepts or codes should be included and why
	InclusionCriteria *string `bson:"inclusionCriteria,omitempty" json:"inclusionCriteria,omitempty"`
	// Criteria describing which concepts or codes should be excluded and why
	ExclusionCriteria *string `bson:"exclusionCriteria,omitempty" json:"exclusionCriteria,omitempty"`
}

type ValueSetComposeIncludeConceptDesignation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Human language of the designation
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Types of uses of designations
	Use *Coding `bson:"use,omitempty" json:"use,omitempty"`
	// Additional ways how this designation would be used
	AdditionalUse []Coding `bson:"additionalUse,omitempty" json:"additionalUse,omitempty"`
	// The text value for this designation
	Value string `bson:"value" json:"value"`
}

type ValueSetExpansionContainsProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to ValueSet.expansion.property.code
	Code custom.Code `bson:"code" json:"code"`
	// Value of the property for this concept
	Value custom.Code `bson:"value" json:"value"`
	// SubProperty value for the concept
	SubProperty []ValueSetExpansionContainsPropertySubProperty `bson:"subProperty,omitempty" json:"subProperty,omitempty"`
}

type ValueSetExpansionContainsPropertySubProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to ValueSet.expansion.property.code
	Code custom.Code `bson:"code" json:"code"`
	// Value of the subproperty for this concept
	Value custom.Code `bson:"value" json:"value"`
}

type ValueSet struct {
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
	// Canonical identifier for this value set, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the value set (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the value set
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this value set (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this value set (human friendly)
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
	// Natural language description of the value set
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for value set (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Indicates whether or not any change to the content logical definition may occur
	Immutable *bool `bson:"immutable,omitempty" json:"immutable,omitempty"`
	// Why this value set is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the ValueSet was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the ValueSet was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the ValueSet is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Who authored the ValueSet
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the ValueSet
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the ValueSet
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the ValueSet
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Content logical definition of the value set (CLD)
	Compose *ValueSetCompose `bson:"compose,omitempty" json:"compose,omitempty"`
	// Used when the value set is "expanded"
	Expansion *ValueSetExpansion `bson:"expansion,omitempty" json:"expansion,omitempty"`
	// Description of the semantic space the Value Set Expansion is intended to cover and should further clarify the text in ValueSet.description
	Scope *ValueSetScope `bson:"scope,omitempty" json:"scope,omitempty"`
}

func (v ValueSet) ResourceType() string {
	return "StructureDefinition"
}
