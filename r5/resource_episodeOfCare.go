// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/EpisodeOfCare
type EpisodeOfCareDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The medical condition that was addressed during the episode of care
	Condition []custom.CodeableReference[Condition] `json:"condition,omitempty"`
	// Role that this diagnosis has within the episode of care (e.g. admission, billing, discharge …)
	Use *CodeableConcept `json:"use,omitempty"`
}

type EpisodeOfCare struct {
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
	// Business Identifier(s) relevant for this EpisodeOfCare
	Identifier []Identifier `json:"identifier,omitempty"`
	// planned | waitlist | active | onhold | finished | cancelled | entered-in-error
	Status custom.Code `json:"status"`
	// Past list of status codes (the current status may be included to cover the start date of the status)
	StatusHistory []EpisodeOfCareStatusHistory `json:"statusHistory,omitempty"`
	// Type/class  - e.g. specialist referral, disease management
	Type []CodeableConcept `json:"type,omitempty"`
	// The list of medical reasons that are expected to be addressed during the episode of care
	Reason []EpisodeOfCareReason `json:"reason,omitempty"`
	// The list of medical conditions that were addressed during the episode of care
	Diagnosis []EpisodeOfCareDiagnosis `json:"diagnosis,omitempty"`
	// The patient who is the focus of this episode of care
	Patient custom.Reference[Patient] `json:"patient"`
	// Organization that assumes responsibility for care coordination
	ManagingOrganization *custom.Reference[Organization] `json:"managingOrganization,omitempty"`
	// Interval during responsibility is assumed
	Period *Period `json:"period,omitempty"`
	// Originating Referral Request(s)
	ReferralRequest []custom.Reference[ServiceRequest] `json:"referralRequest,omitempty"`
	// Care manager/care coordinator for the patient
	CareManager *custom.Reference[EpisodeOfCareCareManager] `json:"careManager,omitempty"`
	// Other practitioners facilitating this episode of care
	CareTeam []custom.Reference[CareTeam] `json:"careTeam,omitempty"`
	// The set of accounts that may be used for billing for this EpisodeOfCare
	Account []custom.Reference[Account] `json:"account,omitempty"`
}

type EpisodeOfCareStatusHistory struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// planned | waitlist | active | onhold | finished | cancelled | entered-in-error
	Status custom.Code `json:"status"`
	// Duration the EpisodeOfCare was in the specified status
	Period Period `json:"period"`
}

type EpisodeOfCareReason struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What the reason value should be used for/as
	Use *CodeableConcept `json:"use,omitempty"`
	// Medical reason to be addressed
	Value []custom.CodeableReference[EpisodeOfCareReasonValue] `json:"value,omitempty"`
}

type OtherEpisodeOfCare EpisodeOfCare

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
	return "EpisodeOfCare"
}

func (e EpisodeOfCare) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEpisodeOfCare
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherEpisodeOfCare: OtherEpisodeOfCare(e), ResourceType: e.ResourceType()})
}

func UnmarshalEpisodeOfCare(b []byte) (res EpisodeOfCare, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
