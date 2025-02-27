// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/account-aggregate
*/type AccountAggregate string

const (
	// This (aggregated) balance is expected to be paid by the Patient
	AccountAggregatePatient AccountAggregate = "patient"
	// This (aggregated) balance is expected to be paid by Insurance coverage(s)
	AccountAggregateInsurance AccountAggregate = "insurance"
	// There is no aggregation on this balance
	AccountAggregateTotal AccountAggregate = "total"
)
