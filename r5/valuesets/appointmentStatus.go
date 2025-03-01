// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/appointmentstatus
*/type AppointmentStatus string

const (
	// None of the participant(s) have finalized their acceptance of the appointment request, and the start/end time might not be set yet.
	AppointmentStatusProposed AppointmentStatus = "proposed"
	// Some or all of the participant(s) have not finalized their acceptance of the appointment request.
	AppointmentStatusPending AppointmentStatus = "pending"
	// All participant(s) have been considered and the appointment is confirmed to go ahead at the date/times specified.
	AppointmentStatusBooked AppointmentStatus = "booked"
	// The patient/patients has/have arrived and is/are waiting to be seen.
	AppointmentStatusArrived AppointmentStatus = "arrived"
	// The planning stages of the appointment are now complete, the encounter resource will exist and will track further status changes. Note that an encounter may exist before the appointment status is fulfilled for many reasons.
	AppointmentStatusFulfilled AppointmentStatus = "fulfilled"
	// The appointment has been cancelled.
	AppointmentStatusCancelled AppointmentStatus = "cancelled"
	// Some or all of the participant(s) have not/did not appear for the appointment (usually the patient).
	AppointmentStatusNoshow AppointmentStatus = "noshow"
	// This instance should not have been part of this patient's medical record.
	AppointmentStatusEnteredInError AppointmentStatus = "entered-in-error"
	// When checked in, all pre-encounter administrative work is complete, and the encounter may begin. (where multiple patients are involved, they are all present).
	AppointmentStatusCheckedIn AppointmentStatus = "checked-in"
	/*
	   The appointment has been placed on a waitlist, to be scheduled/confirmed in the future when a slot/service is available.
	   A specific time might or might not be pre-allocated.
	*/
	AppointmentStatusWaitlist AppointmentStatus = "waitlist"
)
