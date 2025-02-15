// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Organization
type Organization struct {
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
	// Identifies this organization  across multiple systems
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether the organization's record is still in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// Kind of organization
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Name used for the organization
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// A list of alternate names that the organization is known as, or was known as in the past
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty"`
	// Additional details about the Organization that could be displayed as further information to identify the Organization beyond its name
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Official contact details for the Organization
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// The organization of which this organization forms a part
	PartOf *custom.Reference[Organization] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// Technical endpoints providing access to services operated for the organization
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
	// Qualifications, certifications, accreditations, licenses, training, etc. pertaining to the provision of care
	Qualification []OrganizationQualification `bson:"qualification,omitempty" json:"qualification,omitempty"`
}

type OrganizationQualification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// An identifier for this qualification for the organization
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Coded representation of the qualification
	Code CodeableConcept `bson:"code" json:"code"`
	// Period during which the qualification is valid
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Organization that regulates and issues the qualification
	Issuer *custom.Reference[Organization] `bson:"issuer,omitempty" json:"issuer,omitempty"`
}

func (o Organization) ResourceType() string {
	return "StructureDefinition"
}
