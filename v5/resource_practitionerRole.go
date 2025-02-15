// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/PractitionerRole
type PractitionerRole struct {
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
	// Identifiers for a role/location
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Whether this practitioner role record is in active use
	Active *bool `bson:"active,omitempty" json:"active,omitempty"`
	// The period during which the practitioner is authorized to perform in these role(s)
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Practitioner that provides services for the organization
	Practitioner *custom.Reference[Practitioner] `bson:"practitioner,omitempty" json:"practitioner,omitempty"`
	// Organization where the roles are available
	Organization *custom.Reference[Organization] `bson:"organization,omitempty" json:"organization,omitempty"`
	// Roles which this practitioner may perform
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Specific specialty of the practitioner
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	// Location(s) where the practitioner provides care
	Location []custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Healthcare services provided for this role's Organization/Location(s)
	HealthcareService []custom.Reference[HealthcareService] `bson:"healthcareService,omitempty" json:"healthcareService,omitempty"`
	// Official contact details relating to this PractitionerRole
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Collection of characteristics (attributes)
	Characteristic []CodeableConcept `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	// A language the practitioner (in this role) can use in patient communication
	Communication []CodeableConcept `bson:"communication,omitempty" json:"communication,omitempty"`
	// Times the Practitioner is available at this location and/or healthcare service (including exceptions)
	Availability []Availability `bson:"availability,omitempty" json:"availability,omitempty"`
	// Endpoints for interacting with the practitioner in this role
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

func (p PractitionerRole) ResourceType() string {
	return "StructureDefinition"
}
