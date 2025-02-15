// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/RegulatedAuthorization
type RegulatedAuthorization struct {
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
	// Business identifier for the authorization, typically assigned by the authorizing body
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The product type, treatment, facility or activity that is being authorized
	Subject []custom.Reference[RegulatedAuthorizationSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// Overall type of this authorization, for example drug marketing approval, orphan drug designation
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// General textual supporting information
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The territory in which the authorization has been granted
	Region []CodeableConcept `bson:"region,omitempty" json:"region,omitempty"`
	// The status that is authorised e.g. approved. Intermediate states can be tracked with cases and applications
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// The date at which the current status was assigned
	StatusDate *custom.DateTime `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	// The time period in which the regulatory approval etc. is in effect, e.g. a Marketing Authorization includes the date of authorization and/or expiration date
	ValidityPeriod *Period `bson:"validityPeriod,omitempty" json:"validityPeriod,omitempty"`
	// Condition for which the use of the regulated product applies
	Indication []custom.CodeableReference[ClinicalUseDefinition] `bson:"indication,omitempty" json:"indication,omitempty"`
	// The intended use of the product, e.g. prevention, treatment
	IntendedUse *CodeableConcept `bson:"intendedUse,omitempty" json:"intendedUse,omitempty"`
	// The legal/regulatory framework or reasons under which this authorization is granted
	Basis []CodeableConcept `bson:"basis,omitempty" json:"basis,omitempty"`
	// The organization that has been granted this authorization, by the regulator
	Holder *custom.Reference[Organization] `bson:"holder,omitempty" json:"holder,omitempty"`
	// The regulatory authority or authorizing body granting the authorization
	Regulator *custom.Reference[Organization] `bson:"regulator,omitempty" json:"regulator,omitempty"`
	// Additional information or supporting documentation about the authorization
	AttachedDocument []custom.Reference[DocumentReference] `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty"`
	// The case or regulatory procedure for granting or amending a regulated authorization. Note: This area is subject to ongoing review and the workgroup is seeking implementer feedback on its use (see link at bottom of page)
	Case *RegulatedAuthorizationCase `bson:"case,omitempty" json:"case,omitempty"`
}

type RegulatedAuthorizationCase struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Identifier by which this case can be referenced
	Identifier *Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The defining type of case
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The status associated with the case
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// Relevant date for this case
	Date *Period `bson:"date,omitempty" json:"date,omitempty"`
	// Applications submitted to obtain a regulated authorization. Steps within the longer running case or procedure
	Application []interface{} `bson:"application,omitempty" json:"application,omitempty"`
}

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
	return "StructureDefinition"
}
