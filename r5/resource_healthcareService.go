// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/HealthcareService
type HealthcareServiceEligibility struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Coded value for the eligibility
	Code *CodeableConcept `json:"code,omitempty"`
	// Describes the eligibility conditions for the service
	Comment *custom.Markdown `json:"comment,omitempty"`
}

type HealthcareService struct {
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
	// External identifiers for this item
	Identifier []Identifier `json:"identifier,omitempty"`
	// Whether this HealthcareService record is in active use
	Active *bool `json:"active,omitempty"`
	// Organization that provides this service
	ProvidedBy *custom.Reference[Organization] `json:"providedBy,omitempty"`
	// The service within which this service is offered
	OfferedIn []custom.Reference[HealthcareService] `json:"offeredIn,omitempty"`
	// Broad category of service being performed or delivered
	Category []CodeableConcept `json:"category,omitempty"`
	// Type of service that may be delivered or performed
	Type []CodeableConcept `json:"type,omitempty"`
	// Specialties handled by the HealthcareService
	Specialty []CodeableConcept `json:"specialty,omitempty"`
	// Location(s) where service may be provided
	Location []custom.Reference[Location] `json:"location,omitempty"`
	// Description of service as presented to a consumer while searching
	Name *string `json:"name,omitempty"`
	// Additional description and/or any specific issues not covered elsewhere
	Comment *custom.Markdown `json:"comment,omitempty"`
	// Extra details about the service that can't be placed in the other fields
	ExtraDetails *custom.Markdown `json:"extraDetails,omitempty"`
	// Facilitates quick identification of the service
	Photo *Attachment `json:"photo,omitempty"`
	// Official contact details for the HealthcareService
	Contact []ExtendedContactDetail `json:"contact,omitempty"`
	// Location(s) service is intended for/available to
	CoverageArea []custom.Reference[Location] `json:"coverageArea,omitempty"`
	// Conditions under which service is available/offered
	ServiceProvisionCode []CodeableConcept `json:"serviceProvisionCode,omitempty"`
	// Specific eligibility requirements required to use the service
	Eligibility []HealthcareServiceEligibility `json:"eligibility,omitempty"`
	// Programs that this service is applicable to
	Program []CodeableConcept `json:"program,omitempty"`
	// Collection of characteristics (attributes)
	Characteristic []CodeableConcept `json:"characteristic,omitempty"`
	// The language that this service is offered in
	Communication []CodeableConcept `json:"communication,omitempty"`
	// Ways that the service accepts referrals
	ReferralMethod []CodeableConcept `json:"referralMethod,omitempty"`
	// If an appointment is required for access to this service
	AppointmentRequired *bool `json:"appointmentRequired,omitempty"`
	// Times the healthcare service is available (including exceptions)
	Availability []Availability `json:"availability,omitempty"`
	// Technical endpoints providing access to electronic services operated for the healthcare service
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
}

type OtherHealthcareService HealthcareService

func (h HealthcareService) ResourceType() string {
	return "HealthcareService"
}

func (h HealthcareService) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherHealthcareService
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherHealthcareService: OtherHealthcareService(h), ResourceType: h.ResourceType()})
}

func UnmarshalHealthcareService(b []byte) (res HealthcareService, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
