// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/item-type
*/type QuestionnaireItemType string

const (
	// An item with no direct answer but should have at least one child item.
	QuestionnaireItemTypeGroup QuestionnaireItemType = "group"
	// Text for display that will not capture an answer or have child items.
	QuestionnaireItemTypeDisplay QuestionnaireItemType = "display"
	// An item that defines a specific answer to be captured, and which may have child items. (the answer provided in the QuestionnaireResponse should be of the defined datatype).
	QuestionnaireItemTypeQuestion QuestionnaireItemType = "question"
)
