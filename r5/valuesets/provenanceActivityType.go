// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/iso-21089-lifecycle
  - amend
  - originate
  - merge
  - deidentify
  - receive
  - transform
  - verify
*/type ProvenanceActivityType string

const (
	// Occurs when an agent makes any change to record entry content currently residing in storage considered permanent (persistent).
	ProvenanceActivityTypeAmend ProvenanceActivityType = "amend"
	// Occurs when an agent causes the system to scrub record entry content to reduce the association between a set of identifying data and the data subject in a way that might or might not be reversible.
	ProvenanceActivityTypeDeidentify ProvenanceActivityType = "deidentify"
	// Occurs when an agent causes the system to combine or join content from two or more record entries, resulting in a single logical record entry.
	ProvenanceActivityTypeMerge ProvenanceActivityType = "merge"
	// Occurs when an agent causes the system to: a) initiate capture of potential record content, and b) incorporate that content into the storage considered a permanent part of the health record.
	ProvenanceActivityTypeOriginate ProvenanceActivityType = "originate"
	// Occurs when an agent causes the system to a) initiate capture of data content from elsewhere, and b) incorporate that content into the storage considered a permanent part of the health record.
	ProvenanceActivityTypeReceive ProvenanceActivityType = "receive"
	// Occurs when an agent causes the system to change the form, language or code system used to represent record entry content.
	ProvenanceActivityTypeTransform ProvenanceActivityType = "transform"
	// Occurs when an agent causes the system to confirm compliance of data or data objects with regulations, requirements, specifications, or other imposed conditions based on organizational policy.
	ProvenanceActivityTypeVerify ProvenanceActivityType = "verify"
)
