// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Endpoint
type Endpoint struct {
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
	// Identifies this endpoint across multiple systems
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | suspended | error | off | entered-in-error | test
	Status custom.Code `bson:"status" json:"status"`
	// Protocol/Profile/Standard to be used with this endpoint connection
	ConnectionType []CodeableConcept `bson:"connectionType" json:"connectionType"`
	// A name that this endpoint can be identified by
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Additional details about the endpoint that could be displayed as further information to identify the description beyond its name
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// The type of environment(s) exposed at this endpoint
	EnvironmentType []CodeableConcept `bson:"environmentType,omitempty" json:"environmentType,omitempty"`
	// Organization that manages this endpoint (might not be the organization that exposes the endpoint)
	ManagingOrganization *custom.Reference[Organization] `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	// Contact details for source (e.g. troubleshooting)
	Contact []ContactPoint `bson:"contact,omitempty" json:"contact,omitempty"`
	// Interval the endpoint is expected to be operational
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Set of payloads that are provided by this endpoint
	Payload []EndpointPayload `bson:"payload,omitempty" json:"payload,omitempty"`
	// The technical base address for connecting to this endpoint
	Address custom.URL `bson:"address" json:"address"`
	// Usage depends on the channel type
	Header []string `bson:"header,omitempty" json:"header,omitempty"`
}

type EndpointPayload struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of content that may be used at this endpoint (e.g. XDS Discharge summaries)
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Mimetype to send. If not specified, the content could be anything (including no payload, if the connectionType defined this)
	MimeType []custom.Code `bson:"mimeType,omitempty" json:"mimeType,omitempty"`
}

func (e Endpoint) ResourceType() string {
	return "StructureDefinition"
}
