// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/payment-kind
*/type PaymentKind string

const (
	// The payment or adjustment is to an indicated account not to a specific charge.
	PaymentKindDeposit PaymentKind = "deposit"
	// The payment is one of a set of previously agreed payments, for example in fullfilment of a payment plan.
	PaymentKindPeriodicPayment PaymentKind = "periodic-payment"
	// Payment, full or partial, of an invoice or statement provided to the payment issuer.
	PaymentKindOnline PaymentKind = "online"
	// Payment made at an authorized Kiosk.
	PaymentKindKiosk PaymentKind = "kiosk"
)
