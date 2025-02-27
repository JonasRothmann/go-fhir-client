// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/administration-subpotent-reason
*/type MedicationAdministrationSubpotentReason string

const (
	// The full amount of the dose was not administered to the patient.
	MedicationAdministrationSubpotentReasonPartialdose MedicationAdministrationSubpotentReason = "partialdose"
	// The medication experienced a cold chain break.
	MedicationAdministrationSubpotentReasonColdchainbreak MedicationAdministrationSubpotentReason = "coldchainbreak"
	// The medication was recalled by the manufacturer.
	MedicationAdministrationSubpotentReasonRecall MedicationAdministrationSubpotentReason = "recall"
	// The medication experienced adverse storage conditions.
	MedicationAdministrationSubpotentReasonAdversestorage MedicationAdministrationSubpotentReason = "adversestorage"
	// The medication was expired at the time of administration.
	MedicationAdministrationSubpotentReasonExpired MedicationAdministrationSubpotentReason = "expired"
)
