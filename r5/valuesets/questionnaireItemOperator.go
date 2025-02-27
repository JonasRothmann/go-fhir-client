// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/questionnaire-enable-operator
*/type QuestionnaireItemOperator string

const (
	// True if the determination of 'whether an answer exists for the question' is equal to the enableWhen answer (which must be a boolean).
	QuestionnaireItemOperatorExists QuestionnaireItemOperator = "exists"
	// True if at least one answer has a value that is equal to the enableWhen answer.
	QuestionnaireItemOperatorEquals QuestionnaireItemOperator = "="
	// True if no answer has a value that is equal to the enableWhen answer.
	QuestionnaireItemOperatorNotEquals QuestionnaireItemOperator = "!="
	// True if at least one answer has a value that is greater than the enableWhen answer.
	QuestionnaireItemOperatorGreaterThan QuestionnaireItemOperator = ">"
	// True if at least one answer has a value that is less than the enableWhen answer.
	QuestionnaireItemOperatorLessThan QuestionnaireItemOperator = "<"
	// True if at least one answer has a value that is greater or equal to the enableWhen answer.
	QuestionnaireItemOperatorGreaterOrEquals QuestionnaireItemOperator = ">="
	// True if at least one answer has a value that is less or equal to the enableWhen answer.
	QuestionnaireItemOperatorLessOrEquals QuestionnaireItemOperator = "<="
)
