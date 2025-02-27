// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/bodytemp
type ObservationbodytempCode struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Code defined by a terminology system
	Coding []Coding `json:"coding,omitempty"`
	// Plain text representation of the concept
	Text *string `json:"text,omitempty"`
}

type ObservationbodytempReferenceRange struct {
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

type ObservationbodytempComponent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of component observation (code / type)
	Code CodeableConcept `json:"code"`
	// Vital Sign Value
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Vital Sign Value
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Vital Sign Value
	ValueString *string `json:"valueString,omitempty"`
	// Vital Sign Value
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Vital Sign Value
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Vital Sign Value
	ValueRange *Range `json:"valueRange,omitempty"`
	// Vital Sign Value
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Vital Sign Value
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Vital Sign Value
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Vital Sign Value
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Vital Sign Value
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Vital Sign Value
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Vital Sign Value
	ValueReference *custom.Reference[MolecularSequence] `json:"valueReference,omitempty"`
	// Why the component result is missing
	DataAbsentReason *CodeableConcept `json:"dataAbsentReason,omitempty"`
	// High, low, normal, etc
	Interpretation []CodeableConcept `json:"interpretation,omitempty"`
	// Provides guide for interpretation of component result
	ReferenceRange []interface{} `json:"referenceRange,omitempty"`
}

type Observationbodytemp struct {
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
	BasedOn []custom.Reference[ObservationbodytempBasedOn] `json:"basedOn,omitempty"`
	// Triggering observation(s)
	TriggeredBy []ObservationbodytempTriggeredBy `json:"triggeredBy,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[ObservationbodytempPartOf] `json:"partOf,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `json:"status"`
	// Classification of  type of observation
	Category []CodeableConcept `json:"category"`
	// Body Temperature
	Code CodeableConcept `json:"code"`
	// Who and/or what the observation is about
	Subject custom.Reference[Patient] `json:"subject"`
	// What the observation is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `json:"focus,omitempty"`
	// Healthcare event during which this observation is made
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Often just a dateTime for Vital Signs
	EffectiveDateTime *custom.DateTime `json:"effectiveDateTime,omitempty"`
	// Often just a dateTime for Vital Signs
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// Date/Time this version was made available
	Issued *custom.Instant `json:"issued,omitempty"`
	// Who is responsible for the observation
	Performer []custom.Reference[ObservationbodytempPerformer] `json:"performer,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueString *string `json:"valueString,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueRange *Range `json:"valueRange,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
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
	Specimen *custom.Reference[ObservationbodytempSpecimen] `json:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device
	Device *custom.Reference[ObservationbodytempDevice] `json:"device,omitempty"`
	// Provides guide for interpretation
	ReferenceRange []ObservationbodytempReferenceRange `json:"referenceRange,omitempty"`
	// Used when reporting vital signs panel components
	HasMember []custom.Reference[ObservationbodytempHasMember] `json:"hasMember,omitempty"`
	// Related resource from which the observation is made
	DerivedFrom []custom.Reference[ObservationbodytempDerivedFrom] `json:"derivedFrom,omitempty"`
	// Used when reporting systolic and diastolic blood pressure.
	Component []ObservationbodytempComponent `json:"component,omitempty"`
}

type ObservationbodytempTriggeredBy struct {
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

type OtherObservationbodytemp Observationbodytemp

type ObservationbodytempBasedOn interface {
	gofhirclient.Resource

	Is_ObservationbodytempBasedOn()
}

func (c CarePlan) Is_ObservationbodytempBasedOn()                   {}
func (d DeviceRequest) Is_ObservationbodytempBasedOn()              {}
func (i ImmunizationRecommendation) Is_ObservationbodytempBasedOn() {}
func (m MedicationRequest) Is_ObservationbodytempBasedOn()          {}
func (n NutritionOrder) Is_ObservationbodytempBasedOn()             {}
func (s ServiceRequest) Is_ObservationbodytempBasedOn()             {}

type ObservationbodytempPartOf interface {
	gofhirclient.Resource

	Is_ObservationbodytempPartOf()
}

func (m MedicationAdministration) Is_ObservationbodytempPartOf() {}
func (m MedicationDispense) Is_ObservationbodytempPartOf()       {}
func (m MedicationStatement) Is_ObservationbodytempPartOf()      {}
func (p Procedure) Is_ObservationbodytempPartOf()                {}
func (i Immunization) Is_ObservationbodytempPartOf()             {}
func (i ImagingStudy) Is_ObservationbodytempPartOf()             {}
func (g GenomicStudy) Is_ObservationbodytempPartOf()             {}

type ObservationbodytempPerformer interface {
	gofhirclient.Resource

	Is_ObservationbodytempPerformer()
}

func (p Practitioner) Is_ObservationbodytempPerformer()     {}
func (p PractitionerRole) Is_ObservationbodytempPerformer() {}
func (o Organization) Is_ObservationbodytempPerformer()     {}
func (c CareTeam) Is_ObservationbodytempPerformer()         {}
func (p Patient) Is_ObservationbodytempPerformer()          {}
func (r RelatedPerson) Is_ObservationbodytempPerformer()    {}

type ObservationbodytempSpecimen interface {
	gofhirclient.Resource

	Is_ObservationbodytempSpecimen()
}

func (s Specimen) Is_ObservationbodytempSpecimen() {}
func (g Group) Is_ObservationbodytempSpecimen()    {}

type ObservationbodytempDevice interface {
	gofhirclient.Resource

	Is_ObservationbodytempDevice()
}

func (d Device) Is_ObservationbodytempDevice()       {}
func (d DeviceMetric) Is_ObservationbodytempDevice() {}

type ObservationbodytempHasMember interface {
	gofhirclient.Resource

	Is_ObservationbodytempHasMember()
}

func (q QuestionnaireResponse) Is_ObservationbodytempHasMember() {}
func (m MolecularSequence) Is_ObservationbodytempHasMember()     {}

type ObservationbodytempDerivedFrom interface {
	gofhirclient.Resource

	Is_ObservationbodytempDerivedFrom()
}

func (d DocumentReference) Is_ObservationbodytempDerivedFrom()     {}
func (i ImagingStudy) Is_ObservationbodytempDerivedFrom()          {}
func (q QuestionnaireResponse) Is_ObservationbodytempDerivedFrom() {}
func (m MolecularSequence) Is_ObservationbodytempDerivedFrom()     {}

func (o Observationbodytemp) ResourceType() string {
	return "Observationbodytemp"
}

func (o Observationbodytemp) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationbodytemp
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherObservationbodytemp: OtherObservationbodytemp(o), ResourceType: o.ResourceType()})
}

func UnmarshalObservationbodytemp(b []byte) (res Observationbodytemp, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
