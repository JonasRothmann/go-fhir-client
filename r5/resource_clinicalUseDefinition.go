// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ClinicalUseDefinition
type ClinicalUseDefinition struct {
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
	// Business identifier for this issue
	Identifier []Identifier `json:"identifier,omitempty"`
	// indication | contraindication | interaction | undesirable-effect | warning
	Type custom.Code `json:"type"`
	// A categorisation of the issue, primarily for dividing warnings into subject heading areas such as "Pregnancy", "Overdose"
	Category []CodeableConcept `json:"category,omitempty"`
	// The medication, product, substance, device, procedure etc. for which this is an indication
	Subject []custom.Reference[ClinicalUseDefinitionSubject] `json:"subject,omitempty"`
	// Whether this is a current issue or one that has been retired etc
	Status *CodeableConcept `json:"status,omitempty"`
	// Specifics for when this is a contraindication
	Contraindication *ClinicalUseDefinitionContraindication `json:"contraindication,omitempty"`
	// Specifics for when this is an indication
	Indication *ClinicalUseDefinitionIndication `json:"indication,omitempty"`
	// Specifics for when this is an interaction
	Interaction *ClinicalUseDefinitionInteraction `json:"interaction,omitempty"`
	// The population group to which this applies
	Population []custom.Reference[Group] `json:"population,omitempty"`
	// Logic used by the clinical use definition
	Library []custom.Canonical[Library] `json:"library,omitempty"`
	// A possible negative outcome from the use of this treatment
	UndesirableEffect *ClinicalUseDefinitionUndesirableEffect `json:"undesirableEffect,omitempty"`
	// Critical environmental, health or physical risks or hazards. For example 'Do not operate heavy machinery', 'May cause drowsiness'
	Warning *ClinicalUseDefinitionWarning `json:"warning,omitempty"`
}

type ClinicalUseDefinitionContraindication struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The situation that is being documented as contraindicating against this item
	DiseaseSymptomProcedure *custom.CodeableReference[ObservationDefinition] `json:"diseaseSymptomProcedure,omitempty"`
	// The status of the disease or symptom for the contraindication
	DiseaseStatus *custom.CodeableReference[ObservationDefinition] `json:"diseaseStatus,omitempty"`
	// A comorbidity (concurrent condition) or coinfection
	Comorbidity []custom.CodeableReference[ObservationDefinition] `json:"comorbidity,omitempty"`
	// The indication which this is a contraidication for
	Indication []custom.Reference[ClinicalUseDefinition] `json:"indication,omitempty"`
	// An expression that returns true or false, indicating whether the indication is applicable or not, after having applied its other elements
	Applicability *Expression `json:"applicability,omitempty"`
	// Information about use of the product in relation to other therapies described as part of the contraindication
	OtherTherapy []ClinicalUseDefinitionContraindicationOtherTherapy `json:"otherTherapy,omitempty"`
}

type ClinicalUseDefinitionContraindicationOtherTherapy struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of relationship between the product indication/contraindication and another therapy
	RelationshipType CodeableConcept `json:"relationshipType"`
	// Reference to a specific medication, substance etc. as part of an indication or contraindication
	Treatment custom.CodeableReference[ClinicalUseDefinitionContraindicationOtherTherapyTreatment] `json:"treatment"`
}

