// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/Slot
type Slot struct {
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
	// External Ids for this item
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []CodeableConcept `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	// The type of appointments that can be booked into this slot (ideally this would be an identifiable service - which is at a location, rather than the location itself). If provided then this overrides the value provided on the Schedule resource
	ServiceType []custom.CodeableReference[HealthcareService] `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	// The style of appointment or patient that may be booked in the slot (not service type)
	AppointmentType []CodeableConcept `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	// The schedule resource that this slot defines an interval of status information
	Schedule custom.Reference[Schedule] `bson:"schedule" json:"schedule"`
	// busy | free | busy-unavailable | busy-tentative | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Date/Time that the slot is to begin
	Start custom.Instant `bson:"start" json:"start"`
	// Date/Time that the slot is to conclude
	End custom.Instant `bson:"end" json:"end"`
	// This slot has already been overbooked, appointments are unlikely to be accepted for this time
	Overbooked *bool `bson:"overbooked,omitempty" json:"overbooked,omitempty"`
	// Comments on the slot to describe any extended information. Such as custom constraints on the slot
	Comment *string `bson:"comment,omitempty" json:"comment,omitempty"`
}

func (s Slot) ResourceType() string {
	return "StructureDefinition"
}
