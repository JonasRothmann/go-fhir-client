// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/capability-statement-kind
*/type CapabilityStatementKind string

const (
	// The CapabilityStatement instance represents the present capabilities of a specific system instance.  This is the kind returned by /metadata for a FHIR server end-point.
	CapabilityStatementKindInstance CapabilityStatementKind = "instance"
	// The CapabilityStatement instance represents the capabilities of a system or piece of software, independent of a particular installation.
	CapabilityStatementKindCapability CapabilityStatementKind = "capability"
	// The CapabilityStatement instance represents a set of requirements for other systems to meet; e.g. as part of an implementation guide or 'request for proposal'.
	CapabilityStatementKindRequirements CapabilityStatementKind = "requirements"
)
