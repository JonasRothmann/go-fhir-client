// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/claim-use
*/type Use string

const (
	// The treatment is complete and this represents a Claim for the services.
	UseClaim Use = "claim"
	// The treatment is proposed and this represents a Pre-authorization for the services.
	UsePreauthorization Use = "preauthorization"
	// The treatment is proposed and this represents a Pre-determination for the services.
	UsePredetermination Use = "predetermination"
)
