// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Endpoint
type Endpoint struct {
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
	// Identifies this endpoint across multiple systems
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | suspended | error | off | entered-in-error | test
	Status custom.Code `json:"status"`
	// Protocol/Profile/Standard to be used with this endpoint connection
	ConnectionType []CodeableConcept `json:"connectionType"`
	// A name that this endpoint can be identified by
	Name *string `json:"name,omitempty"`
	// Additional details about the endpoint that could be displayed as further information to identify the description beyond its name
	Description *string `json:"description,omitempty"`
	// The type of environment(s) exposed at this endpoint
	EnvironmentType []CodeableConcept `json:"environmentType,omitempty"`
	// Organization that manages this endpoint (might not be the organization that exposes the endpoint)
	ManagingOrganization *custom.Reference[Organization] `json:"managingOrganization,omitempty"`
	// Contact details for source (e.g. troubleshooting)
	Contact []ContactPoint `json:"contact,omitempty"`
	// Interval the endpoint is expected to be operational
	Period *Period `json:"period,omitempty"`
	// Set of payloads that are provided by this endpoint
	Payload []EndpointPayload `json:"payload,omitempty"`
	// The technical base address for connecting to this endpoint
	Address custom.URL `json:"address"`
	// Usage depends on the channel type
	Header []string `json:"header,omitempty"`
}

type EndpointPayload struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of content that may be used at this endpoint (e.g. XDS Discharge summaries)
	Type []CodeableConcept `json:"type,omitempty"`
	// Mimetype to send. If not specified, the content could be anything (including no payload, if the connectionType defined this)
	MimeType []custom.Code `json:"mimeType,omitempty"`
}

type OtherEndpoint Endpoint

func (e Endpoint) ResourceType() string {
	return "Endpoint"
}

func (e Endpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEndpoint
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherEndpoint: OtherEndpoint(e), ResourceType: e.ResourceType()})
}

func UnmarshalEndpoint(b []byte) (res Endpoint, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
