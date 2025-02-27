// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/cited-artifact-status-type
*/type CitedArtifactStatusType string

const (
	// The content was originally constructed or composed.
	CitedArtifactStatusTypeCreated CitedArtifactStatusType = "created"
	// The content was sent to the publisher for consideration of publication.
	CitedArtifactStatusTypeSubmitted CitedArtifactStatusType = "submitted"
	// The content that was not published has been removed from consideration for publishing by the submitter.
	CitedArtifactStatusTypeWithdrawn CitedArtifactStatusType = "withdrawn"
	// The content is awaiting assignment and delivery to reviewer(s).
	CitedArtifactStatusTypePreReview CitedArtifactStatusType = "pre-review"
	// The content is in a state of being reviewed.
	CitedArtifactStatusTypeUnderReview CitedArtifactStatusType = "under-review"
	// The content is in a state between the review(s) being completed and being published.
	CitedArtifactStatusTypePostReviewPrePublished CitedArtifactStatusType = "post-review-pre-published"
	// The content that was not published has been removed from consideration for publishing by a publisher or editor.
	CitedArtifactStatusTypeRejected CitedArtifactStatusType = "rejected"
	// The content is published but future changes to the published version are expected.
	CitedArtifactStatusTypePublishedEarlyForm CitedArtifactStatusType = "published-early-form"
	// The content is published and further changes to the content are not expected.
	CitedArtifactStatusTypePublishedFinalForm CitedArtifactStatusType = "published-final-form"
	// The content that was not published yet has been approved for publication by the publisher and/or editor.
	CitedArtifactStatusTypeAccepted CitedArtifactStatusType = "accepted"
	// The content is retired or considered no longer current but still available as part of the public record.
	CitedArtifactStatusTypeArchived CitedArtifactStatusType = "archived"
	// The content that was published is removed from publication and should no longer be considered part of the public record.
	CitedArtifactStatusTypeRetracted CitedArtifactStatusType = "retracted"
	// The content is considered unfinished or incomplete and not representative of the current state desired by the content creator.
	CitedArtifactStatusTypeDraft CitedArtifactStatusType = "draft"
	// The content is considered complete for its current state by the content creator.
	CitedArtifactStatusTypeActive CitedArtifactStatusType = "active"
	// The content has been approved for a state transition, with the focus of approval described in the text associated with this coding.
	CitedArtifactStatusTypeApproved CitedArtifactStatusType = "approved"
	// The status of the content is not recorded in the metadata.
	CitedArtifactStatusTypeUnknown CitedArtifactStatusType = "unknown"
)
