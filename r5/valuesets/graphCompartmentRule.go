// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/graph-compartment-rule
*/type GraphCompartmentRule string

const (
	// The compartment must be identical (the same literal reference).
	GraphCompartmentRuleIdentical GraphCompartmentRule = "identical"
	// The compartment must be the same - the record must be about the same patient, but the reference may be different.
	GraphCompartmentRuleMatching GraphCompartmentRule = "matching"
	// The compartment must be different.
	GraphCompartmentRuleDifferent GraphCompartmentRule = "different"
	// The compartment rule is defined in the accompanying FHIRPath expression.
	GraphCompartmentRuleCustom GraphCompartmentRule = "custom"
)
