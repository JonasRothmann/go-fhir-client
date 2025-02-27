// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/medicationrequest-status-reason
*/type MedicationRequestStatusReasonCodes string

const (
	// This therapy has been ordered as a backup to a preferred therapy. This order will be released when and if the preferred therapy is unsuccessful.
	MedicationRequestStatusReasonCodesAltchoice MedicationRequestStatusReasonCodes = "altchoice"
	// Clarification is required before the order can be acted upon.
	MedicationRequestStatusReasonCodesClarif MedicationRequestStatusReasonCodes = "clarif"
	// The current level of the medication in the patient's system is too high. The medication is suspended to allow the level to subside to a safer level.
	MedicationRequestStatusReasonCodesDrughigh MedicationRequestStatusReasonCodes = "drughigh"
	// The patient has been admitted to a care facility and their community medications are suspended until hospital discharge.
	MedicationRequestStatusReasonCodesHospadm MedicationRequestStatusReasonCodes = "hospadm"
	// The therapy would interfere with a planned lab test and the therapy is being withdrawn until the test is completed.
	MedicationRequestStatusReasonCodesLabint MedicationRequestStatusReasonCodes = "labint"
	// Patient not available for a period of time due to a scheduled therapy, leave of absence or other reason.
	MedicationRequestStatusReasonCodesPatientNotAvailable MedicationRequestStatusReasonCodes = "non-avail"
	// The patient is pregnant or breast feeding. The therapy will be resumed when the pregnancy is complete and the patient is no longer breastfeeding.
	MedicationRequestStatusReasonCodesPreg MedicationRequestStatusReasonCodes = "preg"
	// The patient is believed to be allergic to a substance that is part of the therapy and the therapy is being temporarily withdrawn to confirm.
	MedicationRequestStatusReasonCodesSalg MedicationRequestStatusReasonCodes = "salg"
	// The drug interacts with a short-term treatment that is more urgently required. This order will be resumed when the short-term treatment is complete.
	MedicationRequestStatusReasonCodesSddi MedicationRequestStatusReasonCodes = "sddi"
	// The drug interacts with a short-term treatment that is more urgently required. This order will be resumed when the short-term treatment is complete.
	MedicationRequestStatusReasonCodesSdupther MedicationRequestStatusReasonCodes = "sdupther"
	// The drug interacts with a short-term treatment that is more urgently required. This order will be resumed when the short-term treatment is complete.
	MedicationRequestStatusReasonCodesSintol MedicationRequestStatusReasonCodes = "sintol"
	// The drug is contraindicated for patients receiving surgery and the patient is scheduled to be admitted for surgery in the near future. The drug will be resumed when the patient has sufficiently recovered from the surgery.
	MedicationRequestStatusReasonCodesSurg MedicationRequestStatusReasonCodes = "surg"
	// The patient was previously receiving a medication contraindicated with the current medication. The current medication will remain on hold until the prior medication has been cleansed from their system.
	MedicationRequestStatusReasonCodesWashout MedicationRequestStatusReasonCodes = "washout"
)
