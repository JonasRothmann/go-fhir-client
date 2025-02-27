// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/address-use
*/type AddressUse string

const (
	// A communication address at a home.
	AddressUseHome AddressUse = "home"
	// An office address. First choice for business related contacts during business hours.
	AddressUseWork AddressUse = "work"
	// A temporary address. The period can provide more detailed information.
	AddressUseTemp AddressUse = "temp"
	// This address is no longer in use (or was never correct but retained for records).
	AddressUseOld AddressUse = "old"
	// An address to be used to send bills, invoices, receipts etc.
	AddressUseBilling AddressUse = "billing"
)
