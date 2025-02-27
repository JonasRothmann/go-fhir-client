// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/questionnaire-answer-constraint
*/type QuestionnaireAnswerConstraint string

const (
	// Only values listed as answerOption or in the expansion of the answerValueSet are permitted
	QuestionnaireAnswerConstraintOptionsOnly QuestionnaireAnswerConstraint = "optionsOnly"
	// In addition to the values listed as answerOption or in the expansion of the answerValueSet, any other values that correspond to the specified item.type are permitted
	QuestionnaireAnswerConstraintOptionsOrType QuestionnaireAnswerConstraint = "optionsOrType"
	// In addition to the values listed as answerOption or in the expansion of the answerValueSet, free-text strings are permitted.  Answers will have a type of 'string'.
	QuestionnaireAnswerConstraintOptionsOrString QuestionnaireAnswerConstraint = "optionsOrString"
)
