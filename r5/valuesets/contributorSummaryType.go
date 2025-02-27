// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/contributor-summary-type
*/type ContributorSummaryType string

const (
	// Display of the author list as a complete string.
	ContributorSummaryTypeAuthorString ContributorSummaryType = "author-string"
	// Display of the list of contributors as a complete string.
	ContributorSummaryTypeContributorshipList ContributorSummaryType = "contributorship-list"
	// Compiled summary of contributions.
	ContributorSummaryTypeContributorshipStatement ContributorSummaryType = "contributorship-statement"
	// Display of the list of acknowledged parties as a complete string.
	ContributorSummaryTypeAcknowledgmentList ContributorSummaryType = "acknowledgement-list"
	// Statement of acknowledgment of contributions beyond those compiled for formal contributorship statements.
	ContributorSummaryTypeAcknowledgmentStatement ContributorSummaryType = "acknowledgment-statement"
	// Statement of financial support for the creation of the cited artifact.
	ContributorSummaryTypeFundingStatement ContributorSummaryType = "funding-statement"
	// Statement of completing interests related to the creation of the cited artifact. Also called conflicts of interest or declaration of interests.
	ContributorSummaryTypeCompetingInterestsStatement ContributorSummaryType = "competing-interests-statement"
)
