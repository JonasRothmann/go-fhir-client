// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Immunization
type Immunization struct {
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
	// Business identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Authority that the immunization event is based on
	BasedOn []custom.Reference[ImmunizationBasedOn] `json:"basedOn,omitempty"`
	// completed | entered-in-error | not-done
	Status custom.Code `json:"status"`
	// Reason for current status
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	// Vaccine administered
	VaccineCode CodeableConcept `json:"vaccineCode"`
	// Product that was administered
	AdministeredProduct *custom.CodeableReference[Medication] `json:"administeredProduct,omitempty"`
	// Vaccine manufacturer
	Manufacturer *custom.CodeableReference[Organization] `json:"manufacturer,omitempty"`
	// Vaccine lot number
	LotNumber *string `json:"lotNumber,omitempty"`
	// Vaccine expiration date
	ExpirationDate *custom.Date `json:"expirationDate,omitempty"`
	// Who was immunized
	Patient custom.Reference[Patient] `json:"patient"`
	// Encounter immunization was part of
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Additional information in support of the immunization
	SupportingInformation []custom.Reference[Resource] `json:"supportingInformation,omitempty"`
	// Vaccine administration date
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// Vaccine administration date
	OccurrenceString *string `json:"occurrenceString,omitempty"`
	// Indicates context the data was captured in
	PrimarySource *bool `json:"primarySource,omitempty"`
	// Indicates the source of a  reported record
	InformationSource *custom.CodeableReference[ImmunizationInformationSource] `json:"informationSource,omitempty"`
	// Where immunization occurred
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Body site vaccine  was administered
	Site *CodeableConcept `json:"site,omitempty"`
	// How vaccine entered body
	Route *CodeableConcept `json:"route,omitempty"`
	// Amount of vaccine administered
	DoseQuantity *Quantity `json:"doseQuantity,omitempty"`
	// Who performed event
	Performer []ImmunizationPerformer `json:"performer,omitempty"`
	// Additional immunization notes
	Note []Annotation `json:"note,omitempty"`
	// Why immunization occurred
	Reason []custom.CodeableReference[ImmunizationReason] `json:"reason,omitempty"`
	// Dose potency
	IsSubpotent *bool `json:"isSubpotent,omitempty"`
	// Reason for being subpotent
	SubpotentReason []CodeableConcept `json:"subpotentReason,omitempty"`
	// Patient eligibility for a specific vaccination program
	ProgramEligibility []ImmunizationProgramEligibility `json:"programEligibility,omitempty"`
	// Funding source for the vaccine
	FundingSource *CodeableConcept `json:"fundingSource,omitempty"`
	// Details of a reaction that follows immunization
	Reaction []ImmunizationReaction `json:"reaction,omitempty"`
	// Protocol followed by the provider
	ProtocolApplied []ImmunizationProtocolApplied `json:"protocolApplied,omitempty"`
}

type ImmunizationPerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// What type of performance was done
	Function *CodeableConcept `json:"function,omitempty"`
	// Individual or organization who was performing
	Actor custom.Reference[ImmunizationPerformerActor] `json:"actor"`
}

type ImmunizationProgramEligibility struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The program that eligibility is declared for
	Program CodeableConcept `json:"program"`
	// The patient's eligibility status for the program
	ProgramStatus CodeableConcept `json:"programStatus"`
}

type ImmunizationReaction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// When reaction started
	Date *custom.DateTime `json:"date,omitempty"`
	// Additional information on reaction
	Manifestation *custom.CodeableReference[Observation] `json:"manifestation,omitempty"`
	// Indicates self-reported reaction
	Reported *bool `json:"reported,omitempty"`
}

type ImmunizationProtocolApplied struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name of vaccine series
	Series *string `json:"series,omitempty"`
	// Who is responsible for publishing the recommendations
	Authority *custom.Reference[Organization] `json:"authority,omitempty"`
	// Vaccine preventatable disease being targeted
	TargetDisease []CodeableConcept `json:"targetDisease,omitempty"`
	// Dose number within series
	DoseNumber string `json:"doseNumber"`
	// Recommended number of doses for immunity
	SeriesDoses *string `json:"seriesDoses,omitempty"`
}

type OtherImmunization Immunization

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
	return "Immunization"
}

func (i Immunization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunization
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherImmunization: OtherImmunization(i), ResourceType: i.ResourceType()})
}

func UnmarshalImmunization(b []byte) (res Immunization, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
