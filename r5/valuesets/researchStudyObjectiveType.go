// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/research-study-objective-type
*/type ResearchStudyObjectiveType string

const (
	// The main question to be answered, and the one that drives any statistical planning for the study—e.g., calculation of the sample size to provide the appropriate power for statistical testing.
	ResearchStudyObjectiveTypePrimary ResearchStudyObjectiveType = "primary"
	// Question to be answered in the study that is of lesser importance than the primary objective.
	ResearchStudyObjectiveTypeSecondary ResearchStudyObjectiveType = "secondary"
	// Exploratory questions to be answered in the study.
	ResearchStudyObjectiveTypeExploratory ResearchStudyObjectiveType = "exploratory"
)
