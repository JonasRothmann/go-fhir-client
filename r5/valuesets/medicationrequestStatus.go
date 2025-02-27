// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/medicationrequest-status
*/type MedicationrequestStatus string

const (
	// The request is 'actionable', but not all actions that are implied by it have occurred yet.
	MedicationrequestStatusActive MedicationrequestStatus = "active"
	// Actions implied by the request are to be temporarily halted. The request might or might not be resumed. May also be called 'suspended'.
	MedicationrequestStatusOnHold MedicationrequestStatus = "on-hold"
	// The request is no longer active and the subject should no longer be taking the medication.
	MedicationrequestStatusEnded MedicationrequestStatus = "ended"
	// The request was recorded against the wrong patient or for some reason should not have been recorded (e.g. wrong medication, wrong dose, etc.). Some of the actions that are implied by the medication request may have occurred. For example, the medication may have been dispensed and the patient may have taken some of the medication.
	MedicationrequestStatusEnteredInError MedicationrequestStatus = "entered-in-error"
	// The request is not yet 'actionable', e.g. it is a work in progress, requires sign-off, verification or needs to be run through decision support process.
	MedicationrequestStatusDraft MedicationrequestStatus = "draft"
	// The authoring/source system does not know which of the status values currently applies for this request. Note: This concept is not to be used for 'other' - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	MedicationrequestStatusUnknown MedicationrequestStatus = "unknown"
)
