// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/InventoryItem
type InventoryItem struct {
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
	// Business identifier for the inventory item
	Identifier []Identifier `json:"identifier,omitempty"`
	// active | inactive | entered-in-error | unknown
	Status custom.Code `json:"status"`
	// Category or class of the item
	Category []CodeableConcept `json:"category,omitempty"`
	// Code designating the specific type of item
	Code []CodeableConcept `json:"code,omitempty"`
	// The item name(s) - the brand name, or common name, functional name, generic name or others
	Name []InventoryItemName `json:"name,omitempty"`
	// Organization(s) responsible for the product
	ResponsibleOrganization []InventoryItemResponsibleOrganization `json:"responsibleOrganization,omitempty"`
	// Descriptive characteristics of the item
	Description *InventoryItemDescription `json:"description,omitempty"`
	// The usage status like recalled, in use, discarded
	InventoryStatus []CodeableConcept `json:"inventoryStatus,omitempty"`
	// The base unit of measure - the unit in which the product is used or counted
	BaseUnit *CodeableConcept `json:"baseUnit,omitempty"`
	// Net content or amount present in the item
	NetContent *Quantity `json:"netContent,omitempty"`
	// Association with other items or products
	Association []InventoryItemAssociation `json:"association,omitempty"`
	// Characteristic of the item
	Characteristic []InventoryItemCharacteristic `json:"characteristic,omitempty"`
	// Instances or occurrences of the product
	Instance *InventoryItemInstance `json:"instance,omitempty"`
	// Link to a product resource used in clinical workflows
	ProductReference *custom.Reference[InventoryItemProductReference] `json:"productReference,omitempty"`
}

type InventoryItemName struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of name e.g. 'brand-name', 'functional-name', 'common-name'
	NameType Coding `json:"nameType"`
	// The language used to express the item name
	Language custom.Code `json:"language"`
	// The name or designation of the item
	Name string `json:"name"`
}

type InventoryItemResponsibleOrganization struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The role of the organization e.g. manufacturer, distributor, or other
	Role CodeableConcept `json:"role"`
	// An organization that is associated with the item
	Organization custom.Reference[Organization] `json:"organization"`
}

type InventoryItemDescription struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The language that is used in the item description
	Language *custom.Code `json:"language,omitempty"`
	// Textual description of the item
	Description *string `json:"description,omitempty"`
}

type InventoryItemAssociation struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The type of association between the device and the other item
	AssociationType CodeableConcept `json:"associationType"`
	// The related item or product
	RelatedItem custom.Reference[InventoryItemAssociationRelatedItem] `json:"relatedItem"`
	// The quantity of the product in this product
	Quantity Ratio `json:"quantity"`
}

type InventoryItemCharacteristic struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The characteristic that is being defined
	CharacteristicType CodeableConcept `json:"characteristicType"`
	// The value of the attribute
	ValueString *string `json:"valueString,omitempty"`
	// The value of the attribute
	ValueInteger *int32 `json:"valueInteger,omitempty"`
	// The value of the attribute
	ValueDecimal *json.Number `json:"valueDecimal,omitempty"`
	// The value of the attribute
	ValueBoolean *bool `json:"valueBoolean,omitempty"`
	// The value of the attribute
	ValueUrl *custom.URL `json:"valueUrl,omitempty"`
	// The value of the attribute
	ValueDateTime *custom.DateTime `json:"valueDateTime,omitempty"`
	// The value of the attribute
	ValueQuantity *Quantity `json:"valueQuantity,omitempty"`
	// The value of the attribute
	ValueRange *Range `json:"valueRange,omitempty"`
	// The value of the attribute
	ValueRatio *Ratio `json:"valueRatio,omitempty"`
	// The value of the attribute
	ValueAnnotation *Annotation `json:"valueAnnotation,omitempty"`
	// The value of the attribute
	ValueAddress *Address `json:"valueAddress,omitempty"`
	// The value of the attribute
	ValueDuration *Duration `json:"valueDuration,omitempty"`
	// The value of the attribute
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
}

type InventoryItemInstance struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The identifier for the physical instance, typically a serial number
	Identifier []Identifier `json:"identifier,omitempty"`
	// The lot or batch number of the item
	LotNumber *string `json:"lotNumber,omitempty"`
	// The expiry date or date and time for the product
	Expiry *custom.DateTime `json:"expiry,omitempty"`
	// The subject that the item is associated with
	Subject *custom.Reference[InventoryItemInstanceSubject] `json:"subject,omitempty"`
	// The location that the item is associated with
	Location *custom.Reference[Location] `json:"location,omitempty"`
}

type OtherInventoryItem InventoryItem

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
	return "InventoryItem"
}

func (i InventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInventoryItem
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherInventoryItem: OtherInventoryItem(i), ResourceType: i.ResourceType()})
}

func UnmarshalInventoryItem(b []byte) (res InventoryItem, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
