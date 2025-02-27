// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/chargeitem-status
*/type ChargeItemStatus string

const (
	// The charge item has been entered, but the charged service is not  yet complete, so it shall not be billed yet but might be used in the context of pre-authorization.
	ChargeItemStatusPlanned ChargeItemStatus = "planned"
	// The charge item is ready for billing.
	ChargeItemStatusBillable ChargeItemStatus = "billable"
	// The charge item has been determined to be not billable (e.g. due to rules associated with the billing code).
	ChargeItemStatusNotBillable ChargeItemStatus = "not-billable"
	// The processing of the charge was aborted.
	ChargeItemStatusAborted ChargeItemStatus = "aborted"
	// The charge item has been billed (e.g. a billing engine has generated financial transactions by applying the associated ruled for the charge item to the context of the Encounter, and placed them into Claims/Invoices.
	ChargeItemStatusBilled ChargeItemStatus = "billed"
	// The charge item has been entered in error and should not be processed for billing.
	ChargeItemStatusEnteredInError ChargeItemStatus = "entered-in-error"
	// The authoring system does not know which of the status values currently applies for this charge item  Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	ChargeItemStatusUnknown ChargeItemStatus = "unknown"
)
