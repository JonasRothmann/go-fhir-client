// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/OrganizationAffiliation
type OrganizationAffiliation struct {
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
	// Business identifiers that are specific to this role
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether this organization affiliation record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// The period during which the participatingOrganization is affiliated with the primary organization
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Organization where the role is available
	Organization *custom.Reference[Organization] `bson:"organization,omitempty" json:"organization,omitempty"`
	// Organization that provides/performs the role (e.g. providing services or is a member of)
	ParticipatingOrganization *custom.Reference[Organization] `bson:"participatingOrganization,omitempty" json:"participatingOrganization,omitempty"`
	// The network in which the participatingOrganization provides the role's services (if defined) at the indicated locations (if defined)
	Network []custom.Reference[Organization] `bson:"network,omitempty" json:"network,omitempty"`
	// Definition of the role the participatingOrganization plays
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Specific specialty of the participatingOrganization in the context of the role
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	// The location(s) at which the role occurs
	Location []custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Healthcare services provided through the role
	HealthcareService []custom.Reference[HealthcareService] `bson:"healthcareService,omitempty" json:"healthcareService,omitempty"`
	// Official contact details at the participatingOrganization relevant to this Affiliation
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Technical endpoints providing access to services operated for this role
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

func (o OrganizationAffiliation) ResourceType() string {
	return "StructureDefinition"
}
