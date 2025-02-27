// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/request-intent
*/type RequestIntent string

const (
	// The request is a suggestion made by someone/something that does not have an intention to ensure it occurs and without providing an authorization to act.
	RequestIntentProposal RequestIntent = "proposal"
	// The request represents an intention to ensure something occurs without providing an authorization for others to act.
	RequestIntentPlan RequestIntent = "plan"
	// The request represents a legally binding instruction authored by a Patient or RelatedPerson.
	RequestIntentDirective RequestIntent = "directive"
	// The request represents a request/demand and authorization for action by the requestor.
	RequestIntentOrder RequestIntent = "order"
	// The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or other constraints among a set of requests.  Refer to [[[RequestOrchestration]]] for additional information on how this status is used.
	RequestIntentOption RequestIntent = "option"
)
