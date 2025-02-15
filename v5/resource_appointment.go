// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Appointment
type AppointmentParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Role of participant in the appointment
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Participation period of the actor
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// The individual, device, location, or service participating in the appointment
	Actor *custom.Reference[AppointmentParticipantActor] `bson:"actor,omitempty" json:"actor,omitempty"`
	// The participant is required to attend (optional when false)
	Required *bool `bson:"required,omitempty" json:"required,omitempty"`
	// accepted | declined | tentative | needs-action
	Status custom.Code `bson:"status" json:"status"`
}

type AppointmentRecurrenceTemplate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The timezone of the occurrences
	Timezone *CodeableConcept `bson:"timezone,omitempty" json:"timezone,omitempty"`
	// The frequency of the recurrence
	RecurrenceType CodeableConcept `bson:"recurrenceType" json:"recurrenceType"`
	// The date when the recurrence should end
	LastOccurrenceDate *custom.Date `bson:"lastOccurrenceDate,omitempty" json:"lastOccurrenceDate,omitempty"`
	// The number of planned occurrences
	OccurrenceCount *uint32 `bson:"occurrenceCount,omitempty" json:"occurrenceCount,omitempty"`
	// Specific dates for a recurring set of appointments (no template)
	OccurrenceDate []custom.Date `bson:"occurrenceDate,omitempty" json:"occurrenceDate,omitempty"`
	// Information about weekly recurring appointments
	WeeklyTemplate *AppointmentRecurrenceTemplateWeeklyTemplate `bson:"weeklyTemplate,omitempty" json:"weeklyTemplate,omitempty"`
	// Information about monthly recurring appointments
	MonthlyTemplate *AppointmentRecurrenceTemplateMonthlyTemplate `bson:"monthlyTemplate,omitempty" json:"monthlyTemplate,omitempty"`
	// Information about yearly recurring appointments
	YearlyTemplate *AppointmentRecurrenceTemplateYearlyTemplate `bson:"yearlyTemplate,omitempty" json:"yearlyTemplate,omitempty"`
	// Any dates that should be excluded from the series
	ExcludingDate []custom.Date `bson:"excludingDate,omitempty" json:"excludingDate,omitempty"`
	// Any recurrence IDs that should be excluded from the recurrence
	ExcludingRecurrenceId []uint32 `bson:"excludingRecurrenceId,omitempty" json:"excludingRecurrenceId,omitempty"`
}

type AppointmentRecurrenceTemplateWeeklyTemplate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Recurs on Mondays
	Monday *bool `bson:"monday,omitempty" json:"monday,omitempty"`
	// Recurs on Tuesday
	Tuesday *bool `bson:"tuesday,omitempty" json:"tuesday,omitempty"`
	// Recurs on Wednesday
	Wednesday *bool `bson:"wednesday,omitempty" json:"wednesday,omitempty"`
	// Recurs on Thursday
	Thursday *bool `bson:"thursday,omitempty" json:"thursday,omitempty"`
	// Recurs on Friday
	Friday *bool `bson:"friday,omitempty" json:"friday,omitempty"`
	// Recurs on Saturday
	Saturday *bool `bson:"saturday,omitempty" json:"saturday,omitempty"`
	// Recurs on Sunday
	Sunday *bool `bson:"sunday,omitempty" json:"sunday,omitempty"`
	// Recurs every nth week
	WeekInterval *uint32 `bson:"weekInterval,omitempty" json:"weekInterval,omitempty"`
}

type AppointmentRecurrenceTemplateMonthlyTemplate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Recurs on a specific day of the month
	DayOfMonth *uint32 `bson:"dayOfMonth,omitempty" json:"dayOfMonth,omitempty"`
	// Indicates which week of the month the appointment should occur
	NthWeekOfMonth *Coding `bson:"nthWeekOfMonth,omitempty" json:"nthWeekOfMonth,omitempty"`
	// Indicates which day of the week the appointment should occur
	DayOfWeek *Coding `bson:"dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`
	// Recurs every nth month
	MonthInterval uint32 `bson:"monthInterval" json:"monthInterval"`
}

type AppointmentRecurrenceTemplateYearlyTemplate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Recurs every nth year
	YearInterval uint32 `bson:"yearInterval" json:"yearInterval"`
}

