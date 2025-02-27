// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Parameters
type Parameters struct {
	// Logical id of this artifact
	Id *string `json:"id,omitempty"`
	// Metadata about the resource
	Meta *Meta `json:"meta,omitempty"`
	// A set of rules under which this content was created
	ImplicitRules *custom.URI `json:"implicitRules,omitempty"`
	// Language of the resource content
	Language *custom.Code `json:"language,omitempty"`
	// Operation Parameter
	Parameter []ParametersParameter `json:"parameter,omitempty"`
}

type ParametersParameter struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Name from the definition
	Name string `json:"name"`
	// If parameter is a data type
	ValueBase64binary *custom.Base64binary `json:"valueBase64binary,omitempty"`
	// If parameter is a data type
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// If parameter is a data type
	ValueCanonical *custom.Canonical[any] `json:"valueCanonical,omitempty"`
	// If parameter is a data type
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// If parameter is a data type
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// If parameter is a data type
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// If parameter is a data type
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// If parameter is a data type
	ValueId *custom.ID `json:"valueId,omitempty"`
	// If parameter is a data type
	ValueInstant *custom.Instant `json:"valueInstant,omitempty"`
	// If parameter is a data type
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// If parameter is a data type
	ValueInteger64 *int64 `json:"valueInteger64,omitempty"`
	// If parameter is a data type
	ValueMarkdown *custom.Markdown `json:"valueMarkdown,omitempty"`
	// If parameter is a data type
	ValueOid *custom.OID `json:"valueOid,omitempty"`
	// If parameter is a data type
	ValuePositiveInt *uint32 `json:"valuePositiveInt,omitempty"`
	// If parameter is a data type
	ValueString *string `json:"valueString,omitempty"`
	// If parameter is a data type
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// If parameter is a data type
	ValueUnsignedInt *uint32 `json:"valueUnsignedInt,omitempty"`
	// If parameter is a data type
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// If parameter is a data type
	ValueUrl *custom.URL `json:"valueUrl,omitempty"`
	// If parameter is a data type
	ValueUuid *custom.UUID `json:"valueUuid,omitempty"`
	// If parameter is a data type
	ValueAddress *Address `json:"valueAddress,omitempty"`
	// If parameter is a data type
	ValueAge *Age `json:"valueAge,omitempty"`
	// If parameter is a data type
	ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`
	// If parameter is a data type
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// If parameter is a data type
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// If parameter is a data type
	ValueCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"valueCodeableReference,omitempty"`
	// If parameter is a data type
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// If parameter is a data type
	ValueContactPoint *ContactPoint `json:"valueContactPoint,omitempty"`
	// If parameter is a data type
	ValueCount *Count `json:"valueCount,omitempty"`
	// If parameter is a data type
	ValueDistance *Distance `json:"valueDistance,omitempty"`
	// If parameter is a data type
	ValueDuration *Duration `json:"valueDuration,omitempty"`
	// If parameter is a data type
	ValueHumanName *HumanName `json:"valueHumanName,omitempty"`
	// If parameter is a data type
	ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`
	// If parameter is a data type
	ValueMoney *Money `json:"valueMoney,omitempty"`
	// If parameter is a data type
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// If parameter is a data type
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// If parameter is a data type
	ValueRange *Range `json:"valueRange,omitempty"`
	// If parameter is a data type
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// If parameter is a data type
	ValueRatioRange *RatioRange `json:"valueRatioRange,omitempty"`
	// If parameter is a data type
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// If parameter is a data type
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// If parameter is a data type
	ValueSignature *Signature `json:"valueSignature,omitempty"`
	// If parameter is a data type
	ValueTiming *Timing `json:"valueTiming,omitempty"`
	// If parameter is a data type
	ValueContactDetail *ContactDetail `json:"valueContactDetail,omitempty"`
	// If parameter is a data type
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement,omitempty"`
	// If parameter is a data type
	ValueExpression *Expression `json:"valueExpression,omitempty"`
	// If parameter is a data type
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	// If parameter is a data type
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact,omitempty"`
	// If parameter is a data type
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition,omitempty"`
	// If parameter is a data type
	ValueUsageContext *UsageContext `json:"valueUsageContext,omitempty"`
	// If parameter is a data type
	ValueAvailability *Availability `json:"valueAvailability,omitempty"`
	// If parameter is a data type
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty"`
	// If parameter is a data type
	ValueDosage *Dosage `json:"valueDosage,omitempty"`
	// If parameter is a data type
	ValueMeta *Meta `json:"valueMeta,omitempty"`
	// If parameter is a whole resource
	Resource *Resource `json:"resource,omitempty"`
	// Named part of a multi-part parameter
	Part []interface{} `json:"part,omitempty"`
}

type OtherParameters Parameters

func (p Parameters) ResourceType() string {
	return "Parameters"
}

func (p Parameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherParameters
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherParameters: OtherParameters(p), ResourceType: p.ResourceType()})
}

func UnmarshalParameters(b []byte) (res Parameters, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
