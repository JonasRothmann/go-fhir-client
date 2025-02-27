// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/device-correctiveactionscope
*/type DeviceCorrectiveActionScope string

const (
	// The corrective action was intended for all units of the same model.
	DeviceCorrectiveActionScopeModel DeviceCorrectiveActionScope = "model"
	// The corrective action was intended for a specific batch of units identified by a lot number.
	DeviceCorrectiveActionScopeLotNumbers DeviceCorrectiveActionScope = "lot-numbers"
	// The corrective action was intended for an individual unit (or a set of units) individually identified by serial number.
	DeviceCorrectiveActionScopeSerialNumbers DeviceCorrectiveActionScope = "serial-numbers"
)
