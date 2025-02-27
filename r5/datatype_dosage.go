// GENERATED CODE – DO NOT EDIT!

package r5

import "encoding/json"

// http://hl7.org/fhir/StructureDefinition/Dosage
type Dosage struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The order of the dosage instructions
	Sequence *int32 `json:"sequence,omitempty"`
	// Free text dosage instructions e.g. SIG
	Text *string `json:"text,omitempty"`
	// Supplemental instruction or warnings to the patient - e.g. "with meals", "may cause drowsiness"
	AdditionalInstruction []CodeableConcept `json:"additionalInstruction,omitempty"`
	// Patient or consumer oriented instructions
	PatientInstruction *string `json:"patientInstruction,omitempty"`
	// When medication should be administered
	Timing *Timing `json:"timing,omitempty"`
	// Take "as needed"
	AsNeeded *bool `json:"asNeeded,omitempty"`
	// Take "as needed" (for x)
	AsNeededFor []CodeableConcept `json:"asNeededFor,omitempty"`
	// Body site to administer to
	Site *CodeableConcept `json:"site,omitempty"`
	// How drug should enter body
	Route *CodeableConcept `json:"route,omitempty"`
	// Technique for administering medication
	Method *CodeableConcept `json:"method,omitempty"`
	// Amount of medication administered, to be administered or typical amount to be administered
	DoseAndRate []DosageDoseAndRate `json:"doseAndRate,omitempty"`
	// Upper limit on medication per unit of time
	MaxDosePerPeriod []Ratio `json:"maxDosePerPeriod,omitempty"`
	// Upper limit on medication per administration
	MaxDosePerAdministration *Quantity `json:"maxDosePerAdministration,omitempty"`
	// Upper limit on medication per lifetime of the patient
	MaxDosePerLifetime *Quantity `json:"maxDosePerLifetime,omitempty"`
}

type DosageDoseAndRate struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// The kind of dose or rate specified
	Type *CodeableConcept `json:"type,omitempty"`
	// Amount of medication per dose
	DoseRange *Range `json:"doseRange,omitempty"`
	// Amount of medication per dose
	DoseQuantity *Quantity `json:"doseQuantity,omitempty"`
	// Amount of medication per unit of time
	RateRatio *Ratio `json:"rateRatio,omitempty"`
	// Amount of medication per unit of time
	RateRange *Range `json:"rateRange,omitempty"`
	// Amount of medication per unit of time
	RateQuantity *Quantity `json:"rateQuantity,omitempty"`
}

type OtherDosage Dosage

func (d Dosage) ResourceType() string {
	return "Dosage"
}

func (d Dosage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDosage
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherDosage: OtherDosage(d), ResourceType: d.ResourceType()})
}

func UnmarshalDosage(b []byte) (res Dosage, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
