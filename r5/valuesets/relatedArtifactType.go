// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/related-artifact-type
*/type RelatedArtifactType string

const (
	// Additional documentation for the knowledge resource. This would include additional instructions on usage as well as additional information on clinical context or appropriateness.
	RelatedArtifactTypeDocumentation RelatedArtifactType = "documentation"
	// The target artifact is a summary of the justification for the knowledge resource including supporting evidence, relevant guidelines, or other clinically important information. This information is intended to provide a way to make the justification for the knowledge resource available to the consumer of interventions or results produced by the knowledge resource.
	RelatedArtifactTypeJustification RelatedArtifactType = "justification"
	// Bibliographic citation for papers, references, or other relevant material for the knowledge resource. This is intended to allow for citation of related material, but that was not necessarily specifically prepared in connection with this knowledge resource.
	RelatedArtifactTypeCitation RelatedArtifactType = "citation"
	// The previous version of the knowledge artifact, used to establish an ordering of versions of an artifact, independent of the status of each version.
	RelatedArtifactTypePredecessor RelatedArtifactType = "predecessor"
	// The subsequent version of the knowledge artfact, used to establish an ordering of versions of an artifact, independent of the status of each version.
	RelatedArtifactTypeSuccessor RelatedArtifactType = "successor"
	// This artifact is derived from the target artifact. This is intended to capture the relationship in which a particular knowledge resource is based on the content of another artifact, but is modified to capture either a different set of overall requirements, or a more specific set of requirements such as those involved in a particular institution or clinical setting. The artifact may be derived from one or more target artifacts.
	RelatedArtifactTypeDerivedFrom RelatedArtifactType = "derived-from"
	// This artifact depends on the target artifact. There is a requirement to use the target artifact in the creation or interpretation of this artifact.
	RelatedArtifactTypeDependsOn RelatedArtifactType = "depends-on"
	// This artifact is composed of the target artifact. This artifact is constructed with the target artifact as a component. The target artifact is a part of this artifact. (A dataset is composed of data.).
	RelatedArtifactTypeComposedOf RelatedArtifactType = "composed-of"
	// This artifact is a part of the target artifact. The target artifact is composed of this artifact (and possibly other artifacts).
	RelatedArtifactTypePartOf RelatedArtifactType = "part-of"
	// This artifact amends or changes the target artifact. This artifact adds additional information that is functionally expected to replace information in the target artifact. This artifact replaces a part but not all of the target artifact.
	RelatedArtifactTypeAmends RelatedArtifactType = "amends"
	// This artifact is amended with or changed by the target artifact. There is information in this artifact that should be functionally replaced with information in the target artifact.
	RelatedArtifactTypeAmendedWith RelatedArtifactType = "amended-with"
	// This artifact adds additional information to the target artifact. The additional information does not replace or change information in the target artifact.
	RelatedArtifactTypeAppends RelatedArtifactType = "appends"
	// This artifact has additional information in the target artifact.
	RelatedArtifactTypeAppendedWith RelatedArtifactType = "appended-with"
	// This artifact cites the target artifact. This may be a bibliographic citation for papers, references, or other relevant material for the knowledge resource. This is intended to allow for citation of related material, but that was not necessarily specifically prepared in connection with this knowledge resource.
	RelatedArtifactTypeCites RelatedArtifactType = "cites"
	// This artifact is cited by the target artifact.
	RelatedArtifactTypeCitedBy RelatedArtifactType = "cited-by"
	// This artifact contains comments about the target artifact.
	RelatedArtifactTypeIsCommentOn RelatedArtifactType = "comments-on"
	// This artifact has comments about it in the target artifact.  The type of comments may be expressed in the targetClassifier element such as reply, review, editorial, feedback, solicited, unsolicited, structured, unstructured.
	RelatedArtifactTypeHasCommentIn RelatedArtifactType = "comment-in"
	// This artifact is a container in which the target artifact is contained. A container is a data structure whose instances are collections of other objects. (A database contains the dataset.).
	RelatedArtifactTypeContains RelatedArtifactType = "contains"
	// This artifact is contained in the target artifact. The target artifact is a data structure whose instances are collections of other objects.
	RelatedArtifactTypeContainedIn RelatedArtifactType = "contained-in"
	// This artifact identifies errors and replacement content for the target artifact.
	RelatedArtifactTypeCorrects RelatedArtifactType = "corrects"
	// This artifact has corrections to it in the target artifact. The target artifact identifies errors and replacement content for this artifact.
	RelatedArtifactTypeCorrectionIn RelatedArtifactType = "correction-in"
	// This artifact replaces or supersedes the target artifact. The target artifact may be considered deprecated.
	RelatedArtifactTypeReplaces RelatedArtifactType = "replaces"
	// This artifact is replaced with or superseded by the target artifact. This artifact may be considered deprecated.
	RelatedArtifactTypeReplacedWith RelatedArtifactType = "replaced-with"
	// This artifact retracts the target artifact. The content that was published in the target artifact should be considered removed from publication and should no longer be considered part of the public record.
	RelatedArtifactTypeRetracts RelatedArtifactType = "retracts"
	// This artifact is retracted by the target artifact. The content that was published in this artifact should be considered removed from publication and should no longer be considered part of the public record.
	RelatedArtifactTypeRetractedBy RelatedArtifactType = "retracted-by"
	// This artifact is a signature of the target artifact.
	RelatedArtifactTypeSigns RelatedArtifactType = "signs"
	// This artifact has characteristics in common with the target artifact. This relationship may be used in systems to “deduplicate” knowledge artifacts from different sources, or in systems to show “similar items”.
	RelatedArtifactTypeSimilarTo RelatedArtifactType = "similar-to"
	// This artifact provides additional support for the target artifact. The type of support  is not documentation as it does not describe, explain, or instruct regarding the target artifact.
	RelatedArtifactTypeSupports RelatedArtifactType = "supports"
	// The target artifact contains additional information related to the knowledge artifact but is not documentation as the additional information does not describe, explain, or instruct regarding the knowledge artifact content or application. This could include an associated dataset.
	RelatedArtifactTypeSupportedWith RelatedArtifactType = "supported-with"
	// This artifact was generated by transforming the target artifact (e.g., format or language conversion). This is intended to capture the relationship in which a particular knowledge resource is based on the content of another artifact, but changes are only apparent in form and there is only one target artifact with the “transforms” relationship type.
	RelatedArtifactTypeTransforms RelatedArtifactType = "transforms"
	// This artifact was transformed into the target artifact (e.g., by format or language conversion).
	RelatedArtifactTypeTransformedInto RelatedArtifactType = "transformed-into"
	// This artifact was generated by transforming a related artifact (e.g., format or language conversion), noted separately with the “transforms” relationship type. This transformation used the target artifact to inform the transformation. The target artifact may be a conversion script or translation guide.
	RelatedArtifactTypeTransformedWith RelatedArtifactType = "transformed-with"
	// This artifact provides additional documentation for the target artifact. This could include additional instructions on usage as well as additional information on clinical context or appropriateness.
	RelatedArtifactTypeDocuments RelatedArtifactType = "documents"
	// The target artifact is a precise description of a concept in this artifact. This may be used when the RelatedArtifact datatype is used in elements contained in this artifact.
	RelatedArtifactTypeSpecificationOf RelatedArtifactType = "specification-of"
	// This artifact was created with the target artifact. The target artifact is a tool or support material used in the creation of the artifact, and not content that the artifact was derived from.
	RelatedArtifactTypeCreatedWith RelatedArtifactType = "created-with"
	// The related artifact is the citation for this artifact.
	RelatedArtifactTypeCiteAs RelatedArtifactType = "cite-as"
)
