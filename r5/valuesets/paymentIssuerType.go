// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/payment-issuertype
*/type PaymentIssuerType string

const (
	// The patient or a party issuing payment on behalf of the patient.
	PaymentIssuerTypePatient PaymentIssuerType = "patient"
	// An insurer, or party acting on their behalf, which is making payment following a contract, direct or indirect, with the patient to pay for healthcare-related services.
	PaymentIssuerTypeInsurance PaymentIssuerType = "insurance"
)
