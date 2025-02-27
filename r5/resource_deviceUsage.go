// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceUsage
type DeviceUsage struct {
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
	// External identifier for this record
	Identifier []Identifier `json:"identifier,omitempty"`
	// Fulfills plan, proposal or order
	BasedOn []custom.Reference[ServiceRequest] `json:"basedOn,omitempty"`
	// active | completed | not-done | entered-in-error +
	Status custom.Code `json:"status"`
	// The category of the statement - classifying how the statement is made
	Category []CodeableConcept `json:"category,omitempty"`
	// Patient using device
	Patient custom.Reference[Patient] `json:"patient"`
	// Supporting information
	DerivedFrom []custom.Reference[DeviceUsageDerivedFrom] `json:"derivedFrom,omitempty"`
	// The encounter or episode of care that establishes the context for this device use statement
	Context *custom.Reference[DeviceUsageContext] `json:"context,omitempty"`
	// How often  the device was used
	TimingTiming *Timing `json:"timingTiming,omitempty"`
	// How often  the device was used
	TimingPeriod *Period `json:"timingPeriod,omitempty"`
	// How often  the device was used
	TimingDateTime *custom.DateTime `json:"timingDateTime,omitempty"`
	// When the statement was made (and recorded)
	DateAsserted *custom.DateTime `json:"dateAsserted,omitempty"`
	// The status of the device usage, for example always, sometimes, never. This is not the same as the status of the statement
	UsageStatus *CodeableConcept `json:"usageStatus,omitempty"`
	// The reason for asserting the usage status - for example forgot, lost, stolen, broken
	UsageReason []CodeableConcept `json:"usageReason,omitempty"`
	// How device is being used
	Adherence *DeviceUsageAdherence `json:"adherence,omitempty"`
	// Who made the statement
	InformationSource *custom.Reference[DeviceUsageInformationSource] `json:"informationSource,omitempty"`
	// Code or Reference to device used
	Device custom.CodeableReference[DeviceUsageDevice] `json:"device"`
	// Why device was used
	Reason []custom.CodeableReference[DeviceUsageReason] `json:"reason,omitempty"`
	// Target body site
	BodySite *custom.CodeableReference[BodyStructure] `json:"bodySite,omitempty"`
	// Addition details (comments, instructions)
	Note []Annotation `json:"note,omitempty"`
}

type DeviceUsageAdherence struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// always | never | sometimes
	Code CodeableConcept `json:"code"`
	// lost | stolen | prescribed | broken | burned | forgot
	Reason []CodeableConcept `json:"reason"`
}

type OtherDeviceUsage DeviceUsage

type DeviceUsageDerivedFrom interface {
	gofhirclient.Resource

	Is_DeviceUsageDerivedFrom()
}

func (s ServiceRequest) Is_DeviceUsageDerivedFrom()        {}
func (p Procedure) Is_DeviceUsageDerivedFrom()             {}
func (c Claim) Is_DeviceUsageDerivedFrom()                 {}
func (o Observation) Is_DeviceUsageDerivedFrom()           {}
func (q QuestionnaireResponse) Is_DeviceUsageDerivedFrom() {}
func (d DocumentReference) Is_DeviceUsageDerivedFrom()     {}

type DeviceUsageContext interface {
	gofhirclient.Resource

	Is_DeviceUsageContext()
}

func (e Encounter) Is_DeviceUsageContext()     {}
func (e EpisodeOfCare) Is_DeviceUsageContext() {}

type DeviceUsageInformationSource interface {
	gofhirclient.Resource

	Is_DeviceUsageInformationSource()
}

func (p Patient) Is_DeviceUsageInformationSource()          {}
func (p Practitioner) Is_DeviceUsageInformationSource()     {}
func (p PractitionerRole) Is_DeviceUsageInformationSource() {}
func (r RelatedPerson) Is_DeviceUsageInformationSource()    {}
func (o Organization) Is_DeviceUsageInformationSource()     {}

type DeviceUsageDevice interface {
	gofhirclient.Resource

	Is_DeviceUsageDevice()
}

func (d Device) Is_DeviceUsageDevice()           {}
func (d DeviceDefinition) Is_DeviceUsageDevice() {}

type DeviceUsageReason interface {
	gofhirclient.Resource

	Is_DeviceUsageReason()
}

func (c Condition) Is_DeviceUsageReason()         {}
func (o Observation) Is_DeviceUsageReason()       {}
func (d DiagnosticReport) Is_DeviceUsageReason()  {}
func (d DocumentReference) Is_DeviceUsageReason() {}
func (p Procedure) Is_DeviceUsageReason()         {}

func (d DeviceUsage) ResourceType() string {
	return "DeviceUsage"
}

func (d DeviceUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceUsage
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDeviceUsage: OtherDeviceUsage(d), ResourceType: d.ResourceType()})
}

func UnmarshalDeviceUsage(b []byte) (res DeviceUsage, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
