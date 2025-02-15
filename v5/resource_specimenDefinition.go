// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/SpecimenDefinition
type SpecimenDefinitionTypeTestedContainer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The material type used for the container
	Material *CodeableConcept `bson:"material,omitempty" json:"material,omitempty"`
	// Kind of container associated with the kind of specimen
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Color of container cap
	Cap *CodeableConcept `bson:"cap,omitempty" json:"cap,omitempty"`
	// The description of the kind of container
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The capacity of this kind of container
	Capacity *Quantity `bson:"capacity,omitempty" json:"capacity,omitempty"`
	// Minimum volume
	MinimumVolume *Quantity `bson:"minimumVolume,omitempty" json:"minimumVolume,omitempty"`
	// Additive associated with container
	Additive []SpecimenDefinitionTypeTestedContainerAdditive `bson:"additive,omitempty" json:"additive,omitempty"`
	// Special processing applied to the container for this specimen type
	Preparation *custom.Markdown `bson:"preparation,omitempty" json:"preparation,omitempty"`
}

type SpecimenDefinitionTypeTestedContainerAdditive struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Additive associated with container
	Additive CodeableConcept `bson:"additive" json:"additive"`
}

type SpecimenDefinitionTypeTestedHandling struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Qualifies the interval of temperature
	TemperatureQualifier *CodeableConcept `bson:"temperatureQualifier,omitempty" json:"temperatureQualifier,omitempty"`
	// Temperature range for these handling instructions
	TemperatureRange *Range `bson:"temperatureRange,omitempty" json:"temperatureRange,omitempty"`
	// Maximum preservation time
	MaxDuration *Duration `bson:"maxDuration,omitempty" json:"maxDuration,omitempty"`
	// Preservation instruction
	Instruction *custom.Markdown `bson:"instruction,omitempty" json:"instruction,omitempty"`
}

type SpecimenDefinition struct {
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
	// Logical canonical URL to reference this SpecimenDefinition (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Business identifier
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the SpecimenDefinition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this {{title}} (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this SpecimenDefinition (Human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Based on FHIR definition of another SpecimenDefinition
	DerivedFromCanonical []custom.Canonical[SpecimenDefinition] `bson:"derivedFromCanonical,omitempty" json:"derivedFromCanonical,omitempty"`
	// Based on external definition
	DerivedFromUri []custom.URI `bson:"derivedFromUri,omitempty" json:"derivedFromUri,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// If this SpecimenDefinition is not for real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Type of subject for specimen collection
	Subject *CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// Date status first applied
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// The name of the individual or organization that published the SpecimenDefinition
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the SpecimenDefinition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Content intends to support these contexts
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for this SpecimenDefinition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this SpecimenDefinition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When SpecimenDefinition was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// The date on which the asset content was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// The effective date range for the SpecimenDefinition
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// Kind of material to collect
	TypeCollected *CodeableConcept `bson:"typeCollected,omitempty" json:"typeCollected,omitempty"`
	// Patient preparation for collection
	PatientPreparation []CodeableConcept `bson:"patientPreparation,omitempty" json:"patientPreparation,omitempty"`
	// Time aspect for collection
	TimeAspect *string `bson:"timeAspect,omitempty" json:"timeAspect,omitempty"`
	// Specimen collection procedure
	Collection []CodeableConcept `bson:"collection,omitempty" json:"collection,omitempty"`
	// Specimen in container intended for testing by lab
	TypeTested []SpecimenDefinitionTypeTested `bson:"typeTested,omitempty" json:"typeTested,omitempty"`
}

type SpecimenDefinitionTypeTested struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Primary or secondary specimen
	IsDerived *bool `bson:"isDerived,omitempty" json:"isDerived,omitempty"`
	// Type of intended specimen
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// preferred | alternate
	Preference custom.Code `bson:"preference" json:"preference"`
	// The specimen's container
	Container *SpecimenDefinitionTypeTestedContainer `bson:"container,omitempty" json:"container,omitempty"`
	// Requirements for specimen delivery and special handling
	Requirement *custom.Markdown `bson:"requirement,omitempty" json:"requirement,omitempty"`
	// The usual time for retaining this kind of specimen
	RetentionTime *Duration `bson:"retentionTime,omitempty" json:"retentionTime,omitempty"`
	// Specimen for single use only
	SingleUse *bool `bson:"singleUse,omitempty" json:"singleUse,omitempty"`
	// Criterion specified for specimen rejection
	RejectionCriterion []CodeableConcept `bson:"rejectionCriterion,omitempty" json:"rejectionCriterion,omitempty"`
	// Specimen handling before testing
	Handling []SpecimenDefinitionTypeTestedHandling `bson:"handling,omitempty" json:"handling,omitempty"`
	// Where the specimen will be tested
	TestingDestination []CodeableConcept `bson:"testingDestination,omitempty" json:"testingDestination,omitempty"`
}

func (s SpecimenDefinition) ResourceType() string {
	return "StructureDefinition"
}
