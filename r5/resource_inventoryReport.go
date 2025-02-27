// GENERATED CODE – DO NOT EDIT!

package r5

import (
	"encoding/json"
	gofhirclient "github.com/jonasrothmann/go-fhir-client"
	custom "github.com/jonasrothmann/go-fhir-client/custom"
)

// http://hl7.org/fhir/StructureDefinition/InventoryReport
type InventoryReport struct {
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
	// Business identifier for the report
	Identifier []Identifier `json:"identifier,omitempty"`
	// draft | requested | active | entered-in-error
	Status custom.Code `json:"status"`
	// snapshot | difference
	CountType custom.Code `json:"countType"`
	// addition | subtraction
	OperationType *CodeableConcept `json:"operationType,omitempty"`
	// The reason for this count - regular count, ad-hoc count, new arrivals, etc
	OperationTypeReason *CodeableConcept `json:"operationTypeReason,omitempty"`
	// When the report has been submitted
	ReportedDateTime custom.DateTime `json:"reportedDateTime"`
	// Who submits the report
	Reporter *custom.Reference[InventoryReportReporter] `json:"reporter,omitempty"`
	// The period the report refers to
	ReportingPeriod *Period `json:"reportingPeriod,omitempty"`
	// An inventory listing section (grouped by any of the attributes)
	InventoryListing []InventoryReportInventoryListing `json:"inventoryListing,omitempty"`
	// A note associated with the InventoryReport
	Note []Annotation `json:"note,omitempty"`
}

type InventoryReportInventoryListing struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// Location of the inventory items
	Location *custom.Reference[Location] `json:"location,omitempty"`
	// The status of the items that are being reported
	ItemStatus *CodeableConcept `json:"itemStatus,omitempty"`
	// The date and time when the items were counted
	CountingDateTime *custom.DateTime `json:"countingDateTime,omitempty"`
	// The item or items in this listing
	Item []InventoryReportInventoryListingItem `json:"item,omitempty"`
}

type InventoryReportInventoryListingItem struct {
	// Unique id for inter-element referencing
	Id *string `json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []Extension `json:"extension,omitempty"`
	// Extensions that cannot be ignored even if unrecognized
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	// The inventory category or classification of the items being reported
	Category *CodeableConcept `json:"category,omitempty"`
	// The quantity of the item or items being reported
	Quantity Quantity `json:"quantity"`
	// The code or reference to the item type
	Item custom.CodeableReference[InventoryReportInventoryListingItemItem] `json:"item"`
}

type OtherInventoryReport InventoryReport

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
func (n NutritionProduct) Is_InventoryReportInventoryListingItemItem()           {}
func (i InventoryItem) Is_InventoryReportInventoryListingItemItem()              {}
func (b BiologicallyDerivedProduct) Is_InventoryReportInventoryListingItemItem() {}

func (i InventoryReport) ResourceType() string {
	return "InventoryReport"
}

func (i InventoryReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInventoryReport
		ResourceType string `json:"resourceType,omitempty"`
	}{OtherInventoryReport: OtherInventoryReport(i), ResourceType: i.ResourceType()})
}

func UnmarshalInventoryReport(b []byte) (res InventoryReport, err error) {
	if err := json.Unmarshal(b, &res); err != nil {
		return res, err
	}
	return res, nil
}
