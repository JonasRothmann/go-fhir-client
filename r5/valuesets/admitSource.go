// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/admit-source
*/type AdmitSource string

const (
	// The Patient has been transferred from another hospital for this encounter.
	AdmitSourceTransferredFromOtherHospital AdmitSource = "hosp-trans"
	// The patient has been transferred from the emergency department within the hospital. This is typically used in the transition to an inpatient encounter
	AdmitSourceEmd AdmitSource = "emd"
	// The patient has been transferred from an outpatient department within the hospital.
	AdmitSourceOutp AdmitSource = "outp"
	// The patient is a newborn and the encounter will track the baby related activities (as opposed to the Mothers encounter - that may be associated using the newborn encounters partof property)
	AdmitSourceBorn AdmitSource = "born"
	// The patient has been admitted due to a referred from a General Practitioner.
	AdmitSourceGp AdmitSource = "gp"
	// The patient has been admitted due to a referred from a Specialist (as opposed to a General Practitioner).
	AdmitSourceMp AdmitSource = "mp"
	// The patient has been transferred from a nursing home.
	AdmitSourceNursing AdmitSource = "nursing"
	// The patient has been transferred from a psychiatric facility.
	AdmitSourcePsych AdmitSource = "psych"
	// The patient has been transferred from a rehabilitation facility or clinic.
	AdmitSourceRehab AdmitSource = "rehab"
	// The patient has been admitted from a source otherwise not specified here.
	AdmitSourceOther AdmitSource = "other"
)
