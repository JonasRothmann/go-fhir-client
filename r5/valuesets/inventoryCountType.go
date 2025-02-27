// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/inventoryreport-counttype
*/type InventoryCountType string

const (
	// The inventory report is a current absolute snapshot, i.e. it represents the quantities at hand.
	InventoryCountTypeSnapshot InventoryCountType = "snapshot"
	// The inventory report is about the difference between a previous count and a current count, i.e. it represents the items that have been added/subtracted from inventory.
	InventoryCountTypeDifference InventoryCountType = "difference"
)
