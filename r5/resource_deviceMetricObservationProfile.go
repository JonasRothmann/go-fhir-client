// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/devicemetricobservation
type DeviceMetricObservationProfile struct {
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
	// Business Identifier for observation
	Identifier []Identifier `json:"identifier,omitempty"`
	// Instantiates FHIR ObservationDefinition
	InstantiatesCanonical *custom.Canonical[ObservationDefinition] `json:"instantiatesCanonical,omitempty"`
	// Instantiates FHIR ObservationDefinition
	InstantiatesReference *custom.Reference[ObservationDefinition] `json:"instantiatesReference,omitempty"`
	// Fulfills plan, proposal or order
	BasedOn []custom.Reference[DeviceMetricObservationProfileBasedOn] `json:"basedOn,omitempty"`
	// Triggering observation(s)
	TriggeredBy []DeviceMetricObservationProfileTriggeredBy `json:"triggeredBy,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[DeviceMetricObservationProfilePartOf] `json:"partOf,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `json:"status"`
	// Classification of  type of observation
	Category []CodeableConcept `json:"category,omitempty"`
	// Type of observation (code / type)
	Code CodeableConcept `json:"code"`
	// Who and/or what the observation is about
	Subject custom.Reference[DeviceMetricObservationProfileSubject] `json:"subject"`
	// What the observation is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `json:"focus,omitempty"`
	// Healthcare event during which this observation is made
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Clinically relevant time/time-period for observation
	EffectiveDateTime *custom.DateTime `json:"effectiveDateTime,omitempty"`
	// Date/Time this version was made available
	Issued *custom.Instant `json:"issued,omitempty"`
	// Who is responsible for the observation
	Performer []custom.Reference[DeviceMetricObservationProfilePerformer] `json:"performer,omitempty"`
	// Actual result
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Actual result
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Actual result
	ValueString *string `json:"valueString,omitempty"`
	// Actual result
	ValueRange *Range `json:"valueRange,omitempty"`
	// Actual result
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Actual result
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Actual result
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Actual result
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Actual result
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Why the result is missing
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	// High, low, normal, etc
	Interpretation *CodeableConcept `json:"interpretation,omitempty"`
	// Comments about the observation
	Note []Annotation `json:"note,omitempty"`
	// Observed body part
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	// Observed body structure
	BodyStructure *custom.Reference[BodyStructure] `json:"bodyStructure,omitempty"`
	// How it was done
	Method *CodeableConcept `json:"method,omitempty"`
	// Specimen used for this observation
	Specimen *custom.Reference[Specimen] `json:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device
	Device custom.Reference[DeviceMetric] `json:"device"`
	// Provides guide for interpretation
	ReferenceRange *DeviceMetricObservationProfileReferenceRange `json:"referenceRange,omitempty"`
	// Related resource that belongs to the Observation group
	HasMember []custom.Reference[Observation] `json:"hasMember,omitempty"`
	// Related resource from which the observation is made
	DerivedFrom []custom.Reference[Observation] `json:"derivedFrom,omitempty"`
	// Component results
	Component []DeviceMetricObservationProfileComponent `json:"component,omitempty"`
}

type DeviceMetricObservationProfileTriggeredBy struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Triggering observation
	Observation custom.Reference[Observation] `json:"observation"`
	// reflex | repeat | re-run
	Type custom.Code `json:"type"`
	// Reason that the observation was triggered
	Reason *string `json:"reason,omitempty"`
}

type DeviceMetricObservationProfileReferenceRange struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Low Range, if relevant
	Low *Quantity `json:"low,omitempty"`
	// High Range, if relevant
	High *Quantity `json:"high,omitempty"`
	// Normal value, if relevant
	NormalValue *CodeableConcept `json:"normalValue,omitempty"`
	// Reference range qualifier
	Type *CodeableConcept `json:"type,omitempty"`
	// Reference range population
	AppliesTo []CodeableConcept `json:"appliesTo,omitempty"`
	// Applicable age range, if relevant
	Age *Range `json:"age,omitempty"`
	// Text based reference range in an observation
	Text *custom.Markdown `json:"text,omitempty"`
}

type DeviceMetricObservationProfileComponent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of component observation (code / type)
	Code CodeableConcept `json:"code"`
	// Actual component result
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Actual component result
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Actual component result
	ValueString *string `json:"valueString,omitempty"`
	// Actual component result
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Actual component result
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Actual component result
	ValueRange *Range `json:"valueRange,omitempty"`
	// Actual component result
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Actual component result
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Actual component result
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Actual component result
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Actual component result
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Actual component result
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Actual component result
	ValueReference *custom.Reference[MolecularSequence] `json:"valueReference,omitempty"`
	// Why the component result is missing
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	// High, low, normal, etc
	Interpretation []CodeableConcept `json:"interpretation,omitempty"`
	// Provides guide for interpretation of component result
	ReferenceRange []interface{} `json:"referenceRange,omitempty"`
}

type OtherDeviceMetricObservationProfile DeviceMetricObservationProfile

type DeviceMetricObservationProfileBasedOn interface {
	gofhirclient.Resource

	Is_DeviceMetricObservationProfileBasedOn()
}

func (c CarePlan) Is_DeviceMetricObservationProfileBasedOn()                   {}
func (d DeviceRequest) Is_DeviceMetricObservationProfileBasedOn()              {}
func (i ImmunizationRecommendation) Is_DeviceMetricObservationProfileBasedOn() {}
func (m MedicationRequest) Is_DeviceMetricObservationProfileBasedOn()          {}
func (n NutritionOrder) Is_DeviceMetricObservationProfileBasedOn()             {}
func (s ServiceRequest) Is_DeviceMetricObservationProfileBasedOn()             {}

type DeviceMetricObservationProfilePartOf interface {
	gofhirclient.Resource

	Is_DeviceMetricObservationProfilePartOf()
}

func (m MedicationAdministration) Is_DeviceMetricObservationProfilePartOf() {}
func (m MedicationDispense) Is_DeviceMetricObservationProfilePartOf()       {}
func (m MedicationStatement) Is_DeviceMetricObservationProfilePartOf()      {}
func (p Procedure) Is_DeviceMetricObservationProfilePartOf()                {}
func (i Immunization) Is_DeviceMetricObservationProfilePartOf()             {}
func (i ImagingStudy) Is_DeviceMetricObservationProfilePartOf()             {}
func (g GenomicStudy) Is_DeviceMetricObservationProfilePartOf()             {}

type DeviceMetricObservationProfileSubject interface {
	gofhirclient.Resource

	Is_DeviceMetricObservationProfileSubject()
}

func (p Patient) Is_DeviceMetricObservationProfileSubject() {}
func (d Device) Is_DeviceMetricObservationProfileSubject()  {}

type DeviceMetricObservationProfilePerformer interface {
	gofhirclient.Resource

	Is_DeviceMetricObservationProfilePerformer()
}

func (p Practitioner) Is_DeviceMetricObservationProfilePerformer()     {}
func (p PractitionerRole) Is_DeviceMetricObservationProfilePerformer() {}
func (o Organization) Is_DeviceMetricObservationProfilePerformer()     {}
func (c CareTeam) Is_DeviceMetricObservationProfilePerformer()         {}
func (p Patient) Is_DeviceMetricObservationProfilePerformer()          {}
func (r RelatedPerson) Is_DeviceMetricObservationProfilePerformer()    {}

func (d DeviceMetricObservationProfile) ResourceType() string {
	return "DeviceMetricObservationProfile"
}

func (d DeviceMetricObservationProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceMetricObservationProfile
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDeviceMetricObservationProfile: OtherDeviceMetricObservationProfile(d), ResourceType: d.ResourceType()})
}

func UnmarshalDeviceMetricObservationProfile(b []byte) (res DeviceMetricObservationProfile, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
