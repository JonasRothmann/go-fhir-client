// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/claiminformationcategory
*/type ClaimInformationCategoryCodes string

const (
	// Codes conveying additional situation and condition information.
	ClaimInformationCategoryCodesInfo ClaimInformationCategoryCodes = "info"
	// Discharge status and discharge to locations.
	ClaimInformationCategoryCodesDischarge ClaimInformationCategoryCodes = "discharge"
	// Period, start or end dates of aspects of the Condition.
	ClaimInformationCategoryCodesOnset ClaimInformationCategoryCodes = "onset"
	// Nature and date of the related event e.g. Last exam, service, X-ray etc.
	ClaimInformationCategoryCodesRelated ClaimInformationCategoryCodes = "related"
	// Insurance policy exceptions.
	ClaimInformationCategoryCodesException ClaimInformationCategoryCodes = "exception"
	// Materials being forwarded, e.g. Models, molds, images, documents.
	ClaimInformationCategoryCodesMaterial ClaimInformationCategoryCodes = "material"
	// Materials attached such as images, documents and resources.
	ClaimInformationCategoryCodesAttachment ClaimInformationCategoryCodes = "attachment"
	// Teeth which are missing for any reason, for example: prior extraction, never developed.
	ClaimInformationCategoryCodesMissingtooth ClaimInformationCategoryCodes = "missingtooth"
	// The type of prosthesis and date of supply if a previously supplied prosthesis.
	ClaimInformationCategoryCodesProsthesis ClaimInformationCategoryCodes = "prosthesis"
	// Other information identified by the type.system.
	ClaimInformationCategoryCodesOther ClaimInformationCategoryCodes = "other"
	// An indication that the patient was hospitalized, the period if known otherwise a Yes/No (boolean).
	ClaimInformationCategoryCodesHospitalized ClaimInformationCategoryCodes = "hospitalized"
	// An indication that the patient was unable to work, the period if known otherwise a Yes/No (boolean).
	ClaimInformationCategoryCodesEmploymentimpacted ClaimInformationCategoryCodes = "employmentimpacted"
	// The external cause of an illness or injury.
	ClaimInformationCategoryCodesExternalcause ClaimInformationCategoryCodes = "externalcause"
	// The reason for the patient visit.
	ClaimInformationCategoryCodesPatientreasonforvisit ClaimInformationCategoryCodes = "patientreasonforvisit"
)
