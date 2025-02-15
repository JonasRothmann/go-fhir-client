// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Library
type Library struct {
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
	// Canonical identifier for this library, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Additional identifier for the library
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the library
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this library (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this library (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Subordinate title of the library
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// logic-library | model-definition | asset-collection | module-definition
	Type CodeableConcept `bson:"type" json:"type"`
	// Type of individual the library content is focused on
	Subject *CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the library
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for library (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this library is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Describes the clinical usage of the library
	Usage *custom.Markdown `bson:"usage,omitempty" json:"usage,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When the library was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// When the library was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// When the library is expected to be used
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Who authored the content
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Parameters defined by the library
	Parameter []ParameterDefinition `bson:"parameter,omitempty" json:"parameter,omitempty"`
	// What data is referenced by this library
	DataRequirement []DataRequirement `bson:"dataRequirement,omitempty" json:"dataRequirement,omitempty"`
	// Contents of the library, either embedded or referenced
	Content []Attachment `bson:"content,omitempty" json:"content,omitempty"`
}

func (l Library) ResourceType() string {
	return "StructureDefinition"
}
