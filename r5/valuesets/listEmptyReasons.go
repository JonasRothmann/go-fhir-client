// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/list-empty-reason
*/type ListEmptyReasons string

const (
	// Clinical judgment that there are no known items for this list after reasonable investigation. Note that this a positive statement by a clinical user, and not a default position asserted by a computer system in the lack of other information. Example uses:  * For allergies: the patient or patient's agent/guardian has asserted that he/she is not aware of any allergies (NKA - nil known allergies)  * For medications: the patient or patient's agent/guardian has asserted that the patient is known to be taking no medications  * For diagnoses, problems and procedures: the patient or patient's agent/guardian has asserted that there is no known event to record.
	ListEmptyReasonsNilknown ListEmptyReasons = "nilknown"
	// The investigation to find out whether there are items for this list has not occurred.
	ListEmptyReasonsNotasked ListEmptyReasons = "notasked"
	// The content of the list was not provided due to privacy or confidentiality concerns. Note that it should not be assumed that this means that the particular information in question was withheld due to its contents - it can also be a policy decision.
	ListEmptyReasonsWithheld ListEmptyReasons = "withheld"
	// Information to populate this list cannot be obtained; e.g. unconscious patient.
	ListEmptyReasonsUnavailable ListEmptyReasons = "unavailable"
	// The work to populate this list has not yet begun.
	ListEmptyReasonsNotstarted ListEmptyReasons = "notstarted"
	// This list has now closed or has ceased to be relevant or useful.
	ListEmptyReasonsClosed ListEmptyReasons = "closed"
)
