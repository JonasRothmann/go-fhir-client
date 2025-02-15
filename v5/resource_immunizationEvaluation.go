// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/ImmunizationEvaluation
type ImmunizationEvaluation struct {
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
	// completed | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// Who this evaluation is for
	Patient custom.Reference[Patient] `bson:"patient" json:"patient"`
	// Date evaluation was performed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// Who is responsible for publishing the recommendations
	Authority *custom.Reference[Organization] `bson:"authority,omitempty" json:"authority,omitempty"`
	// The vaccine preventable disease schedule being evaluated
	TargetDisease CodeableConcept `bson:"targetDisease" json:"targetDisease"`
	// Immunization being evaluated
	ImmunizationEvent custom.Reference[Immunization] `bson:"immunizationEvent" json:"immunizationEvent"`
	// Status of the dose relative to published recommendations
	DoseStatus CodeableConcept `bson:"doseStatus" json:"doseStatus"`
	// Reason why the doese is considered valid, invalid or some other status
	DoseStatusReason []CodeableConcept `bson:"doseStatusReason,omitempty" json:"doseStatusReason,omitempty"`
	// Evaluation notes
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Name of vaccine series
	Series *string `bson:"series,omitempty" json:"series,omitempty"`
	// Dose number within series
	DoseNumber *string `bson:"doseNumber,omitempty" json:"doseNumber,omitempty"`
	// Recommended number of doses for immunity
	SeriesDoses *string `bson:"seriesDoses,omitempty" json:"seriesDoses,omitempty"`
}

func (i ImmunizationEvaluation) ResourceType() string {
	return "StructureDefinition"
}
