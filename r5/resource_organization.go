// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
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
	// Identifies this organization  across multiple systems
	Identifier []Identifier `json:"identifier,omitempty"`
	// Whether the organization's record is still in active use
	Active *bool `json:"active,omitempty"`
	// Kind of organization
	Type []CodeableConcept `json:"type,omitempty"`
	// Name used for the organization
	Name *string `json:"name,omitempty"`
	// A list of alternate names that the organization is known as, or was known as in the past
	Alias []string `json:"alias,omitempty"`
	// Additional details about the Organization that could be displayed as further information to identify the Organization beyond its name
	Description *custom.Markdown `json:"description,omitempty"`
	// Official contact details for the Organization
	Contact []ExtendedContactDetail `json:"contact,omitempty"`
	// The organization of which this organization forms a part
	PartOf *custom.Reference[Organization] `json:"partOf,omitempty"`
	// Technical endpoints providing access to services operated for the organization
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
	// Qualifications, certifications, accreditations, licenses, training, etc. pertaining to the provision of care
	Qualification []OrganizationQualification `json:"qualification,omitempty"`
}

type OrganizationQualification struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An identifier for this qualification for the organization
	Identifier []Identifier `json:"identifier,omitempty"`
	// Coded representation of the qualification
	Code CodeableConcept `json:"code"`
	// Period during which the qualification is valid
	Period *Period `json:"period,omitempty"`
	// Organization that regulates and issues the qualification
	Issuer *custom.Reference[Organization] `json:"issuer,omitempty"`
}

type OtherOrganization Organization

func (o Organization) ResourceType() string {
	return "Organization"
}

func (o Organization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganization
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherOrganization: OtherOrganization(o), ResourceType: o.ResourceType()})
}

func UnmarshalOrganization(b []byte) (res Organization, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
