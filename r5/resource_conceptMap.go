// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ConceptMap
type ConceptMapGroupElementTarget struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code that identifies the target element
	Code *custom.Code `json:"code,omitempty"`
	// Display for the code
	Display *string `json:"display,omitempty"`
	// Identifies the set of target concepts
	ValueSet *custom.Canonical[ValueSet] `json:"valueSet,omitempty"`
	// related-to | equivalent | source-is-narrower-than-target | source-is-broader-than-target | not-related-to
	Relationship custom.Code `json:"relationship"`
	// Description of status/issues in mapping
	Comment *string `json:"comment,omitempty"`
	// Property value for the source -> target mapping
	Property []ConceptMapGroupElementTargetProperty `json:"property,omitempty"`
	// Other properties required for this mapping
	DependsOn []ConceptMapGroupElementTargetDependsOn `json:"dependsOn,omitempty"`
	// Other data elements that this mapping also produces
	Product []interface{} `json:"product,omitempty"`
}

type ConceptMapGroupElementTargetDependsOn struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A reference to a mapping attribute defined in ConceptMap.additionalAttribute
	Attribute custom.Code `json:"attribute"`
	// Value of the referenced data element
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Value of the referenced data element
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Value of the referenced data element
	ValueString *string `json:"valueString,omitempty"`
	// Value of the referenced data element
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Value of the referenced data element
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The mapping depends on a data element with a value from this value set
	ValueSet *custom.Canonical[ValueSet] `json:"valueSet,omitempty"`
}

type ConceptMapGroupUnmapped struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// use-source-code | fixed | other-map
	Mode custom.Code `json:"mode"`
	// Fixed code when mode = fixed
	Code *custom.Code `json:"code,omitempty"`
	// Display for the code
	Display *string `json:"display,omitempty"`
	// Fixed code set when mode = fixed
	ValueSet *custom.Canonical[ValueSet] `json:"valueSet,omitempty"`
	// related-to | equivalent | source-is-narrower-than-target | source-is-broader-than-target | not-related-to
	Relationship *custom.Code `json:"relationship,omitempty"`
	// canonical reference to an additional ConceptMap to use for mapping if the source concept is unmapped
	OtherMap *custom.Canonical[ConceptMap] `json:"otherMap,omitempty"`
}

type ConceptMap struct {
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
	// Canonical identifier for this concept map, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Additional identifier for the concept map
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the concept map
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this concept map (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this concept map (human friendly)
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
	// Natural language description of the concept map
	Description *custom.Markdown `json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for concept map (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this concept map is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the ConceptMap was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the ConceptMap was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the ConceptMap is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `json:"topic,omitempty"`
	// Who authored the ConceptMap
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the ConceptMap
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the ConceptMap
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the ConceptMap
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Additional properties of the mapping
	Property []ConceptMapProperty `json:"property,omitempty"`
	// Definition of an additional attribute to act as a data source or target
	AdditionalAttribute []ConceptMapAdditionalAttribute `json:"additionalAttribute,omitempty"`
	// The source value set that contains the concepts that are being mapped
	SourceScopeUri *custom.URI `json:"sourceScopeUri,omitempty"`
	// The source value set that contains the concepts that are being mapped
	SourceScopeCanonical *custom.Canonical[ValueSet] `json:"sourceScopeCanonical,omitempty"`
	// The target value set which provides context for the mappings
	TargetScopeUri *custom.URI `json:"targetScopeUri,omitempty"`
	// The target value set which provides context for the mappings
	TargetScopeCanonical *custom.Canonical[ValueSet] `json:"targetScopeCanonical,omitempty"`
	// Same source and target systems
	Group []ConceptMapGroup `json:"group,omitempty"`
}

type ConceptMapAdditionalAttribute struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies this additional attribute through this resource
	Code custom.Code `json:"code"`
	// Formal identifier for the data element referred to in this attribte
	Uri *custom.URI `json:"uri,omitempty"`
	// Why the additional attribute is defined, and/or what the data element it refers to is
	Description *string `json:"description,omitempty"`
	// code | Coding | string | boolean | Quantity
	Type custom.Code `json:"type"`
}

type ConceptMapGroupElement struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies element being mapped
	Code *custom.Code `json:"code,omitempty"`
	// Display for the code
	Display *string `json:"display,omitempty"`
	// Identifies the set of concepts being mapped
	ValueSet *custom.Canonical[ValueSet] `json:"valueSet,omitempty"`
	// No mapping to a target concept for this source concept
	NoMap *bool `json:"noMap,omitempty"`
	// Concept in target system for element
	Target []ConceptMapGroupElementTarget `json:"target,omitempty"`
}

type ConceptMapProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifies the property on the mappings, and when referred to in the $translate operation
	Code custom.Code `json:"code"`
	// Formal identifier for the property
	Uri *custom.URI `json:"uri,omitempty"`
	// Why the property is defined, and/or what it conveys
	Description *string `json:"description,omitempty"`
	// Coding | string | integer | boolean | dateTime | decimal | code
	Type custom.Code `json:"type"`
	// The CodeSystem from which code values come
	System *custom.Canonical[CodeSystem] `json:"system,omitempty"`
}

type ConceptMapGroup struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Source system where concepts to be mapped are defined
	Source *custom.Canonical[CodeSystem] `json:"source,omitempty"`
	// Target system that the concepts are to be mapped to
	Target *custom.Canonical[CodeSystem] `json:"target,omitempty"`
	// Mappings for a concept from the source set
	Element []ConceptMapGroupElement `json:"element"`
	// What to do when there is no mapping target for the source concept and ConceptMap.group.element.noMap is not true
	Unmapped *ConceptMapGroupUnmapped `json:"unmapped,omitempty"`
}

type ConceptMapGroupElementTargetProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference to ConceptMap.property.code
	Code custom.Code `json:"code"`
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
	// Value of the property for this concept
	ValueCode *custom.Code `json:"valueCode,omitempty"`
}

type OtherConceptMap ConceptMap

func (c ConceptMap) ResourceType() string {
	return "ConceptMap"
}

func (c ConceptMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConceptMap
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherConceptMap: OtherConceptMap(c), ResourceType: c.ResourceType()})
}

func UnmarshalConceptMap(b []byte) (res ConceptMap, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
