// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/payment-type
*/type PaymentTypeCodes string

const (
	// The amount is partial or complete settlement of the amounts due.
	PaymentTypeCodesPayment PaymentTypeCodes = "payment"
	// The amount is an adjustment regarding claims already paid.
	PaymentTypeCodesAdjustment PaymentTypeCodes = "adjustment"
	// The amount is an advance against future claims.
	PaymentTypeCodesAdvance PaymentTypeCodes = "advance"
)
