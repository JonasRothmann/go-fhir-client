// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/questionnaire-disabled-display
*/type QuestionnaireItemDisabledDisplay string

const (
	// The item (and its children) should not be visible to the user at all.
	QuestionnaireItemDisabledDisplayHidden QuestionnaireItemDisabledDisplay = "hidden"
	// The item (and possibly its children) should not be selectable or editable but should still be visible - to allow the user to see what questions *could* have been completed had other answers caused the item to be enabled.
	QuestionnaireItemDisabledDisplayProtected QuestionnaireItemDisabledDisplay = "protected"
)
