// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicationRequest
type MedicationRequest struct {
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
	// External ids for this request
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// A plan or request that is fulfilled in whole or in part by this medication request
	BasedOn []custom.Reference[MedicationRequestBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Reference to an order/prescription that is being replaced by this MedicationRequest
	PriorPrescription *custom.Reference[MedicationRequest] `bson:"priorPrescription,omitempty" json:"priorPrescription,omitempty"`
	// Composite request this is part of
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	// active | on-hold | ended | stopped | completed | cancelled | entered-in-error | draft | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// When the status was changed
	StatusChanged *custom.DateTime `bson:"statusChanged,omitempty" json:"statusChanged,omitempty"`
	// proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `bson:"intent" json:"intent"`
	// Grouping or category of medication request
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// True if patient is to stop taking or not to start taking the medication
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	// Medication to be taken
	Medication custom.CodeableReference[Medication] `bson:"medication" json:"medication"`
	// Individual or group for whom the medication has been requested
	Subject custom.Reference[MedicationRequestSubject] `bson:"subject" json:"subject"`
	// The person or organization who provided the information about this request, if the source is someone other than the requestor
	InformationSource []custom.Reference[MedicationRequestInformationSource] `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	// Encounter created as part of encounter/admission/stay
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Information to support fulfilling of the medication
	SupportingInformation []custom.Reference[Resource] `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	// When request was initially authored
	AuthoredOn *custom.DateTime `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	// Who/What requested the Request
	Requester *custom.Reference[MedicationRequestRequester] `bson:"requester,omitempty" json:"requester,omitempty"`
	// Reported rather than primary record
	Reported *bool `bson:"reported,omitempty" json:"reported,omitempty"`
	// Desired kind of performer of the medication administration
	PerformerType *CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty"`
	// Intended performer of administration
	Performer []custom.Reference[MedicationRequestPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Intended type of device for the administration
	Device []custom.CodeableReference[DeviceDefinition] `bson:"device,omitempty" json:"device,omitempty"`
	// Person who entered the request
	Recorder *custom.Reference[MedicationRequestRecorder] `bson:"recorder,omitempty" json:"recorder,omitempty"`
	// Reason or indication for ordering or not ordering the medication
	Reason []custom.CodeableReference[MedicationRequestReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Overall pattern of medication administration
	CourseOfTherapyType *CodeableConcept `bson:"courseOfTherapyType,omitempty" json:"courseOfTherapyType,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[MedicationRequestInsurance] `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Information about the prescription
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Full representation of the dosage instructions
	RenderedDosageInstruction *custom.Markdown `bson:"renderedDosageInstruction,omitempty" json:"renderedDosageInstruction,omitempty"`
	// Period over which the medication is to be taken
	EffectiveDosePeriod *Period `bson:"effectiveDosePeriod,omitempty" json:"effectiveDosePeriod,omitempty"`
	// Specific instructions for how the medication should be taken
	DosageInstruction []Dosage `bson:"dosageInstruction,omitempty" json:"dosageInstruction,omitempty"`
	// Medication supply authorization
	DispenseRequest *MedicationRequestDispenseRequest `bson:"dispenseRequest,omitempty" json:"dispenseRequest,omitempty"`
	// Any restrictions on medication substitution
	Substitution *MedicationRequestSubstitution `bson:"substitution,omitempty" json:"substitution,omitempty"`
	// A list of events of interest in the lifecycle
	EventHistory []custom.Reference[Provenance] `bson:"eventHistory,omitempty" json:"eventHistory,omitempty"`
}

type MedicationRequestDispenseRequest struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// First fill details
	InitialFill *MedicationRequestDispenseRequestInitialFill `bson:"initialFill,omitempty" json:"initialFill,omitempty"`
	// Minimum period of time between dispenses
	DispenseInterval *Duration `bson:"dispenseInterval,omitempty" json:"dispenseInterval,omitempty"`
	// Time period supply is authorized for
	ValidityPeriod *Period `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	// Number of refills authorized
	NumberOfRepeatsAllowed *uint32 `bson:"numberOfRepeatsAllowed,omitempty" json:"numberOfRepeatsAllowed,omitempty"`
	// Amount of medication to supply per dispense
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Number of days supply per dispense
	ExpectedSupplyDuration *Duration `bson:"expectedSupplyDuration,omitempty" json:"expectedSupplyDuration,omitempty"`
	// Intended performer of dispense
	Dispenser *custom.Reference[Organization] `bson:"dispenser,omitempty" json:"dispenser,omitempty"`
	// Additional information for the dispenser
	DispenserInstruction []Annotation `bson:"dispenserInstruction,omitempty" json:"dispenserInstruction,omitempty"`
	// Type of adherence packaging to use for the dispense
	DoseAdministrationAid *CodeableConcept `bson:"doseAdministrationAid,omitempty" json:"doseAdministrationAid,omitempty"`
}

type MedicationRequestDispenseRequestInitialFill struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// First fill quantity
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// First fill duration
	Duration *Duration `bson:"duration,omitempty" json:"duration,omitempty"`
}

type MedicationRequestSubstitution struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Whether substitution is allowed or not
	Allowed bool `bson:"allowed" json:"allowed"`
	// Why should (not) substitution be made
	Reason *CodeableConcept `bson:"reason,omitempty" json:"reason,omitempty"`
}

type MedicationRequestBasedOn interface {
	gofhirclient.Resource

	Is_MedicationRequestBasedOn()
}

func (c CarePlan) Is_MedicationRequestBasedOn()                   {}
func (m MedicationRequest) Is_MedicationRequestBasedOn()          {}
func (s ServiceRequest) Is_MedicationRequestBasedOn()             {}
func (i ImmunizationRecommendation) Is_MedicationRequestBasedOn() {}

type MedicationRequestSubject interface {
	gofhirclient.Resource

	Is_MedicationRequestSubject()
}

func (p Patient) Is_MedicationRequestSubject() {}
func (g Group) Is_MedicationRequestSubject()   {}

type MedicationRequestInformationSource interface {
	gofhirclient.Resource

	Is_MedicationRequestInformationSource()
}

func (p Patient) Is_MedicationRequestInformationSource()          {}
func (p Practitioner) Is_MedicationRequestInformationSource()     {}
func (p PractitionerRole) Is_MedicationRequestInformationSource() {}
func (r RelatedPerson) Is_MedicationRequestInformationSource()    {}
func (o Organization) Is_MedicationRequestInformationSource()     {}

type MedicationRequestRequester interface {
	gofhirclient.Resource

	Is_MedicationRequestRequester()
}

func (p Practitioner) Is_MedicationRequestRequester()     {}
func (p PractitionerRole) Is_MedicationRequestRequester() {}
func (o Organization) Is_MedicationRequestRequester()     {}
func (p Patient) Is_MedicationRequestRequester()          {}
func (r RelatedPerson) Is_MedicationRequestRequester()    {}
func (d Device) Is_MedicationRequestRequester()           {}

type MedicationRequestPerformer interface {
	gofhirclient.Resource

	Is_MedicationRequestPerformer()
}

func (p Practitioner) Is_MedicationRequestPerformer()      {}
func (p PractitionerRole) Is_MedicationRequestPerformer()  {}
func (o Organization) Is_MedicationRequestPerformer()      {}
func (p Patient) Is_MedicationRequestPerformer()           {}
func (d DeviceDefinition) Is_MedicationRequestPerformer()  {}
func (r RelatedPerson) Is_MedicationRequestPerformer()     {}
func (c CareTeam) Is_MedicationRequestPerformer()          {}
func (h HealthcareService) Is_MedicationRequestPerformer() {}

type MedicationRequestRecorder interface {
	gofhirclient.Resource

	Is_MedicationRequestRecorder()
}

func (p Practitioner) Is_MedicationRequestRecorder()     {}
func (p PractitionerRole) Is_MedicationRequestRecorder() {}

type MedicationRequestReason interface {
	gofhirclient.Resource

	Is_MedicationRequestReason()
}

func (c Condition) Is_MedicationRequestReason()   {}
func (o Observation) Is_MedicationRequestReason() {}

type MedicationRequestInsurance interface {
	gofhirclient.Resource

	Is_MedicationRequestInsurance()
}

func (c Coverage) Is_MedicationRequestInsurance()      {}
func (c ClaimResponse) Is_MedicationRequestInsurance() {}

func (m MedicationRequest) ResourceType() string {
	return "StructureDefinition"
}
