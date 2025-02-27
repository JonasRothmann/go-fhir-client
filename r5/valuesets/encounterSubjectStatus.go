// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/encounter-subject-status
*/type EncounterSubjectStatus string

const (
	// The subject has arrived at the physical or virtual location and is ready to receive care.
	EncounterSubjectStatusArrived EncounterSubjectStatus = "arrived"
	// The subject has been seen triaged by staff and is waiting for further care.
	EncounterSubjectStatusTriaged EncounterSubjectStatus = "triaged"
	// The subject is present and commenced receiving care.  This can include periods of waiting between care.
	EncounterSubjectStatusReceivingCare EncounterSubjectStatus = "receiving-care"
	// The subject has left the physical or virtual location, but is expected to return and resume care as part of this encounter.
	EncounterSubjectStatusOnLeave EncounterSubjectStatus = "on-leave"
	// The subject has left the physical or virtual location, and is not expected to return as part of this encounter.
	EncounterSubjectStatusDeparted EncounterSubjectStatus = "departed"
)
