// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DocumentReference
type DocumentReferenceAttester struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// personal | professional | legal | official
	Mode CodeableConcept `json:"mode"`
	// When the document was attested
	Time *custom.DateTime `json:"time,omitempty"`
	// Who attested the document
	Party *custom.Reference[DocumentReferenceAttesterParty] `json:"party,omitempty"`
}

type DocumentReferenceRelatesTo struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The relationship type with another document
	Code CodeableConcept `json:"code"`
	// Target of the relationship
	Target custom.Reference[DocumentReference] `json:"target"`
}

type DocumentReferenceContent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Where to access the document
	Attachment Attachment `json:"attachment"`
	// Content profile rules for the document
	Profile []DocumentReferenceContentProfile `json:"profile,omitempty"`
}

type DocumentReferenceContentProfile struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Code|uri|canonical
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Code|uri|canonical
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Code|uri|canonical
	ValueCanonical *custom.Canonical[any] `json:"valueCanonical,omitempty"`
}

type DocumentReference struct {
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
	// Business identifiers for the document
	Identifier []Identifier `json:"identifier,omitempty"`
	// An explicitly assigned identifer of a variation of the content in the DocumentReference
	Version *string `json:"version,omitempty"`
	// Procedure that caused this media to be created
	BasedOn []custom.Reference[DocumentReferenceBasedOn] `json:"basedOn,omitempty"`
	// current | superseded | entered-in-error
	Status custom.Code `json:"status"`
	// registered | partial | preliminary | final | amended | corrected | appended | cancelled | entered-in-error | deprecated | unknown
	DocStatus *custom.Code `json:"docStatus,omitempty"`
	// Imaging modality used
	Modality []CodeableConcept `json:"modality,omitempty"`
	// Kind of document (LOINC if possible)
	Type *CodeableConcept `json:"type,omitempty"`
	// Categorization of document
	Category []CodeableConcept `json:"category,omitempty"`
	// Who/what is the subject of the document
	Subject *custom.Reference[Resource] `json:"subject,omitempty"`
	// Context of the document content
	Context []custom.Reference[DocumentReferenceContext] `json:"context,omitempty"`
	// Main clinical acts documented
	Event []custom.CodeableReference[gofhirclient.Resource] `json:"event,omitempty"`
	// Body part included
	BodySite []custom.CodeableReference[BodyStructure] `json:"bodySite,omitempty"`
	// Kind of facility where patient was seen
	FacilityType *CodeableConcept `json:"facilityType,omitempty"`
	// Additional details about where the content was created (e.g. clinical specialty)
	PracticeSetting *CodeableConcept `json:"practiceSetting,omitempty"`
	// Time of service that is being documented
	Period *Period `json:"period,omitempty"`
	// When this document reference was created
	Date *custom.Instant `json:"date,omitempty"`
	// Who and/or what authored the document
	Author []custom.Reference[DocumentReferenceAuthor] `json:"author,omitempty"`
	// Attests to accuracy of the document
	Attester []DocumentReferenceAttester `json:"attester,omitempty"`
	// Organization which maintains the document
	Custodian *custom.Reference[Organization] `json:"custodian,omitempty"`
	// Relationships to other documents
	RelatesTo []DocumentReferenceRelatesTo `json:"relatesTo,omitempty"`
	// Human-readable description
	Description *custom.Markdown `json:"description,omitempty"`
	// Document security-tags
	SecurityLabel []CodeableConcept `json:"securityLabel,omitempty"`
	// Document referenced
	Content []DocumentReferenceContent `json:"content"`
}

type OtherDocumentReference DocumentReference

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
	return "DocumentReference"
}

func (d DocumentReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDocumentReference
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDocumentReference: OtherDocumentReference(d), ResourceType: d.ResourceType()})
}

func UnmarshalDocumentReference(b []byte) (res DocumentReference, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
