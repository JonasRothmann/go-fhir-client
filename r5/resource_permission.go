// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Permission
type PermissionJustification struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The regulatory grounds upon which this Permission builds
	Basis []CodeableConcept `json:"basis,omitempty"`
	// Justifing rational
	Evidence []custom.Reference[Resource] `json:"evidence,omitempty"`
}

type PermissionRule struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// deny | permit
	Type *custom.Code `json:"type,omitempty"`
	// The selection criteria to identify data that is within scope of this provision
	Data []PermissionRuleData `json:"data,omitempty"`
	// A description or definition of which activities are allowed to be done on the data
	Activity []PermissionRuleActivity `json:"activity,omitempty"`
	// What limits apply to the use of the data
	Limit []CodeableConcept `json:"limit,omitempty"`
}

type PermissionRuleData struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Explicit FHIR Resource references
	Resource []PermissionRuleDataResource `json:"resource,omitempty"`
	// Security tag code on .meta.security
	Security []Coding `json:"security,omitempty"`
	// Timeframe encompasing data create/update
	Period []Period `json:"period,omitempty"`
	// Expression identifying the data
	Expression *Expression `json:"expression,omitempty"`
}

type PermissionRuleDataResource struct {
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

type PermissionRuleActivity struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Authorized actor(s)
	Actor []custom.Reference[PermissionRuleActivityActor] `json:"actor,omitempty"`
	// Actions controlled by this rule
	Action []CodeableConcept `json:"action,omitempty"`
	// The purpose for which the permission is given
	Purpose []CodeableConcept `json:"purpose,omitempty"`
}

type Permission struct {
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
	// active | entered-in-error | draft | rejected
	Status custom.Code `json:"status"`
	// The person or entity that asserts the permission
	Asserter *custom.Reference[PermissionAsserter] `json:"asserter,omitempty"`
	// The date that permission was asserted
	Date []custom.DateTime `json:"date,omitempty"`
	// The period in which the permission is active
	Validity *Period `json:"validity,omitempty"`
	// The asserted justification for using the data
	Justification *PermissionJustification `json:"justification,omitempty"`
	// deny-overrides | permit-overrides | ordered-deny-overrides | ordered-permit-overrides | deny-unless-permit | permit-unless-deny
	Combining custom.Code `json:"combining"`
	// Constraints to the Permission
	Rule []PermissionRule `json:"rule,omitempty"`
}

type OtherPermission Permission

type PermissionAsserter interface {
	gofhirclient.Resource

	Is_PermissionAsserter()
}

func (p Practitioner) Is_PermissionAsserter()      {}
func (p PractitionerRole) Is_PermissionAsserter()  {}
func (o Organization) Is_PermissionAsserter()      {}
func (c CareTeam) Is_PermissionAsserter()          {}
func (p Patient) Is_PermissionAsserter()           {}
func (r RelatedPerson) Is_PermissionAsserter()     {}
func (h HealthcareService) Is_PermissionAsserter() {}

type PermissionRuleActivityActor interface {
	gofhirclient.Resource

	Is_PermissionRuleActivityActor()
}

func (d Device) Is_PermissionRuleActivityActor()           {}
func (g Group) Is_PermissionRuleActivityActor()            {}
func (c CareTeam) Is_PermissionRuleActivityActor()         {}
func (o Organization) Is_PermissionRuleActivityActor()     {}
func (p Patient) Is_PermissionRuleActivityActor()          {}
func (p Practitioner) Is_PermissionRuleActivityActor()     {}
func (r RelatedPerson) Is_PermissionRuleActivityActor()    {}
func (p PractitionerRole) Is_PermissionRuleActivityActor() {}

func (p Permission) ResourceType() string {
	return "Permission"
}

func (p Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPermission
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPermission: OtherPermission(p), ResourceType: p.ResourceType()})
}

func UnmarshalPermission(b []byte) (res Permission, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
