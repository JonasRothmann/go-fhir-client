// GENERATED CODE – DO NOT EDIT!

package v5

import custom "github.com/jonasrothmann/go-fhir-client/custom"

// http://hl7.org/fhir/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinition struct {
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
	// An identifier for the administrable product
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `bson:"status" json:"status"`
	// References a product from which one or more of the constituent parts of that product can be prepared and used as described by this administrable product
	FormOf []custom.Reference[MedicinalProductDefinition] `bson:"formOf,omitempty" json:"formOf,omitempty"`
	// The dose form of the final product after necessary reconstitution or processing
	AdministrableDoseForm *CodeableConcept `bson:"administrableDoseForm,omitempty" json:"administrableDoseForm,omitempty"`
	// The presentation type in which this item is given to a patient. e.g. for a spray - 'puff'
	UnitOfPresentation *CodeableConcept `bson:"unitOfPresentation,omitempty" json:"unitOfPresentation,omitempty"`
	// Indicates the specific manufactured items that are part of the 'formOf' product that are used in the preparation of this specific administrable form
	ProducedFrom []custom.Reference[ManufacturedItemDefinition] `bson:"producedFrom,omitempty" json:"producedFrom,omitempty"`
	// The ingredients of this administrable medicinal product. This is only needed if the ingredients are not specified either using ManufacturedItemDefiniton, or using by incoming references from the Ingredient resource
	Ingredient []CodeableConcept `bson:"ingredient,omitempty" json:"ingredient,omitempty"`
	// A device that is integral to the medicinal product, in effect being considered as an "ingredient" of the medicinal product
	Device *custom.Reference[DeviceDefinition] `bson:"device,omitempty" json:"device,omitempty"`
	// A general description of the product, when in its final form, suitable for administration e.g. effervescent blue liquid, to be swallowed
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// Characteristics e.g. a product's onset of action
	Property []AdministrableProductDefinitionProperty `bson:"property,omitempty" json:"property,omitempty"`
	// The path by which the product is taken into or makes contact with the body
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `bson:"routeOfAdministration" json:"routeOfAdministration"`
}

type AdministrableProductDefinitionProperty struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// A code expressing the type of characteristic
	Type CodeableConcept `bson:"type" json:"type"`
	// A value for the characteristic
	Value *CodeableConcept `bson:"value,omitempty" json:"value,omitempty"`
	// The status of characteristic e.g. assigned or pending
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministration struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Coded expression for the route
	Code CodeableConcept `bson:"code" json:"code"`
	// The first dose (dose quantity) administered can be specified for the product
	FirstDose *Quantity `bson:"firstDose,omitempty" json:"firstDose,omitempty"`
	// The maximum single dose that can be administered
	MaxSingleDose *Quantity `bson:"maxSingleDose,omitempty" json:"maxSingleDose,omitempty"`
	// The maximum dose quantity to be administered in any one 24-h period
	MaxDosePerDay *Quantity `bson:"maxDosePerDay,omitempty" json:"maxDosePerDay,omitempty"`
	// The maximum dose per treatment period that can be administered
	MaxDosePerTreatmentPeriod *Ratio `bson:"maxDosePerTreatmentPeriod,omitempty" json:"maxDosePerTreatmentPeriod,omitempty"`
	// The maximum treatment period during which the product can be administered
	MaxTreatmentPeriod *Duration `bson:"maxTreatmentPeriod,omitempty" json:"maxTreatmentPeriod,omitempty"`
	// A species for which this route applies
	TargetSpecies []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `bson:"targetSpecies,omitempty" json:"targetSpecies,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Coded expression for the species
	Code CodeableConcept `bson:"code" json:"code"`
	// A species specific time during which consumption of animal product is not appropriate
	WithdrawalPeriod []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `bson:"withdrawalPeriod,omitempty" json:"withdrawalPeriod,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of tissue for which the withdrawal period applies, e.g. meat, milk
	Tissue CodeableConcept `bson:"tissue" json:"tissue"`
	// A value for the time
	Value Quantity `bson:"value" json:"value"`
	// Extra information about the withdrawal period
	SupportingInformation *string `bson:"supportingInformation,omitempty" json:"supportingInformation,omitempty"`
}

func (a AdministrableProductDefinition) ResourceType() string {
	return "StructureDefinition"
}
