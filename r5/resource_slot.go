// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Slot
type Slot struct {
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
	// External Ids for this item
	Identifier []Identifier `json:"identifier,omitempty"`
	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []CodeableConcept `json:"serviceCategory,omitempty"`
	// The type of appointments that can be booked into this slot (ideally this would be an identifiable service - which is at a location, rather than the location itself). If provided then this overrides the value provided on the Schedule resource
	ServiceType []custom.CodeableReference[HealthcareService] `json:"serviceType,omitempty"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []CodeableConcept `json:"specialty,omitempty"`
	// The style of appointment or patient that may be booked in the slot (not service type)
	AppointmentType []CodeableConcept `json:"appointmentType,omitempty"`
	// The schedule resource that this slot defines an interval of status information
	Schedule custom.Reference[Schedule] `json:"schedule"`
	// busy | free | busy-unavailable | busy-tentative | entered-in-error
	Status custom.Code `json:"status"`
	// Date/Time that the slot is to begin
	Start custom.Instant `json:"start"`
	// Date/Time that the slot is to conclude
	End custom.Instant `json:"end"`
	// This slot has already been overbooked, appointments are unlikely to be accepted for this time
	Overbooked *bool `json:"overbooked,omitempty"`
	// Comments on the slot to describe any extended information. Such as custom constraints on the slot
	Comment *string `json:"comment,omitempty"`
}

type OtherSlot Slot

func (s Slot) ResourceType() string {
	return "Slot"
}

func (s Slot) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSlot
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSlot: OtherSlot(s), ResourceType: s.ResourceType()})
}

func UnmarshalSlot(b []byte) (res Slot, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
