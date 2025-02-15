// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/CodeSystem
type CodeSystem struct {
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
	// Canonical identifier for this code system, represented as a URI (globally unique) (Coding.system)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the code system (business identifier)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the code system (Coding.version)
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this code system (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this code system (human friendly)
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
	// Natural language description of the code system
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for code system (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this code system is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the CodeSystem was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the CodeSystem was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the CodeSystem is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Who authored the CodeSystem
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the CodeSystem
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the CodeSystem
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the CodeSystem
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// If code comparison is case sensitive
	CaseSensitive *bool `bson:"caseSensitive,omitempty" json:"caseSensitive,omitempty"`
	// Canonical reference to the value set with entire code system
	ValueSet *custom.Canonical[ValueSet] `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	// grouped-by | is-a | part-of | classified-with
	HierarchyMeaning *custom.Code `bson:"hierarchyMeaning,omitempty" json:"hierarchyMeaning,omitempty"`
	// If code system defines a compositional grammar
	Compositional *bool `bson:"compositional,omitempty" json:"compositional,omitempty"`
	// If definitions are not stable
	VersionNeeded *bool `bson:"versionNeeded,omitempty" json:"versionNeeded,omitempty"`
	// not-present | example | fragment | complete | supplement
	Content custom.Code `bson:"content" json:"content"`
	// Canonical URL of Code System this adds designations and properties to
	Supplements *custom.Canonical[CodeSystem] `bson:"supplements,omitempty" json:"supplements,omitempty"`
	// Total concepts in the code system
	Count *uint32 `bson:"count,omitempty" json:"count,omitempty"`
	// Filter that can be used in a value set
	Filter []CodeSystemFilter `bson:"filter,omitempty" json:"filter,omitempty"`
	// Additional information supplied about each concept
	Property []CodeSystemProperty `bson:"property,omitempty" json:"property,omitempty"`
	// Concepts in the code system
	Concept []CodeSystemConcept `bson:"concept,omitempty" json:"concept,omitempty"`
}

type CodeSystemFilter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code that identifies the filter
	Code custom.Code `bson:"code" json:"code"`
	// How or why the filter is used
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// = | is-a | descendent-of | is-not-a | regex | in | not-in | generalizes | child-of | descendent-leaf | exists
	Operator []custom.Code `bson:"operator" json:"operator"`
	// What to use for the value
	Value string `bson:"value" json:"value"`
}

type CodeSystemProperty struct {
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
	// Why the property is defined, and/or what it conveys
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// code | Coding | string | integer | boolean | dateTime | decimal
	Type custom.Code `bson:"type" json:"type"`
}

type CodeSystemConcept struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code that identifies concept
	Code custom.Code `bson:"code" json:"code"`
	// Text to display to the user
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// Formal definition
	Definition *string `bson:"definition,omitempty" json:"definition,omitempty"`
	// Additional representations for the concept
	Designation []CodeSystemConceptDesignation `bson:"designation,omitempty" json:"designation,omitempty"`
	// Property value for the concept
	Property []CodeSystemConceptProperty `bson:"property,omitempty" json:"property,omitempty"`
	// Child Concepts (is-a/contains/categorizes)
	Concept []interface{} `bson:"concept,omitempty" json:"concept,omitempty"`
}

type CodeSystemConceptDesignation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Human language of the designation
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Details how this designation would be used
	Use *Coding `bson:"use,omitempty" json:"use,omitempty"`
	// Additional ways how this designation would be used
	AdditionalUse []Coding `bson:"additionalUse,omitempty" json:"additionalUse,omitempty"`
	// The text value for this designation
	Value string `bson:"value" json:"value"`
}

type CodeSystemConceptProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to CodeSystem.property.code
	Code custom.Code `bson:"code" json:"code"`
	// Value of the property for this concept
	Value custom.Code `bson:"value" json:"value"`
}

func (c CodeSystem) ResourceType() string {
	return "StructureDefinition"
}
