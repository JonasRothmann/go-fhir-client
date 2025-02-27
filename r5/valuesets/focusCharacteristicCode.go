// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/focus-characteristic-code
*/type FocusCharacteristicCode string

const (
	// Used to reference a specific article.
	FocusCharacteristicCodeCitation FocusCharacteristicCode = "citation"
	// Used to denote a focus on clinical outcomes, ie evidence variable in role of outcome (measured variable) as observed is considered a "clinical outcome" (patient-important outcome such as mortality, symptoms, function or quality of life).
	FocusCharacteristicCodeObservedOutcomesAreClinicalOutcomes FocusCharacteristicCode = "clinical-outcomes-observed"
	// The population of interest.
	FocusCharacteristicCodePopulation FocusCharacteristicCode = "population"
	// The exposure of interest, such as an intervention.
	FocusCharacteristicCodeExposure FocusCharacteristicCode = "exposure"
	// The comparator (intervention or control state) of interest.
	FocusCharacteristicCodeComparator FocusCharacteristicCode = "comparator"
	// the outcome of interest.
	FocusCharacteristicCodeOutcome FocusCharacteristicCode = "outcome"
	// Any medication exposures. A subset of exposures or interventions that are medications.
	FocusCharacteristicCodeMedicationExposures FocusCharacteristicCode = "medication-exposures"
	// Type of research study, such as randomized trial or case-control study.
	FocusCharacteristicCodeStudyType FocusCharacteristicCode = "study-type"
)
