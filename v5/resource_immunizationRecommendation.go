// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendation struct {
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
	// Who this profile is for
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Date recommendation(s) created
	Date custom.DateTime `bson:"date" json:"date"`
	// Who is responsible for protocol
	Authority *custom.Reference[Organization] `bson:"authority,omitempty" json:"authority,omitempty"`
	// Vaccine administration recommendations
	Recommendation []ImmunizationRecommendationRecommendation `bson:"recommendation" json:"recommendation"`
}

type ImmunizationRecommendationRecommendation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Vaccine  or vaccine group recommendation applies to
	VaccineCode []CodeableConcept `bson:"vaccineCode,omitempty" json:"vaccineCode,omitempty"`
	// Disease to be immunized against
	TargetDisease []CodeableConcept `bson:"targetDisease,omitempty" json:"targetDisease,omitempty"`
	// Vaccine which is contraindicated to fulfill the recommendation
	ContraindicatedVaccineCode []CodeableConcept `bson:"contraindicatedVaccineCode,omitempty" json:"contraindicatedVaccineCode,omitempty"`
	// Vaccine recommendation status
	ForecastStatus CodeableConcept `bson:"forecastStatus" json:"forecastStatus"`
	// Vaccine administration status reason
	ForecastReason []CodeableConcept `bson:"forecastReason,omitempty" json:"forecastReason,omitempty"`
	// Dates governing proposed immunization
	DateCriterion []ImmunizationRecommendationRecommendationDateCriterion `bson:"dateCriterion,omitempty" json:"dateCriterion,omitempty"`
	// Protocol details
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Name of vaccination series
	Series *string `bson:"series,omitempty" json:"series,omitempty"`
	// Recommended dose number within series
	DoseNumber *string `bson:"doseNumber,omitempty" json:"doseNumber,omitempty"`
	// Recommended number of doses for immunity
	SeriesDoses *string `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty"`
	// Past immunizations supporting recommendation
	SupportingImmunization []custom.Reference[ImmunizationRecommendationRecommendationSupportingImmunization] `bson:"supportingImmunization,omitempty" json:"supportingImmunization,omitempty"`
	// Patient observations supporting recommendation
	SupportingPatientInformation []custom.Reference[Resource] `bson:"supportingPatientInformation,omitempty" json:"supportingPatientInformation,omitempty"`
}

type ImmunizationRecommendationRecommendationDateCriterion struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of date
	Code CodeableConcept `bson:"code" json:"code"`
	// Recommended date
	Value custom.DateTime `bson:"value" json:"value"`
}

type ImmunizationRecommendationRecommendationSupportingImmunization interface {
	gofhirclient.Resource

	Is_ImmunizationRecommendationRecommendationSupportingImmunization()
}

func (i Immunization) Is_ImmunizationRecommendationRecommendationSupportingImmunization()           {}
func (i ImmunizationEvaluation) Is_ImmunizationRecommendationRecommendationSupportingImmunization() {}

func (i ImmunizationRecommendation) ResourceType() string {
	return "StructureDefinition"
}
