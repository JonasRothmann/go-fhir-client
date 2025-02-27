// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/metric-operational-status
*/type DeviceMetricOperationalStatus string

const (
	// The DeviceMetric is operating and will generate Observations.
	DeviceMetricOperationalStatusOn DeviceMetricOperationalStatus = "on"
	// The DeviceMetric is not operating.
	DeviceMetricOperationalStatusOff DeviceMetricOperationalStatus = "off"
	// The DeviceMetric is operating, but will not generate any Observations.
	DeviceMetricOperationalStatusStandby DeviceMetricOperationalStatus = "standby"
	// The DeviceMetric was entered in error.
	DeviceMetricOperationalStatusEnteredInError DeviceMetricOperationalStatus = "entered-in-error"
)
