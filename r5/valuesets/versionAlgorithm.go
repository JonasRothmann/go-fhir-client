// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/version-algorithm
*/type VersionAlgorithm string

const (
	// Uses the semantic versioning scheme as defined in [semver.org](http://semver.org).
	VersionAlgorithmSemver VersionAlgorithm = "semver"
	// Versions are integers and ordered numerically
	VersionAlgorithmInteger VersionAlgorithm = "integer"
	// Simple alphabetic sort on a case-insensitive and accent-insensitive basis.  (Sorting of different cases or accented versions of a character is indeterminate)
	VersionAlgorithmAlpha VersionAlgorithm = "alpha"
	// Versions are expressed as an ISO date/time syntax (including syntaxes with only portions of a date)
	VersionAlgorithmDate VersionAlgorithm = "date"
	// Sorted according to the algorithm defined here: [naturalordersort.org](http://www.naturalordersort.org/)
	VersionAlgorithmNatural VersionAlgorithm = "natural"
)
