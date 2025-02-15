// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ObservationDefinition
type ObservationDefinition struct {
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
	// Logical canonical URL to reference this ObservationDefinition (globally unique)
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Business identifier of the ObservationDefinition
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Business version of the ObservationDefinition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithm *string `bson:"versionAlgorithm,omitempty" json:"versionAlgorithm,omitempty"`
	// Name for this ObservationDefinition (computer friendly)
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Name for this ObservationDefinition (human friendly)
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// If for testing purposes, not real usage
	Experimental *bool `bson:"experimental,omitempty" json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// The name of the individual or organization that published the ObservationDefinition
	Publisher *string `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `bson:"contact,omitempty" json:"contact,omitempty"`
	// Natural language description of the ObservationDefinition
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Content intends to support these contexts
	UseContext []UsageContext `bson:"useContext,omitempty" json:"useContext,omitempty"`
	// Intended jurisdiction for this ObservationDefinition (if applicable)
	Jurisdiction []CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	// Why this ObservationDefinition is defined
	Purpose *custom.Markdown `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `bson:"copyrightLabel,omitempty" json:"copyrightLabel,omitempty"`
	// When ObservationDefinition was approved by publisher
	ApprovalDate *custom.Date `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	// Date on which the asset content was last reviewed by the publisher
	LastReviewDate *custom.Date `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	// The effective date range for the ObservationDefinition
	EffectivePeriod *Period `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	// Based on FHIR definition of another observation
	DerivedFromCanonical []custom.Canonical[ObservationDefinition] `bson:"derivedFromCanonical,omitempty" json:"derivedFromCanonical,omitempty"`
	// Based on external definition
	DerivedFromUri []custom.URI `bson:"derivedFromUri,omitempty" json:"derivedFromUri,omitempty"`
	// Type of subject for the defined observation
	Subject []CodeableConcept `bson:"subject,omitempty" json:"subject,omitempty"`
	// Desired kind of performer for such kind of observation
	PerformerType *CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty"`
	// General type of observation
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Type of observation
	Code CodeableConcept `bson:"code" json:"code"`
	// Quantity | CodeableConcept | string | boolean | integer | Range | Ratio | SampledData | time | dateTime | Period
	PermittedDataType []custom.Code `bson:"permittedDataType,omitempty" json:"permittedDataType,omitempty"`
	// Multiple results allowed for conforming observations
	MultipleResultsAllowed *bool `bson:"multipleResultsAllowed,omitempty" json:"multipleResultsAllowed,omitempty"`
	// Body part to be observed
	BodySite *CodeableConcept `bson:"bodySite,omitempty" json:"bodySite,omitempty"`
	// Method used to produce the observation
	Method *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	// Kind of specimen used by this type of observation
	Specimen []custom.Reference[SpecimenDefinition] `bson:"specimen,omitempty" json:"specimen,omitempty"`
	// Measurement device or model of device
	Device []custom.Reference[ObservationDefinitionDevice] `bson:"device,omitempty" json:"device,omitempty"`
	// The preferred name to be used when reporting the observation results
	PreferredReportName *string `bson:"preferredReportName,omitempty" json:"preferredReportName,omitempty"`
	// Unit for quantitative results
	PermittedUnit []Coding `bson:"permittedUnit,omitempty" json:"permittedUnit,omitempty"`
	// Set of qualified values for observation results
	QualifiedValue []ObservationDefinitionQualifiedValue `bson:"qualifiedValue,omitempty" json:"qualifiedValue,omitempty"`
	// Definitions of related resources belonging to this kind of observation group
	HasMember []custom.Reference[ObservationDefinitionHasMember] `bson:"hasMember,omitempty" json:"hasMember,omitempty"`
	// Component results
	Component []ObservationDefinitionComponent `bson:"component,omitempty" json:"component,omitempty"`
}

type ObservationDefinitionQualifiedValue struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Context qualifier for the set of qualified values
	Context *CodeableConcept `bson:"context,omitempty" json:"context,omitempty"`
	// Targetted population for the set of qualified values
	AppliesTo []CodeableConcept `bson:"appliesTo,omitempty" json:"appliesTo,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `bson:"gender,omitempty" json:"gender,omitempty"`
	// Applicable age range for the set of qualified values
	Age *Range `bson:"age,omitempty" json:"age,omitempty"`
	// Applicable gestational age range for the set of qualified values
	GestationalAge *Range `bson:"gestationalAge,omitempty" json:"gestationalAge,omitempty"`
	// Condition associated with the set of qualified values
	Condition *string `bson:"condition,omitempty" json:"condition,omitempty"`
	// reference | critical | absolute
	RangeCategory *custom.Code `bson:"rangeCategory,omitempty" json:"rangeCategory,omitempty"`
	// The range for continuous or ordinal observations
	Range *Range `bson:"range,omitempty" json:"range,omitempty"`
	// Value set of valid coded values as part of this set of qualified values
	ValidCodedValueSet *custom.Canonical[ValueSet] `bson:"validCodedValueSet,omitempty" json:"validCodedValueSet,omitempty"`
	// Value set of normal coded values as part of this set of qualified values
	NormalCodedValueSet *custom.Canonical[ValueSet] `bson:"normalCodedValueSet,omitempty" json:"normalCodedValueSet,omitempty"`
	// Value set of abnormal coded values as part of this set of qualified values
	AbnormalCodedValueSet *custom.Canonical[ValueSet] `bson:"abnormalCodedValueSet,omitempty" json:"abnormalCodedValueSet,omitempty"`
	// Value set of critical coded values as part of this set of qualified values
	CriticalCodedValueSet *custom.Canonical[ValueSet] `bson:"criticalCodedValueSet,omitempty" json:"criticalCodedValueSet,omitempty"`
}

type ObservationDefinitionComponent struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Type of observation
	Code CodeableConcept `bson:"code" json:"code"`
	// Quantity | CodeableConcept | string | boolean | integer | Range | Ratio | SampledData | time | dateTime | Period
	PermittedDataType []custom.Code `bson:"permittedDataType,omitempty" json:"permittedDataType,omitempty"`
	// Unit for quantitative results
	PermittedUnit []Coding `bson:"permittedUnit,omitempty" json:"permittedUnit,omitempty"`
	// Set of qualified values for observation results
	QualifiedValue []interface{} `bson:"qualifiedValue,omitempty" json:"qualifiedValue,omitempty"`
}

type ObservationDefinitionDevice interface {
	gofhirclient.Resource

	Is_ObservationDefinitionDevice()
}

func (d DeviceDefinition) Is_ObservationDefinitionDevice() {}
func (d Device) Is_ObservationDefinitionDevice()           {}

type ObservationDefinitionHasMember interface {
	gofhirclient.Resource

	Is_ObservationDefinitionHasMember()
}

func (o ObservationDefinition) Is_ObservationDefinitionHasMember() {}
func (q Questionnaire) Is_ObservationDefinitionHasMember()         {}

func (o ObservationDefinition) ResourceType() string {
	return "StructureDefinition"
}
