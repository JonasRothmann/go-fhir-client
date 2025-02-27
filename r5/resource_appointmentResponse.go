// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AppointmentResponse
type AppointmentResponse struct {
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
	// Appointment this response relates to
	Appointment custom.Reference[Appointment] `json:"appointment"`
	// Indicator for a counter proposal
	ProposedNewTime *bool `json:"proposedNewTime,omitempty"`
	// Time from appointment, or requested new start time
	Start *custom.Instant `json:"start,omitempty"`
	// Time from appointment, or requested new end time
	End *custom.Instant `json:"end,omitempty"`
	// Role of participant in the appointment
	ParticipantType []CodeableConcept `json:"participantType,omitempty"`
	// Person(s), Location, HealthcareService, or Device
	Actor *custom.Reference[AppointmentResponseActor] `json:"actor,omitempty"`
	// accepted | declined | tentative | needs-action | entered-in-error
	ParticipantStatus custom.Code `json:"participantStatus"`
	// Additional comments
	Comment *custom.Markdown `json:"comment,omitempty"`
	// This response is for all occurrences in a recurring request
	Recurring *bool `json:"recurring,omitempty"`
	// Original date within a recurring request
	OccurrenceDate *custom.Date `json:"occurrenceDate,omitempty"`
	// The recurrence ID of the specific recurring request
	RecurrenceId *uint32 `json:"recurrenceId,omitempty"`
}

type OtherAppointmentResponse AppointmentResponse

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
	return "AppointmentResponse"
}

func (a AppointmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointmentResponse
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAppointmentResponse: OtherAppointmentResponse(a), ResourceType: a.ResourceType()})
}

func UnmarshalAppointmentResponse(b []byte) (res AppointmentResponse, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
