// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/observation-status
*/type ObservationStatus string

const (
	// The existence of the observation is registered, but there is no result yet available.
	ObservationStatusRegistered ObservationStatus = "registered"
	// This is an initial or interim observation: data may be incomplete or unverified.
	ObservationStatusPreliminary ObservationStatus = "preliminary"
	// The observation is complete and there are no further actions needed. Additional information such "released", "signed", etc. would be represented using [Provenance](provenance.html) which provides not only the act but also the actors and dates and other related data. These act states would be associated with an observation status of `preliminary` until they are all completed and then a status of `final` would be applied.
	ObservationStatusFinal ObservationStatus = "final"
	// Subsequent to being Final, the observation has been modified subsequent.  This includes updates/new information and corrections.
	ObservationStatusAmended ObservationStatus = "amended"
	// The observation is unavailable because the measurement was not started or not completed (also sometimes called "aborted").
	ObservationStatusCancelled ObservationStatus = "cancelled"
	// The observation has been withdrawn following previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	ObservationStatusEnteredInError ObservationStatus = "entered-in-error"
	// The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for "other" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which.
	ObservationStatusUnknown ObservationStatus = "unknown"
)
