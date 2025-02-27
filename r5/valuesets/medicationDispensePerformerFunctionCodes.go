// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/medicationdispense-performer-function
*/type MedicationDispensePerformerFunctionCodes string

const (
	// Recorded the details of the request
	MedicationDispensePerformerFunctionCodesDataenterer MedicationDispensePerformerFunctionCodes = "dataenterer"
	// Prepared the medication.
	MedicationDispensePerformerFunctionCodesPackager MedicationDispensePerformerFunctionCodes = "packager"
	// Performed initial quality assurance on the prepared medication
	MedicationDispensePerformerFunctionCodesChecker MedicationDispensePerformerFunctionCodes = "checker"
	// Performed the final quality assurance on the prepared medication against the request. Typically, this is a pharmacist function.
	MedicationDispensePerformerFunctionCodesFinalchecker MedicationDispensePerformerFunctionCodes = "finalchecker"
	// Provided the drug information to the patient at the time of dispensing.
	MedicationDispensePerformerFunctionCodesCounsellor MedicationDispensePerformerFunctionCodes = "counsellor"
)
