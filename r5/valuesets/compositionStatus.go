// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/composition-status
*/type CompositionStatus string

const (
	// The existence of the composition is registered, but there is nothing yet available.
	CompositionStatusRegistered CompositionStatus = "registered"
	// This is a partial (e.g. initial, interim or preliminary) composition: data in the composition may be incomplete or unverified.
	CompositionStatusPartial CompositionStatus = "partial"
	// This version of the composition is complete and verified by an appropriate person and no further work is planned. Any subsequent updates would be on a new version of the composition.
	CompositionStatusFinal CompositionStatus = "final"
	// The composition content or the referenced resources have been modified (edited or added to) subsequent to being released as "final" and the composition is complete and verified by an authorized person.
	CompositionStatusAmended CompositionStatus = "amended"
	// The composition is unavailable because the measurement was not started or not completed (also sometimes called "aborted").
	CompositionStatusCancelled CompositionStatus = "cancelled"
	// The composition or document was originally created/issued in error, and this is an amendment that marks that the entire series should not be considered as valid.
	CompositionStatusEnteredInError CompositionStatus = "entered-in-error"
	// This composition has been withdrawn or superseded and should no longer be used.
	CompositionStatusDeprecated CompositionStatus = "deprecated"
	// The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	CompositionStatusUnknown CompositionStatus = "unknown"
)
