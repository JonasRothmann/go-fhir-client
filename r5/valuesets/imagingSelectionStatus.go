// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/imagingselection-status
*/type ImagingSelectionStatus string

const (
	// The selected resources are available..
	ImagingSelectionStatusAvailable ImagingSelectionStatus = "available"
	// The imaging selection has been withdrawn following a release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	ImagingSelectionStatusEnteredInError ImagingSelectionStatus = "entered-in-error"
	// The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	ImagingSelectionStatusUnknown ImagingSelectionStatus = "unknown"
)
