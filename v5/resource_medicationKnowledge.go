// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/MedicationKnowledge
type MedicationKnowledgeCost struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The date range for which the cost is effective
	EffectiveDate []Period `bson:"effectiveDate,omitempty" json:"effectiveDate,omitempty"`
	// The category of the cost information
	Type CodeableConcept `bson:"type" json:"type"`
	// The source or owner for the price information
	Source *string `bson:"source,omitempty" json:"source,omitempty"`
	// The price or category of the cost of the medication
	Cost Money `bson:"cost" json:"cost"`
}

type MedicationKnowledgeStorageGuideline struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference to additional information
	Reference *custom.URI `bson:"reference,omitempty" json:"reference,omitempty"`
	// Additional storage notes
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Duration remains stable
	StabilityDuration *Duration `bson:"stabilityDuration,omitempty" json:"stabilityDuration,omitempty"`
	// Setting or value of environment for adequate storage
	EnvironmentalSetting []MedicationKnowledgeStorageGuidelineEnvironmentalSetting `bson:"environmentalSetting,omitempty" json:"environmentalSetting,omitempty"`
}

type MedicationKnowledgeStorageGuidelineEnvironmentalSetting struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Categorization of the setting
	Type CodeableConcept `bson:"type" json:"type"`
	// Value of the setting
	Value Quantity `bson:"value" json:"value"`
}

type MedicationKnowledgeRegulatorySubstitution struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Specifies the type of substitution allowed
	Type CodeableConcept `bson:"type" json:"type"`
	// Specifies if regulation allows for changes in the medication when dispensing
	Allowed bool `bson:"allowed" json:"allowed"`
}

type MedicationKnowledge struct {
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
	// Business identifier for this medication
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Code that identifies this medication
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// active | entered-in-error | inactive
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Creator or owner of the knowledge or information about the medication
	Author *custom.Reference[Organization] `bson:"author,omitempty" json:"author,omitempty"`
	// Codes that identify the different jurisdictions for which the information of this resource was created
	IntendedJurisdiction []CodeableConcept `bson:"intendedJurisdiction,omitempty" json:"intendedJurisdiction,omitempty"`
	// A name associated with the medication being described
	Name []string `bson:"name,omitempty" json:"name,omitempty"`
	// Associated or related medication information
	RelatedMedicationKnowledge []MedicationKnowledgeRelatedMedicationKnowledge `bson:"relatedMedicationKnowledge,omitempty" json:"relatedMedicationKnowledge,omitempty"`
	// The set of medication resources that are associated with this medication
	AssociatedMedication []custom.Reference[Medication] `bson:"associatedMedication,omitempty" json:"associatedMedication,omitempty"`
	// Category of the medication or product
	ProductType []CodeableConcept `bson:"productType,omitempty" json:"productType,omitempty"`
	// Associated documentation about the medication
	Monograph []MedicationKnowledgeMonograph `bson:"monograph,omitempty" json:"monograph,omitempty"`
	// The instructions for preparing the medication
	PreparationInstruction *custom.Markdown `bson:"preparationInstruction,omitempty" json:"preparationInstruction,omitempty"`
	// The pricing of the medication
	Cost []MedicationKnowledgeCost `bson:"cost,omitempty" json:"cost,omitempty"`
	// Program under which a medication is reviewed
	MonitoringProgram []MedicationKnowledgeMonitoringProgram `bson:"monitoringProgram,omitempty" json:"monitoringProgram,omitempty"`
	// Guidelines or protocols for administration of the medication for an indication
	IndicationGuideline []MedicationKnowledgeIndicationGuideline `bson:"indicationGuideline,omitempty" json:"indicationGuideline,omitempty"`
	// Categorization of the medication within a formulary or classification system
	MedicineClassification []MedicationKnowledgeMedicineClassification `bson:"medicineClassification,omitempty" json:"medicineClassification,omitempty"`
	// Details about packaged medications
	Packaging []MedicationKnowledgePackaging `bson:"packaging,omitempty" json:"packaging,omitempty"`
	// Potential clinical issue with or between medication(s)
	ClinicalUseIssue []custom.Reference[ClinicalUseDefinition] `bson:"clinicalUseIssue,omitempty" json:"clinicalUseIssue,omitempty"`
	// How the medication should be stored
	StorageGuideline []MedicationKnowledgeStorageGuideline `bson:"storageGuideline,omitempty" json:"storageGuideline,omitempty"`
	// Regulatory information about a medication
	Regulatory []MedicationKnowledgeRegulatory `bson:"regulatory,omitempty" json:"regulatory,omitempty"`
	// Minimal definition information about the medication
	Definitional *MedicationKnowledgeDefinitional `bson:"definitional,omitempty" json:"definitional,omitempty"`
}

type MedicationKnowledgeMedicineClassification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of category for the medication (for example, therapeutic classification, therapeutic sub-classification)
	Type CodeableConcept `bson:"type" json:"type"`
	// The source of the classification
	Source *string `bson:"source,omitempty" json:"source,omitempty"`
	// Specific category assigned to the medication
	Classification []CodeableConcept `bson:"classification,omitempty" json:"classification,omitempty"`
}

