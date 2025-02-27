// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/search-comparator
*/type SearchComparator string

const (
	// the value for the parameter in the resource is equal to the provided value.
	SearchComparatorEq SearchComparator = "eq"
	// the value for the parameter in the resource is not equal to the provided value.
	SearchComparatorNe SearchComparator = "ne"
	// the value for the parameter in the resource is greater than the provided value.
	SearchComparatorGt SearchComparator = "gt"
	// the value for the parameter in the resource is less than the provided value.
	SearchComparatorLt SearchComparator = "lt"
	// the value for the parameter in the resource is greater or equal to the provided value.
	SearchComparatorGe SearchComparator = "ge"
	// the value for the parameter in the resource is less or equal to the provided value.
	SearchComparatorLe SearchComparator = "le"
	// the value for the parameter in the resource starts after the provided value.
	SearchComparatorSa SearchComparator = "sa"
	// the value for the parameter in the resource ends before the provided value.
	SearchComparatorEb SearchComparator = "eb"
	// the value for the parameter in the resource is approximately the same to the provided value.
	SearchComparatorAp SearchComparator = "ap"
)
