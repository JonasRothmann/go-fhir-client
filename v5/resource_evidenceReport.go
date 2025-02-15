// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EvidenceReport
type EvidenceReportRelatesToTarget struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Target of the relationship URL
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Target of the relationship Identifier
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Target of the relationship Display
	Display *custom.Markdown `bson:"display,omitempty" json:"display,omitempty"`
	// Target of the relationship Resource reference
	Resource *custom.Reference[Resource] `bson:"resource,omitempty" json:"resource,omitempty"`
}

type EvidenceReportSection struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for section (e.g. for ToC)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Classification of section (recommended)
	Focus *CodeableConcept `bson:"focus,omitempty" json:"focus,omitempty"`
	// Classification of section by Resource
	FocusReference *custom.Reference[Resource] `bson:"focusReference,omitempty" json:"focusReference,omitempty"`
	// Who and/or what authored the section
	Author []custom.Reference[EvidenceReportSectionAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Text summary of the section, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// working | snapshot | changes
	Mode *custom.Code `bson:"mode,omitempty" json:"mode,omitempty"`
	// Order of section entries
	OrderedBy *CodeableConcept `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	// Extensible classifiers as content
	EntryClassifier []CodeableConcept `bson:"entryClassifier,omitempty" json:"entryClassifier,omitempty"`
	// Reference to resources as content
	EntryReference []custom.Reference[Resource] `bson:"entryReference,omitempty" json:"entryReference,omitempty"`
	// Quantity as content
	EntryQuantity []Quantity `bson:"entryQuantity,omitempty" json:"entryQuantity,omitempty"`
	// Why the section is empty
	EmptyReason *CodeableConcept `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	// Nested Section
	Section []interface{} `bson:"section,omitempty" json:"section,omitempty"`
}

type EvidenceReport struct {
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
	// Canonical identifier for this EvidenceReport, represented as a globally unique URI
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Unique identifier for the evidence report
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Identifiers for articles that may relate to more than one evidence report
	RelatedIdentifier []Identifier `bson:"relatedIdentifier,omitempty" json:"relatedIdentifier,omitempty"`
	// Citation for this report
	CiteAs *custom.Reference[Citation] `bson:"citeAs,omitempty" json:"citeAs,omitempty"`
	// Kind of report
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Used for footnotes and annotations
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Link, description or reference to artifact associated with the report
	RelatedArtifact []RelatedArtifact `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	// Focus of the report
	Subject EvidenceReportSubject `bson:"subject" json:"subject"`
	// Name of the publisher/steward (organization or individual)
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Who authored the content
	Author []ContactDetail `bson:"author,omitempty" json:"author,omitempty"`
	// Who edited the content
	Editor []ContactDetail `bson:"editor,omitempty" json:"editor,omitempty"`
	// Who reviewed the content
	Reviewer []ContactDetail `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	// Who endorsed the content
	Endorser []ContactDetail `bson:"endorser,omitempty" json:"endorser,omitempty"`
	// Relationships to other compositions/documents
	RelatesTo []EvidenceReportRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	// Composition is broken into sections
	Section []EvidenceReportSection `bson:"section,omitempty" json:"section,omitempty"`
}

type EvidenceReportSubject struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Characteristic
	Characteristic []EvidenceReportSubjectCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	// Footnotes and/or explanatory notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type EvidenceReportSubjectCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Characteristic code
	Code CodeableConcept `bson:"code" json:"code"`
	// Characteristic value
	Value custom.Reference[Resource] `bson:"value" json:"value"`
	// Is used to express not the characteristic
	Exclude *bool `bson:"exclude,omitempty" json:"exclude,omitempty"`
	// Timeframe for the characteristic
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type EvidenceReportRelatesTo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// replaces | amends | appends | transforms | replacedWith | amendedWith | appendedWith | transformedWith
	Code custom.Code `bson:"code" json:"code"`
	// Target of the relationship
	Target EvidenceReportRelatesToTarget `bson:"target" json:"target"`
}

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
	return "StructureDefinition"
}
