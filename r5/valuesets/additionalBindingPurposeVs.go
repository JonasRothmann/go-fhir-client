// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/additional-binding-purpose
*/type AdditionalBindingPurposeVS string

const (
	// A required binding, for use when the binding strength is 'extensible' or 'preferred'
	AdditionalBindingPurposeVSMaximum AdditionalBindingPurposeVS = "maximum"
	// The minimum allowable value set - any conformant system SHALL support all these codes
	AdditionalBindingPurposeVSMinimum AdditionalBindingPurposeVS = "minimum"
	// This value set is used as a required binding (in addition to the base binding (not a replacement), usually in a particular usage context)
	AdditionalBindingPurposeVSRequired AdditionalBindingPurposeVS = "required"
	// This value set is used as an extensible binding (in addition to the base binding (not a replacement), usually in a particular usage context)
	AdditionalBindingPurposeVSExtensible AdditionalBindingPurposeVS = "extensible"
	// This value set is a candidate to substitute for the overall conformance value set in some situations; usually these are defined in the documentation
	AdditionalBindingPurposeVSCandidate AdditionalBindingPurposeVS = "candidate"
	// New records are required to use this value set, but legacy records may use other codes. The definition of 'new record' is difficult, since systems often create new records based on pre-existing data. Usually 'current' bindings are mandated by an external authority that makes clear rules around this
	AdditionalBindingPurposeVSCurrent AdditionalBindingPurposeVS = "current"
	// This is the value set that is preferred in a given context (documentation should explain why)
	AdditionalBindingPurposeVSPreferred AdditionalBindingPurposeVS = "preferred"
	// This value set is provided for user look up in a given context. Typically, these valuesets only include a subset of codes relevant for input in a context
	AdditionalBindingPurposeVSUi AdditionalBindingPurposeVS = "ui"
	// This value set is a good set of codes to start with when designing your system
	AdditionalBindingPurposeVSStarter AdditionalBindingPurposeVS = "starter"
	// This value set is a component of the base value set. Usually this is called out so that documentation can be written about a portion of the value set
	AdditionalBindingPurposeVSComponent AdditionalBindingPurposeVS = "component"
)
