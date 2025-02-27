// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/nutritionproduct-status
*/type NutritionProductStatus string

const (
	// The product can be used.
	NutritionProductStatusActive NutritionProductStatus = "active"
	// The product is not expected or allowed to be used.
	NutritionProductStatusInactive NutritionProductStatus = "inactive"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	NutritionProductStatusEnteredInError NutritionProductStatus = "entered-in-error"
)
