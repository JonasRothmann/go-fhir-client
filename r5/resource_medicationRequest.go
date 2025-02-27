// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/MedicationRequest
type MedicationRequest struct {
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
	// External ids for this request
	Identifier []Identifier `json:"identifier,omitempty"`
	// A plan or request that is fulfilled in whole or in part by this medication request
	BasedOn []custom.Reference[MedicationRequestBasedOn] `json:"basedOn,omitempty"`
	// Reference to an order/prescription that is being replaced by this MedicationRequest
	PriorPrescription *custom.Reference[MedicationRequest] `json:"priorPrescription,omitempty"`
	// Composite request this is part of
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	// active | on-hold | ended | stopped | completed | cancelled | entered-in-error | draft | unknown
	Status custom.Code `json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	// When the status was changed
	StatusChanged *custom.DateTime `json:"statusChanged,omitempty"`
	// proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `json:"intent"`
	// Grouping or category of medication request
	Category []CodeableConcept `json:"category,omitempty"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// True if patient is to stop taking or not to start taking the medication
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	// Medication to be taken
	Medication custom.CodeableReference[Medication] `json:"medication"`
	// Individual or group for whom the medication has been requested
	Subject custom.Reference[MedicationRequestSubject] `json:"subject"`
	// The person or organization who provided the information about this request, if the source is someone other than the requestor
	InformationSource []custom.Reference[MedicationRequestInformationSource] `json:"informationSource,omitempty"`
	// Encounter created as part of encounter/admission/stay
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Information to support fulfilling of the medication
	SupportingInformation []custom.Reference[Resource] `json:"supportingInformation,omitempty"`
	// When request was initially authored
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Who/What requested the Request
	Requester *custom.Reference[MedicationRequestRequester] `json:"requester,omitempty"`
	// Reported rather than primary record
	Reported *bool `json:"reported,omitempty"`
	// Desired kind of performer of the medication administration
	PerformerType *CodeableConcept `json:"performerType,omitempty"`
	// Intended performer of administration
	Performer []custom.Reference[MedicationRequestPerformer] `json:"performer,omitempty"`
	// Intended type of device for the administration
	Device []custom.CodeableReference[DeviceDefinition] `json:"device,omitempty"`
	// Person who entered the request
	Recorder *custom.Reference[MedicationRequestRecorder] `json:"recorder,omitempty"`
	// Reason or indication for ordering or not ordering the medication
	Reason []custom.CodeableReference[MedicationRequestReason] `json:"reason,omitempty"`
	// Overall pattern of medication administration
	CourseOfTherapyType *CodeableConcept `json:"courseOfTherapyType,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[MedicationRequestInsurance] `json:"insurance,omitempty"`
	// Information about the prescription
	Note []Annotation `json:"note,omitempty"`
	// Full representation of the dosage instructions
	RenderedDosageInstruction *custom.Markdown `json:"renderedDosageInstruction,omitempty"`
	// Period over which the medication is to be taken
	EffectiveDosePeriod *Period `json:"effectiveDosePeriod,omitempty"`
	// Specific instructions for how the medication should be taken
	DosageInstruction []Dosage `json:"dosageInstruction,omitempty"`
	// Medication supply authorization
	DispenseRequest *MedicationRequestDispenseRequest `json:"dispenseRequest,omitempty"`
	// Any restrictions on medication substitution
	Substitution *MedicationRequestSubstitution `json:"substitution,omitempty"`
	// A list of events of interest in the lifecycle
	EventHistory []custom.Reference[Provenance] `json:"eventHistory,omitempty"`
}

type MedicationRequestDispenseRequest struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// First fill details
	InitialFill *MedicationRequestDispenseRequestInitialFill `json:"initialFill,omitempty"`
	// Minimum period of time between dispenses
	DispenseInterval *Duration `json:"dispenseInterval,omitempty"`
	// Time period supply is authorized for
	ValidityPeriod *Period `json:"validityPeriod,omitempty"`
	// Number of refills authorized
	NumberOfRepeatsAllowed *uint32 `json:"numberOfRepeatsAllowed,omitempty"`
	// Amount of medication to supply per dispense
	Quantity *Quantity `json:"quantity,omitempty"`
	// Number of days supply per dispense
	ExpectedSupplyDuration *Duration `json:"expectedSupplyDuration,omitempty"`
	// Intended performer of dispense
	Dispenser *custom.Reference[Organization] `json:"dispenser,omitempty"`
	// Additional information for the dispenser
	DispenserInstruction []Annotation `json:"dispenserInstruction,omitempty"`
	// Type of adherence packaging to use for the dispense
	DoseAdministrationAid *CodeableConcept `json:"doseAdministrationAid,omitempty"`
}

type MedicationRequestDispenseRequestInitialFill struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// First fill quantity
	Quantity *Quantity `json:"quantity,omitempty"`
	// First fill duration
	Duration *Duration `json:"duration,omitempty"`
}

type MedicationRequestSubstitution struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Whether substitution is allowed or not
	AllowedBoolean *bool `json:"allowedBoolean,omitempty"`
	// Whether substitution is allowed or not
	AllowedCodeableConcept *CodeableConcept `json:"allowedCodeableConcept,omitempty"`
	// Why should (not) substitution be made
	Reason *CodeableConcept `json:"reason,omitempty"`
}

type OtherMedicationRequest MedicationRequest

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
	return "MedicationRequest"
}

func (m MedicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicationRequest
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherMedicationRequest: OtherMedicationRequest(m), ResourceType: m.ResourceType()})
}

func UnmarshalMedicationRequest(b []byte) (res MedicationRequest, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
