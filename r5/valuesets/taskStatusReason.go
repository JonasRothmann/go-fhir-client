// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/task-status-reason
*/type TaskStatusReason string

const (
	// An item nessary for task completion is missing.
	TaskStatusReasonMissing TaskStatusReason = "missing"
	// Something related to the task was misidentified.
	TaskStatusReasonMisidentified TaskStatusReason = "misidentified"
	// A piece of equipment necessary for completion of the task is malfunctioning.
	TaskStatusReasonEquipmentIssue TaskStatusReason = "equipment-issue"
	// Something in the environment is preventing task completion.
	TaskStatusReasonEnvironmentalIssue TaskStatusReason = "environmental-issue"
	// Key personnel necessary for task completion are not present.
	TaskStatusReasonPersonnelIssue TaskStatusReason = "personnel-issue"
)
