// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Consent
type ConsentPolicyBasis struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Reference backing policy resource
	Reference *custom.Reference[Resource] `bson:"reference,omitempty" json:"reference,omitempty"`
	// URL to a computable backing policy
	Url *custom.URL `bson:"url,omitempty" json:"url,omitempty"`
}

type ConsentVerification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Has been verified
	Verified bool `bson:"verified" json:"verified"`
	// Business case of verification
	VerificationType *CodeableConcept `bson:"verificationType,omitempty" json:"verificationType,omitempty"`
	// Person conducting verification
	VerifiedBy *custom.Reference[ConsentVerificationVerifiedBy] `bson:"verifiedBy,omitempty" json:"verifiedBy,omitempty"`
	// Person who verified
	VerifiedWith *custom.Reference[ConsentVerificationVerifiedWith] `bson:"verifiedWith,omitempty" json:"verifiedWith,omitempty"`
	// When consent verified
	VerificationDate []custom.DateTime `bson:"verificationDate,omitempty" json:"verificationDate,omitempty"`
}

type ConsentProvision struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Timeframe for this provision
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Who|what controlled by this provision (or group, by role)
	Actor []ConsentProvisionActor `bson:"actor,omitempty" json:"actor,omitempty"`
	// Actions controlled by this provision
	Action []CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	// Security Labels that define affected resources
	SecurityLabel []Coding `bson:"securityLabel,omitempty" json:"securityLabel,omitempty"`
	// Context of activities covered by this provision
	Purpose []Coding `bson:"purpose,omitempty" json:"purpose,omitempty"`
	// e.g. Resource Type, Profile, CDA, etc
	DocumentType []Coding `bson:"documentType,omitempty" json:"documentType,omitempty"`
	// e.g. Resource Type, Profile, etc
	ResourceType []Coding `bson:"resourceType,omitempty" json:"resourceType,omitempty"`
	// e.g. LOINC or SNOMED CT code, etc. in the content
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// Timeframe for data controlled by this provision
	DataPeriod *Period `bson:"dataPeriod,omitempty" json:"dataPeriod,omitempty"`
	// Data controlled by this provision
	Data []ConsentProvisionData `bson:"data,omitempty" json:"data,omitempty"`
	// A computable expression of the consent
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty"`
	// Nested Exception Provisions
	Provision []interface{} `bson:"provision,omitempty" json:"provision,omitempty"`
}

type ConsentProvisionActor struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// How the actor is involved
	Role *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	// Resource for the actor (or group, by role)
	Reference *custom.Reference[ConsentProvisionActorReference] `bson:"reference,omitempty" json:"reference,omitempty"`
}

type ConsentProvisionData struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// instance | related | dependents | authoredby
	Meaning custom.Code `bson:"meaning" json:"meaning"`
	// The actual data reference
	Reference custom.Reference[Resource] `bson:"reference" json:"reference"`
}

