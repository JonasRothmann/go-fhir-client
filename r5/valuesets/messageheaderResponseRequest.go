// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/messageheader-response-request
*/type MessageheaderResponseRequest string

const (
	// initiator expects a response for this message.
	MessageheaderResponseRequestAlways MessageheaderResponseRequest = "always"
	// initiator expects a response only if in error.
	MessageheaderResponseRequestErrorRejectConditionsOnly MessageheaderResponseRequest = "on-error"
	// initiator does not expect a response.
	MessageheaderResponseRequestNever MessageheaderResponseRequest = "never"
	// initiator expects a response only if successful.
	MessageheaderResponseRequestSuccessfulCompletionOnly MessageheaderResponseRequest = "on-success"
)
