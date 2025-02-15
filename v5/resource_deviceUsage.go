// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/DeviceUsage
type DeviceUsage struct {
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
	// External identifier for this record
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Fulfills plan, proposal or order
	BasedOn []custom.Reference[ServiceRequest] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// active | completed | not-done | entered-in-error +
	Status custom.Code `bson:"status" json:"status"`
	// The category of the statement - classifying how the statement is made
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Patient using device
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Supporting information
	DerivedFrom []custom.Reference[DeviceUsageDerivedFrom] `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	// The encounter or episode of care that establishes the context for this device use statement
	Context *custom.Reference[DeviceUsageContext] `bson:"context,omitempty" json:"context,omitempty"`
	// How often  the device was used
	Timing *Timing `bson:"timing,omitempty" json:"timing,omitempty"`
	// When the statement was made (and recorded)
	DateAsserted *custom.DateTime `bson:"dateAsserted,omitempty" json:"dateAsserted,omitempty"`
	// The status of the device usage, for example always, sometimes, never. This is not the same as the status of the statement
	UsageStatus *CodeableConcept `bson:"usageStatus,omitempty" json:"usageStatus,omitempty"`
	// The reason for asserting the usage status - for example forgot, lost, stolen, broken
	UsageReason []CodeableConcept `bson:"usageReason,omitempty" json:"usageReason,omitempty"`
	// How device is being used
	Adherence *DeviceUsageAdherence `bson:"adherence,omitempty" json:"adherence,omitempty"`
	// Who made the statement
	InformationSource *custom.Reference[DeviceUsageInformationSource] `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	// Code or Reference to device used
	Device custom.CodeableReference[DeviceUsageDevice] `bson:"device" json:"device"`
	// Why device was used
	Reason []custom.CodeableReference[DeviceUsageReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Target body site
	BodySite *custom.CodeableReference[BodyStructure] `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Addition details (comments, instructions)
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type DeviceUsageAdherence struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// always | never | sometimes
	Code CodeableConcept `bson:"code" json:"code"`
	// lost | stolen | prescribed | broken | burned | forgot
	Reason []CodeableConcept `bson:"reason" json:"reason"`
}

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
	return "StructureDefinition"
}
