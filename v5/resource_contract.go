// GENERATED CODE – DO NOT EDIT!

package v5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Contract
type ContractTermOffer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Offer business ID
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Offer Recipient
	Party []ContractTermOfferParty `bson:"party,omitempty" json:"party,omitempty"`
	// Negotiable offer asset
	Topic *custom.Reference[Resource] `bson:"topic,omitempty" json:"topic,omitempty"`
	// Contract Offer Type or Form
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Accepting party choice
	Decision *CodeableConcept `bson:"decision,omitempty" json:"decision,omitempty"`
	// How decision is conveyed
	DecisionMode []CodeableConcept `bson:"decisionMode,omitempty" json:"decisionMode,omitempty"`
	// Response to offer text
	Answer []ContractTermOfferAnswer `bson:"answer,omitempty" json:"answer,omitempty"`
	// Human readable offer text
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// Pointer to text
	LinkId []string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Offer restriction numbers
	SecurityLabelNumber []uint32 `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}

type ContractTermOfferParty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Referenced entity
	Reference []custom.Reference[ContractTermOfferPartyReference] `bson:"reference" json:"reference"`
	// Participant engagement type
	Role CodeableConcept `bson:"role" json:"role"`
}

type ContractTermOfferAnswer struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The actual answer response
	Value bool `bson:"value" json:"value"`
}

type ContractTermAsset struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Range of asset
	Scope *CodeableConcept `bson:"scope,omitempty" json:"scope,omitempty"`
	// Asset category
	Type []CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Associated entities
	TypeReference []custom.Reference[Resource] `bson:"typeReference,omitempty" json:"typeReference,omitempty"`
	// Asset sub-category
	Subtype []CodeableConcept `bson:"subtype,omitempty" json:"subtype,omitempty"`
	// Kinship of the asset
	Relationship *Coding `bson:"relationship,omitempty" json:"relationship,omitempty"`
	// Circumstance of the asset
	Context []ContractTermAssetContext `bson:"context,omitempty" json:"context,omitempty"`
	// Quality desctiption of asset
	Condition *string `bson:"condition,omitempty" json:"condition,omitempty"`
	// Asset availability types
	PeriodType []CodeableConcept `bson:"periodType,omitempty" json:"periodType,omitempty"`
	// Time period of the asset
	Period []Period `bson:"period,omitempty" json:"period,omitempty"`
	// Time period
	UsePeriod []Period `bson:"usePeriod,omitempty" json:"usePeriod,omitempty"`
	// Asset clause or question text
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// Pointer to asset text
	LinkId []string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Response to assets
	Answer []interface{} `bson:"answer,omitempty" json:"answer,omitempty"`
	// Asset restriction numbers
	SecurityLabelNumber []uint32 `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
	// Contract Valued Item List
	ValuedItem []ContractTermAssetValuedItem `bson:"valuedItem,omitempty" json:"valuedItem,omitempty"`
}

type ContractFriendly struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Easily comprehended representation of this Contract
	Content Attachment `bson:"content" json:"content"`
}

type ContractRule struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Computable Contract Rules
	Content Attachment `bson:"content" json:"content"`
}

type ContractContentDefinition struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Content structure and use
	Type CodeableConcept `bson:"type" json:"type"`
	// Detailed Content Type Definition
	SubType *CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty"`
	// Publisher Entity
	Publisher *custom.Reference[ContractContentDefinitionPublisher] `bson:"publisher,omitempty" json:"publisher,omitempty"`
	// When published
	PublicationDate *custom.DateTime `bson:"publicationDate,omitempty" json:"publicationDate,omitempty"`
	// amended | appended | cancelled | disputed | entered-in-error | executable +
	PublicationStatus custom.Code `bson:"publicationStatus" json:"publicationStatus"`
	// Publication Ownership
	Copyright *custom.Markdown `bson:"copyright,omitempty" json:"copyright,omitempty"`
}

