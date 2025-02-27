// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Observation
type Observation struct {
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
	BasedOn []custom.Reference[ObservationBasedOn] `json:"basedOn,omitempty"`
	// Triggering observation(s)
	TriggeredBy []ObservationTriggeredBy `json:"triggeredBy,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[ObservationPartOf] `json:"partOf,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `json:"status"`
	// Classification of  type of observation
	Category []CodeableConcept `json:"category,omitempty"`
	// Type of observation (code / type)
	Code CodeableConcept `json:"code"`
	// Who and/or what the observation is about
	Subject *custom.Reference[ObservationSubject] `json:"subject,omitempty"`
	// What the observation is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `json:"focus,omitempty"`
	// Healthcare event during which this observation is made
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Clinically relevant time/time-period for observation
	EffectiveDateTime *custom.DateTime `json:"effectiveDateTime,omitempty"`
	// Clinically relevant time/time-period for observation
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// Clinically relevant time/time-period for observation
	EffectiveTiming *Timing `json:"effectiveTiming,omitempty"`
	// Clinically relevant time/time-period for observation
	EffectiveInstant *custom.Instant `json:"effectiveInstant,omitempty"`
	// Date/Time this version was made available
	Issued *custom.Instant `json:"issued,omitempty"`
	// Who is responsible for the observation
	Performer []custom.Reference[ObservationPerformer] `json:"performer,omitempty"`
	// Actual result
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Actual result
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Actual result
	ValueString *string `json:"valueString,omitempty"`
	// Actual result
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Actual result
	ValueInteger *int32 `json:"valueInteger,omitempty"`
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
	// Actual result
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Actual result
	ValueReference *custom.Reference[MolecularSequence] `json:"valueReference,omitempty"`
	// Why the result is missing
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	// High, low, normal, etc
	Interpretation []CodeableConcept `json:"interpretation,omitempty"`
	// Comments about the observation
	Note []Annotation `json:"note,omitempty"`
	// Observed body part
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	// Observed body structure
	BodyStructure *custom.Reference[BodyStructure] `json:"bodyStructure,omitempty"`
	// How it was done
	Method *CodeableConcept `json:"method,omitempty"`
	// Specimen used for this observation
	Specimen *custom.Reference[ObservationSpecimen] `json:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device
	Device *custom.Reference[ObservationDevice] `json:"device,omitempty"`
	// Provides guide for interpretation
	ReferenceRange []ObservationReferenceRange `json:"referenceRange,omitempty"`
	// Related resource that belongs to the Observation group
	HasMember []custom.Reference[ObservationHasMember] `json:"hasMember,omitempty"`
	// Related resource from which the observation is made
	DerivedFrom []custom.Reference[ObservationDerivedFrom] `json:"derivedFrom,omitempty"`
	// Component results
	Component []ObservationComponent `json:"component,omitempty"`
}

type ObservationTriggeredBy struct {
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

type ObservationReferenceRange struct {
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

type ObservationComponent struct {
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

type OtherObservation Observation

type ObservationBasedOn interface {
	gofhirclient.Resource

	Is_ObservationBasedOn()
}

func (c CarePlan) Is_ObservationBasedOn()                   {}
func (d DeviceRequest) Is_ObservationBasedOn()              {}
func (i ImmunizationRecommendation) Is_ObservationBasedOn() {}
func (m MedicationRequest) Is_ObservationBasedOn()          {}
func (n NutritionOrder) Is_ObservationBasedOn()             {}
func (s ServiceRequest) Is_ObservationBasedOn()             {}

type ObservationPartOf interface {
	gofhirclient.Resource

	Is_ObservationPartOf()
}

func (m MedicationAdministration) Is_ObservationPartOf() {}
func (m MedicationDispense) Is_ObservationPartOf()       {}
func (m MedicationStatement) Is_ObservationPartOf()      {}
func (p Procedure) Is_ObservationPartOf()                {}
func (i Immunization) Is_ObservationPartOf()             {}
func (i ImagingStudy) Is_ObservationPartOf()             {}
func (g GenomicStudy) Is_ObservationPartOf()             {}

type ObservationSubject interface {
	gofhirclient.Resource

	Is_ObservationSubject()
}

func (p Patient) Is_ObservationSubject()                    {}
func (g Group) Is_ObservationSubject()                      {}
func (d Device) Is_ObservationSubject()                     {}
func (l Location) Is_ObservationSubject()                   {}
func (o Organization) Is_ObservationSubject()               {}
func (p Procedure) Is_ObservationSubject()                  {}
func (p Practitioner) Is_ObservationSubject()               {}
func (m Medication) Is_ObservationSubject()                 {}
func (s Substance) Is_ObservationSubject()                  {}
func (b BiologicallyDerivedProduct) Is_ObservationSubject() {}
func (n NutritionProduct) Is_ObservationSubject()           {}

type ObservationPerformer interface {
	gofhirclient.Resource

	Is_ObservationPerformer()
}

func (p Practitioner) Is_ObservationPerformer()     {}
func (p PractitionerRole) Is_ObservationPerformer() {}
func (o Organization) Is_ObservationPerformer()     {}
func (c CareTeam) Is_ObservationPerformer()         {}
func (p Patient) Is_ObservationPerformer()          {}
func (r RelatedPerson) Is_ObservationPerformer()    {}

type ObservationSpecimen interface {
	gofhirclient.Resource

	Is_ObservationSpecimen()
}

func (s Specimen) Is_ObservationSpecimen() {}
func (g Group) Is_ObservationSpecimen()    {}

type ObservationDevice interface {
	gofhirclient.Resource

	Is_ObservationDevice()
}

func (d Device) Is_ObservationDevice()       {}
func (d DeviceMetric) Is_ObservationDevice() {}

type ObservationHasMember interface {
	gofhirclient.Resource

	Is_ObservationHasMember()
}

func (o Observation) Is_ObservationHasMember()           {}
func (q QuestionnaireResponse) Is_ObservationHasMember() {}
func (m MolecularSequence) Is_ObservationHasMember()     {}

type ObservationDerivedFrom interface {
	gofhirclient.Resource

	Is_ObservationDerivedFrom()
}

func (d DocumentReference) Is_ObservationDerivedFrom()     {}
func (i ImagingStudy) Is_ObservationDerivedFrom()          {}
func (i ImagingSelection) Is_ObservationDerivedFrom()      {}
func (q QuestionnaireResponse) Is_ObservationDerivedFrom() {}
func (o Observation) Is_ObservationDerivedFrom()           {}
func (m MolecularSequence) Is_ObservationDerivedFrom()     {}
func (g GenomicStudy) Is_ObservationDerivedFrom()          {}

func (o Observation) ResourceType() string {
	return "Observation"
}

func (o Observation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherObservation: OtherObservation(o), ResourceType: o.ResourceType()})
}

func UnmarshalObservation(b []byte) (res Observation, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
