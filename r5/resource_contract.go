// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Contract
type ContractSigner struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Contract Signatory Role
	Type Coding `json:"type"`
	// Contract Signatory Party
	Party custom.Reference[ContractSignerParty] `json:"party"`
	// Contract Documentation Signature
	Signature []Signature `json:"signature"`
}

type ContractFriendly struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Easily comprehended representation of this Contract
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`
	// Easily comprehended representation of this Contract
	ContentReference *custom.Reference[ContractFriendlyContentReference] `json:"contentReference,omitempty"`
}

type ContractLegal struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Contract Legal Text
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`
	// Contract Legal Text
	ContentReference *custom.Reference[ContractLegalContentReference] `json:"contentReference,omitempty"`
}

type Contract struct {
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
	// Contract number
	Identifier []Identifier `json:"identifier,omitempty"`
	// Basal definition
	Url *custom.URI `json:"url,omitempty"`
	// Business edition
	Version *string `json:"version,omitempty"`
	// amended | appended | cancelled | disputed | entered-in-error | executable +
	Status *custom.Code `json:"status,omitempty"`
	// Negotiation status
	LegalState *CodeableConcept `json:"legalState,omitempty"`
	// Source Contract Definition
	InstantiatesCanonical *custom.Reference[Contract] `json:"instantiatesCanonical,omitempty"`
	// External Contract Definition
	InstantiatesUri *custom.URI `json:"instantiatesUri,omitempty"`
	// Content derived from the basal information
	ContentDerivative *CodeableConcept `json:"contentDerivative,omitempty"`
	// When this Contract was issued
	Issued *custom.DateTime `json:"issued,omitempty"`
	// Effective time
	Applies *Period `json:"applies,omitempty"`
	// Contract cessation cause
	ExpirationType *CodeableConcept `json:"expirationType,omitempty"`
	// Contract Target Entity
	Subject []custom.Reference[Resource] `json:"subject,omitempty"`
	// Authority under which this Contract has standing
	Authority []custom.Reference[Organization] `json:"authority,omitempty"`
	// A sphere of control governed by an authoritative jurisdiction, organization, or person
	Domain []custom.Reference[Location] `json:"domain,omitempty"`
	// Specific Location
	Site []custom.Reference[Location] `json:"site,omitempty"`
	// Computer friendly designation
	Name *string `json:"name,omitempty"`
	// Human Friendly name
	Title *string `json:"title,omitempty"`
	// Subordinate Friendly name
	Subtitle *string `json:"subtitle,omitempty"`
	// Acronym or short name
	Alias []string `json:"alias,omitempty"`
	// Source of Contract
	Author *custom.Reference[ContractAuthor] `json:"author,omitempty"`
	// Range of Legal Concerns
	Scope *CodeableConcept `json:"scope,omitempty"`
	// Focus of contract interest
	TopicCodeableConcept *CodeableConcept `json:"topicCodeableConcept,omitempty"`
	// Focus of contract interest
	TopicReference *custom.Reference[Resource] `json:"topicReference,omitempty"`
	// Legal instrument category
	Type *CodeableConcept `json:"type,omitempty"`
	// Subtype within the context of type
	SubType []CodeableConcept `json:"subType,omitempty"`
	// Contract precursor content
	ContentDefinition *ContractContentDefinition `json:"contentDefinition,omitempty"`
	// Contract Term List
	Term []ContractTerm `json:"term,omitempty"`
	// Extra Information
	SupportingInfo []custom.Reference[Resource] `json:"supportingInfo,omitempty"`
	// Key event in Contract History
	RelevantHistory []custom.Reference[Provenance] `json:"relevantHistory,omitempty"`
	// Contract Signatory
	Signer []ContractSigner `json:"signer,omitempty"`
	// Contract Friendly Language
	Friendly []ContractFriendly `json:"friendly,omitempty"`
	// Contract Legal Language
	Legal []ContractLegal `json:"legal,omitempty"`
	// Computable Contract Language
	Rule []ContractRule `json:"rule,omitempty"`
	// Binding Contract
	LegallyBindingAttachment *Attachment `json:"legallyBindingAttachment,omitempty"`
	// Binding Contract
	LegallyBindingReference *custom.Reference[ContractLegallyBindingReference] `json:"legallyBindingReference,omitempty"`
}

