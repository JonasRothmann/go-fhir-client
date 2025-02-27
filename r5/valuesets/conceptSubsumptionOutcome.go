// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/concept-subsumption-outcome
*/type ConceptSubsumptionOutcome string

const (
	// A equivalent to B if A subsumes B and B subsumes A
	ConceptSubsumptionOutcomeEquivalent ConceptSubsumptionOutcome = "equivalent"
	// A subsumes B if there is a subsumption relationship between A and B
	ConceptSubsumptionOutcomeSubsumes ConceptSubsumptionOutcome = "subsumes"
	// A subsumed by B if B subsumes A
	ConceptSubsumptionOutcomeSubsumedBy ConceptSubsumptionOutcome = "subsumed-by"
	// Neither A subsumes B nor B subsumes A
	ConceptSubsumptionOutcomeNotSubsumed ConceptSubsumptionOutcome = "not-subsumed"
)
