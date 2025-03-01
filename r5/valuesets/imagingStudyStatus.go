// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/imagingstudy-status
*/type ImagingStudyStatus string

const (
	// The existence of the imaging study is registered, but there is nothing yet available.
	ImagingStudyStatusRegistered ImagingStudyStatus = "registered"
	// At least one instance has been associated with this imaging study.
	ImagingStudyStatusAvailable ImagingStudyStatus = "available"
	// The imaging study is unavailable because the imaging study was not started or not completed (also sometimes called "aborted").
	ImagingStudyStatusCancelled ImagingStudyStatus = "cancelled"
	// The imaging study has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	ImagingStudyStatusEnteredInError ImagingStudyStatus = "entered-in-error"
	// The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	ImagingStudyStatusUnknown ImagingStudyStatus = "unknown"
)
