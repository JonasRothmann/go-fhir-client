// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/eligibilityresponse-purpose
*/type EligibilityResponsePurpose string

const (
	// The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested.
	EligibilityResponsePurposeCoverageAuthRequirements EligibilityResponsePurpose = "auth-requirements"
	// The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested.
	EligibilityResponsePurposeBenefits EligibilityResponsePurpose = "benefits"
	// The insurer is requested to report on any coverages which they are aware of in addition to any specifed.
	EligibilityResponsePurposeDiscovery EligibilityResponsePurpose = "discovery"
	// A check that the specified coverages are in-force is requested.
	EligibilityResponsePurposeValidation EligibilityResponsePurpose = "validation"
)
