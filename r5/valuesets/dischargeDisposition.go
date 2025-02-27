// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/discharge-disposition
*/type DischargeDisposition string

const (
	// The patient was dicharged and has indicated that they are going to return home afterwards.
	DischargeDispositionHome DischargeDisposition = "home"
	// The patient was discharged and has indicated that they are going to return home afterwards, but not the patient's home - e.g. a family member's home.
	DischargeDispositionAlternativeHome DischargeDisposition = "alt-home"
	// The patient was transferred to another healthcare facility.
	DischargeDispositionOtherHealthcareFacility DischargeDisposition = "other-hcf"
	// The patient has been discharged into palliative care.
	DischargeDispositionHosp DischargeDisposition = "hosp"
	// The patient has been discharged into long-term care where is likely to be monitored through an ongoing episode-of-care.
	DischargeDispositionLong DischargeDisposition = "long"
	// The patient self discharged against medical advice.
	DischargeDispositionAadvice DischargeDisposition = "aadvice"
	// The patient has deceased during this encounter.
	DischargeDispositionExp DischargeDisposition = "exp"
	// The patient has been transferred to a psychiatric facility.
	DischargeDispositionPsy DischargeDisposition = "psy"
	// The patient was discharged and is to receive post acute care rehabilitation services.
	DischargeDispositionRehab DischargeDisposition = "rehab"
	// The patient has been discharged to a skilled nursing facility for the patient to receive additional care.
	DischargeDispositionSnf DischargeDisposition = "snf"
	// The discharge disposition has not otherwise defined.
	DischargeDispositionOth DischargeDisposition = "oth"
)
