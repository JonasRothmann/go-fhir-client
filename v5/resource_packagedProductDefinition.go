// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/PackagedProductDefinition
type PackagedProductDefinition struct {
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
	// A unique identifier for this package as whole - not for the content of the package
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// A name for this package. Typically as listed in a drug formulary, catalogue, inventory etc
	Name *string `bson:"name,omitempty" json:"name,omitempty"`
	// A high level category e.g. medicinal product, raw material, shipping container etc
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// The product that this is a pack for
	PackageFor []custom.Reference[MedicinalProductDefinition] `bson:"packageFor,omitempty" json:"packageFor,omitempty"`
	// The status within the lifecycle of this item. High level - not intended to duplicate details elsewhere e.g. legal status, or authorization/marketing status
	Status *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	// The date at which the given status became applicable
	StatusDate *custom.DateTime `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	// A total of the complete count of contained items of a particular type/form, independent of sub-packaging or organization. This can be considered as the pack size. See also packaging.containedItem.amount (especially the long definition)
	ContainedItemQuantity []Quantity `bson:"containedItemQuantity,omitempty" json:"containedItemQuantity,omitempty"`
	// Textual description. Note that this is not the name of the package or product
	Description *custom.Markdown `bson:"description,omitempty" json:"description,omitempty"`
	// The legal status of supply of the packaged item as classified by the regulator
	LegalStatusOfSupply []PackagedProductDefinitionLegalStatusOfSupply `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty"`
	// Allows specifying that an item is on the market for sale, or that it is not available, and the dates and locations associated
	MarketingStatus []MarketingStatus `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty"`
	// Identifies if the drug product is supplied with another item such as a diluent or adjuvant
	CopackagedIndicator *bool `bson:"copackagedIndicator,omitempty" json:"copackagedIndicator,omitempty"`
	// Manufacturer of this package type (multiple means these are all possible manufacturers)
	Manufacturer []custom.Reference[Organization] `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// Additional information or supporting documentation about the packaged product
	AttachedDocument []custom.Reference[DocumentReference] `bson:"attachedDocument,omitempty" json:"attachedDocument,omitempty"`
	// A packaging item, as a container for medically related items, possibly with other packaging items within, or a packaging component, such as bottle cap
	Packaging *PackagedProductDefinitionPackaging `bson:"packaging,omitempty" json:"packaging,omitempty"`
	// Allows the key features to be recorded, such as "hospital pack", "nurse prescribable"
	Characteristic []interface{} `bson:"characteristic,omitempty" json:"characteristic,omitempty"`
}

type PackagedProductDefinitionLegalStatusOfSupply struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The actual status of supply. In what situation this package type may be supplied for use
	Code *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	// The place where the legal status of supply applies
	Jurisdiction *CodeableConcept `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
}

type PackagedProductDefinitionPackaging struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// An identifier that is specific to this particular part of the packaging. Including possibly a Data Carrier Identifier
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// The physical type of the container of the items
	Type *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	// Is this a part of the packaging (e.g. a cap or bottle stopper), rather than the packaging itself (e.g. a bottle or vial)
	ComponentPart *bool `bson:"componentPart,omitempty" json:"componentPart,omitempty"`
	// The quantity of this level of packaging in the package that contains it (with the outermost level being 1)
	Quantity *int32 `bson:"quantity,omitempty" json:"quantity,omitempty"`
	// Material type of the package item
	Material []CodeableConcept `bson:"material,omitempty" json:"material,omitempty"`
	// A possible alternate material for this part of the packaging, that is allowed to be used instead of the usual material
	AlternateMaterial []CodeableConcept `bson:"alternateMaterial,omitempty" json:"alternateMaterial,omitempty"`
	// Shelf Life and storage information
	ShelfLifeStorage []ProductShelfLife `bson:"shelfLifeStorage,omitempty" json:"shelfLifeStorage,omitempty"`
	// Manufacturer of this packaging item (multiple means these are all potential manufacturers)
	Manufacturer []custom.Reference[Organization] `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	// General characteristics of this item
	Property []PackagedProductDefinitionPackagingProperty `bson:"property,omitempty" json:"property,omitempty"`
	// The item(s) within the packaging
	ContainedItem []PackagedProductDefinitionPackagingContainedItem `bson:"containedItem,omitempty" json:"containedItem,omitempty"`
	// Allows containers (and parts of containers) within containers, still as a part of single packaged product
	Packaging []interface{} `bson:"packaging,omitempty" json:"packaging,omitempty"`
}

type PackagedProductDefinitionPackagingProperty struct {
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
}

type PackagedProductDefinitionPackagingContainedItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The actual item(s) of medication, as manufactured, or a device, or other medically related item (food, biologicals, raw materials, medical fluids, gases etc.), as contained in the package
	Item custom.CodeableReference[PackagedProductDefinitionPackagingContainedItemItem] `bson:"item" json:"item"`
	// The number of this type of item within this packaging or for continuous items such as liquids it is the quantity (for example 25ml). See also PackagedProductDefinition.containedItemQuantity (especially the long definition)
	Amount *Quantity `bson:"amount,omitempty" json:"amount,omitempty"`
}

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
	return "StructureDefinition"
}
