// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinitionPackaging struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// An identifier that is specific to this particular part of the packaging. Including possibly a Data Carrier Identifier
	Identifier []Identifier `json:"identifier,omitempty"`
	// The physical type of the container of the items
	Type *CodeableConcept `json:"type,omitempty"`
	// Is this a part of the packaging (e.g. a cap or bottle stopper), rather than the packaging itself (e.g. a bottle or vial)
	ComponentPart *bool `json:"componentPart,omitempty"`
	// The quantity of this level of packaging in the package that contains it (with the outermost level being 1)
	Quantity *int32 `json:"quantity,omitempty"`
	// Material type of the package item
	Material []CodeableConcept `json:"material,omitempty"`
	// A possible alternate material for this part of the packaging, that is allowed to be used instead of the usual material
	AlternateMaterial []CodeableConcept `json:"alternateMaterial,omitempty"`
	// Shelf Life and storage information
	ShelfLifeStorage []ProductShelfLife `json:"shelfLifeStorage,omitempty"`
	// Manufacturer of this packaging item (multiple means these are all potential manufacturers)
	Manufacturer []custom.Reference[Organization] `json:"manufacturer,omitempty"`
	// General characteristics of this item
	Property []PackagedProductDefinitionPackagingProperty `json:"property,omitempty"`
	// The item(s) within the packaging
	ContainedItem []PackagedProductDefinitionPackagingContainedItem `json:"containedItem,omitempty"`
	// Allows containers (and parts of containers) within containers, still as a part of single packaged product
	Packaging []interface{} `json:"packaging,omitempty"`
}

type PackagedProductDefinitionPackagingProperty struct {
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
	ValueAttachment *Attachment `json:"valueAttachment,omitempty"`
}

type PackagedProductDefinitionPackagingContainedItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The actual item(s) of medication, as manufactured, or a device, or other medically related item (food, biologicals, raw materials, medical fluids, gases etc.), as contained in the package
	Item custom.CodeableReference[PackagedProductDefinitionPackagingContainedItemItem] `json:"item"`
	// The number of this type of item within this packaging or for continuous items such as liquids it is the quantity (for example 25ml). See also PackagedProductDefinition.containedItemQuantity (especially the long definition)
	Amount *Quantity `json:"amount,omitempty"`
}

type PackagedProductDefinition struct {
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
	// A unique identifier for this package as whole - not for the content of the package
	Identifier []Identifier `json:"identifier,omitempty"`
	// A name for this package. Typically as listed in a drug formulary, catalogue, inventory etc
	Name *string `json:"name,omitempty"`
	// A high level category e.g. medicinal product, raw material, shipping container etc
	Type *CodeableConcept `json:"type,omitempty"`
	// The product that this is a pack for
	PackageFor []custom.Reference[MedicinalProductDefinition] `json:"packageFor,omitempty"`
	// The status within the lifecycle of this item. High level - not intended to duplicate details elsewhere e.g. legal status, or authorization/marketing status
	Status *CodeableConcept `json:"status,omitempty"`
	// The date at which the given status became applicable
	StatusDate *custom.DateTime `json:"statusDate,omitempty"`
	// A total of the complete count of contained items of a particular type/form, independent of sub-packaging or organization. This can be considered as the pack size. See also packaging.containedItem.amount (especially the long definition)
	ContainedItemQuantity []Quantity `json:"containedItemQuantity,omitempty"`
	// Textual description. Note that this is not the name of the package or product
	Description *custom.Markdown `json:"description,omitempty"`
	// The legal status of supply of the packaged item as classified by the regulator
	LegalStatusOfSupply []PackagedProductDefinitionLegalStatusOfSupply `json:"legalStatusOfSupply,omitempty"`
	// Allows specifying that an item is on the market for sale, or that it is not available, and the dates and locations associated
	MarketingStatus []MarketingStatus `json:"marketingStatus,omitempty"`
	// Identifies if the drug product is supplied with another item such as a diluent or adjuvant
	CopackagedIndicator *bool `json:"copackagedIndicator,omitempty"`
	// Manufacturer of this package type (multiple means these are all possible manufacturers)
	Manufacturer []custom.Reference[Organization] `json:"manufacturer,omitempty"`
	// Additional information or supporting documentation about the packaged product
	AttachedDocument []custom.Reference[DocumentReference] `json:"attachedDocument,omitempty"`
	// A packaging item, as a container for medically related items, possibly with other packaging items within, or a packaging component, such as bottle cap
	Packaging *PackagedProductDefinitionPackaging `json:"packaging,omitempty"`
	// Allows the key features to be recorded, such as "hospital pack", "nurse prescribable"
	Characteristic []interface{} `json:"characteristic,omitempty"`
}

type PackagedProductDefinitionLegalStatusOfSupply struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The actual status of supply. In what situation this package type may be supplied for use
	Code *CodeableConcept `json:"code,omitempty"`
	// The place where the legal status of supply applies
	Jurisdiction *CodeableConcept `json:"jurisdiction,omitempty"`
}

type OtherPackagedProductDefinition PackagedProductDefinition

type PackagedProductDefinitionPackagingContainedItemItem interface {
	gofhirclient.Resource

	Is_PackagedProductDefinitionPackagingContainedItemItem()
}

func (m ManufacturedItemDefinition) Is_PackagedProductDefinitionPackagingContainedItemItem() {}
func (d DeviceDefinition) Is_PackagedProductDefinitionPackagingContainedItemItem()           {}
func (p PackagedProductDefinition) Is_PackagedProductDefinitionPackagingContainedItemItem()  {}
func (b BiologicallyDerivedProduct) Is_PackagedProductDefinitionPackagingContainedItemItem() {}
func (n NutritionProduct) Is_PackagedProductDefinitionPackagingContainedItemItem()           {}

func (p PackagedProductDefinition) ResourceType() string {
	return "PackagedProductDefinition"
}

func (p PackagedProductDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPackagedProductDefinition
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherPackagedProductDefinition: OtherPackagedProductDefinition(p), ResourceType: p.ResourceType()})
}

func UnmarshalPackagedProductDefinition(b []byte) (res PackagedProductDefinition, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
