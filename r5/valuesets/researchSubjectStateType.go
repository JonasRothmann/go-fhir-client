// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/research-subject-state-type
*/type ResearchSubjectStateType string

const (
	// One of a set of milestone events that once they have occurred remain true thereafter.  For example once a subject has reached the "Signed Informed Consent" milestone they achieve a status of "Consented" that will be true thereafter, even when they have left the study.  For a subject a number of these states can be simulataneously true and should be recorded.
	ResearchSubjectStateTypeMilestone ResearchSubjectStateType = "Milestone"
	// This is a status that can only have a single value at a given point in time.  The subject is either on-study or off-study.
	ResearchSubjectStateTypeEnrollment ResearchSubjectStateType = "Enrollment"
)
