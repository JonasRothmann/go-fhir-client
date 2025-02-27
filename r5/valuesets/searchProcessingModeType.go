// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/search-processingmode
*/type SearchProcessingModeType string

const (
	// The search parameter is derived directly from the selected nodes based on the type definitions.
	SearchProcessingModeTypeNormal SearchProcessingModeType = "normal"
	// The search parameter is derived by a phonetic transform from the selected nodes.
	SearchProcessingModeTypePhonetic SearchProcessingModeType = "phonetic"
	// The interpretation of the xpath statement is unknown (and can't be automated).
	SearchProcessingModeTypeOther SearchProcessingModeType = "other"
)
