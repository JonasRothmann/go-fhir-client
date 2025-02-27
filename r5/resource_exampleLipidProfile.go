// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ldlcholesterol
type ExampleLipidProfile struct {
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
	BasedOn []custom.Reference[ExampleLipidProfileBasedOn] `json:"basedOn,omitempty"`
	// Triggering observation(s)
	TriggeredBy []ExampleLipidProfileTriggeredBy `json:"triggeredBy,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[ExampleLipidProfilePartOf] `json:"partOf,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `json:"status"`
	// Classification of  type of observation
	Category []CodeableConcept `json:"category,omitempty"`
	// LDL Cholesterol -measured or calculated per code
	Code CodeableConcept `json:"code"`
	// Who and/or what the observation is about
	Subject *custom.Reference[ExampleLipidProfileSubject] `json:"subject,omitempty"`
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
	Performer []custom.Reference[ExampleLipidProfilePerformer] `json:"performer,omitempty"`
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
	// + | ++ | +++ | - | -- | ---
	Interpretation *CodeableConcept `json:"interpretation,omitempty"`
	// Comments about result
	Note []Annotation `json:"note,omitempty"`
	// Observed body part
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	// Observed body structure
	BodyStructure *custom.Reference[BodyStructure] `json:"bodyStructure,omitempty"`
	// How it was done
	Method *CodeableConcept `json:"method,omitempty"`
	// Specimen used for this observation
	Specimen *custom.Reference[ExampleLipidProfileSpecimen] `json:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device
	Device *custom.Reference[ExampleLipidProfileDevice] `json:"device,omitempty"`
	// Provides guide for interpretation
	ReferenceRange ExampleLipidProfileReferenceRange `json:"referenceRange"`
	// Related resource that belongs to the Observation group
	HasMember *custom.Reference[ExampleLipidProfileHasMember] `json:"hasMember,omitempty"`
	// Related resource from which the observation is made
	DerivedFrom *custom.Reference[ExampleLipidProfileDerivedFrom] `json:"derivedFrom,omitempty"`
	// Component results
	Component []ExampleLipidProfileComponent `json:"component,omitempty"`
}

type ExampleLipidProfileTriggeredBy struct {
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

type ExampleLipidProfileReferenceRange struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Low Range, if relevant
	Low *Quantity `json:"low,omitempty"`
	// A fixed quantity (no comparator)
	High Quantity `json:"high"`
	// Normal value, if relevant
	NormalValue *CodeableConcept `json:"normalValue,omitempty"`
	// Reference range qualifier
	Type *CodeableConcept `json:"type,omitempty"`
	// Reference range population
	AppliesTo *CodeableConcept `json:"appliesTo,omitempty"`
	// Applicable age range, if relevant
	Age *Range `json:"age,omitempty"`
	// Text based reference range in an observation
	Text *custom.Markdown `json:"text,omitempty"`
}

type ExampleLipidProfileComponent struct {
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

type OtherExampleLipidProfile ExampleLipidProfile

type ExampleLipidProfileBasedOn interface {
	gofhirclient.Resource

	Is_ExampleLipidProfileBasedOn()
}

func (c CarePlan) Is_ExampleLipidProfileBasedOn()                   {}
func (d DeviceRequest) Is_ExampleLipidProfileBasedOn()              {}
func (i ImmunizationRecommendation) Is_ExampleLipidProfileBasedOn() {}
func (m MedicationRequest) Is_ExampleLipidProfileBasedOn()          {}
func (n NutritionOrder) Is_ExampleLipidProfileBasedOn()             {}
func (s ServiceRequest) Is_ExampleLipidProfileBasedOn()             {}

type ExampleLipidProfilePartOf interface {
	gofhirclient.Resource

	Is_ExampleLipidProfilePartOf()
}

func (m MedicationAdministration) Is_ExampleLipidProfilePartOf() {}
func (m MedicationDispense) Is_ExampleLipidProfilePartOf()       {}
func (m MedicationStatement) Is_ExampleLipidProfilePartOf()      {}
func (p Procedure) Is_ExampleLipidProfilePartOf()                {}
func (i Immunization) Is_ExampleLipidProfilePartOf()             {}
func (i ImagingStudy) Is_ExampleLipidProfilePartOf()             {}
func (g GenomicStudy) Is_ExampleLipidProfilePartOf()             {}

type ExampleLipidProfileSubject interface {
	gofhirclient.Resource

	Is_ExampleLipidProfileSubject()
}

func (p Patient) Is_ExampleLipidProfileSubject()                    {}
func (g Group) Is_ExampleLipidProfileSubject()                      {}
func (d Device) Is_ExampleLipidProfileSubject()                     {}
func (l Location) Is_ExampleLipidProfileSubject()                   {}
func (o Organization) Is_ExampleLipidProfileSubject()               {}
func (p Procedure) Is_ExampleLipidProfileSubject()                  {}
func (p Practitioner) Is_ExampleLipidProfileSubject()               {}
func (m Medication) Is_ExampleLipidProfileSubject()                 {}
func (s Substance) Is_ExampleLipidProfileSubject()                  {}
func (b BiologicallyDerivedProduct) Is_ExampleLipidProfileSubject() {}
func (n NutritionProduct) Is_ExampleLipidProfileSubject()           {}

type ExampleLipidProfilePerformer interface {
	gofhirclient.Resource

	Is_ExampleLipidProfilePerformer()
}

func (p Practitioner) Is_ExampleLipidProfilePerformer()     {}
func (p PractitionerRole) Is_ExampleLipidProfilePerformer() {}
func (o Organization) Is_ExampleLipidProfilePerformer()     {}
func (c CareTeam) Is_ExampleLipidProfilePerformer()         {}
func (p Patient) Is_ExampleLipidProfilePerformer()          {}
func (r RelatedPerson) Is_ExampleLipidProfilePerformer()    {}

type ExampleLipidProfileSpecimen interface {
	gofhirclient.Resource

	Is_ExampleLipidProfileSpecimen()
}

func (s Specimen) Is_ExampleLipidProfileSpecimen() {}
func (g Group) Is_ExampleLipidProfileSpecimen()    {}

type ExampleLipidProfileDevice interface {
	gofhirclient.Resource

	Is_ExampleLipidProfileDevice()
}

func (d Device) Is_ExampleLipidProfileDevice()       {}
func (d DeviceMetric) Is_ExampleLipidProfileDevice() {}

type ExampleLipidProfileHasMember interface {
	gofhirclient.Resource

	Is_ExampleLipidProfileHasMember()
}

func (o Observation) Is_ExampleLipidProfileHasMember()           {}
func (q QuestionnaireResponse) Is_ExampleLipidProfileHasMember() {}
func (m MolecularSequence) Is_ExampleLipidProfileHasMember()     {}

type ExampleLipidProfileDerivedFrom interface {
	gofhirclient.Resource

	Is_ExampleLipidProfileDerivedFrom()
}

func (d DocumentReference) Is_ExampleLipidProfileDerivedFrom()     {}
func (i ImagingStudy) Is_ExampleLipidProfileDerivedFrom()          {}
func (i ImagingSelection) Is_ExampleLipidProfileDerivedFrom()      {}
func (q QuestionnaireResponse) Is_ExampleLipidProfileDerivedFrom() {}
func (o Observation) Is_ExampleLipidProfileDerivedFrom()           {}
func (m MolecularSequence) Is_ExampleLipidProfileDerivedFrom()     {}
func (g GenomicStudy) Is_ExampleLipidProfileDerivedFrom()          {}

func (e ExampleLipidProfile) ResourceType() string {
	return "ExampleLipidProfile"
}

func (e ExampleLipidProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExampleLipidProfile
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherExampleLipidProfile: OtherExampleLipidProfile(e), ResourceType: e.ResourceType()})
}

func UnmarshalExampleLipidProfile(b []byte) (res ExampleLipidProfile, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
