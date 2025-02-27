// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/history-absent-reason
*/type FamilyHistoryAbsentReason string

const (
	// Patient does not know the subject, e.g. the biological parent of an adopted patient.
	FamilyHistoryAbsentReasonSubjectUnknown FamilyHistoryAbsentReason = "subject-unknown"
	// The patient withheld or refused to share the information.
	FamilyHistoryAbsentReasonWithheld FamilyHistoryAbsentReason = "withheld"
	// Information cannot be obtained; e.g. unconscious patient.
	FamilyHistoryAbsentReasonUnableToObtain FamilyHistoryAbsentReason = "unable-to-obtain"
	// Patient does not have the information now, but can provide the information at a later date.
	FamilyHistoryAbsentReasonDeferred FamilyHistoryAbsentReason = "deferred"
)
