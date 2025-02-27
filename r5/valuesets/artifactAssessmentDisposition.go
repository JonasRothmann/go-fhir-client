// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/artifactassessment-disposition
*/type ArtifactAssessmentDisposition string

const (
	// The comment is unresolved
	ArtifactAssessmentDispositionUnresolved ArtifactAssessmentDisposition = "unresolved"
	// The comment is not persuasive (rejected in full)
	ArtifactAssessmentDispositionNotPersuasive ArtifactAssessmentDisposition = "not-persuasive"
	// The comment is persuasive (accepted in full)
	ArtifactAssessmentDispositionPersuasive ArtifactAssessmentDisposition = "persuasive"
	// The comment is persuasive with modification (partially accepted)
	ArtifactAssessmentDispositionPersuasiveWithModification ArtifactAssessmentDisposition = "persuasive-with-modification"
	// The comment is not persuasive with modification (partially rejected)
	ArtifactAssessmentDispositionNotPersuasiveWithModification ArtifactAssessmentDisposition = "not-persuasive-with-modification"
)
