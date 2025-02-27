// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/task-intent
- http://hl7.org/fhir/request-intent
  - proposal
  - plan
  - order
  - original-order
  - reflex-order
  - filler-order
  - instance-order
  - option
*/type TaskIntent string

const (
	// The intent is not known.  When dealing with Task, it's not always known (or relevant) how the task was initiated - i.e. whether it was proposed, planned, ordered or just done spontaneously.
	TaskIntentUnknown TaskIntent = "unknown"
	// The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act.
	TaskIntentProposal TaskIntent = "proposal"
	// The request represents an intention to ensure something occurs without providing an authorization for others to act.
	TaskIntentPlan TaskIntent = "plan"
	// The request represents a request/demand and authorization for action by the requestor.
	TaskIntentOrder TaskIntent = "order"
	// The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestOrchestration]]] for additional information on how this status is used.
	TaskIntentOption TaskIntent = "option"
)