type ContractContentDefinition struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Content structure and use
	Type CodeableConcept `json:"type"`
	// Detailed Content Type Definition
	SubType *CodeableConcept `json:"subType,omitempty"`
	// Publisher Entity
	Publisher *custom.Reference[ContractContentDefinitionPublisher] `json:"publisher,omitempty"`
	// When published
	PublicationDate *custom.DateTime `json:"publicationDate,omitempty"`
	// amended | appended | cancelled | disputed | entered-in-error | executable +
	PublicationStatus custom.Code `json:"publicationStatus"`
	// Publication Ownership
	Copyright *custom.Markdown `json:"copyright,omitempty"`
}

type ContractTermAction struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// True if the term prohibits the  action
	DoNotPerform *bool `json:"doNotPerform,omitempty"`
	// Type or form of the action
	Type CodeableConcept `json:"type"`
	// Entity of the action
	Subject []ContractTermActionSubject `json:"subject,omitempty"`
	// Purpose for the Contract Term Action
	Intent CodeableConcept `json:"intent"`
	// Pointer to specific item
	LinkId []string `json:"linkId,omitempty"`
	// State of the action
	Status CodeableConcept `json:"status"`
	// Episode associated with action
	Context *custom.Reference[ContractTermActionContext] `json:"context,omitempty"`
	// Pointer to specific item
	ContextLinkId []string `json:"contextLinkId,omitempty"`
	// When action happens
	OccurrenceDateTime *custom.DateTime `json:"occurrenceDateTime,omitempty"`
	// When action happens
	OccurrencePeriod *Period `json:"occurrencePeriod,omitempty"`
	// When action happens
	OccurrenceTiming *Timing `json:"occurrenceTiming,omitempty"`
	// Who asked for action
	Requester []custom.Reference[ContractTermActionRequester] `json:"requester,omitempty"`
	// Pointer to specific item
	RequesterLinkId []string `json:"requesterLinkId,omitempty"`
	// Kind of service performer
	PerformerType []CodeableConcept `json:"performerType,omitempty"`
	// Competency of the performer
	PerformerRole *CodeableConcept `json:"performerRole,omitempty"`
	// Actor that wil execute (or not) the action
	Performer *custom.Reference[ContractTermActionPerformer] `json:"performer,omitempty"`
	// Pointer to specific item
	PerformerLinkId []string `json:"performerLinkId,omitempty"`
	// Why is action (not) needed?
	Reason []custom.CodeableReference[ContractTermActionReason] `json:"reason,omitempty"`
	// Pointer to specific item
	ReasonLinkId []string `json:"reasonLinkId,omitempty"`
	// Comments about the action
	Note []Annotation `json:"note,omitempty"`
	// Action restriction numbers
	SecurityLabelNumber []uint32 `json:"securityLabelNumber,omitempty"`
}

type ContractTermAssetContext struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Creator,custodian or owner
	Reference *custom.Reference[Resource] `json:"reference,omitempty"`
	// Codeable asset context
	Code []CodeableConcept `json:"code,omitempty"`
	// Context description
	Text *string `json:"text,omitempty"`
}

type ContractTermOfferAnswer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The actual answer response
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// The actual answer response
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// The actual answer response
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// The actual answer response
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// The actual answer response
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// The actual answer response
	ValueTime *custom.Time `json:"valueTime,omitempty"`
	// The actual answer response
	ValueString *string `json:"valueString,omitempty"`
	// The actual answer response
	ValueUri *custom.URI `json:"valueUri,omitempty"`
	// The actual answer response
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// The actual answer response
	ValueCoding *Coding `json:"valueCoding,omitempty"`
	// The actual answer response
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The actual answer response
	ValueReference *custom.Reference[Resource] `json:"valueReference,omitempty"`
}

type ContractTermAssetValuedItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Contract Valued Item Type
	EntityCodeableConcept *CodeableConcept `json:"entityCodeableConcept,omitempty"`
	// Contract Valued Item Type
	EntityReference *custom.Reference[Resource] `json:"entityReference,omitempty"`
	// Contract Valued Item Number
	Identifier *Identifier `json:"identifier,omitempty"`
	// Contract Valued Item Effective Tiem
	EffectiveTime *custom.DateTime `json:"effectiveTime,omitempty"`
	// Count of Contract Valued Items
	Quantity *Quantity `json:"quantity,omitempty"`
	// Contract Valued Item fee, charge, or cost
	UnitPrice *Money `json:"unitPrice,omitempty"`
	// Contract Valued Item Price Scaling Factor
	Factor *json.Number `json:"factor,omitempty"`
	// Contract Valued Item Difficulty Scaling Factor
	Points *json.Number `json:"points,omitempty"`
	// Total Contract Valued Item Value
	Net *Money `json:"net,omitempty"`
	// Terms of valuation
	Payment *string `json:"payment,omitempty"`
	// When payment is due
	PaymentDate *custom.DateTime `json:"paymentDate,omitempty"`
	// Who will make payment
	Responsible *custom.Reference[ContractTermAssetValuedItemResponsible] `json:"responsible,omitempty"`
	// Who will receive payment
	Recipient *custom.Reference[ContractTermAssetValuedItemRecipient] `json:"recipient,omitempty"`
	// Pointer to specific item
	LinkId []string `json:"linkId,omitempty"`
	// Security Labels that define affected terms
	SecurityLabelNumber []uint32 `json:"securityLabelNumber,omitempty"`
}

type ContractTermSecurityLabel struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Link to Security Labels
	Number []uint32 `json:"number,omitempty"`
	// Confidentiality Protection
	Classification Coding `json:"classification"`
	// Applicable Policy
	Category []Coding `json:"category,omitempty"`
	// Handling Instructions
	Control []Coding `json:"control,omitempty"`
}

type ContractTermOffer struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Offer business ID
	Identifier []Identifier `json:"identifier,omitempty"`
	// Offer Recipient
	Party []ContractTermOfferParty `json:"party,omitempty"`
	// Negotiable offer asset
	Topic *custom.Reference[Resource] `json:"topic,omitempty"`
	// Contract Offer Type or Form
	Type *CodeableConcept `json:"type,omitempty"`
	// Accepting party choice
	Decision *CodeableConcept `json:"decision,omitempty"`
	// How decision is conveyed
	DecisionMode []CodeableConcept `json:"decisionMode,omitempty"`
	// Response to offer text
	Answer []ContractTermOfferAnswer `json:"answer,omitempty"`
	// Human readable offer text
	Text *string `json:"text,omitempty"`
	// Pointer to text
	LinkId []string `json:"linkId,omitempty"`
	// Offer restriction numbers
	SecurityLabelNumber []uint32 `json:"securityLabelNumber,omitempty"`
}

type ContractTermOfferParty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Referenced entity
	Reference []custom.Reference[ContractTermOfferPartyReference] `json:"reference"`
	// Participant engagement type
	Role CodeableConcept `json:"role"`
}

type ContractRule struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Computable Contract Rules
	ContentAttachment *Attachment `json:"contentAttachment,omitempty"`
	// Computable Contract Rules
	ContentReference *custom.Reference[DocumentReference] `json:"contentReference,omitempty"`
}

type ContractTerm struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Contract Term Number
	Identifier *Identifier `json:"identifier,omitempty"`
	// Contract Term Issue Date Time
	Issued *custom.DateTime `json:"issued,omitempty"`
	// Contract Term Effective Time
	Applies *Period `json:"applies,omitempty"`
	// Term Concern
	TopicCodeableConcept *CodeableConcept `json:"topicCodeableConcept,omitempty"`
	// Term Concern
	TopicReference *custom.Reference[Resource] `json:"topicReference,omitempty"`
	// Contract Term Type or Form
	Type *CodeableConcept `json:"type,omitempty"`
	// Contract Term Type specific classification
	SubType *CodeableConcept `json:"subType,omitempty"`
	// Term Statement
	Text *string `json:"text,omitempty"`
	// Protection for the Term
	SecurityLabel []ContractTermSecurityLabel `json:"securityLabel,omitempty"`
	// Context of the Contract term
	Offer ContractTermOffer `json:"offer"`
	// Contract Term Asset List
	Asset []ContractTermAsset `json:"asset,omitempty"`
	// Entity being ascribed responsibility
	Action []ContractTermAction `json:"action,omitempty"`
	// Nested Contract Term Group
	Group []interface{} `json:"group,omitempty"`
}

