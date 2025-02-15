// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EpisodeOfCare
type EpisodeOfCare struct {
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
	// Business Identifier(s) relevant for this EpisodeOfCare
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// planned | waitlist | active | onhold | finished | cancelled | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Past list of status codes (the current status may be included to cover the start date of the status)
	StatusHistory []EpisodeOfCareStatusHistory `bson:"statusHistory,omitempty" json:"statusHistory,omitempty"`
	// Type/class  - e.g. specialist referral, disease management
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The list of medical reasons that are expected to be addressed during the episode of care
	Reason []EpisodeOfCareReason `bson:"reason,omitempty" json:"reason,omitempty"`
	// The list of medical conditions that were addressed during the episode of care
	Diagnosis []EpisodeOfCareDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	// The patient who is the focus of this episode of care
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Organization that assumes responsibility for care coordination
	ManagingOrganization *custom.Reference[Organization] `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	// Interval during responsibility is assumed
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Originating Referral Request(s)
	ReferralRequest []custom.Reference[ServiceRequest] `bson:"referralRequest,omitempty" json:"referralRequest,omitempty"`
	// Care manager/care coordinator for the patient
	CareManager *custom.Reference[EpisodeOfCareCareManager] `bson:"careManager,omitempty" json:"careManager,omitempty"`
	// Other practitioners facilitating this episode of care
	CareTeam []custom.Reference[CareTeam] `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	// The set of accounts that may be used for billing for this EpisodeOfCare
	Account []custom.Reference[Account] `bson:"account,omitempty" json:"account,omitempty"`
}

type EpisodeOfCareStatusHistory struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// planned | waitlist | active | onhold | finished | cancelled | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Duration the EpisodeOfCare was in the specified status
	Period Period `bson:"period" json:"period"`
}

type EpisodeOfCareReason struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What the reason value should be used for/as
	Use *CodeableConcept `bson:"use,omitempty" json:"use,omitempty"`
	// Medical reason to be addressed
	Value []custom.CodeableReference[EpisodeOfCareReasonValue] `bson:"value,omitempty" json:"value,omitempty"`
}

type EpisodeOfCareDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The medical condition that was addressed during the episode of care
	Condition []custom.CodeableReference[Condition] `bson:"condition,omitempty" json:"condition,omitempty"`
	// Role that this diagnosis has within the episode of care (e.g. admission, billing, discharge …)
	Use *CodeableConcept `bson:"use,omitempty" json:"use,omitempty"`
}

type EpisodeOfCareReasonValue interface {
	gofhirclient.Resource

	Is_EpisodeOfCareReasonValue()
}

func (c Condition) Is_EpisodeOfCareReasonValue()         {}
func (p Procedure) Is_EpisodeOfCareReasonValue()         {}
func (o Observation) Is_EpisodeOfCareReasonValue()       {}
func (h HealthcareService) Is_EpisodeOfCareReasonValue() {}

type EpisodeOfCareCareManager interface {
	gofhirclient.Resource

	Is_EpisodeOfCareCareManager()
}

func (p Practitioner) Is_EpisodeOfCareCareManager()     {}
func (p PractitionerRole) Is_EpisodeOfCareCareManager() {}

func (e EpisodeOfCare) ResourceType() string {
	return "StructureDefinition"
}
