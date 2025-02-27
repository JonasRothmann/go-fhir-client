// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/coverage-copay-type
*/type CoverageCopayTypeCodes string

const (
	// An office visit for a general practitioner of a discipline.
	CoverageCopayTypeCodesGpvisit CoverageCopayTypeCodes = "gpvisit"
	// An office visit for a specialist practitioner of a discipline
	CoverageCopayTypeCodesSpvisit CoverageCopayTypeCodes = "spvisit"
	// An episode in an emergency department.
	CoverageCopayTypeCodesEmergency CoverageCopayTypeCodes = "emergency"
	// An episode of an Inpatient hospital stay.
	CoverageCopayTypeCodesInpthosp CoverageCopayTypeCodes = "inpthosp"
	// A visit held where the patient is remote relative to the practitioner, e.g. by phone, computer or video conference.
	CoverageCopayTypeCodesTelevisit CoverageCopayTypeCodes = "televisit"
	// A visit to an urgent care facility - typically a community care clinic.
	CoverageCopayTypeCodesUrgentcare CoverageCopayTypeCodes = "urgentcare"
	// A standard percentage applied to all classes or service or product not otherwise specified.
	CoverageCopayTypeCodesCopaypct CoverageCopayTypeCodes = "copaypct"
	// A standard fixed currency amount applied to all classes or service or product not otherwise specified.
	CoverageCopayTypeCodesCopay CoverageCopayTypeCodes = "copay"
	// The accumulated amount of patient payment before the coverage begins to pay for services.
	CoverageCopayTypeCodesDeductible CoverageCopayTypeCodes = "deductible"
	// The maximum amout of payment for services which a patient, or family, is expected to incur - typically annually.
	CoverageCopayTypeCodesMaxoutofpocket CoverageCopayTypeCodes = "maxoutofpocket"
)
