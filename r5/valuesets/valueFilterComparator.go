// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/search-comparator
  - eq
  - gt
  - lt
  - ge
  - le
  - sa
  - eb
*/type ValueFilterComparator string

const (
	// the value for the parameter in the resource is equal to the provided value.
	ValueFilterComparatorEq ValueFilterComparator = "eq"
	// the value for the parameter in the resource is greater than the provided value.
	ValueFilterComparatorGt ValueFilterComparator = "gt"
	// the value for the parameter in the resource is less than the provided value.
	ValueFilterComparatorLt ValueFilterComparator = "lt"
	// the value for the parameter in the resource is greater or equal to the provided value.
	ValueFilterComparatorGe ValueFilterComparator = "ge"
	// the value for the parameter in the resource is less or equal to the provided value.
	ValueFilterComparatorLe ValueFilterComparator = "le"
	// the value for the parameter in the resource starts after the provided value.
	ValueFilterComparatorSa ValueFilterComparator = "sa"
	// the value for the parameter in the resource ends before the provided value.
	ValueFilterComparatorEb ValueFilterComparator = "eb"
)
