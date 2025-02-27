// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/observation-triggeredbytype
*/type TriggeredBytype string

const (
	// Performance of one or more other tests depending on the results of the initial test.  This may include collection of additional specimen. While a new ServiceRequest is not required to perform the additional test, where it is still needed (e.g., requesting another laboratory to perform the reflex test), the Observation.basedOn would reference the new ServiceRequest that requested the additional test to be performed as well as the original ServiceRequest to reflect the one that provided the authorization.
	TriggeredBytypeReflex TriggeredBytype = "reflex"
	// Performance of the same test again with the same parameters/settings/solution.
	TriggeredBytypeRepeat TriggeredBytype = "repeat"
	// Performance of the same test but with different parameters/settings/solution.
	TriggeredBytypeReRunPerPolicy TriggeredBytype = "re-run"
)
