// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Procedure
type Procedure struct {
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
	// External Identifiers for this procedure
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[ProcedureInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// A request for this procedure
	BasedOn []custom.Reference[ProcedureBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[ProcedurePartOf] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// Classification of the procedure
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Identification of the procedure
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Individual or entity the procedure was performed on
	Subject custom.Reference[ProcedureSubject] `bson:"subject" json:"subject"`
	// Who is the target of the procedure when it is not the subject of record only
	Focus *custom.Reference[ProcedureFocus] `bson:"focus,omitempty" json:"focus,omitempty"`
	// The Encounter during which this Procedure was created
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// When the procedure occurred or is occurring
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// When the procedure was first captured in the subject's record
	Recorded *custom.DateTime `bson:"recorded,omitempty" json:"recorded,omitempty"`
	// Who recorded the procedure
	Recorder *custom.Reference[ProcedureRecorder] `bson:"recorder,omitempty" json:"recorder,omitempty"`
	// Reported rather than primary record
	Reported *bool `bson:"reported,omitempty" json:"reported,omitempty"`
	// Who performed the procedure and what they did
	Performer []ProcedurePerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Where the procedure happened
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// The justification that the procedure was performed
	Reason []custom.CodeableReference[ProcedureReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Target body sites
	BodySite []CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// The result of procedure
	Outcome *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	// Any report resulting from the procedure
	Report []custom.Reference[ProcedureReport] `bson:"report,omitempty" json:"report,omitempty"`
	// Complication following the procedure
	Complication []custom.CodeableReference[Condition] `bson:"complication,omitempty" json:"complication,omitempty"`
	// Instructions for follow up
	FollowUp []CodeableConcept `bson:"followUp,omitempty" json:"followUp,omitempty"`
	// Additional information about the procedure
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Manipulated, implanted, or removed device
	FocalDevice []ProcedureFocalDevice `bson:"focalDevice,omitempty" json:"focalDevice,omitempty"`
	// Items used during procedure
	Used []custom.CodeableReference[ProcedureUsed] `bson:"used,omitempty" json:"used,omitempty"`
	// Extra information relevant to the procedure
	SupportingInfo []custom.Reference[Resource] `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
}

type ProcedurePerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of performance
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Who performed the procedure
	Actor custom.Reference[ProcedurePerformerActor] `bson:"actor" json:"actor"`
	// Organization the device or practitioner was acting for
	OnBehalfOf *custom.Reference[Organization] `bson:"onBehalfOf,omitempty" json:"onBehalfOf,omitempty"`
	// When the performer performed the procedure
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
}

type ProcedureFocalDevice struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Kind of change to device
	Action *CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	// Device that was changed
	Manipulated custom.Reference[Device] `bson:"manipulated" json:"manipulated"`
}

type ProcedureInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_ProcedureInstantiatesCanonical()
}

func (p PlanDefinition) Is_ProcedureInstantiatesCanonical()      {}
func (a ActivityDefinition) Is_ProcedureInstantiatesCanonical()  {}
func (m Measure) Is_ProcedureInstantiatesCanonical()             {}
func (o OperationDefinition) Is_ProcedureInstantiatesCanonical() {}
func (q Questionnaire) Is_ProcedureInstantiatesCanonical()       {}

type ProcedureBasedOn interface {
	gofhirclient.Resource

	Is_ProcedureBasedOn()
}

func (c CarePlan) Is_ProcedureBasedOn()       {}
func (s ServiceRequest) Is_ProcedureBasedOn() {}

type ProcedurePartOf interface {
	gofhirclient.Resource

	Is_ProcedurePartOf()
}

func (p Procedure) Is_ProcedurePartOf()                {}
func (o Observation) Is_ProcedurePartOf()              {}
func (m MedicationAdministration) Is_ProcedurePartOf() {}

type ProcedureSubject interface {
	gofhirclient.Resource

	Is_ProcedureSubject()
}

func (p Patient) Is_ProcedureSubject()      {}
func (g Group) Is_ProcedureSubject()        {}
func (d Device) Is_ProcedureSubject()       {}
func (p Practitioner) Is_ProcedureSubject() {}
func (o Organization) Is_ProcedureSubject() {}
func (l Location) Is_ProcedureSubject()     {}

type ProcedureFocus interface {
	gofhirclient.Resource

	Is_ProcedureFocus()
}

func (p Patient) Is_ProcedureFocus()          {}
func (g Group) Is_ProcedureFocus()            {}
func (r RelatedPerson) Is_ProcedureFocus()    {}
func (p Practitioner) Is_ProcedureFocus()     {}
func (o Organization) Is_ProcedureFocus()     {}
func (c CareTeam) Is_ProcedureFocus()         {}
func (p PractitionerRole) Is_ProcedureFocus() {}
func (s Specimen) Is_ProcedureFocus()         {}

type ProcedureRecorder interface {
	gofhirclient.Resource

	Is_ProcedureRecorder()
}

func (p Patient) Is_ProcedureRecorder()          {}
func (r RelatedPerson) Is_ProcedureRecorder()    {}
func (p Practitioner) Is_ProcedureRecorder()     {}
func (p PractitionerRole) Is_ProcedureRecorder() {}

type ProcedurePerformerActor interface {
	gofhirclient.Resource

	Is_ProcedurePerformerActor()
}

func (p Practitioner) Is_ProcedurePerformerActor()      {}
func (p PractitionerRole) Is_ProcedurePerformerActor()  {}
func (o Organization) Is_ProcedurePerformerActor()      {}
func (p Patient) Is_ProcedurePerformerActor()           {}
func (r RelatedPerson) Is_ProcedurePerformerActor()     {}
func (d Device) Is_ProcedurePerformerActor()            {}
func (c CareTeam) Is_ProcedurePerformerActor()          {}
func (h HealthcareService) Is_ProcedurePerformerActor() {}

type ProcedureReason interface {
	gofhirclient.Resource

	Is_ProcedureReason()
}

func (c Condition) Is_ProcedureReason()         {}
func (o Observation) Is_ProcedureReason()       {}
func (p Procedure) Is_ProcedureReason()         {}
func (d DiagnosticReport) Is_ProcedureReason()  {}
func (d DocumentReference) Is_ProcedureReason() {}

type ProcedureReport interface {
	gofhirclient.Resource

	Is_ProcedureReport()
}

func (d DiagnosticReport) Is_ProcedureReport()  {}
func (d DocumentReference) Is_ProcedureReport() {}
func (c Composition) Is_ProcedureReport()       {}

type ProcedureUsed interface {
	gofhirclient.Resource

	Is_ProcedureUsed()
}

func (d Device) Is_ProcedureUsed()                     {}
func (m Medication) Is_ProcedureUsed()                 {}
func (s Substance) Is_ProcedureUsed()                  {}
func (b BiologicallyDerivedProduct) Is_ProcedureUsed() {}

func (p Procedure) ResourceType() string {
	return "StructureDefinition"
}
