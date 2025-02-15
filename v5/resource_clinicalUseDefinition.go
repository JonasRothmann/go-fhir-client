// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinitionInteraction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The specific medication, product, food etc. or laboratory test that interacts
	Interactant []ClinicalUseDefinitionInteractionInteractant `bson:"interactant,omitempty" json:"interactant,omitempty"`
	// The type of the interaction e.g. drug-drug interaction, drug-lab test interaction
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The effect of the interaction, for example "reduced gastric absorption of primary medication"
	Effect *custom.CodeableReference[ObservationDefinition] `bson:"effect,omitempty" json:"effect,omitempty"`
	// The incidence of the interaction, e.g. theoretical, observed
	Incidence *CodeableConcept `bson:"incidence,omitempty" json:"incidence,omitempty"`
	// Actions for managing the interaction
	Management []CodeableConcept `bson:"management,omitempty" json:"management,omitempty"`
}

type ClinicalUseDefinitionInteractionInteractant struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The specific medication, product, food etc. or laboratory test that interacts
	Item custom.Reference[ClinicalUseDefinitionInteractionInteractantItem] `bson:"item" json:"item"`
}

type ClinicalUseDefinitionUndesirableEffect struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The situation in which the undesirable effect may manifest
	SymptomConditionEffect *custom.CodeableReference[ObservationDefinition] `bson:"symptomConditionEffect,omitempty" json:"symptomConditionEffect,omitempty"`
	// High level classification of the effect
	Classification *CodeableConcept `bson:"classification,omitempty" json:"classification,omitempty"`
	// How often the effect is seen
	FrequencyOfOccurrence *CodeableConcept `bson:"frequencyOfOccurrence,omitempty" json:"frequencyOfOccurrence,omitempty"`
}

type ClinicalUseDefinitionWarning struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A textual definition of this warning, with formatting
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// A coded or unformatted textual definition of this warning
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
}

type ClinicalUseDefinition struct {
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
	// Business identifier for this issue
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// indication | contraindication | interaction | undesirable-effect | warning
	Type custom.Code `bson:"type" json:"type"`
	// A categorisation of the issue, primarily for dividing warnings into subject heading areas such as "Pregnancy", "Overdose"
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// The medication, product, substance, device, procedure etc. for which this is an indication
	Subject []custom.Reference[ClinicalUseDefinitionSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Whether this is a current issue or one that has been retired etc
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// Specifics for when this is a contraindication
	Contraindication *ClinicalUseDefinitionContraindication `bson:"contraindication,omitempty" json:"contraindication,omitempty"`
	// Specifics for when this is an indication
	Indication *ClinicalUseDefinitionIndication `bson:"indication,omitempty" json:"indication,omitempty"`
	// Specifics for when this is an interaction
	Interaction *ClinicalUseDefinitionInteraction `bson:"interaction,omitempty" json:"interaction,omitempty"`
	// The population group to which this applies
	Population []custom.Reference[Group] `bson:"population,omitempty" json:"population,omitempty"`
	// Logic used by the clinical use definition
	Library []custom.Canonical[Library] `bson:"library,omitempty" json:"library,omitempty"`
	// A possible negative outcome from the use of this treatment
	UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `bson:"undesirableEffect,omitempty" json:"undesirableEffect,omitempty"`
	// Critical environmental, health or physical risks or hazards. For example 'Do not operate heavy machinery', 'May cause drowsiness'
	Warning *ClinicalUseDefinitionWarning `bson:"warning,omitempty" json:"warning,omitempty"`
}

type ClinicalUseDefinitionContraindication struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The situation that is being documented as contraindicating against this item
	DiseaseSymptomProcedure *custom.CodeableReference[ObservationDefinition] `bson:"diseaseSymptomProcedure,omitempty" json:"diseaseSymptomProcedure,omitempty"`
	// The status of the disease or symptom for the contraindication
	DiseaseStatus *custom.CodeableReference[ObservationDefinition] `bson:"diseaseStatus,omitempty" json:"diseaseStatus,omitempty"`
	// A comorbidity (concurrent condition) or coinfection
	Comorbidity []custom.CodeableReference[ObservationDefinition] `bson:"comorbidity,omitempty" json:"comorbidity,omitempty"`
	// The indication which this is a contraidication for
	Indication []custom.Reference[ClinicalUseDefinition] `bson:"indication,omitempty" json:"indication,omitempty"`
	// An expression that returns true or false, indicating whether the indication is applicable or not, after having applied its other elements
	Applicability *Expression `bson:"applicability,omitempty" json:"applicability,omitempty"`
	// Information about use of the product in relation to other therapies described as part of the contraindication
	OtherTherapy []ClinicalUseDefinitionContraindicationOtherTherapy `bson:"otherTherapy,omitempty" json:"otherTherapy,omitempty"`
}

