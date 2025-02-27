// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/publication-status
*/type PublicationStatus string

const (
	// This resource is still under development and is not yet considered to be ready for normal use.
	PublicationStatusDraft PublicationStatus = "draft"
	// This resource is ready for normal use.
	PublicationStatusActive PublicationStatus = "active"
	// This resource has been withdrawn or superseded and should no longer be used.
	PublicationStatusRetired PublicationStatus = "retired"
	// The authoring system does not know which of the status values currently applies for this resource.  Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	PublicationStatusUnknown PublicationStatus = "unknown"
)
