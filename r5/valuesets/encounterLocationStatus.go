// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/encounter-location-status
*/type EncounterLocationStatus string

const (
	// The patient is planned to be moved to this location at some point in the future.
	EncounterLocationStatusPlanned EncounterLocationStatus = "planned"
	// The patient is currently at this location, or was between the period specified.A system may update these records when the patient leaves the location to either reserved, or completed.
	EncounterLocationStatusActive EncounterLocationStatus = "active"
	// This location is held empty for this patient.
	EncounterLocationStatusReserved EncounterLocationStatus = "reserved"
	// The patient was at this location during the period specified.Not to be used when the patient is currently at the location.
	EncounterLocationStatusCompleted EncounterLocationStatus = "completed"
)
