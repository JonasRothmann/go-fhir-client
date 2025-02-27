// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/response-code
*/type ResponseType string

const (
	// The message was accepted and processed without error.
	ResponseTypeOk ResponseType = "ok"
	// Some internal unexpected error occurred - wait and try again. Note - this is usually used for things like database unavailable, which may be expected to resolve, though human intervention may be required.
	ResponseTypeTransientError ResponseType = "transient-error"
	// The message was rejected because of a problem with the content. There is no point in re-sending without change. The response narrative SHALL describe the issue.
	ResponseTypeFatalError ResponseType = "fatal-error"
)
