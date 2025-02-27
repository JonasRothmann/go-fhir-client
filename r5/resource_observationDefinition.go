// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ObservationDefinition
type ObservationDefinition struct {
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
	// Logical canonical URL to reference this ObservationDefinition (globally unique)
	Url *custom.URI `json:"url,omitempty"`
	// Business identifier of the ObservationDefinition
	Identifier *Identifier `json:"identifier,omitempty"`
	// Business version of the ObservationDefinition
	Version *string `json:"version,omitempty"`
	// How to compare versions
	VersionAlgorithmString *string `json:"versionAlgorithmString,omitempty"`
	// How to compare versions
	VersionAlgorithmCoding *Coding `json:"versionAlgorithmCoding,omitempty"`
	// Name for this ObservationDefinition (computer friendly)
	Name *string `json:"name,omitempty"`
	// Name for this ObservationDefinition (human friendly)
	Title *string `json:"title,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// If for testing purposes, not real usage
	Experimental *bool `json:"experimental,omitempty"`
	// Date last changed
	Date *custom.DateTime `json:"date,omitempty"`
	// The name of the individual or organization that published the ObservationDefinition
	Publisher *string `json:"publisher,omitempty"`
	// Contact details for the publisher
	Contact []ContactDetail `json:"contact,omitempty"`
	// Natural language description of the ObservationDefinition
	Description *custom.Markdown `json:"description,omitempty"`
	// Content intends to support these contexts
	UseContext []UsageContext `json:"useContext,omitempty"`
	// Intended jurisdiction for this ObservationDefinition (if applicable)
	Jurisdiction []CodeableConcept `json:"jurisdiction,omitempty"`
	// Why this ObservationDefinition is defined
	Purpose *custom.Markdown `json:"purpose,omitempty"`
	// Use and/or publishing restrictions
	Copyright *custom.Markdown `json:"copyright,omitempty"`
	// Copyright holder and year(s)
	CopyrightLabel *string `json:"copyrightLabel,omitempty"`
	// When ObservationDefinition was approved by publisher
	ApprovalDate *custom.Date `json:"approvalDate,omitempty"`
	// Date on which the asset content was last reviewed by the publisher
	LastReviewDate *custom.Date `json:"lastReviewDate,omitempty"`
	// The effective date range for the ObservationDefinition
	EffectivePeriod *Period `json:"effectivePeriod,omitempty"`
	// Based on FHIR definition of another observation
	DerivedFromCanonical []custom.Canonical[ObservationDefinition] `json:"derivedFromCanonical,omitempty"`
	// Based on external definition
	DerivedFromUri []custom.URI `json:"derivedFromUri,omitempty"`
	// Type of subject for the defined observation
	Subject []CodeableConcept `json:"subject,omitempty"`
	// Desired kind of performer for such kind of observation
	PerformerType *CodeableConcept `json:"performerType,omitempty"`
	// General type of observation
	Category []CodeableConcept `json:"category,omitempty"`
	// Type of observation
	Code CodeableConcept `json:"code"`
	// Quantity | CodeableConcept | string | boolean | integer | Range | Ratio | SampledData | time | dateTime | Period
	PermittedDataType []custom.Code `json:"permittedDataType,omitempty"`
	// Multiple results allowed for conforming observations
	MultipleResultsAllowed *bool `json:"multipleResultsAllowed,omitempty"`
	// Body part to be observed
	BodySite *CodeableConcept `json:"bodySite,omitempty"`
	// Method used to produce the observation
	Method *CodeableConcept `json:"method,omitempty"`
	// Kind of specimen used by this type of observation
	Specimen []custom.Reference[SpecimenDefinition] `json:"specimen,omitempty"`
	// Measurement device or model of device
	Device []custom.Reference[ObservationDefinitionDevice] `json:"device,omitempty"`
	// The preferred name to be used when reporting the observation results
	PreferredReportName *string `json:"preferredReportName,omitempty"`
	// Unit for quantitative results
	PermittedUnit []Coding `json:"permittedUnit,omitempty"`
	// Set of qualified values for observation results
	QualifiedValue []ObservationDefinitionQualifiedValue `json:"qualifiedValue,omitempty"`
	// Definitions of related resources belonging to this kind of observation group
	HasMember []custom.Reference[ObservationDefinitionHasMember] `json:"hasMember,omitempty"`
	// Component results
	Component []ObservationDefinitionComponent `json:"component,omitempty"`
}

type ObservationDefinitionQualifiedValue struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Context qualifier for the set of qualified values
	Context *CodeableConcept `json:"context,omitempty"`
	// Targetted population for the set of qualified values
	AppliesTo []CodeableConcept `json:"appliesTo,omitempty"`
	// male | female | other | unknown
	Gender *custom.Code `json:"gender,omitempty"`
	// Applicable age range for the set of qualified values
	Age *Range `json:"age,omitempty"`
	// Applicable gestational age range for the set of qualified values
	GestationalAge *Range `json:"gestationalAge,omitempty"`
	// Condition associated with the set of qualified values
	Condition *string `json:"condition,omitempty"`
	// reference | critical | absolute
	RangeCategory *custom.Code `json:"rangeCategory,omitempty"`
	// The range for continuous or ordinal observations
	Range *Range `json:"range,omitempty"`
	// Value set of valid coded values as part of this set of qualified values
	ValidCodedValueSet *custom.Canonical[ValueSet] `json:"validCodedValueSet,omitempty"`
	// Value set of normal coded values as part of this set of qualified values
	NormalCodedValueSet *custom.Canonical[ValueSet] `json:"normalCodedValueSet,omitempty"`
	// Value set of abnormal coded values as part of this set of qualified values
	AbnormalCodedValueSet *custom.Canonical[ValueSet] `json:"abnormalCodedValueSet,omitempty"`
	// Value set of critical coded values as part of this set of qualified values
	CriticalCodedValueSet *custom.Canonical[ValueSet] `json:"criticalCodedValueSet,omitempty"`
}

type ObservationDefinitionComponent struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of observation
	Code CodeableConcept `json:"code"`
	// Quantity | CodeableConcept | string | boolean | integer | Range | Ratio | SampledData | time | dateTime | Period
	PermittedDataType []custom.Code `json:"permittedDataType,omitempty"`
	// Unit for quantitative results
	PermittedUnit []Coding `json:"permittedUnit,omitempty"`
	// Set of qualified values for observation results
	QualifiedValue []interface{} `json:"qualifiedValue,omitempty"`
}

type OtherObservationDefinition ObservationDefinition

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
	return "ObservationDefinition"
}

func (o ObservationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherObservationDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherObservationDefinition: OtherObservationDefinition(o), ResourceType: o.ResourceType()})
}

func UnmarshalObservationDefinition(b []byte) (res ObservationDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
