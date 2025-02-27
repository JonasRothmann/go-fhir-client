// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/goal-status
*/type GoalLifecycleStatus string

const (
	// A goal is proposed for this patient.
	GoalLifecycleStatusProposed GoalLifecycleStatus = "proposed"
	// A goal is planned for this patient.
	GoalLifecycleStatusPlanned GoalLifecycleStatus = "planned"
	// A proposed goal was accepted or acknowledged.
	GoalLifecycleStatusAccepted GoalLifecycleStatus = "accepted"
	// The goal has been abandoned.
	GoalLifecycleStatusCancelled GoalLifecycleStatus = "cancelled"
	// The goal was entered in error and voided.
	GoalLifecycleStatusEnteredInError GoalLifecycleStatus = "entered-in-error"
	// A proposed goal was rejected.
	GoalLifecycleStatusRejected GoalLifecycleStatus = "rejected"
)
