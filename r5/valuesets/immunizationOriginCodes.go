// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/immunization-origin
  - provider
  - record
  - recall
  - school
*/type ImmunizationOriginCodes string

const (
	// The data for the immunization event originated with another provider.
	ImmunizationOriginCodesProvider ImmunizationOriginCodes = "provider"
	// The data for the immunization event originated with a written record for the patient.
	ImmunizationOriginCodesRecord ImmunizationOriginCodes = "record"
	// The data for the immunization event originated from the recollection of the patient or parent/guardian of the patient.
	ImmunizationOriginCodesRecall ImmunizationOriginCodes = "recall"
	// The data for the immunization event originated with a school record for the patient.
	ImmunizationOriginCodesSchool ImmunizationOriginCodes = "school"
)