type ClinicalUseDefinitionIndication struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The situation that is being documented as an indicaton for this item
	DiseaseSymptomProcedure *custom.CodeableReference[ObservationDefinition] `json:"diseaseSymptomProcedure,omitempty"`
	// The status of the disease or symptom for the indication
	DiseaseStatus *custom.CodeableReference[ObservationDefinition] `json:"diseaseStatus,omitempty"`
	// A comorbidity or coinfection as part of the indication
	Comorbidity []custom.CodeableReference[ObservationDefinition] `json:"comorbidity,omitempty"`
	// The intended effect, aim or strategy to be achieved
	IntendedEffect *custom.CodeableReference[ObservationDefinition] `json:"intendedEffect,omitempty"`
	// Timing or duration information
	DurationRange *Range `json:"durationRange,omitempty"`
	// Timing or duration information
	DurationString *string `json:"durationString,omitempty"`
	// An unwanted side effect or negative outcome of the subject of this resource when being used for this indication
	UndesirableEffect []custom.Reference[ClinicalUseDefinition] `json:"undesirableEffect,omitempty"`
	// An expression that returns true or false, indicating whether the indication is applicable or not, after having applied its other elements
	Applicability *Expression `json:"applicability,omitempty"`
	// The use of the medicinal product in relation to other therapies described as part of the indication
	OtherTherapy []interface{} `json:"otherTherapy,omitempty"`
}

type ClinicalUseDefinitionInteraction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The specific medication, product, food etc. or laboratory test that interacts
	Interactant []ClinicalUseDefinitionInteractionInteractant `json:"interactant,omitempty"`
	// The type of the interaction e.g. drug-drug interaction, drug-lab test interaction
	Type *CodeableConcept `json:"type,omitempty"`
	// The effect of the interaction, for example "reduced gastric absorption of primary medication"
	Effect *custom.CodeableReference[ObservationDefinition] `json:"effect,omitempty"`
	// The incidence of the interaction, e.g. theoretical, observed
	Incidence *CodeableConcept `json:"incidence,omitempty"`
	// Actions for managing the interaction
	Management []CodeableConcept `json:"management,omitempty"`
}

type ClinicalUseDefinitionInteractionInteractant struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The specific medication, product, food etc. or laboratory test that interacts
	ItemReference *custom.Reference[ClinicalUseDefinitionInteractionInteractantItemReference] `json:"itemReference,omitempty"`
	// The specific medication, product, food etc. or laboratory test that interacts
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
}

type ClinicalUseDefinitionUndesirableEffect struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The situation in which the undesirable effect may manifest
	SymptomConditionEffect *custom.CodeableReference[ObservationDefinition] `json:"symptomConditionEffect,omitempty"`
	// High level classification of the effect
	Classification *CodeableConcept `json:"classification,omitempty"`
	// How often the effect is seen
	FrequencyOfOccurrence *CodeableConcept `json:"frequencyOfOccurrence,omitempty"`
}

type ClinicalUseDefinitionWarning struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A textual definition of this warning, with formatting
	Description *custom.Markdown `json:"description,omitempty"`
	// A coded or unformatted textual definition of this warning
	Code *CodeableConcept `json:"code,omitempty"`
}

type OtherClinicalUseDefinition ClinicalUseDefinition

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

type ClinicalUseDefinitionInteractionInteractantItemReference interface {
	gofhirclient.Resource

	Is_ClinicalUseDefinitionInteractionInteractantItemReference()
}

func (m MedicinalProductDefinition) Is_ClinicalUseDefinitionInteractionInteractantItemReference() {}
func (m Medication) Is_ClinicalUseDefinitionInteractionInteractantItemReference()                 {}
func (s Substance) Is_ClinicalUseDefinitionInteractionInteractantItemReference()                  {}
func (n NutritionProduct) Is_ClinicalUseDefinitionInteractionInteractantItemReference()           {}
func (b BiologicallyDerivedProduct) Is_ClinicalUseDefinitionInteractionInteractantItemReference() {}
func (o ObservationDefinition) Is_ClinicalUseDefinitionInteractionInteractantItemReference()      {}

func (c ClinicalUseDefinition) ResourceType() string {
	return "ClinicalUseDefinition"
}

func (c ClinicalUseDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalUseDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherClinicalUseDefinition: OtherClinicalUseDefinition(c), ResourceType: c.ResourceType()})
}

func UnmarshalClinicalUseDefinition(b []byte) (res ClinicalUseDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
