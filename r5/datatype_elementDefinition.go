// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/ElementDefinition
type ElementDefinitionMapping struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Reference to mapping declaration
	Identity custom.ID `json:"identity"`
	// Computable language of mapping
	Language *custom.Code `json:"language,omitempty"`
	// Details of the mapping
	Map string `json:"map"`
	// Comments about the mapping or its use
	Comment *custom.Markdown `json:"comment,omitempty"`
}

type ElementDefinitionSlicing struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Element values that are used to distinguish the slices
	Discriminator []ElementDefinitionSlicingDiscriminator `json:"discriminator,omitempty"`
	// Text description of how slicing works (or not)
	Description *string `json:"description,omitempty"`
	// If elements must be in same order as slices
	Ordered *bool `json:"ordered,omitempty"`
	// closed | open | openAtEnd
	Rules custom.Code `json:"rules"`
}

type ElementDefinitionBase struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Path that identifies the base element
	Path string `json:"path"`
	// Min cardinality of the base element
	Min uint32 `json:"min"`
	// Max cardinality of the base element
	Max string `json:"max"`
}

type ElementDefinitionExample struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Describes the purpose of this example
	Label string `json:"label"`
	// Value of Example (one of allowed types)
	ValueBase64binary *custom.Base64binary `json:"valueBase64binary,omitempty"`
	// Value of Example (one of allowed types)
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Value of Example (one of allowed types)
	ValueCanonical *custom.Canonical[any] `json:"valueCanonical,omitempty"`
	// Value of Example (one of allowed types)
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Value of Example (one of allowed types)
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Value of Example (one of allowed types)
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Value of Example (one of allowed types)
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Value of Example (one of allowed types)
	ValueId *custom.ID `json:"valueId,omitempty"`
	// Value of Example (one of allowed types)
	ValueInstant *custom.Instant `json:"valueInstant,omitempty"`
	// Value of Example (one of allowed types)
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Value of Example (one of allowed types)
	ValueInteger64 *int64 `json:"valueInteger64,omitempty"`
	// Value of Example (one of allowed types)
	ValueMarkdown *custom.Markdown `json:"valueMarkdown,omitempty"`
	// Value of Example (one of allowed types)
	ValueOid *custom.OID `json:"valueOid,omitempty"`
	// Value of Example (one of allowed types)
	ValuePositiveInt *uint32 `json:"valuePositiveInt,omitempty"`
	// Value of Example (one of allowed types)
	ValueString *string `json:"valueString,omitempty"`
	// Value of Example (one of allowed types)
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Value of Example (one of allowed types)
	ValueUnsignedInt *uint32 `json:"valueUnsignedInt,omitempty"`
	// Value of Example (one of allowed types)
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Value of Example (one of allowed types)
	ValueUrl *custom.URL `json:"valueUrl,omitempty"`
	// Value of Example (one of allowed types)
	ValueUuid *custom.UUID `json:"valueUuid,omitempty"`
	// Value of Example (one of allowed types)
	ValueAddress *Address `json:"valueAddress,omitempty"`
	// Value of Example (one of allowed types)
	ValueAge *Age `json:"valueAge,omitempty"`
	// Value of Example (one of allowed types)
	ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`
	// Value of Example (one of allowed types)
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Value of Example (one of allowed types)
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Value of Example (one of allowed types)
	ValueCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"valueCodeableReference,omitempty"`
	// Value of Example (one of allowed types)
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Value of Example (one of allowed types)
	ValueContactPoint *ContactPoint `json:"valueContactPoint,omitempty"`
	// Value of Example (one of allowed types)
	ValueCount *Count `json:"valueCount,omitempty"`
	// Value of Example (one of allowed types)
	ValueDistance *Distance `json:"valueDistance,omitempty"`
	// Value of Example (one of allowed types)
	ValueDuration *Duration `json:"valueDuration,omitempty"`
	// Value of Example (one of allowed types)
	ValueHumanName *HumanName `json:"valueHumanName,omitempty"`
	// Value of Example (one of allowed types)
	ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`
	// Value of Example (one of allowed types)
	ValueMoney *Money `json:"valueMoney,omitempty"`
	// Value of Example (one of allowed types)
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Value of Example (one of allowed types)
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Value of Example (one of allowed types)
	ValueRange *Range `json:"valueRange,omitempty"`
	// Value of Example (one of allowed types)
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Value of Example (one of allowed types)
	ValueRatioRange *RatioRange `json:"valueRatioRange,omitempty"`
	// Value of Example (one of allowed types)
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// Value of Example (one of allowed types)
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Value of Example (one of allowed types)
	ValueSignature *Signature `json:"valueSignature,omitempty"`
	// Value of Example (one of allowed types)
	ValueTiming *Timing `json:"valueTiming,omitempty"`
	// Value of Example (one of allowed types)
	ValueContactDetail *ContactDetail `json:"valueContactDetail,omitempty"`
	// Value of Example (one of allowed types)
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement,omitempty"`
	// Value of Example (one of allowed types)
	ValueExpression *Expression `json:"valueExpression,omitempty"`
	// Value of Example (one of allowed types)
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	// Value of Example (one of allowed types)
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact,omitempty"`
	// Value of Example (one of allowed types)
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition,omitempty"`
	// Value of Example (one of allowed types)
	ValueUsageContext *UsageContext `json:"valueUsageContext,omitempty"`
	// Value of Example (one of allowed types)
	ValueAvailability *Availability `json:"valueAvailability,omitempty"`
	// Value of Example (one of allowed types)
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty"`
	// Value of Example (one of allowed types)
	ValueDosage *Dosage `json:"valueDosage,omitempty"`
	// Value of Example (one of allowed types)
	ValueMeta *Meta `json:"valueMeta,omitempty"`
}

