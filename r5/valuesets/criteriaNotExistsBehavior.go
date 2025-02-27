// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/subscriptiontopic-cr-behavior
*/type CriteriaNotExistsBehavior string

const (
	// The requested conditional statement will pass if a matching state does not exist (e.g., previous state during create).
	CriteriaNotExistsBehaviorTestPasses CriteriaNotExistsBehavior = "test-passes"
	// The requested conditional statement will fail if a matching state does not exist (e.g., previous state during create).
	CriteriaNotExistsBehaviorTestFails CriteriaNotExistsBehavior = "test-fails"
)
