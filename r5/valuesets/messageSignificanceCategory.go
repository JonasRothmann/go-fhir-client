// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/message-significance-category
*/type MessageSignificanceCategory string

const (
	// The message represents/requests a change that should not be processed more than once; e.g., making a booking for an appointment.
	MessageSignificanceCategoryConsequence MessageSignificanceCategory = "consequence"
	// The message represents a response to query for current information. Retrospective processing is wrong and/or wasteful.
	MessageSignificanceCategoryCurrency MessageSignificanceCategory = "currency"
	// The content is not necessarily intended to be current, and it can be reprocessed, though there may be version issues created by processing old notifications.
	MessageSignificanceCategoryNotification MessageSignificanceCategory = "notification"
)
