// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/catalog
type ProfileForCatalog struct {
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
	// Canonical identifier for this Composition, represented as a URI (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Version-independent identifier for the Composition
	Identifier []Identifier `json:"identifier,omitempty"`
	// An explicitly assigned identifer of a variation of the content in the Composition
	Version *string `json:"version,omitempty"`
	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | deprecated | unknown
	Status custom.Code `json:"status"`
	// The type of document - a Catalog
	Type CodeableConcept `json:"type"`
	// The Content of the section
	Category CodeableConcept `json:"category"`
	// Who and/or what the composition is about
	Subject *custom.Reference[Resource] `json:"subject,omitempty"`
	// Context of the Composition
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Who and/or what authored the composition
	Author []custom.Reference[ProfileForCatalogAuthor] `json:"author"`
	// Name for this Composition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Human Readable name/title
	Title string `json:"title"`
	// For any additional notes
	Note []Annotation `json:"note,omitempty"`
	// Attests to accuracy of composition
	Attester []ProfileForCatalogAttester `json:"attester,omitempty"`
	// Organization which maintains the composition
	Custodian *custom.Reference[Organization] `json:"custodian,omitempty"`
	// Relationships to other compositions/documents
	RelatesTo []RelatedArtifact `json:"relatesTo,omitempty"`
	// The clinical service(s) being documented
	Event []ProfileForCatalogEvent `json:"event,omitempty"`
	// Composition is broken into sections
	Section []ProfileForCatalogSection `json:"section,omitempty"`
}

type ProfileForCatalogAttester struct {
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
	Party *custom.Reference[ProfileForCatalogAttesterParty] `json:"party,omitempty"`
}

type ProfileForCatalogEvent struct {
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

type ProfileForCatalogSection struct {
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
	Author []custom.Reference[ProfileForCatalogSectionAuthor] `json:"author,omitempty"`
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

type OtherProfileForCatalog ProfileForCatalog

type ProfileForCatalogAuthor interface {
	gofhirclient.Resource

	Is_ProfileForCatalogAuthor()
}

func (p Practitioner) Is_ProfileForCatalogAuthor()     {}
func (p PractitionerRole) Is_ProfileForCatalogAuthor() {}
func (d Device) Is_ProfileForCatalogAuthor()           {}
func (p Patient) Is_ProfileForCatalogAuthor()          {}
func (r RelatedPerson) Is_ProfileForCatalogAuthor()    {}
func (o Organization) Is_ProfileForCatalogAuthor()     {}

type ProfileForCatalogAttesterParty interface {
	gofhirclient.Resource

	Is_ProfileForCatalogAttesterParty()
}

func (p Patient) Is_ProfileForCatalogAttesterParty()          {}
func (r RelatedPerson) Is_ProfileForCatalogAttesterParty()    {}
func (p Practitioner) Is_ProfileForCatalogAttesterParty()     {}
func (p PractitionerRole) Is_ProfileForCatalogAttesterParty() {}
func (o Organization) Is_ProfileForCatalogAttesterParty()     {}

type ProfileForCatalogSectionAuthor interface {
	gofhirclient.Resource

	Is_ProfileForCatalogSectionAuthor()
}

func (p Practitioner) Is_ProfileForCatalogSectionAuthor()     {}
func (p PractitionerRole) Is_ProfileForCatalogSectionAuthor() {}
func (d Device) Is_ProfileForCatalogSectionAuthor()           {}
func (p Patient) Is_ProfileForCatalogSectionAuthor()          {}
func (r RelatedPerson) Is_ProfileForCatalogSectionAuthor()    {}
func (o Organization) Is_ProfileForCatalogSectionAuthor()     {}

func (p ProfileForCatalog) ResourceType() string {
	return "ProfileForCatalog"
}

func (p ProfileForCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherProfileForCatalog
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherProfileForCatalog: OtherProfileForCatalog(p), ResourceType: p.ResourceType()})
}

func UnmarshalProfileForCatalog(b []byte) (res ProfileForCatalog, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
