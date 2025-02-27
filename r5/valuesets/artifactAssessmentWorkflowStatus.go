// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/artifactassessment-workflow-status
*/type ArtifactAssessmentWorkflowStatus string

const (
	// The comment has been submitted, but the responsible party has not yet been determined, or the responsible party has not yet determined the next steps to be taken.
	ArtifactAssessmentWorkflowStatusSubmitted ArtifactAssessmentWorkflowStatus = "submitted"
	// The comment has been triaged, meaning the responsible party has been determined and next steps have been identified to address the comment.
	ArtifactAssessmentWorkflowStatusTriaged ArtifactAssessmentWorkflowStatus = "triaged"
	// The comment is waiting for input from a specific party before next steps can be taken.
	ArtifactAssessmentWorkflowStatusWaitingForInput ArtifactAssessmentWorkflowStatus = "waiting-for-input"
	// The comment has been resolved and no changes resulted from the resolution
	ArtifactAssessmentWorkflowStatusResolvedNoChange ArtifactAssessmentWorkflowStatus = "resolved-no-change"
	// The comment has been resolved and changes are required to address the comment
	ArtifactAssessmentWorkflowStatusResolvedChangeRequired ArtifactAssessmentWorkflowStatus = "resolved-change-required"
	// The comment is acceptable, but resolution of the comment and application of any associated changes have been deferred
	ArtifactAssessmentWorkflowStatusDeferred ArtifactAssessmentWorkflowStatus = "deferred"
	// The comment is a duplicate of another comment already received
	ArtifactAssessmentWorkflowStatusDuplicate ArtifactAssessmentWorkflowStatus = "duplicate"
	// The comment is resolved and any necessary changes have been applied
	ArtifactAssessmentWorkflowStatusApplied ArtifactAssessmentWorkflowStatus = "applied"
	// The necessary changes to the artifact have been published in a new version of the artifact
	ArtifactAssessmentWorkflowStatusPublished ArtifactAssessmentWorkflowStatus = "published"
	// The assessment was entered in error
	ArtifactAssessmentWorkflowStatusEnteredInError ArtifactAssessmentWorkflowStatus = "entered-in-error"
)
