// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
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
	// Appointment this response relates to
	Appointment custom.Reference[Appointment] `bson:"appointment" json:"appointment"`
	// Indicator for a counter proposal
	ProposedNewTime *bool `bson:"proposedNewTime,omitempty" json:"proposedNewTime,omitempty"`
	// Time from appointment, or requested new start time
	Start *custom.Instant `bson:"start,omitempty" json:"start,omitempty"`
	// Time from appointment, or requested new end time
	End *custom.Instant `bson:"end,omitempty" json:"end,omitempty"`
	// Role of participant in the appointment
	ParticipantType []CodeableConcept `bson:"participantType,omitempty" json:"participantType,omitempty"`
	// Person(s), Location, HealthcareService, or Device
	Actor *custom.Reference[AppointmentResponseActor] `bson:"actor,omitempty" json:"actor,omitempty"`
	// accepted | declined | tentative | needs-action | entered-in-error
	ParticipantStatus custom.Code `bson:"participantStatus" json:"participantStatus"`
	// Additional comments
	Comment *custom.Markdown `bson:"comment,omitempty" json:"comment,omitempty"`
	// This response is for all occurrences in a recurring request
	Recurring *bool `bson:"recurring,omitempty" json:"recurring,omitempty"`
	// Original date within a recurring request
	OccurrenceDate *custom.Date `bson:"occurrenceDate,omitempty" json:"occurrenceDate,omitempty"`
	// The recurrence ID of the specific recurring request
	RecurrenceId *uint32 `bson:"recurrenceId,omitempty" json:"recurrenceId,omitempty"`
}

type AppointmentResponseActor interface {
	gofhirclient.Resource

	Is_AppointmentResponseActor()
}

func (p Patient) Is_AppointmentResponseActor()           {}
func (g Group) Is_AppointmentResponseActor()             {}
func (p Practitioner) Is_AppointmentResponseActor()      {}
func (p PractitionerRole) Is_AppointmentResponseActor()  {}
func (r RelatedPerson) Is_AppointmentResponseActor()     {}
func (d Device) Is_AppointmentResponseActor()            {}
func (h HealthcareService) Is_AppointmentResponseActor() {}
func (l Location) Is_AppointmentResponseActor()          {}

func (a AppointmentResponse) ResourceType() string {
	return "StructureDefinition"
}
