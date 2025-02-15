// GENERATED CODE – DO NOT EDIT!

package v5

import (
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/InventoryReport
type InventoryReport struct {
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
	// Business identifier for the report
	Identifier []Identifier `bson:"identifier,omitempty" json:"identifier,omitempty"`
	// draft | requested | active | entered-in-error
	Status custom.Code `bson:"status" json:"status"`
	// snapshot | difference
	CountType custom.Code `bson:"countType" json:"countType"`
	// addition | subtraction
	OperationType *CodeableConcept `bson:"operationType,omitempty" json:"operationType,omitempty"`
	// The reason for this count - regular count, ad-hoc count, new arrivals, etc
	OperationTypeReason *CodeableConcept `bson:"operationTypeReason,omitempty" json:"operationTypeReason,omitempty"`
	// When the report has been submitted
	ReportedDateTime custom.DateTime `bson:"reportedDateTime" json:"reportedDateTime"`
	// Who submits the report
	Reporter *custom.Reference[InventoryReportReporter] `bson:"reporter,omitempty" json:"reporter,omitempty"`
	// The period the report refers to
	ReportingPeriod *Period `bson:"reportingPeriod,omitempty" json:"reportingPeriod,omitempty"`
	// An inventory listing section (grouped by any of the attributes)
	InventoryListing []InventoryReportInventoryListing `bson:"inventoryListing,omitempty" json:"inventoryListing,omitempty"`
	// A note associated with the InventoryReport
	Note []Annotation `bson:"note,omitempty" json:"note,omitempty"`
}

type InventoryReportInventoryListing struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// Location of the inventory items
	Location *custom.Reference[Location] `bson:"location,omitempty" json:"location,omitempty"`
	// The status of the items that are being reported
	ItemStatus *CodeableConcept `bson:"itemStatus,omitempty" json:"itemStatus,omitempty"`
	// The date and time when the items were counted
	CountingDateTime *custom.DateTime `bson:"countingDateTime,omitempty" json:"countingDateTime,omitempty"`
	// The item or items in this listing
	Item []InventoryReportInventoryListingItem `bson:"item,omitempty" json:"item,omitempty"`
}

type InventoryReportInventoryListingItem struct {
	// Unique id for inter-element referencing
	Id *string `bson:"id,omitempty" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	// The inventory category or classification of the items being reported
	Category *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	// The quantity of the item or items being reported
	Quantity Quantity `bson:"quantity" json:"quantity"`
	// The code or reference to the item type
	Item custom.CodeableReference[InventoryReportInventoryListingItemItem] `bson:"item" json:"item"`
}

type InventoryReportReporter interface {
	gofhirclient.Resource

	Is_InventoryReportReporter()
}

func (p Practitioner) Is_InventoryReportReporter()  {}
func (p Patient) Is_InventoryReportReporter()       {}
func (r RelatedPerson) Is_InventoryReportReporter() {}
func (d Device) Is_InventoryReportReporter()        {}

type InventoryReportInventoryListingItemItem interface {
	gofhirclient.Resource

	Is_InventoryReportInventoryListingItemItem()
}

func (m Medication) Is_InventoryReportInventoryListingItemItem()                 {}
func (d Device) Is_InventoryReportInventoryListingItemItem()                     {}
func (m Medication) Is_InventoryReportInventoryListingItemItem()                 {}
func (n NutritionProduct) Is_InventoryReportInventoryListingItemItem()           {}
func (i InventoryItem) Is_InventoryReportInventoryListingItemItem()              {}
func (b BiologicallyDerivedProduct) Is_InventoryReportInventoryListingItemItem() {}
func (i InventoryItem) Is_InventoryReportInventoryListingItemItem()              {}

func (i InventoryReport) ResourceType() string {
	return "StructureDefinition"
}
