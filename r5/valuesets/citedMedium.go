// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/cited-medium
*/type CitedMedium string

const (
	// Online publication in a periodic release. Used to match NLM JournalIssue CitedMedium code for online version.
	CitedMediumInternet CitedMedium = "internet"
	// Print publication in a periodic release. Used to match NLM JournalIssue CitedMedium code for print version.
	CitedMediumPrint CitedMedium = "print"
	// Publication in a physical device for electronic data storage, organized in issues for periodic release.
	CitedMediumOfflineDigitalStorage CitedMedium = "offline-digital-storage"
	// Online publication without any periodic release. Used for article specific publication date which could be the same as or different from journal issue publication date.
	CitedMediumInternetWithoutIssue CitedMedium = "internet-without-issue"
	// Print publication without any periodic release.
	CitedMediumPrintWithoutIssue CitedMedium = "print-without-issue"
	// Publication in a physical device for electronic data storage, without any periodic release.
	CitedMediumOfflineDigitalStorageWithoutIssue CitedMedium = "offline-digital-storage-without-issue"
)
