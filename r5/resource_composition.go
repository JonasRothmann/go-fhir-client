// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Composition
type Composition struct {
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
	Subject []custom.Reference[Resource] `json:"subject,omitempty"`
	// Context of the Composition
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Composition editing time
	Date custom.DateTime `json:"date"`
	// The context that the content is intended to support
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Who and/or what authored the composition
	Author []custom.Reference[CompositionAuthor] `json:"author"`
	// Name for this Composition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Human Readable name/title
	Title string `json:"title"`
	// For any additional notes
	Note []Annotation `json:"note,omitempty"`
	// Attests to accuracy of composition
	Attester []CompositionAttester `json:"attester,omitempty"`
	// Organization which maintains the composition
	Custodian *custom.Reference[Organization] `json:"custodian,omitempty"`
	// Relationships to other compositions/documents
	RelatesTo []RelatedArtifact `json:"relatesTo,omitempty"`
	// The clinical service(s) being documented
	Event []CompositionEvent `json:"event,omitempty"`
	// Composition is broken into sections
	Section []CompositionSection `json:"section,omitempty"`
}

type CompositionAttester struct {
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
	Party *custom.Reference[CompositionAttesterParty] `json:"party,omitempty"`
}

type CompositionEvent struct {
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

type CompositionSection struct {
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
	Author []custom.Reference[CompositionSectionAuthor] `json:"author,omitempty"`
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

type OtherComposition Composition

type CompositionAuthor interface {
	gofhirclient.Resource

	Is_CompositionAuthor()
}

func (p Practitioner) Is_CompositionAuthor()     {}
func (p PractitionerRole) Is_CompositionAuthor() {}
func (d Device) Is_CompositionAuthor()           {}
func (p Patient) Is_CompositionAuthor()          {}
func (r RelatedPerson) Is_CompositionAuthor()    {}
func (o Organization) Is_CompositionAuthor()     {}

type CompositionAttesterParty interface {
	gofhirclient.Resource

	Is_CompositionAttesterParty()
}

func (p Patient) Is_CompositionAttesterParty()          {}
func (r RelatedPerson) Is_CompositionAttesterParty()    {}
func (p Practitioner) Is_CompositionAttesterParty()     {}
func (p PractitionerRole) Is_CompositionAttesterParty() {}
func (o Organization) Is_CompositionAttesterParty()     {}

type CompositionSectionAuthor interface {
	gofhirclient.Resource

	Is_CompositionSectionAuthor()
}

func (p Practitioner) Is_CompositionSectionAuthor()     {}
func (p PractitionerRole) Is_CompositionSectionAuthor() {}
func (d Device) Is_CompositionSectionAuthor()           {}
func (p Patient) Is_CompositionSectionAuthor()          {}
func (r RelatedPerson) Is_CompositionSectionAuthor()    {}
func (o Organization) Is_CompositionSectionAuthor()     {}

func (c Composition) ResourceType() string {
	return "Composition"
}

func (c Composition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherComposition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherComposition: OtherComposition(c), ResourceType: c.ResourceType()})
}

func UnmarshalComposition(b []byte) (res Composition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
