// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/v3-ActReason
  - TREAT
  - HPAYMT
  - ETREAT

- http://terminology.hl7.org/CodeSystem/v3-ActCode
  - NOAUTH
  - DELAU
  - NORDSCLCD
*/type SecurityLabelEventExamples string

const (
	// To perform one or more operations on information for conducting financial or contractual activities related to payment for provision of health care.
	SecurityLabelEventExamplesHealthcarePayment SecurityLabelEventExamples = "HPAYMT"
	// To perform one or more operations on information for provision of health care.
	SecurityLabelEventExamplesTreatment SecurityLabelEventExamples = "TREAT"
	// To perform one or more operations on information for provision of immediately needed health care for an emergent condition.
	SecurityLabelEventExamplesEmergencyTreatment SecurityLabelEventExamples = "ETREAT"
	// Custodian system must remove target information from access after use.
	SecurityLabelEventExamplesDeleteAfterUse SecurityLabelEventExamples = "DELAU"
	// Prohibition on disclosure without information subject's authorization.
	SecurityLabelEventExamplesNoDisclosureWithoutSubjectAuthorization SecurityLabelEventExamples = "NOAUTH"
	// Prohibition on redisclosure without patient consent directive.
	SecurityLabelEventExamplesNoRedisclosureWithoutConsentDirective SecurityLabelEventExamples = "NORDSCLCD"
)
