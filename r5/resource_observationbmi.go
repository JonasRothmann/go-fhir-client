// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/bmi
type Observationbmi struct {
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
	BasedOn []custom.Reference[ObservationbmiBasedOn] `json:"basedOn,omitempty"`
	// Triggering observation(s)
	TriggeredBy []ObservationbmiTriggeredBy `json:"triggeredBy,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[ObservationbmiPartOf] `json:"partOf,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `json:"status"`
	// Classification of  type of observation
	Category []CodeableConcept `json:"category"`
	// Body Mass Index (BMI)
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
	Performer []custom.Reference[ObservationbmiPerformer] `json:"performer,omitempty"`
	// Vital Signs value are recorded using the Quantity data type. For supporting observations such as Cuff size could use other datatypes such as CodeableConcept.
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
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
	Specimen *custom.Reference[ObservationbmiSpecimen] `json:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device
	Device *custom.Reference[ObservationbmiDevice] `json:"device,omitempty"`
	// Provides guide for interpretation
	ReferenceRange []ObservationbmiReferenceRange `json:"referenceRange,omitempty"`
	// Used when reporting vital signs panel components
	HasMember []custom.Reference[ObservationbmiHasMember] `json:"hasMember,omitempty"`
	// Related resource from which the observation is made
	DerivedFrom []custom.Reference[ObservationbmiDerivedFrom] `json:"derivedFrom,omitempty"`
	// Used when reporting systolic and diastolic blood pressure.
	Component []ObservationbmiComponent `json:"component,omitempty"`
}

type ObservationbmiTriggeredBy struct {
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

type ObservationbmiCode struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Code defined by a terminology system
	Coding []Coding `json:"coding,omitempty"`
	// Plain text representation of the concept
	Text *string `json:"text,omitempty"`
}

type ObservationbmiReferenceRange struct {
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

type ObservationbmiComponent struct {
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

type OtherObservationbmi Observationbmi

type ObservationbmiBasedOn interface {
	gofhirclient.Resource

	Is_ObservationbmiBasedOn()
}

func (c CarePlan) Is_ObservationbmiBasedOn()                   {}
func (d DeviceRequest) Is_ObservationbmiBasedOn()              {}
func (i ImmunizationRecommendation) Is_ObservationbmiBasedOn() {}
func (m MedicationRequest) Is_ObservationbmiBasedOn()          {}
func (n NutritionOrder) Is_ObservationbmiBasedOn()             {}
func (s ServiceRequest) Is_ObservationbmiBasedOn()             {}

type ObservationbmiPartOf interface {
	gofhirclient.Resource

	Is_ObservationbmiPartOf()
}

func (m MedicationAdministration) Is_ObservationbmiPartOf() {}
func (m MedicationDispense) Is_ObservationbmiPartOf()       {}
func (m MedicationStatement) Is_ObservationbmiPartOf()      {}
func (p Procedure) Is_ObservationbmiPartOf()                {}
func (i Immunization) Is_ObservationbmiPartOf()             {}
func (i ImagingStudy) Is_ObservationbmiPartOf()             {}
func (g GenomicStudy) Is_ObservationbmiPartOf()             {}

type ObservationbmiPerformer interface {
	gofhirclient.Resource

	Is_ObservationbmiPerformer()
}

func (p Practitioner) Is_ObservationbmiPerformer()     {}
func (p PractitionerRole) Is_ObservationbmiPerformer() {}
func (o Organization) Is_ObservationbmiPerformer()     {}
func (c CareTeam) Is_ObservationbmiPerformer()         {}
func (p Patient) Is_ObservationbmiPerformer()          {}
func (r RelatedPerson) Is_ObservationbmiPerformer()    {}

type ObservationbmiSpecimen interface {
	gofhirclient.Resource

	Is_ObservationbmiSpecimen()
}

func (s Specimen) Is_ObservationbmiSpecimen() {}
func (g Group) Is_ObservationbmiSpecimen()    {}

type ObservationbmiDevice interface {
	gofhirclient.Resource

	Is_ObservationbmiDevice()
}

func (d Device) Is_ObservationbmiDevice()       {}
func (d DeviceMetric) Is_ObservationbmiDevice() {}

type ObservationbmiHasMember interface {
	gofhirclient.Resource

	Is_ObservationbmiHasMember()
}

func (q QuestionnaireResponse) Is_ObservationbmiHasMember() {}
func (m MolecularSequence) Is_ObservationbmiHasMember()     {}

type ObservationbmiDerivedFrom interface {
	gofhirclient.Resource

	Is_ObservationbmiDerivedFrom()
}

func (d DocumentReference) Is_ObservationbmiDerivedFrom()     {}
func (i ImagingStudy) Is_ObservationbmiDerivedFrom()          {}
func (q QuestionnaireResponse) Is_ObservationbmiDerivedFrom() {}
func (m MolecularSequence) Is_ObservationbmiDerivedFrom()     {}

func (o Observationbmi) ResourceType() string {
	return "Observationbmi"
}

func (o Observationbmi) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationbmi
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherObservationbmi: OtherObservationbmi(o), ResourceType: o.ResourceType()})
}

func UnmarshalObservationbmi(b []byte) (res Observationbmi, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
