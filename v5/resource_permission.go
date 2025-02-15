// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/Permission
type PermissionRuleActivity struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Authorized actor(s)
	Actor []custom.Reference[PermissionRuleActivityActor] `bson:"actor,omitempty" json:"actor,omitempty"`
	// Actions controlled by this rule
	Action []CodeableConcept `bson:"action,omitempty" json:"action,omitempty"`
	// The purpose for which the permission is given
	Purpose []CodeableConcept `bson:"purpose,omitempty" json:"purpose,omitempty"`
}

type Permission struct {
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
	// active | entered-in-error | draft | rejected
	Status custom.Code `bson:"status" json:"status"`
	// The person or entity that asserts the permission
	Asserter *custom.Reference[PermissionAsserter] `bson:"asserter,omitempty" json:"asserter,omitempty"`
	// The date that permission was asserted
	Date []custom.DateTime `bson:"date,omitempty" json:"date,omitempty"`
	// The period in which the permission is active
	Validity *Period `bson:"validity,omitempty" json:"validity,omitempty"`
	// The asserted justification for using the data
	Justification *PermissionJustification `bson:"justification,omitempty" json:"justification,omitempty"`
	// deny-overrides | permit-overrides | ordered-deny-overrides | ordered-permit-overrides | deny-unless-permit | permit-unless-deny
	Combining custom.Code `bson:"combining" json:"combining"`
	// Constraints to the Permission
	Rule []PermissionRule `bson:"rule,omitempty" json:"rule,omitempty"`
}

type PermissionJustification struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The regulatory grounds upon which this Permission builds
	Basis []CodeableConcept `bson:"basis,omitempty" json:"basis,omitempty"`
	// Justifing rational
	Evidence []custom.Reference[Resource] `bson:"evidence,omitempty" json:"evidence,omitempty"`
}

type PermissionRule struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// deny | permit
	Type *custom.Code `bson:"type,omitempty" json:"type,omitempty"`
	// The selection criteria to identify data that is within scope of this provision
	Data []PermissionRuleData `bson:"data,omitempty" json:"data,omitempty"`
	// A description or definition of which activities are allowed to be done on the data
	Activity []PermissionRuleActivity `bson:"activity,omitempty" json:"activity,omitempty"`
	// What limits apply to the use of the data
	Limit []CodeableConcept `bson:"limit,omitempty" json:"limit,omitempty"`
}

type PermissionRuleData struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Explicit FHIR Resource references
	Resource []PermissionRuleDataResource `bson:"resource,omitempty" json:"resource,omitempty"`
	// Security tag code on .meta.security
	Security []Coding `bson:"security,omitempty" json:"security,omitempty"`
	// Timeframe encompasing data create/update
	Period []Period `bson:"period,omitempty" json:"period,omitempty"`
	// Expression identifying the data
	Expression *Expression `bson:"expression,omitempty" json:"expression,omitempty"`
}

type PermissionRuleDataResource struct {
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
	return "StructureDefinition"
}
