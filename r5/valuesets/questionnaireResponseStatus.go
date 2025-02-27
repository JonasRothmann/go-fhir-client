// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/questionnaire-answers-status
*/type QuestionnaireResponseStatus string

const (
	// This QuestionnaireResponse has been partially filled out with answers but changes or additions are still expected to be made to it.
	QuestionnaireResponseStatusInProgress QuestionnaireResponseStatus = "in-progress"
	// This QuestionnaireResponse has been filled out with answers and the current content is regarded as definitive.
	QuestionnaireResponseStatusCompleted QuestionnaireResponseStatus = "completed"
	// This QuestionnaireResponse has been filled out with answers, then marked as complete, yet changes or additions have been made to it afterwards.
	QuestionnaireResponseStatusAmended QuestionnaireResponseStatus = "amended"
	// This QuestionnaireResponse was entered in error and voided.
	QuestionnaireResponseStatusEnteredInError QuestionnaireResponseStatus = "entered-in-error"
	// This QuestionnaireResponse has been partially filled out with answers but has been abandoned. No subsequent changes can be made.
	QuestionnaireResponseStatusStopped QuestionnaireResponseStatus = "stopped"
)
