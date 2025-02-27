// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/adjudication
*/type AdjudicationValueCodes string

const (
	// The total submitted amount for the claim or group or line item.
	AdjudicationValueCodesSubmitted AdjudicationValueCodes = "submitted"
	// Patient Co-Payment
	AdjudicationValueCodesCopay AdjudicationValueCodes = "copay"
	// Amount of the change which is considered for adjudication.
	AdjudicationValueCodesEligible AdjudicationValueCodes = "eligible"
	// Amount deducted from the eligible amount prior to adjudication.
	AdjudicationValueCodesDeductible AdjudicationValueCodes = "deductible"
	// The amount of deductible which could not allocated to other line items.
	AdjudicationValueCodesUnallocdeduct AdjudicationValueCodes = "unallocdeduct"
	// Eligible Percentage.
	AdjudicationValueCodesEligpercent AdjudicationValueCodes = "eligpercent"
	// The amount of tax.
	AdjudicationValueCodesTax AdjudicationValueCodes = "tax"
	// Amount payable under the coverage
	AdjudicationValueCodesBenefit AdjudicationValueCodes = "benefit"
)
