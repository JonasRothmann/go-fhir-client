// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ArtifactAssessment
type ArtifactAssessment struct {
	// Logical id of this artifact
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `bson:"meta,omitempty" json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `bson:"contained,omitempty" json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Additional identifier for the artifact assessment
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// A short title for the assessment for use in displaying and selecting
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// How to cite the comment or rating
	CiteAs *custom.Reference[Citation] `bson:"citeAs,omitempty" json:"citeAs,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// When the artifact assessment was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the artifact assessment was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// The artifact assessed, commented upon or rated
	Artifact custom.Reference[Resource] `bson:"artifact" json:"artifact"`
	// Comment, classifier, or rating content
	Content []ArtifactAssessmentContent `bson:"content,omitempty" json:"content,omitempty"`
	// submitted | triaged | waiting-for-input | resolved-no-change | resolved-change-required | deferred | duplicate | applied | published | entered-in-error
	WorkflowStatus *custom.Code `bson:"workflowStatus,omitempty" json:"workflowStatus,omitempty"`
	// unresolved | not-persuasive | persuasive | persuasive-with-modification | not-persuasive-with-modification
	Disposition *custom.Code `bson:"disposition,omitempty" json:"disposition,omitempty"`
}

type ArtifactAssessmentContent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// comment | classifier | rating | container | response | change-request
	InformationType *custom.Code `bson:"informationType,omitempty" json:"informationType,omitempty"`
	// Brief summary of the content
	Summary *custom.Markdown `bson:"summary,omitempty" json:"summary,omitempty"`
	// What type of content
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Rating, classifier, or assessment
	Classifier []CodeableConcept `bson:"classifier,omitempty" json:"classifier,omitempty"`
	// Quantitative rating
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Who authored the content
	Author *custom.Reference[ArtifactAssessmentContentAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// What the comment is directed to
	Path []custom.URI `bson:"path,omitempty" json:"path,omitempty"`
	// Additional information
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Acceptable to publicly share the resource content
	FreeToShare *bool `bson:"freeToShare,omitempty" json:"freeToShare,omitempty"`
	// Contained content
	Component []interface{} `bson:"component,omitempty" json:"component,omitempty"`
}

type ArtifactAssessmentContentAuthor interface {
	gofhirclient.Resource

	Is_ArtifactAssessmentContentAuthor()
}

func (p Patient) Is_ArtifactAssessmentContentAuthor()          {}
func (p Practitioner) Is_ArtifactAssessmentContentAuthor()     {}
func (p PractitionerRole) Is_ArtifactAssessmentContentAuthor() {}
func (o Organization) Is_ArtifactAssessmentContentAuthor()     {}
func (d Device) Is_ArtifactAssessmentContentAuthor()           {}

func (a ArtifactAssessment) ResourceType() string {
	return "StructureDefinition"
}
