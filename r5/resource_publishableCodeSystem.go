// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/publishablecodesystem
type PublishableCodeSystem struct {
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
	// Canonical identifier for this code system, represented as a URI (globally unique) (Coding.system)
	Url custom.URI `json:"url"`
	// Additional identifier for the code system (business identifier)
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the code system (Coding.version)
	Version string `json:"version"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this code system (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this code system (human friendly)
	Title string `json:"title"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental bool `json:"experimental"`
	// Date last changed
	Date custom.DateTime `json:"date"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the code system
	Description custom.Markdown `json:"description"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for code system (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this code system is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the CodeSystem was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the CodeSystem was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the CodeSystem is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `json:"topic,omitempty"`
	// Who authored the CodeSystem
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the CodeSystem
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the CodeSystem
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the CodeSystem
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// If code comparison is case sensitive
	CaseSensitive bool `json:"caseSensitive"`
	// Canonical reference to the value set with entire code system
	ValueSet *custom.Canonical[ValueSet] `json:"valueSet,omitempty"`
	// grouped-by | is-a | part-of | classified-with
	HierarchyMeaning *custom.Code `json:"hierarchyMeaning,omitempty"`
	// If code system defines a compositional grammar
	Compositional *bool `json:"compositional,omitempty"`
	// If definitions are not stable
	VersionNeeded *bool `json:"versionNeeded,omitempty"`
	// not-present | example | fragment | complete | supplement
	Content custom.Code `json:"content"`
	// Canonical URL of Code System this adds designations and properties to
	Supplements *custom.Canonical[CodeSystem] `json:"supplements,omitempty"`
	// Total concepts in the code system
	Count *uint32 `json:"count,omitempty"`
	// Filter that can be used in a value set
	Filter []PublishableCodeSystemFilter `json:"filter,omitempty"`
	// Additional information supplied about each concept
	Property []PublishableCodeSystemProperty `json:"property,omitempty"`
	// Concepts in the code system
	Concept []PublishableCodeSystemConcept `json:"concept,omitempty"`
}

type PublishableCodeSystemFilter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code that identifies the filter
	Code custom.Code `json:"code"`
	// How or why the filter is used
	Description *string `json:"description,omitempty"`
	// = | is-a | descendent-of | is-not-a | regex | in | not-in | generalizes | child-of | descendent-leaf | exists
	Operator []custom.Code `json:"operator"`
	// What to use for the value
	Value string `json:"value"`
}

type PublishableCodeSystemProperty struct {
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
	// Why the property is defined, and/or what it conveys
	Description *string `json:"description,omitempty"`
	// code | Coding | string | integer | boolean | dateTime | decimal
	Type custom.Code `json:"type"`
}

type PublishableCodeSystemConcept struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code that identifies concept
	Code custom.Code `json:"code"`
	// Text to display to the user
	Display *string `json:"display,omitempty"`
	// Formal definition
	Definition *string `json:"definition,omitempty"`
	// Additional representations for the concept
	Designation []PublishableCodeSystemConceptDesignation `json:"designation,omitempty"`
	// Property value for the concept
	Property []PublishableCodeSystemConceptProperty `json:"property,omitempty"`
	// Child Concepts (is-a/contains/categorizes)
	Concept []interface{} `json:"concept,omitempty"`
}

type PublishableCodeSystemConceptDesignation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Human language of the designation
	Language *custom.Code `json:"language,omitempty"`
	// Details how this designation would be used
	Use *Coding `json:"use,omitempty"`
	// Additional ways how this designation would be used
	AdditionalUse []Coding `json:"additionalUse,omitempty"`
	// The text value for this designation
	Value string `json:"value"`
}

type PublishableCodeSystemConceptProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to CodeSystem.property.code
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
}

type OtherPublishableCodeSystem PublishableCodeSystem

func (p PublishableCodeSystem) ResourceType() string {
	return "PublishableCodeSystem"
}

func (p PublishableCodeSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPublishableCodeSystem
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPublishableCodeSystem: OtherPublishableCodeSystem(p), ResourceType: p.ResourceType()})
}

func UnmarshalPublishableCodeSystem(b []byte) (res PublishableCodeSystem, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
