// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/SpecimenDefinition
type SpecimenDefinition struct {
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
	// Logical canonical URL to reference this SpecimenDefinition (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Business identifier
	Identifier *Identifier `json:"identifier,omitempty"`
	// Business version of the SpecimenDefinition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this {{title}} (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this SpecimenDefinition (Human friendly)
	Title *string `json:"title,omitempty"`
	// Based on FHIR definition of another SpecimenDefinition
	DerivedFromCanonical []custom.Canonical[SpecimenDefinition] `json:"derivedFromCanonical,omitempty"`
	// Based on external definition
	DerivedFromUri []custom.URI `json:"derivedFromUri,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// If this SpecimenDefinition is not for real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Type of subject for specimen collection
	SubjectCodeableConcept *CodeableConcept `json:"subjectCodeableConcept,omitempty"`
	// Type of subject for specimen collection
	SubjectReference *custom.Reference[Group] `json:"subjectReference,omitempty"`
	// Date status first applied
	Date *custom.DateTime `json:"date,omitempty"`
	// The name of the individual or organization that published the SpecimenDefinition
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the SpecimenDefinition
	Description *custom.Markdown `json:"description,omitempty"`
	// Content intends to support these contexts
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for this SpecimenDefinition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this SpecimenDefinition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When SpecimenDefinition was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// The date on which the asset content was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// The effective date range for the SpecimenDefinition
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// Kind of material to collect
	TypeCollected *CodeableConcept `json:"typeCollected,omitempty"`
	// Patient preparation for collection
	PatientPreparation []CodeableConcept `json:"patientPreparation,omitempty"`
	// Time aspect for collection
	TimeAspect *string `json:"timeAspect,omitempty"`
	// Specimen collection procedure
	Collection []CodeableConcept `json:"collection,omitempty"`
	// Specimen in container intended for testing by lab
	TypeTested []SpecimenDefinitionTypeTested `json:"typeTested,omitempty"`
}

type SpecimenDefinitionTypeTested struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Primary or secondary specimen
	IsDerived *bool `json:"isDerived,omitempty"`
	// Type of intended specimen
	Type *CodeableConcept `json:"type,omitempty"`
	// preferred | alternate
	Preference custom.Code `json:"preference"`
	// The specimen's container
	Container *SpecimenDefinitionTypeTestedContainer `json:"container,omitempty"`
	// Requirements for specimen delivery and special handling
	Requirement *custom.Markdown `json:"requirement,omitempty"`
	// The usual time for retaining this kind of specimen
	RetentionTime *Duration `json:"retentionTime,omitempty"`
	// Specimen for single use only
	SingleUse *bool `json:"singleUse,omitempty"`
	// Criterion specified for specimen rejection
	RejectionCriterion []CodeableConcept `json:"rejectionCriterion,omitempty"`
	// Specimen handling before testing
	Handling []SpecimenDefinitionTypeTestedHandling `json:"handling,omitempty"`
	// Where the specimen will be tested
	TestingDestination []CodeableConcept `json:"testingDestination,omitempty"`
}

type SpecimenDefinitionTypeTestedContainer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The material type used for the container
	Material *CodeableConcept `json:"material,omitempty"`
	// Kind of container associated with the kind of specimen
	Type *CodeableConcept `json:"type,omitempty"`
	// Color of container cap
	Cap *CodeableConcept `json:"cap,omitempty"`
	// The description of the kind of container
	Description *custom.Markdown `json:"description,omitempty"`
	// The capacity of this kind of container
	Capacity *Quantity `json:"capacity,omitempty"`
	// Minimum volume
	MinimumVolumeQuantity *Quantity `json:"minimumVolumeQuantity,omitempty"`
	// Minimum volume
	MinimumVolumeString *string `json:"minimumVolumeString,omitempty"`
	// Additive associated with container
	Additive []SpecimenDefinitionTypeTestedContainerAdditive `json:"additive,omitempty"`
	// Special processing applied to the container for this specimen type
	Preparation *custom.Markdown `json:"preparation,omitempty"`
}

type SpecimenDefinitionTypeTestedContainerAdditive struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Additive associated with container
	AdditiveCodeableConcept *CodeableConcept `json:"additiveCodeableConcept,omitempty"`
	// Additive associated with container
	AdditiveReference *custom.Reference[SubstanceDefinition] `json:"additiveReference,omitempty"`
}

type SpecimenDefinitionTypeTestedHandling struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Qualifies the interval of temperature
	TemperatureQualifier *CodeableConcept `json:"temperatureQualifier,omitempty"`
	// Temperature range for these handling instructions
	TemperatureRange *Range `json:"temperatureRange,omitempty"`
	// Maximum preservation time
	MaxDuration *Duration `json:"maxDuration,omitempty"`
	// Preservation instruction
	Instruction *custom.Markdown `json:"instruction,omitempty"`
}

type OtherSpecimenDefinition SpecimenDefinition

func (s SpecimenDefinition) ResourceType() string {
	return "SpecimenDefinition"
}

func (s SpecimenDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSpecimenDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherSpecimenDefinition: OtherSpecimenDefinition(s), ResourceType: s.ResourceType()})
}

func UnmarshalSpecimenDefinition(b []byte) (res SpecimenDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
