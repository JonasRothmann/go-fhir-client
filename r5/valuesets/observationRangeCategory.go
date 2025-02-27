// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/observation-range-category
*/type ObservationRangeCategory string

const (
	// Reference (Normal) Range for Ordinal and Continuous Observations.
	ObservationRangeCategoryReference ObservationRangeCategory = "reference"
	// Critical Range for Ordinal and Continuous Observations. Results outside this range are critical.
	ObservationRangeCategoryCritical ObservationRangeCategory = "critical"
	// Absolute Range for Ordinal and Continuous Observations. Results outside this range are not possible.
	ObservationRangeCategoryAbsolute ObservationRangeCategory = "absolute"
)
