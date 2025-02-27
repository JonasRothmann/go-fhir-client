// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/concept-map-relationship
*/type ConceptMapRelationship string

const (
	// The concepts are related to each other, but the exact relationship is not known.
	ConceptMapRelationshipRelatedTo ConceptMapRelationship = "related-to"
	// This is an explicit assertion that the target concept is not related to the source concept.
	ConceptMapRelationshipNotRelatedTo ConceptMapRelationship = "not-related-to"
)
