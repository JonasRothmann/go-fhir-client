// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/observation-status
  - preliminary
  - final
  - entered-in-error

- http://hl7.org/fhir/detectedissue-status
  - mitigated
*/type DetectedIssueStatus string

const (
	// This is an initial or interim observation: data may be incomplete or unverified.
	DetectedIssueStatusPreliminary DetectedIssueStatus = "preliminary"
	// The observation is complete and there are no further actions needed. Additional information such "released", "signed", etc. would be represented using [Provenance](provenance.html) which provides not only the act but also the actors and dates and other related data. These act states would be associated with an observation status of `preliminary` until they are all completed and then a status of `final` would be applied.
	DetectedIssueStatusFinal DetectedIssueStatus = "final"
	// The observation has been withdrawn following previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "cancelled" rather than "entered-in-error".).
	DetectedIssueStatusEnteredInError DetectedIssueStatus = "entered-in-error"
	// Indicates the detected issue has been mitigated
	DetectedIssueStatusMitigated DetectedIssueStatus = "mitigated"
)
