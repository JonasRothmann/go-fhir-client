// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Appointment
type Appointment struct {
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
	// proposed | pending | booked | arrived | fulfilled | cancelled | noshow | entered-in-error | checked-in | waitlist
	Status custom.Code `json:"status"`
	// The coded reason for the appointment being cancelled
	CancellationReason *CodeableConcept `json:"cancellationReason,omitempty"`
	// Classification when becoming an encounter
	Class []CodeableConcept `json:"class,omitempty"`
	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []CodeableConcept `json:"serviceCategory,omitempty"`
	// The specific service that is to be performed during this appointment
	ServiceType []custom.CodeableReference[HealthcareService] `json:"serviceType,omitempty"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []CodeableConcept `json:"specialty,omitempty"`
	// The style of appointment or patient that has been booked in the slot (not service type)
	AppointmentType *CodeableConcept `json:"appointmentType,omitempty"`
	// Reason this appointment is scheduled
	Reason []custom.CodeableReference[AppointmentReason] `json:"reason,omitempty"`
	// Used to make informed decisions if needing to re-prioritize
	Priority *CodeableConcept `json:"priority,omitempty"`
	// Shown on a subject line in a meeting request, or appointment list
	Description *string `json:"description,omitempty"`
	// Appointment replaced by this Appointment
	Replaces []custom.Reference[Appointment] `json:"replaces,omitempty"`
	// Connection details of a virtual service (e.g. conference call)
	VirtualService []VirtualServiceDetail `json:"virtualService,omitempty"`
	// Additional information to support the appointment
	SupportingInformation []custom.Reference[Resource] `json:"supportingInformation,omitempty"`
	// The previous appointment in a series
	PreviousAppointment *custom.Reference[Appointment] `json:"previousAppointment,omitempty"`
	// The originating appointment in a recurring set of appointments
	OriginatingAppointment *custom.Reference[Appointment] `json:"originatingAppointment,omitempty"`
	// When appointment is to take place
	Start *custom.Instant `json:"start,omitempty"`
	// When appointment is to conclude
	End *custom.Instant `json:"end,omitempty"`
	// Can be less than start/end (e.g. estimate)
	MinutesDuration *uint32 `json:"minutesDuration,omitempty"`
	// Potential date/time interval(s) requested to allocate the appointment within
	RequestedPeriod []Period `json:"requestedPeriod,omitempty"`
	// The slots that this appointment is filling
	Slot []custom.Reference[Slot] `json:"slot,omitempty"`
	// The set of accounts that may be used for billing for this Appointment
	Account []custom.Reference[Account] `json:"account,omitempty"`
	// The date that this appointment was initially created
	Created *custom.DateTime `json:"created,omitempty"`
	// When the appointment was cancelled
	CancellationDate *custom.DateTime `json:"cancellationDate,omitempty"`
	// Additional comments
	Note []Annotation `json:"note,omitempty"`
	// Detailed information and instructions for the patient
	PatientInstruction []custom.CodeableReference[AppointmentPatientInstruction] `json:"patientInstruction,omitempty"`
	// The request this appointment is allocated to assess
	BasedOn []custom.Reference[AppointmentBasedOn] `json:"basedOn,omitempty"`
	// The patient or group associated with the appointment
	Subject *custom.Reference[AppointmentSubject] `json:"subject,omitempty"`
	// Participants involved in appointment
	Participant []AppointmentParticipant `json:"participant"`
	// The sequence number in the recurrence
	RecurrenceId *uint32 `json:"recurrenceId,omitempty"`
	// Indicates that this appointment varies from a recurrence pattern
	OccurrenceChanged *bool `json:"occurrenceChanged,omitempty"`
	// Details of the recurrence pattern/template used to generate occurrences
	RecurrenceTemplate []AppointmentRecurrenceTemplate `json:"recurrenceTemplate,omitempty"`
}

type AppointmentParticipant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Role of participant in the appointment
	Type []CodeableConcept `json:"type,omitempty"`
	// Participation period of the actor
	Period *Period `json:"period,omitempty"`
	// The individual, device, location, or service participating in the appointment
	Actor *custom.Reference[AppointmentParticipantActor] `json:"actor,omitempty"`
	// The participant is required to attend (optional when false)
	Required *bool `json:"required,omitempty"`
	// accepted | declined | tentative | needs-action
	Status custom.Code `json:"status"`
}

type AppointmentRecurrenceTemplate struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The timezone of the occurrences
	Timezone *CodeableConcept `json:"timezone,omitempty"`
	// The frequency of the recurrence
	RecurrenceType CodeableConcept `json:"recurrenceType"`
	// The date when the recurrence should end
	LastOccurrenceDate *custom.Date `json:"lastOccurrenceDate,omitempty"`
	// The number of planned occurrences
	OccurrenceCount *uint32 `json:"occurrenceCount,omitempty"`
	// Specific dates for a recurring set of appointments (no template)
	OccurrenceDate []custom.Date `json:"occurrenceDate,omitempty"`
	// Information about weekly recurring appointments
	WeeklyTemplate *AppointmentRecurrenceTemplateWeeklyTemplate `json:"weeklyTemplate,omitempty"`
	// Information about monthly recurring appointments
	MonthlyTemplate *AppointmentRecurrenceTemplateMonthlyTemplate `json:"monthlyTemplate,omitempty"`
	// Information about yearly recurring appointments
	YearlyTemplate *AppointmentRecurrenceTemplateYearlyTemplate `json:"yearlyTemplate,omitempty"`
	// Any dates that should be excluded from the series
	ExcludingDate []custom.Date `json:"excludingDate,omitempty"`
	// Any recurrence IDs that should be excluded from the recurrence
	ExcludingRecurrenceId []uint32 `json:"excludingRecurrenceId,omitempty"`
}

type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Recurs on Mondays
	Monday *bool `json:"monday,omitempty"`
	// Recurs on Tuesday
	Tuesday *bool `json:"tuesday,omitempty"`
	// Recurs on Wednesday
	Wednesday *bool `json:"wednesday,omitempty"`
	// Recurs on Thursday
	Thursday *bool `json:"thursday,omitempty"`
	// Recurs on Friday
	Friday *bool `json:"friday,omitempty"`
	// Recurs on Saturday
	Saturday *bool `json:"saturday,omitempty"`
	// Recurs on Sunday
	Sunday *bool `json:"sunday,omitempty"`
	// Recurs every nth week
	WeekInterval *uint32 `json:"weekInterval,omitempty"`
}

type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Recurs on a specific day of the month
	DayOfMonth *uint32 `json:"dayOfMonth,omitempty"`
	// Indicates which week of the month the appointment should occur
	NthWeekOfMonth *Coding `json:"nthWeekOfMonth,omitempty"`
	// Indicates which day of the week the appointment should occur
	DayOfWeek *Coding `json:"dayOfWeek,omitempty"`
	// Recurs every nth month
	MonthInterval uint32 `json:"monthInterval"`
}

type AppointmentRecurrenceTemplateYearlyTemplate struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Recurs every nth year
	YearInterval uint32 `json:"yearInterval"`
}

type OtherAppointment Appointment

type AppointmentReason interface {
	gofhirclient.Resource

	Is_AppointmentReason()
}

func (c Condition) Is_AppointmentReason()                  {}
func (p Procedure) Is_AppointmentReason()                  {}
func (o Observation) Is_AppointmentReason()                {}
func (i ImmunizationRecommendation) Is_AppointmentReason() {}

type AppointmentPatientInstruction interface {
	gofhirclient.Resource

	Is_AppointmentPatientInstruction()
}

func (d DocumentReference) Is_AppointmentPatientInstruction() {}
func (b Binary) Is_AppointmentPatientInstruction()            {}
func (c Communication) Is_AppointmentPatientInstruction()     {}

type AppointmentBasedOn interface {
	gofhirclient.Resource

	Is_AppointmentBasedOn()
}

func (c CarePlan) Is_AppointmentBasedOn()          {}
func (d DeviceRequest) Is_AppointmentBasedOn()     {}
func (m MedicationRequest) Is_AppointmentBasedOn() {}
func (s ServiceRequest) Is_AppointmentBasedOn()    {}

type AppointmentSubject interface {
	gofhirclient.Resource

	Is_AppointmentSubject()
}

func (p Patient) Is_AppointmentSubject() {}
func (g Group) Is_AppointmentSubject()   {}

type AppointmentParticipantActor interface {
	gofhirclient.Resource

	Is_AppointmentParticipantActor()
}

func (p Patient) Is_AppointmentParticipantActor()           {}
func (g Group) Is_AppointmentParticipantActor()             {}
func (p Practitioner) Is_AppointmentParticipantActor()      {}
func (p PractitionerRole) Is_AppointmentParticipantActor()  {}
func (c CareTeam) Is_AppointmentParticipantActor()          {}
func (r RelatedPerson) Is_AppointmentParticipantActor()     {}
func (d Device) Is_AppointmentParticipantActor()            {}
func (h HealthcareService) Is_AppointmentParticipantActor() {}
func (l Location) Is_AppointmentParticipantActor()          {}

func (a Appointment) ResourceType() string {
	return "Appointment"
}

func (a Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointment
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAppointment: OtherAppointment(a), ResourceType: a.ResourceType()})
}

func UnmarshalAppointment(b []byte) (res Appointment, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
