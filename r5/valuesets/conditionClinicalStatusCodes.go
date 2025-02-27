// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/condition-clinical
*/type ConditionClinicalStatusCodes string

const (
	// The subject is currently experiencing the condition or situation, there is evidence of the condition or situation, or considered to be a significant risk.
	ConditionClinicalStatusCodesActive ConditionClinicalStatusCodes = "active"
	// The subject is no longer experiencing the condition or situation and there is no longer evidence or appreciable risk of the condition or situation.
	ConditionClinicalStatusCodesInactive ConditionClinicalStatusCodes = "inactive"
	// The authoring/source system does not know which of the status values currently applies for this condition. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	ConditionClinicalStatusCodesUnknown ConditionClinicalStatusCodes = "unknown"
)
