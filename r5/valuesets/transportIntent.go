// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/transport-intent
- http://hl7.org/fhir/request-intent
  - proposal
  - plan
  - order
  - original-order
  - reflex-order
  - filler-order
  - instance-order
  - option
*/type TransportIntent string

const (
	// The intent is not known.  When dealing with Transport, it's not always known (or relevant) how the transport was initiated - i.e. whether it was proposed, planned, ordered or just done spontaneously.
	TransportIntentUnknown TransportIntent = "unknown"
	// The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act.
	TransportIntentProposal TransportIntent = "proposal"
	// The request represents an intention to ensure something occurs without providing an authorization for others to act.
	TransportIntentPlan TransportIntent = "plan"
	// The request represents a request/demand and authorization for action by the requestor.
	TransportIntentOrder TransportIntent = "order"
	// The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestOrchestration]]] for additional information on how this status is used.
	TransportIntentOption TransportIntent = "option"
)
