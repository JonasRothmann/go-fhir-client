// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/provenance-entity-role
*/type ProvenanceEntityRole string

const (
	// An entity that is used by the activity to produce a new version of that entity.
	ProvenanceEntityRoleRevision ProvenanceEntityRole = "revision"
	// An entity that is copied in full or part by an agent that is not the author of the entity.
	ProvenanceEntityRoleQuotation ProvenanceEntityRole = "quotation"
	// An entity that is used as input to the activity that produced the target.
	ProvenanceEntityRoleSource ProvenanceEntityRole = "source"
	// The record resulting from this event adheres to the protocol, guideline, order set or other definition represented by this entity.
	ProvenanceEntityRoleInstantiates ProvenanceEntityRole = "instantiates"
	// An entity that is removed from accessibility, usually through the DELETE operator.
	ProvenanceEntityRoleRemoval ProvenanceEntityRole = "removal"
)
