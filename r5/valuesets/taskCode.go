// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/task-code
*/type TaskCode string

const (
	// Take what actions are needed to transition the focus resource from 'draft' to 'active' or 'in-progress', as appropriate for the resource type.  This may involve additing additional content, approval, validation, etc.
	TaskCodeApprove TaskCode = "approve"
	// Act to perform the actions described in the focus request.  This might result in a 'more assertive' request (order for a plan or proposal, filler order for a placer order), but is intend to eventually result in events.  The degree of fulfillment requested might be limited by Task.restriction.
	TaskCodeFulfill TaskCode = "fulfill"
	// Act to perform the actions defined in the focus definition resource (ActivityDefinition, PlanDefinition, Questionnaire, etc.)  For PlanDefinition or ActivityDefinition, this might result in a 'more assertive' request (order for a plan or proposal, filler order for a placer order), but is intend to eventually result in events.  For Questionnaire, this would result in a QuestionnaireResponse - and possibly resources constructed using data extracted from the response.  The degree of fulfillment requested might be limited by Task.restriction.
	TaskCodeInstantiate TaskCode = "instantiate"
	// Abort, cancel or withdraw the focal resource, as appropriate for the type of resource.
	TaskCodeAbort TaskCode = "abort"
	// Replace the focal resource with the specified input resource
	TaskCodeReplace TaskCode = "replace"
	// Update the focal resource of the owning system to reflect the content specified as the Task.focus
	TaskCodeChange TaskCode = "change"
	// Transition the focal resource from 'active' or 'in-progress' to 'suspended'
	TaskCodeSuspend TaskCode = "suspend"
	// Transition the focal resource from 'suspended' to 'active' or 'in-progress' as appropriate for the resource type.
	TaskCodeResume TaskCode = "resume"
)