type ContractTermAsset struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Range of asset
	Scope *CodeableConcept `json:"scope,omitempty"`
	// Asset category
	Type []CodeableConcept `json:"type,omitempty"`
	// Associated entities
	TypeReference []custom.Reference[Resource] `json:"typeReference,omitempty"`
	// Asset sub-category
	Subtype []CodeableConcept `json:"subtype,omitempty"`
	// Kinship of the asset
	Relationship *Coding `json:"relationship,omitempty"`
	// Circumstance of the asset
	Context []ContractTermAssetContext `json:"context,omitempty"`
	// Quality desctiption of asset
	Condition *string `json:"condition,omitempty"`
	// Asset availability types
	PeriodType []CodeableConcept `json:"periodType,omitempty"`
	// Time period of the asset
	Period []Period `json:"period,omitempty"`
	// Time period
	UsePeriod []Period `json:"usePeriod,omitempty"`
	// Asset clause or question text
	Text *string `json:"text,omitempty"`
	// Pointer to asset text
	LinkId []string `json:"linkId,omitempty"`
	// Response to assets
	Answer []interface{} `json:"answer,omitempty"`
	// Asset restriction numbers
	SecurityLabelNumber []uint32 `json:"securityLabelNumber,omitempty"`
	// Contract Valued Item List
	ValuedItem []ContractTermAssetValuedItem `json:"valuedItem,omitempty"`
}

type ContractTermActionSubject struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Entity of the action
	Reference []custom.Reference[ContractTermActionSubjectReference] `json:"reference"`
	// Role type of the agent
	Role *CodeableConcept `json:"role,omitempty"`
}

type OtherContract Contract

type ContractAuthor interface {
	gofhirclient.Resource

	Is_ContractAuthor()
}

func (p Patient) Is_ContractAuthor()          {}
func (p Practitioner) Is_ContractAuthor()     {}
func (p PractitionerRole) Is_ContractAuthor() {}
func (o Organization) Is_ContractAuthor()     {}

type ContractContentDefinitionPublisher interface {
	gofhirclient.Resource

	Is_ContractContentDefinitionPublisher()
}

func (p Practitioner) Is_ContractContentDefinitionPublisher()     {}
func (p PractitionerRole) Is_ContractContentDefinitionPublisher() {}
func (o Organization) Is_ContractContentDefinitionPublisher()     {}

type ContractTermOfferPartyReference interface {
	gofhirclient.Resource

	Is_ContractTermOfferPartyReference()
}

func (p Patient) Is_ContractTermOfferPartyReference()          {}
func (r RelatedPerson) Is_ContractTermOfferPartyReference()    {}
func (p Practitioner) Is_ContractTermOfferPartyReference()     {}
func (p PractitionerRole) Is_ContractTermOfferPartyReference() {}
func (d Device) Is_ContractTermOfferPartyReference()           {}
func (g Group) Is_ContractTermOfferPartyReference()            {}
func (o Organization) Is_ContractTermOfferPartyReference()     {}

type ContractTermAssetValuedItemResponsible interface {
	gofhirclient.Resource

	Is_ContractTermAssetValuedItemResponsible()
}

func (o Organization) Is_ContractTermAssetValuedItemResponsible()     {}
func (p Patient) Is_ContractTermAssetValuedItemResponsible()          {}
func (p Practitioner) Is_ContractTermAssetValuedItemResponsible()     {}
func (p PractitionerRole) Is_ContractTermAssetValuedItemResponsible() {}
func (r RelatedPerson) Is_ContractTermAssetValuedItemResponsible()    {}

type ContractTermAssetValuedItemRecipient interface {
	gofhirclient.Resource

	Is_ContractTermAssetValuedItemRecipient()
}

func (o Organization) Is_ContractTermAssetValuedItemRecipient()     {}
func (p Patient) Is_ContractTermAssetValuedItemRecipient()          {}
func (p Practitioner) Is_ContractTermAssetValuedItemRecipient()     {}
func (p PractitionerRole) Is_ContractTermAssetValuedItemRecipient() {}
func (r RelatedPerson) Is_ContractTermAssetValuedItemRecipient()    {}

type ContractTermActionSubjectReference interface {
	gofhirclient.Resource

	Is_ContractTermActionSubjectReference()
}

