// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/issue-type
*/type IssueType string

const (
	// Content invalid against the specification or a profile.
	IssueTypeInvalid IssueType = "invalid"
	// An authentication/authorization/permissions issue of some kind.
	IssueTypeSecurity IssueType = "security"
	// Processing issues. These are expected to be final e.g. there is no point resubmitting the same content unchanged.
	IssueTypeProcessing IssueType = "processing"
	// Transient processing issues. The system receiving the message may be able to resubmit the same content once an underlying issue is resolved.
	IssueTypeTransient IssueType = "transient"
	// A message unrelated to the processing success of the completed operation (examples of the latter include things like reminders of password expiry, system maintenance times, etc.).
	IssueTypeInformational IssueType = "informational"
	// The operation completed successfully.
	IssueTypeSuccess IssueType = "success"
)
