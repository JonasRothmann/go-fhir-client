// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/medicationdispense-status-reason
*/type MedicationDispenseStatusReasonCodes string

const (
	// The order has been stopped by the prescriber but this fact has not necessarily captured electronically. Example: A verbal stop, a fax, etc.
	MedicationDispenseStatusReasonCodesFrr01 MedicationDispenseStatusReasonCodes = "frr01"
	// Order has not been fulfilled within a reasonable amount of time, and might not be current.
	MedicationDispenseStatusReasonCodesFrr02 MedicationDispenseStatusReasonCodes = "frr02"
	// Data needed to safely act on the order which was expected to become available independent of the order is not yet available. Example: Lab results, diagnostic imaging, etc.
	MedicationDispenseStatusReasonCodesFrr03 MedicationDispenseStatusReasonCodes = "frr03"
	// Product not available or manufactured. Cannot supply.
	MedicationDispenseStatusReasonCodesFrr04 MedicationDispenseStatusReasonCodes = "frr04"
	// The dispenser has ethical, religious or moral objections to fulfilling the order/dispensing the product.
	MedicationDispenseStatusReasonCodesFrr05 MedicationDispenseStatusReasonCodes = "frr05"
	// Fulfiller not able to provide appropriate care associated with fulfilling the order. Example: Therapy requires ongoing monitoring by fulfiller and fulfiller will be ending practice, leaving town, unable to schedule necessary time, etc.
	MedicationDispenseStatusReasonCodesFrr06 MedicationDispenseStatusReasonCodes = "frr06"
	// This therapy has been ordered as a backup to a preferred therapy. This order will be released when and if the preferred therapy is unsuccessful.
	MedicationDispenseStatusReasonCodesAltchoice MedicationDispenseStatusReasonCodes = "altchoice"
	// Clarification is required before the order can be acted upon.
	MedicationDispenseStatusReasonCodesClarif MedicationDispenseStatusReasonCodes = "clarif"
	// The current level of the medication in the patient's system is too high. The medication is suspended to allow the level to subside to a safer level.
	MedicationDispenseStatusReasonCodesDrughigh MedicationDispenseStatusReasonCodes = "drughigh"
	// The patient has been admitted to a care facility and their community medications are suspended until hospital discharge.
	MedicationDispenseStatusReasonCodesHospadm MedicationDispenseStatusReasonCodes = "hospadm"
	// The therapy would interfere with a planned lab test and the therapy is being withdrawn until the test is completed.
	MedicationDispenseStatusReasonCodesLabint MedicationDispenseStatusReasonCodes = "labint"
	// Patient not available for a period of time due to a scheduled therapy, leave of absence or other reason.
	MedicationDispenseStatusReasonCodesPatientNotAvailable MedicationDispenseStatusReasonCodes = "non-avail"
	// The patient is pregnant or breast feeding. The therapy will be resumed when the pregnancy is complete and the patient is no longer breastfeeding.
	MedicationDispenseStatusReasonCodesPreg MedicationDispenseStatusReasonCodes = "preg"
	// The patient is believed to be allergic to a substance that is part of the therapy and the therapy is being temporarily withdrawn to confirm.
	MedicationDispenseStatusReasonCodesSaig MedicationDispenseStatusReasonCodes = "saig"
	// The drug interacts with a short-term treatment that is more urgently required. This order will be resumed when the short-term treatment is complete.
	MedicationDispenseStatusReasonCodesSddi MedicationDispenseStatusReasonCodes = "sddi"
	// Another short-term co-occurring therapy fulfills the same purpose as this therapy. This therapy will be resumed when the co-occuring therapy is complete.
	MedicationDispenseStatusReasonCodesSdupther MedicationDispenseStatusReasonCodes = "sdupther"
	// The patient is believed to have an intolerance to a substance that is part of the therapy and the therapy is being temporarily withdrawn to confirm.
	MedicationDispenseStatusReasonCodesSintol MedicationDispenseStatusReasonCodes = "sintol"
	// The drug is contraindicated for patients receiving surgery and the patient is scheduled to be admitted for surgery in the near future. The drug will be resumed when the patient has sufficiently recovered from the surgery.
	MedicationDispenseStatusReasonCodesSurg MedicationDispenseStatusReasonCodes = "surg"
	// The patient was previously receiving a medication contraindicated with the current medication. The current medication will remain on hold until the prior medication has been cleansed from their system.
	MedicationDispenseStatusReasonCodesWashout MedicationDispenseStatusReasonCodes = "washout"
	// Drug out of stock. Cannot supply.
	MedicationDispenseStatusReasonCodesOutofstock MedicationDispenseStatusReasonCodes = "outofstock"
	// Drug no longer marketed Cannot supply.
	MedicationDispenseStatusReasonCodesOffmarket MedicationDispenseStatusReasonCodes = "offmarket"
)
