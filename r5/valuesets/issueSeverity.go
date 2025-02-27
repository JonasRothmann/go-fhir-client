// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/issue-severity
*/type IssueSeverity string

const (
	// The issue caused the action to fail and no further checking could be performed.
	IssueSeverityFatal IssueSeverity = "fatal"
	// The issue is sufficiently important to cause the action to fail.
	IssueSeverityError IssueSeverity = "error"
	// The issue is not important enough to cause the action to fail but may cause it to be performed suboptimally or in a way that is not as desired.
	IssueSeverityWarning IssueSeverity = "warning"
	// The issue has no relation to the degree of success of the action.
	IssueSeverityInformation IssueSeverity = "information"
	// The operation completed successfully.
	IssueSeveritySuccess IssueSeverity = "success"
)
