// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/AdministrableProductDefinition
type AdministrableProductDefinition struct {
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
	// An identifier for the administrable product
	Identifier []Identifier `json:"identifier,omitempty"`
	// draft | active | retired | unknown
	Status custom.Code `json:"status"`
	// References a product from which one or more of the constituent parts of that product can be prepared and used as described by this administrable product
	FormOf []custom.Reference[MedicinalProductDefinition] `json:"formOf,omitempty"`
	// The dose form of the final product after necessary reconstitution or processing
	AdministrableDoseForm *CodeableConcept `json:"administrableDoseForm,omitempty"`
	// The presentation type in which this item is given to a patient. e.g. for a spray - 'puff'
	UnitOfPresentation *CodeableConcept `json:"unitOfPresentation,omitempty"`
	// Indicates the specific manufactured items that are part of the 'formOf' product that are used in the preparation of this specific administrable form
	ProducedFrom []custom.Reference[ManufacturedItemDefinition] `json:"producedFrom,omitempty"`
	// The ingredients of this administrable medicinal product. This is only needed if the ingredients are not specified either using ManufacturedItemDefiniton, or using by incoming references from the Ingredient resource
	Ingredient []CodeableConcept `json:"ingredient,omitempty"`
	// A device that is integral to the medicinal product, in effect being considered as an "ingredient" of the medicinal product
	Device *custom.Reference[DeviceDefinition] `json:"device,omitempty"`
	// A general description of the product, when in its final form, suitable for administration e.g. effervescent blue liquid, to be swallowed
	Description *custom.Markdown `json:"description,omitempty"`
	// Characteristics e.g. a product's onset of action
	Property []AdministrableProductDefinitionProperty `json:"property,omitempty"`
	// The path by which the product is taken into or makes contact with the body
	RouteOfAdministration []AdministrableProductDefinitionRouteOfAdministration `json:"routeOfAdministration"`
}

type AdministrableProductDefinitionProperty struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// A code expressing the type of characteristic
	Type CodeableConcept `json:"type"`
	// A value for the characteristic
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	// A value for the characteristic
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// A value for the characteristic
	ValueDate *custom.Date `json:"valueDate,omitempty"`
	// A value for the characteristic
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// A value for the characteristic
	ValueMarkdown *custom.Markdown `json:"valueMarkdown,omitempty"`
	// A value for the characteristic
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
	// A value for the characteristic
	ValueReference *custom.Reference[Binary] `json:"valueReference,omitempty"`
	// The status of characteristic e.g. assigned or pending
	Status *CodeableConcept `json:"status,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministration struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Coded expression for the route
	Code CodeableConcept `json:"code"`
	// The first dose (dose quantity) administered can be specified for the product
	FirstDose *Quantity `json:"firstDose,omitempty"`
	// The maximum single dose that can be administered
	MaxSingleDose *Quantity `json:"maxSingleDose,omitempty"`
	// The maximum dose quantity to be administered in any one 24-h period
	MaxDosePerDay *Quantity `json:"maxDosePerDay,omitempty"`
	// The maximum dose per treatment period that can be administered
	MaxDosePerTreatmentPeriod *Ratio `json:"maxDosePerTreatmentPeriod,omitempty"`
	// The maximum treatment period during which the product can be administered
	MaxTreatmentPeriod *Duration `json:"maxTreatmentPeriod,omitempty"`
	// A species for which this route applies
	TargetSpecies []AdministrableProductDefinitionRouteOfAdministrationTargetSpecies `json:"targetSpecies,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpecies struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Coded expression for the species
	Code CodeableConcept `json:"code"`
	// A species specific time during which consumption of animal product is not appropriate
	WithdrawalPeriod []AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod `json:"withdrawalPeriod,omitempty"`
}

type AdministrableProductDefinitionRouteOfAdministrationTargetSpeciesWithdrawalPeriod struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of tissue for which the withdrawal period applies, e.g. meat, milk
	Tissue CodeableConcept `json:"tissue"`
	// A value for the time
	Value Quantity `json:"value"`
	// Extra information about the withdrawal period
	SupportingInformation *string `json:"supportingInformation,omitempty"`
}

type OtherAdministrableProductDefinition AdministrableProductDefinition

func (a AdministrableProductDefinition) ResourceType() string {
	return "AdministrableProductDefinition"
}

func (a AdministrableProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAdministrableProductDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherAdministrableProductDefinition: OtherAdministrableProductDefinition(a), ResourceType: a.ResourceType()})
}

func UnmarshalAdministrableProductDefinition(b []byte) (res AdministrableProductDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
