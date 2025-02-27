// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/evidence-variable-event
*/type EvidenceVariableEvent string

const (
	// The time of enrollment for the study participant
	EvidenceVariableEventStudyStart EvidenceVariableEvent = "study-start"
	// The time of initiation of the treatment
	EvidenceVariableEventStartOfTreatment EvidenceVariableEvent = "treatment-start"
	// The time of first detection of the condition
	EvidenceVariableEventDetectionOfCondition EvidenceVariableEvent = "condition-detection"
	// The time of first treatment of the condition
	EvidenceVariableEventTreatmentOfCondition EvidenceVariableEvent = "condition-treatment"
	// The time of admission to the hospital
	EvidenceVariableEventHospitalAdmission EvidenceVariableEvent = "hospital-admission"
	// The time of discharge from the hospital
	EvidenceVariableEventHospitalDischarge EvidenceVariableEvent = "hospital-discharge"
	// The time of surgery
	EvidenceVariableEventOperativeProcedure EvidenceVariableEvent = "operative-procedure"
)