type Appointment struct {
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
	// proposed | pending | booked | arrived | fulfilled | cancelled | noshow | entered-in-error | checked-in | waitlist
	Status custom.Code `bson:"status" json:"status"`
	// The coded reason for the appointment being cancelled
	CancellationReason *CodeableConcept `bson:"cancellationReason,omitempty" json:"cancellationReason,omitempty"`
	// Classification when becoming an encounter
	Class []CodeableConcept `bson:"class,omitempty" json:"class,omitempty"`
	// A broad categorization of the service that is to be performed during this appointment
	ServiceCategory []CodeableConcept `bson:"serviceCategory,omitempty" json:"serviceCategory,omitempty"`
	// The specific service that is to be performed during this appointment
	ServiceType []custom.CodeableReference[HealthcareService] `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	// The specialty of a practitioner that would be required to perform the service requested in this appointment
	Specialty []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	// The style of appointment or patient that has been booked in the slot (not service type)
	AppointmentType *CodeableConcept `bson:"appointmentType,omitempty" json:"appointmentType,omitempty"`
	// Reason this appointment is scheduled
	Reason []custom.CodeableReference[AppointmentReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Used to make informed decisions if needing to re-prioritize
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty"`
	// Shown on a subject line in a meeting request, or appointment list
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
	// Appointment replaced by this Appointment
	Replaces []custom.Reference[Appointment] `bson:"replaces,omitempty" json:"replaces,omitempty"`
	// Connection details of a virtual service (e.g. conference call)
	VirtualService []VirtualServiceDetail `bson:"virtualService,omitempty" json:"virtualService,omitempty"`
	// Additional information to support the appointment
	SupportingInformation []custom.Reference[Resource] `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	// The previous appointment in a series
	PreviousAppointment *custom.Reference[Appointment] `bson:"previousAppointment,omitempty" json:"previousAppointment,omitempty"`
	// The originating appointment in a recurring set of appointments
	OriginatingAppointment *custom.Reference[Appointment] `bson:"originatingAppointment,omitempty" json:"originatingAppointment,omitempty"`
	// When appointment is to take place
	Start *custom.Instant `bson:"start,omitempty" json:"start,omitempty"`
	// When appointment is to conclude
	End *custom.Instant `bson:"end,omitempty" json:"end,omitempty"`
	// Can be less than start/end (e.g. estimate)
	MinutesDuration *uint32 `bson:"minutesDuration,omitempty" json:"minutesDuration,omitempty"`
	// Potential date/time interval(s) requested to allocate the appointment within
	RequestedPeriod []Period `bson:"requestedPeriod,omitempty" json:"requestedPeriod,omitempty"`
	// The slots that this appointment is filling
	Slot []custom.Reference[Slot] `bson:"slot,omitempty" json:"slot,omitempty"`
	// The set of accounts that may be used for billing for this Appointment
	Account []custom.Reference[Account] `bson:"account,omitempty" json:"account,omitempty"`
	// The date that this appointment was initially created
	Created *custom.DateTime `bson:"created,omitempty" json:"created,omitempty"`
	// When the appointment was cancelled
	CancellationDate *custom.DateTime `bson:"cancellationDate,omitempty" json:"cancellationDate,omitempty"`
	// Additional comments
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Detailed information and instructions for the patient
	PatientInstruction []custom.CodeableReference[AppointmentPatientInstruction] `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	// The request this appointment is allocated to assess
	BasedOn []custom.Reference[AppointmentBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// The patient or group associated with the appointment
	Subject *custom.Reference[AppointmentSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Participants involved in appointment
	Participant []AppointmentParticipant `bson:"participant" json:"participant"`
	// The sequence number in the recurrence
	RecurrenceId *uint32 `bson:"recurrenceId,omitempty" json:"recurrenceId,omitempty"`
	// Indicates that this appointment varies from a recurrence pattern
	OccurrenceChanged *bool `bson:"occurrenceChanged,omitempty" json:"occurrenceChanged,omitempty"`
	// Details of the recurrence pattern/template used to generate occurrences
	RecurrenceTemplate []AppointmentRecurrenceTemplate `bson:"recurrenceTemplate,omitempty" json:"recurrenceTemplate,omitempty"`
}

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
	return "StructureDefinition"
}
