// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Encounter
type EncounterParticipant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Role of participant in encounter
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Period of time during the encounter that the participant participated
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// The individual, device, or service participating in the encounter
	Actor *custom.Reference[EncounterParticipantActor] `bson:"actor,omitempty" json:"actor,omitempty"`
}

type EncounterReason struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What the reason value should be used for/as
	Use []CodeableConcept `bson:"use,omitempty" json:"use,omitempty"`
	// Reason the encounter takes place (core or reference)
	Value []custom.CodeableReference[EncounterReasonValue] `bson:"value,omitempty" json:"value,omitempty"`
}

type EncounterDiagnosis struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The diagnosis relevant to the encounter
	Condition []custom.CodeableReference[Condition] `bson:"condition,omitempty" json:"condition,omitempty"`
	// Role that this diagnosis has within the encounter (e.g. admission, billing, discharge …)
	Use []CodeableConcept `bson:"use,omitempty" json:"use,omitempty"`
}

type EncounterAdmission struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Pre-admission identifier
	PreAdmissionIdentifier *Identifier `bson:"preAdmissionIdentifier,omitempty" json:"preAdmissionIdentifier,omitempty"`
	// The location/organization from which the patient came before admission
	Origin *custom.Reference[EncounterAdmissionOrigin] `bson:"origin,omitempty" json:"origin,omitempty"`
	// From where patient was admitted (physician referral, transfer)
	AdmitSource *CodeableConcept `bson:"admitSource,omitempty" json:"admitSource,omitempty"`
	// Indicates that the patient is being re-admitted
	ReAdmission *CodeableConcept `bson:"reAdmission,omitempty" json:"reAdmission,omitempty"`
	// Location/organization to which the patient is discharged
	Destination *custom.Reference[EncounterAdmissionDestination] `bson:"destination,omitempty" json:"destination,omitempty"`
	// Category or kind of location after discharge
	DischargeDisposition *CodeableConcept `bson:"dischargeDisposition,omitempty" json:"dischargeDisposition,omitempty"`
}

