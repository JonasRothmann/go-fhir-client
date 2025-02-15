// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Location
type Location struct {
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
	// Unique code or number identifying the location to its users
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | suspended | inactive
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// The operational status of the location (typically only for a bed/room)
	OperationalStatus *Coding `bson:"operationalStatus,omitempty" json:"operationalStatus,omitempty"`
	// Name of the location as used by humans
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// A list of alternate names that the location is known as, or was known as, in the past
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty"`
	// Additional details about the location that could be displayed as further information to identify the location beyond its name
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// instance | kind
	Mode *custom.Code `bson:"mode,omitempty" json:"mode,omitempty"`
	// Type of function performed
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Official contact details for the location
	Contact []ExtendedContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Physical location
	Address *Address `bson:"address,omitempty" json:"address,omitempty"`
	// Physical form of the location
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty"`
	// The absolute geographic location
	Position *LocationPosition `bson:"position,omitempty" json:"position,omitempty"`
	// Organization responsible for provisioning and upkeep
	ManagingOrganization *custom.Reference[Organization] `bson:"managingOrganization,omitempty" json:"managingOrganization,omitempty"`
	// Another Location this one is physically a part of
	PartOf *custom.Reference[Location] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// Collection of characteristics (attributes)
	Characteristic []CodeableConcept `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	// What days/times during a week is this location usually open (including exceptions)
	HoursOfOperation []Availability `bson:"hoursOfOperation,omitempty" json:"hoursOfOperation,omitempty"`
	// Connection details of a virtual service (e.g. conference call)
	VirtualService []VirtualServiceDetail `bson:"virtualService,omitempty" json:"virtualService,omitempty"`
	// Technical endpoints providing access to services operated for the location
	Endpoint []custom.Reference[Endpoint] `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}

type LocationPosition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Longitude with WGS84 datum
	Longitude json.Number `bson:"longitude" json:"longitude"`
	// Latitude with WGS84 datum
	Latitude json.Number `bson:"latitude" json:"latitude"`
	// Altitude with WGS84 datum
	Altitude *json.Number `bson:"altitude,omitempty" json:"altitude,omitempty"`
}

func (l Location) ResourceType() string {
	return "StructureDefinition"
}
