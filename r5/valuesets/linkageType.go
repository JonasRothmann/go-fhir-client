// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/linkage-type
*/type LinkageType string

const (
	// The resource represents the "source of truth" (from the perspective of this Linkage resource) for the underlying event/condition/etc.
	LinkageTypeSource LinkageType = "source"
	// The resource represents an alternative view of the underlying event/condition/etc.  The resource may still be actively maintained, even though it is not considered to be the source of truth.
	LinkageTypeAlternate LinkageType = "alternate"
	// The resource represents an obsolete record of the underlying event/condition/etc.  It is not expected to be actively maintained.
	LinkageTypeHistorical LinkageType = "historical"
)
