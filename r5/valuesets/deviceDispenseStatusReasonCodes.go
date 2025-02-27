// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/CodeSystem/devicedispense-status-reason
*/type DeviceDispenseStatusReasonCodes string

const (
	// The device was not dispensed because it was not available.
	DeviceDispenseStatusReasonCodesOutOfStock DeviceDispenseStatusReasonCodes = "out-of-stock"
	// The device was not dispensed because it is of-market - for example not authorized, withdrawn or recalled.
	DeviceDispenseStatusReasonCodesOffMarket DeviceDispenseStatusReasonCodes = "off-market"
	// The device was not dispensed because a contraindication was found - for example pregnancy, allergy to a device component...
	DeviceDispenseStatusReasonCodesContraindication DeviceDispenseStatusReasonCodes = "contraindication"
	// The device was not dispensed because an incompatibility has been found with the device or between the device and other devices being used in the same context.
	DeviceDispenseStatusReasonCodesIncompatibleDevice DeviceDispenseStatusReasonCodes = "incompatible-device"
	// The device was not dispensed because the order has expired or been invalidated.
	DeviceDispenseStatusReasonCodesOrderExpired DeviceDispenseStatusReasonCodes = "order-expired"
	// The device not dispensed because there was a verbal order.
	DeviceDispenseStatusReasonCodesVerbalOrder DeviceDispenseStatusReasonCodes = "verbal-order"
)