func (p Patient) Is_ContractTermActionSubjectReference()          {}
func (r RelatedPerson) Is_ContractTermActionSubjectReference()    {}
func (p Practitioner) Is_ContractTermActionSubjectReference()     {}
func (p PractitionerRole) Is_ContractTermActionSubjectReference() {}
func (d Device) Is_ContractTermActionSubjectReference()           {}
func (g Group) Is_ContractTermActionSubjectReference()            {}
func (o Organization) Is_ContractTermActionSubjectReference()     {}

type ContractTermActionContext interface {
	gofhirclient.Resource

	Is_ContractTermActionContext()
}

func (e Encounter) Is_ContractTermActionContext()     {}
func (e EpisodeOfCare) Is_ContractTermActionContext() {}

type ContractTermActionRequester interface {
	gofhirclient.Resource

	Is_ContractTermActionRequester()
}

func (p Patient) Is_ContractTermActionRequester()          {}
func (r RelatedPerson) Is_ContractTermActionRequester()    {}
func (p Practitioner) Is_ContractTermActionRequester()     {}
func (p PractitionerRole) Is_ContractTermActionRequester() {}
func (d Device) Is_ContractTermActionRequester()           {}
func (g Group) Is_ContractTermActionRequester()            {}
func (o Organization) Is_ContractTermActionRequester()     {}

type ContractTermActionPerformer interface {
	gofhirclient.Resource

	Is_ContractTermActionPerformer()
}

func (r RelatedPerson) Is_ContractTermActionPerformer()    {}
func (p Patient) Is_ContractTermActionPerformer()          {}
func (p Practitioner) Is_ContractTermActionPerformer()     {}
func (p PractitionerRole) Is_ContractTermActionPerformer() {}
func (c CareTeam) Is_ContractTermActionPerformer()         {}
func (d Device) Is_ContractTermActionPerformer()           {}
func (s Substance) Is_ContractTermActionPerformer()        {}
func (o Organization) Is_ContractTermActionPerformer()     {}
func (l Location) Is_ContractTermActionPerformer()         {}

type ContractTermActionReason interface {
	gofhirclient.Resource

	Is_ContractTermActionReason()
}

func (c Condition) Is_ContractTermActionReason()             {}
func (o Observation) Is_ContractTermActionReason()           {}
func (d DiagnosticReport) Is_ContractTermActionReason()      {}
func (d DocumentReference) Is_ContractTermActionReason()     {}
func (q Questionnaire) Is_ContractTermActionReason()         {}
func (q QuestionnaireResponse) Is_ContractTermActionReason() {}

type ContractSignerParty interface {
	gofhirclient.Resource

	Is_ContractSignerParty()
}

func (o Organization) Is_ContractSignerParty()     {}
func (p Patient) Is_ContractSignerParty()          {}
func (p Practitioner) Is_ContractSignerParty()     {}
func (p PractitionerRole) Is_ContractSignerParty() {}
func (r RelatedPerson) Is_ContractSignerParty()    {}

type ContractFriendlyContentReference interface {
	gofhirclient.Resource

	Is_ContractFriendlyContentReference()
}

func (c Composition) Is_ContractFriendlyContentReference()           {}
func (d DocumentReference) Is_ContractFriendlyContentReference()     {}
func (q QuestionnaireResponse) Is_ContractFriendlyContentReference() {}

type ContractLegalContentReference interface {
	gofhirclient.Resource

	Is_ContractLegalContentReference()
}

func (c Composition) Is_ContractLegalContentReference()           {}
func (d DocumentReference) Is_ContractLegalContentReference()     {}
func (q QuestionnaireResponse) Is_ContractLegalContentReference() {}

type ContractLegallyBindingReference interface {
	gofhirclient.Resource

	Is_ContractLegallyBindingReference()
}

func (c Composition) Is_ContractLegallyBindingReference()           {}
func (d DocumentReference) Is_ContractLegallyBindingReference()     {}
func (q QuestionnaireResponse) Is_ContractLegallyBindingReference() {}
func (c Contract) Is_ContractLegallyBindingReference()              {}

func (c Contract) ResourceType() string {
	return "Contract"
}

func (c Contract) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherContract
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherContract: OtherContract(c), ResourceType: c.ResourceType()})
}

func UnmarshalContract(b []byte) (res Contract, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
