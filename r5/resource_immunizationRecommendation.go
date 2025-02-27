// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ImmunizationRecommendation
type ImmunizationRecommendation struct {
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
	// Who this profile is for
	Patient custom.Reference[Patient] `json:"patient"`
	// Date recommendation(s) created
	Date custom.DateTime `json:"date"`
	// Who is responsible for protocol
	Authority *custom.Reference[Organization] `json:"authority,omitempty"`
	// Vaccine administration recommendations
	Recommendation []ImmunizationRecommendationRecommendation `json:"recommendation"`
}

type ImmunizationRecommendationRecommendation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Vaccine  or vaccine group recommendation applies to
	VaccineCode []CodeableConcept `json:"vaccineCode,omitempty"`
	// Disease to be immunized against
	TargetDisease []CodeableConcept `json:"targetDisease,omitempty"`
	// Vaccine which is contraindicated to fulfill the recommendation
	ContraindicatedVaccineCode []CodeableConcept `json:"contraindicatedVaccineCode,omitempty"`
	// Vaccine recommendation status
	ForecastStatus CodeableConcept `json:"forecastStatus"`
	// Vaccine administration status reason
	ForecastReason []CodeableConcept `json:"forecastReason,omitempty"`
	// Dates governing proposed immunization
	DateCriterion []ImmunizationRecommendationRecommendationDateCriterion `json:"dateCriterion,omitempty"`
	// Protocol details
	Description *custom.Markdown `json:"description,omitempty"`
	// Name of vaccination series
	Series *string `json:"series,omitempty"`
	// Recommended dose number within series
	DoseNumber *string `json:"doseNumber,omitempty"`
	// Recommended number of doses for immunity
	SeriesDoses *string `json:"seriesDoses,omitempty"`
	// Past immunizations supporting recommendation
	SupportingImmunization []custom.Reference[ImmunizationRecommendationRecommendationSupportingImmunization] `json:"supportingImmunization,omitempty"`
	// Patient observations supporting recommendation
	SupportingPatientInformation []custom.Reference[Resource] `json:"supportingPatientInformation,omitempty"`
}

type ImmunizationRecommendationRecommendationDateCriterion struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of date
	Code CodeableConcept `json:"code"`
	// Recommended date
	Value custom.DateTime `json:"value"`
}

type OtherImmunizationRecommendation ImmunizationRecommendation

type ImmunizationRecommendationRecommendationSupportingImmunization interface {
	gofhirclient.Resource

	Is_ImmunizationRecommendationRecommendationSupportingImmunization()
}

func (i Immunization) Is_ImmunizationRecommendationRecommendationSupportingImmunization()           {}
func (i ImmunizationEvaluation) Is_ImmunizationRecommendationRecommendationSupportingImmunization() {}

func (i ImmunizationRecommendation) ResourceType() string {
	return "ImmunizationRecommendation"
}

func (i ImmunizationRecommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationRecommendation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherImmunizationRecommendation: OtherImmunizationRecommendation(i), ResourceType: i.ResourceType()})
}

func UnmarshalImmunizationRecommendation(b []byte) (res ImmunizationRecommendation, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
