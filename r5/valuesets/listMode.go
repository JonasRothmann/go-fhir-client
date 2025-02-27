// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/list-mode
*/type ListMode string

const (
	// This list is the master list, maintained in an ongoing fashion with regular updates as the real-world list it is tracking changes.
	ListModeWorking ListMode = "working"
	// This list was prepared as a snapshot. It should not be assumed to be current.
	ListModeSnapshot ListMode = "snapshot"
	// A point-in-time list that shows what changes have been made or recommended.  E.g. a discharge medication list showing what was added and removed during an encounter.
	ListModeChanges ListMode = "changes"
)
