// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/biologicallyderivedproductdispense-match-status
*/type BiologicallyDerivedProductDispenseMatchStatus string

const (
	// The product has been serologically or electronically crossmatched for the recipient
	BiologicallyDerivedProductDispenseMatchStatusCrossmatched BiologicallyDerivedProductDispenseMatchStatus = "crossmatched"
	// The product has been selected for the specific use of the recipient
	BiologicallyDerivedProductDispenseMatchStatusSelected BiologicallyDerivedProductDispenseMatchStatus = "selected"
	// No specific matching has been carried out
	BiologicallyDerivedProductDispenseMatchStatusUnmatched BiologicallyDerivedProductDispenseMatchStatus = "unmatched"
	// The product has been selected through crossmatching as least incompatible
	BiologicallyDerivedProductDispenseMatchStatusLeastIncompatible BiologicallyDerivedProductDispenseMatchStatus = "least-incompatible"
)
