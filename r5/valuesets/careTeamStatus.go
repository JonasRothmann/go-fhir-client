// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/care-team-status
*/type CareTeamStatus string

const (
	// The care team has been drafted and proposed, but not yet participating in the coordination and delivery of patient care.
	CareTeamStatusProposed CareTeamStatus = "proposed"
	// The care team is currently participating in the coordination and delivery of care.
	CareTeamStatusActive CareTeamStatus = "active"
	// The care team is temporarily on hold or suspended and not participating in the coordination and delivery of care.
	CareTeamStatusSuspended CareTeamStatus = "suspended"
	// The care team was, but is no longer, participating in the coordination and delivery of care.
	CareTeamStatusInactive CareTeamStatus = "inactive"
	// The care team should have never existed.
	CareTeamStatusEnteredInError CareTeamStatus = "entered-in-error"
)
