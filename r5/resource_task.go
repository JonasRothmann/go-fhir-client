// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Task
type TaskRestriction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// How many times to repeat
	Repetitions *uint32 `json:"repetitions,omitempty"`
	// When fulfillment is sought
	Period *Period `json:"period,omitempty"`
	// For whom is fulfillment sought?
	Recipient []custom.Reference[TaskRestrictionRecipient] `json:"recipient,omitempty"`
}

type TaskInput struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Label for the input
	Type CodeableConcept `json:"type"`
	// Content to use in performing the task
	ValueBase64binary *custom.Base64binary `json:"valueBase64binary,omitempty"`
	// Content to use in performing the task
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// Content to use in performing the task
	ValueCanonical *custom.Canonical[any] `json:"valueCanonical,omitempty"`
	// Content to use in performing the task
	ValueCode *custom.Code `json:"valueCode,omitempty"`
	// Content to use in performing the task
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// Content to use in performing the task
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// Content to use in performing the task
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// Content to use in performing the task
	ValueId *custom.ID `json:"valueId,omitempty"`
	// Content to use in performing the task
	ValueInstant *custom.Instant `json:"valueInstant,omitempty"`
	// Content to use in performing the task
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// Content to use in performing the task
	ValueInteger64 *int64 `json:"valueInteger64,omitempty"`
	// Content to use in performing the task
	ValueMarkdown *custom.Markdown `json:"valueMarkdown,omitempty"`
	// Content to use in performing the task
	ValueOid *custom.OID `json:"valueOid,omitempty"`
	// Content to use in performing the task
	ValuePositiveInt *uint32 `json:"valuePositiveInt,omitempty"`
	// Content to use in performing the task
	ValueString *string `json:"valueString,omitempty"`
	// Content to use in performing the task
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// Content to use in performing the task
	ValueUnsignedInt *uint32 `json:"valueUnsignedInt,omitempty"`
	// Content to use in performing the task
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// Content to use in performing the task
	ValueUrl *custom.URL `json:"valueUrl,omitempty"`
	// Content to use in performing the task
	ValueUuid *custom.UUID `json:"valueUuid,omitempty"`
	// Content to use in performing the task
	ValueAddress *Address `json:"valueAddress,omitempty"`
	// Content to use in performing the task
	ValueAge *Age `json:"valueAge,omitempty"`
	// Content to use in performing the task
	ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`
	// Content to use in performing the task
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// Content to use in performing the task
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// Content to use in performing the task
	ValueCodeableReference *custom.CodeableReference[gofhirclient.Resource] `json:"valueCodeableReference,omitempty"`
	// Content to use in performing the task
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// Content to use in performing the task
	ValueContactPoint *ContactPoint `json:"valueContactPoint,omitempty"`
	// Content to use in performing the task
	ValueCount *Count `json:"valueCount,omitempty"`
	// Content to use in performing the task
	ValueDistance *Distance `json:"valueDistance,omitempty"`
	// Content to use in performing the task
	ValueDuration *Duration `json:"valueDuration,omitempty"`
	// Content to use in performing the task
	ValueHumanName *HumanName `json:"valueHumanName,omitempty"`
	// Content to use in performing the task
	ValueIdentifier *Identifier `json:"valueIdentifier,omitempty"`
	// Content to use in performing the task
	ValueMoney *Money `json:"valueMoney,omitempty"`
	// Content to use in performing the task
	ValuePeriod *Period `json:"valuePeriod,omitempty"`
	// Content to use in performing the task
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// Content to use in performing the task
	ValueRange *Range `json:"valueRange,omitempty"`
	// Content to use in performing the task
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// Content to use in performing the task
	ValueRatioRange *RatioRange `json:"valueRatioRange,omitempty"`
	// Content to use in performing the task
	ValueReference *custom.Reference[gofhirclient.Resource] `json:"valueReference,omitempty"`
	// Content to use in performing the task
	ValueSampledData *SampledData `json:"valueSampledData,omitempty"`
	// Content to use in performing the task
	ValueSignature *Signature `json:"valueSignature,omitempty"`
	// Content to use in performing the task
	ValueTiming *Timing `json:"valueTiming,omitempty"`
	// Content to use in performing the task
	ValueContactDetail *ContactDetail `json:"valueContactDetail,omitempty"`
	// Content to use in performing the task
	ValueDataRequirement *DataRequirement `json:"valueDataRequirement,omitempty"`
	// Content to use in performing the task
	ValueExpression *Expression `json:"valueExpression,omitempty"`
	// Content to use in performing the task
	ValueParameterDefinition *ParameterDefinition `json:"valueParameterDefinition,omitempty"`
	// Content to use in performing the task
	ValueRelatedArtifact *RelatedArtifact `json:"valueRelatedArtifact,omitempty"`
	// Content to use in performing the task
	ValueTriggerDefinition *TriggerDefinition `json:"valueTriggerDefinition,omitempty"`
	// Content to use in performing the task
	ValueUsageContext *UsageContext `json:"valueUsageContext,omitempty"`
	// Content to use in performing the task
	ValueAvailability *Availability `json:"valueAvailability,omitempty"`
	// Content to use in performing the task
	ValueExtendedContactDetail *ExtendedContactDetail `json:"valueExtendedContactDetail,omitempty"`
	// Content to use in performing the task
	ValueDosage *Dosage `json:"valueDosage,omitempty"`
	// Content to use in performing the task
	ValueMeta *Meta `json:"valueMeta,omitempty"`
}

type TaskOutput struct {
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

type Task struct {
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
	// Task Instance Identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// Formal definition of task
	InstantiatesCanonical *custom.Canonical[ActivityDefinition] `json:"instantiatesCanonical,omitempty"`
	// Formal definition of task
	InstantiatesUri *custom.URI `json:"instantiatesUri,omitempty"`
	// Request fulfilled by this task
	BasedOn []custom.Reference[Resource] `json:"basedOn,omitempty"`
	// Requisition or grouper id
	GroupIdentifier *Identifier `json:"groupIdentifier,omitempty"`
	// Composite task
	PartOf []custom.Reference[Task] `json:"partOf,omitempty"`
	// draft | requested | received | accepted | +
	Status custom.Code `json:"status"`
	// Reason for current status
	StatusReason *custom.CodeableReference[gofhirclient.Resource] `json:"statusReason,omitempty"`
	// E.g. "Specimen collected", "IV prepped"
	BusinessStatus *CodeableConcept `json:"businessStatus,omitempty"`
	// unknown | proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option
	Intent custom.Code `json:"intent"`
	// routine | urgent | asap | stat
	Priority *custom.Code `json:"priority,omitempty"`
	// True if Task is prohibiting action
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	// Task Type
	Code *CodeableConcept `json:"code,omitempty"`
	// Human-readable explanation of task
	Description *string `json:"description,omitempty"`
	// What task is acting on
	Focus *custom.Reference[Resource] `json:"focus,omitempty"`
	// Beneficiary of the Task
	For *custom.Reference[Resource] `json:"for,omitempty"`
	// Healthcare event during which this task originated
	Encounter *custom.Reference[Encounter] `json:"encounter,omitempty"`
	// When the task should be performed
	RequestedPeriod *Period `json:"requestedPeriod,omitempty"`
	// Start and end time of execution
	ExecutionPeriod *Period `json:"executionPeriod,omitempty"`
	// Task Creation Date
	AuthoredOn *custom.DateTime `json:"authoredOn,omitempty"`
	// Task Last Modified Date
	LastModified *custom.DateTime `json:"lastModified,omitempty"`
	// Who is asking for task to be done
	Requester *custom.Reference[TaskRequester] `json:"requester,omitempty"`
	// Who should perform Task
	RequestedPerformer []custom.CodeableReference[TaskRequestedPerformer] `json:"requestedPerformer,omitempty"`
	// Responsible individual
	Owner *custom.Reference[TaskOwner] `json:"owner,omitempty"`
	// Who or what performed the task
	Performer []TaskPerformer `json:"performer,omitempty"`
	// Where task occurs
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// Why task is needed
	Reason []custom.CodeableReference[gofhirclient.Resource] `json:"reason,omitempty"`
	// Associated insurance coverage
	Insurance []custom.Reference[TaskInsurance] `json:"insurance,omitempty"`
	// Comments made about the task
	Note []Annotation `json:"note,omitempty"`
	// Key events in history of the Task
	RelevantHistory []custom.Reference[Provenance] `json:"relevantHistory,omitempty"`
	// Constraints on fulfillment tasks
	Restriction *TaskRestriction `json:"restriction,omitempty"`
	// Information used to perform task
	Input []TaskInput `json:"input,omitempty"`
	// Information produced as part of task
	Output []TaskOutput `json:"output,omitempty"`
}

type TaskPerformer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Type of performance
	Function *CodeableConcept `json:"function,omitempty"`
	// Who performed the task
	Actor custom.Reference[TaskPerformerActor] `json:"actor"`
}

type OtherTask Task

type TaskRequester interface {
	gofhirclient.Resource

	Is_TaskRequester()
}

func (d Device) Is_TaskRequester()           {}
func (o Organization) Is_TaskRequester()     {}
func (p Patient) Is_TaskRequester()          {}
func (p Practitioner) Is_TaskRequester()     {}
func (p PractitionerRole) Is_TaskRequester() {}
func (r RelatedPerson) Is_TaskRequester()    {}

type TaskRequestedPerformer interface {
	gofhirclient.Resource

	Is_TaskRequestedPerformer()
}

func (p Practitioner) Is_TaskRequestedPerformer()      {}
func (p PractitionerRole) Is_TaskRequestedPerformer()  {}
func (o Organization) Is_TaskRequestedPerformer()      {}
func (c CareTeam) Is_TaskRequestedPerformer()          {}
func (h HealthcareService) Is_TaskRequestedPerformer() {}
func (p Patient) Is_TaskRequestedPerformer()           {}
func (d Device) Is_TaskRequestedPerformer()            {}
func (r RelatedPerson) Is_TaskRequestedPerformer()     {}

type TaskOwner interface {
	gofhirclient.Resource

	Is_TaskOwner()
}

func (p Practitioner) Is_TaskOwner()     {}
func (p PractitionerRole) Is_TaskOwner() {}
func (o Organization) Is_TaskOwner()     {}
func (c CareTeam) Is_TaskOwner()         {}
func (p Patient) Is_TaskOwner()          {}
func (r RelatedPerson) Is_TaskOwner()    {}

type TaskPerformerActor interface {
	gofhirclient.Resource

	Is_TaskPerformerActor()
}

func (p Practitioner) Is_TaskPerformerActor()     {}
func (p PractitionerRole) Is_TaskPerformerActor() {}
func (o Organization) Is_TaskPerformerActor()     {}
func (c CareTeam) Is_TaskPerformerActor()         {}
func (p Patient) Is_TaskPerformerActor()          {}
func (r RelatedPerson) Is_TaskPerformerActor()    {}

type TaskInsurance interface {
	gofhirclient.Resource

	Is_TaskInsurance()
}

func (c Coverage) Is_TaskInsurance()      {}
func (c ClaimResponse) Is_TaskInsurance() {}

type TaskRestrictionRecipient interface {
	gofhirclient.Resource

	Is_TaskRestrictionRecipient()
}

func (p Patient) Is_TaskRestrictionRecipient()          {}
func (p Practitioner) Is_TaskRestrictionRecipient()     {}
func (p PractitionerRole) Is_TaskRestrictionRecipient() {}
func (r RelatedPerson) Is_TaskRestrictionRecipient()    {}
func (g Group) Is_TaskRestrictionRecipient()            {}
func (o Organization) Is_TaskRestrictionRecipient()     {}

func (t Task) ResourceType() string {
	return "Task"
}

func (t Task) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTask
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherTask: OtherTask(t), ResourceType: t.ResourceType()})
}

func UnmarshalTask(b []byte) (res Task, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
