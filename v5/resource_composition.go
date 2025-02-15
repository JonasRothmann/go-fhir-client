// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Composition
type Composition struct {
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
	// Canonical identifier for this Composition, represented as a URI (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Version-independent identifier for the Composition
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// An explicitly assigned identifer of a variation of the content in the Composition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | deprecated | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Kind of composition (LOINC if possible)
	Type CodeableConcept `bson:"type" json:"type"`
	// Categorization of Composition
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Who and/or what the composition is about
	Subject []custom.Reference[Resource] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Context of the Composition
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Composition editing time
	Date custom.DateTime `bson:"date" json:"date"`
	// The context that the content is intended to support
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Who and/or what authored the composition
	Author []custom.Reference[CompositionAuthor] `bson:"author" json:"author"`
	// Name for this Composition (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Human Readable name/title
	Title string `bson:"title" json:"title"`
	// For any additional notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Attests to accuracy of composition
	Attester []CompositionAttester `bson:"attester,omitempty" json:"attester,omitempty"`
	// Organization which maintains the composition
	Custodian *custom.Reference[Organization] `bson:"custodian,omitempty" json:"custodian,omitempty"`
	// Relationships to other compositions/documents
	RelatesTo []RelatedArtifact `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	// The clinical service(s) being documented
	Event []CompositionEvent `bson:"event,omitempty" json:"event,omitempty"`
	// Composition is broken into sections
	Section []CompositionSection `bson:"section,omitempty" json:"section,omitempty"`
}

type CompositionAttester struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// personal | professional | legal | official
	Mode CodeableConcept `bson:"mode" json:"mode"`
	// When the composition was attested
	Time *custom.DateTime `bson:"time,omitempty" json:"time,omitempty"`
	// Who attested the composition
	Party *custom.Reference[CompositionAttesterParty] `bson:"party,omitempty" json:"party,omitempty"`
}

type CompositionEvent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The period covered by the documentation
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// The event(s) being documented, as code(s), reference(s), or both
	Detail []custom.CodeableReference[Resource] `bson:"detail,omitempty" json:"detail,omitempty"`
}

type CompositionSection struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Label for section (e.g. for ToC)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Classification of section (recommended)
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Who and/or what authored the section
	Author []custom.Reference[CompositionSectionAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Who/what the section is about, when it is not about the subject of composition
	Focus *custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// Text summary of the section, for human interpretation
	Text *Narrative `bson:"text,omitempty" json:"text,omitempty"`
	// Order of section entries
	OrderedBy *CodeableConcept `bson:"orderedBy,omitempty" json:"orderedBy,omitempty"`
	// A reference to data that supports this section
	Entry []custom.Reference[Resource] `bson:"entry,omitempty" json:"entry,omitempty"`
	// Why the section is empty
	EmptyReason *CodeableConcept `bson:"emptyReason,omitempty" json:"emptyReason,omitempty"`
	// Nested Section
	Section []interface{} `bson:"section,omitempty" json:"section,omitempty"`
}

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
	return "StructureDefinition"
}
