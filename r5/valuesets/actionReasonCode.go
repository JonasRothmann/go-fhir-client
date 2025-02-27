// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/action-reason-code
*/type ActionReasonCode string

const (
	// The action should be performed because the patient was determined to be off pathway.
	ActionReasonCodeOffPathway ActionReasonCode = "off-pathway"
	// The action should be performed based on a particular risk assessment.
	ActionReasonCodeRiskAssessment ActionReasonCode = "risk-assessment"
	// The action should be performed to address a detected care gap.
	ActionReasonCodeCareGapDetected ActionReasonCode = "care-gap"
	// The action should be performed to address a detected potential drug-drug interaction.
	ActionReasonCodeDrugDrugInteraction ActionReasonCode = "drug-drug-interaction"
	// The action should be performed to bring the patient's care in line with a quality measure.
	ActionReasonCodeQualityMeasure ActionReasonCode = "quality-measure"
)