type ClinicalUseDefinitionContraindicationOtherTherapy struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of relationship between the product indication/contraindication and another therapy
	RelationshipType CodeableConcept `bson:"relationshipType" json:"relationshipType"`
	// Reference to a specific medication, substance etc. as part of an indication or contraindication
	Treatment custom.CodeableReference[ClinicalUseDefinitionContraindicationOtherTherapyTreatment] `bson:"treatment" json:"treatment"`
}

type ClinicalUseDefinitionIndication struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The situation that is being documented as an indicaton for this item
	DiseaseSymptomProcedure *custom.CodeableReference[ObservationDefinition] `bson:"diseaseSymptomProcedure,omitempty" json:"diseaseSymptomProcedure,omitempty"`
	// The status of the disease or symptom for the indication
	DiseaseStatus *custom.CodeableReference[ObservationDefinition] `bson:"diseaseStatus,omitempty" json:"diseaseStatus,omitempty"`
	// A comorbidity or coinfection as part of the indication
	Comorbidity []custom.CodeableReference[ObservationDefinition] `bson:"comorbidity,omitempty" json:"comorbidity,omitempty"`
	// The intended effect, aim or strategy to be achieved
	IntendedEffect *custom.CodeableReference[ObservationDefinition] `bson:"intendedEffect,omitempty" json:"intendedEffect,omitempty"`
	// Timing or duration information
	Duration *Range `bson:"duration,omitempty" json:"duration,omitempty"`
	// An unwanted side effect or negative outcome of the subject of this resource when being used for this indication
	UndesirableEffect []custom.Reference[ClinicalUseDefinition] `bson:"undesirableEffect,omitempty" json:"undesirableEffect,omitempty"`
	// An expression that returns true or false, indicating whether the indication is applicable or not, after having applied its other elements
	Applicability *Expression `bson:"applicability,omitempty" json:"applicability,omitempty"`
	// The use of the medicinal product in relation to other therapies described as part of the indication
	OtherTherapy []interface{} `bson:"otherTherapy,omitempty" json:"otherTherapy,omitempty"`
}

type ClinicalUseDefinitionSubject interface {
	gofhirclient.Resource

	Is_ClinicalUseDefinitionSubject()
}

func (m MedicinalProductDefinition) Is_ClinicalUseDefinitionSubject() {}
func (m Medication) Is_ClinicalUseDefinitionSubject()                 {}
func (a ActivityDefinition) Is_ClinicalUseDefinitionSubject()         {}
func (p PlanDefinition) Is_ClinicalUseDefinitionSubject()             {}
func (d Device) Is_ClinicalUseDefinitionSubject()                     {}
func (d DeviceDefinition) Is_ClinicalUseDefinitionSubject()           {}
func (s Substance) Is_ClinicalUseDefinitionSubject()                  {}
func (n NutritionProduct) Is_ClinicalUseDefinitionSubject()           {}
func (b BiologicallyDerivedProduct) Is_ClinicalUseDefinitionSubject() {}

type ClinicalUseDefinitionContraindicationOtherTherapyTreatment interface {
	gofhirclient.Resource

	Is_ClinicalUseDefinitionContraindicationOtherTherapyTreatment()
}

func (m MedicinalProductDefinition) Is_ClinicalUseDefinitionContraindicationOtherTherapyTreatment() {}
func (m Medication) Is_ClinicalUseDefinitionContraindicationOtherTherapyTreatment()                 {}
func (s Substance) Is_ClinicalUseDefinitionContraindicationOtherTherapyTreatment()                  {}
func (s SubstanceDefinition) Is_ClinicalUseDefinitionContraindicationOtherTherapyTreatment()        {}
func (n NutritionProduct) Is_ClinicalUseDefinitionContraindicationOtherTherapyTreatment()           {}
func (b BiologicallyDerivedProduct) Is_ClinicalUseDefinitionContraindicationOtherTherapyTreatment() {}
func (a ActivityDefinition) Is_ClinicalUseDefinitionContraindicationOtherTherapyTreatment()         {}

type ClinicalUseDefinitionInteractionInteractantItem interface {
	gofhirclient.Resource

	Is_ClinicalUseDefinitionInteractionInteractantItem()
}

func (m MedicinalProductDefinition) Is_ClinicalUseDefinitionInteractionInteractantItem() {}
func (m Medication) Is_ClinicalUseDefinitionInteractionInteractantItem()                 {}
func (s Substance) Is_ClinicalUseDefinitionInteractionInteractantItem()                  {}
func (n NutritionProduct) Is_ClinicalUseDefinitionInteractionInteractantItem()           {}
func (b BiologicallyDerivedProduct) Is_ClinicalUseDefinitionInteractionInteractantItem() {}
func (o ObservationDefinition) Is_ClinicalUseDefinitionInteractionInteractantItem()      {}

func (c ClinicalUseDefinition) ResourceType() string {
	return "StructureDefinition"
}
