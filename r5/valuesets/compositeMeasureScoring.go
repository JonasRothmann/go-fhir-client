// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/composite-measure-scoring
*/type CompositeMeasureScoring string

const (
	// Opportunity scoring combines the scores from component measures by combining the numerators and denominators for each component.
	CompositeMeasureScoringOpportunity CompositeMeasureScoring = "opportunity"
	// All-or-nothing scoring includes an individual in the numerator of the composite measure if they are in the numerators of all of the component measures in which they are in the denominator.
	CompositeMeasureScoringAllOrNothing CompositeMeasureScoring = "all-or-nothing"
	// Linear scoring gives an individual a score based on the number of numerators in which they appear.
	CompositeMeasureScoringLinear CompositeMeasureScoring = "linear"
	// Weighted scoring gives an individual a score based on a weighted factor for each component numerator in which they appear.
	CompositeMeasureScoringWeighted CompositeMeasureScoring = "weighted"
)
