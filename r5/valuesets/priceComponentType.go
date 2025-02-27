// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/price-component-type
*/type PriceComponentType string

const (
	// the amount is the base price used for calculating the total price before applying surcharges, discount or taxes.
	PriceComponentTypeBase PriceComponentType = "base"
	// the amount is a surcharge applied on the base price.
	PriceComponentTypeSurcharge PriceComponentType = "surcharge"
	// the amount is a deduction applied on the base price.
	PriceComponentTypeDeduction PriceComponentType = "deduction"
	// the amount is a discount applied on the base price.
	PriceComponentTypeDiscount PriceComponentType = "discount"
	// the amount is the tax component of the total price.
	PriceComponentTypeTax PriceComponentType = "tax"
	// the amount is of informational character, it has not been applied in the calculation of the total price.
	PriceComponentTypeInformational PriceComponentType = "informational"
)
