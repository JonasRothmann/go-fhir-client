// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/vitalspanel
type ObservationvitalspanelTriggeredBy struct {
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

type ObservationvitalspanelCode struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Code defined by a terminology system
	Coding []Coding `json:"coding,omitempty"`
	// Plain text representation of the concept
	Text *string `json:"text,omitempty"`
}

type ObservationvitalspanelReferenceRange struct {
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

type ObservationvitalspanelComponent struct {
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

type Observationvitalspanel struct {
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
	BasedOn []custom.Reference[ObservationvitalspanelBasedOn] `json:"basedOn,omitempty"`
	// Triggering observation(s)
	TriggeredBy []ObservationvitalspanelTriggeredBy `json:"triggeredBy,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[ObservationvitalspanelPartOf] `json:"partOf,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `json:"status"`
	// Classification of  type of observation
	Category []CodeableConcept `json:"category"`
	// Vital Signs Panel
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
	Performer []custom.Reference[ObservationvitalspanelPerformer] `json:"performer,omitempty"`
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
	Specimen *custom.Reference[ObservationvitalspanelSpecimen] `json:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device
	Device *custom.Reference[ObservationvitalspanelDevice] `json:"device,omitempty"`
	// Provides guide for interpretation
	ReferenceRange []ObservationvitalspanelReferenceRange `json:"referenceRange,omitempty"`
	// Used when reporting vital signs panel components
	HasMember []custom.Reference[ObservationvitalspanelHasMember] `json:"hasMember"`
	// Related resource from which the observation is made
	DerivedFrom []custom.Reference[ObservationvitalspanelDerivedFrom] `json:"derivedFrom,omitempty"`
	// Used when reporting systolic and diastolic blood pressure.
	Component []ObservationvitalspanelComponent `json:"component,omitempty"`
}

type OtherObservationvitalspanel Observationvitalspanel

type ObservationvitalspanelBasedOn interface {
	gofhirclient.Resource

	Is_ObservationvitalspanelBasedOn()
}

func (c CarePlan) Is_ObservationvitalspanelBasedOn()                   {}
func (d DeviceRequest) Is_ObservationvitalspanelBasedOn()              {}
func (i ImmunizationRecommendation) Is_ObservationvitalspanelBasedOn() {}
func (m MedicationRequest) Is_ObservationvitalspanelBasedOn()          {}
func (n NutritionOrder) Is_ObservationvitalspanelBasedOn()             {}
func (s ServiceRequest) Is_ObservationvitalspanelBasedOn()             {}

type ObservationvitalspanelPartOf interface {
	gofhirclient.Resource

	Is_ObservationvitalspanelPartOf()
}

func (m MedicationAdministration) Is_ObservationvitalspanelPartOf() {}
func (m MedicationDispense) Is_ObservationvitalspanelPartOf()       {}
func (m MedicationStatement) Is_ObservationvitalspanelPartOf()      {}
func (p Procedure) Is_ObservationvitalspanelPartOf()                {}
func (i Immunization) Is_ObservationvitalspanelPartOf()             {}
func (i ImagingStudy) Is_ObservationvitalspanelPartOf()             {}
func (g GenomicStudy) Is_ObservationvitalspanelPartOf()             {}

type ObservationvitalspanelPerformer interface {
	gofhirclient.Resource

	Is_ObservationvitalspanelPerformer()
}

func (p Practitioner) Is_ObservationvitalspanelPerformer()     {}
func (p PractitionerRole) Is_ObservationvitalspanelPerformer() {}
func (o Organization) Is_ObservationvitalspanelPerformer()     {}
func (c CareTeam) Is_ObservationvitalspanelPerformer()         {}
func (p Patient) Is_ObservationvitalspanelPerformer()          {}
func (r RelatedPerson) Is_ObservationvitalspanelPerformer()    {}

type ObservationvitalspanelSpecimen interface {
	gofhirclient.Resource

	Is_ObservationvitalspanelSpecimen()
}

func (s Specimen) Is_ObservationvitalspanelSpecimen() {}
func (g Group) Is_ObservationvitalspanelSpecimen()    {}

type ObservationvitalspanelDevice interface {
	gofhirclient.Resource

	Is_ObservationvitalspanelDevice()
}

func (d Device) Is_ObservationvitalspanelDevice()       {}
func (d DeviceMetric) Is_ObservationvitalspanelDevice() {}

type ObservationvitalspanelHasMember interface {
	gofhirclient.Resource

	Is_ObservationvitalspanelHasMember()
}

func (q QuestionnaireResponse) Is_ObservationvitalspanelHasMember() {}
func (m MolecularSequence) Is_ObservationvitalspanelHasMember()     {}

type ObservationvitalspanelDerivedFrom interface {
	gofhirclient.Resource

	Is_ObservationvitalspanelDerivedFrom()
}

func (d DocumentReference) Is_ObservationvitalspanelDerivedFrom()     {}
func (i ImagingStudy) Is_ObservationvitalspanelDerivedFrom()          {}
func (q QuestionnaireResponse) Is_ObservationvitalspanelDerivedFrom() {}
func (m MolecularSequence) Is_ObservationvitalspanelDerivedFrom()     {}

func (o Observationvitalspanel) ResourceType() string {
	return "Observationvitalspanel"
}

func (o Observationvitalspanel) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationvitalspanel
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherObservationvitalspanel: OtherObservationvitalspanel(o), ResourceType: o.ResourceType()})
}

func UnmarshalObservationvitalspanel(b []byte) (res Observationvitalspanel, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
