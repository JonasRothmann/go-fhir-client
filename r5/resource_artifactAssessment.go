// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ArtifactAssessment
type ArtifactAssessment struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Text summary of the resource, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Contained, inline Resources
	Contained []Resource `json:"contained,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Additional identifier for the artifact assessment
	Identifier []Identifier `json:"identifier,omitempty"`
	// A short title for the assessment for use in displaying and selecting
	Title *string `json:"title,omitempty"`
	// How to cite the comment or rating
	CiteAsReference *custom.Reference[Citation] `json:"citeAsReference,omitempty"`
	// How to cite the comment or rating
	CiteAsMarkdown *custom.Markdown `json:"citeAsMarkdown,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// When the artifact assessment was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the artifact assessment was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// The artifact assessed, commented upon or rated
	ArtifactReference *custom.Reference[Resource] `json:"artifactReference,omitempty"`
	// The artifact assessed, commented upon or rated
	ArtifactCanonical *custom.Canonical[any] `json:"artifactCanonical,omitempty"`
	// The artifact assessed, commented upon or rated
	ArtifactUri *custom.URI `json:"artifactUri,omitempty"`
	// Comment, classifier, or rating content
	Content []ArtifactAssessmentContent `json:"content,omitempty"`
	// submitted | triaged | waiting-for-input | resolved-no-change | resolved-change-required | deferred | duplicate | applied | published | entered-in-error
	WorkflowStatus *custom.Code `json:"workflowStatus,omitempty"`
	// unresolved | not-persuasive | persuasive | persuasive-with-modification | not-persuasive-with-modification
	Disposition *custom.Code `json:"disposition,omitempty"`
}

type ArtifactAssessmentContent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// comment | classifier | rating | container | response | change-request
	InformationType *custom.Code `json:"informationType,omitempty"`
	// Brief summary of the content
	Summary *custom.Markdown `json:"summary,omitempty"`
	// What type of content
	Type *CodeableConcept `json:"type,omitempty"`
	// Rating, classifier, or assessment
	Classifier []CodeableConcept `json:"classifier,omitempty"`
	// Quantitative rating
	Quantity *Quantity `json:"quantity,omitempty"`
	// Who authored the content
	Author *custom.Reference[ArtifactAssessmentContentAuthor] `json:"author,omitempty"`
	// What the comment is directed to
	Path []custom.URI `json:"path,omitempty"`
	// Additional information
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Acceptable to publicly share the resource content
	FreeToShare *bool `json:"freeToShare,omitempty"`
	// Contained content
	Component []interface{} `json:"component,omitempty"`
}

type OtherArtifactAssessment ArtifactAssessment

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
	return "ArtifactAssessment"
}

func (a ArtifactAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherArtifactAssessment
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherArtifactAssessment: OtherArtifactAssessment(a), ResourceType: a.ResourceType()})
}

func UnmarshalArtifactAssessment(b []byte) (res ArtifactAssessment, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
