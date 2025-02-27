// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Transport
type TransportInput struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for the input
	Type CodeableConcept `json:"type"`
	// Content to use in performing the transport
	ValueBase64binary *custom.Base64binary `json:"valueBase64binary,omitempty"`
	// Content to use in performing the transport
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Content to use in performing the transport
	ValueCanonical *custom.Canonical[any] `json:"valueCanonical,omitempty"`
	// Content to use in performing the transport
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Content to use in performing the transport
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Content to use in performing the transport
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Content to use in performing the transport
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Content to use in performing the transport
	ValueId *custom.ID `json:"valueId,omitempty"`
	// Content to use in performing the transport
	ValueInstant *custom.Instant `json:"valueInstant,omitempty"`
	// Content to use in performing the transport
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Content to use in performing the transport
	ValueInteger64 *int64 `json:"valueInteger64,omitempty"`
	// Content to use in performing the transport
	ValueMarkdown *custom.Markdown `json:"valueMarkdown,omitempty"`
	// Content to use in performing the transport
	ValueOid *custom.OID `json:"valueOid,omitempty"`
	// Content to use in performing the transport
	ValuePositiveInt *uint32 `json:"valuePositiveInt,omitempty"`
	// Content to use in performing the transport
	ValueString *string `json:"valueString,omitempty"`
	// Content to use in performing the transport
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Content to use in performing the transport
	ValueUnsignedInt *uint32 `json:"valueUnsignedInt,omitempty"`
	// Content to use in performing the transport
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Content to use in performing the transport
	ValueUrl *custom.URL `json:"valueUrl,omitempty"`
	// Content to use in performing the transport
	ValueUuid *custom.UUID `json:"valueUuid,omitempty"`
	// Content to use in performing the transport
	ValueAddress *Address `json:"valueAddress,omitempty"`
	// Content to use in performing the transport
	ValueAge *Age `json:"valueAge,omitempty"`
	// Content to use in performing the transport
	ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`
	// Content to use in performing the transport
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Content to use in performing the transport
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Content to use in performing the transport
	ValueCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"valueCodeableReference,omitempty"`
	// Content to use in performing the transport
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Content to use in performing the transport
	ValueContactPoint *ContactPoint `json:"valueContactPoint,omitempty"`
	// Content to use in performing the transport
	ValueCount *Count `json:"valueCount,omitempty"`
	// Content to use in performing the transport
	ValueDistance *Distance `json:"valueDistance,omitempty"`
	// Content to use in performing the transport
	ValueDuration *Duration `json:"valueDuration,omitempty"`
	// Content to use in performing the transport
	ValueHumanName *HumanName `json:"valueHumanName,omitempty"`
	// Content to use in performing the transport
	ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`
	// Content to use in performing the transport
	ValueMoney *Money `json:"valueMoney,omitempty"`
	// Content to use in performing the transport
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Content to use in performing the transport
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Content to use in performing the transport
	ValueRange *Range `json:"valueRange,omitempty"`
	// Content to use in performing the transport
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Content to use in performing the transport
	ValueRatioRange *RatioRange `json:"valueRatioRange,omitempty"`
	// Content to use in performing the transport
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// Content to use in performing the transport
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Content to use in performing the transport
	ValueSignature *Signature `json:"valueSignature,omitempty"`
	// Content to use in performing the transport
	ValueTiming *Timing `json:"valueTiming,omitempty"`
	// Content to use in performing the transport
	ValueContactDetail *ContactDetail `json:"valueContactDetail,omitempty"`
	// Content to use in performing the transport
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement,omitempty"`
	// Content to use in performing the transport
	ValueExpression *Expression `json:"valueExpression,omitempty"`
	// Content to use in performing the transport
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	// Content to use in performing the transport
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact,omitempty"`
	// Content to use in performing the transport
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition,omitempty"`
	// Content to use in performing the transport
	ValueUsageContext *UsageContext `json:"valueUsageContext,omitempty"`
	// Content to use in performing the transport
	ValueAvailability *Availability `json:"valueAvailability,omitempty"`
	// Content to use in performing the transport
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty"`
	// Content to use in performing the transport
	ValueDosage *Dosage `json:"valueDosage,omitempty"`
	// Content to use in performing the transport
	ValueMeta *Meta `json:"valueMeta,omitempty"`
}

