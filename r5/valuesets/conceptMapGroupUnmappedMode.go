// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/conceptmap-unmapped-mode
*/type ConceptMapGroupUnmappedMode string

const (
	// Use the code as provided in the $translate request in one of the following input parameters: sourceCode, sourceCoding, sourceCodeableConcept.
	ConceptMapGroupUnmappedModeUseProvidedSourceCode ConceptMapGroupUnmappedMode = "use-source-code"
	// Use the code(s) explicitly provided in the group.unmapped 'code' or 'valueSet' element.
	ConceptMapGroupUnmappedModeFixed ConceptMapGroupUnmappedMode = "fixed"
	// Use the map identified by the canonical URL in the url element.
	ConceptMapGroupUnmappedModeOtherMap ConceptMapGroupUnmappedMode = "other-map"
)
