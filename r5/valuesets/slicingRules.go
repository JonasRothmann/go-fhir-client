// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/resource-slicing-rules
*/type SlicingRules string

const (
	// No additional content is allowed other than that described by the slices in this profile.
	SlicingRulesClosed SlicingRules = "closed"
	// Additional content is allowed anywhere in the list.
	SlicingRulesOpen SlicingRules = "open"
	// Additional content is allowed, but only at the end of the list. Note that using this requires that the slices be ordered, which makes it hard to share uses. This should only be done where absolutely required.
	SlicingRulesOpenAtEnd SlicingRules = "openAtEnd"
)
