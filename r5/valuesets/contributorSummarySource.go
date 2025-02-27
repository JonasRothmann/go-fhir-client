// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/contributor-summary-source
*/type ContributorSummarySource string

const (
	// Data copied by machine from publisher data.
	ContributorSummarySourcePublisherProvided ContributorSummarySource = "publisher-data"
	// Data copied by human from article text.
	ContributorSummarySourceCopiedFromArticle ContributorSummarySource = "article-copy"
	// Data copied by machine from citation manager data.
	ContributorSummarySourceReportedByCitationManager ContributorSummarySource = "citation-manager"
	// Custom format (may be described in text note).
	ContributorSummarySourceCustom ContributorSummarySource = "custom"
)
