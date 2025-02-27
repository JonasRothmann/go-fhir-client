// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-participant-function
*/type ActionParticipantFunction string

const (
	// The participant is the performer of the action.
	ActionParticipantFunctionPerformer ActionParticipantFunction = "performer"
	// The participant is the author of the result of the action.
	ActionParticipantFunctionAuthor ActionParticipantFunction = "author"
	// The participant is reviewing the result of the action.
	ActionParticipantFunctionReviewer ActionParticipantFunction = "reviewer"
	// The participant is a witness to the action being performed.
	ActionParticipantFunctionWitness ActionParticipantFunction = "witness"
)
