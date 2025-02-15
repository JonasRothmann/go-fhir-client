// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ConceptMap
type ConceptMap struct {
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
	// Canonical identifier for this concept map, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the concept map
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the concept map
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this concept map (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this concept map (human friendly)
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
	// Natural language description of the concept map
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for concept map (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this concept map is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the ConceptMap was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the ConceptMap was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the ConceptMap is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Who authored the ConceptMap
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the ConceptMap
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the ConceptMap
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the ConceptMap
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Additional properties of the mapping
	Property []ConceptMapProperty `bson:"property,omitempty" json:"property,omitempty"`
	// Definition of an additional attribute to act as a data source or target
	AdditionalAttribute []ConceptMapAdditionalAttribute `bson:"additionalAttribute,omitempty" json:"additionalAttribute,omitempty"`
	// The source value set that contains the concepts that are being mapped
	SourceScope *custom.URI `bson:"sourceScope,omitempty" json:"sourceScope,omitempty"`
	// The target value set which provides context for the mappings
	TargetScope *custom.URI `bson:"targetScope,omitempty" json:"targetScope,omitempty"`
	// Same source and target systems
	Group []ConceptMapGroup `bson:"group,omitempty" json:"group,omitempty"`
}

type ConceptMapGroupElementTarget struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code that identifies the target element
	Code *custom.Code `bson:"code,omitempty" json:"code,omitempty"`
	// Display for the code
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// Identifies the set of target concepts
	ValueSet *custom.Canonical[ValueSet] `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	// related-to | equivalent | source-is-narrower-than-target | source-is-broader-than-target | not-related-to
	Relationship custom.Code `bson:"relationship" json:"relationship"`
	// Description of status/issues in mapping
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty"`
	// Property value for the source -> target mapping
	Property []ConceptMapGroupElementTargetProperty `bson:"property,omitempty" json:"property,omitempty"`
	// Other properties required for this mapping
	DependsOn []ConceptMapGroupElementTargetDependsOn `bson:"dependsOn,omitempty" json:"dependsOn,omitempty"`
	// Other data elements that this mapping also produces
	Product []interface{} `bson:"product,omitempty" json:"product,omitempty"`
}

type ConceptMapGroupElementTargetProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to ConceptMap.property.code
	Code custom.Code `bson:"code" json:"code"`
	// Value of the property for this concept
	Value Coding `bson:"value" json:"value"`
}

type ConceptMapGroupUnmapped struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// use-source-code | fixed | other-map
	Mode custom.Code `bson:"mode" json:"mode"`
	// Fixed code when mode = fixed
	Code *custom.Code `bson:"code,omitempty" json:"code,omitempty"`
	// Display for the code
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// Fixed code set when mode = fixed
	ValueSet *custom.Canonical[ValueSet] `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	// related-to | equivalent | source-is-narrower-than-target | source-is-broader-than-target | not-related-to
	Relationship *custom.Code `bson:"relationship,omitempty" json:"relationship,omitempty"`
	// canonical reference to an additional ConceptMap to use for mapping if the source concept is unmapped
	OtherMap *custom.Canonical[ConceptMap] `bson:"otherMap,omitempty" json:"otherMap,omitempty"`
}

type ConceptMapProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifies the property on the mappings, and when referred to in the $translate operation
	Code custom.Code `bson:"code" json:"code"`
	// Formal identifier for the property
	Uri *custom.URI `bson:"uri,omitempty" json:"uri,omitempty"`
	// Why the property is defined, and/or what it conveys
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Coding | string | integer | boolean | dateTime | decimal | code
	Type custom.Code `bson:"type" json:"type"`
	// The CodeSystem from which code values come
	System *custom.Canonical[CodeSystem] `bson:"system,omitempty" json:"system,omitempty"`
}

type ConceptMapAdditionalAttribute struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifies this additional attribute through this resource
	Code custom.Code `bson:"code" json:"code"`
	// Formal identifier for the data element referred to in this attribte
	Uri *custom.URI `bson:"uri,omitempty" json:"uri,omitempty"`
	// Why the additional attribute is defined, and/or what the data element it refers to is
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// code | Coding | string | boolean | Quantity
	Type custom.Code `bson:"type" json:"type"`
}

type ConceptMapGroup struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Source system where concepts to be mapped are defined
	Source *custom.Canonical[CodeSystem] `bson:"source,omitempty" json:"source,omitempty"`
	// Target system that the concepts are to be mapped to
	Target *custom.Canonical[CodeSystem] `bson:"target,omitempty" json:"target,omitempty"`
	// Mappings for a concept from the source set
	Element []ConceptMapGroupElement `bson:"element" json:"element"`
	// What to do when there is no mapping target for the source concept and ConceptMap.group.element.noMap is not true
	Unmapped *ConceptMapGroupUnmapped `bson:"unmapped,omitempty" json:"unmapped,omitempty"`
}

type ConceptMapGroupElement struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifies element being mapped
	Code *custom.Code `bson:"code,omitempty" json:"code,omitempty"`
	// Display for the code
	Display *string `bson:"display,omitempty" json:"display,omitempty"`
	// Identifies the set of concepts being mapped
	ValueSet *custom.Canonical[ValueSet] `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	// No mapping to a target concept for this source concept
	NoMap *bool `bson:"noMap,omitempty" json:"noMap,omitempty"`
	// Concept in target system for element
	Target []ConceptMapGroupElementTarget `bson:"target,omitempty" json:"target,omitempty"`
}

type ConceptMapGroupElementTargetDependsOn struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A reference to a mapping attribute defined in ConceptMap.additionalAttribute
	Attribute custom.Code `bson:"attribute" json:"attribute"`
	// Value of the referenced data element
	Value *custom.Code `bson:"value,omitempty" json:"value,omitempty"`
	// The mapping depends on a data element with a value from this value set
	ValueSet *custom.Canonical[ValueSet] `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
}

func (c ConceptMap) ResourceType() string {
	return "StructureDefinition"
}