type TransportOutput struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for output
	Type CodeableConcept `json:"type"`
	// Result of output
	ValueBase64binary *custom.Base64binary `json:"valueBase64binary,omitempty"`
	// Result of output
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Result of output
	ValueCanonical *custom.Canonical[any] `json:"valueCanonical,omitempty"`
	// Result of output
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Result of output
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Result of output
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Result of output
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Result of output
	ValueId *custom.ID `json:"valueId,omitempty"`
	// Result of output
	ValueInstant *custom.Instant `json:"valueInstant,omitempty"`
	// Result of output
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Result of output
	ValueInteger64 *int64 `json:"valueInteger64,omitempty"`
	// Result of output
	ValueMarkdown *custom.Markdown `json:"valueMarkdown,omitempty"`
	// Result of output
	ValueOid *custom.OID `json:"valueOid,omitempty"`
	// Result of output
	ValuePositiveInt *uint32 `json:"valuePositiveInt,omitempty"`
	// Result of output
	ValueString *string `json:"valueString,omitempty"`
	// Result of output
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Result of output
	ValueUnsignedInt *uint32 `json:"valueUnsignedInt,omitempty"`
	// Result of output
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Result of output
	ValueUrl *custom.URL `json:"valueUrl,omitempty"`
	// Result of output
	ValueUuid *custom.UUID `json:"valueUuid,omitempty"`
	// Result of output
	ValueAddress *Address `json:"valueAddress,omitempty"`
	// Result of output
	ValueAge *Age `json:"valueAge,omitempty"`
	// Result of output
	ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`
	// Result of output
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Result of output
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Result of output
	ValueCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"valueCodeableReference,omitempty"`
	// Result of output
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Result of output
	ValueContactPoint *ContactPoint `json:"valueContactPoint,omitempty"`
	// Result of output
	ValueCount *Count `json:"valueCount,omitempty"`
	// Result of output
	ValueDistance *Distance `json:"valueDistance,omitempty"`
	// Result of output
	ValueDuration *Duration `json:"valueDuration,omitempty"`
	// Result of output
	ValueHumanName *HumanName `json:"valueHumanName,omitempty"`
	// Result of output
	ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`
	// Result of output
	ValueMoney *Money `json:"valueMoney,omitempty"`
	// Result of output
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Result of output
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Result of output
	ValueRange *Range `json:"valueRange,omitempty"`
	// Result of output
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Result of output
	ValueRatioRange *RatioRange `json:"valueRatioRange,omitempty"`
	// Result of output
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// Result of output
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Result of output
	ValueSignature *Signature `json:"valueSignature,omitempty"`
	// Result of output
	ValueTiming *Timing `json:"valueTiming,omitempty"`
	// Result of output
	ValueContactDetail *ContactDetail `json:"valueContactDetail,omitempty"`
	// Result of output
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement,omitempty"`
	// Result of output
	ValueExpression *Expression `json:"valueExpression,omitempty"`
	// Result of output
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	// Result of output
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact,omitempty"`
	// Result of output
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition,omitempty"`
	// Result of output
	ValueUsageContext *UsageContext `json:"valueUsageContext,omitempty"`
	// Result of output
	ValueAvailability *Availability `json:"valueAvailability,omitempty"`
	// Result of output
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty"`
	// Result of output
	ValueDosage *Dosage `json:"valueDosage,omitempty"`
	// Result of output
	ValueMeta *Meta `json:"valueMeta,omitempty"`
}

