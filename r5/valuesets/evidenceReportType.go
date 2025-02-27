// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/evidence-report-type
*/type EvidenceReportType string

const (
	// The report is primarily a listing of classifiers about the report subject.
	EvidenceReportTypeClassification EvidenceReportType = "classification"
	// The report is a composition of results generated in response to a search query.
	EvidenceReportTypeSearchResults EvidenceReportType = "search-results"
	// The report is a composition containing one or more FHIR resources in the content.
	EvidenceReportTypeResourceCompilation EvidenceReportType = "resources-compiled"
	// The report is a structured representation of text.
	EvidenceReportTypeStructuredText EvidenceReportType = "text-structured"
)
