// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/diagnosis-role
*/type DiagnosisRole string

const (
	// The diagnoses documented for administrative purposes as the basis for a hospital or other institutional admission
	DiagnosisRoleAdmissionDiagnosis DiagnosisRole = "AD"
	// The diagnoses documented for administrative purposes at the time of hospital or other institutional discharge
	DiagnosisRoleDischargeDiagnosis   DiagnosisRole = "DD"
	DiagnosisRoleChiefComplaint       DiagnosisRole = "CC"
	DiagnosisRoleComorbidityDiagnosis DiagnosisRole = "CM"
	DiagnosisRolePreOpDiagnosis       DiagnosisRole = "pre-op"
	DiagnosisRolePostOpDiagnosis      DiagnosisRole = "post-op"
	// The diagnosis documented for billing purposes
	DiagnosisRoleBilling DiagnosisRole = "billing"
)
