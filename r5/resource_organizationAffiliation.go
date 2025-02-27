// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/OrganizationAffiliation
type OrganizationAffiliation struct {
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
	// Business identifiers that are specific to this role
	Identifier []Identifier `json:"identifier,omitempty"`
	// Whether this organization affiliation record is in active use
	Active *bool `json:"active,omitempty"`
	// The period during which the participatingOrganization is affiliated with the primary organization
	Period *Period `json:"period,omitempty"`
	// Organization where the role is available
	Organization *custom.Reference[Organization] `json:"organization,omitempty"`
	// Organization that provides/performs the role (e.g. providing services or is a member of)
	ParticipatingOrganization *custom.Reference[Organization] `json:"participatingOrganization,omitempty"`
	// The network in which the participatingOrganization provides the role's services (if defined) at the indicated locations (if defined)
	Network []custom.Reference[Organization] `json:"network,omitempty"`
	// Definition of the role the participatingOrganization plays
	Code []CodeableConcept `json:"code,omitempty"`
	// Specific specialty of the participatingOrganization in the context of the role
	Specialty []CodeableConcept `json:"specialty,omitempty"`
	// The location(s) at which the role occurs
	Location []custom.Reference[Location] `json:"location,omitempty"`
	// Healthcare services provided through the role
	HealthcareService []custom.Reference[HealthcareService] `json:"healthcareService,omitempty"`
	// Official contact details at the participatingOrganization relevant to this Affiliation
	Contact []ExtendedContactDetail `json:"contact,omitempty"`
	// Technical endpoints providing access to services operated for this role
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
}

type OtherOrganizationAffiliation OrganizationAffiliation

func (o OrganizationAffiliation) ResourceType() string {
	return "OrganizationAffiliation"
}

func (o OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganizationAffiliation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherOrganizationAffiliation: OtherOrganizationAffiliation(o), ResourceType: o.ResourceType()})
}

func UnmarshalOrganizationAffiliation(b []byte) (res OrganizationAffiliation, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
