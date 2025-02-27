// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/RegulatedAuthorization
type RegulatedAuthorization struct {
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
	// Business identifier for the authorization, typically assigned by the authorizing body
	Identifier []Identifier `json:"identifier,omitempty"`
	// The product type, treatment, facility or activity that is being authorized
	Subject []custom.Reference[RegulatedAuthorizationSubject] `json:"subject,omitempty"`
	// Overall type of this authorization, for example drug marketing approval, orphan drug designation
	Type *CodeableConcept `json:"type,omitempty"`
	// General textual supporting information
	Description *custom.Markdown `json:"description,omitempty"`
	// The territory in which the authorization has been granted
	Region []CodeableConcept `json:"region,omitempty"`
	// The status that is authorised e.g. approved. Intermediate states can be tracked with cases and applications
	Status *CodeableConcept `json:"status,omitempty"`
	// The date at which the current status was assigned
	StatusDate *custom.DateTime `json:"statusDate,omitempty"`
	// The time period in which the regulatory approval etc. is in effect, e.g. a Marketing Authorization includes the date of authorization and/or expiration date
	ValidityPeriod *Period `json:"validityPeriod,omitempty"`
	// Condition for which the use of the regulated product applies
	Indication []custom.CodeableReference[ClinicalUseDefinition] `json:"indication,omitempty"`
	// The intended use of the product, e.g. prevention, treatment
	IntendedUse *CodeableConcept `json:"intendedUse,omitempty"`
	// The legal/regulatory framework or reasons under which this authorization is granted
	Basis []CodeableConcept `json:"basis,omitempty"`
	// The organization that has been granted this authorization, by the regulator
	Holder *custom.Reference[Organization] `json:"holder,omitempty"`
	// The regulatory authority or authorizing body granting the authorization
	Regulator *custom.Reference[Organization] `json:"regulator,omitempty"`
	// Additional information or supporting documentation about the authorization
	AttachedDocument []custom.Reference[DocumentReference] `json:"attachedDocument,omitempty"`
	// The case or regulatory procedure for granting or amending a regulated authorization. Note: This area is subject to ongoing review and the workgroup is seeking implementer feedback on its use (see link at bottom of page)
	Case *RegulatedAuthorizationCase `json:"case,omitempty"`
}

type RegulatedAuthorizationCase struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Identifier by which this case can be referenced
	Identifier *Identifier `json:"identifier,omitempty"`
	// The defining type of case
	Type *CodeableConcept `json:"type,omitempty"`
	// The status associated with the case
	Status *CodeableConcept `json:"status,omitempty"`
	// Relevant date for this case
	DatePeriod *Period `json:"datePeriod,omitempty"`
	// Relevant date for this case
	DateDateTime *custom.DateTime `json:"dateDateTime,omitempty"`
	// Applications submitted to obtain a regulated authorization. Steps within the longer running case or procedure
	Application []interface{} `json:"application,omitempty"`
}

type OtherRegulatedAuthorization RegulatedAuthorization

type RegulatedAuthorizationSubject interface {
	gofhirclient.Resource

	Is_RegulatedAuthorizationSubject()
}

func (m MedicinalProductDefinition) Is_RegulatedAuthorizationSubject() {}
func (b BiologicallyDerivedProduct) Is_RegulatedAuthorizationSubject() {}
func (n NutritionProduct) Is_RegulatedAuthorizationSubject()           {}
func (p PackagedProductDefinition) Is_RegulatedAuthorizationSubject()  {}
func (m ManufacturedItemDefinition) Is_RegulatedAuthorizationSubject() {}
func (i Ingredient) Is_RegulatedAuthorizationSubject()                 {}
func (s SubstanceDefinition) Is_RegulatedAuthorizationSubject()        {}
func (d DeviceDefinition) Is_RegulatedAuthorizationSubject()           {}
func (r ResearchStudy) Is_RegulatedAuthorizationSubject()              {}
func (a ActivityDefinition) Is_RegulatedAuthorizationSubject()         {}
func (p PlanDefinition) Is_RegulatedAuthorizationSubject()             {}
func (o ObservationDefinition) Is_RegulatedAuthorizationSubject()      {}
func (p Practitioner) Is_RegulatedAuthorizationSubject()               {}
func (o Organization) Is_RegulatedAuthorizationSubject()               {}
func (l Location) Is_RegulatedAuthorizationSubject()                   {}

func (r RegulatedAuthorization) ResourceType() string {
	return "RegulatedAuthorization"
}

func (r RegulatedAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRegulatedAuthorization
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherRegulatedAuthorization: OtherRegulatedAuthorization(r), ResourceType: r.ResourceType()})
}

func UnmarshalRegulatedAuthorization(b []byte) (res RegulatedAuthorization, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
