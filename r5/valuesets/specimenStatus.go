// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/specimen-status
*/type SpecimenStatus string

const (
	// The physical specimen is present and in good condition.
	SpecimenStatusAvailable SpecimenStatus = "available"
	// There is no physical specimen because it is either lost, destroyed or consumed.
	SpecimenStatusUnavailable SpecimenStatus = "unavailable"
	// The specimen cannot be used because of a quality issue such as a broken container, contamination, or too old.
	SpecimenStatusUnsatisfactory SpecimenStatus = "unsatisfactory"
	// The specimen was entered in error and therefore nullified.
	SpecimenStatusEnteredInError SpecimenStatus = "entered-in-error"
)