type ContractTerm struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Contract Term Number
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Contract Term Issue Date Time
	Issued *custom.DateTime `bson:"issued,omitempty" json:"issued,omitempty"`
	// Contract Term Effective Time
	Applies *Period `bson:"applies,omitempty" json:"applies,omitempty"`
	// Term Concern
	Topic *CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Contract Term Type or Form
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Contract Term Type specific classification
	SubType *CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty"`
	// Term Statement
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
	// Protection for the Term
	SecurityLabel []ContractTermSecurityLabel `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	// Context of the Contract term
	Offer ContractTermOffer `bson:"offer" json:"offer"`
	// Contract Term Asset List
	Asset []ContractTermAsset `bson:"asset,omitempty" json:"asset,omitempty"`
	// Entity being ascribed responsibility
	Action []ContractTermAction `bson:"action,omitempty" json:"action,omitempty"`
	// Nested Contract Term Group
	Group []interface{} `bson:"group,omitempty" json:"group,omitempty"`
}

type ContractTermAssetContext struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Creator,custodian or owner
	Reference *custom.Reference[Resource] `bson:"reference,omitempty" json:"reference,omitempty"`
	// Codeable asset context
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Context description
	Text *string `bson:"text,omitempty" json:"text,omitempty"`
}

type ContractTermAssetValuedItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Contract Valued Item Type
	Entity *CodeableConcept `bson:"entity,omitempty" json:"entity,omitempty"`
	// Contract Valued Item Number
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Contract Valued Item Effective Tiem
	EffectiveTime *custom.DateTime `bson:"effectiveTime,omitempty" json:"effectiveTime,omitempty"`
	// Count of Contract Valued Items
	Quantity *Quantity `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Contract Valued Item fee, charge, or cost
	UnitPrice *Money `bson:"unitPrice,omitempty" json:"unitPrice,omitempty"`
	// Contract Valued Item Price Scaling Factor
	Factor *json.Number `bson:"factor,omitempty" json:"factor,omitempty"`
	// Contract Valued Item Difficulty Scaling Factor
	Points *json.Number `bson:"points,omitempty" json:"points,omitempty"`
	// Total Contract Valued Item Value
	Net *Money `bson:"net,omitempty" json:"net,omitempty"`
	// Terms of valuation
	Payment *string `bson:"payment,omitempty" json:"payment,omitempty"`
	// When payment is due
	PaymentDate *custom.DateTime `bson:"paymentDate,omitempty" json:"paymentDate,omitempty"`
	// Who will make payment
	Responsible *custom.Reference[ContractTermAssetValuedItemResponsible] `bson:"responsible,omitempty" json:"responsible,omitempty"`
	// Who will receive payment
	Recipient *custom.Reference[ContractTermAssetValuedItemRecipient] `bson:"recipient,omitempty" json:"recipient,omitempty"`
	// Pointer to specific item
	LinkId []string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// Security Labels that define affected terms
	SecurityLabelNumber []uint32 `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}

type ContractLegal struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Contract Legal Text
	Content Attachment `bson:"content" json:"content"`
}

type ContractTermActionSubject struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Entity of the action
	Reference []custom.Reference[ContractTermActionSubjectReference] `bson:"reference" json:"reference"`
	// Role type of the agent
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
}

type Contract struct {
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
	// Contract number
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// Basal definition
	Url *custom.URI `bson:"url,omitempty" json:"url,omitempty"`
	// Business edition
	Version *string `bson:"version,omitempty" json:"version,omitempty"`
	// amended | appended | cancelled | disputed | entered-in-error | executable +
	Status *custom.Code `bson:"status,omitempty" json:"status,omitempty"`
	// Negotiation status
	LegalState *CodeableConcept `bson:"legalState,omitempty" json:"legalState,omitempty"`
	// Source Contract Definition
	InstantiatesCanonical *custom.Reference[Contract] `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	// External Contract Definition
	InstantiatesUri *custom.URI `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	// Content derived from the basal information
	ContentDerivative *CodeableConcept `bson:"contentDerivative,omitempty" json:"contentDerivative,omitempty"`
	// When this Contract was issued
	Issued *custom.DateTime `bson:"issued,omitempty" json:"issued,omitempty"`
	// Effective time
	Applies *Period `bson:"applies,omitempty" json:"applies,omitempty"`
	// Contract cessation cause
	ExpirationType *CodeableConcept `bson:"expirationType,omitempty" json:"expirationType,omitempty"`
	// Contract Target Entity
	Subject []custom.Reference[Resource] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Authority under which this Contract has standing
	Authority []custom.Reference[Organization] `bson:"authority,omitempty" json:"authority,omitempty"`
	// A sphere of control governed by an authoritative jurisdiction, organization, or person
	Domain []custom.Reference[Location] `bson:"domain,omitempty" json:"domain,omitempty"`
	// Specific Location
	Site []custom.Reference[Location] `bson:"site,omitempty" json:"site,omitempty"`
	// Computer friendly designation
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// Human Friendly name
	Title *string `bson:"title,omitempty" json:"title,omitempty"`
	// Subordinate Friendly name
	Subtitle *string `bson:"subtitle,omitempty" json:"subtitle,omitempty"`
	// Acronym or short name
	Alias []string `bson:"alias,omitempty" json:"alias,omitempty"`
	// Source of Contract
	Author *custom.Reference[ContractAuthor] `bson:"author,omitempty" json:"author,omitempty"`
	// Range of Legal Concerns
	Scope *CodeableConcept `bson:"scope,omitempty" json:"scope,omitempty"`
	// Focus of contract interest
	Topic *CodeableConcept `bson:"topic,omitempty" json:"topic,omitempty"`
	// Legal instrument category
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Subtype within the context of type
	SubType []CodeableConcept `bson:"subType,omitempty" json:"subType,omitempty"`
	// Contract precursor content
	ContentDefinition *ContractContentDefinition `bson:"contentDefinition,omitempty" json:"contentDefinition,omitempty"`
	// Contract Term List
	Term []ContractTerm `bson:"term,omitempty" json:"term,omitempty"`
	// Extra Information
	SupportingInfo []custom.Reference[Resource] `bson:"supportingInfo,omitempty" json:"supportingInfo,omitempty"`
	// Key event in Contract History
	RelevantHistory []custom.Reference[Provenance] `bson:"relevantHistory,omitempty" json:"relevantHistory,omitempty"`
	// Contract Signatory
	Signer []ContractSigner `bson:"signer,omitempty" json:"signer,omitempty"`
	// Contract Friendly Language
	Friendly []ContractFriendly `bson:"friendly,omitempty" json:"friendly,omitempty"`
	// Contract Legal Language
	Legal []ContractLegal `bson:"legal,omitempty" json:"legal,omitempty"`
	// Computable Contract Language
	Rule []ContractRule `bson:"rule,omitempty" json:"rule,omitempty"`
	// Binding Contract
	LegallyBinding *Attachment `bson:"legallyBinding,omitempty" json:"legallyBinding,omitempty"`
}

type ContractTermSecurityLabel struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Link to Security Labels
	Number []uint32 `bson:"number,omitempty" json:"number,omitempty"`
	// Confidentiality Protection
	Classification Coding `bson:"classification" json:"classification"`
	// Applicable Policy
	Category []Coding `bson:"category,omitempty" json:"category,omitempty"`
	// Handling Instructions
	Control []Coding `bson:"control,omitempty" json:"control,omitempty"`
}

type ContractTermAction struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// True if the term prohibits the  action
	DoNotPerform *bool `bson:"doNotPerform,omitempty" json:"doNotPerform,omitempty"`
	// Type or form of the action
	Type CodeableConcept `bson:"type" json:"type"`
	// Entity of the action
	Subject []ContractTermActionSubject `bson:"subject,omitempty" json:"subject,omitempty"`
	// Purpose for the Contract Term Action
	Intent CodeableConcept `bson:"intent" json:"intent"`
	// Pointer to specific item
	LinkId []string `bson:"linkId,omitempty" json:"linkId,omitempty"`
	// State of the action
	Status CodeableConcept `bson:"status" json:"status"`
	// Episode associated with action
	Context *custom.Reference[ContractTermActionContext] `bson:"context,omitempty" json:"context,omitempty"`
	// Pointer to specific item
	ContextLinkId []string `bson:"contextLinkId,omitempty" json:"contextLinkId,omitempty"`
	// When action happens
	Occurrence *custom.DateTime `bson:"occurrence,omitempty" json:"occurrence,omitempty"`
	// Who asked for action
	Requester []custom.Reference[ContractTermActionRequester] `bson:"requester,omitempty" json:"requester,omitempty"`
	// Pointer to specific item
	RequesterLinkId []string `bson:"requesterLinkId,omitempty" json:"requesterLinkId,omitempty"`
	// Kind of service performer
	PerformerType []CodeableConcept `bson:"performerType,omitempty" json:"performerType,omitempty"`
	// Competency of the performer
	PerformerRole *CodeableConcept `bson:"performerRole,omitempty" json:"performerRole,omitempty"`
	// Actor that wil execute (or not) the action
	Performer *custom.Reference[ContractTermActionPerformer] `bson:"performer,omitempty" json:"performer,omitempty"`
	// Pointer to specific item
	PerformerLinkId []string `bson:"performerLinkId,omitempty" json:"performerLinkId,omitempty"`
	// Why is action (not) needed?
	Reason []custom.CodeableReference[ContractTermActionReason] `bson:"reason,omitempty" json:"reason,omitempty"`
	// Pointer to specific item
	ReasonLinkId []string `bson:"reasonLinkId,omitempty" json:"reasonLinkId,omitempty"`
	// Comments about the action
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
	// Action restriction numbers
	SecurityLabelNumber []uint32 `bson:"securityLabelNumber,omitempty" json:"securityLabelNumber,omitempty"`
}

type ContractSigner struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Contract Signatory Role
	Type Coding `bson:"type" json:"type"`
	// Contract Signatory Party
	Party custom.Reference[ContractSignerParty] `bson:"party" json:"party"`
	// Contract Documentation Signature
	Signature []Signature `bson:"signature" json:"signature"`
}

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

func (c Contract) ResourceType() string {
	return "StructureDefinition"
}
