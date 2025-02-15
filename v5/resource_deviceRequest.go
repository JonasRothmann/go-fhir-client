// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceRequest
type DeviceRequest struct {
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
	// External Request identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Instantiates FHIR protocol or definition
	InstantiatesCanonical []custom.Canonical[DeviceRequestInstantiatesCanonical] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// Instantiates external protocol or definition
	InstantiatesUri []custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// What request fulfills
	BasedOn []custom.Reference[Resource] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// What request replaces
	Replaces []custom.Reference[DeviceRequest] `bson:"replaces,omitempty" json:"replaces,omitempty"`
	// Identifier of composite request
	GroupIdentifier *Identifier `bson:"groupIdentifier,omitempty" json:"groupIdentifier,omitempty"`
	// draft | active | on-hold | revoked | completed | entered-in-error | unknown
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `bson:"intent" json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `bson:"priority,omitempty" json:"priority,omitempty"`
	// True if the request is to stop or not to start using the device
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	// Device requested
	Code custom.CodeableReference[DeviceRequestCode] `bson:"code" json:"code"`
	// Quantity of devices to supply
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Device details
	Parameter []DeviceRequestParameter `bson:"parameter,omitempty" json:"parameter,omitempty"`
	// Focus of request
	Subject custom.Reference[DeviceRequestSubject] `bson:"subject" json:"subject"`
	// Encounter motivating request
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Desired time or schedule for use
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// When recorded
	AuthoredOn *custom.DateTime `bson:"authoredOn,omitempty" json:"authoredOn,omitempty"`
	// Who/what submitted the device request
	Requester *custom.Reference[DeviceRequestRequester] `bson:"requester,omitempty" json:"requester,omitempty"`
	// Requested Filler
	Performer *custom.CodeableReference[DeviceRequestPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Coded/Linked Reason for request
	Reason []custom.CodeableReference[DeviceRequestReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// PRN status of request
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
	// Device usage reason
	AsNeededFor *CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[DeviceRequestInsurance] `bson:"insurance,omitempty" json:"insurance,omitempty"`
	// Additional clinical information
	SupportingInfo []custom.Reference[Resource] `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Notes or comments
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Request provenance
	RelevantHistory []custom.Reference[Provenance] `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
}

type DeviceRequestParameter struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Device detail
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Value of detail
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

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
	return "StructureDefinition"
}
