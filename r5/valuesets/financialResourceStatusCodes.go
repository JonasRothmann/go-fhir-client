// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/fm-status
*/type FinancialResourceStatusCodes string

const (
	// The instance is currently in-force.
	FinancialResourceStatusCodesActive FinancialResourceStatusCodes = "active"
	// The instance is withdrawn, rescinded or reversed.
	FinancialResourceStatusCodesCancelled FinancialResourceStatusCodes = "cancelled"
	// A new instance the contents of which is not complete.
	FinancialResourceStatusCodesDraft FinancialResourceStatusCodes = "draft"
	// The instance was entered in error.
	FinancialResourceStatusCodesEnteredInError FinancialResourceStatusCodes = "entered-in-error"
)
