//go:generate go-enum --marshal

package valuesets

/*
https://build.fhir.org/valueset-issue-severity.html

ENUM(Unset = "",fatal,error,warning,information,success)
*/
type IssueSeverity string

/*
https://build.fhir.org/valueset-issue-type.html
*/
type IssueType string

const (

	// Content invalid against the specification or a profile.
	IssueTypeInvalid IssueType = "invalid"

	// A structural issue in the content such as wrong namespace, unable to parse the content completely, invalid syntax, etc.
	IssueTypeStructure IssueType = "structure"

	// A required element is missing.
	IssueTypeRequired IssueType = "required"

	// An element or header value is invalid.
	IssueTypeValue IssueType = "value"

	// A content validation rule failed - e.g. a schematron rule.
	IssueTypeInvariant IssueType = "invariant"

	// An authentication/authorization/permissions issue of some kind.
	IssueTypeSecurity IssueType = "security"

	// The client needs to initiate an authentication process.
	IssueTypeLogin IssueType = "login"

	// The user or system was not able to be authenticated (either there is no process, or the proferred token is unacceptable).
	IssueTypeUnknown IssueType = "unknown"

	// User session expired; a login may be required.
	IssueTypeExpired IssueType = "expired"

	// The user does not have the rights to perform this action.
	IssueTypeForbidden IssueType = "forbidden"

	// Some information was not or might not have been returned due to business rules, consent or privacy rules, or access permission constraints.
	// This information may be accessible through alternate processes.
	IssueTypeSuppressed IssueType = "suppressed"

	// Processing issues. These are expected to be final e.g. there is no point resubmitting the same content unchanged.
	IssueTypeProcessing IssueType = "processing"

	// The interaction, operation, resource or profile is not supported.
	IssueTypeNotSupported IssueType = "not-supported"

	// An attempt was made to create a duplicate record.
	IssueTypeDuplicate IssueType = "duplicate"

	// Multiple matching records were found when the operation required only one match.
	IssueTypeMultipleMatches IssueType = "multiple-matches"

	// The reference provided was not found. In a pure RESTful environment, this would be an HTTP 404 error,
	// but this code may be used where the content is not found further into the application architecture.
	IssueTypeNotFound IssueType = "not-found"

	// The reference pointed to content (usually a resource) that has been deleted.
	IssueTypeDeleted IssueType = "deleted"

	// Provided content is too long (typically, this is a denial of service protection type of error).
	IssueTypeTooLong IssueType = "too-long"

	// The code or system could not be understood, or it was not valid in the context of a particular ValueSet.code.
	IssueTypeCodeInvalid IssueType = "code-invalid"

	// An extension was found that was not acceptable, could not be resolved, or a modifierExtension was not recognized.
	IssueTypeExtension IssueType = "extension"

	// The operation was stopped to protect server resources; e.g. a request for a value set expansion on all of SNOMED CT.
	IssueTypeTooCostly IssueType = "too-costly"

	// The content/operation failed to pass some business rule and so could not proceed.
	IssueTypeBusinessRule IssueType = "business-rule"

	// Content could not be accepted because of an edit conflict (i.e. version aware updates).
	// (In a pure RESTful environment, this would be an HTTP 409 error, but this code may be used where the conflict is discovered further into the application architecture.)
	IssueTypeConflict IssueType = "conflict"

	// Some search filters might not have applied on all results.
	// Data may have been included that does not meet all of the filters listed in the self Bundle.link.
	IssueTypeLimitedFilter IssueType = "limited-filter"

	// Transient processing issues. The system receiving the message may be able to resubmit the same content once an underlying issue is resolved.
	IssueTypeTransient IssueType = "transient"

	// A resource/record locking failure (usually in an underlying database).
	IssueTypeLockError IssueType = "lock-error"

	// The persistent store is unavailable; e.g. the database is down for maintenance or similar action,
	// and the interaction or operation cannot be processed.
	IssueTypeNoStore IssueType = "no-store"

	// An unexpected internal error has occurred.
	IssueTypeException IssueType = "exception"

	// An internal timeout has occurred.
	IssueTypeTimeout IssueType = "timeout"

	// Not all data sources typically accessed could be reached or responded in time,
	// so the returned information might not be complete (applies to search interactions and some operations).
	IssueTypeIncomplete IssueType = "incomplete"

	// The system is not prepared to handle this request due to load management.
	IssueTypeThrottled IssueType = "throttled"

	// A message unrelated to the processing success of the completed operation (examples include reminders of password expiry,
	// system maintenance times, etc.).
	IssueTypeInformational IssueType = "informational"

	// The operation completed successfully.
	IssueTypeSuccess IssueType = "success"
)

var issueTypeToMessage = map[IssueType]string{
	IssueTypeInvalid:         "Invalid Content",
	IssueTypeStructure:       "Structural Issue",
	IssueTypeRequired:        "Required element missing",
	IssueTypeValue:           "Element value invalid",
	IssueTypeInvariant:       "Validation rule failed",
	IssueTypeSecurity:        "Security Problem",
	IssueTypeLogin:           "Login Required",
	IssueTypeUnknown:         "Unknown User",
	IssueTypeExpired:         "Session Expired",
	IssueTypeForbidden:       "Forbidden",
	IssueTypeSuppressed:      "Information Suppressed",
	IssueTypeProcessing:      "Processing Failure",
	IssueTypeNotSupported:    "Content not supported",
	IssueTypeDuplicate:       "Duplicate",
	IssueTypeMultipleMatches: "Multiple Matches",
	IssueTypeNotFound:        "Not Found",
	IssueTypeDeleted:         "Deleted",
	IssueTypeTooLong:         "Content Too Long",
	IssueTypeCodeInvalid:     "Invalid Code",
	IssueTypeExtension:       "Unacceptable Extension",
	IssueTypeTooCostly:       "Operation Too Costly",
	IssueTypeBusinessRule:    "Business Rule Violation",
	IssueTypeConflict:        "Edit Version Conflict",
	IssueTypeLimitedFilter:   "Limited Filter Application",
	IssueTypeTransient:       "Transient Issue",
	IssueTypeLockError:       "Lock Error",
	IssueTypeNoStore:         "No Store Available",
	IssueTypeException:       "Exception",
	IssueTypeTimeout:         "Timeout",
	IssueTypeIncomplete:      "Incomplete Results",
	IssueTypeThrottled:       "Throttled",
	IssueTypeInformational:   "Informational Note",
	IssueTypeSuccess:         "Operation Successful",
}

func (i IssueType) Message() string {
	return issueTypeToMessage[i]
}
