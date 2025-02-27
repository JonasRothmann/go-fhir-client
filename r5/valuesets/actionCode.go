// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-code
*/type ActionCode string

const (
	// The action indicates that a particular message should be sent to a participant in the process.
	ActionCodeSendAMessage ActionCode = "send-message"
	// The action indicates that information should be collected from a participant in the process.
	ActionCodeCollectInformation ActionCode = "collect-information"
	// The action indicates that a particular medication should be prescribed to the patient.
	ActionCodePrescribeAMedication ActionCode = "prescribe-medication"
	// The action indicates that a particular immunization should be performed.
	ActionCodeRecommendAnImmunization ActionCode = "recommend-immunization"
	// The action indicates that a particular service should be provided.
	ActionCodeOrderAService ActionCode = "order-service"
	// The action indicates that a particular diagnosis should be proposed.
	ActionCodeProposeADiagnosis ActionCode = "propose-diagnosis"
	// The action indicates that a particular detected issue should be recorded.
	ActionCodeRecordADetectedIssue ActionCode = "record-detected-issue"
	// The action indicates that a particular inference should be recorded.
	ActionCodeRecordAnInference ActionCode = "record-inference"
	// The action indicates that a particular flag should be reported.
	ActionCodeReportAFlag ActionCode = "report-flag"
)