type ElementDefinitionBinding struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// required | extensible | preferred | example
	Strength custom.Code `json:"strength"`
	// Intended use of codes in the bound value set
	Description *custom.Markdown `json:"description,omitempty"`
	// Source of value set
	ValueSet *custom.Canonical[ValueSet] `json:"valueSet,omitempty"`
	// Additional Bindings - more rules about the binding
	Additional []ElementDefinitionBindingAdditional `json:"additional,omitempty"`
}

type ElementDefinitionBindingAdditional struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// maximum | minimum | required | extensible | candidate | current | preferred | ui | starter | component
	Purpose custom.Code `json:"purpose"`
	// The value set for the additional binding
	ValueSet custom.Canonical[ValueSet] `json:"valueSet"`
	// Documentation of the purpose of use of the binding
	Documentation *custom.Markdown `json:"documentation,omitempty"`
	// Concise documentation - for summary tables
	ShortDoco *string `json:"shortDoco,omitempty"`
	// Qualifies the usage - jurisdiction, gender, workflow status etc.
	Usage []UsageContext `json:"usage,omitempty"`
	// Whether binding can applies to all repeats, or just one
	Any *bool `json:"any,omitempty"`
}

type ElementDefinition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Path of the element in the hierarchy of elements
	Path string `json:"path"`
	// xmlAttr | xmlText | typeAttr | cdaText | xhtml
	Representation []custom.Code `json:"representation,omitempty"`
	// Name for this particular element (in a set of slices)
	SliceName *string `json:"sliceName,omitempty"`
	// If this slice definition constrains an inherited slice definition (or not)
	SliceIsConstraining *bool `json:"sliceIsConstraining,omitempty"`
	// Name for element to display with or prompt for element
	Label *string `json:"label,omitempty"`
	// Corresponding codes in terminologies
	Code []Coding `json:"code,omitempty"`
	// This element is sliced - slices follow
	Slicing *ElementDefinitionSlicing `json:"slicing,omitempty"`
	// Concise definition for space-constrained presentation
	Short *string `json:"short,omitempty"`
	// Full formal definition as narrative text
	Definition *custom.Markdown `json:"definition,omitempty"`
	// Comments about the use of this element
	Comment *custom.Markdown `json:"comment,omitempty"`
	// Why this resource has been created
	Requirements *custom.Markdown `json:"requirements,omitempty"`
	// Other names
	Alias []string `json:"alias,omitempty"`
	// Minimum Cardinality
	Min *uint32 `json:"min,omitempty"`
	// Maximum Cardinality (a number or *)
	Max *string `json:"max,omitempty"`
	// Base definition information for tools
	Base *ElementDefinitionBase `json:"base,omitempty"`
	// Reference to definition of content for the element
	ContentReference *custom.URI `json:"contentReference,omitempty"`
	// Data type and Profile for this element
	Type []ElementDefinitionType `json:"type,omitempty"`
	// Specified value if missing from instance
	DefaultValueBase64binary *custom.Base64binary `json:"defaultValueBase64binary,omitempty"`
	// Specified value if missing from instance
	DefaultValueBoolean *bool `json:"defaultValueBoolean,omitempty"`
	// Specified value if missing from instance
	DefaultValueCanonical *custom.Canonical[any] `json:"defaultValueCanonical,omitempty"`
	// Specified value if missing from instance
	DefaultValueCode *custom.Code `json:"defaultValueCode,omitempty"`
	// Specified value if missing from instance
	DefaultValueDate *custom.Date `json:"defaultValueDate,omitempty"`
	// Specified value if missing from instance
	DefaultValueDateTime *custom.DateTime `json:"defaultValueDateTime,omitempty"`
	// Specified value if missing from instance
	DefaultValueDecimal *json.Number `json:"defaultValueDecimal,omitempty"`
	// Specified value if missing from instance
	DefaultValueId *custom.ID `json:"defaultValueId,omitempty"`
	// Specified value if missing from instance
	DefaultValueInstant *custom.Instant `json:"defaultValueInstant,omitempty"`
	// Specified value if missing from instance
	DefaultValueInteger *int32 `json:"defaultValueInteger,omitempty"`
	// Specified value if missing from instance
	DefaultValueInteger64 *int64 `json:"defaultValueInteger64,omitempty"`
	// Specified value if missing from instance
	DefaultValueMarkdown *custom.Markdown `json:"defaultValueMarkdown,omitempty"`
	// Specified value if missing from instance
	DefaultValueOid *custom.OID `json:"defaultValueOid,omitempty"`
	// Specified value if missing from instance
	DefaultValuePositiveInt *uint32 `json:"defaultValuePositiveInt,omitempty"`
	// Specified value if missing from instance
	DefaultValueString *string `json:"defaultValueString,omitempty"`
	// Specified value if missing from instance
	DefaultValueTime *custom.Time `json:"defaultValueTime,omitempty"`
	// Specified value if missing from instance
	DefaultValueUnsignedInt *uint32 `json:"defaultValueUnsignedInt,omitempty"`
	// Specified value if missing from instance
	DefaultValueUri *custom.URI `json:"defaultValueUri,omitempty"`
	// Specified value if missing from instance
	DefaultValueUrl *custom.URL `json:"defaultValueUrl,omitempty"`
	// Specified value if missing from instance
	DefaultValueUuid *custom.UUID `json:"defaultValueUuid,omitempty"`
	// Specified value if missing from instance
	DefaultValueAddress *Address `json:"defaultValueAddress,omitempty"`
	// Specified value if missing from instance
	DefaultValueAge *Age `json:"defaultValueAge,omitempty"`
	// Specified value if missing from instance
	DefaultValueAnnotation *Annotation `json:"defaultValueAnnotation,omitempty"`
	// Specified value if missing from instance
	DefaultValueAttachment *Attachment `json:"defaultValueAttachment,omitempty"`
	// Specified value if missing from instance
	DefaultValueCodeableConcept *CodeableConcept `json:"defaultValueCodeableConcept,omitempty"`
	// Specified value if missing from instance
	DefaultValueCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"defaultValueCodeableReference,omitempty"`
	// Specified value if missing from instance
	DefaultValueCoding *Coding `json:"defaultValueCoding,omitempty"`
	// Specified value if missing from instance
	DefaultValueContactPoint *ContactPoint `json:"defaultValueContactPoint,omitempty"`
	// Specified value if missing from instance
	DefaultValueCount *Count `json:"defaultValueCount,omitempty"`
	// Specified value if missing from instance
	DefaultValueDistance *Distance `json:"defaultValueDistance,omitempty"`
	// Specified value if missing from instance
	DefaultValueDuration *Duration `json:"defaultValueDuration,omitempty"`
	// Specified value if missing from instance
	DefaultValueHumanName *HumanName `json:"defaultValueHumanName,omitempty"`
	// Specified value if missing from instance
	DefaultValueIdentifier *Identifier `json:"defaultValueIdentifier,omitempty"`
	// Specified value if missing from instance
	DefaultValueMoney *Money `json:"defaultValueMoney,omitempty"`
	// Specified value if missing from instance
	DefaultValuePeriod *Period `json:"defaultValuePeriod,omitempty"`
	// Specified value if missing from instance
	DefaultValueQuantity *Quantity `json:"defaultValueQuantity,omitempty"`
	// Specified value if missing from instance
	DefaultValueRange *Range `json:"defaultValueRange,omitempty"`
	// Specified value if missing from instance
	DefaultValueRatio *Ratio `json:"defaultValueRatio,omitempty"`
	// Specified value if missing from instance
	DefaultValueRatioRange *RatioRange `json:"defaultValueRatioRange,omitempty"`
	// Specified value if missing from instance
	DefaultValueReference *custom.Reference[gofhirclient.Resource] `json:"defaultValueReference,omitempty"`
	// Specified value if missing from instance
	DefaultValueSampledData *SampledData `json:"defaultValueSampledData,omitempty"`
	// Specified value if missing from instance
	DefaultValueSignature *Signature `json:"defaultValueSignature,omitempty"`
	// Specified value if missing from instance
	DefaultValueTiming *Timing `json:"defaultValueTiming,omitempty"`
	// Specified value if missing from instance
	DefaultValueContactDetail *ContactDetail `json:"defaultValueContactDetail,omitempty"`
	// Specified value if missing from instance
	DefaultValueDataRequirement *DataRequirement `json:"defaultValueDataRequirement,omitempty"`
	// Specified value if missing from instance
	DefaultValueExpression *Expression `json:"defaultValueExpression,omitempty"`
	// Specified value if missing from instance
	DefaultValueParameterDefinition *ParameterDefinition `json:"defaultValueParameterDefinition,omitempty"`
	// Specified value if missing from instance
	DefaultValueRelatedArtifact *RelatedArtifact `json:"defaultValueRelatedArtifact,omitempty"`
	// Specified value if missing from instance
	DefaultValueTriggerDefinition *TriggerDefinition `json:"defaultValueTriggerDefinition,omitempty"`
	// Specified value if missing from instance
	DefaultValueUsageContext *UsageContext `json:"defaultValueUsageContext,omitempty"`
	// Specified value if missing from instance
	DefaultValueAvailability *Availability `json:"defaultValueAvailability,omitempty"`
	// Specified value if missing from instance
	DefaultValueExtendedContactDetail *ExtendedContactDetail `json:"defaultValueExtendedContactDetail,omitempty"`
	// Specified value if missing from instance
	DefaultValueDosage *Dosage `json:"defaultValueDosage,omitempty"`
	// Specified value if missing from instance
	DefaultValueMeta *Meta `json:"defaultValueMeta,omitempty"`
	// Implicit meaning when this element is missing
	MeaningWhenMissing *custom.Markdown `json:"meaningWhenMissing,omitempty"`
	// What the order of the elements means
	OrderMeaning *string `json:"orderMeaning,omitempty"`
	// Value must be exactly this
	FixedBase64binary *custom.Base64binary `json:"fixedBase64binary,omitempty"`
	// Value must be exactly this
	FixedBoolean *bool `json:"fixedBoolean,omitempty"`
	// Value must be exactly this
	FixedCanonical *custom.Canonical[any] `json:"fixedCanonical,omitempty"`
	// Value must be exactly this
	FixedCode *custom.Code `json:"fixedCode,omitempty"`
	// Value must be exactly this
	FixedDate *custom.Date `json:"fixedDate,omitempty"`
	// Value must be exactly this
	FixedDateTime *custom.DateTime `json:"fixedDateTime,omitempty"`
	// Value must be exactly this
	FixedDecimal *json.Number `json:"fixedDecimal,omitempty"`
	// Value must be exactly this
	FixedId *custom.ID `json:"fixedId,omitempty"`
	// Value must be exactly this
	FixedInstant *custom.Instant `json:"fixedInstant,omitempty"`
	// Value must be exactly this
	FixedInteger *int32 `json:"fixedInteger,omitempty"`
	// Value must be exactly this
	FixedInteger64 *int64 `json:"fixedInteger64,omitempty"`
	// Value must be exactly this
	FixedMarkdown *custom.Markdown `json:"fixedMarkdown,omitempty"`
	// Value must be exactly this
	FixedOid *custom.OID `json:"fixedOid,omitempty"`
	// Value must be exactly this
	FixedPositiveInt *uint32 `json:"fixedPositiveInt,omitempty"`
	// Value must be exactly this
	FixedString *string `json:"fixedString,omitempty"`
	// Value must be exactly this
	FixedTime *custom.Time `json:"fixedTime,omitempty"`
	// Value must be exactly this
	FixedUnsignedInt *uint32 `json:"fixedUnsignedInt,omitempty"`
	// Value must be exactly this
	FixedUri *custom.URI `json:"fixedUri,omitempty"`
	// Value must be exactly this
	FixedUrl *custom.URL `json:"fixedUrl,omitempty"`
	// Value must be exactly this
	FixedUuid *custom.UUID `json:"fixedUuid,omitempty"`
	// Value must be exactly this
	FixedAddress *Address `json:"fixedAddress,omitempty"`
	// Value must be exactly this
	FixedAge *Age `json:"fixedAge,omitempty"`
	// Value must be exactly this
	FixedAnnotation *Annotation `json:"fixedAnnotation,omitempty"`
	// Value must be exactly this
	FixedAttachment *Attachment `json:"fixedAttachment,omitempty"`
	// Value must be exactly this
	FixedCodeableConcept *CodeableConcept `json:"fixedCodeableConcept,omitempty"`
	// Value must be exactly this
	FixedCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"fixedCodeableReference,omitempty"`
	// Value must be exactly this
	FixedCoding *Coding `json:"fixedCoding,omitempty"`
	// Value must be exactly this
	FixedContactPoint *ContactPoint `json:"fixedContactPoint,omitempty"`
	// Value must be exactly this
	FixedCount *Count `json:"fixedCount,omitempty"`
	// Value must be exactly this
	FixedDistance *Distance `json:"fixedDistance,omitempty"`
	// Value must be exactly this
	FixedDuration *Duration `json:"fixedDuration,omitempty"`
	// Value must be exactly this
	FixedHumanName *HumanName `json:"fixedHumanName,omitempty"`
	// Value must be exactly this
	FixedIdentifier *Identifier `json:"fixedIdentifier,omitempty"`
	// Value must be exactly this
	FixedMoney *Money `json:"fixedMoney,omitempty"`
	// Value must be exactly this
	FixedPeriod *Period `json:"fixedPeriod,omitempty"`
	// Value must be exactly this
	FixedQuantity *Quantity `json:"fixedQuantity,omitempty"`
	// Value must be exactly this
	FixedRange *Range `json:"fixedRange,omitempty"`
	// Value must be exactly this
	FixedRatio *Ratio `json:"fixedRatio,omitempty"`
	// Value must be exactly this
	FixedRatioRange *RatioRange `json:"fixedRatioRange,omitempty"`
	// Value must be exactly this
	FixedReference *custom.Reference[gofhirclient.Resource] `json:"fixedReference,omitempty"`
	// Value must be exactly this
	FixedSampledData *SampledData `json:"fixedSampledData,omitempty"`
	// Value must be exactly this
	FixedSignature *Signature `json:"fixedSignature,omitempty"`
	// Value must be exactly this
	FixedTiming *Timing `json:"fixedTiming,omitempty"`
	// Value must be exactly this
	FixedContactDetail *ContactDetail `json:"fixedContactDetail,omitempty"`
	// Value must be exactly this
	FixedDataRequirement *DataRequirement `json:"fixedDataRequirement,omitempty"`
	// Value must be exactly this
	FixedExpression *Expression `json:"fixedExpression,omitempty"`
	// Value must be exactly this
	FixedParameterDefinition *ParameterDefinition `json:"fixedParameterDefinition,omitempty"`
	// Value must be exactly this
	FixedRelatedArtifact *RelatedArtifact `json:"fixedRelatedArtifact,omitempty"`
	// Value must be exactly this
	FixedTriggerDefinition *TriggerDefinition `json:"fixedTriggerDefinition,omitempty"`
	// Value must be exactly this
	FixedUsageContext *UsageContext `json:"fixedUsageContext,omitempty"`
	// Value must be exactly this
	FixedAvailability *Availability `json:"fixedAvailability,omitempty"`
	// Value must be exactly this
	FixedExtendedContactDetail *ExtendedContactDetail `json:"fixedExtendedContactDetail,omitempty"`
	// Value must be exactly this
	FixedDosage *Dosage `json:"fixedDosage,omitempty"`
	// Value must be exactly this
	FixedMeta *Meta `json:"fixedMeta,omitempty"`
	// Value must have at least these property values
	PatternBase64binary *custom.Base64binary `json:"patternBase64binary,omitempty"`
	// Value must have at least these property values
	PatternBoolean *bool `json:"patternBoolean,omitempty"`
	// Value must have at least these property values
	PatternCanonical *custom.Canonical[any] `json:"patternCanonical,omitempty"`
	// Value must have at least these property values
	PatternCode *custom.Code `json:"patternCode,omitempty"`
	// Value must have at least these property values
	PatternDate *custom.Date `json:"patternDate,omitempty"`
	// Value must have at least these property values
	PatternDateTime *custom.DateTime `json:"patternDateTime,omitempty"`
	// Value must have at least these property values
	PatternDecimal *json.Number `json:"patternDecimal,omitempty"`
	// Value must have at least these property values
	PatternId *custom.ID `json:"patternId,omitempty"`
	// Value must have at least these property values
	PatternInstant *custom.Instant `json:"patternInstant,omitempty"`
	// Value must have at least these property values
	PatternInteger *int32 `json:"patternInteger,omitempty"`
	// Value must have at least these property values
	PatternInteger64 *int64 `json:"patternInteger64,omitempty"`
	// Value must have at least these property values
	PatternMarkdown *custom.Markdown `json:"patternMarkdown,omitempty"`
	// Value must have at least these property values
	PatternOid *custom.OID `json:"patternOid,omitempty"`
	// Value must have at least these property values
	PatternPositiveInt *uint32 `json:"patternPositiveInt,omitempty"`
	// Value must have at least these property values
	PatternString *string `json:"patternString,omitempty"`
	// Value must have at least these property values
	PatternTime *custom.Time `json:"patternTime,omitempty"`
	// Value must have at least these property values
	PatternUnsignedInt *uint32 `json:"patternUnsignedInt,omitempty"`
	// Value must have at least these property values
	PatternUri *custom.URI `json:"patternUri,omitempty"`
	// Value must have at least these property values
	PatternUrl *custom.URL `json:"patternUrl,omitempty"`
	// Value must have at least these property values
	PatternUuid *custom.UUID `json:"patternUuid,omitempty"`
	// Value must have at least these property values
	PatternAddress *Address `json:"patternAddress,omitempty"`
	// Value must have at least these property values
	PatternAge *Age `json:"patternAge,omitempty"`
	// Value must have at least these property values
	PatternAnnotation *Annotation `json:"patternAnnotation,omitempty"`
	// Value must have at least these property values
	PatternAttachment *Attachment `json:"patternAttachment,omitempty"`
	// Value must have at least these property values
	PatternCodeableConcept *CodeableConcept `json:"patternCodeableConcept,omitempty"`
	// Value must have at least these property values
	PatternCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"patternCodeableReference,omitempty"`
	// Value must have at least these property values
	PatternCoding *Coding `json:"patternCoding,omitempty"`
	// Value must have at least these property values
	PatternContactPoint *ContactPoint `json:"patternContactPoint,omitempty"`
	// Value must have at least these property values
	PatternCount *Count `json:"patternCount,omitempty"`
	// Value must have at least these property values
	PatternDistance *Distance `json:"patternDistance,omitempty"`
	// Value must have at least these property values
	PatternDuration *Duration `json:"patternDuration,omitempty"`
	// Value must have at least these property values
	PatternHumanName *HumanName `json:"patternHumanName,omitempty"`
	// Value must have at least these property values
	PatternIdentifier *Identifier `json:"patternIdentifier,omitempty"`
	// Value must have at least these property values
	PatternMoney *Money `json:"patternMoney,omitempty"`
	// Value must have at least these property values
	PatternPeriod *Period `json:"patternPeriod,omitempty"`
	// Value must have at least these property values
	PatternQuantity *Quantity `json:"patternQuantity,omitempty"`
	// Value must have at least these property values
	PatternRange *Range `json:"patternRange,omitempty"`
	// Value must have at least these property values
	PatternRatio *Ratio `json:"patternRatio,omitempty"`
	// Value must have at least these property values
	PatternRatioRange *RatioRange `json:"patternRatioRange,omitempty"`
	// Value must have at least these property values
	PatternReference *custom.Reference[gofhirclient.Resource] `json:"patternReference,omitempty"`
	// Value must have at least these property values
	PatternSampledData *SampledData `json:"patternSampledData,omitempty"`
	// Value must have at least these property values
	PatternSignature *Signature `json:"patternSignature,omitempty"`
	// Value must have at least these property values
	PatternTiming *Timing `json:"patternTiming,omitempty"`
	// Value must have at least these property values
	PatternContactDetail *ContactDetail `json:"patternContactDetail,omitempty"`
	// Value must have at least these property values
	PatternDataRequirement *DataRequirement `json:"patternDataRequirement,omitempty"`
	// Value must have at least these property values
	PatternExpression *Expression `json:"patternExpression,omitempty"`
	// Value must have at least these property values
	PatternParameterDefinition *ParameterDefinition `json:"patternParameterDefinition,omitempty"`
	// Value must have at least these property values
	PatternRelatedArtifact *RelatedArtifact `json:"patternRelatedArtifact,omitempty"`
	// Value must have at least these property values
	PatternTriggerDefinition *TriggerDefinition `json:"patternTriggerDefinition,omitempty"`
	// Value must have at least these property values
	PatternUsageContext *UsageContext `json:"patternUsageContext,omitempty"`
	// Value must have at least these property values
	PatternAvailability *Availability `json:"patternAvailability,omitempty"`
	// Value must have at least these property values
	PatternExtendedContactDetail *ExtendedContactDetail `json:"patternExtendedContactDetail,omitempty"`
	// Value must have at least these property values
	PatternDosage *Dosage `json:"patternDosage,omitempty"`
	// Value must have at least these property values
	PatternMeta *Meta `json:"patternMeta,omitempty"`
	// Example value (as defined for type)
	Example []ElementDefinitionExample `json:"example,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueDate *custom.Date `json:"minValueDate,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueDateTime *custom.DateTime `json:"minValueDateTime,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueInstant *custom.Instant `json:"minValueInstant,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueTime *custom.Time `json:"minValueTime,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueDecimal *json.Number `json:"minValueDecimal,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueInteger *int32 `json:"minValueInteger,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueInteger64 *int64 `json:"minValueInteger64,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValuePositiveInt *uint32 `json:"minValuePositiveInt,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueUnsignedInt *uint32 `json:"minValueUnsignedInt,omitempty"`
	// Minimum Allowed Value (for some types)
	MinValueQuantity *Quantity `json:"minValueQuantity,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueDate *custom.Date `json:"maxValueDate,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueDateTime *custom.DateTime `json:"maxValueDateTime,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueInstant *custom.Instant `json:"maxValueInstant,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueTime *custom.Time `json:"maxValueTime,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueDecimal *json.Number `json:"maxValueDecimal,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueInteger *int32 `json:"maxValueInteger,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueInteger64 *int64 `json:"maxValueInteger64,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValuePositiveInt *uint32 `json:"maxValuePositiveInt,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueUnsignedInt *uint32 `json:"maxValueUnsignedInt,omitempty"`
	// Maximum Allowed Value (for some types)
	MaxValueQuantity *Quantity `json:"maxValueQuantity,omitempty"`
	// Max length for string type data
	MaxLength *int32 `json:"maxLength,omitempty"`
	// Reference to invariant about presence
	Condition []custom.ID `json:"condition,omitempty"`
	// Condition that must evaluate to true
	Constraint []ElementDefinitionConstraint `json:"constraint,omitempty"`
	// For primitives, that a value must be present - not replaced by an extension
	MustHaveValue *bool `json:"mustHaveValue,omitempty"`
	// Extensions that are allowed to replace a primitive value
	ValueAlternatives []custom.Canonical[StructureDefinition] `json:"valueAlternatives,omitempty"`
	// If the element must be supported (discouraged - see obligations)
	MustSupport *bool `json:"mustSupport,omitempty"`
	// If this modifies the meaning of other elements
	IsModifier *bool `json:"isModifier,omitempty"`
	// Reason that this element is marked as a modifier
	IsModifierReason *string `json:"isModifierReason,omitempty"`
	// Include when _summary = true?
	IsSummary *bool `json:"isSummary,omitempty"`
	// ValueSet details if this is coded
	Binding *ElementDefinitionBinding `json:"binding,omitempty"`
	// Map element to another set of definitions
	Mapping []ElementDefinitionMapping `json:"mapping,omitempty"`
}

type ElementDefinitionSlicingDiscriminator struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// value | exists | type | profile | position
	Type custom.Code `json:"type"`
	// Path to element value
	Path string `json:"path"`
}

type ElementDefinitionType struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Data type or Resource (reference to definition)
	Code custom.URI `json:"code"`
	// Profiles (StructureDefinition or IG) - one must apply
	Profile []custom.Canonical[ElementDefinitionTypeProfile] `json:"profile,omitempty"`
	// Profile (StructureDefinition or IG) on the Reference/canonical target - one must apply
	TargetProfile []custom.Canonical[ElementDefinitionTypeTargetProfile] `json:"targetProfile,omitempty"`
	// contained | referenced | bundled - how aggregated
	Aggregation []custom.Code `json:"aggregation,omitempty"`
	// either | independent | specific
	Versioning *custom.Code `json:"versioning,omitempty"`
}

type ElementDefinitionConstraint struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Target of 'condition' reference above
	Key custom.ID `json:"key"`
	// Why this constraint is necessary or appropriate
	Requirements *custom.Markdown `json:"requirements,omitempty"`
	// error | warning
	Severity custom.Code `json:"severity"`
	// Suppress warning or hint in profile
	Suppress *bool `json:"suppress,omitempty"`
	// Human description of constraint
	Human string `json:"human"`
	// FHIRPath expression of constraint
	Expression *string `json:"expression,omitempty"`
	// Reference to original source of constraint
	Source *custom.Canonical[StructureDefinition] `json:"source,omitempty"`
}

type OtherElementDefinition ElementDefinition

type ElementDefinitionTypeProfile interface {
	gofhirclient.Resource

	Is_ElementDefinitionTypeProfile()
}

func (s StructureDefinition) Is_ElementDefinitionTypeProfile() {}
func (i ImplementationGuide) Is_ElementDefinitionTypeProfile() {}

type ElementDefinitionTypeTargetProfile interface {
	gofhirclient.Resource

	Is_ElementDefinitionTypeTargetProfile()
}

func (s StructureDefinition) Is_ElementDefinitionTypeTargetProfile() {}
func (i ImplementationGuide) Is_ElementDefinitionTypeTargetProfile() {}

func (e ElementDefinition) ResourceType() string {
	return "ElementDefinition"
}

func (e ElementDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherElementDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherElementDefinition: OtherElementDefinition(e), ResourceType: e.ResourceType()})
}

func UnmarshalElementDefinition(b []byte) (res ElementDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
