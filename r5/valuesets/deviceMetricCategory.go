// GENERATED CODE – DO NOT EDIT!

package valuesets

/*
Composed of following code systems:
- http://hl7.org/fhir/metric-category
*/type DeviceMetricCategory string

const (
	// Observations generated for this DeviceMetric are measured.
	DeviceMetricCategoryMeasurement DeviceMetricCategory = "measurement"
	// Observations generated for this DeviceMetric is a setting that will influence the behavior of the Device.
	DeviceMetricCategorySetting DeviceMetricCategory = "setting"
	// Observations generated for this DeviceMetric are calculated.
	DeviceMetricCategoryCalculation DeviceMetricCategory = "calculation"
	// The category of this DeviceMetric is unspecified.
	DeviceMetricCategoryUnspecified DeviceMetricCategory = "unspecified"
)
