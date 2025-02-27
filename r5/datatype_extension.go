// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Extension
type Extension struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// identifies the meaning of the extension
	Url string `json:"url"`
	// Value of extension
	ValueBase64binary *custom.Base64binary `json:"valueBase64binary,omitempty"`
	// Value of extension
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Value of extension
	ValueCanonical *custom.Canonical[any] `json:"valueCanonical,omitempty"`
	// Value of extension
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Value of extension
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Value of extension
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Value of extension
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Value of extension
	ValueId *custom.ID `json:"valueId,omitempty"`
	// Value of extension
	ValueInstant *custom.Instant `json:"valueInstant,omitempty"`
	// Value of extension
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Value of extension
	ValueInteger64 *int64 `json:"valueInteger64,omitempty"`
	// Value of extension
	ValueMarkdown *custom.Markdown `json:"valueMarkdown,omitempty"`
	// Value of extension
	ValueOid *custom.OID `json:"valueOid,omitempty"`
	// Value of extension
	ValuePositiveInt *uint32 `json:"valuePositiveInt,omitempty"`
	// Value of extension
	ValueString *string `json:"valueString,omitempty"`
	// Value of extension
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Value of extension
	ValueUnsignedInt *uint32 `json:"valueUnsignedInt,omitempty"`
	// Value of extension
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Value of extension
	ValueUrl *custom.URL `json:"valueUrl,omitempty"`
	// Value of extension
	ValueUuid *custom.UUID `json:"valueUuid,omitempty"`
	// Value of extension
	ValueAddress *Address `json:"valueAddress,omitempty"`
	// Value of extension
	ValueAge *Age `json:"valueAge,omitempty"`
	// Value of extension
	ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`
	// Value of extension
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Value of extension
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Value of extension
	ValueCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"valueCodeableReference,omitempty"`
	// Value of extension
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Value of extension
	ValueContactPoint *ContactPoint `json:"valueContactPoint,omitempty"`
	// Value of extension
	ValueCount *Count `json:"valueCount,omitempty"`
	// Value of extension
	ValueDistance *Distance `json:"valueDistance,omitempty"`
	// Value of extension
	ValueDuration *Duration `json:"valueDuration,omitempty"`
	// Value of extension
	ValueHumanName *HumanName `json:"valueHumanName,omitempty"`
	// Value of extension
	ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`
	// Value of extension
	ValueMoney *Money `json:"valueMoney,omitempty"`
	// Value of extension
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Value of extension
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Value of extension
	ValueRange *Range `json:"valueRange,omitempty"`
	// Value of extension
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Value of extension
	ValueRatioRange *RatioRange `json:"valueRatioRange,omitempty"`
	// Value of extension
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// Value of extension
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Value of extension
	ValueSignature *Signature `json:"valueSignature,omitempty"`
	// Value of extension
	ValueTiming *Timing `json:"valueTiming,omitempty"`
	// Value of extension
	ValueContactDetail *ContactDetail `json:"valueContactDetail,omitempty"`
	// Value of extension
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement,omitempty"`
	// Value of extension
	ValueExpression *Expression `json:"valueExpression,omitempty"`
	// Value of extension
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	// Value of extension
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact,omitempty"`
	// Value of extension
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition,omitempty"`
	// Value of extension
	ValueUsageContext *UsageContext `json:"valueUsageContext,omitempty"`
	// Value of extension
	ValueAvailability *Availability `json:"valueAvailability,omitempty"`
	// Value of extension
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty"`
	// Value of extension
	ValueDosage *Dosage `json:"valueDosage,omitempty"`
	// Value of extension
	ValueMeta *Meta `json:"valueMeta,omitempty"`
}

type OtherExtension Extension

func (e Extension) ResourceType() string {
	return "Extension"
}

func (e Extension) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherExtension
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherExtension: OtherExtension(e), ResourceType: e.ResourceType()})
}

func UnmarshalExtension(b []byte) (res Extension, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
