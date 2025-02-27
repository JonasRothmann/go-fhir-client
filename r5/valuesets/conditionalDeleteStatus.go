// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/conditional-delete-status
*/type ConditionalDeleteStatus string

const (
	// No support for conditional deletes.
	ConditionalDeleteStatusNotSupported ConditionalDeleteStatus = "not-supported"
	// Conditional deletes are supported, but only single resources at a time.
	ConditionalDeleteStatusSingle ConditionalDeleteStatus = "single"
	// Conditional deletes are supported, and multiple resources can be deleted in a single interaction.
	ConditionalDeleteStatusMultiple ConditionalDeleteStatus = "multiple"
)
