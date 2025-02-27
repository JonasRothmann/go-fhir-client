// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EvidenceReport
type EvidenceReportSubjectCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Characteristic code
	Code CodeableConcept `json:"code"`
	// Characteristic value
	ValueReference *custom.Reference[Resource] `json:"valueReference,omitempty"`
	// Characteristic value
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Characteristic value
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Characteristic value
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Characteristic value
	ValueRange *Range `json:"valueRange,omitempty"`
	// Is used to express not the characteristic
	Exclude *bool `json:"exclude,omitempty"`
	// Timeframe for the characteristic
	Period *Period `json:"period,omitempty"`
}

type EvidenceReportRelatesTo struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// replaces | amends | appends | transforms | replacedWith | amendedWith | appendedWith | transformedWith
	Code custom.Code `json:"code"`
	// Target of the relationship
	Target EvidenceReportRelatesToTarget `json:"target"`
}

type EvidenceReportRelatesToTarget struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Target of the relationship URL
	Url *custom.URI `json:"url,omitempty"`
	// Target of the relationship Identifier
	Identifier *Identifier `json:"identifier,omitempty"`
	// Target of the relationship Display
	Display *custom.Markdown `json:"display,omitempty"`
	// Target of the relationship Resource reference
	Resource *custom.Reference[Resource] `json:"resource,omitempty"`
}

type EvidenceReportSection struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for section (e.g. for ToC)
	Title *string `json:"title,omitempty"`
	// Classification of section (recommended)
	Focus *CodeableConcept `json:"focus,omitempty"`
	// Classification of section by Resource
	FocusReference *custom.Reference[Resource] `json:"focusReference,omitempty"`
	// Who and/or what authored the section
	Author []custom.Reference[EvidenceReportSectionAuthor] `json:"author,omitempty"`
	// Text summary of the section, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// working | snapshot | changes
	Mode *custom.Code `json:"mode,omitempty"`
	// Order of section entries
	OrderedBy *CodeableConcept `json:"orderedBy,omitempty"`
	// Extensible classifiers as content
	EntryClassifier []CodeableConcept `json:"entryClassifier,omitempty"`
	// Reference to resources as content
	EntryReference []custom.Reference[Resource] `json:"entryReference,omitempty"`
	// Quantity as content
	EntryQuantity []Quantity `json:"entryQuantity,omitempty"`
	// Why the section is empty
	EmptyReason *CodeableConcept `json:"emptyReason,omitempty"`
	// Nested Section
	Section []interface{} `json:"section,omitempty"`
}

type EvidenceReport struct {
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
	// Canonical identifier for this EvidenceReport, represented as a globally unique URI
	Url *custom.URI `json:"url,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Unique identifier for the evidence report
	Identifier []Identifier `json:"identifier,omitempty"`
	// Identifiers for articles that may relate to more than one evidence report
	RelatedIdentifier []Identifier `json:"relatedIdentifier,omitempty"`
	// Citation for this report
	CiteAsReference *custom.Reference[Citation] `json:"citeAsReference,omitempty"`
	// Citation for this report
	CiteAsMarkdown *custom.Markdown `json:"citeAsMarkdown,omitempty"`
	// Kind of report
	Type *CodeableConcept `json:"type,omitempty"`
	// Used for footnotes and annotations
	Note []Annotation `json:"note,omitempty"`
	// Link, description or reference to artifact associated with the report
	RelatedArtifact []RelatedArtifact `json:"relatedArtifact,omitempty"`
	// Focus of the report
	Subject EvidenceReportSubject `json:"subject"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Who authored the content
	Author []ContactDetail `json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `json:"endorser,omitempty"`
	// Relationships to other compositions/documents
	RelatesTo []EvidenceReportRelatesTo `json:"relatesTo,omitempty"`
	// Composition is broken into sections
	Section []EvidenceReportSection `json:"section,omitempty"`
}

type EvidenceReportSubject struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Characteristic
	Characteristic []EvidenceReportSubjectCharacteristic `json:"characteristic,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `json:"note,omitempty"`
}

type OtherEvidenceReport EvidenceReport

type EvidenceReportSectionAuthor interface {
	gofhirclient.Resource

	Is_EvidenceReportSectionAuthor()
}

func (p Patient) Is_EvidenceReportSectionAuthor()          {}
func (p Practitioner) Is_EvidenceReportSectionAuthor()     {}
func (p PractitionerRole) Is_EvidenceReportSectionAuthor() {}
func (r RelatedPerson) Is_EvidenceReportSectionAuthor()    {}
func (d Device) Is_EvidenceReportSectionAuthor()           {}
func (g Group) Is_EvidenceReportSectionAuthor()            {}
func (o Organization) Is_EvidenceReportSectionAuthor()     {}

func (e EvidenceReport) ResourceType() string {
	return "EvidenceReport"
}

func (e EvidenceReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceReport
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherEvidenceReport: OtherEvidenceReport(e), ResourceType: e.ResourceType()})
}

func UnmarshalEvidenceReport(b []byte) (res EvidenceReport, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
