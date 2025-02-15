// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/HealthcareService
type HealthcareService struct {
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
	// External identifiers for this item
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether this HealthcareService record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// Organization that provides this service
	ProvidedBy *custom.Reference[Organization] `bson:"providedBy,omitempty" json:"providedBy,omitempty"`
	// The service within which this service is offered
	OfferedIn []custom.Reference[HealthcareService] `bson:"offeredIn,omitempty" json:"offeredIn,omitempty"`
	// Broad category of service being performed or delivered
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Type of service that may be delivered or performed
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Specialties handled by the HealthcareService
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	// Location(s) where service may be provided
	Location []custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Description of service as presented to a consumer while searching
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Additional description and/or any specific issues not covered elsewhere
	Comment *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
	// Extra details about the service that can't be placed in the other fields
	ExtraDetails *custom.Markdown `bson:"extraDetails,omitempty" json:"extraDetails,omitempty"`
	// Facilitates quick identification of the service
	Photo *Attachment `bson:"photo,omitempty" json:"photo,omitempty"`
	// Official contact details for the HealthcareService
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Location(s) service is intended for/available to
	CoverageArea []custom.Reference[Location] `bson:"coverageArea,omitempty" json:"coverageArea,omitempty"`
	// Conditions under which service is available/offered
	ServiceProvisionCode []CodeableConcept `bson:"serviceProvisionCode,omitempty" json:"serviceProvisionCode,omitempty"`
	// Specific eligibility requirements required to use the service
	Eligibility []HealthcareServiceEligibility `bson:"eligibility,omitempty" json:"eligibility,omitempty"`
	// Programs that this service is applicable to
	Program []CodeableConcept `bson:"program,omitempty" json:"program,omitempty"`
	// Collection of characteristics (attributes)
	Characteristic []CodeableConcept `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	// The language that this service is offered in
	Communication []CodeableConcept `bson:"communication,omitempty" json:"communication,omitempty"`
	// Ways that the service accepts referrals
	ReferralMethod []CodeableConcept `bson:"referralMethod,omitempty" json:"referralMethod,omitempty"`
	// If an appointment is required for access to this service
	AppointmentRequired *bool `bson:"appointmentRequired,omitempty" json:"appointmentRequired,omitempty"`
	// Times the healthcare service is available (including exceptions)
	Availability []Availability `bson:"availability,omitempty" json:"availability,omitempty"`
	// Technical endpoints providing access to electronic services operated for the healthcare service
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type HealthcareServiceEligibility struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Coded value for the eligibility
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Describes the eligibility conditions for the service
	Comment *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
}

func (h HealthcareService) ResourceType() string {
	return "StructureDefinition"
}
