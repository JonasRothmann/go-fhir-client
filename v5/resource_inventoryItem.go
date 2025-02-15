// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/InventoryItem
type InventoryItem struct {
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
	// Business identifier for the inventory item
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// active | inactive | entered-in-error | unknown
	Status custom.Code `bson:"status" json:"status"`
	// Category or class of the item
	Category []CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// Code designating the specific type of item
	Code []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// The item name(s) - the brand name, or common name, functional name, generic name or others
	Name []InventoryItemName `bson:"name,omitempty" json:"name,omitempty"`
	// Organization(s) responsible for the product
	ResponsibleOrganization []InventoryItemResponsibleOrganization `bson:"responsibleOrganization,omitempty" json:"responsibleOrganization,omitempty"`
	// Descriptive characteristics of the item
	Description *InventoryItemDescription `bson:"description,omitempty" json:"description,omitempty"`
	// The usage status like recalled, in use, discarded
	InventoryStatus []CodeableConcept `bson:"inventoryStatus,omitempty" json:"inventoryStatus,omitempty"`
	// The base unit of measure - the unit in which the product is used or counted
	BaseUnit *CodeableConcept `bson:"baseUnit,omitempty" json:"baseUnit,omitempty"`
	// Net content or amount present in the item
	NetContent *Quantity `bson:"netContent,omitempty" json:"netContent,omitempty"`
	// Association with other items or products
	Association []InventoryItemAssociation `bson:"association,omitempty" json:"association,omitempty"`
	// Characteristic of the item
	Characteristic []InventoryItemCharacteristic `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
	// Instances or occurrences of the product
	Instance *InventoryItemInstance `bson:"instance,omitempty" json:"instance,omitempty"`
	// Link to a product resource used in clinical workflows
	ProductReference *custom.Reference[InventoryItemProductReference] `bson:"productReference,omitempty" json:"productReference,omitempty"`
}

type InventoryItemName struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of name e.g. 'brand-name', 'functional-name', 'common-name'
	NameType Coding `bson:"nameType" json:"nameType"`
	// The language used to express the item name
	Language custom.Code `bson:"language" json:"language"`
	// The name or designation of the item
	Name string `bson:"name" json:"name"`
}

type InventoryItemResponsibleOrganization struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The role of the organization e.g. manufacturer, distributor, or other
	Role CodeableConcept `bson:"role" json:"role"`
	// An organization that is associated with the item
	Organization custom.Reference[Organization] `bson:"organization" json:"organization"`
}

type InventoryItemDescription struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The language that is used in the item description
	Language *custom.Code `bson:"language,omitempty" json:"language,omitempty"`
	// Textual description of the item
	Description *string `bson:"description,omitempty" json:"description,omitempty"`
}

type InventoryItemAssociation struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The type of association between the device and the other item
	AssociationType CodeableConcept `bson:"associationType" json:"associationType"`
	// The related item or product
	RelatedItem custom.Reference[InventoryItemAssociationRelatedItem] `bson:"relatedItem" json:"relatedItem"`
	// The quantity of the product in this product
	Quantity Ratio `bson:"quantity" json:"quantity"`
}

type InventoryItemCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The characteristic that is being defined
	CharacteristicType CodeableConcept `bson:"characteristicType" json:"characteristicType"`
	// The value of the attribute
	Value string `bson:"value" json:"value"`
}

type InventoryItemInstance struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The identifier for the physical instance, typically a serial number
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The lot or batch number of the item
	LotNumber *string `bson:"lotNumber,omitempty" json:"lotNumber,omitempty"`
	// The expiry date or date and time for the product
	Expiry *custom.DateTime `bson:"expiry,omitempty" json:"expiry,omitempty"`
	// The subject that the item is associated with
	Subject *custom.Reference[InventoryItemInstanceSubject] `bson:"subject,omitempty" json:"subject,omitempty"`
	// The location that the item is associated with
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
}

type InventoryItemAssociationRelatedItem interface {
	gofhirclient.Resource

	Is_InventoryItemAssociationRelatedItem()
}

func (i InventoryItem) Is_InventoryItemAssociationRelatedItem()              {}
func (m Medication) Is_InventoryItemAssociationRelatedItem()                 {}
func (m MedicationKnowledge) Is_InventoryItemAssociationRelatedItem()        {}
func (d Device) Is_InventoryItemAssociationRelatedItem()                     {}
func (d DeviceDefinition) Is_InventoryItemAssociationRelatedItem()           {}
func (n NutritionProduct) Is_InventoryItemAssociationRelatedItem()           {}
func (b BiologicallyDerivedProduct) Is_InventoryItemAssociationRelatedItem() {}

type InventoryItemInstanceSubject interface {
	gofhirclient.Resource

	Is_InventoryItemInstanceSubject()
}

func (p Patient) Is_InventoryItemInstanceSubject()      {}
func (o Organization) Is_InventoryItemInstanceSubject() {}

type InventoryItemProductReference interface {
	gofhirclient.Resource

	Is_InventoryItemProductReference()
}

func (m Medication) Is_InventoryItemProductReference()                 {}
func (d Device) Is_InventoryItemProductReference()                     {}
func (n NutritionProduct) Is_InventoryItemProductReference()           {}
func (b BiologicallyDerivedProduct) Is_InventoryItemProductReference() {}

func (i InventoryItem) ResourceType() string {
	return "StructureDefinition"
}
