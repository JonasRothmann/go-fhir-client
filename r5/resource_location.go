// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Location
type Location struct {
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
	// Unique code or number identifying the location to its users
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | suspended | inactive
	Status *custom.Code `json:"status,omitempty"`
	// The operational status of the location (typically only for a bed/room)
	OperationalStatus *Coding `json:"operationalStatus,omitempty"`
	// Name of the location as used by humans
	Name *string `json:"name,omitempty"`
	// A list of alternate names that the location is known as, or was known as, in the past
	Alias []string `json:"alias,omitempty"`
	// Additional details about the location that could be displayed as further information to identify the location beyond its name
	Description *custom.Markdown `json:"description,omitempty"`
	// instance | kind
	Mode *custom.Code `json:"mode,omitempty"`
	// Type of function performed
	Type []CodeableConcept `json:"type,omitempty"`
	// Official contact details for the location
	Contact []ExtendedContactDetail `json:"contact,omitempty"`
	// Physical location
	Address *Address `json:"address,omitempty"`
	// Physical form of the location
	Form *CodeableConcept `json:"form,omitempty"`
	// The absolute geographic location
	Position *LocationPosition `json:"position,omitempty"`
	// Organization responsible for provisioning and upkeep
	ManagingOrganization *custom.Reference[Organization] `json:"managingOrganization,omitempty"`
	// Another Location this one is physically a part of
	PartOf *custom.Reference[Location] `json:"partOf,omitempty"`
	// Collection of characteristics (attributes)
	Characteristic []CodeableConcept `json:"characteristic,omitempty"`
	// What days/times during a week is this location usually open (including exceptions)
	HoursOfOperation []Availability `json:"hoursOfOperation,omitempty"`
	// Connection details of a virtual service (e.g. conference call)
	VirtualService []VirtualServiceDetail `json:"virtualService,omitempty"`
	// Technical endpoints providing access to services operated for the location
	Endpoint []custom.Reference[Endpoint] `json:"endpoint,omitempty"`
}

type LocationPosition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Longitude with WGS84 datum
	Longitude json.Number `json:"longitude"`
	// Latitude with WGS84 datum
	Latitude json.Number `json:"latitude"`
	// Altitude with WGS84 datum
	Altitude *json.Number `json:"altitude,omitempty"`
}

type OtherLocation Location

func (l Location) ResourceType() string {
	return "Location"
}

func (l Location) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherLocation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherLocation: OtherLocation(l), ResourceType: l.ResourceType()})
}

func UnmarshalLocation(b []byte) (res Location, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
