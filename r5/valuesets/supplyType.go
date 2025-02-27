// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://terminology.hl7.org/CodeSystem/supply-kind
*/type SupplyType string

const (
	// Supply is stored and requested from central supply.
	SupplyTypeCentral SupplyType = "central"
	// Supply is not onsite and must be requested from an outside vendor using a non-stock requisition.
	SupplyTypeNonstock SupplyType = "nonstock"
)
