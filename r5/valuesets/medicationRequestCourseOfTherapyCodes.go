// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/medicationrequest-course-of-therapy
*/type MedicationRequestCourseOfTherapyCodes string

const (
	// A medication which is expected to be continued beyond the present order and which the patient should be assumed to be taking unless explicitly stopped.
	MedicationRequestCourseOfTherapyCodesContinuous MedicationRequestCourseOfTherapyCodes = "continuous"
	// A medication which the patient is only expected to consume for the duration of the current order and which is not expected to be renewed.
	MedicationRequestCourseOfTherapyCodesAcute MedicationRequestCourseOfTherapyCodes = "acute"
	// A medication which is expected to be used on a part time basis at certain times of the year
	MedicationRequestCourseOfTherapyCodesSeasonal MedicationRequestCourseOfTherapyCodes = "seasonal"
)