type Consent struct {
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
	// Identifier for this record (external references)
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// draft | active | inactive | not-done | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Classification of the consent statement - for indexing/retrieval
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Who the consent applies to
	Subject *custom.Reference[ConsentSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Fully executed date of the consent
	Date *custom.Date `bson:"date,omitempty" json:"date,omitempty"`
	// Effective period for this Consent
	Period *Period `bson:"period,omitempty" json:"period,omitempty"`
	// Who is granting rights according to the policy and rules
	Grantor []custom.Reference[ConsentGrantor] `bson:"grantor,omitempty" json:"grantor,omitempty"`
	// Who is agreeing to the policy and rules
	Grantee []custom.Reference[ConsentGrantee] `bson:"grantee,omitempty" json:"grantee,omitempty"`
	// Consent workflow management
	Manager []custom.Reference[ConsentManager] `bson:"manager,omitempty" json:"manager,omitempty"`
	// Consent Enforcer
	Controller []custom.Reference[ConsentController] `bson:"controller,omitempty" json:"controller,omitempty"`
	// Source from which this consent is taken
	SourceAttachment []Attachment `bson:"sourceAttachment,omitempty" json:"sourceAttachment,omitempty"`
	// Source from which this consent is taken
	SourceReference []custom.Reference[ConsentSourceReference] `bson:"sourceReference,omitempty" json:"sourceReference,omitempty"`
	// Regulations establishing base Consent
	RegulatoryBasis []CodeableConcept `bson:"regulatoryBasis,omitempty" json:"regulatoryBasis,omitempty"`
	// Computable version of the backing policy
	PolicyBasis *ConsentPolicyBasis `bson:"policyBasis,omitempty" json:"policyBasis,omitempty"`
	// Human Readable Policy
	PolicyText []custom.Reference[DocumentReference] `bson:"policyText,omitempty" json:"policyText,omitempty"`
	// Consent Verified by patient or family
	Verification []ConsentVerification `bson:"verification,omitempty" json:"verification,omitempty"`
	// deny | permit
	Decision *custom.Code `bson:"decision,omitempty" json:"decision,omitempty"`
	// Constraints to the base Consent.policyRule/Consent.policy
	Provision []ConsentProvision `bson:"provision,omitempty" json:"provision,omitempty"`
}

type ConsentSubject interface {
	gofhirclient.Resource

	Is_ConsentSubject()
}

func (p Patient) Is_ConsentSubject()      {}
func (p Practitioner) Is_ConsentSubject() {}
func (g Group) Is_ConsentSubject()        {}

type ConsentGrantor interface {
	gofhirclient.Resource

	Is_ConsentGrantor()
}

func (c CareTeam) Is_ConsentGrantor()          {}
func (h HealthcareService) Is_ConsentGrantor() {}
func (o Organization) Is_ConsentGrantor()      {}
func (p Patient) Is_ConsentGrantor()           {}
func (p Practitioner) Is_ConsentGrantor()      {}
func (r RelatedPerson) Is_ConsentGrantor()     {}
func (p PractitionerRole) Is_ConsentGrantor()  {}

type ConsentGrantee interface {
	gofhirclient.Resource

	Is_ConsentGrantee()
}

func (c CareTeam) Is_ConsentGrantee()          {}
func (h HealthcareService) Is_ConsentGrantee() {}
func (o Organization) Is_ConsentGrantee()      {}
func (p Patient) Is_ConsentGrantee()           {}
func (p Practitioner) Is_ConsentGrantee()      {}
func (r RelatedPerson) Is_ConsentGrantee()     {}
func (p PractitionerRole) Is_ConsentGrantee()  {}

type ConsentManager interface {
	gofhirclient.Resource

	Is_ConsentManager()
}

func (h HealthcareService) Is_ConsentManager() {}
func (o Organization) Is_ConsentManager()      {}
func (p Patient) Is_ConsentManager()           {}
func (p Practitioner) Is_ConsentManager()      {}

type ConsentController interface {
	gofhirclient.Resource

	Is_ConsentController()
}

func (h HealthcareService) Is_ConsentController() {}
func (o Organization) Is_ConsentController()      {}
func (p Patient) Is_ConsentController()           {}
func (p Practitioner) Is_ConsentController()      {}

type ConsentSourceReference interface {
	gofhirclient.Resource

	Is_ConsentSourceReference()
}

func (c Consent) Is_ConsentSourceReference()               {}
func (d DocumentReference) Is_ConsentSourceReference()     {}
func (c Contract) Is_ConsentSourceReference()              {}
func (q QuestionnaireResponse) Is_ConsentSourceReference() {}

type ConsentVerificationVerifiedBy interface {
	gofhirclient.Resource

	Is_ConsentVerificationVerifiedBy()
}

func (o Organization) Is_ConsentVerificationVerifiedBy()     {}
func (p Practitioner) Is_ConsentVerificationVerifiedBy()     {}
func (p PractitionerRole) Is_ConsentVerificationVerifiedBy() {}

type ConsentVerificationVerifiedWith interface {
	gofhirclient.Resource

	Is_ConsentVerificationVerifiedWith()
}

func (p Patient) Is_ConsentVerificationVerifiedWith()       {}
func (r RelatedPerson) Is_ConsentVerificationVerifiedWith() {}

type ConsentProvisionActorReference interface {
	gofhirclient.Resource

	Is_ConsentProvisionActorReference()
}

func (d Device) Is_ConsentProvisionActorReference()           {}
func (g Group) Is_ConsentProvisionActorReference()            {}
func (c CareTeam) Is_ConsentProvisionActorReference()         {}
func (o Organization) Is_ConsentProvisionActorReference()     {}
func (p Patient) Is_ConsentProvisionActorReference()          {}
func (p Practitioner) Is_ConsentProvisionActorReference()     {}
func (r RelatedPerson) Is_ConsentProvisionActorReference()    {}
func (p PractitionerRole) Is_ConsentProvisionActorReference() {}

func (c Consent) ResourceType() string {
	return "StructureDefinition"
}
