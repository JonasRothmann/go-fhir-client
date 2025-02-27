// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Consent
type ConsentProvisionData struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// instance | related | dependents | authoredby
	Meaning custom.Code `json:"meaning"`
	// The actual data reference
	Reference custom.Reference[Resource] `json:"reference"`
}

type Consent struct {
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
	// Identifier for this record (external references)
	Identifier []Identifier `json:"identifier,omitempty"`
	// draft | active | inactive | not-done | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Classification of the consent statement - for indexing/retrieval
	Category []CodeableConcept `json:"category,omitempty"`
	// Who the consent applies to
	Subject *custom.Reference[ConsentSubject] `json:"subject,omitempty"`
	// Fully executed date of the consent
	Date *custom.Date `json:"date,omitempty"`
	// Effective period for this Consent
	Period *Period `json:"period,omitempty"`
	// Who is granting rights according to the policy and rules
	Grantor []custom.Reference[ConsentGrantor] `json:"grantor,omitempty"`
	// Who is agreeing to the policy and rules
	Grantee []custom.Reference[ConsentGrantee] `json:"grantee,omitempty"`
	// Consent workflow management
	Manager []custom.Reference[ConsentManager] `json:"manager,omitempty"`
	// Consent Enforcer
	Controller []custom.Reference[ConsentController] `json:"controller,omitempty"`
	// Source from which this consent is taken
	SourceAttachment []Attachment `json:"sourceAttachment,omitempty"`
	// Source from which this consent is taken
	SourceReference []custom.Reference[ConsentSourceReference] `json:"sourceReference,omitempty"`
	// Regulations establishing base Consent
	RegulatoryBasis []CodeableConcept `json:"regulatoryBasis,omitempty"`
	// Computable version of the backing policy
	PolicyBasis *ConsentPolicyBasis `json:"policyBasis,omitempty"`
	// Human Readable Policy
	PolicyText []custom.Reference[DocumentReference] `json:"policyText,omitempty"`
	// Consent Verified by patient or family
	Verification []ConsentVerification `json:"verification,omitempty"`
	// deny | permit
	Decision *custom.Code `json:"decision,omitempty"`
	// Constraints to the base Consent.policyRule/Consent.policy
	Provision []ConsentProvision `json:"provision,omitempty"`
}

type ConsentPolicyBasis struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Reference backing policy resource
	Reference *custom.Reference[Resource] `json:"reference,omitempty"`
	// URL to a computable backing policy
	Url *custom.URL `json:"url,omitempty"`
}

type ConsentVerification struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Has been verified
	Verified bool `json:"verified"`
	// Business case of verification
	VerificationType *CodeableConcept `json:"verificationType,omitempty"`
	// Person conducting verification
	VerifiedBy *custom.Reference[ConsentVerificationVerifiedBy] `json:"verifiedBy,omitempty"`
	// Person who verified
	VerifiedWith *custom.Reference[ConsentVerificationVerifiedWith] `json:"verifiedWith,omitempty"`
	// When consent verified
	VerificationDate []custom.DateTime `json:"verificationDate,omitempty"`
}

type ConsentProvision struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Timeframe for this provision
	Period *Period `json:"period,omitempty"`
	// Who|what controlled by this provision (or group, by role)
	Actor []ConsentProvisionActor `json:"actor,omitempty"`
	// Actions controlled by this provision
	Action []CodeableConcept `json:"action,omitempty"`
	// Security Labels that define affected resources
	SecurityLabel []Coding `json:"securityLabel,omitempty"`
	// Context of activities covered by this provision
	Purpose []Coding `json:"purpose,omitempty"`
	// e.g. Resource Type, Profile, CDA, etc
	DocumentType []Coding `json:"documentType,omitempty"`
	// e.g. Resource Type, Profile, etc
	ResourceType []Coding `json:"resourceType,omitempty"`
	// e.g. LOINC or SNOMED CT code, etc. in the content
	Code []CodeableConcept `json:"code,omitempty"`
	// Timeframe for data controlled by this provision
	DataPeriod *Period `json:"dataPeriod,omitempty"`
	// Data controlled by this provision
	Data []ConsentProvisionData `json:"data,omitempty"`
	// A computable expression of the consent
	Expression *Expression `json:"expression,omitempty"`
	// Nested Exception Provisions
	Provision []interface{} `json:"provision,omitempty"`
}

type ConsentProvisionActor struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// How the actor is involved
	Role *CodeableConcept `json:"role,omitempty"`
	// Resource for the actor (or group, by role)
	Reference *custom.Reference[ConsentProvisionActorReference] `json:"reference,omitempty"`
}

type OtherConsent Consent

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
	return "Consent"
}

func (c Consent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConsent
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherConsent: OtherConsent(c), ResourceType: c.ResourceType()})
}

func UnmarshalConsent(b []byte) (res Consent, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
