// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/resource-aggregation-mode
*/type AggregationMode string

const (
	// The reference is a local reference to a contained resource.
	AggregationModeContained AggregationMode = "contained"
	// The reference to a resource that has to be resolved externally to the resource that includes the reference.
	AggregationModeReferenced AggregationMode = "referenced"
)
