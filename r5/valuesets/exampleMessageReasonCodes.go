// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/message-reasons-encounter
*/type ExampleMessageReasonCodes string

const (
	// The patient has been admitted.
	ExampleMessageReasonCodesAdmit ExampleMessageReasonCodes = "admit"
	// The patient has been discharged.
	ExampleMessageReasonCodesDischarge ExampleMessageReasonCodes = "discharge"
	// The patient has temporarily left the institution.
	ExampleMessageReasonCodesAbsent ExampleMessageReasonCodes = "absent"
	// The patient has returned from a temporary absence.
	ExampleMessageReasonCodesReturn ExampleMessageReasonCodes = "return"
	// The patient has been moved to a new location.
	ExampleMessageReasonCodesMoved ExampleMessageReasonCodes = "moved"
	// Encounter details have been updated (e.g. to correct a coding error).
	ExampleMessageReasonCodesEdit ExampleMessageReasonCodes = "edit"
)
