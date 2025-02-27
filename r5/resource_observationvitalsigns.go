// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/vitalsigns
type ObservationvitalsignsComponent struct {
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

type Observationvitalsigns struct {
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
	BasedOn []custom.Reference[ObservationvitalsignsBasedOn] `json:"basedOn,omitempty"`
	// Triggering observation(s)
	TriggeredBy []ObservationvitalsignsTriggeredBy `json:"triggeredBy,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[ObservationvitalsignsPartOf] `json:"partOf,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `json:"status"`
	// Classification of  type of observation
	Category []CodeableConcept `json:"category"`
	// Coded Responses from C-CDA Vital Sign Results
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
	Performer []custom.Reference[ObservationvitalsignsPerformer] `json:"performer,omitempty"`
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
	Specimen *custom.Reference[ObservationvitalsignsSpecimen] `json:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device
	Device *custom.Reference[ObservationvitalsignsDevice] `json:"device,omitempty"`
	// Provides guide for interpretation
	ReferenceRange []ObservationvitalsignsReferenceRange `json:"referenceRange,omitempty"`
	// Used when reporting vital signs panel components
	HasMember []custom.Reference[ObservationvitalsignsHasMember] `json:"hasMember,omitempty"`
	// Related resource from which the observation is made
	DerivedFrom []custom.Reference[ObservationvitalsignsDerivedFrom] `json:"derivedFrom,omitempty"`
	// Used when reporting systolic and diastolic blood pressure.
	Component []ObservationvitalsignsComponent `json:"component,omitempty"`
}

type ObservationvitalsignsTriggeredBy struct {
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

type ObservationvitalsignsReferenceRange struct {
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

type OtherObservationvitalsigns Observationvitalsigns

type ObservationvitalsignsBasedOn interface {
	gofhirclient.Resource

	Is_ObservationvitalsignsBasedOn()
}

func (c CarePlan) Is_ObservationvitalsignsBasedOn()                   {}
func (d DeviceRequest) Is_ObservationvitalsignsBasedOn()              {}
func (i ImmunizationRecommendation) Is_ObservationvitalsignsBasedOn() {}
func (m MedicationRequest) Is_ObservationvitalsignsBasedOn()          {}
func (n NutritionOrder) Is_ObservationvitalsignsBasedOn()             {}
func (s ServiceRequest) Is_ObservationvitalsignsBasedOn()             {}

type ObservationvitalsignsPartOf interface {
	gofhirclient.Resource

	Is_ObservationvitalsignsPartOf()
}

func (m MedicationAdministration) Is_ObservationvitalsignsPartOf() {}
func (m MedicationDispense) Is_ObservationvitalsignsPartOf()       {}
func (m MedicationStatement) Is_ObservationvitalsignsPartOf()      {}
func (p Procedure) Is_ObservationvitalsignsPartOf()                {}
func (i Immunization) Is_ObservationvitalsignsPartOf()             {}
func (i ImagingStudy) Is_ObservationvitalsignsPartOf()             {}
func (g GenomicStudy) Is_ObservationvitalsignsPartOf()             {}

type ObservationvitalsignsPerformer interface {
	gofhirclient.Resource

	Is_ObservationvitalsignsPerformer()
}

func (p Practitioner) Is_ObservationvitalsignsPerformer()     {}
func (p PractitionerRole) Is_ObservationvitalsignsPerformer() {}
func (o Organization) Is_ObservationvitalsignsPerformer()     {}
func (c CareTeam) Is_ObservationvitalsignsPerformer()         {}
func (p Patient) Is_ObservationvitalsignsPerformer()          {}
func (r RelatedPerson) Is_ObservationvitalsignsPerformer()    {}

type ObservationvitalsignsSpecimen interface {
	gofhirclient.Resource

	Is_ObservationvitalsignsSpecimen()
}

func (s Specimen) Is_ObservationvitalsignsSpecimen() {}
func (g Group) Is_ObservationvitalsignsSpecimen()    {}

type ObservationvitalsignsDevice interface {
	gofhirclient.Resource

	Is_ObservationvitalsignsDevice()
}

func (d Device) Is_ObservationvitalsignsDevice()       {}
func (d DeviceMetric) Is_ObservationvitalsignsDevice() {}

type ObservationvitalsignsHasMember interface {
	gofhirclient.Resource

	Is_ObservationvitalsignsHasMember()
}

func (q QuestionnaireResponse) Is_ObservationvitalsignsHasMember() {}
func (m MolecularSequence) Is_ObservationvitalsignsHasMember()     {}

type ObservationvitalsignsDerivedFrom interface {
	gofhirclient.Resource

	Is_ObservationvitalsignsDerivedFrom()
}

func (d DocumentReference) Is_ObservationvitalsignsDerivedFrom()     {}
func (i ImagingStudy) Is_ObservationvitalsignsDerivedFrom()          {}
func (q QuestionnaireResponse) Is_ObservationvitalsignsDerivedFrom() {}
func (m MolecularSequence) Is_ObservationvitalsignsDerivedFrom()     {}

func (o Observationvitalsigns) ResourceType() string {
	return "Observationvitalsigns"
}

func (o Observationvitalsigns) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationvitalsigns
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherObservationvitalsigns: OtherObservationvitalsigns(o), ResourceType: o.ResourceType()})
}

func UnmarshalObservationvitalsigns(b []byte) (res Observationvitalsigns, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
