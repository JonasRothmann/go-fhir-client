// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/map-source-list-mode
*/type StructureMapSourceListMode string

const (
	// Only process this rule for the first in the list.
	StructureMapSourceListModeFirst StructureMapSourceListMode = "first"
	// Process this rule for all but the first.
	StructureMapSourceListModeNotFirst StructureMapSourceListMode = "not_first"
	// Only process this rule for the last in the list.
	StructureMapSourceListModeLast StructureMapSourceListMode = "last"
	// Process this rule for all but the last.
	StructureMapSourceListModeNotLast StructureMapSourceListMode = "not_last"
	// Only process this rule is there is only item.
	StructureMapSourceListModeOnlyOne StructureMapSourceListMode = "only_one"
)
