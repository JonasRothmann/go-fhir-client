// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ImmunizationEvaluation
type ImmunizationEvaluation struct {
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
	// completed | entered-in-error
	Status custom.Code `json:"status"`
	// Who this evaluation is for
	Patient custom.Reference[Patient] `json:"patient"`
	// Date evaluation was performed
	Date *custom.DateTime `json:"date,omitempty"`
	// Who is responsible for publishing the recommendations
	Authority *custom.Reference[Organization] `json:"authority,omitempty"`
	// The vaccine preventable disease schedule being evaluated
	TargetDisease CodeableConcept `json:"targetDisease"`
	// Immunization being evaluated
	ImmunizationEvent custom.Reference[Immunization] `json:"immunizationEvent"`
	// Status of the dose relative to published recommendations
	DoseStatus CodeableConcept `json:"doseStatus"`
	// Reason why the doese is considered valid, invalid or some other status
	DoseStatusReason []CodeableConcept `json:"doseStatusReason,omitempty"`
	// Evaluation notes
	Description *custom.Markdown `json:"description,omitempty"`
	// Name of vaccine series
	Series *string `json:"series,omitempty"`
	// Dose number within series
	DoseNumber *string `json:"doseNumber,omitempty"`
	// Recommended number of doses for immunity
	SeriesDoses *string `json:"seriesDoses,omitempty"`
}

type OtherImmunizationEvaluation ImmunizationEvaluation

func (i ImmunizationEvaluation) ResourceType() string {
	return "ImmunizationEvaluation"
}

func (i ImmunizationEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationEvaluation
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherImmunizationEvaluation: OtherImmunizationEvaluation(i), ResourceType: i.ResourceType()})
}

func UnmarshalImmunizationEvaluation(b []byte) (res ImmunizationEvaluation, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
