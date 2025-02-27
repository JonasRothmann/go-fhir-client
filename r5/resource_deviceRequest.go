// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceRequest
type DeviceRequest struct {
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
	// External Request identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[DeviceRequestInstantiatesCanonical] `json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `json:"instantiatesUri,omitempty"`
	// What request fulfills
	BasedOn []custom.Reference[Resource] `json:"basedOn,omitempty"`
	// What request replaces
	Replaces []custom.Reference[DeviceRequest] `json:"replaces,omitempty"`
	// Identifier of composite request
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status *custom.Code `json:"status,omitempty"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// True if the request is to stop or not to start using the device
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	// Device requested
	Code custom.CodeableReference[DeviceRequestCode] `json:"code"`
	// Quantity of devices to supply
	Quantity *int32 `json:"quantity,omitempty"`
	// Device details
	Parameter []DeviceRequestParameter `json:"parameter,omitempty"`
	// Focus of request
	Subject custom.Reference[DeviceRequestSubject] `json:"subject"`
	// Encounter motivating request
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Desired time or schedule for use
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// Desired time or schedule for use
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// Desired time or schedule for use
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	// When recorded
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Who/what submitted the device request
	Requester *custom.Reference[DeviceRequestRequester] `json:"requester,omitempty"`
	// Requested Filler
	Performer *custom.CodeableReference[DeviceRequestPerformer] `json:"performer,omitempty"`
	// Coded/Linked Reason for request
	Reason []custom.CodeableReference[DeviceRequestReason] `json:"reason,omitempty"`
	// PRN status of request
	AsNeeded *bool `json:"asNeeded,omitempty"`
	// Device usage reason
	AsNeededFor *CodeableConcept `json:"asNeededFor,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[DeviceRequestInsurance] `json:"insurance,omitempty"`
	// Additional clinical information
	SupportingInfo []custom.Reference[Resource] `json:"supportingInfo,omitempty"`
	// Notes or comments
	Note []Annotation `json:"note,omitempty"`
	// Request provenance
	RelevantHistory []custom.Reference[Provenance] `json:"relevantHistory,omitempty"`
}

type DeviceRequestParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Device detail
	Code *CodeableConcept `json:"code,omitempty"`
	// Value of detail
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Value of detail
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Value of detail
	ValueRange *Range `json:"valueRange,omitempty"`
	// Value of detail
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
}

type OtherDeviceRequest DeviceRequest

type DeviceRequestInstantiatesCanonical interface {
	gofhirclient.Resource

	Is_DeviceRequestInstantiatesCanonical()
}

func (a ActivityDefinition) Is_DeviceRequestInstantiatesCanonical() {}
func (p PlanDefinition) Is_DeviceRequestInstantiatesCanonical()     {}

type DeviceRequestCode interface {
	gofhirclient.Resource

	Is_DeviceRequestCode()
}

func (d Device) Is_DeviceRequestCode()           {}
func (d DeviceDefinition) Is_DeviceRequestCode() {}

type DeviceRequestSubject interface {
	gofhirclient.Resource

	Is_DeviceRequestSubject()
}

func (p Patient) Is_DeviceRequestSubject()  {}
func (g Group) Is_DeviceRequestSubject()    {}
func (l Location) Is_DeviceRequestSubject() {}
func (d Device) Is_DeviceRequestSubject()   {}

type DeviceRequestRequester interface {
	gofhirclient.Resource

	Is_DeviceRequestRequester()
}

func (d Device) Is_DeviceRequestRequester()           {}
func (p Practitioner) Is_DeviceRequestRequester()     {}
func (p PractitionerRole) Is_DeviceRequestRequester() {}
func (o Organization) Is_DeviceRequestRequester()     {}

type DeviceRequestPerformer interface {
	gofhirclient.Resource

	Is_DeviceRequestPerformer()
}

func (p Practitioner) Is_DeviceRequestPerformer()      {}
func (p PractitionerRole) Is_DeviceRequestPerformer()  {}
func (o Organization) Is_DeviceRequestPerformer()      {}
func (c CareTeam) Is_DeviceRequestPerformer()          {}
func (h HealthcareService) Is_DeviceRequestPerformer() {}
func (p Patient) Is_DeviceRequestPerformer()           {}
func (d Device) Is_DeviceRequestPerformer()            {}
func (r RelatedPerson) Is_DeviceRequestPerformer()     {}

type DeviceRequestReason interface {
	gofhirclient.Resource

	Is_DeviceRequestReason()
}

func (c Condition) Is_DeviceRequestReason()         {}
func (o Observation) Is_DeviceRequestReason()       {}
func (d DiagnosticReport) Is_DeviceRequestReason()  {}
func (d DocumentReference) Is_DeviceRequestReason() {}

type DeviceRequestInsurance interface {
	gofhirclient.Resource

	Is_DeviceRequestInsurance()
}

func (c Coverage) Is_DeviceRequestInsurance()      {}
func (c ClaimResponse) Is_DeviceRequestInsurance() {}

func (d DeviceRequest) ResourceType() string {
	return "DeviceRequest"
}

func (d DeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceRequest
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDeviceRequest: OtherDeviceRequest(d), ResourceType: d.ResourceType()})
}

func UnmarshalDeviceRequest(b []byte) (res DeviceRequest, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
