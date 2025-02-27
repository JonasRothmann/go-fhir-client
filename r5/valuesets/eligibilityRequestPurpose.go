// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/eligibilityrequest-purpose
*/type EligibilityRequestPurpose string

const (
	// The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested.
	EligibilityRequestPurposeCoverageAuthRequirements EligibilityRequestPurpose = "auth-requirements"
	// The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested.
	EligibilityRequestPurposeBenefits EligibilityRequestPurpose = "benefits"
	// The insurer is requested to report on any coverages which they are aware of in addition to any specifed.
	EligibilityRequestPurposeDiscovery EligibilityRequestPurpose = "discovery"
	// A check that the specified coverages are in-force is requested.
	EligibilityRequestPurposeValidation EligibilityRequestPurpose = "validation"
)
