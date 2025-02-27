// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/med-admin-perform-function
*/type MedicationAdministrationPerformerFunctionCodes string

const (
	// A person, non-person living subject, organization or device that who actually and principally carries out the action
	MedicationAdministrationPerformerFunctionCodesPerformer MedicationAdministrationPerformerFunctionCodes = "performer"
	// A person who verifies the correctness and appropriateness of the service (plan, order, event, etc.) and hence takes on accountability.
	MedicationAdministrationPerformerFunctionCodesVerifier MedicationAdministrationPerformerFunctionCodes = "verifier"
	// A person witnessing the action happening without doing anything. A witness is not necessarily aware, much less approves of anything stated in the service event. Example for a witness is students watching an operation or an advanced directive witness.
	MedicationAdministrationPerformerFunctionCodesWitness MedicationAdministrationPerformerFunctionCodes = "witness"
)
