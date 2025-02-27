// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-relationship-type
*/type ActionRelationshipType string

const (
	// The action must be performed before the related action.
	ActionRelationshipTypeBefore ActionRelationshipType = "before"
	// The action must be performed concurrent with the related action.
	ActionRelationshipTypeConcurrent ActionRelationshipType = "concurrent"
	// The action must be performed after the related action.
	ActionRelationshipTypeAfter ActionRelationshipType = "after"
)
