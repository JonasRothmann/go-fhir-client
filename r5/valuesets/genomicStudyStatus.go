// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/genomicstudy-status
*/type GenomicStudyStatus string

const (
	// The existence of the genomic study is registered, but there is nothing yet available.
	GenomicStudyStatusRegistered GenomicStudyStatus = "registered"
	// At least one instance has been associated with this genomic study.
	GenomicStudyStatusAvailable GenomicStudyStatus = "available"
	// The genomic study is unavailable because the genomic study was not started or not completed (also sometimes called "aborted").
	GenomicStudyStatusCancelled GenomicStudyStatus = "cancelled"
	// The genomic study has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	GenomicStudyStatusEnteredInError GenomicStudyStatus = "entered-in-error"
	// The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, it's just not known which one.
	GenomicStudyStatusUnknown GenomicStudyStatus = "unknown"
)
