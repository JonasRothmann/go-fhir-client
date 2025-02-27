// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/supplydelivery-status
*/type SupplyDeliveryStatus string

const (
	// Supply has been requested, but not delivered.
	SupplyDeliveryStatusInProgress SupplyDeliveryStatus = "in-progress"
	// Supply has been delivered ("completed").
	SupplyDeliveryStatusCompleted SupplyDeliveryStatus = "completed"
	// Delivery was not completed.
	SupplyDeliveryStatusAbandoned SupplyDeliveryStatus = "abandoned"
	// This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be "abandoned" rather than "entered-in-error".).
	SupplyDeliveryStatusEnteredInError SupplyDeliveryStatus = "entered-in-error"
)
