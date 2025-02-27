// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/slotstatus
*/type SlotStatus string

const (
	// Indicates that the time interval is busy because one  or more events have been scheduled for that interval.
	SlotStatusBusy SlotStatus = "busy"
	// Indicates that the time interval is free for scheduling.
	SlotStatusFree SlotStatus = "free"
	// Indicates that the time interval is busy and that the interval cannot be scheduled.
	SlotStatusBusyUnavailable SlotStatus = "busy-unavailable"
	// Indicates that the time interval is busy because one or more events have been tentatively scheduled for that interval.
	SlotStatusBusyTentative SlotStatus = "busy-tentative"
	// This instance should not have been part of this patient's medical record.
	SlotStatusEnteredInError SlotStatus = "entered-in-error"
)
