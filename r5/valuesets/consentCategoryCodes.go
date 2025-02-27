// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/consentscope
- http://terminology.hl7.org/CodeSystem/consentcategorycodes
- http://loinc.org
  - 59284-0
  - 57016-8
  - 57017-6
  - 64292-6
*/type ConsentCategoryCodes string

const (
	// Actions to be taken if they are no longer able to make decisions for themselves
	ConsentCategoryCodesAdr ConsentCategoryCodes = "adr"
	// Consent to participate in research protocol and information sharing required
	ConsentCategoryCodesResearch ConsentCategoryCodes = "research"
	// Agreement to collect, access, use or disclose (share) information
	ConsentCategoryCodesPrivacyConsent ConsentCategoryCodes = "patient-privacy"
	// Consent to undergo a specific treatment
	ConsentCategoryCodesTreatment ConsentCategoryCodes = "treatment"
	// Any instructions, written or given verbally by a patient in anticipation of potential need for medical treatment.
	ConsentCategoryCodesAcd ConsentCategoryCodes = "acd"
	// A legal document, signed by both the patient and their provider, stating a desire not to have CPR initiated in case of a cardiac event.
	ConsentCategoryCodesDnr ConsentCategoryCodes = "dnr"
	// Opt-in to disclosure of health information for emergency only consent directive. Comment: This general consent directive specifically limits disclosure of health information for purpose of emergency treatment. Additional parameters may further limit the disclosure to specific users, roles, duration, types of information, and impose uses obligations.
	ConsentCategoryCodesEmrgonly ConsentCategoryCodes = "emrgonly"
	// Patient authored document communicating to patient's health care provider(s) instructions about the patient's goals, preferences, and priorities if the patient is diagnosed as being terminally ill and in a persistent vegetative state, is in a permanently unconscious condition or is otherwise unable to communicate thoe instructions.
	ConsentCategoryCodesHcd ConsentCategoryCodes = "hcd"
	// Acknowledgement of custodian notice of privacy practices. Usage Notes: This type of consent directive acknowledges a custodian's notice of privacy practices including its permitted collection, access, use and disclosure of health information to users and for purposes of use specified.
	ConsentCategoryCodesNpp ConsentCategoryCodes = "npp"
	// The Physician Order for Life-Sustaining Treatment form records a person's health care wishes for end of life emergency treatment and translates them into an order by the physician. It must be reviewed and signed by both the patient and the physician, Advanced Registered Nurse Practitioner or Physician Assistant.
	ConsentCategoryCodesPolst ConsentCategoryCodes = "polst"
	// Consent to have healthcare information in an electronic health record accessed for research purposes.
	ConsentCategoryCodesResearch2 ConsentCategoryCodes = "research"
	// Consent to have de-identified healthcare information in an electronic health record that is accessed for research purposes, but without consent to re-identify the information under any circumstance.
	ConsentCategoryCodesRsdid ConsentCategoryCodes = "rsdid"
	// Consent to have de-identified healthcare information in an electronic health record that is accessed for research purposes re-identified under specific circumstances outlined in the consent.
	ConsentCategoryCodesRsreid ConsentCategoryCodes = "rsreid"
)