type MedicationKnowledgeRegulatory struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Specifies the authority of the regulation
	RegulatoryAuthority custom.Reference[Organization] `bson:"regulatoryAuthority" json:"regulatoryAuthority"`
	// Specifies if changes are allowed when dispensing a medication from a regulatory perspective
	Substitution []MedicationKnowledgeRegulatorySubstitution `bson:"substitution,omitempty" json:"substitution,omitempty"`
	// Specifies the schedule of a medication in jurisdiction
	Schedule []CodeableConcept `bson:"schedule,omitempty" json:"schedule,omitempty"`
	// The maximum number of units of the medication that can be dispensed in a period
	MaxDispense *MedicationKnowledgeRegulatoryMaxDispense `bson:"maxDispense,omitempty" json:"maxDispense,omitempty"`
}

type MedicationKnowledgeDefinitionalDrugCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Code specifying the type of characteristic of medication
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Description of the characteristic
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

type MedicationKnowledgeIndicationGuidelineDosingGuideline struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Intention of the treatment
	TreatmentIntent *CodeableConcept `bson:"treatmentIntent,omitempty" json:"treatmentIntent,omitempty"`
	// Dosage for the medication for the specific guidelines
	Dosage []MedicationKnowledgeIndicationGuidelineDosingGuidelineDosage `bson:"dosage,omitempty" json:"dosage,omitempty"`
	// Type of treatment the guideline applies to
	AdministrationTreatment *CodeableConcept `bson:"administrationTreatment,omitempty" json:"administrationTreatment,omitempty"`
	// Characteristics of the patient that are relevant to the administration guidelines
	PatientCharacteristic []MedicationKnowledgeIndicationGuidelineDosingGuidelinePatientCharacteristic `bson:"patientCharacteristic,omitempty" json:"patientCharacteristic,omitempty"`
}

type MedicationKnowledgeRegulatoryMaxDispense struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The maximum number of units of the medication that can be dispensed
	Quantity Quantity `bson:"quantity" json:"quantity"`
	// The period that applies to the maximum number of units
	Period *Duration `bson:"period,omitempty" json:"period,omitempty"`
}

type MedicationKnowledgeDefinitionalIngredient struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Substances contained in the medication
	Item custom.CodeableReference[Substance] `bson:"item" json:"item"`
	// A code that defines the type of ingredient, active, base, etc
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Quantity of ingredient present
	Strength *Ratio `bson:"strength,omitempty" json:"strength,omitempty"`
}

type MedicationKnowledgeRelatedMedicationKnowledge struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Category of medicationKnowledge
	Type CodeableConcept `bson:"type" json:"type"`
	// Associated documentation about the associated medication knowledge
	Reference []custom.Reference[MedicationKnowledge] `bson:"reference" json:"reference"`
}

type MedicationKnowledgeMonitoringProgram struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of program under which the medication is monitored
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Name of the reviewing program
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
}

type MedicationKnowledgeIndicationGuideline struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Indication for use that applies to the specific administration guideline
	Indication []custom.CodeableReference[ClinicalUseDefinition] `bson:"indication,omitempty" json:"indication,omitempty"`
	// Guidelines for dosage of the medication
	DosingGuideline []MedicationKnowledgeIndicationGuidelineDosingGuideline `bson:"dosingGuideline,omitempty" json:"dosingGuideline,omitempty"`
}

type MedicationKnowledgeIndicationGuidelineDosingGuidelineDosage struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Category of dosage for a medication
	Type CodeableConcept `bson:"type" json:"type"`
	// Dosage for the medication for the specific guidelines
	Dosage []Dosage `bson:"dosage" json:"dosage"`
}

type MedicationKnowledgeIndicationGuidelineDosingGuidelinePatientCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Categorization of specific characteristic that is relevant to the administration guideline
	Type CodeableConcept `bson:"type" json:"type"`
	// The specific characteristic
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
}

type MedicationKnowledgePackaging struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Cost of the packaged medication
	Cost []interface{} `bson:"cost,omitempty" json:"cost,omitempty"`
	// The packaged medication that is being priced
	PackagedProduct *custom.Reference[PackagedProductDefinition] `bson:"packagedProduct,omitempty" json:"packagedProduct,omitempty"`
}

type MedicationKnowledgeDefinitional struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Definitional resources that provide more information about this medication
	Definition []custom.Reference[MedicinalProductDefinition] `bson:"definition,omitempty" json:"definition,omitempty"`
	// powder | tablets | capsule +
	DoseForm *CodeableConcept `bson:"doseForm,omitempty" json:"doseForm,omitempty"`
	// The intended or approved route of administration
	IntendedRoute []CodeableConcept `bson:"intendedRoute,omitempty" json:"intendedRoute,omitempty"`
	// Active or inactive ingredient
	Ingredient []MedicationKnowledgeDefinitionalIngredient `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	// Specifies descriptive properties of the medicine
	DrugCharacteristic []MedicationKnowledgeDefinitionalDrugCharacteristic `bson:"drugCharacteristic,omitempty" json:"drugCharacteristic,omitempty"`
}

type MedicationKnowledgeMonograph struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The category of medication document
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Associated documentation about the medication
	Source *custom.Reference[DocumentReference] `bson:"source,omitempty" json:"source,omitempty"`
}

func (m MedicationKnowledge) ResourceType() string {
	return "StructureDefinition"
}
