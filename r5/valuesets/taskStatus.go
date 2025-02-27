// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/task-status
*/type TaskStatus string

const (
	// The task is not yet ready to be acted upon.
	TaskStatusDraft TaskStatus = "draft"
	// The task is ready to be acted upon and action is sought.
	TaskStatusRequested TaskStatus = "requested"
	// A potential performer has claimed ownership of the task and is evaluating whether to perform it.
	TaskStatusReceived TaskStatus = "received"
	// The potential performer has agreed to execute the task but has not yet started work.
	TaskStatusAccepted TaskStatus = "accepted"
	// The potential performer who claimed ownership of the task has decided not to execute it prior to performing any action.
	TaskStatusRejected TaskStatus = "rejected"
	// The task is ready to be performed, but no action has yet been taken.  Used in place of requested/received/accepted/rejected when request assignment and acceptance is a given.
	TaskStatusReady TaskStatus = "ready"
	// The task was not completed.
	TaskStatusCancelled TaskStatus = "cancelled"
	// The task has been started but is not yet complete.
	TaskStatusInProgress TaskStatus = "in-progress"
	// The task has been started but work has been paused.
	TaskStatusOnHold TaskStatus = "on-hold"
	// The task was attempted but could not be completed due to some error.
	TaskStatusFailed TaskStatus = "failed"
	// The task has been completed.
	TaskStatusCompleted TaskStatus = "completed"
	// The task should never have existed and is retained only because of the possibility it may have used.
	TaskStatusEnteredInError TaskStatus = "entered-in-error"
)
