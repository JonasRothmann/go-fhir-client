// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/medicationrequest-intent
*/type MedicationRequestIntent string

const (
	// The request is a suggestion made by someone/something that doesn't have an intention to ensure it occurs and without providing an authorization to act
	MedicationRequestIntentProposal MedicationRequestIntent = "proposal"
	// The request represents an intention to ensure something occurs without providing an authorization for others to act.
	MedicationRequestIntentPlan MedicationRequestIntent = "plan"
	// The request represents a request/demand and authorization for action
	MedicationRequestIntentOrder MedicationRequestIntent = "order"
	// The request represents a component or option for a RequestOrchestration that establishes timing, conditionality and/or  other constraints among a set of requests.
	MedicationRequestIntentOption MedicationRequestIntent = "option"
)