type Transport struct {
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
	// External identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Formal definition of transport
	InstantiatesCanonical *custom.Canonical[ActivityDefinition] `json:"instantiatesCanonical,omitempty"`
	// Formal definition of transport
	InstantiatesUri *custom.URI `json:"instantiatesUri,omitempty"`
	// Request fulfilled by this transport
	BasedOn []custom.Reference[Resource] `json:"basedOn,omitempty"`
	// Requisition or grouper id
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	// Part of referenced event
	PartOf []custom.Reference[Transport] `json:"partOf,omitempty"`
	// in-progress | completed | abandoned | cancelled | planned | entered-in-error
	Status *custom.Code `json:"status,omitempty"`
	// Reason for current status
	StatusReason *CodeableConcept `json:"statusReason,omitempty"`
	// unknown | proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// Transport Type
	Code *CodeableConcept `json:"code,omitempty"`
	// Human-readable explanation of transport
	Description *string `json:"description,omitempty"`
	// What transport is acting on
	Focus *custom.Reference[Resource] `json:"focus,omitempty"`
	// Beneficiary of the Transport
	For *custom.Reference[Resource] `json:"for,omitempty"`
	// Healthcare event during which this transport originated
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// Completion time of the event (the occurrence)
	CompletionTime *custom.DateTime `json:"completionTime,omitempty"`
	// Transport Creation Date
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Transport Last Modified Date
	LastModified *custom.DateTime `json:"lastModified,omitempty"`
	// Who is asking for transport to be done
	Requester *custom.Reference[TransportRequester] `json:"requester,omitempty"`
	// Requested performer
	PerformerType []CodeableConcept `json:"performerType,omitempty"`
	// Responsible individual
	Owner *custom.Reference[TransportOwner] `json:"owner,omitempty"`
	// Where transport occurs
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[TransportInsurance] `json:"insurance,omitempty"`
	// Comments made about the transport
	Note []Annotation `json:"note,omitempty"`
	// Key events in history of the Transport
	RelevantHistory []custom.Reference[Provenance] `json:"relevantHistory,omitempty"`
	// Constraints on fulfillment transports
	Restriction *TransportRestriction `json:"restriction,omitempty"`
	// Information used to perform transport
	Input []TransportInput `json:"input,omitempty"`
	// Information produced as part of transport
	Output []TransportOutput `json:"output,omitempty"`
	// The desired location
	RequestedLocation custom.Reference[Location] `json:"requestedLocation"`
	// The entity current location
	CurrentLocation custom.Reference[Location] `json:"currentLocation"`
	// Why transport is needed
	Reason *custom.CodeableReference[Resource] `json:"reason,omitempty"`
	// Parent (or preceding) transport
	History *custom.Reference[Transport] `json:"history,omitempty"`
}

type TransportRestriction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// How many times to repeat
	Repetitions *uint32 `json:"repetitions,omitempty"`
	// When fulfillment sought
	Period *Period `json:"period,omitempty"`
	// For whom is fulfillment sought?
	Recipient []custom.Reference[TransportRestrictionRecipient] `json:"recipient,omitempty"`
}

type OtherTransport Transport

type TransportRequester interface {
	gofhirclient.Resource

	Is_TransportRequester()
}

func (d Device) Is_TransportRequester()           {}
func (o Organization) Is_TransportRequester()     {}
func (p Patient) Is_TransportRequester()          {}
func (p Practitioner) Is_TransportRequester()     {}
func (p PractitionerRole) Is_TransportRequester() {}
func (r RelatedPerson) Is_TransportRequester()    {}

type TransportOwner interface {
	gofhirclient.Resource

	Is_TransportOwner()
}

func (p Practitioner) Is_TransportOwner()      {}
func (p PractitionerRole) Is_TransportOwner()  {}
func (o Organization) Is_TransportOwner()      {}
func (c CareTeam) Is_TransportOwner()          {}
func (h HealthcareService) Is_TransportOwner() {}
func (p Patient) Is_TransportOwner()           {}
func (d Device) Is_TransportOwner()            {}
func (r RelatedPerson) Is_TransportOwner()     {}

type TransportInsurance interface {
	gofhirclient.Resource

	Is_TransportInsurance()
}

func (c Coverage) Is_TransportInsurance()      {}
func (c ClaimResponse) Is_TransportInsurance() {}

type TransportRestrictionRecipient interface {
	gofhirclient.Resource

	Is_TransportRestrictionRecipient()
}

func (p Patient) Is_TransportRestrictionRecipient()          {}
func (p Practitioner) Is_TransportRestrictionRecipient()     {}
func (p PractitionerRole) Is_TransportRestrictionRecipient() {}
func (r RelatedPerson) Is_TransportRestrictionRecipient()    {}
func (g Group) Is_TransportRestrictionRecipient()            {}
func (o Organization) Is_TransportRestrictionRecipient()     {}

func (t Transport) ResourceType() string {
	return "Transport"
}

func (t Transport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTransport
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherTransport: OtherTransport(t), ResourceType: t.ResourceType()})
}

func UnmarshalTransport(b []byte) (res Transport, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
