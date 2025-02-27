// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/adverse-event-category
*/type AdverseEventCategory string

const (
	// The adverse event pertains to an event involving the wrong patient, who was not the intended subject.
	AdverseEventCategoryWrongPatient AdverseEventCategory = "wrong-patient"
	// The adverse event pertains to a procedure mishap.
	AdverseEventCategoryProcedureMishap AdverseEventCategory = "procedure-mishap"
	// The adverse event pertains to a medication mishap, such as wrong dose, route, rate, or duration of administration.
	AdverseEventCategoryMedicationMishap AdverseEventCategory = "medication-mishap"
	// The adverse event pertains to a device.
	AdverseEventCategoryDevice AdverseEventCategory = "device"
	// The adverse event pertains to an unsafe physical environment.
	AdverseEventCategoryUnsafePhysicalEnvironment AdverseEventCategory = "unsafe-physical-environment"
	// The adverse event pertains to a hospital acquired infection.
	AdverseEventCategoryHospitalAcquiredInfection AdverseEventCategory = "hospital-aquired-infection"
	// The adverse event pertains to the wrong body site.
	AdverseEventCategoryWrongBodySite AdverseEventCategory = "wrong-body-site"
)
