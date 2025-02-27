// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/filter-operator
*/type FilterOperator string

const (
	// The specified property of the code equals the provided value.
	FilterOperatorEquals FilterOperator = "="
	// Includes all concept ids that have a transitive is-a relationship with the concept Id provided as the value, including the provided concept itself (include descendant codes and self).
	FilterOperatorIsABySubsumption FilterOperator = "is-a"
	// Includes all concept ids that have a transitive is-a relationship with the concept Id provided as the value, excluding the provided concept itself (i.e. include descendant codes only).
	FilterOperatorDescendentOfBySubsumption FilterOperator = "descendent-of"
	// The specified property of the code does not have an is-a relationship with the provided value.
	FilterOperatorNotIsABySubsumption FilterOperator = "is-not-a"
	// The specified property of the code  matches the regex specified in the provided value.
	FilterOperatorRegex FilterOperator = "regex"
	// The specified property of the code is in the set of codes or concepts specified in the provided value (comma-separated list).
	FilterOperatorIn FilterOperator = "in"
	// The specified property of the code is not in the set of codes or concepts specified in the provided value (comma-separated list).
	FilterOperatorNotInSet FilterOperator = "not-in"
	// Includes all concept ids that have a transitive is-a relationship from the concept Id provided as the value, including the provided concept itself (i.e. include ancestor codes and self).
	FilterOperatorGeneralizes FilterOperator = "generalizes"
	// Only concepts with a direct hierarchical relationship to the index code and no other concepts. This does not include the index code in the output.
	FilterOperatorChildOf FilterOperator = "child-of"
	// Includes concept ids that have a transitive is-a relationship with the concept Id provided as the value, but which do not have any concept ids with transitive is-a relationships with themselves.
	FilterOperatorDescendentLeaf FilterOperator = "descendent-leaf"
	// The specified property of the code has at least one value (if the specified value is true; if the specified value is false, then matches when the specified property of the code has no values).
	FilterOperatorExists FilterOperator = "exists"
)
