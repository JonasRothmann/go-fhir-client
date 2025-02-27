// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/fhirpathlibrary
type FhirPathLibrary struct {
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
	// Extension
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Canonical identifier for this library, represented as a URI (globally unique)
	Url custom.URI `json:"url"`
	// Additional identifier for the library
	Identifier []Identifier `json:"identifier,omitempty"`
	// Business version of the library
	Version string `json:"version"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this library (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this library (human friendly)
	Title string `json:"title"`
	// Subordinate title of the library
	Subtitle *string `json:"subtitle,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// For testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// logic-library | model-definition | asset-collection | module-definition
	Type CodeableConcept `json:"type"`
	// Type of individual the library content is focused on
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	// Type of individual the library content is focused on
	SubjectReference *custom.Reference[Group] `json:"subjectReference,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the library
	Description custom.Markdown `json:"description"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for library (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this library is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Describes the clinical usage of the library
	Usage *custom.Markdown `json:"usage,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When the library was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// When the library was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// When the library is expected to be used
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// E.g. Education, Treatment, Assessment, etc
	Topic []CodeableConcept `json:"topic,omitempty"`
	// Who authored the content
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Additional documentation, citations, etc
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Parameters defined by the library
	Parameter []ParameterDefinition `json:"parameter,omitempty"`
	// What data is referenced by this library
	DataRequirement []DataRequirement `json:"dataRequirement,omitempty"`
	// Contents of the library, either embedded or referenced
	Content []Attachment `json:"content,omitempty"`
}

type FhirPathLibraryContent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Mime type of the content, with charset etc.
	ContentType *custom.Code `json:"contentType,omitempty"`
	// Human language of the content (BCP-47)
	Language *custom.Code `json:"language,omitempty"`
	// Data inline, base64ed
	Data *custom.Base64binary `json:"data,omitempty"`
	// Uri where the data can be found
	Url *custom.URL `json:"url,omitempty"`
	// Number of bytes of content (if url provided)
	Size *int64 `json:"size,omitempty"`
	// Hash of the data (sha-1, base64ed)
	Hash *custom.Base64binary `json:"hash,omitempty"`
	// Label to display in place of the data
	Title *string `json:"title,omitempty"`
	// Date attachment was first created
	Creation *custom.DateTime `json:"creation,omitempty"`
	// Height of the image in pixels (photo/video)
	Height *uint32 `json:"height,omitempty"`
	// Width of the image in pixels (photo/video)
	Width *uint32 `json:"width,omitempty"`
	// Number of frames if > 1 (photo)
	Frames *uint32 `json:"frames,omitempty"`
	// Length in seconds (audio / video)
	Duration *json.Number `json:"duration,omitempty"`
	// Number of printed pages
	Pages *uint32 `json:"pages,omitempty"`
}

type OtherFhirPathLibrary FhirPathLibrary

func (f FhirPathLibrary) ResourceType() string {
	return "FHIRPathLibrary"
}

func (f FhirPathLibrary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFhirPathLibrary
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherFhirPathLibrary: OtherFhirPathLibrary(f), ResourceType: f.ResourceType()})
}

func UnmarshalFhirPathLibrary(b []byte) (res FhirPathLibrary, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
