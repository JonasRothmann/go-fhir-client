// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/measure-aggregate-method
*/type MeasureAggregateMethod string

const (
	// The measure score is determined by adding together the observations derived from the measure population.
	MeasureAggregateMethodSum MeasureAggregateMethod = "sum"
	// The measure score is determined by taking the average of the observations derived from the measure population.
	MeasureAggregateMethodAverage MeasureAggregateMethod = "average"
	// The measure score is determined by taking the median of the observations derived from the measure population.
	MeasureAggregateMethodMedian MeasureAggregateMethod = "median"
	// The measure score is determined by taking the minimum of the observations derived from the measure population.
	MeasureAggregateMethodMinimum MeasureAggregateMethod = "minimum"
	// The measure score is determined by taking the maximum of the observations derived from the measure population.
	MeasureAggregateMethodMaximum MeasureAggregateMethod = "maximum"
	// The measure score is determined as the number of observations derived from the measure population.
	MeasureAggregateMethodCount MeasureAggregateMethod = "count"
)
