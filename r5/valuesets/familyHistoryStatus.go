// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/history-status
*/type FamilyHistoryStatus string

const (
	// Some health information is known and captured, but not complete - see notes for details.
	FamilyHistoryStatusPartial FamilyHistoryStatus = "partial"
	// All available related health information is captured as of the date (and possibly time) when the family member history was taken.
	FamilyHistoryStatusCompleted FamilyHistoryStatus = "completed"
	// This instance should not have been part of this patient's medical record.
	FamilyHistoryStatusEnteredInError FamilyHistoryStatus = "entered-in-error"
	// Health information for this family member is unavailable/unknown.
	FamilyHistoryStatusHealthUnknown FamilyHistoryStatus = "health-unknown"
)
