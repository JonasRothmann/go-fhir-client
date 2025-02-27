// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/definition-method
*/type DefinitionMethod string

const (
	// Any method of routinely determining whether or not specific outcomes (e.g. adverse events) have occurred, for example through a standard questionnaire, regular investigator assessment, regular laboratory testing, or other method
	DefinitionMethodSystematicAssessment DefinitionMethod = "systematic-assessment"
	// Any non-systematic method for determining whether or not adverse events have occurred, such as self-reporting by participants or occasional assessment/testing
	DefinitionMethodNonSystematicAssessment DefinitionMethod = "non-systematic-assessment"
	// Aggregated using mean of observed values.
	DefinitionMethodMean DefinitionMethod = "mean"
	// Aggregated using median of observed values.
	DefinitionMethodMedian DefinitionMethod = "median"
	// Aggregated using mean of means (e.g. study mean values).
	DefinitionMethodMeanOfMeans DefinitionMethod = "mean-of-mean"
	// Aggregated using mean of medians (e.g. study median values).
	DefinitionMethodMeanOfMedians DefinitionMethod = "mean-of-median"
	// Aggregated using median of means (e.g. study mean values).
	DefinitionMethodMedianOfMeans DefinitionMethod = "median-of-mean"
	// Aggregated using median of medians (e.g. study median values).
	DefinitionMethodMedianOfMedians DefinitionMethod = "median-of-median"
)
