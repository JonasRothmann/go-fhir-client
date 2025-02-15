// GENERATED CODE – DO NOT EDIT!

package v5

// http://hl7.org/fhir/StructureDefinition/Dosage
type Dosage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The order of the dosage instructions
	Sequence *int32 `bson:"sequence,omitempty" json:"sequence,omitempty"`
	// Free text dosage instructions e.g. SIG
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// Supplemental instruction or warnings to the patient - e.g. "with meals", "may cause drowsiness"
	AdditionalInstruction []CodeableConcept `bson:"additionalInstruction,omitempty" json:"additionalInstruction,omitempty"`
	// Patient or consumer oriented instructions
	PatientInstruction *string `bson:"patientInstruction,omitempty" json:"patientInstruction,omitempty"`
	// When medication should be administered
	Timing *Timing `bson:"timing,omitempty" json:"timing,omitempty"`
	// Take "as needed"
	AsNeeded *bool `bson:"asNeeded,omitempty" json:"asNeeded,omitempty"`
	// Take "as needed" (for x)
	AsNeededFor []CodeableConcept `bson:"asNeededFor,omitempty" json:"asNeededFor,omitempty"`
	// Body site to administer to
	Site *CodeableConcept `bson:"site,omitempty" json:"site,omitempty"`
	// How drug should enter body
	Route *CodeableConcept `bson:"route,omitempty" json:"route,omitempty"`
	// Technique for administering medication
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Amount of medication administered, to be administered or typical amount to be administered
	DoseAndRate []DosageDoseAndRate `bson:"doseAndRate,omitempty" json:"doseAndRate,omitempty"`
	// Upper limit on medication per unit of time
	MaxDosePerPeriod []Ratio `bson:"maxDosePerPeriod,omitempty" json:"maxDosePerPeriod,omitempty"`
	// Upper limit on medication per administration
	MaxDosePerAdministration *Quantity `bson:"maxDosePerAdministration,omitempty" json:"maxDosePerAdministration,omitempty"`
	// Upper limit on medication per lifetime of the patient
	MaxDosePerLifetime *Quantity `bson:"maxDosePerLifetime,omitempty" json:"maxDosePerLifetime,omitempty"`
}

type DosageDoseAndRate struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// The kind of dose or rate specified
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Amount of medication per dose
	Dose *Range `bson:"dose,omitempty" json:"dose,omitempty"`
	// Amount of medication per unit of time
	Rate *Ratio `bson:"rate,omitempty" json:"rate,omitempty"`
}

func (d Dosage) ResourceType() string {
	return "StructureDefinition"
}
