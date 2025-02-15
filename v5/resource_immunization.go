// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Immunization
type ImmunizationReaction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// When reaction started
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Additional information on reaction
	Manifestation *custom.CodeableReference[Observation] `bson:"manifestation,omitempty" json:"manifestation,omitempty"`
	// Indicates self-reported reaction
	Reported *bool `bson:"reported,omitempty" json:"reported,omitempty"`
}

type ImmunizationProtocolApplied struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Name of vaccine series
	Series *string `bson:"series,omitempty" json:"series,omitempty"`
	// Who is responsible for publishing the recommendations
	Authority *custom.Reference[Organization] `bson:"authority,omitempty" json:"authority,omitempty"`
	// Vaccine preventatable disease being targeted
	TargetDisease []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	// Dose number within series
	DoseNumber string `bson:"doseNumber" json:"doseNumber"`
	// Recommended number of doses for immunity
	SeriesDoses *string `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty"`
}

type Immunization struct {
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
	// Business identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Authority that the immunization event is based on
	BasedOn []custom.Reference[ImmunizationBasedOn] `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	// completed | entered-in-error | not-done
	Status custom.Code `bson:"status" json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `bson:"statusReason,omitempty" json:"statusReason,omitempty"`
	// Vaccine administered
	VaccineCode CodeableConcept `bson:"vaccineCode" json:"vaccineCode"`
	// Product that was administered
	AdministeredProduct *custom.CodeableReference[Medication] `bson:"administeredProduct,omitempty" json:"administeredProduct,omitempty"`
	// Vaccine manufacturer
	Manufacturer *custom.CodeableReference[Organization] `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// Vaccine lot number
	LotNumber *string `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	// Vaccine expiration date
	ExpirationDate *custom.Date `bson:"expirationDate,omitempty" json:"expirationDate,omitempty"`
	// Who was immunized
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Encounter immunization was part of
	Encounter *custom.Reference[Encounter] `bson:"encounter,omitempty" json:"encounter,omitempty"`
	// Additional information in support of the immunization
	SupportingInformation []custom.Reference[Resource] `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
	// Vaccine administration date
	Occurrence custom.DateTime `bson:"occurrence" json:"occurrence"`
	// Indicates context the data was captured in
	PrimarySource *bool `bson:"primarySource,omitempty" json:"primarySource,omitempty"`
	// Indicates the source of a  reported record
	InformationSource *custom.CodeableReference[ImmunizationInformationSource] `bson:"informationSource,omitempty" json:"informationSource,omitempty"`
	// Where immunization occurred
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// Body site vaccine  was administered
	Site *CodeableConcept `bson:"site,omitempty" json:"site,omitempty"`
	// How vaccine entered body
	Route *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	// Amount of vaccine administered
	DoseQuantity *Quantity `bson:"doseQuantity,omitempty" json:"doseQuantity,omitempty"`
	// Who performed event
	Performer []ImmunizationPerformer `bson:"performer,omitempty" json:"performer,omitempty"`
	// Additional immunization notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Why immunization occurred
	Reason []custom.CodeableReference[ImmunizationReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Dose potency
	IsSubpotent *bool `bson:"isSubpotent,omitempty" json:"isSubpotent,omitempty"`
	// Reason for being subpotent
	SubpotentReason []CodeableConcept `bson:"subpotentReason,omitempty" json:"subpotentReason,omitempty"`
	// Patient eligibility for a specific vaccination program
	ProgramEligibility []ImmunizationProgramEligibility `bson:"programEligibility,omitempty" json:"programEligibility,omitempty"`
	// Funding source for the vaccine
	FundingSource *CodeableConcept `bson:"fundingSource,omitempty" json:"fundingSource,omitempty"`
	// Details of a reaction that follows immunization
	Reaction []ImmunizationReaction `bson:"reaction,omitempty" json:"reaction,omitempty"`
	// Protocol followed by the provider
	ProtocolApplied []ImmunizationProtocolApplied `bson:"protocolApplied,omitempty" json:"protocolApplied,omitempty"`
}

type ImmunizationPerformer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// What type of performance was done
	Function *CodeableConcept `bson:"function,omitempty" json:"function,omitempty"`
	// Individual or organization who was performing
	Actor custom.Reference[ImmunizationPerformerActor] `bson:"actor" json:"actor"`
}

type ImmunizationProgramEligibility struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The program that eligibility is declared for
	Program CodeableConcept `bson:"program" json:"program"`
	// The patient's eligibility status for the program
	ProgramStatus CodeableConcept `bson:"programStatus" json:"programStatus"`
}

type ImmunizationBasedOn interface {
	gofhirclient.Resource

	Is_ImmunizationBasedOn()
}

func (c CarePlan) Is_ImmunizationBasedOn()                   {}
func (m MedicationRequest) Is_ImmunizationBasedOn()          {}
func (s ServiceRequest) Is_ImmunizationBasedOn()             {}
func (i ImmunizationRecommendation) Is_ImmunizationBasedOn() {}

type ImmunizationInformationSource interface {
	gofhirclient.Resource

	Is_ImmunizationInformationSource()
}

func (p Patient) Is_ImmunizationInformationSource()          {}
func (p Practitioner) Is_ImmunizationInformationSource()     {}
func (p PractitionerRole) Is_ImmunizationInformationSource() {}
func (r RelatedPerson) Is_ImmunizationInformationSource()    {}
func (o Organization) Is_ImmunizationInformationSource()     {}

type ImmunizationPerformerActor interface {
	gofhirclient.Resource

	Is_ImmunizationPerformerActor()
}

func (p Practitioner) Is_ImmunizationPerformerActor()     {}
func (p PractitionerRole) Is_ImmunizationPerformerActor() {}
func (o Organization) Is_ImmunizationPerformerActor()     {}
func (p Patient) Is_ImmunizationPerformerActor()          {}
func (r RelatedPerson) Is_ImmunizationPerformerActor()    {}

type ImmunizationReason interface {
	gofhirclient.Resource

	Is_ImmunizationReason()
}

func (c Condition) Is_ImmunizationReason()        {}
func (o Observation) Is_ImmunizationReason()      {}
func (d DiagnosticReport) Is_ImmunizationReason() {}

func (i Immunization) ResourceType() string {
	return "StructureDefinition"
}
