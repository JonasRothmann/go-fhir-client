// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Observation
type Observation struct {
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
	// Business Identifier for observation
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Instantiates FHIR ObservationDefinition
	Instantiates *custom.Canonical[ObservationDefinition] `bson:"instantiates,omitempty" json:"instantiates,omitempty"`
	// Fulfills plan, proposal or order
	BasedOn []custom.Reference[ObservationBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// Triggering observation(s)
	TriggeredBy []ObservationTriggeredBy `bson:"triggeredBy,omitempty" json:"triggeredBy,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[ObservationPartOf] `bson:"partOf,omitempty" json:"partOf,omitempty"`
	// registered | preliminary | final | amended +
	Status custom.Code `bson:"status" json:"status"`
	// Classification of  type of observation
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Type of observation (code / type)
	Code CodeableConcept `bson:"code" json:"code"`
	// Who and/or what the observation is about
	Subject *custom.Reference[ObservationSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// What the observation is about, when it is not about the subject of record
	Focus []custom.Reference[Resource] `bson:"focus,omitempty" json:"focus,omitempty"`
	// Healthcare event during which this observation is made
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Clinically relevant time/time-period for observation
	Effective *custom.DateTime `bson:"effective,omitempty" json:"effective,omitempty"`
	// Date/Time this version was made available
	Issued *custom.Instant `bson:"issued,omitempty" json:"issued,omitempty"`
	// Who is responsible for the observation
	Performer []custom.Reference[ObservationPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Actual result
	Value *Quantity `bson:"value,omitempty" json:"value,omitempty"`
	// Why the result is missing
	DataAbsentReason *CodeableConcept `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	// High, low, normal, etc
	Interpretation []CodeableConcept `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	// Comments about the observation
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Observed body part
	BodySite *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Observed body structure
	BodyStructure *custom.Reference[BodyStructure] `bson:"bodyStructure,omitempty" json:"bodyStructure,omitempty"`
	// How it was done
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Specimen used for this observation
	Specimen *custom.Reference[ObservationSpecimen] `bson:"specimen,omitempty" json:"specimen,omitempty"`
	// A reference to the device that generates the measurements or the device settings for the device
	Device *custom.Reference[ObservationDevice] `bson:"device,omitempty" json:"device,omitempty"`
	// Provides guide for interpretation
	ReferenceRange []ObservationReferenceRange `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
	// Related resource that belongs to the Observation group
	HasMember []custom.Reference[ObservationHasMember] `bson:"hasMember,omitempty" json:"hasMember,omitempty"`
	// Related resource from which the observation is made
	DerivedFrom []custom.Reference[ObservationDerivedFrom] `bson:"derivedFrom,omitempty" json:"derivedFrom,omitempty"`
	// Component results
	Component []ObservationComponent `bson:"component,omitempty" json:"component,omitempty"`
}

type ObservationTriggeredBy struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Triggering observation
	Observation custom.Reference[Observation] `bson:"observation" json:"observation"`
	// reflex | repeat | re-run
	Type custom.Code `bson:"type" json:"type"`
	// Reason that the observation was triggered
	Reason *string `bson:"reason,omitempty" json:"reason,omitempty"`
}

type ObservationReferenceRange struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Low Range, if relevant
	Low *Quantity `bson:"low,omitempty" json:"low,omitempty"`
	// High Range, if relevant
	High *Quantity `bson:"high,omitempty" json:"high,omitempty"`
	// Normal value, if relevant
	NormalValue *CodeableConcept `bson:"normalValue,omitempty" json:"normalValue,omitempty"`
	// Reference range qualifier
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Reference range population
	AppliesTo []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	// Applicable age range, if relevant
	Age *Range `bson:"age,omitempty" json:"age,omitempty"`
	// Text based reference range in an observation
	Text *custom.Markdown `bson:"text,omitempty" json:"text,omitempty"`
}

type ObservationComponent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of component observation (code / type)
	Code CodeableConcept `bson:"code" json:"code"`
	// Actual component result
	Value *Quantity `bson:"value,omitempty" json:"value,omitempty"`
	// Why the component result is missing
	DataAbsentReason *CodeableConcept `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	// High, low, normal, etc
	Interpretation []CodeableConcept `bson:"interpretation,omitempty" json:"interpretation,omitempty"`
	// Provides guide for interpretation of component result
	ReferenceRange []interface{} `bson:"referenceRange,omitempty" json:"referenceRange,omitempty"`
}

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
	return "StructureDefinition"
}
