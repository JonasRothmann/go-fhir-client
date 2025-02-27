// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/product-status
*/type ProductStatus string

const (
	// The product can be used.
	ProductStatusActive ProductStatus = "active"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	ProductStatusEnteredInError ProductStatus = "entered-in-error"
)
