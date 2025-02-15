// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DocumentReference
type DocumentReference struct {
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
	// Business identifiers for the document
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// An explicitly assigned identifer of a variation of the content in the DocumentReference
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// Procedure that caused this media to be created
	BasedOn []custom.Reference[DocumentReferenceBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// current | superseded | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | deprecated | unknown
	DocStatus *custom.Code `bson:"docStatus,omitempty" json:"docStatus,omitempty"`
	// Imaging modality used
	Modality []CodeableConcept `bson:"modality,omitempty" json:"modality,omitempty"`
	// Kind of document (LOINC if possible)
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Categorization of document
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Who/what is the subject of the document
	Subject *custom.Reference[Resource] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Context of the document content
	Context []custom.Reference[DocumentReferenceContext] `bson:"context,omitempty" json:"context,omitempty"`
	// Main clinical acts documented
	Event []custom.CodeableReference `bson:"event,omitempty" json:"event,omitempty"`
	// Body part included
	BodySite []custom.CodeableReference[BodyStructure] `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Kind of facility where patient was seen
	FacilityType *CodeableConcept `bson:"facilityType,omitempty" json:"facilityType,omitempty"`
	// Additional details about where the content was created (e.g. clinical specialty)
	PracticeSetting *CodeableConcept `bson:"practiceSetting,omitempty" json:"practiceSetting,omitempty"`
	// Time of service that is being documented
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// When this document reference was created
	Date *custom.Instant `bson:"date,omitempty" json:"date,omitempty"`
	// Who and/or what authored the document
	Author []custom.Reference[DocumentReferenceAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Attests to accuracy of the document
	Attester []DocumentReferenceAttester `bson:"attester,omitempty" json:"attester,omitempty"`
	// Organization which maintains the document
	Custodian *custom.Reference[Organization] `bson:"custodian,omitempty" json:"custodian,omitempty"`
	// Relationships to other documents
	RelatesTo []DocumentReferenceRelatesTo `bson:"relatesTo,omitempty" json:"relatesTo,omitempty"`
	// Human-readable description
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Document security-tags
	SecurityLabel []CodeableConcept `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	// Document referenced
	Content []DocumentReferenceContent `bson:"content" json:"content"`
}

type DocumentReferenceAttester struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// personal | professional | legal | official
	Mode CodeableConcept `bson:"mode" json:"mode"`
	// When the document was attested
	Time *custom.DateTime `bson:"time,omitempty" json:"time,omitempty"`
	// Who attested the document
	Party *custom.Reference[DocumentReferenceAttesterParty] `bson:"party,omitempty" json:"party,omitempty"`
}

type DocumentReferenceRelatesTo struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The relationship type with another document
	Code CodeableConcept `bson:"code" json:"code"`
	// Target of the relationship
	Target custom.Reference[DocumentReference] `bson:"target" json:"target"`
}

type DocumentReferenceContent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Where to access the document
	Attachment Attachment `bson:"attachment" json:"attachment"`
	// Content profile rules for the document
	Profile []DocumentReferenceContentProfile `bson:"profile,omitempty" json:"profile,omitempty"`
}

type DocumentReferenceContentProfile struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code|uri|canonical
	Value Coding `bson:"value" json:"value"`
}

type DocumentReferenceBasedOn interface {
	gofhirclient.Resource

	Is_DocumentReferenceBasedOn()
}

func (a Appointment) Is_DocumentReferenceBasedOn()                {}
func (a AppointmentResponse) Is_DocumentReferenceBasedOn()        {}
func (c CarePlan) Is_DocumentReferenceBasedOn()                   {}
func (c Claim) Is_DocumentReferenceBasedOn()                      {}
func (c CommunicationRequest) Is_DocumentReferenceBasedOn()       {}
func (c Contract) Is_DocumentReferenceBasedOn()                   {}
func (c CoverageEligibilityRequest) Is_DocumentReferenceBasedOn() {}
func (d DeviceRequest) Is_DocumentReferenceBasedOn()              {}
func (e EnrollmentRequest) Is_DocumentReferenceBasedOn()          {}
func (i ImmunizationRecommendation) Is_DocumentReferenceBasedOn() {}
func (m MedicationRequest) Is_DocumentReferenceBasedOn()          {}
func (n NutritionOrder) Is_DocumentReferenceBasedOn()             {}
func (r RequestOrchestration) Is_DocumentReferenceBasedOn()       {}
func (s ServiceRequest) Is_DocumentReferenceBasedOn()             {}
func (s SupplyRequest) Is_DocumentReferenceBasedOn()              {}
func (v VisionPrescription) Is_DocumentReferenceBasedOn()         {}

type DocumentReferenceContext interface {
	gofhirclient.Resource

	Is_DocumentReferenceContext()
}

func (a Appointment) Is_DocumentReferenceContext()   {}
func (e Encounter) Is_DocumentReferenceContext()     {}
func (e EpisodeOfCare) Is_DocumentReferenceContext() {}

type DocumentReferenceAuthor interface {
	gofhirclient.Resource

	Is_DocumentReferenceAuthor()
}

func (p Practitioner) Is_DocumentReferenceAuthor()     {}
func (p PractitionerRole) Is_DocumentReferenceAuthor() {}
func (o Organization) Is_DocumentReferenceAuthor()     {}
func (d Device) Is_DocumentReferenceAuthor()           {}
func (p Patient) Is_DocumentReferenceAuthor()          {}
func (r RelatedPerson) Is_DocumentReferenceAuthor()    {}
func (c CareTeam) Is_DocumentReferenceAuthor()         {}

type DocumentReferenceAttesterParty interface {
	gofhirclient.Resource

	Is_DocumentReferenceAttesterParty()
}

func (p Patient) Is_DocumentReferenceAttesterParty()          {}
func (r RelatedPerson) Is_DocumentReferenceAttesterParty()    {}
func (p Practitioner) Is_DocumentReferenceAttesterParty()     {}
func (p PractitionerRole) Is_DocumentReferenceAttesterParty() {}
func (o Organization) Is_DocumentReferenceAttesterParty()     {}

func (d DocumentReference) ResourceType() string {
	return "StructureDefinition"
}
