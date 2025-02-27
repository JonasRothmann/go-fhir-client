// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/clinicaldocument
type ClinicalDocument struct {
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
	// Canonical identifier for this Composition, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Version-independent identifier for the Composition
	Identifier []Identifier `json:"identifier,omitempty"`
	// An explicitly assigned identifer of a variation of the content in the Composition
	Version *string `json:"version,omitempty"`
	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | deprecated | unknown
	Status custom.Code `json:"status"`
	// Kind of composition (LOINC if possible)
	Type CodeableConcept `json:"type"`
	// Categorization of Composition
	Category []CodeableConcept `json:"category,omitempty"`
	// Who and/or what the composition is about
	Subject []custom.Reference[ClinicalDocumentSubject] `json:"subject,omitempty"`
	// Context of the Composition
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Composition editing time
	Date custom.DateTime `json:"date"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Who and/or what authored the composition
	Author []custom.Reference[ClinicalDocumentAuthor] `json:"author"`
	// Name for this Composition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Human Readable name/title
	Title string `json:"title"`
	// For any additional notes
	Note []Annotation `json:"note,omitempty"`
	// Attests to accuracy of composition
	Attester []ClinicalDocumentAttester `json:"attester,omitempty"`
	// Organization which maintains the composition
	Custodian *custom.Reference[Organization] `json:"custodian,omitempty"`
	// Relationships to other compositions/documents
	RelatesTo []RelatedArtifact `json:"relatesTo,omitempty"`
	// The clinical service(s) being documented
	Event []ClinicalDocumentEvent `json:"event,omitempty"`
	// Composition is broken into sections
	Section []ClinicalDocumentSection `json:"section,omitempty"`
}

type ClinicalDocumentAttester struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// personal | professional | legal | official
	Mode CodeableConcept `json:"mode"`
	// When the composition was attested
	Time *custom.DateTime `json:"time,omitempty"`
	// Who attested the composition
	Party *custom.Reference[ClinicalDocumentAttesterParty] `json:"party,omitempty"`
}

type ClinicalDocumentEvent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The period covered by the documentation
	Period *Period `json:"period,omitempty"`
	// The event(s) being documented, as code(s), reference(s), or both
	Detail []custom.CodeableReference[Resource] `json:"detail,omitempty"`
}

type ClinicalDocumentSection struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for section (e.g. for ToC)
	Title *string `json:"title,omitempty"`
	// Classification of section (recommended)
	Code *CodeableConcept `json:"code,omitempty"`
	// Who and/or what authored the section
	Author []custom.Reference[ClinicalDocumentSectionAuthor] `json:"author,omitempty"`
	// Who/what the section is about, when it is not about the subject of composition
	Focus *custom.Reference[Resource] `json:"focus,omitempty"`
	// Text summary of the section, for human interpretation
	Text *Narrative `json:"text,omitempty"`
	// Order of section entries
	OrderedBy *CodeableConcept `json:"orderedBy,omitempty"`
	// A reference to data that supports this section
	Entry []custom.Reference[Resource] `json:"entry,omitempty"`
	// Why the section is empty
	EmptyReason *CodeableConcept `json:"emptyReason,omitempty"`
	// Nested Section
	Section []interface{} `json:"section,omitempty"`
}

type OtherClinicalDocument ClinicalDocument

type ClinicalDocumentSubject interface {
	gofhirclient.Resource

	Is_ClinicalDocumentSubject()
}

func (p Patient) Is_ClinicalDocumentSubject()      {}
func (p Practitioner) Is_ClinicalDocumentSubject() {}
func (g Group) Is_ClinicalDocumentSubject()        {}
func (d Device) Is_ClinicalDocumentSubject()       {}
func (l Location) Is_ClinicalDocumentSubject()     {}

type ClinicalDocumentAuthor interface {
	gofhirclient.Resource

	Is_ClinicalDocumentAuthor()
}

func (p Practitioner) Is_ClinicalDocumentAuthor()     {}
func (p PractitionerRole) Is_ClinicalDocumentAuthor() {}
func (d Device) Is_ClinicalDocumentAuthor()           {}
func (p Patient) Is_ClinicalDocumentAuthor()          {}
func (r RelatedPerson) Is_ClinicalDocumentAuthor()    {}
func (o Organization) Is_ClinicalDocumentAuthor()     {}

type ClinicalDocumentAttesterParty interface {
	gofhirclient.Resource

	Is_ClinicalDocumentAttesterParty()
}

func (p Patient) Is_ClinicalDocumentAttesterParty()          {}
func (r RelatedPerson) Is_ClinicalDocumentAttesterParty()    {}
func (p Practitioner) Is_ClinicalDocumentAttesterParty()     {}
func (p PractitionerRole) Is_ClinicalDocumentAttesterParty() {}
func (o Organization) Is_ClinicalDocumentAttesterParty()     {}

type ClinicalDocumentSectionAuthor interface {
	gofhirclient.Resource

	Is_ClinicalDocumentSectionAuthor()
}

func (p Practitioner) Is_ClinicalDocumentSectionAuthor()     {}
func (p PractitionerRole) Is_ClinicalDocumentSectionAuthor() {}
func (d Device) Is_ClinicalDocumentSectionAuthor()           {}
func (p Patient) Is_ClinicalDocumentSectionAuthor()          {}
func (r RelatedPerson) Is_ClinicalDocumentSectionAuthor()    {}
func (o Organization) Is_ClinicalDocumentSectionAuthor()     {}

func (c ClinicalDocument) ResourceType() string {
	return "ClinicalDocument"
}

func (c ClinicalDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalDocument
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherClinicalDocument: OtherClinicalDocument(c), ResourceType: c.ResourceType()})
}

func UnmarshalClinicalDocument(b []byte) (res ClinicalDocument, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