type EncounterLocation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Location the encounter takes place
	Location custom.Reference[Location] `bson:"location" json:"location"`
	// planned | active | reserved | completed
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// The physical type of the location (usually the level in the location hierarchy - bed, room, ward, virtual etc.)
	Form *CodeableConcept `bson:"form,omitempty" json:"form,omitempty"`
	// Time period during which the patient was present at the location
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type Encounter struct {
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
	// Identifier(s) by which this encounter is known
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// planned | in-progress | on-hold | discharged | completed | cancelled | discontinued | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Classification of patient encounter context - e.g. Inpatient, outpatient
	Class []CodeableConcept `bson:"class,omitempty" json:"class,omitempty"`
	// Indicates the urgency of the encounter
	Priority *CodeableConcept `bson:"priority,omitempty" json:"priority,omitempty"`
	// Specific type of encounter (e.g. e-mail consultation, surgical day-care, ...)
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Specific type of service
	ServiceType []custom.CodeableReference[HealthcareService] `bson:"serviceType,omitempty" json:"serviceType,omitempty"`
	// The patient or group related to this encounter
	Subject *custom.Reference[EncounterSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// The current status of the subject in relation to the Encounter
	SubjectStatus *CodeableConcept `bson:"subjectStatus,omitempty" json:"subjectStatus,omitempty"`
	// Episode(s) of care that this encounter should be recorded against
	EpisodeOfCare []custom.Reference[EpisodeOfCare] `bson:"episodeOfCare,omitempty" json:"episodeOfCare,omitempty"`
	// The request that initiated this encounter
	BasedOn []custom.Reference[EncounterBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// The group(s) that are allocated to participate in this encounter
	CareTeam []custom.Reference[CareTeam] `bson:"careTeam,omitempty" json:"careTeam,omitempty"`
	// Another Encounter this encounter is part of
	PartOf *custom.Reference[Encounter] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// The organization (facility) responsible for this encounter
	ServiceProvider *custom.Reference[Organization] `bson:"serviceProvider,omitempty" json:"serviceProvider,omitempty"`
	// List of participants involved in the encounter
	Participant []EncounterParticipant `bson:"participant,omitempty" json:"participant,omitempty"`
	// The appointment that scheduled this encounter
	Appointment []custom.Reference[Appointment] `bson:"appointment,omitempty" json:"appointment,omitempty"`
	// Connection details of a virtual service (e.g. conference call)
	VirtualService []VirtualServiceDetail `bson:"virtualService,omitempty" json:"virtualService,omitempty"`
	// The actual start and end time of the encounter
	ActualPeriod *Period `bson:"actualPeriod,omitempty" json:"actualPeriod,omitempty"`
	// The planned start date/time (or admission date) of the encounter
	PlannedStartDate *custom.DateTime `bson:"plannedStartDate,omitempty" json:"plannedStartDate,omitempty"`
	// The planned end date/time (or discharge date) of the encounter
	PlannedEndDate *custom.DateTime `bson:"plannedEndDate,omitempty" json:"plannedEndDate,omitempty"`
	// Actual quantity of time the encounter lasted (less time absent)
	Length *Duration `bson:"length,omitempty" json:"length,omitempty"`
	// The list of medical reasons that are expected to be addressed during the episode of care
	Reason []EncounterReason `bson:"reason,omitempty" json:"reason,omitempty"`
	// The list of diagnosis relevant to this encounter
	Diagnosis []EncounterDiagnosis `bson:"diagnosis,omitempty" json:"diagnosis,omitempty"`
	// The set of accounts that may be used for billing for this Encounter
	Account []custom.Reference[Account] `bson:"account,omitempty" json:"account,omitempty"`
	// Diet preferences reported by the patient
	DietPreference []CodeableConcept `bson:"dietPreference,omitempty" json:"dietPreference,omitempty"`
	// Wheelchair, translator, stretcher, etc
	SpecialArrangement []CodeableConcept `bson:"specialArrangement,omitempty" json:"specialArrangement,omitempty"`
	// Special courtesies (VIP, board member)
	SpecialCourtesy []CodeableConcept `bson:"specialCourtesy,omitempty" json:"specialCourtesy,omitempty"`
	// Details about the admission to a healthcare service
	Admission *EncounterAdmission `bson:"admission,omitempty" json:"admission,omitempty"`
	// List of locations where the patient has been
	Location []EncounterLocation `bson:"location,omitempty" json:"location,omitempty"`
}

type EncounterSubject interface {
	gofhirclient.Resource

	Is_EncounterSubject()
}

func (p Patient) Is_EncounterSubject() {}
func (g Group) Is_EncounterSubject()   {}

type EncounterBasedOn interface {
	gofhirclient.Resource

	Is_EncounterBasedOn()
}

func (c CarePlan) Is_EncounterBasedOn()          {}
func (d DeviceRequest) Is_EncounterBasedOn()     {}
func (m MedicationRequest) Is_EncounterBasedOn() {}
func (s ServiceRequest) Is_EncounterBasedOn()    {}

type EncounterParticipantActor interface {
	gofhirclient.Resource

	Is_EncounterParticipantActor()
}

func (p Patient) Is_EncounterParticipantActor()           {}
func (g Group) Is_EncounterParticipantActor()             {}
func (r RelatedPerson) Is_EncounterParticipantActor()     {}
func (p Practitioner) Is_EncounterParticipantActor()      {}
func (p PractitionerRole) Is_EncounterParticipantActor()  {}
func (d Device) Is_EncounterParticipantActor()            {}
func (h HealthcareService) Is_EncounterParticipantActor() {}

type EncounterReasonValue interface {
	gofhirclient.Resource

	Is_EncounterReasonValue()
}

func (c Condition) Is_EncounterReasonValue()                  {}
func (d DiagnosticReport) Is_EncounterReasonValue()           {}
func (o Observation) Is_EncounterReasonValue()                {}
func (i ImmunizationRecommendation) Is_EncounterReasonValue() {}
func (p Procedure) Is_EncounterReasonValue()                  {}

type EncounterAdmissionOrigin interface {
	gofhirclient.Resource

	Is_EncounterAdmissionOrigin()
}

func (l Location) Is_EncounterAdmissionOrigin()     {}
func (o Organization) Is_EncounterAdmissionOrigin() {}

type EncounterAdmissionDestination interface {
	gofhirclient.Resource

	Is_EncounterAdmissionDestination()
}

func (l Location) Is_EncounterAdmissionDestination()     {}
func (o Organization) Is_EncounterAdmissionDestination() {}

func (e Encounter) ResourceType() string {
	return "StructureDefinition"
}
