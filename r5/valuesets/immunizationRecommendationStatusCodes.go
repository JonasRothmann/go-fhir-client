// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/immunization-recommendation-status
*/type ImmunizationRecommendationStatusCodes string

const (
	// The patient is due for their next vaccination.
	ImmunizationRecommendationStatusCodesDue ImmunizationRecommendationStatusCodes = "due"
	// The patient is considered overdue for their next vaccination.
	ImmunizationRecommendationStatusCodesOverdue ImmunizationRecommendationStatusCodes = "overdue"
	// The patient is immune to the target disease and further immunization against the disease is not likely to provide benefit.
	ImmunizationRecommendationStatusCodesImmune ImmunizationRecommendationStatusCodes = "immune"
	// The patient is contraindicated for futher doses.
	ImmunizationRecommendationStatusCodesContraindicated ImmunizationRecommendationStatusCodes = "contraindicated"
	// The patient is fully protected and no further doses are recommended.
	ImmunizationRecommendationStatusCodesComplete ImmunizationRecommendationStatusCodes = "complete"
)
