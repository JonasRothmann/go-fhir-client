// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/encounter-special-arrangements
*/type SpecialArrangements string

const (
	// The patient requires a wheelchair to be made available for the encounter.
	SpecialArrangementsWheel SpecialArrangements = "wheel"
	// An additional bed made available for a person accompanying the patient, for example a parent accompanying a child.
	SpecialArrangementsAdditionalBedding SpecialArrangements = "add-bed"
	// The patient is not fluent in the local language and requires an interpreter to be available. Refer to the Patient.Language property for the type of interpreter required.
	SpecialArrangementsInt SpecialArrangements = "int"
	// A person who accompanies a patient to provide assistive services necessary for the patient's care during the encounter.
	SpecialArrangementsAtt SpecialArrangements = "att"
	// The patient has a guide dog and the location used for the encounter should be able to support the presence of the service animal.
	SpecialArrangementsDog SpecialArrangements = "dog"
)
