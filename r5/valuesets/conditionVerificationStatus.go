// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/condition-ver-status
*/type ConditionVerificationStatus string

const (
	// There is not sufficient evidence to assert the presence of the subject's condition.
	ConditionVerificationStatusUnconfirmed ConditionVerificationStatus = "unconfirmed"
	// There is sufficient evidence to assert the presence of the subject's condition.
	ConditionVerificationStatusConfirmed ConditionVerificationStatus = "confirmed"
	// This condition has been ruled out by subsequent diagnostic and clinical evidence.
	ConditionVerificationStatusRefuted ConditionVerificationStatus = "refuted"
	// The statement was entered in error and is not valid.
	ConditionVerificationStatusEnteredInError ConditionVerificationStatus = "entered-in-error"
)
