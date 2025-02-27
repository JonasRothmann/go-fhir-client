// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/constraint-severity
*/type ConstraintSeverity string

const (
	// If the constraint is violated, the resource is not conformant.
	ConstraintSeverityError ConstraintSeverity = "error"
	// If the constraint is violated, the resource is conformant, but it is not necessarily following best practice.
	ConstraintSeverityWarning ConstraintSeverity = "warning"
)
