// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/search-entry-mode
*/type SearchEntryMode string

const (
	// This resource matched the search specification.
	SearchEntryModeMatch SearchEntryMode = "match"
	// This resource is returned because it is referred to from another resource in the search set.
	SearchEntryModeInclude SearchEntryMode = "include"
	// An OperationOutcome that provides additional information about the processing of a search.
	SearchEntryModeOutcome SearchEntryMode = "outcome"
)
