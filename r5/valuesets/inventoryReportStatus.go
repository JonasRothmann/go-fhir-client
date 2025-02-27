// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/inventoryreport-status
*/type InventoryReportStatus string

const (
	// The existence of the report is registered, but it is still without content or only some preliminary content.
	InventoryReportStatusDraft InventoryReportStatus = "draft"
	// The inventory report has been requested but there is no data available.
	InventoryReportStatusRequested InventoryReportStatus = "requested"
	// This report is submitted as current.
	InventoryReportStatusActive InventoryReportStatus = "active"
	// The report has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it.
	InventoryReportStatusEnteredInError InventoryReportStatus = "entered-in-error"
)
