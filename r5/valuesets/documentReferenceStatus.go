// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/document-reference-status
*/type DocumentReferenceStatus string

const (
	// This is the current reference for this document.
	DocumentReferenceStatusCurrent DocumentReferenceStatus = "current"
	// This reference has been superseded by another reference.
	DocumentReferenceStatusSuperseded DocumentReferenceStatus = "superseded"
	// This reference was created in error.
	DocumentReferenceStatusEnteredInError DocumentReferenceStatus = "entered-in-error"
)
