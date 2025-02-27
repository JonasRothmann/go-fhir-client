// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/issue-severity
*/type AuditEventOutcome string

const (
	// The issue caused the action to fail and no further checking could be performed.
	AuditEventOutcomeFatal AuditEventOutcome = "fatal"
	// The issue is sufficiently important to cause the action to fail.
	AuditEventOutcomeError AuditEventOutcome = "error"
	// The issue is not important enough to cause the action to fail but may cause it to be performed suboptimally or in a way that is not as desired.
	AuditEventOutcomeWarning AuditEventOutcome = "warning"
	// The issue has no relation to the degree of success of the action.
	AuditEventOutcomeInformation AuditEventOutcome = "information"
	// The operation completed successfully.
	AuditEventOutcomeSuccess AuditEventOutcome = "success"
)
