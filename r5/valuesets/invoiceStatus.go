// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/invoice-status
*/type InvoiceStatus string

const (
	// the invoice has been prepared but not yet finalized.
	InvoiceStatusDraft InvoiceStatus = "draft"
	// the invoice has been finalized and sent to the recipient.
	InvoiceStatusIssued InvoiceStatus = "issued"
	// the invoice has been balaced / completely paid.
	InvoiceStatusBalanced InvoiceStatus = "balanced"
	// the invoice was cancelled.
	InvoiceStatusCancelled InvoiceStatus = "cancelled"
	// the invoice was determined as entered in error before it was issued.
	InvoiceStatusEnteredInError InvoiceStatus = "entered-in-error"
)
